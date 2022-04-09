// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: ci.proto

package pb

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

type BuildPlanModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectID  string `protobuf:"bytes,3,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Tag        string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	DockerFile string `protobuf:"bytes,5,opt,name=dockerFile,proto3" json:"dockerFile,omitempty"`
}

func (x *BuildPlanModel) Reset() {
	*x = BuildPlanModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildPlanModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildPlanModel) ProtoMessage() {}

func (x *BuildPlanModel) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildPlanModel.ProtoReflect.Descriptor instead.
func (*BuildPlanModel) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{0}
}

func (x *BuildPlanModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BuildPlanModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildPlanModel) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *BuildPlanModel) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *BuildPlanModel) GetDockerFile() string {
	if x != nil {
		return x.DockerFile
	}
	return ""
}

type GetBuildPlanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBuildPlanReq) Reset() {
	*x = GetBuildPlanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuildPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuildPlanReq) ProtoMessage() {}

func (x *GetBuildPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuildPlanReq.ProtoReflect.Descriptor instead.
func (*GetBuildPlanReq) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{1}
}

func (x *GetBuildPlanReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBuildPlanRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectID  string `protobuf:"bytes,3,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Tag        string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	DockerFile string `protobuf:"bytes,5,opt,name=dockerFile,proto3" json:"dockerFile,omitempty"`
}

func (x *GetBuildPlanRsp) Reset() {
	*x = GetBuildPlanRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuildPlanRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuildPlanRsp) ProtoMessage() {}

func (x *GetBuildPlanRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuildPlanRsp.ProtoReflect.Descriptor instead.
func (*GetBuildPlanRsp) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{2}
}

func (x *GetBuildPlanRsp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBuildPlanRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBuildPlanRsp) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *GetBuildPlanRsp) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *GetBuildPlanRsp) GetDockerFile() string {
	if x != nil {
		return x.DockerFile
	}
	return ""
}

type AddBuildPlanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProjectID  string `protobuf:"bytes,2,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Tag        string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	DockerFile string `protobuf:"bytes,4,opt,name=dockerFile,proto3" json:"dockerFile,omitempty"`
}

func (x *AddBuildPlanReq) Reset() {
	*x = AddBuildPlanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBuildPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBuildPlanReq) ProtoMessage() {}

func (x *AddBuildPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBuildPlanReq.ProtoReflect.Descriptor instead.
func (*AddBuildPlanReq) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{3}
}

func (x *AddBuildPlanReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddBuildPlanReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *AddBuildPlanReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AddBuildPlanReq) GetDockerFile() string {
	if x != nil {
		return x.DockerFile
	}
	return ""
}

type AddBuildPlanRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddBuildPlanRsp) Reset() {
	*x = AddBuildPlanRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBuildPlanRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBuildPlanRsp) ProtoMessage() {}

func (x *AddBuildPlanRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBuildPlanRsp.ProtoReflect.Descriptor instead.
func (*AddBuildPlanRsp) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{4}
}

func (x *AddBuildPlanRsp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBuildPlanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag        string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	DockerFile string `protobuf:"bytes,4,opt,name=dockerFile,proto3" json:"dockerFile,omitempty"`
}

func (x *UpdateBuildPlanReq) Reset() {
	*x = UpdateBuildPlanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBuildPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBuildPlanReq) ProtoMessage() {}

func (x *UpdateBuildPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBuildPlanReq.ProtoReflect.Descriptor instead.
func (*UpdateBuildPlanReq) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBuildPlanReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBuildPlanReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBuildPlanReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *UpdateBuildPlanReq) GetDockerFile() string {
	if x != nil {
		return x.DockerFile
	}
	return ""
}

type UpdateBuildPlanRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBuildPlanRsp) Reset() {
	*x = UpdateBuildPlanRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBuildPlanRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBuildPlanRsp) ProtoMessage() {}

func (x *UpdateBuildPlanRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBuildPlanRsp.ProtoReflect.Descriptor instead.
func (*UpdateBuildPlanRsp) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{6}
}

type DeleteBuildPlanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBuildPlanReq) Reset() {
	*x = DeleteBuildPlanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBuildPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBuildPlanReq) ProtoMessage() {}

func (x *DeleteBuildPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBuildPlanReq.ProtoReflect.Descriptor instead.
func (*DeleteBuildPlanReq) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBuildPlanReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBuildPlanRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBuildPlanRsp) Reset() {
	*x = DeleteBuildPlanRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBuildPlanRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBuildPlanRsp) ProtoMessage() {}

func (x *DeleteBuildPlanRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBuildPlanRsp.ProtoReflect.Descriptor instead.
func (*DeleteBuildPlanRsp) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{8}
}

type ListBuildPlanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID string `protobuf:"bytes,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
}

func (x *ListBuildPlanReq) Reset() {
	*x = ListBuildPlanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildPlanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildPlanReq) ProtoMessage() {}

func (x *ListBuildPlanReq) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildPlanReq.ProtoReflect.Descriptor instead.
func (*ListBuildPlanReq) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{9}
}

func (x *ListBuildPlanReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

type ListBuildPlanRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64             `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Result []*BuildPlanModel `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListBuildPlanRsp) Reset() {
	*x = ListBuildPlanRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildPlanRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildPlanRsp) ProtoMessage() {}

func (x *ListBuildPlanRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildPlanRsp.ProtoReflect.Descriptor instead.
func (*ListBuildPlanRsp) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{10}
}

func (x *ListBuildPlanRsp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListBuildPlanRsp) GetResult() []*BuildPlanModel {
	if x != nil {
		return x.Result
	}
	return nil
}

type BuildReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID   string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	BuildPlanID string `protobuf:"bytes,2,opt,name=buildPlanID,proto3" json:"buildPlanID,omitempty"`
}

func (x *BuildReq) Reset() {
	*x = BuildReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildReq) ProtoMessage() {}

func (x *BuildReq) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildReq.ProtoReflect.Descriptor instead.
func (*BuildReq) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{11}
}

func (x *BuildReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *BuildReq) GetBuildPlanID() string {
	if x != nil {
		return x.BuildPlanID
	}
	return ""
}

type BuildRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuildRsp) Reset() {
	*x = BuildRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRsp) ProtoMessage() {}

func (x *BuildRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ci_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRsp.ProtoReflect.Descriptor instead.
func (*BuildRsp) Descriptor() ([]byte, []int) {
	return file_ci_proto_rawDescGZIP(), []int{12}
}

var File_ci_proto protoreflect.FileDescriptor

var file_ci_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x63, 0x69, 0x22, 0x84,
	0x01, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x22, 0x75, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x30, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x22, 0x54, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x69, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x4a, 0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x22, 0x0a, 0x0a, 0x08,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x73, 0x70, 0x32, 0xe5, 0x02, 0x0a, 0x02, 0x63, 0x69, 0x12,
	0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x13, 0x2e, 0x63, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0c, 0x41, 0x64, 0x64,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x13, 0x2e, 0x63, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x63, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x63, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x63, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x14, 0x2e, 0x63, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x63, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x63, 0x69, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x63, 0x69, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x73, 0x70,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ci_proto_rawDescOnce sync.Once
	file_ci_proto_rawDescData = file_ci_proto_rawDesc
)

func file_ci_proto_rawDescGZIP() []byte {
	file_ci_proto_rawDescOnce.Do(func() {
		file_ci_proto_rawDescData = protoimpl.X.CompressGZIP(file_ci_proto_rawDescData)
	})
	return file_ci_proto_rawDescData
}

var file_ci_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_ci_proto_goTypes = []interface{}{
	(*BuildPlanModel)(nil),     // 0: ci.BuildPlanModel
	(*GetBuildPlanReq)(nil),    // 1: ci.GetBuildPlanReq
	(*GetBuildPlanRsp)(nil),    // 2: ci.GetBuildPlanRsp
	(*AddBuildPlanReq)(nil),    // 3: ci.AddBuildPlanReq
	(*AddBuildPlanRsp)(nil),    // 4: ci.AddBuildPlanRsp
	(*UpdateBuildPlanReq)(nil), // 5: ci.UpdateBuildPlanReq
	(*UpdateBuildPlanRsp)(nil), // 6: ci.UpdateBuildPlanRsp
	(*DeleteBuildPlanReq)(nil), // 7: ci.DeleteBuildPlanReq
	(*DeleteBuildPlanRsp)(nil), // 8: ci.DeleteBuildPlanRsp
	(*ListBuildPlanReq)(nil),   // 9: ci.ListBuildPlanReq
	(*ListBuildPlanRsp)(nil),   // 10: ci.ListBuildPlanRsp
	(*BuildReq)(nil),           // 11: ci.BuildReq
	(*BuildRsp)(nil),           // 12: ci.BuildRsp
}
var file_ci_proto_depIdxs = []int32{
	0,  // 0: ci.ListBuildPlanRsp.result:type_name -> ci.BuildPlanModel
	1,  // 1: ci.ci.GetBuildPlan:input_type -> ci.GetBuildPlanReq
	3,  // 2: ci.ci.AddBuildPlan:input_type -> ci.AddBuildPlanReq
	5,  // 3: ci.ci.UpdateBuildPlan:input_type -> ci.UpdateBuildPlanReq
	7,  // 4: ci.ci.DeleteBuildPlan:input_type -> ci.DeleteBuildPlanReq
	9,  // 5: ci.ci.ListBuildPlan:input_type -> ci.ListBuildPlanReq
	11, // 6: ci.ci.BuildImage:input_type -> ci.BuildReq
	2,  // 7: ci.ci.GetBuildPlan:output_type -> ci.GetBuildPlanRsp
	4,  // 8: ci.ci.AddBuildPlan:output_type -> ci.AddBuildPlanRsp
	6,  // 9: ci.ci.UpdateBuildPlan:output_type -> ci.UpdateBuildPlanRsp
	8,  // 10: ci.ci.DeleteBuildPlan:output_type -> ci.DeleteBuildPlanRsp
	10, // 11: ci.ci.ListBuildPlan:output_type -> ci.ListBuildPlanRsp
	12, // 12: ci.ci.BuildImage:output_type -> ci.BuildRsp
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_ci_proto_init() }
func file_ci_proto_init() {
	if File_ci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildPlanModel); i {
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
		file_ci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBuildPlanReq); i {
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
		file_ci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBuildPlanRsp); i {
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
		file_ci_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBuildPlanReq); i {
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
		file_ci_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBuildPlanRsp); i {
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
		file_ci_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBuildPlanReq); i {
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
		file_ci_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBuildPlanRsp); i {
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
		file_ci_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBuildPlanReq); i {
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
		file_ci_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBuildPlanRsp); i {
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
		file_ci_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildPlanReq); i {
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
		file_ci_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildPlanRsp); i {
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
		file_ci_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildReq); i {
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
		file_ci_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRsp); i {
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
			RawDescriptor: file_ci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ci_proto_goTypes,
		DependencyIndexes: file_ci_proto_depIdxs,
		MessageInfos:      file_ci_proto_msgTypes,
	}.Build()
	File_ci_proto = out.File
	file_ci_proto_rawDesc = nil
	file_ci_proto_goTypes = nil
	file_ci_proto_depIdxs = nil
}
