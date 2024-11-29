// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: payment/v1/payment.proto

package paymentv1

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

type PaymentStatus int32

const (
	PaymentStatus_PAYMENT_STATUS_UNSPECIFIED PaymentStatus = 0
	PaymentStatus_PAYMENT_STATUS_INIT        PaymentStatus = 1
	PaymentStatus_PAYMENT_STATUS_SUCCESS     PaymentStatus = 2
	PaymentStatus_PAYMENT_STATUS_FAILED      PaymentStatus = 3
	PaymentStatus_PAYMENT_STATUS_REFUND      PaymentStatus = 4
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "PAYMENT_STATUS_UNSPECIFIED",
		1: "PAYMENT_STATUS_INIT",
		2: "PAYMENT_STATUS_SUCCESS",
		3: "PAYMENT_STATUS_FAILED",
		4: "PAYMENT_STATUS_REFUND",
	}
	PaymentStatus_value = map[string]int32{
		"PAYMENT_STATUS_UNSPECIFIED": 0,
		"PAYMENT_STATUS_INIT":        1,
		"PAYMENT_STATUS_SUCCESS":     2,
		"PAYMENT_STATUS_FAILED":      3,
		"PAYMENT_STATUS_REFUND":      4,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_v1_payment_proto_enumTypes[0].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_payment_v1_payment_proto_enumTypes[0]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

type GetPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizTradeNo string `protobuf:"bytes,1,opt,name=biz_trade_no,json=bizTradeNo,proto3" json:"biz_trade_no,omitempty"`
}

func (x *GetPaymentRequest) Reset() {
	*x = GetPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentRequest) ProtoMessage() {}

func (x *GetPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *GetPaymentRequest) GetBizTradeNo() string {
	if x != nil {
		return x.BizTradeNo
	}
	return ""
}

type GetPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status PaymentStatus `protobuf:"varint,1,opt,name=status,proto3,enum=payment.v1.PaymentStatus" json:"status,omitempty"`
}

func (x *GetPaymentResponse) Reset() {
	*x = GetPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentResponse) ProtoMessage() {}

func (x *GetPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{1}
}

func (x *GetPaymentResponse) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_PAYMENT_STATUS_UNSPECIFIED
}

type NativePrePayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      *Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BizTradeNo  string  `protobuf:"bytes,2,opt,name=biz_trade_no,json=bizTradeNo,proto3" json:"biz_trade_no,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *NativePrePayRequest) Reset() {
	*x = NativePrePayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NativePrePayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NativePrePayRequest) ProtoMessage() {}

func (x *NativePrePayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NativePrePayRequest.ProtoReflect.Descriptor instead.
func (*NativePrePayRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{2}
}

func (x *NativePrePayRequest) GetAmount() *Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *NativePrePayRequest) GetBizTradeNo() string {
	if x != nil {
		return x.BizTradeNo
	}
	return ""
}

func (x *NativePrePayRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{3}
}

func (x *Amount) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Amount) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type NativePrePayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeUrl string `protobuf:"bytes,1,opt,name=code_url,json=codeUrl,proto3" json:"code_url,omitempty"`
}

func (x *NativePrePayResponse) Reset() {
	*x = NativePrePayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NativePrePayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NativePrePayResponse) ProtoMessage() {}

func (x *NativePrePayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NativePrePayResponse.ProtoReflect.Descriptor instead.
func (*NativePrePayResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{4}
}

func (x *NativePrePayResponse) GetCodeUrl() string {
	if x != nil {
		return x.CodeUrl
	}
	return ""
}

type RefundPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizTradeNo   string `protobuf:"bytes,1,opt,name=biz_trade_no,json=bizTradeNo,proto3" json:"biz_trade_no,omitempty"`
	RefundReason string `protobuf:"bytes,2,opt,name=refund_reason,json=refundReason,proto3" json:"refund_reason,omitempty"`
}

func (x *RefundPaymentRequest) Reset() {
	*x = RefundPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundPaymentRequest) ProtoMessage() {}

func (x *RefundPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundPaymentRequest.ProtoReflect.Descriptor instead.
func (*RefundPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{5}
}

func (x *RefundPaymentRequest) GetBizTradeNo() string {
	if x != nil {
		return x.BizTradeNo
	}
	return ""
}

func (x *RefundPaymentRequest) GetRefundReason() string {
	if x != nil {
		return x.RefundReason
	}
	return ""
}

type RefundPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefundPaymentResponse) Reset() {
	*x = RefundPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_v1_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundPaymentResponse) ProtoMessage() {}

func (x *RefundPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundPaymentResponse.ProtoReflect.Descriptor instead.
func (*RefundPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{6}
}

var File_payment_v1_payment_proto protoreflect.FileDescriptor

var file_payment_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62,
	0x69, 0x7a, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x69, 0x7a, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22, 0x47, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x4e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x72, 0x65, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x69,
	0x7a, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x69, 0x7a, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x31, 0x0a, 0x14, 0x4e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x65, 0x50, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x64, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x5d, 0x0a,
	0x14, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x69, 0x7a, 0x5f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x7a,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x9a, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15,
	0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44,
	0x10, 0x04, 0x32, 0x8c, 0x02, 0x0a, 0x14, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x4e,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x65, 0x50, 0x61, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50,
	0x72, 0x65, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x50, 0x72, 0x65, 0x50, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xa0, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x77, 0x6f, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_v1_payment_proto_rawDescOnce sync.Once
	file_payment_v1_payment_proto_rawDescData = file_payment_v1_payment_proto_rawDesc
)

func file_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_v1_payment_proto_rawDescData)
	})
	return file_payment_v1_payment_proto_rawDescData
}

var file_payment_v1_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_payment_v1_payment_proto_goTypes = []interface{}{
	(PaymentStatus)(0),            // 0: payment.v1.PaymentStatus
	(*GetPaymentRequest)(nil),     // 1: payment.v1.GetPaymentRequest
	(*GetPaymentResponse)(nil),    // 2: payment.v1.GetPaymentResponse
	(*NativePrePayRequest)(nil),   // 3: payment.v1.NativePrePayRequest
	(*Amount)(nil),                // 4: payment.v1.Amount
	(*NativePrePayResponse)(nil),  // 5: payment.v1.NativePrePayResponse
	(*RefundPaymentRequest)(nil),  // 6: payment.v1.RefundPaymentRequest
	(*RefundPaymentResponse)(nil), // 7: payment.v1.RefundPaymentResponse
}
var file_payment_v1_payment_proto_depIdxs = []int32{
	0, // 0: payment.v1.GetPaymentResponse.status:type_name -> payment.v1.PaymentStatus
	4, // 1: payment.v1.NativePrePayRequest.amount:type_name -> payment.v1.Amount
	3, // 2: payment.v1.WechatPaymentService.NativePrePay:input_type -> payment.v1.NativePrePayRequest
	1, // 3: payment.v1.WechatPaymentService.GetPayment:input_type -> payment.v1.GetPaymentRequest
	6, // 4: payment.v1.WechatPaymentService.RefundPayment:input_type -> payment.v1.RefundPaymentRequest
	5, // 5: payment.v1.WechatPaymentService.NativePrePay:output_type -> payment.v1.NativePrePayResponse
	2, // 6: payment.v1.WechatPaymentService.GetPayment:output_type -> payment.v1.GetPaymentResponse
	7, // 7: payment.v1.WechatPaymentService.RefundPayment:output_type -> payment.v1.RefundPaymentResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_v1_payment_proto_init() }
func file_payment_v1_payment_proto_init() {
	if File_payment_v1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_v1_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentRequest); i {
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
		file_payment_v1_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentResponse); i {
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
		file_payment_v1_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NativePrePayRequest); i {
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
		file_payment_v1_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_payment_v1_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NativePrePayResponse); i {
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
		file_payment_v1_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundPaymentRequest); i {
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
		file_payment_v1_payment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundPaymentResponse); i {
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
			RawDescriptor: file_payment_v1_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_v1_payment_proto_goTypes,
		DependencyIndexes: file_payment_v1_payment_proto_depIdxs,
		EnumInfos:         file_payment_v1_payment_proto_enumTypes,
		MessageInfos:      file_payment_v1_payment_proto_msgTypes,
	}.Build()
	File_payment_v1_payment_proto = out.File
	file_payment_v1_payment_proto_rawDesc = nil
	file_payment_v1_payment_proto_goTypes = nil
	file_payment_v1_payment_proto_depIdxs = nil
}
