// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: child_profile.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Create
type ChildGender int32

const (
	ChildGender_CHILD_GENDER_UNSPECIFIED ChildGender = 0
	ChildGender_CHILD_GENDER_BOY         ChildGender = 1
	ChildGender_CHILD_GENDER_GIRL        ChildGender = 2
)

// Enum value maps for ChildGender.
var (
	ChildGender_name = map[int32]string{
		0: "CHILD_GENDER_UNSPECIFIED",
		1: "CHILD_GENDER_BOY",
		2: "CHILD_GENDER_GIRL",
	}
	ChildGender_value = map[string]int32{
		"CHILD_GENDER_UNSPECIFIED": 0,
		"CHILD_GENDER_BOY":         1,
		"CHILD_GENDER_GIRL":        2,
	}
)

func (x ChildGender) Enum() *ChildGender {
	p := new(ChildGender)
	*p = x
	return p
}

func (x ChildGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChildGender) Descriptor() protoreflect.EnumDescriptor {
	return file_child_profile_proto_enumTypes[0].Descriptor()
}

func (ChildGender) Type() protoreflect.EnumType {
	return &file_child_profile_proto_enumTypes[0]
}

func (x ChildGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChildGender.Descriptor instead.
func (ChildGender) EnumDescriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{0}
}

type ChildProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender    ChildGender            `protobuf:"varint,4,opt,name=gender,proto3,enum=child_profile.ChildGender" json:"gender,omitempty"`
	Birthday  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	FamilyId  string                 `protobuf:"bytes,6,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *ChildProfile) Reset() {
	*x = ChildProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildProfile) ProtoMessage() {}

func (x *ChildProfile) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildProfile.ProtoReflect.Descriptor instead.
func (*ChildProfile) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{0}
}

func (x *ChildProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChildProfile) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ChildProfile) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ChildProfile) GetGender() ChildGender {
	if x != nil {
		return x.Gender
	}
	return ChildGender_CHILD_GENDER_UNSPECIFIED
}

func (x *ChildProfile) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *ChildProfile) GetFamilyId() string {
	if x != nil {
		return x.FamilyId
	}
	return ""
}

// Find one
type FindOneChildProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FamilyId string `protobuf:"bytes,2,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *FindOneChildProfileRequest) Reset() {
	*x = FindOneChildProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneChildProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneChildProfileRequest) ProtoMessage() {}

func (x *FindOneChildProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneChildProfileRequest.ProtoReflect.Descriptor instead.
func (*FindOneChildProfileRequest) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{1}
}

func (x *FindOneChildProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FindOneChildProfileRequest) GetFamilyId() string {
	if x != nil {
		return x.FamilyId
	}
	return ""
}

type FindOneChildProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ChildProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *FindOneChildProfileResponse) Reset() {
	*x = FindOneChildProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneChildProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneChildProfileResponse) ProtoMessage() {}

func (x *FindOneChildProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneChildProfileResponse.ProtoReflect.Descriptor instead.
func (*FindOneChildProfileResponse) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneChildProfileResponse) GetProfile() *ChildProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type CreateChildProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender    ChildGender            `protobuf:"varint,3,opt,name=gender,proto3,enum=child_profile.ChildGender" json:"gender,omitempty"`
	Birthday  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	FamilyId  string                 `protobuf:"bytes,5,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *CreateChildProfileRequest) Reset() {
	*x = CreateChildProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChildProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChildProfileRequest) ProtoMessage() {}

func (x *CreateChildProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChildProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateChildProfileRequest) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{3}
}

func (x *CreateChildProfileRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateChildProfileRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateChildProfileRequest) GetGender() ChildGender {
	if x != nil {
		return x.Gender
	}
	return ChildGender_CHILD_GENDER_UNSPECIFIED
}

func (x *CreateChildProfileRequest) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *CreateChildProfileRequest) GetFamilyId() string {
	if x != nil {
		return x.FamilyId
	}
	return ""
}

type CreateChildProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ChildProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CreateChildProfileResponse) Reset() {
	*x = CreateChildProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChildProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChildProfileResponse) ProtoMessage() {}

func (x *CreateChildProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChildProfileResponse.ProtoReflect.Descriptor instead.
func (*CreateChildProfileResponse) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{4}
}

func (x *CreateChildProfileResponse) GetProfile() *ChildProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// Update
type UpdateChildProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ChildProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateChildProfileRequest) Reset() {
	*x = UpdateChildProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChildProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChildProfileRequest) ProtoMessage() {}

func (x *UpdateChildProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChildProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateChildProfileRequest) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateChildProfileRequest) GetProfile() *ChildProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type UpdateChildProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ChildProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateChildProfileResponse) Reset() {
	*x = UpdateChildProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChildProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChildProfileResponse) ProtoMessage() {}

func (x *UpdateChildProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChildProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateChildProfileResponse) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateChildProfileResponse) GetProfile() *ChildProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// Delete
type DeleteChildProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FamilyId string `protobuf:"bytes,2,opt,name=family_id,json=familyId,proto3" json:"family_id,omitempty"`
}

func (x *DeleteChildProfileRequest) Reset() {
	*x = DeleteChildProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChildProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChildProfileRequest) ProtoMessage() {}

func (x *DeleteChildProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChildProfileRequest.ProtoReflect.Descriptor instead.
func (*DeleteChildProfileRequest) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteChildProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteChildProfileRequest) GetFamilyId() string {
	if x != nil {
		return x.FamilyId
	}
	return ""
}

type DeleteChildProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteChildProfileResponse) Reset() {
	*x = DeleteChildProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_child_profile_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChildProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChildProfileResponse) ProtoMessage() {}

func (x *DeleteChildProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_child_profile_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChildProfileResponse.ProtoReflect.Descriptor instead.
func (*DeleteChildProfileResponse) Descriptor() ([]byte, []int) {
	return file_child_profile_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteChildProfileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_child_profile_proto protoreflect.FileDescriptor

var file_child_profile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x1a, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e,
	0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xe0, 0x01, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22,
	0x53, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x53, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x48, 0x0a,
	0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a,
	0x58, 0x0a, 0x0b, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x18, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x4f, 0x59,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x47, 0x45, 0x4e, 0x44,
	0x45, 0x52, 0x5f, 0x47, 0x49, 0x52, 0x4c, 0x10, 0x02, 0x32, 0xcc, 0x03, 0x0a, 0x13, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6e, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_child_profile_proto_rawDescOnce sync.Once
	file_child_profile_proto_rawDescData = file_child_profile_proto_rawDesc
)

func file_child_profile_proto_rawDescGZIP() []byte {
	file_child_profile_proto_rawDescOnce.Do(func() {
		file_child_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_child_profile_proto_rawDescData)
	})
	return file_child_profile_proto_rawDescData
}

var file_child_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_child_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_child_profile_proto_goTypes = []interface{}{
	(ChildGender)(0),                    // 0: child_profile.ChildGender
	(*ChildProfile)(nil),                // 1: child_profile.ChildProfile
	(*FindOneChildProfileRequest)(nil),  // 2: child_profile.FindOneChildProfileRequest
	(*FindOneChildProfileResponse)(nil), // 3: child_profile.FindOneChildProfileResponse
	(*CreateChildProfileRequest)(nil),   // 4: child_profile.CreateChildProfileRequest
	(*CreateChildProfileResponse)(nil),  // 5: child_profile.CreateChildProfileResponse
	(*UpdateChildProfileRequest)(nil),   // 6: child_profile.UpdateChildProfileRequest
	(*UpdateChildProfileResponse)(nil),  // 7: child_profile.UpdateChildProfileResponse
	(*DeleteChildProfileRequest)(nil),   // 8: child_profile.DeleteChildProfileRequest
	(*DeleteChildProfileResponse)(nil),  // 9: child_profile.DeleteChildProfileResponse
	(*timestamppb.Timestamp)(nil),       // 10: google.protobuf.Timestamp
}
var file_child_profile_proto_depIdxs = []int32{
	0,  // 0: child_profile.ChildProfile.gender:type_name -> child_profile.ChildGender
	10, // 1: child_profile.ChildProfile.birthday:type_name -> google.protobuf.Timestamp
	1,  // 2: child_profile.FindOneChildProfileResponse.profile:type_name -> child_profile.ChildProfile
	0,  // 3: child_profile.CreateChildProfileRequest.gender:type_name -> child_profile.ChildGender
	10, // 4: child_profile.CreateChildProfileRequest.birthday:type_name -> google.protobuf.Timestamp
	1,  // 5: child_profile.CreateChildProfileResponse.profile:type_name -> child_profile.ChildProfile
	1,  // 6: child_profile.UpdateChildProfileRequest.profile:type_name -> child_profile.ChildProfile
	1,  // 7: child_profile.UpdateChildProfileResponse.profile:type_name -> child_profile.ChildProfile
	2,  // 8: child_profile.ChildProfileService.FindOneChildProfile:input_type -> child_profile.FindOneChildProfileRequest
	4,  // 9: child_profile.ChildProfileService.CreateChildProfile:input_type -> child_profile.CreateChildProfileRequest
	6,  // 10: child_profile.ChildProfileService.UpdateChildProfile:input_type -> child_profile.UpdateChildProfileRequest
	8,  // 11: child_profile.ChildProfileService.DeleteChildProfile:input_type -> child_profile.DeleteChildProfileRequest
	3,  // 12: child_profile.ChildProfileService.FindOneChildProfile:output_type -> child_profile.FindOneChildProfileResponse
	5,  // 13: child_profile.ChildProfileService.CreateChildProfile:output_type -> child_profile.CreateChildProfileResponse
	7,  // 14: child_profile.ChildProfileService.UpdateChildProfile:output_type -> child_profile.UpdateChildProfileResponse
	9,  // 15: child_profile.ChildProfileService.DeleteChildProfile:output_type -> child_profile.DeleteChildProfileResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_child_profile_proto_init() }
func file_child_profile_proto_init() {
	if File_child_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_child_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildProfile); i {
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
		file_child_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneChildProfileRequest); i {
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
		file_child_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneChildProfileResponse); i {
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
		file_child_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChildProfileRequest); i {
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
		file_child_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChildProfileResponse); i {
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
		file_child_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChildProfileRequest); i {
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
		file_child_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChildProfileResponse); i {
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
		file_child_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChildProfileRequest); i {
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
		file_child_profile_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChildProfileResponse); i {
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
			RawDescriptor: file_child_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_child_profile_proto_goTypes,
		DependencyIndexes: file_child_profile_proto_depIdxs,
		EnumInfos:         file_child_profile_proto_enumTypes,
		MessageInfos:      file_child_profile_proto_msgTypes,
	}.Build()
	File_child_profile_proto = out.File
	file_child_profile_proto_rawDesc = nil
	file_child_profile_proto_goTypes = nil
	file_child_profile_proto_depIdxs = nil
}