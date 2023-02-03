// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.0
// source: grpc/proto/images_handler.proto

package api

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateDate string `protobuf:"bytes,2,opt,name=createDate,proto3" json:"createDate,omitempty"`
	UpdateDate string `protobuf:"bytes,3,opt,name=updateDate,proto3" json:"updateDate,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *Image) GetUpdateDate() string {
	if x != nil {
		return x.UpdateDate
	}
	return ""
}

type SetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetImageRequest) Reset() {
	*x = SetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetImageRequest) ProtoMessage() {}

func (x *SetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetImageRequest.ProtoReflect.Descriptor instead.
func (*SetImageRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{1}
}

func (x *SetImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetImageResponse) Reset() {
	*x = SetImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetImageResponse) ProtoMessage() {}

func (x *SetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetImageResponse.ProtoReflect.Descriptor instead.
func (*SetImageResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{2}
}

type GetImagesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetImagesListRequest) Reset() {
	*x = GetImagesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesListRequest) ProtoMessage() {}

func (x *GetImagesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesListRequest.ProtoReflect.Descriptor instead.
func (*GetImagesListRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{3}
}

type GetImagesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetImagesListResponse) Reset() {
	*x = GetImagesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesListResponse) ProtoMessage() {}

func (x *GetImagesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesListResponse.ProtoReflect.Descriptor instead.
func (*GetImagesListResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{4}
}

func (x *GetImagesListResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{5}
}

func (x *GetImageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetImageResponse) Reset() {
	*x = GetImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_images_handler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageResponse) ProtoMessage() {}

func (x *GetImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_images_handler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageResponse.ProtoReflect.Descriptor instead.
func (*GetImageResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_images_handler_proto_rawDescGZIP(), []int{6}
}

func (x *GetImageResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_grpc_proto_images_handler_proto protoreflect.FileDescriptor

var file_grpc_proto_images_handler_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10,
	0x53, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x32, 0xd5, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_images_handler_proto_rawDescOnce sync.Once
	file_grpc_proto_images_handler_proto_rawDescData = file_grpc_proto_images_handler_proto_rawDesc
)

func file_grpc_proto_images_handler_proto_rawDescGZIP() []byte {
	file_grpc_proto_images_handler_proto_rawDescOnce.Do(func() {
		file_grpc_proto_images_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_images_handler_proto_rawDescData)
	})
	return file_grpc_proto_images_handler_proto_rawDescData
}

var file_grpc_proto_images_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_proto_images_handler_proto_goTypes = []interface{}{
	(*Image)(nil),                 // 0: proto.Image
	(*SetImageRequest)(nil),       // 1: proto.SetImageRequest
	(*SetImageResponse)(nil),      // 2: proto.SetImageResponse
	(*GetImagesListRequest)(nil),  // 3: proto.GetImagesListRequest
	(*GetImagesListResponse)(nil), // 4: proto.GetImagesListResponse
	(*GetImageRequest)(nil),       // 5: proto.GetImageRequest
	(*GetImageResponse)(nil),      // 6: proto.GetImageResponse
}
var file_grpc_proto_images_handler_proto_depIdxs = []int32{
	0, // 0: proto.GetImagesListResponse.image:type_name -> proto.Image
	0, // 1: proto.GetImageResponse.image:type_name -> proto.Image
	1, // 2: proto.ImagesHandler.SetImage:input_type -> proto.SetImageRequest
	3, // 3: proto.ImagesHandler.GetImagesList:input_type -> proto.GetImagesListRequest
	5, // 4: proto.ImagesHandler.GetImage:input_type -> proto.GetImageRequest
	2, // 5: proto.ImagesHandler.SetImage:output_type -> proto.SetImageResponse
	4, // 6: proto.ImagesHandler.GetImagesList:output_type -> proto.GetImagesListResponse
	6, // 7: proto.ImagesHandler.GetImage:output_type -> proto.GetImageResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_images_handler_proto_init() }
func file_grpc_proto_images_handler_proto_init() {
	if File_grpc_proto_images_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_images_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_grpc_proto_images_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetImageRequest); i {
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
		file_grpc_proto_images_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetImageResponse); i {
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
		file_grpc_proto_images_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesListRequest); i {
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
		file_grpc_proto_images_handler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesListResponse); i {
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
		file_grpc_proto_images_handler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
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
		file_grpc_proto_images_handler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageResponse); i {
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
			RawDescriptor: file_grpc_proto_images_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_images_handler_proto_goTypes,
		DependencyIndexes: file_grpc_proto_images_handler_proto_depIdxs,
		MessageInfos:      file_grpc_proto_images_handler_proto_msgTypes,
	}.Build()
	File_grpc_proto_images_handler_proto = out.File
	file_grpc_proto_images_handler_proto_rawDesc = nil
	file_grpc_proto_images_handler_proto_goTypes = nil
	file_grpc_proto_images_handler_proto_depIdxs = nil
}
