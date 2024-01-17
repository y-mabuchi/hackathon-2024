// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: family.proto

package generated

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

type Family struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Family) Reset() {
	*x = Family{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Family) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Family) ProtoMessage() {}

func (x *Family) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Family.ProtoReflect.Descriptor instead.
func (*Family) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{0}
}

func (x *Family) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Family) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Find one
type FindOneFamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneFamilyRequest) Reset() {
	*x = FindOneFamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneFamilyRequest) ProtoMessage() {}

func (x *FindOneFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneFamilyRequest.ProtoReflect.Descriptor instead.
func (*FindOneFamilyRequest) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{1}
}

func (x *FindOneFamilyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneFamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family *Family `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
}

func (x *FindOneFamilyResponse) Reset() {
	*x = FindOneFamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneFamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneFamilyResponse) ProtoMessage() {}

func (x *FindOneFamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneFamilyResponse.ProtoReflect.Descriptor instead.
func (*FindOneFamilyResponse) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneFamilyResponse) GetFamily() *Family {
	if x != nil {
		return x.Family
	}
	return nil
}

// Create
type CreateFamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateFamilyRequest) Reset() {
	*x = CreateFamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFamilyRequest) ProtoMessage() {}

func (x *CreateFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFamilyRequest.ProtoReflect.Descriptor instead.
func (*CreateFamilyRequest) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFamilyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateFamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family *Family `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
}

func (x *CreateFamilyResponse) Reset() {
	*x = CreateFamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFamilyResponse) ProtoMessage() {}

func (x *CreateFamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFamilyResponse.ProtoReflect.Descriptor instead.
func (*CreateFamilyResponse) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFamilyResponse) GetFamily() *Family {
	if x != nil {
		return x.Family
	}
	return nil
}

// Update
type UpdateFamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family *Family `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
}

func (x *UpdateFamilyRequest) Reset() {
	*x = UpdateFamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFamilyRequest) ProtoMessage() {}

func (x *UpdateFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFamilyRequest.ProtoReflect.Descriptor instead.
func (*UpdateFamilyRequest) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFamilyRequest) GetFamily() *Family {
	if x != nil {
		return x.Family
	}
	return nil
}

type UpdateFamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family *Family `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
}

func (x *UpdateFamilyResponse) Reset() {
	*x = UpdateFamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFamilyResponse) ProtoMessage() {}

func (x *UpdateFamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFamilyResponse.ProtoReflect.Descriptor instead.
func (*UpdateFamilyResponse) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFamilyResponse) GetFamily() *Family {
	if x != nil {
		return x.Family
	}
	return nil
}

// Delete
type DeleteFamilyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFamilyRequest) Reset() {
	*x = DeleteFamilyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFamilyRequest) ProtoMessage() {}

func (x *DeleteFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFamilyRequest.ProtoReflect.Descriptor instead.
func (*DeleteFamilyRequest) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFamilyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFamilyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteFamilyResponse) Reset() {
	*x = DeleteFamilyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_family_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFamilyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFamilyResponse) ProtoMessage() {}

func (x *DeleteFamilyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_family_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFamilyResponse.ProtoReflect.Descriptor instead.
func (*DeleteFamilyResponse) Descriptor() ([]byte, []int) {
	return file_family_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFamilyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_family_proto protoreflect.FileDescriptor

var file_family_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x2c, 0x0a, 0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x15,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x29, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52,
	0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52,
	0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xc6, 0x02, 0x0a, 0x0d, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x12, 0x1c, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x12, 0x1b, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12,
	0x1b, 0x2e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1b, 0x2e, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_family_proto_rawDescOnce sync.Once
	file_family_proto_rawDescData = file_family_proto_rawDesc
)

func file_family_proto_rawDescGZIP() []byte {
	file_family_proto_rawDescOnce.Do(func() {
		file_family_proto_rawDescData = protoimpl.X.CompressGZIP(file_family_proto_rawDescData)
	})
	return file_family_proto_rawDescData
}

var file_family_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_family_proto_goTypes = []interface{}{
	(*Family)(nil),                // 0: family.Family
	(*FindOneFamilyRequest)(nil),  // 1: family.FindOneFamilyRequest
	(*FindOneFamilyResponse)(nil), // 2: family.FindOneFamilyResponse
	(*CreateFamilyRequest)(nil),   // 3: family.CreateFamilyRequest
	(*CreateFamilyResponse)(nil),  // 4: family.CreateFamilyResponse
	(*UpdateFamilyRequest)(nil),   // 5: family.UpdateFamilyRequest
	(*UpdateFamilyResponse)(nil),  // 6: family.UpdateFamilyResponse
	(*DeleteFamilyRequest)(nil),   // 7: family.DeleteFamilyRequest
	(*DeleteFamilyResponse)(nil),  // 8: family.DeleteFamilyResponse
}
var file_family_proto_depIdxs = []int32{
	0, // 0: family.FindOneFamilyResponse.family:type_name -> family.Family
	0, // 1: family.CreateFamilyResponse.family:type_name -> family.Family
	0, // 2: family.UpdateFamilyRequest.family:type_name -> family.Family
	0, // 3: family.UpdateFamilyResponse.family:type_name -> family.Family
	1, // 4: family.FamilyService.FindOneFamily:input_type -> family.FindOneFamilyRequest
	3, // 5: family.FamilyService.CreateFamily:input_type -> family.CreateFamilyRequest
	5, // 6: family.FamilyService.UpdateFamily:input_type -> family.UpdateFamilyRequest
	7, // 7: family.FamilyService.DeleteFamily:input_type -> family.DeleteFamilyRequest
	2, // 8: family.FamilyService.FindOneFamily:output_type -> family.FindOneFamilyResponse
	4, // 9: family.FamilyService.CreateFamily:output_type -> family.CreateFamilyResponse
	6, // 10: family.FamilyService.UpdateFamily:output_type -> family.UpdateFamilyResponse
	8, // 11: family.FamilyService.DeleteFamily:output_type -> family.DeleteFamilyResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_family_proto_init() }
func file_family_proto_init() {
	if File_family_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_family_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Family); i {
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
		file_family_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneFamilyRequest); i {
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
		file_family_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneFamilyResponse); i {
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
		file_family_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFamilyRequest); i {
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
		file_family_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFamilyResponse); i {
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
		file_family_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFamilyRequest); i {
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
		file_family_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFamilyResponse); i {
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
		file_family_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFamilyRequest); i {
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
		file_family_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFamilyResponse); i {
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
			RawDescriptor: file_family_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_family_proto_goTypes,
		DependencyIndexes: file_family_proto_depIdxs,
		MessageInfos:      file_family_proto_msgTypes,
	}.Build()
	File_family_proto = out.File
	file_family_proto_rawDesc = nil
	file_family_proto_goTypes = nil
	file_family_proto_depIdxs = nil
}