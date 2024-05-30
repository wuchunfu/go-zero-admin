package middleware

import (
	"bytes"
	"github.com/feihua/zero-admin/rpc/sys/client/operatelogservice"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/ua-parser/uap-go/uaparser"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
	"net/http"
	"time"
)

type AddLogMiddleware struct {
	Sys operatelogservice.OperateLogService
}

func NewAddLogMiddleware(Sys operatelogservice.OperateLogService) *AddLogMiddleware {
	return &AddLogMiddleware{Sys: Sys}
}

// Handle 参考chatgpt实现
func (m *AddLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uri := r.RequestURI
		if uri == "/api/sys/user/login" || uri == "/api/sys/upload" {
			logx.WithContext(r.Context()).Infof("Request: %s %s", r.Method, uri)
			next(w, r)
			return
		}

		userName := r.Context().Value("userName").(string)
		deptName := r.Context().Value("deptName").(string)

		startTime := time.Now()

		// 读取请求主体
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("Failed to read request body: %v", err)
		}

		// 创建一个新的请求主体用于后续读取
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 打印请求参数日志
		logx.WithContext(r.Context()).Infof("Request: %s %s %s", r.Method, uri, body)

		// 创建一个自定义的 ResponseWriter，用于记录响应
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}

		// 调用下一个处理器，捕获响应
		next(recorder, r)

		// 打印响应日志
		responseBoy := string(recorder.body)
		// 响应参数较多,可以不打印
		//logx.WithContext(r.Context()).Infof("Response: %s %s %s", r.Method, r.RequestURI, responseBoy)

		userAgent := r.Header.Get("User-Agent")
		parser := uaparser.NewFromSaved()
		ua := parser.Parse(userAgent)

		browser := ua.UserAgent.Family + " " + ua.UserAgent.Major
		os := ua.Os.Family + " " + ua.Os.Major
		// 打印请求和响应耗时
		duration := time.Since(startTime)
		_, _ = m.Sys.AddOperateLog(r.Context(), &sysclient.AddOperateLogReq{
			DeptName:          deptName,
			OperationIp:       httpx.GetRemoteAddr(r),
			OperationName:     userName,
			OperationParams:   string(body),
			OperationResponse: responseBoy,
			UseTime:           duration.Milliseconds(),
			OperationType:     "",
			OperationUrl:      uri,
			RequestMethod:     r.Method,
			Title:             "",
			Os:                os,
			Browser:           browser,
		})
	}
}

// 自定义的 ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader 重写 WriteHeader 方法，捕获状态码
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// 重写 Write 方法，捕获响应数据
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}
