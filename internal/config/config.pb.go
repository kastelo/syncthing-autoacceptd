// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/config.proto

package config

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

type Mode int32

const (
	Mode_SENDRECEIVE Mode = 0
	Mode_SENDONLY    Mode = 1
	Mode_RECEIVEONLY Mode = 2
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "SENDRECEIVE",
		1: "SENDONLY",
		2: "RECEIVEONLY",
	}
	Mode_value = map[string]int32{
		"SENDRECEIVE": 0,
		"SENDONLY":    1,
		"RECEIVEONLY": 2,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_config_proto_enumTypes[0].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_proto_config_proto_enumTypes[0]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0}
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern []*DevicePattern `protobuf:"bytes,1,rep,name=pattern,proto3" json:"pattern,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetPattern() []*DevicePattern {
	if x != nil {
		return x.Pattern
	}
	return nil
}

type DevicePattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folder     []*FolderPattern `protobuf:"bytes,1,rep,name=folder,proto3" json:"folder,omitempty"`
	AcceptCidr []string         `protobuf:"bytes,2,rep,name=accept_cidr,json=acceptCidr,proto3" json:"accept_cidr,omitempty"`
}

func (x *DevicePattern) Reset() {
	*x = DevicePattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicePattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicePattern) ProtoMessage() {}

func (x *DevicePattern) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicePattern.ProtoReflect.Descriptor instead.
func (*DevicePattern) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *DevicePattern) GetFolder() []*FolderPattern {
	if x != nil {
		return x.Folder
	}
	return nil
}

func (x *DevicePattern) GetAcceptCidr() []string {
	if x != nil {
		return x.AcceptCidr
	}
	return nil
}

type FolderPattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mode Mode   `protobuf:"varint,2,opt,name=mode,proto3,enum=config.Mode" json:"mode,omitempty"`
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FolderPattern) Reset() {
	*x = FolderPattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FolderPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FolderPattern) ProtoMessage() {}

func (x *FolderPattern) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FolderPattern.ProtoReflect.Descriptor instead.
func (*FolderPattern) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{2}
}

func (x *FolderPattern) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FolderPattern) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_SENDRECEIVE
}

func (x *FolderPattern) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_proto_config_proto protoreflect.FileDescriptor

var file_proto_config_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x40, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x22, 0x5f,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12,
	0x2d, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x69, 0x64, 0x72, 0x22,
	0x55, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x36, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x45, 0x4e, 0x44, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x45, 0x4e, 0x44, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x42, 0x33,
	0x5a, 0x31, 0x6b, 0x61, 0x73, 0x74, 0x65, 0x6c, 0x6f, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x79,
	0x6e, 0x63, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x75, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_proto_rawDescOnce sync.Once
	file_proto_config_proto_rawDescData = file_proto_config_proto_rawDesc
)

func file_proto_config_proto_rawDescGZIP() []byte {
	file_proto_config_proto_rawDescOnce.Do(func() {
		file_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_proto_rawDescData)
	})
	return file_proto_config_proto_rawDescData
}

var file_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_config_proto_goTypes = []interface{}{
	(Mode)(0),             // 0: config.Mode
	(*Configuration)(nil), // 1: config.Configuration
	(*DevicePattern)(nil), // 2: config.DevicePattern
	(*FolderPattern)(nil), // 3: config.FolderPattern
}
var file_proto_config_proto_depIdxs = []int32{
	2, // 0: config.Configuration.pattern:type_name -> config.DevicePattern
	3, // 1: config.DevicePattern.folder:type_name -> config.FolderPattern
	0, // 2: config.FolderPattern.mode:type_name -> config.Mode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_config_proto_init() }
func file_proto_config_proto_init() {
	if File_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicePattern); i {
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
		file_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FolderPattern); i {
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
			RawDescriptor: file_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_proto_goTypes,
		DependencyIndexes: file_proto_config_proto_depIdxs,
		EnumInfos:         file_proto_config_proto_enumTypes,
		MessageInfos:      file_proto_config_proto_msgTypes,
	}.Build()
	File_proto_config_proto = out.File
	file_proto_config_proto_rawDesc = nil
	file_proto_config_proto_goTypes = nil
	file_proto_config_proto_depIdxs = nil
}
