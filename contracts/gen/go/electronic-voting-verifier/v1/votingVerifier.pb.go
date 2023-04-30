// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: votingVerifier.proto

package v1

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

type GetPublicKeyForVotingIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VotingID int32 `protobuf:"varint,1,opt,name=votingID,proto3" json:"votingID,omitempty"`
}

func (x *GetPublicKeyForVotingIDRequest) Reset() {
	*x = GetPublicKeyForVotingIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_votingVerifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicKeyForVotingIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicKeyForVotingIDRequest) ProtoMessage() {}

func (x *GetPublicKeyForVotingIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_votingVerifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicKeyForVotingIDRequest.ProtoReflect.Descriptor instead.
func (*GetPublicKeyForVotingIDRequest) Descriptor() ([]byte, []int) {
	return file_votingVerifier_proto_rawDescGZIP(), []int{0}
}

func (x *GetPublicKeyForVotingIDRequest) GetVotingID() int32 {
	if x != nil {
		return x.VotingID
	}
	return 0
}

type GetPublicKeyForVotingIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKeyBytes []byte `protobuf:"bytes,1,opt,name=publicKeyBytes,proto3" json:"publicKeyBytes,omitempty"`
}

func (x *GetPublicKeyForVotingIDResponse) Reset() {
	*x = GetPublicKeyForVotingIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_votingVerifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicKeyForVotingIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicKeyForVotingIDResponse) ProtoMessage() {}

func (x *GetPublicKeyForVotingIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_votingVerifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicKeyForVotingIDResponse.ProtoReflect.Descriptor instead.
func (*GetPublicKeyForVotingIDResponse) Descriptor() ([]byte, []int) {
	return file_votingVerifier_proto_rawDescGZIP(), []int{1}
}

func (x *GetPublicKeyForVotingIDResponse) GetPublicKeyBytes() []byte {
	if x != nil {
		return x.PublicKeyBytes
	}
	return nil
}

var File_votingVerifier_proto protoreflect.FileDescriptor

var file_votingVerifier_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x32,
	0x6e, 0x0a, 0x0e, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x5c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61,
	0x6b, 0x6f, 0x73, 0x61, 0x72, 0x65, 0x76, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e,
	0x69, 0x63, 0x2d, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x76, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_votingVerifier_proto_rawDescOnce sync.Once
	file_votingVerifier_proto_rawDescData = file_votingVerifier_proto_rawDesc
)

func file_votingVerifier_proto_rawDescGZIP() []byte {
	file_votingVerifier_proto_rawDescOnce.Do(func() {
		file_votingVerifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_votingVerifier_proto_rawDescData)
	})
	return file_votingVerifier_proto_rawDescData
}

var file_votingVerifier_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_votingVerifier_proto_goTypes = []interface{}{
	(*GetPublicKeyForVotingIDRequest)(nil),  // 0: GetPublicKeyForVotingIDRequest
	(*GetPublicKeyForVotingIDResponse)(nil), // 1: GetPublicKeyForVotingIDResponse
}
var file_votingVerifier_proto_depIdxs = []int32{
	0, // 0: VotingVerifier.GetPublicKeyForVotingID:input_type -> GetPublicKeyForVotingIDRequest
	1, // 1: VotingVerifier.GetPublicKeyForVotingID:output_type -> GetPublicKeyForVotingIDResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_votingVerifier_proto_init() }
func file_votingVerifier_proto_init() {
	if File_votingVerifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_votingVerifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicKeyForVotingIDRequest); i {
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
		file_votingVerifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicKeyForVotingIDResponse); i {
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
			RawDescriptor: file_votingVerifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_votingVerifier_proto_goTypes,
		DependencyIndexes: file_votingVerifier_proto_depIdxs,
		MessageInfos:      file_votingVerifier_proto_msgTypes,
	}.Build()
	File_votingVerifier_proto = out.File
	file_votingVerifier_proto_rawDesc = nil
	file_votingVerifier_proto_goTypes = nil
	file_votingVerifier_proto_depIdxs = nil
}
