// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: aocfinder.proto

package proto

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

type WorkingOn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year               int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Day                int32 `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Part               int32 `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
	CorrespondingIssue int32 `protobuf:"varint,4,opt,name=corresponding_issue,json=correspondingIssue,proto3" json:"corresponding_issue,omitempty"`
}

func (x *WorkingOn) Reset() {
	*x = WorkingOn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aocfinder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkingOn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkingOn) ProtoMessage() {}

func (x *WorkingOn) ProtoReflect() protoreflect.Message {
	mi := &file_aocfinder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkingOn.ProtoReflect.Descriptor instead.
func (*WorkingOn) Descriptor() ([]byte, []int) {
	return file_aocfinder_proto_rawDescGZIP(), []int{0}
}

func (x *WorkingOn) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *WorkingOn) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *WorkingOn) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *WorkingOn) GetCorrespondingIssue() int32 {
	if x != nil {
		return x.CorrespondingIssue
	}
	return 0
}

var File_aocfinder_proto protoreflect.FileDescriptor

var file_aocfinder_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6f, 0x63, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x6f, 0x63, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x09,
	0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f,
	0x61, 0x6f, 0x63, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aocfinder_proto_rawDescOnce sync.Once
	file_aocfinder_proto_rawDescData = file_aocfinder_proto_rawDesc
)

func file_aocfinder_proto_rawDescGZIP() []byte {
	file_aocfinder_proto_rawDescOnce.Do(func() {
		file_aocfinder_proto_rawDescData = protoimpl.X.CompressGZIP(file_aocfinder_proto_rawDescData)
	})
	return file_aocfinder_proto_rawDescData
}

var file_aocfinder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aocfinder_proto_goTypes = []interface{}{
	(*WorkingOn)(nil), // 0: aocfinder.WorkingOn
}
var file_aocfinder_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aocfinder_proto_init() }
func file_aocfinder_proto_init() {
	if File_aocfinder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aocfinder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkingOn); i {
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
			RawDescriptor: file_aocfinder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aocfinder_proto_goTypes,
		DependencyIndexes: file_aocfinder_proto_depIdxs,
		MessageInfos:      file_aocfinder_proto_msgTypes,
	}.Build()
	File_aocfinder_proto = out.File
	file_aocfinder_proto_rawDesc = nil
	file_aocfinder_proto_goTypes = nil
	file_aocfinder_proto_depIdxs = nil
}
