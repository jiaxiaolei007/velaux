// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/proto/catalogservice/service.proto

package catalogservice

import (
	proto "github.com/golang/protobuf/proto"
	model "github.com/oam-dev/velacp/pkg/datastore/model"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PutCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalog *model.Catalog `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
}

func (x *PutCatalogRequest) Reset() {
	*x = PutCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutCatalogRequest) ProtoMessage() {}

func (x *PutCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutCatalogRequest.ProtoReflect.Descriptor instead.
func (*PutCatalogRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{0}
}

func (x *PutCatalogRequest) GetCatalog() *model.Catalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

type PutCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutCatalogResponse) Reset() {
	*x = PutCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutCatalogResponse) ProtoMessage() {}

func (x *PutCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutCatalogResponse.ProtoReflect.Descriptor instead.
func (*PutCatalogResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{1}
}

type GetCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalog *model.Catalog `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
}

func (x *GetCatalogResponse) Reset() {
	*x = GetCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogResponse) ProtoMessage() {}

func (x *GetCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogResponse.ProtoReflect.Descriptor instead.
func (*GetCatalogResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCatalogResponse) GetCatalog() *model.Catalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

type ListCatalogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCatalogsRequest) Reset() {
	*x = ListCatalogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsRequest) ProtoMessage() {}

func (x *ListCatalogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsRequest.ProtoReflect.Descriptor instead.
func (*ListCatalogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{4}
}

type ListCatalogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalogs []*model.Catalog `protobuf:"bytes,1,rep,name=catalogs,proto3" json:"catalogs,omitempty"`
}

func (x *ListCatalogsResponse) Reset() {
	*x = ListCatalogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsResponse) ProtoMessage() {}

func (x *ListCatalogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsResponse.ProtoReflect.Descriptor instead.
func (*ListCatalogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListCatalogsResponse) GetCatalogs() []*model.Catalog {
	if x != nil {
		return x.Catalogs
	}
	return nil
}

type DelCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DelCatalogRequest) Reset() {
	*x = DelCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelCatalogRequest) ProtoMessage() {}

func (x *DelCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelCatalogRequest.ProtoReflect.Descriptor instead.
func (*DelCatalogRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{6}
}

func (x *DelCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DelCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelCatalogResponse) Reset() {
	*x = DelCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelCatalogResponse) ProtoMessage() {}

func (x *DelCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelCatalogResponse.ProtoReflect.Descriptor instead.
func (*DelCatalogResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{7}
}

type SyncCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SyncCatalogRequest) Reset() {
	*x = SyncCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCatalogRequest) ProtoMessage() {}

func (x *SyncCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCatalogRequest.ProtoReflect.Descriptor instead.
func (*SyncCatalogRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{8}
}

func (x *SyncCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SyncCatalogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncCatalogResponse) Reset() {
	*x = SyncCatalogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCatalogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCatalogResponse) ProtoMessage() {}

func (x *SyncCatalogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_catalogservice_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCatalogResponse.ProtoReflect.Descriptor instead.
func (*SyncCatalogResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_catalogservice_service_proto_rawDescGZIP(), []int{9}
}

var File_pkg_proto_catalogservice_service_proto protoreflect.FileDescriptor

var file_pkg_proto_catalogservice_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x50, 0x75,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x75, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x4b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65,
	0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x08, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x27,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a,
	0x12, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8c,
	0x05, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x77, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12,
	0x32, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x32, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x76,
	0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x92, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x34, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x76, 0x65, 0x6c,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x77, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x32, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x76, 0x65, 0x6c,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7a, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x12, 0x33, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x61, 0x6d, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x76, 0x65, 0x6c, 0x61, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_catalogservice_service_proto_rawDescOnce sync.Once
	file_pkg_proto_catalogservice_service_proto_rawDescData = file_pkg_proto_catalogservice_service_proto_rawDesc
)

func file_pkg_proto_catalogservice_service_proto_rawDescGZIP() []byte {
	file_pkg_proto_catalogservice_service_proto_rawDescOnce.Do(func() {
		file_pkg_proto_catalogservice_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_catalogservice_service_proto_rawDescData)
	})
	return file_pkg_proto_catalogservice_service_proto_rawDescData
}

var file_pkg_proto_catalogservice_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_proto_catalogservice_service_proto_goTypes = []interface{}{
	(*PutCatalogRequest)(nil),    // 0: vela.api.service.catalogservice.PutCatalogRequest
	(*PutCatalogResponse)(nil),   // 1: vela.api.service.catalogservice.PutCatalogResponse
	(*GetCatalogRequest)(nil),    // 2: vela.api.service.catalogservice.GetCatalogRequest
	(*GetCatalogResponse)(nil),   // 3: vela.api.service.catalogservice.GetCatalogResponse
	(*ListCatalogsRequest)(nil),  // 4: vela.api.service.catalogservice.ListCatalogsRequest
	(*ListCatalogsResponse)(nil), // 5: vela.api.service.catalogservice.ListCatalogsResponse
	(*DelCatalogRequest)(nil),    // 6: vela.api.service.catalogservice.DelCatalogRequest
	(*DelCatalogResponse)(nil),   // 7: vela.api.service.catalogservice.DelCatalogResponse
	(*SyncCatalogRequest)(nil),   // 8: vela.api.service.catalogservice.SyncCatalogRequest
	(*SyncCatalogResponse)(nil),  // 9: vela.api.service.catalogservice.SyncCatalogResponse
	(*model.Catalog)(nil),        // 10: vela.api.model.Catalog
}
var file_pkg_proto_catalogservice_service_proto_depIdxs = []int32{
	10, // 0: vela.api.service.catalogservice.PutCatalogRequest.catalog:type_name -> vela.api.model.Catalog
	10, // 1: vela.api.service.catalogservice.GetCatalogResponse.catalog:type_name -> vela.api.model.Catalog
	10, // 2: vela.api.service.catalogservice.ListCatalogsResponse.catalogs:type_name -> vela.api.model.Catalog
	0,  // 3: vela.api.service.catalogservice.CatalogService.PutCatalog:input_type -> vela.api.service.catalogservice.PutCatalogRequest
	2,  // 4: vela.api.service.catalogservice.CatalogService.GetCatalog:input_type -> vela.api.service.catalogservice.GetCatalogRequest
	4,  // 5: vela.api.service.catalogservice.CatalogService.ListCatalogs:input_type -> vela.api.service.catalogservice.ListCatalogsRequest
	6,  // 6: vela.api.service.catalogservice.CatalogService.DelCatalog:input_type -> vela.api.service.catalogservice.DelCatalogRequest
	8,  // 7: vela.api.service.catalogservice.CatalogService.SyncCatalog:input_type -> vela.api.service.catalogservice.SyncCatalogRequest
	1,  // 8: vela.api.service.catalogservice.CatalogService.PutCatalog:output_type -> vela.api.service.catalogservice.PutCatalogResponse
	3,  // 9: vela.api.service.catalogservice.CatalogService.GetCatalog:output_type -> vela.api.service.catalogservice.GetCatalogResponse
	5,  // 10: vela.api.service.catalogservice.CatalogService.ListCatalogs:output_type -> vela.api.service.catalogservice.ListCatalogsResponse
	7,  // 11: vela.api.service.catalogservice.CatalogService.DelCatalog:output_type -> vela.api.service.catalogservice.DelCatalogResponse
	9,  // 12: vela.api.service.catalogservice.CatalogService.SyncCatalog:output_type -> vela.api.service.catalogservice.SyncCatalogResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_catalogservice_service_proto_init() }
func file_pkg_proto_catalogservice_service_proto_init() {
	if File_pkg_proto_catalogservice_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_catalogservice_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutCatalogRequest); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutCatalogResponse); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogRequest); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogResponse); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogsRequest); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogsResponse); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelCatalogRequest); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelCatalogResponse); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCatalogRequest); i {
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
		file_pkg_proto_catalogservice_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCatalogResponse); i {
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
			RawDescriptor: file_pkg_proto_catalogservice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_catalogservice_service_proto_goTypes,
		DependencyIndexes: file_pkg_proto_catalogservice_service_proto_depIdxs,
		MessageInfos:      file_pkg_proto_catalogservice_service_proto_msgTypes,
	}.Build()
	File_pkg_proto_catalogservice_service_proto = out.File
	file_pkg_proto_catalogservice_service_proto_rawDesc = nil
	file_pkg_proto_catalogservice_service_proto_goTypes = nil
	file_pkg_proto_catalogservice_service_proto_depIdxs = nil
}
