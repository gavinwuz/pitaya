// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: arg.proto

package protos

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

type Arg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Arg) Reset() {
	*x = Arg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Arg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Arg) ProtoMessage() {}

func (x *Arg) ProtoReflect() protoreflect.Message {
	mi := &file_arg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Arg.ProtoReflect.Descriptor instead.
func (*Arg) Descriptor() ([]byte, []int) {
	return file_arg_proto_rawDescGZIP(), []int{0}
}

func (x *Arg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_arg_proto protoreflect.FileDescriptor

var file_arg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x17, 0x0a, 0x03, 0x41, 0x72,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arg_proto_rawDescOnce sync.Once
	file_arg_proto_rawDescData = file_arg_proto_rawDesc
)

func file_arg_proto_rawDescGZIP() []byte {
	file_arg_proto_rawDescOnce.Do(func() {
		file_arg_proto_rawDescData = protoimpl.X.CompressGZIP(file_arg_proto_rawDescData)
	})
	return file_arg_proto_rawDescData
}

var file_arg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_arg_proto_goTypes = []interface{}{
	(*Arg)(nil), // 0: worker_protos.Arg
}
var file_arg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arg_proto_init() }
func file_arg_proto_init() {
	if File_arg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Arg); i {
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
			RawDescriptor: file_arg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_arg_proto_goTypes,
		DependencyIndexes: file_arg_proto_depIdxs,
		MessageInfos:      file_arg_proto_msgTypes,
	}.Build()
	File_arg_proto = out.File
	file_arg_proto_rawDesc = nil
	file_arg_proto_goTypes = nil
	file_arg_proto_depIdxs = nil
}
