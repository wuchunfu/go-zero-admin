// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package userservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddDeptReq                = sysclient.AddDeptReq
	AddDeptResp               = sysclient.AddDeptResp
	AddDictItemReq            = sysclient.AddDictItemReq
	AddDictItemResp           = sysclient.AddDictItemResp
	AddDictTypeReq            = sysclient.AddDictTypeReq
	AddDictTypeResp           = sysclient.AddDictTypeResp
	AddMenuReq                = sysclient.AddMenuReq
	AddMenuResp               = sysclient.AddMenuResp
	AddOperateLogReq          = sysclient.AddOperateLogReq
	AddOperateLogResp         = sysclient.AddOperateLogResp
	AddPostReq                = sysclient.AddPostReq
	AddPostResp               = sysclient.AddPostResp
	AddRoleReq                = sysclient.AddRoleReq
	AddRoleResp               = sysclient.AddRoleResp
	AddUserReq                = sysclient.AddUserReq
	AddUserResp               = sysclient.AddUserResp
	DeleteDeptReq             = sysclient.DeleteDeptReq
	DeleteDeptResp            = sysclient.DeleteDeptResp
	DeleteDictItemReq         = sysclient.DeleteDictItemReq
	DeleteDictItemResp        = sysclient.DeleteDictItemResp
	DeleteDictTypeReq         = sysclient.DeleteDictTypeReq
	DeleteDictTypeResp        = sysclient.DeleteDictTypeResp
	DeleteLoginLogReq         = sysclient.DeleteLoginLogReq
	DeleteLoginLogResp        = sysclient.DeleteLoginLogResp
	DeleteMenuReq             = sysclient.DeleteMenuReq
	DeleteMenuResp            = sysclient.DeleteMenuResp
	DeleteOperateLogReq       = sysclient.DeleteOperateLogReq
	DeleteOperateLogResp      = sysclient.DeleteOperateLogResp
	DeletePostReq             = sysclient.DeletePostReq
	DeletePostResp            = sysclient.DeletePostResp
	DeleteRoleReq             = sysclient.DeleteRoleReq
	DeleteRoleResp            = sysclient.DeleteRoleResp
	DeleteUserReq             = sysclient.DeleteUserReq
	DeleteUserResp            = sysclient.DeleteUserResp
	DeptListData              = sysclient.DeptListData
	DictItemListData          = sysclient.DictItemListData
	DictTypeListData          = sysclient.DictTypeListData
	InfoReq                   = sysclient.InfoReq
	InfoResp                  = sysclient.InfoResp
	LoginLogListData          = sysclient.LoginLogListData
	LoginReq                  = sysclient.LoginReq
	LoginResp                 = sysclient.LoginResp
	MenuListData              = sysclient.MenuListData
	MenuListTree              = sysclient.MenuListTree
	OperateLogListData        = sysclient.OperateLogListData
	PostListData              = sysclient.PostListData
	QueryDeptAndPostListReq   = sysclient.QueryDeptAndPostListReq
	QueryDeptAndPostListResp  = sysclient.QueryDeptAndPostListResp
	QueryDeptDetailReq        = sysclient.QueryDeptDetailReq
	QueryDeptDetailResp       = sysclient.QueryDeptDetailResp
	QueryDeptListReq          = sysclient.QueryDeptListReq
	QueryDeptListResp         = sysclient.QueryDeptListResp
	QueryDictItemDetailReq    = sysclient.QueryDictItemDetailReq
	QueryDictItemDetailResp   = sysclient.QueryDictItemDetailResp
	QueryDictItemListReq      = sysclient.QueryDictItemListReq
	QueryDictItemListResp     = sysclient.QueryDictItemListResp
	QueryDictTypeDetailReq    = sysclient.QueryDictTypeDetailReq
	QueryDictTypeDetailResp   = sysclient.QueryDictTypeDetailResp
	QueryDictTypeListReq      = sysclient.QueryDictTypeListReq
	QueryDictTypeListResp     = sysclient.QueryDictTypeListResp
	QueryLoginLogDetailReq    = sysclient.QueryLoginLogDetailReq
	QueryLoginLogDetailResp   = sysclient.QueryLoginLogDetailResp
	QueryLoginLogListReq      = sysclient.QueryLoginLogListReq
	QueryLoginLogListResp     = sysclient.QueryLoginLogListResp
	QueryMenuDetailReq        = sysclient.QueryMenuDetailReq
	QueryMenuDetailResp       = sysclient.QueryMenuDetailResp
	QueryMenuListReq          = sysclient.QueryMenuListReq
	QueryMenuListResp         = sysclient.QueryMenuListResp
	QueryOperateLogDetailReq  = sysclient.QueryOperateLogDetailReq
	QueryOperateLogDetailResp = sysclient.QueryOperateLogDetailResp
	QueryOperateLogListReq    = sysclient.QueryOperateLogListReq
	QueryOperateLogListResp   = sysclient.QueryOperateLogListResp
	QueryPostDetailReq        = sysclient.QueryPostDetailReq
	QueryPostDetailResp       = sysclient.QueryPostDetailResp
	QueryPostListReq          = sysclient.QueryPostListReq
	QueryPostListResp         = sysclient.QueryPostListResp
	QueryRoleDetailReq        = sysclient.QueryRoleDetailReq
	QueryRoleDetailResp       = sysclient.QueryRoleDetailResp
	QueryRoleListReq          = sysclient.QueryRoleListReq
	QueryRoleListResp         = sysclient.QueryRoleListResp
	QueryRoleMenuListReq      = sysclient.QueryRoleMenuListReq
	QueryRoleMenuListResp     = sysclient.QueryRoleMenuListResp
	QueryUserDetailReq        = sysclient.QueryUserDetailReq
	QueryUserDetailResp       = sysclient.QueryUserDetailResp
	QueryUserListReq          = sysclient.QueryUserListReq
	QueryUserListResp         = sysclient.QueryUserListResp
	QueryUserRoleListReq      = sysclient.QueryUserRoleListReq
	QueryUserRoleListResp     = sysclient.QueryUserRoleListResp
	ReSetPasswordReq          = sysclient.ReSetPasswordReq
	ReSetPasswordResp         = sysclient.ReSetPasswordResp
	RoleListData              = sysclient.RoleListData
	UpdateDeptReq             = sysclient.UpdateDeptReq
	UpdateDeptResp            = sysclient.UpdateDeptResp
	UpdateDeptStatusReq       = sysclient.UpdateDeptStatusReq
	UpdateDeptStatusResp      = sysclient.UpdateDeptStatusResp
	UpdateDictItemReq         = sysclient.UpdateDictItemReq
	UpdateDictItemResp        = sysclient.UpdateDictItemResp
	UpdateDictItemStatusReq   = sysclient.UpdateDictItemStatusReq
	UpdateDictItemStatusResp  = sysclient.UpdateDictItemStatusResp
	UpdateDictTypeReq         = sysclient.UpdateDictTypeReq
	UpdateDictTypeResp        = sysclient.UpdateDictTypeResp
	UpdateDictTypeStatusReq   = sysclient.UpdateDictTypeStatusReq
	UpdateDictTypeStatusResp  = sysclient.UpdateDictTypeStatusResp
	UpdateMenuReq             = sysclient.UpdateMenuReq
	UpdateMenuResp            = sysclient.UpdateMenuResp
	UpdateMenuRoleReq         = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp        = sysclient.UpdateMenuRoleResp
	UpdateMenuStatusReq       = sysclient.UpdateMenuStatusReq
	UpdateMenuStatusResp      = sysclient.UpdateMenuStatusResp
	UpdatePostReq             = sysclient.UpdatePostReq
	UpdatePostResp            = sysclient.UpdatePostResp
	UpdatePostStatusReq       = sysclient.UpdatePostStatusReq
	UpdatePostStatusResp      = sysclient.UpdatePostStatusResp
	UpdateRoleReq             = sysclient.UpdateRoleReq
	UpdateRoleResp            = sysclient.UpdateRoleResp
	UpdateRoleStatusReq       = sysclient.UpdateRoleStatusReq
	UpdateRoleStatusResp      = sysclient.UpdateRoleStatusResp
	UpdateUserReq             = sysclient.UpdateUserReq
	UpdateUserResp            = sysclient.UpdateUserResp
	UpdateUserRoleListReq     = sysclient.UpdateUserRoleListReq
	UpdateUserRoleListResp    = sysclient.UpdateUserRoleListResp
	UpdateUserStatusReq       = sysclient.UpdateUserStatusReq
	UpdateUserStatusResp      = sysclient.UpdateUserStatusResp
	UserListData              = sysclient.UserListData

	UserService interface {
		// 用户登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 获取用户个人信息
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		// 重置用户密码
		ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error)
		// 添加用户信息表
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
		// 删除用户信息表
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
		// 更新用户信息表
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		// 更新用户信息表状态
		UpdateUserStatus(ctx context.Context, in *UpdateUserStatusReq, opts ...grpc.CallOption) (*UpdateUserStatusResp, error)
		// 查询用户信息表详情
		QueryUserDetail(ctx context.Context, in *QueryUserDetailReq, opts ...grpc.CallOption) (*QueryUserDetailResp, error)
		// 查询用户信息表列表
		QueryUserList(ctx context.Context, in *QueryUserListReq, opts ...grpc.CallOption) (*QueryUserListResp, error)
		// 查询用户与角色的关联
		QueryUserRoleList(ctx context.Context, in *QueryUserRoleListReq, opts ...grpc.CallOption) (*QueryUserRoleListResp, error)
		// 更新用户与角色的关联
		UpdateUserRoleList(ctx context.Context, in *UpdateUserRoleListReq, opts ...grpc.CallOption) (*UpdateUserRoleListResp, error)
		// 查询所有部门和岗位
		QueryDeptAndPostList(ctx context.Context, in *QueryDeptAndPostListReq, opts ...grpc.CallOption) (*QueryDeptAndPostListResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// 用户登录
func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 获取用户个人信息
func (m *defaultUserService) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

// 重置用户密码
func (m *defaultUserService) ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.ReSetPassword(ctx, in, opts...)
}

// 添加用户信息表
func (m *defaultUserService) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

// 删除用户信息表
func (m *defaultUserService) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

// 更新用户信息表
func (m *defaultUserService) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

// 更新用户信息表状态
func (m *defaultUserService) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusReq, opts ...grpc.CallOption) (*UpdateUserStatusResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}

// 查询用户信息表详情
func (m *defaultUserService) QueryUserDetail(ctx context.Context, in *QueryUserDetailReq, opts ...grpc.CallOption) (*QueryUserDetailResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.QueryUserDetail(ctx, in, opts...)
}

// 查询用户信息表列表
func (m *defaultUserService) QueryUserList(ctx context.Context, in *QueryUserListReq, opts ...grpc.CallOption) (*QueryUserListResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.QueryUserList(ctx, in, opts...)
}

// 查询用户与角色的关联
func (m *defaultUserService) QueryUserRoleList(ctx context.Context, in *QueryUserRoleListReq, opts ...grpc.CallOption) (*QueryUserRoleListResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.QueryUserRoleList(ctx, in, opts...)
}

// 更新用户与角色的关联
func (m *defaultUserService) UpdateUserRoleList(ctx context.Context, in *UpdateUserRoleListReq, opts ...grpc.CallOption) (*UpdateUserRoleListResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUserRoleList(ctx, in, opts...)
}

// 查询所有部门和岗位
func (m *defaultUserService) QueryDeptAndPostList(ctx context.Context, in *QueryDeptAndPostListReq, opts ...grpc.CallOption) (*QueryDeptAndPostListResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.QueryDeptAndPostList(ctx, in, opts...)
}
