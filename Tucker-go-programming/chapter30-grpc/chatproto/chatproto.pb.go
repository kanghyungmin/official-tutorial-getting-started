// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: chatproto.proto

package chatproto

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

type ChatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMsg) Reset() {
	*x = ChatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMsg) ProtoMessage() {}

func (x *ChatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chatproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMsg.ProtoReflect.Descriptor instead.
func (*ChatMsg) Descriptor() ([]byte, []int) {
	return file_chatproto_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMsg) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_chatproto_proto protoreflect.FileDescriptor

var file_chatproto_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x43, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x73, 0x67, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatproto_proto_rawDescOnce sync.Once
	file_chatproto_proto_rawDescData = file_chatproto_proto_rawDesc
)

func file_chatproto_proto_rawDescGZIP() []byte {
	file_chatproto_proto_rawDescOnce.Do(func() {
		file_chatproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatproto_proto_rawDescData)
	})
	return file_chatproto_proto_rawDescData
}

var file_chatproto_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chatproto_proto_goTypes = []any{
	(*ChatMsg)(nil), // 0: chatproto.ChatMsg
}
var file_chatproto_proto_depIdxs = []int32{
	0, // 0: chatproto.ChatService.Chat:input_type -> chatproto.ChatMsg
	0, // 1: chatproto.ChatService.Chat:output_type -> chatproto.ChatMsg
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatproto_proto_init() }
func file_chatproto_proto_init() {
	if File_chatproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatproto_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChatMsg); i {
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
			RawDescriptor: file_chatproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatproto_proto_goTypes,
		DependencyIndexes: file_chatproto_proto_depIdxs,
		MessageInfos:      file_chatproto_proto_msgTypes,
	}.Build()
	File_chatproto_proto = out.File
	file_chatproto_proto_rawDesc = nil
	file_chatproto_proto_goTypes = nil
	file_chatproto_proto_depIdxs = nil
}
