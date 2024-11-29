// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: captcha/v1/captcha.proto

package captchav1

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

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Biz   string `protobuf:"bytes,1,opt,name=biz,proto3" json:"biz,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_v1_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_v1_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_captcha_v1_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

func (x *SendRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_v1_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_v1_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_captcha_v1_captcha_proto_rawDescGZIP(), []int{1}
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Biz     string `protobuf:"bytes,1,opt,name=biz,proto3" json:"biz,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Captcha string `protobuf:"bytes,3,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_v1_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_v1_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_captcha_v1_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetBiz() string {
	if x != nil {
		return x.Biz
	}
	return ""
}

func (x *VerifyRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *VerifyRequest) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer bool `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captcha_v1_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captcha_v1_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_captcha_v1_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResponse) GetAnswer() bool {
	if x != nil {
		return x.Answer
	}
	return false
}

var File_captcha_v1_captcha_proto protoreflect.FileDescriptor

var file_captcha_v1_captcha_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x0e, 0x0a,
	0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a,
	0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x7a,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x22, 0x28, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0x8c, 0x01, 0x0a, 0x0e, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa0, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x77, 0x6f, 0x6f,
	0x2f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0b, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_captcha_v1_captcha_proto_rawDescOnce sync.Once
	file_captcha_v1_captcha_proto_rawDescData = file_captcha_v1_captcha_proto_rawDesc
)

func file_captcha_v1_captcha_proto_rawDescGZIP() []byte {
	file_captcha_v1_captcha_proto_rawDescOnce.Do(func() {
		file_captcha_v1_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_captcha_v1_captcha_proto_rawDescData)
	})
	return file_captcha_v1_captcha_proto_rawDescData
}

var file_captcha_v1_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_captcha_v1_captcha_proto_goTypes = []interface{}{
	(*SendRequest)(nil),    // 0: captcha.v1.SendRequest
	(*SendResponse)(nil),   // 1: captcha.v1.SendResponse
	(*VerifyRequest)(nil),  // 2: captcha.v1.VerifyRequest
	(*VerifyResponse)(nil), // 3: captcha.v1.VerifyResponse
}
var file_captcha_v1_captcha_proto_depIdxs = []int32{
	0, // 0: captcha.v1.CaptchaService.Send:input_type -> captcha.v1.SendRequest
	2, // 1: captcha.v1.CaptchaService.Verify:input_type -> captcha.v1.VerifyRequest
	1, // 2: captcha.v1.CaptchaService.Send:output_type -> captcha.v1.SendResponse
	3, // 3: captcha.v1.CaptchaService.Verify:output_type -> captcha.v1.VerifyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_captcha_v1_captcha_proto_init() }
func file_captcha_v1_captcha_proto_init() {
	if File_captcha_v1_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_captcha_v1_captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_captcha_v1_captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_captcha_v1_captcha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_captcha_v1_captcha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
			RawDescriptor: file_captcha_v1_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_captcha_v1_captcha_proto_goTypes,
		DependencyIndexes: file_captcha_v1_captcha_proto_depIdxs,
		MessageInfos:      file_captcha_v1_captcha_proto_msgTypes,
	}.Build()
	File_captcha_v1_captcha_proto = out.File
	file_captcha_v1_captcha_proto_rawDesc = nil
	file_captcha_v1_captcha_proto_goTypes = nil
	file_captcha_v1_captcha_proto_depIdxs = nil
}
