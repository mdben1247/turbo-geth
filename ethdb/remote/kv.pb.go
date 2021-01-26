// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: remote/kv.proto

package remote

import (
	proto "github.com/golang/protobuf/proto"
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

type Op int32

const (
	Op_FIRST           Op = 0
	Op_FIRST_DUP       Op = 1
	Op_SEEK            Op = 2
	Op_SEEK_BOTH       Op = 3
	Op_CURRENT         Op = 4
	Op_GET_MULTIPLE    Op = 5
	Op_LAST            Op = 6
	Op_LAST_DUP        Op = 7
	Op_NEXT            Op = 8
	Op_NEXT_DUP        Op = 9
	Op_NEXT_MULTIPLE   Op = 10
	Op_NEXT_NO_DUP     Op = 11
	Op_PREV            Op = 12
	Op_PREV_DUP        Op = 13
	Op_PREV_NO_DUP     Op = 14
	Op_SEEK_EXACT      Op = 15
	Op_SEEK_BOTH_EXACT Op = 16
	Op_OPEN            Op = 30
	Op_CLOSE           Op = 31
)

// Enum value maps for Op.
var (
	Op_name = map[int32]string{
		0:  "FIRST",
		1:  "FIRST_DUP",
		2:  "SEEK",
		3:  "SEEK_BOTH",
		4:  "CURRENT",
		5:  "GET_MULTIPLE",
		6:  "LAST",
		7:  "LAST_DUP",
		8:  "NEXT",
		9:  "NEXT_DUP",
		10: "NEXT_MULTIPLE",
		11: "NEXT_NO_DUP",
		12: "PREV",
		13: "PREV_DUP",
		14: "PREV_NO_DUP",
		15: "SEEK_EXACT",
		16: "SEEK_BOTH_EXACT",
		30: "OPEN",
		31: "CLOSE",
	}
	Op_value = map[string]int32{
		"FIRST":           0,
		"FIRST_DUP":       1,
		"SEEK":            2,
		"SEEK_BOTH":       3,
		"CURRENT":         4,
		"GET_MULTIPLE":    5,
		"LAST":            6,
		"LAST_DUP":        7,
		"NEXT":            8,
		"NEXT_DUP":        9,
		"NEXT_MULTIPLE":   10,
		"NEXT_NO_DUP":     11,
		"PREV":            12,
		"PREV_DUP":        13,
		"PREV_NO_DUP":     14,
		"SEEK_EXACT":      15,
		"SEEK_BOTH_EXACT": 16,
		"OPEN":            30,
		"CLOSE":           31,
	}
)

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}

func (x Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Op) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_kv_proto_enumTypes[0].Descriptor()
}

func (Op) Type() protoreflect.EnumType {
	return &file_remote_kv_proto_enumTypes[0]
}

func (x Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Op.Descriptor instead.
func (Op) EnumDescriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{0}
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op         Op     `protobuf:"varint,1,opt,name=op,proto3,enum=remote.Op" json:"op,omitempty"`
	BucketName string `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Cursor     uint32 `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	K          []byte `protobuf:"bytes,4,opt,name=k,proto3" json:"k,omitempty"`
	V          []byte `protobuf:"bytes,5,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{0}
}

func (x *Cursor) GetOp() Op {
	if x != nil {
		return x.Op
	}
	return Op_FIRST
}

func (x *Cursor) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *Cursor) GetCursor() uint32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *Cursor) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Cursor) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K        []byte `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	V        []byte `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
	CursorID uint32 `protobuf:"varint,3,opt,name=cursorID,proto3" json:"cursorID,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{1}
}

func (x *Pair) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Pair) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *Pair) GetCursorID() uint32 {
	if x != nil {
		return x.CursorID
	}
	return 0
}

var File_remote_kv_proto protoreflect.FileDescriptor

var file_remote_kv_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x78, 0x0a, 0x06, 0x43, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x01, 0x76, 0x22, 0x3e, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x49, 0x44, 0x2a, 0x8d, 0x02, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49,
	0x52, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x44,
	0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x45, 0x4b, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x45, 0x45, 0x4b, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45,
	0x54, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04,
	0x4c, 0x41, 0x53, 0x54, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x44,
	0x55, 0x50, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x58, 0x54, 0x10, 0x08, 0x12, 0x0c,
	0x0a, 0x08, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x10, 0x0a, 0x12,
	0x0f, 0x0a, 0x0b, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x0b,
	0x12, 0x08, 0x0a, 0x04, 0x50, 0x52, 0x45, 0x56, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52,
	0x45, 0x56, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x45, 0x56,
	0x5f, 0x4e, 0x4f, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x45,
	0x4b, 0x5f, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x45,
	0x4b, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x10, 0x12, 0x08,
	0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x1e, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x10, 0x1f, 0x32, 0x2c, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x26, 0x0a, 0x02, 0x54, 0x78, 0x12,
	0x0e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x1a,
	0x0c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x29, 0x0a, 0x10, 0x69, 0x6f, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x2d, 0x67, 0x65,
	0x74, 0x68, 0x2e, 0x64, 0x62, 0x42, 0x02, 0x4b, 0x56, 0x50, 0x01, 0x5a, 0x0f, 0x2e, 0x2f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_kv_proto_rawDescOnce sync.Once
	file_remote_kv_proto_rawDescData = file_remote_kv_proto_rawDesc
)

func file_remote_kv_proto_rawDescGZIP() []byte {
	file_remote_kv_proto_rawDescOnce.Do(func() {
		file_remote_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_kv_proto_rawDescData)
	})
	return file_remote_kv_proto_rawDescData
}

var file_remote_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remote_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_remote_kv_proto_goTypes = []interface{}{
	(Op)(0),        // 0: remote.Op
	(*Cursor)(nil), // 1: remote.Cursor
	(*Pair)(nil),   // 2: remote.Pair
}
var file_remote_kv_proto_depIdxs = []int32{
	0, // 0: remote.Cursor.op:type_name -> remote.Op
	1, // 1: remote.KV.Tx:input_type -> remote.Cursor
	2, // 2: remote.KV.Tx:output_type -> remote.Pair
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_remote_kv_proto_init() }
func file_remote_kv_proto_init() {
	if File_remote_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_remote_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
			RawDescriptor: file_remote_kv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_kv_proto_goTypes,
		DependencyIndexes: file_remote_kv_proto_depIdxs,
		EnumInfos:         file_remote_kv_proto_enumTypes,
		MessageInfos:      file_remote_kv_proto_msgTypes,
	}.Build()
	File_remote_kv_proto = out.File
	file_remote_kv_proto_rawDesc = nil
	file_remote_kv_proto_goTypes = nil
	file_remote_kv_proto_depIdxs = nil
}
