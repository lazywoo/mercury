// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: follow/v1/follow.proto

package followv1

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

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Follower int64 `protobuf:"varint,2,opt,name=follower,proto3" json:"follower,omitempty"`
	Followee int64 `protobuf:"varint,3,opt,name=followee,proto3" json:"followee,omitempty"`
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{0}
}

func (x *Relation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Relation) GetFollower() int64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

func (x *Relation) GetFollowee() int64 {
	if x != nil {
		return x.Followee
	}
	return 0
}

type Statics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerCount int64 `protobuf:"varint,1,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	FolloweeCount int64 `protobuf:"varint,2,opt,name=followee_count,json=followeeCount,proto3" json:"followee_count,omitempty"`
}

func (x *Statics) Reset() {
	*x = Statics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statics) ProtoMessage() {}

func (x *Statics) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statics.ProtoReflect.Descriptor instead.
func (*Statics) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{1}
}

func (x *Statics) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *Statics) GetFolloweeCount() int64 {
	if x != nil {
		return x.FolloweeCount
	}
	return 0
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followee int64 `protobuf:"varint,1,opt,name=followee,proto3" json:"followee,omitempty"`
	Follower int64 `protobuf:"varint,2,opt,name=follower,proto3" json:"follower,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{2}
}

func (x *FollowRequest) GetFollowee() int64 {
	if x != nil {
		return x.Followee
	}
	return 0
}

func (x *FollowRequest) GetFollower() int64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{3}
}

type CancelFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followee int64 `protobuf:"varint,1,opt,name=followee,proto3" json:"followee,omitempty"`
	Follower int64 `protobuf:"varint,2,opt,name=follower,proto3" json:"follower,omitempty"`
}

func (x *CancelFollowRequest) Reset() {
	*x = CancelFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelFollowRequest) ProtoMessage() {}

func (x *CancelFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelFollowRequest.ProtoReflect.Descriptor instead.
func (*CancelFollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{4}
}

func (x *CancelFollowRequest) GetFollowee() int64 {
	if x != nil {
		return x.Followee
	}
	return 0
}

func (x *CancelFollowRequest) GetFollower() int64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

type CancelFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelFollowResponse) Reset() {
	*x = CancelFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelFollowResponse) ProtoMessage() {}

func (x *CancelFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelFollowResponse.ProtoReflect.Descriptor instead.
func (*CancelFollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{5}
}

type GetFolloweeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follower int64 `protobuf:"varint,1,opt,name=follower,proto3" json:"follower,omitempty"`
	Offset   int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetFolloweeRequest) Reset() {
	*x = GetFolloweeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFolloweeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFolloweeRequest) ProtoMessage() {}

func (x *GetFolloweeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFolloweeRequest.ProtoReflect.Descriptor instead.
func (*GetFolloweeRequest) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{6}
}

func (x *GetFolloweeRequest) GetFollower() int64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

func (x *GetFolloweeRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetFolloweeRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFolloweeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowRelation []*Relation `protobuf:"bytes,1,rep,name=follow_relation,json=followRelation,proto3" json:"follow_relation,omitempty"`
}

func (x *GetFolloweeResponse) Reset() {
	*x = GetFolloweeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFolloweeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFolloweeResponse) ProtoMessage() {}

func (x *GetFolloweeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFolloweeResponse.ProtoReflect.Descriptor instead.
func (*GetFolloweeResponse) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{7}
}

func (x *GetFolloweeResponse) GetFollowRelation() []*Relation {
	if x != nil {
		return x.FollowRelation
	}
	return nil
}

type GetFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followee int64 `protobuf:"varint,1,opt,name=followee,proto3" json:"followee,omitempty"`
	Offset   int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetFollowerRequest) Reset() {
	*x = GetFollowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerRequest) ProtoMessage() {}

func (x *GetFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerRequest.ProtoReflect.Descriptor instead.
func (*GetFollowerRequest) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{8}
}

func (x *GetFollowerRequest) GetFollowee() int64 {
	if x != nil {
		return x.Followee
	}
	return 0
}

func (x *GetFollowerRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetFollowerRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFollowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowRelation []*Relation `protobuf:"bytes,1,rep,name=follow_relation,json=followRelation,proto3" json:"follow_relation,omitempty"`
}

func (x *GetFollowerResponse) Reset() {
	*x = GetFollowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowerResponse) ProtoMessage() {}

func (x *GetFollowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowerResponse.ProtoReflect.Descriptor instead.
func (*GetFollowerResponse) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{9}
}

func (x *GetFollowerResponse) GetFollowRelation() []*Relation {
	if x != nil {
		return x.FollowRelation
	}
	return nil
}

type GetRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followee int64 `protobuf:"varint,1,opt,name=followee,proto3" json:"followee,omitempty"`
	Follower int64 `protobuf:"varint,2,opt,name=follower,proto3" json:"follower,omitempty"`
}

func (x *GetRelationRequest) Reset() {
	*x = GetRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationRequest) ProtoMessage() {}

func (x *GetRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationRequest.ProtoReflect.Descriptor instead.
func (*GetRelationRequest) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{10}
}

func (x *GetRelationRequest) GetFollowee() int64 {
	if x != nil {
		return x.Followee
	}
	return 0
}

func (x *GetRelationRequest) GetFollower() int64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

type GetRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowRelation *Relation `protobuf:"bytes,1,opt,name=follow_relation,json=followRelation,proto3" json:"follow_relation,omitempty"`
}

func (x *GetRelationResponse) Reset() {
	*x = GetRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationResponse) ProtoMessage() {}

func (x *GetRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationResponse.ProtoReflect.Descriptor instead.
func (*GetRelationResponse) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{11}
}

func (x *GetRelationResponse) GetFollowRelation() *Relation {
	if x != nil {
		return x.FollowRelation
	}
	return nil
}

type GetStaticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetStaticsRequest) Reset() {
	*x = GetStaticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsRequest) ProtoMessage() {}

func (x *GetStaticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsRequest.ProtoReflect.Descriptor instead.
func (*GetStaticsRequest) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{12}
}

func (x *GetStaticsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetStaticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statics *Statics `protobuf:"bytes,1,opt,name=statics,proto3" json:"statics,omitempty"`
}

func (x *GetStaticsResponse) Reset() {
	*x = GetStaticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_v1_follow_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticsResponse) ProtoMessage() {}

func (x *GetStaticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_v1_follow_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticsResponse.ProtoReflect.Descriptor instead.
func (*GetStaticsResponse) Descriptor() ([]byte, []int) {
	return file_follow_v1_follow_proto_rawDescGZIP(), []int{13}
}

func (x *GetStaticsResponse) GetStatics() *Statics {
	if x != nil {
		return x.Statics
	}
	return nil
}

var File_follow_v1_follow_proto protoreflect.FileDescriptor

var file_follow_v1_follow_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x22, 0x52, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x22, 0x57, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x47, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x13, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x32, 0xd4, 0x03, 0x0a, 0x0d, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1e, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x12, 0x1d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x98, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x61, 0x7a, 0x79, 0x77, 0x6f, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31,
	0x3b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follow_v1_follow_proto_rawDescOnce sync.Once
	file_follow_v1_follow_proto_rawDescData = file_follow_v1_follow_proto_rawDesc
)

func file_follow_v1_follow_proto_rawDescGZIP() []byte {
	file_follow_v1_follow_proto_rawDescOnce.Do(func() {
		file_follow_v1_follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_follow_v1_follow_proto_rawDescData)
	})
	return file_follow_v1_follow_proto_rawDescData
}

var file_follow_v1_follow_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_follow_v1_follow_proto_goTypes = []interface{}{
	(*Relation)(nil),             // 0: follow.v1.Relation
	(*Statics)(nil),              // 1: follow.v1.Statics
	(*FollowRequest)(nil),        // 2: follow.v1.FollowRequest
	(*FollowResponse)(nil),       // 3: follow.v1.FollowResponse
	(*CancelFollowRequest)(nil),  // 4: follow.v1.CancelFollowRequest
	(*CancelFollowResponse)(nil), // 5: follow.v1.CancelFollowResponse
	(*GetFolloweeRequest)(nil),   // 6: follow.v1.GetFolloweeRequest
	(*GetFolloweeResponse)(nil),  // 7: follow.v1.GetFolloweeResponse
	(*GetFollowerRequest)(nil),   // 8: follow.v1.GetFollowerRequest
	(*GetFollowerResponse)(nil),  // 9: follow.v1.GetFollowerResponse
	(*GetRelationRequest)(nil),   // 10: follow.v1.GetRelationRequest
	(*GetRelationResponse)(nil),  // 11: follow.v1.GetRelationResponse
	(*GetStaticsRequest)(nil),    // 12: follow.v1.GetStaticsRequest
	(*GetStaticsResponse)(nil),   // 13: follow.v1.GetStaticsResponse
}
var file_follow_v1_follow_proto_depIdxs = []int32{
	0,  // 0: follow.v1.GetFolloweeResponse.follow_relation:type_name -> follow.v1.Relation
	0,  // 1: follow.v1.GetFollowerResponse.follow_relation:type_name -> follow.v1.Relation
	0,  // 2: follow.v1.GetRelationResponse.follow_relation:type_name -> follow.v1.Relation
	1,  // 3: follow.v1.GetStaticsResponse.statics:type_name -> follow.v1.Statics
	2,  // 4: follow.v1.FollowService.Follow:input_type -> follow.v1.FollowRequest
	4,  // 5: follow.v1.FollowService.CancelFollow:input_type -> follow.v1.CancelFollowRequest
	6,  // 6: follow.v1.FollowService.GetFollowee:input_type -> follow.v1.GetFolloweeRequest
	8,  // 7: follow.v1.FollowService.GetFollower:input_type -> follow.v1.GetFollowerRequest
	10, // 8: follow.v1.FollowService.GetRelation:input_type -> follow.v1.GetRelationRequest
	12, // 9: follow.v1.FollowService.GetStatics:input_type -> follow.v1.GetStaticsRequest
	3,  // 10: follow.v1.FollowService.Follow:output_type -> follow.v1.FollowResponse
	5,  // 11: follow.v1.FollowService.CancelFollow:output_type -> follow.v1.CancelFollowResponse
	7,  // 12: follow.v1.FollowService.GetFollowee:output_type -> follow.v1.GetFolloweeResponse
	9,  // 13: follow.v1.FollowService.GetFollower:output_type -> follow.v1.GetFollowerResponse
	11, // 14: follow.v1.FollowService.GetRelation:output_type -> follow.v1.GetRelationResponse
	13, // 15: follow.v1.FollowService.GetStatics:output_type -> follow.v1.GetStaticsResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_follow_v1_follow_proto_init() }
func file_follow_v1_follow_proto_init() {
	if File_follow_v1_follow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follow_v1_follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
		file_follow_v1_follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statics); i {
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
		file_follow_v1_follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_follow_v1_follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_follow_v1_follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelFollowRequest); i {
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
		file_follow_v1_follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelFollowResponse); i {
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
		file_follow_v1_follow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFolloweeRequest); i {
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
		file_follow_v1_follow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFolloweeResponse); i {
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
		file_follow_v1_follow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerRequest); i {
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
		file_follow_v1_follow_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowerResponse); i {
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
		file_follow_v1_follow_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationRequest); i {
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
		file_follow_v1_follow_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationResponse); i {
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
		file_follow_v1_follow_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticsRequest); i {
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
		file_follow_v1_follow_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticsResponse); i {
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
			RawDescriptor: file_follow_v1_follow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follow_v1_follow_proto_goTypes,
		DependencyIndexes: file_follow_v1_follow_proto_depIdxs,
		MessageInfos:      file_follow_v1_follow_proto_msgTypes,
	}.Build()
	File_follow_v1_follow_proto = out.File
	file_follow_v1_follow_proto_rawDesc = nil
	file_follow_v1_follow_proto_goTypes = nil
	file_follow_v1_follow_proto_depIdxs = nil
}
