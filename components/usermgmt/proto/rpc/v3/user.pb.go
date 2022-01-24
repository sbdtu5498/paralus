// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/rpc/v3/user.proto

package rpcv3

import (
	v3 "github.com/RafaySystems/rcloud-base/components/usermgmt/proto/types/userpb/v3"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_v3_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_v3_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_v3_user_proto_rawDescGZIP(), []int{0}
}

var File_proto_rpc_v3_user_proto protoreflect.FileDescriptor

var file_proto_rpc_v3_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x61, 0x66, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x91, 0x0a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0xdf, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1d,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x92, 0x01,
	0x92, 0x41, 0x36, 0x4a, 0x34, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x69, 0x73, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x53, 0x22,
	0x4e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0xa4, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x21,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x56, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x12, 0x4e, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x8f, 0x02, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0xc5, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xbe, 0x01, 0x12, 0x5d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f,
	0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x5a, 0x5d, 0x12, 0x5b,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x98, 0x02, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0xcb, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0xc4, 0x01, 0x1a, 0x5d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x60, 0x1a, 0x5b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xd2, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xfe, 0x01, 0x92, 0x41, 0x36,
	0x4a, 0x34, 0x0a, 0x03, 0x32, 0x30, 0x34, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73,
	0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xbe, 0x01, 0x2a, 0x5d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f,
	0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x5a, 0x5d, 0x2a, 0x5b,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x7d,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x64, 0x7d, 0x42, 0xca, 0x04, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61,
	0x66, 0x61, 0x79, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x70, 0x63, 0x76, 0x33, 0xa2, 0x02, 0x03,
	0x52, 0x44, 0x52, 0xaa, 0x02, 0x10, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e,
	0x52, 0x70, 0x63, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x10, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44,
	0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x1c, 0x52, 0x61, 0x66, 0x61,
	0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x52, 0x61, 0x66, 0x61, 0x79,
	0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x56, 0x33, 0x92, 0x41,
	0xe2, 0x02, 0x12, 0x2b, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0b, 0x0a,
	0x09, 0x52, 0x61, 0x66, 0x61, 0x79, 0x20, 0x44, 0x65, 0x76, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x2a,
	0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30,
	0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03,
	0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x5a, 0x38, 0x0a, 0x25, 0x0a, 0x0a, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x08, 0x02, 0x1a, 0x11, 0x58,
	0x2d, 0x52, 0x41, 0x46, 0x41, 0x59, 0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b, 0x45, 0x59, 0x49, 0x44,
	0x20, 0x02, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x02, 0x08, 0x01, 0x62, 0x1f, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01,
	0xc0, 0xe3, 0x1e, 0x01, 0xc8, 0xe3, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_v3_user_proto_rawDescOnce sync.Once
	file_proto_rpc_v3_user_proto_rawDescData = file_proto_rpc_v3_user_proto_rawDesc
)

func file_proto_rpc_v3_user_proto_rawDescGZIP() []byte {
	file_proto_rpc_v3_user_proto_rawDescOnce.Do(func() {
		file_proto_rpc_v3_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_v3_user_proto_rawDescData)
	})
	return file_proto_rpc_v3_user_proto_rawDescData
}

var file_proto_rpc_v3_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_rpc_v3_user_proto_goTypes = []interface{}{
	(*DeleteUserResponse)(nil), // 0: rafay.dev.rpc.v3.DeleteUserResponse
	(*v3.User)(nil),            // 1: rafay.dev.types.user.v3.User
	(*v3.UserList)(nil),        // 2: rafay.dev.types.user.v3.UserList
}
var file_proto_rpc_v3_user_proto_depIdxs = []int32{
	1, // 0: rafay.dev.rpc.v3.User.CreateUser:input_type -> rafay.dev.types.user.v3.User
	1, // 1: rafay.dev.rpc.v3.User.GetUsers:input_type -> rafay.dev.types.user.v3.User
	1, // 2: rafay.dev.rpc.v3.User.GetUser:input_type -> rafay.dev.types.user.v3.User
	1, // 3: rafay.dev.rpc.v3.User.UpdateUser:input_type -> rafay.dev.types.user.v3.User
	1, // 4: rafay.dev.rpc.v3.User.DeleteUser:input_type -> rafay.dev.types.user.v3.User
	1, // 5: rafay.dev.rpc.v3.User.CreateUser:output_type -> rafay.dev.types.user.v3.User
	2, // 6: rafay.dev.rpc.v3.User.GetUsers:output_type -> rafay.dev.types.user.v3.UserList
	1, // 7: rafay.dev.rpc.v3.User.GetUser:output_type -> rafay.dev.types.user.v3.User
	1, // 8: rafay.dev.rpc.v3.User.UpdateUser:output_type -> rafay.dev.types.user.v3.User
	0, // 9: rafay.dev.rpc.v3.User.DeleteUser:output_type -> rafay.dev.rpc.v3.DeleteUserResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rpc_v3_user_proto_init() }
func file_proto_rpc_v3_user_proto_init() {
	if File_proto_rpc_v3_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_v3_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
			RawDescriptor: file_proto_rpc_v3_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_v3_user_proto_goTypes,
		DependencyIndexes: file_proto_rpc_v3_user_proto_depIdxs,
		MessageInfos:      file_proto_rpc_v3_user_proto_msgTypes,
	}.Build()
	File_proto_rpc_v3_user_proto = out.File
	file_proto_rpc_v3_user_proto_rawDesc = nil
	file_proto_rpc_v3_user_proto_goTypes = nil
	file_proto_rpc_v3_user_proto_depIdxs = nil
}