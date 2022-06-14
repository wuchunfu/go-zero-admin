// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: pay.proto

package payclient

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnifiedOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string `protobuf:"bytes,1,opt,name=BusinessId,proto3" json:"BusinessId,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Remarks    string `protobuf:"bytes,3,opt,name=Remarks,proto3" json:"Remarks,omitempty"`
	Code       string `protobuf:"bytes,4,opt,name=Code,proto3" json:"Code,omitempty"`
	MerId      string `protobuf:"bytes,5,opt,name=MerId,proto3" json:"MerId,omitempty"`
	PayType    string `protobuf:"bytes,6,opt,name=PayType,proto3" json:"PayType,omitempty"`
}

func (x *UnifiedOrderReq) Reset() {
	*x = UnifiedOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnifiedOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnifiedOrderReq) ProtoMessage() {}

func (x *UnifiedOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnifiedOrderReq.ProtoReflect.Descriptor instead.
func (*UnifiedOrderReq) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{0}
}

func (x *UnifiedOrderReq) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *UnifiedOrderReq) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *UnifiedOrderReq) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *UnifiedOrderReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UnifiedOrderReq) GetMerId() string {
	if x != nil {
		return x.MerId
	}
	return ""
}

func (x *UnifiedOrderReq) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

type UnifiedOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      string `protobuf:"bytes,1,opt,name=AppId,proto3" json:"AppId,omitempty"`
	PartnerId  string `protobuf:"bytes,2,opt,name=PartnerId,proto3" json:"PartnerId,omitempty"`
	PrepayId   string `protobuf:"bytes,3,opt,name=prepayId,proto3" json:"prepayId,omitempty"`
	PackageStr string `protobuf:"bytes,4,opt,name=PackageStr,proto3" json:"PackageStr,omitempty"`
	NonceStr   string `protobuf:"bytes,5,opt,name=NonceStr,proto3" json:"NonceStr,omitempty"`
	Timestamp  string `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Sign       string `protobuf:"bytes,7,opt,name=Sign,proto3" json:"Sign,omitempty"`
	MWebUrl    string `protobuf:"bytes,8,opt,name=MWebUrl,proto3" json:"MWebUrl,omitempty"`
}

func (x *UnifiedOrderResp) Reset() {
	*x = UnifiedOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnifiedOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnifiedOrderResp) ProtoMessage() {}

func (x *UnifiedOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnifiedOrderResp.ProtoReflect.Descriptor instead.
func (*UnifiedOrderResp) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{1}
}

func (x *UnifiedOrderResp) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *UnifiedOrderResp) GetPartnerId() string {
	if x != nil {
		return x.PartnerId
	}
	return ""
}

func (x *UnifiedOrderResp) GetPrepayId() string {
	if x != nil {
		return x.PrepayId
	}
	return ""
}

func (x *UnifiedOrderResp) GetPackageStr() string {
	if x != nil {
		return x.PackageStr
	}
	return ""
}

func (x *UnifiedOrderResp) GetNonceStr() string {
	if x != nil {
		return x.NonceStr
	}
	return ""
}

func (x *UnifiedOrderResp) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *UnifiedOrderResp) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *UnifiedOrderResp) GetMWebUrl() string {
	if x != nil {
		return x.MWebUrl
	}
	return ""
}

type H5UnifiedOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MWebUrl string `protobuf:"bytes,1,opt,name=MWebUrl,proto3" json:"MWebUrl,omitempty"`
}

func (x *H5UnifiedOrderResp) Reset() {
	*x = H5UnifiedOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H5UnifiedOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H5UnifiedOrderResp) ProtoMessage() {}

func (x *H5UnifiedOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H5UnifiedOrderResp.ProtoReflect.Descriptor instead.
func (*H5UnifiedOrderResp) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{2}
}

func (x *H5UnifiedOrderResp) GetMWebUrl() string {
	if x != nil {
		return x.MWebUrl
	}
	return ""
}

type OrderQueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string `protobuf:"bytes,1,opt,name=BusinessId,proto3" json:"BusinessId,omitempty"`
	MerId      string `protobuf:"bytes,2,opt,name=MerId,proto3" json:"MerId,omitempty"`
	PayType    string `protobuf:"bytes,6,opt,name=PayType,proto3" json:"PayType,omitempty"`
}

func (x *OrderQueryReq) Reset() {
	*x = OrderQueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderQueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderQueryReq) ProtoMessage() {}

func (x *OrderQueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderQueryReq.ProtoReflect.Descriptor instead.
func (*OrderQueryReq) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{3}
}

func (x *OrderQueryReq) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *OrderQueryReq) GetMerId() string {
	if x != nil {
		return x.MerId
	}
	return ""
}

func (x *OrderQueryReq) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

type OrderQueryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayMessage string `protobuf:"bytes,1,opt,name=PayMessage,proto3" json:"PayMessage,omitempty"`
	PayStatus  string `protobuf:"bytes,2,opt,name=PayStatus,proto3" json:"PayStatus,omitempty"`
}

func (x *OrderQueryResp) Reset() {
	*x = OrderQueryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderQueryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderQueryResp) ProtoMessage() {}

func (x *OrderQueryResp) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderQueryResp.ProtoReflect.Descriptor instead.
func (*OrderQueryResp) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{4}
}

func (x *OrderQueryResp) GetPayMessage() string {
	if x != nil {
		return x.PayMessage
	}
	return ""
}

func (x *OrderQueryResp) GetPayStatus() string {
	if x != nil {
		return x.PayStatus
	}
	return ""
}

var File_pay_proto protoreflect.FileDescriptor

var file_pay_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x61, 0x79,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x0f, 0x55, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xea, 0x01, 0x0a, 0x10, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65,
	0x70, 0x61, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65,
	0x70, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x53, 0x74,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x53, 0x74,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53,
	0x69, 0x67, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x57, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x57, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x22, 0x2e, 0x0a,
	0x12, 0x48, 0x35, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x57, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x57, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x22, 0x5f, 0x0a,
	0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x4d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4e,
	0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xac,
	0x02, 0x0a, 0x03, 0x50, 0x61, 0x79, 0x12, 0x4a, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x55, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0e, 0x48, 0x35, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x35, 0x55,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x49, 0x0a, 0x0e, 0x4a, 0x73, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0a, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2f, 0x70, 0x61, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pay_proto_rawDescOnce sync.Once
	file_pay_proto_rawDescData = file_pay_proto_rawDesc
)

func file_pay_proto_rawDescGZIP() []byte {
	file_pay_proto_rawDescOnce.Do(func() {
		file_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_pay_proto_rawDescData)
	})
	return file_pay_proto_rawDescData
}

var file_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pay_proto_goTypes = []interface{}{
	(*UnifiedOrderReq)(nil),    // 0: payclient.UnifiedOrderReq
	(*UnifiedOrderResp)(nil),   // 1: payclient.UnifiedOrderResp
	(*H5UnifiedOrderResp)(nil), // 2: payclient.H5UnifiedOrderResp
	(*OrderQueryReq)(nil),      // 3: payclient.OrderQueryReq
	(*OrderQueryResp)(nil),     // 4: payclient.OrderQueryResp
}
var file_pay_proto_depIdxs = []int32{
	0, // 0: payclient.Pay.AppUnifiedOrder:input_type -> payclient.UnifiedOrderReq
	0, // 1: payclient.Pay.H5UnifiedOrder:input_type -> payclient.UnifiedOrderReq
	0, // 2: payclient.Pay.JsUnifiedOrder:input_type -> payclient.UnifiedOrderReq
	3, // 3: payclient.Pay.OrderQuery:input_type -> payclient.OrderQueryReq
	1, // 4: payclient.Pay.AppUnifiedOrder:output_type -> payclient.UnifiedOrderResp
	2, // 5: payclient.Pay.H5UnifiedOrder:output_type -> payclient.H5UnifiedOrderResp
	1, // 6: payclient.Pay.JsUnifiedOrder:output_type -> payclient.UnifiedOrderResp
	4, // 7: payclient.Pay.OrderQuery:output_type -> payclient.OrderQueryResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pay_proto_init() }
func file_pay_proto_init() {
	if File_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnifiedOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnifiedOrderResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H5UnifiedOrderResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderQueryReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderQueryResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pay_proto_goTypes,
		DependencyIndexes: file_pay_proto_depIdxs,
		MessageInfos:      file_pay_proto_msgTypes,
	}.Build()
	File_pay_proto = out.File
	file_pay_proto_rawDesc = nil
	file_pay_proto_goTypes = nil
	file_pay_proto_depIdxs = nil
}