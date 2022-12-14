// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: tages_service/images/v1/image.proto

package pb_tages_service_images

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

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size       uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	ImageBytes []byte `protobuf:"bytes,4,opt,name=image_bytes,json=imageBytes,proto3" json:"image_bytes,omitempty"`
	CreatedAt  int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tages_service_images_v1_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_tages_service_images_v1_image_proto_msgTypes[0]
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
	return file_tages_service_images_v1_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Image) GetImageBytes() []byte {
	if x != nil {
		return x.ImageBytes
	}
	return nil
}

func (x *Image) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Image) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_tages_service_images_v1_image_proto protoreflect.FileDescriptor

var file_tages_service_images_v1_image_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x9e,
	0x01, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x61, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65,
	0x72, 0x6d, 0x61, 0x6e, 0x42, 0x6f, 0x67, 0x61, 0x74, 0x6f, 0x76, 0x2f, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x74, 0x61,
	0x67, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tages_service_images_v1_image_proto_rawDescOnce sync.Once
	file_tages_service_images_v1_image_proto_rawDescData = file_tages_service_images_v1_image_proto_rawDesc
)

func file_tages_service_images_v1_image_proto_rawDescGZIP() []byte {
	file_tages_service_images_v1_image_proto_rawDescOnce.Do(func() {
		file_tages_service_images_v1_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_tages_service_images_v1_image_proto_rawDescData)
	})
	return file_tages_service_images_v1_image_proto_rawDescData
}

var file_tages_service_images_v1_image_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tages_service_images_v1_image_proto_goTypes = []interface{}{
	(*Image)(nil), // 0: tages_service.images.v1.Image
}
var file_tages_service_images_v1_image_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tages_service_images_v1_image_proto_init() }
func file_tages_service_images_v1_image_proto_init() {
	if File_tages_service_images_v1_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tages_service_images_v1_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tages_service_images_v1_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tages_service_images_v1_image_proto_goTypes,
		DependencyIndexes: file_tages_service_images_v1_image_proto_depIdxs,
		MessageInfos:      file_tages_service_images_v1_image_proto_msgTypes,
	}.Build()
	File_tages_service_images_v1_image_proto = out.File
	file_tages_service_images_v1_image_proto_rawDesc = nil
	file_tages_service_images_v1_image_proto_goTypes = nil
	file_tages_service_images_v1_image_proto_depIdxs = nil
}
