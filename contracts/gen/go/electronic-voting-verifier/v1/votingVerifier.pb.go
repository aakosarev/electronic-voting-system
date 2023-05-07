// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: votingVerifier.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type SignBlindedAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         int32  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	VotingID       int32  `protobuf:"varint,2,opt,name=votingID,proto3" json:"votingID,omitempty"`
	BlindedAddress []byte `protobuf:"bytes,3,opt,name=blindedAddress,proto3" json:"blindedAddress,omitempty"`
}

func (x *SignBlindedAddressRequest) Reset() {
	*x = SignBlindedAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_votingVerifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignBlindedAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignBlindedAddressRequest) ProtoMessage() {}

func (x *SignBlindedAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_votingVerifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignBlindedAddressRequest.ProtoReflect.Descriptor instead.
func (*SignBlindedAddressRequest) Descriptor() ([]byte, []int) {
	return file_votingVerifier_proto_rawDescGZIP(), []int{2}
}

func (x *SignBlindedAddressRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SignBlindedAddressRequest) GetVotingID() int32 {
	if x != nil {
		return x.VotingID
	}
	return 0
}

func (x *SignBlindedAddressRequest) GetBlindedAddress() []byte {
	if x != nil {
		return x.BlindedAddress
	}
	return nil
}

type SignBlindedAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedBlindedAddress []byte `protobuf:"bytes,1,opt,name=signedBlindedAddress,proto3" json:"signedBlindedAddress,omitempty"`
}

func (x *SignBlindedAddressResponse) Reset() {
	*x = SignBlindedAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_votingVerifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignBlindedAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignBlindedAddressResponse) ProtoMessage() {}

func (x *SignBlindedAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_votingVerifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignBlindedAddressResponse.ProtoReflect.Descriptor instead.
func (*SignBlindedAddressResponse) Descriptor() ([]byte, []int) {
	return file_votingVerifier_proto_rawDescGZIP(), []int{3}
}

func (x *SignBlindedAddressResponse) GetSignedBlindedAddress() []byte {
	if x != nil {
		return x.SignedBlindedAddress
	}
	return nil
}

type RegisterAddressToVotingBySignedTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VotingID    int32  `protobuf:"varint,1,opt,name=votingID,proto3" json:"votingID,omitempty"`
	Token       []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	SignedToken []byte `protobuf:"bytes,3,opt,name=signedToken,proto3" json:"signedToken,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RegisterAddressToVotingBySignedTokenRequest) Reset() {
	*x = RegisterAddressToVotingBySignedTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_votingVerifier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAddressToVotingBySignedTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAddressToVotingBySignedTokenRequest) ProtoMessage() {}

func (x *RegisterAddressToVotingBySignedTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_votingVerifier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAddressToVotingBySignedTokenRequest.ProtoReflect.Descriptor instead.
func (*RegisterAddressToVotingBySignedTokenRequest) Descriptor() ([]byte, []int) {
	return file_votingVerifier_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterAddressToVotingBySignedTokenRequest) GetVotingID() int32 {
	if x != nil {
		return x.VotingID
	}
	return 0
}

func (x *RegisterAddressToVotingBySignedTokenRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *RegisterAddressToVotingBySignedTokenRequest) GetSignedToken() []byte {
	if x != nil {
		return x.SignedToken
	}
	return nil
}

func (x *RegisterAddressToVotingBySignedTokenRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GenerateRSAKeyPairForVotingIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VotingID int32 `protobuf:"varint,1,opt,name=votingID,proto3" json:"votingID,omitempty"`
}

func (x *GenerateRSAKeyPairForVotingIDRequest) Reset() {
	*x = GenerateRSAKeyPairForVotingIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_votingVerifier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRSAKeyPairForVotingIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRSAKeyPairForVotingIDRequest) ProtoMessage() {}

func (x *GenerateRSAKeyPairForVotingIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_votingVerifier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRSAKeyPairForVotingIDRequest.ProtoReflect.Descriptor instead.
func (*GenerateRSAKeyPairForVotingIDRequest) Descriptor() ([]byte, []int) {
	return file_votingVerifier_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateRSAKeyPairForVotingIDRequest) GetVotingID() int32 {
	if x != nil {
		return x.VotingID
	}
	return 0
}

var File_votingVerifier_proto protoreflect.FileDescriptor

var file_votingVerifier_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x22, 0x49, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x77, 0x0a, 0x19,
	0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x0e, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x1a, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c, 0x69,
	0x6e, 0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6c, 0x69,
	0x6e, 0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x2b, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x56, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x24, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x46, 0x6f, 0x72, 0x56, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x32, 0x8b, 0x03, 0x0a, 0x0e, 0x56, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x53, 0x69,
	0x67, 0x6e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x24, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x56, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5e, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x46, 0x6f, 0x72,
	0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x46, 0x6f, 0x72,
	0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61, 0x6b, 0x6f, 0x73, 0x61, 0x72, 0x65, 0x76, 0x2f,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72,
	0x6f, 0x6e, 0x69, 0x63, 0x2d, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_votingVerifier_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_votingVerifier_proto_goTypes = []interface{}{
	(*GetPublicKeyForVotingIDRequest)(nil),              // 0: GetPublicKeyForVotingIDRequest
	(*GetPublicKeyForVotingIDResponse)(nil),             // 1: GetPublicKeyForVotingIDResponse
	(*SignBlindedAddressRequest)(nil),                   // 2: SignBlindedAddressRequest
	(*SignBlindedAddressResponse)(nil),                  // 3: SignBlindedAddressResponse
	(*RegisterAddressToVotingBySignedTokenRequest)(nil), // 4: RegisterAddressToVotingBySignedTokenRequest
	(*GenerateRSAKeyPairForVotingIDRequest)(nil),        // 5: GenerateRSAKeyPairForVotingIDRequest
	(*emptypb.Empty)(nil),                               // 6: google.protobuf.Empty
}
var file_votingVerifier_proto_depIdxs = []int32{
	0, // 0: VotingVerifier.GetPublicKeyForVotingID:input_type -> GetPublicKeyForVotingIDRequest
	2, // 1: VotingVerifier.SignBlindedAddress:input_type -> SignBlindedAddressRequest
	4, // 2: VotingVerifier.RegisterAddressToVotingBySignedToken:input_type -> RegisterAddressToVotingBySignedTokenRequest
	5, // 3: VotingVerifier.GenerateRSAKeyPairForVotingID:input_type -> GenerateRSAKeyPairForVotingIDRequest
	1, // 4: VotingVerifier.GetPublicKeyForVotingID:output_type -> GetPublicKeyForVotingIDResponse
	3, // 5: VotingVerifier.SignBlindedAddress:output_type -> SignBlindedAddressResponse
	6, // 6: VotingVerifier.RegisterAddressToVotingBySignedToken:output_type -> google.protobuf.Empty
	6, // 7: VotingVerifier.GenerateRSAKeyPairForVotingID:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
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
		file_votingVerifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignBlindedAddressRequest); i {
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
		file_votingVerifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignBlindedAddressResponse); i {
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
		file_votingVerifier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAddressToVotingBySignedTokenRequest); i {
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
		file_votingVerifier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRSAKeyPairForVotingIDRequest); i {
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
			NumMessages:   6,
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
