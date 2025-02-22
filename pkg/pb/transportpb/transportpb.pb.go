// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: transportpb/transportpb.proto

package transportpb

import (
	messagepb "github.com/filecoin-project/mir/pkg/pb/messagepb"
	_ "github.com/filecoin-project/mir/pkg/pb/mir"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Event_SendMessage
	//	*Event_MessageReceived
	Type isEvent_Type `protobuf_oneof:"Type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportpb_transportpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_transportpb_transportpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_transportpb_transportpb_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetSendMessage() *SendMessage {
	if x, ok := x.GetType().(*Event_SendMessage); ok {
		return x.SendMessage
	}
	return nil
}

func (x *Event) GetMessageReceived() *MessageReceived {
	if x, ok := x.GetType().(*Event_MessageReceived); ok {
		return x.MessageReceived
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_SendMessage struct {
	SendMessage *SendMessage `protobuf:"bytes,1,opt,name=send_message,json=sendMessage,proto3,oneof"`
}

type Event_MessageReceived struct {
	MessageReceived *MessageReceived `protobuf:"bytes,2,opt,name=message_received,json=messageReceived,proto3,oneof"`
}

func (*Event_SendMessage) isEvent_Type() {}

func (*Event_MessageReceived) isEvent_Type() {}

type SendMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg          *messagepb.Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Destinations []string           `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *SendMessage) Reset() {
	*x = SendMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportpb_transportpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessage) ProtoMessage() {}

func (x *SendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transportpb_transportpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessage.ProtoReflect.Descriptor instead.
func (*SendMessage) Descriptor() ([]byte, []int) {
	return file_transportpb_transportpb_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessage) GetMsg() *messagepb.Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SendMessage) GetDestinations() []string {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type MessageReceived struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string             `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Msg  *messagepb.Message `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MessageReceived) Reset() {
	*x = MessageReceived{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportpb_transportpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReceived) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReceived) ProtoMessage() {}

func (x *MessageReceived) ProtoReflect() protoreflect.Message {
	mi := &file_transportpb_transportpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReceived.ProtoReflect.Descriptor instead.
func (*MessageReceived) Descriptor() ([]byte, []int) {
	return file_transportpb_transportpb_proto_rawDescGZIP(), []int{2}
}

func (x *MessageReceived) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MessageReceived) GetMsg() *messagepb.Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_transportpb_transportpb_proto protoreflect.FileDescriptor

var file_transportpb_transportpb_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x3d, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x49,
	0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x3a, 0x04, 0x90, 0xa6, 0x1d, 0x01, 0x42,
	0x0c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0x93, 0x01,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x58, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x34, 0x82, 0xa6, 0x1d, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f,
	0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52,
	0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x04, 0x98,
	0xa6, 0x1d, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0x82, 0xa6, 0x1d, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x24, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transportpb_transportpb_proto_rawDescOnce sync.Once
	file_transportpb_transportpb_proto_rawDescData = file_transportpb_transportpb_proto_rawDesc
)

func file_transportpb_transportpb_proto_rawDescGZIP() []byte {
	file_transportpb_transportpb_proto_rawDescOnce.Do(func() {
		file_transportpb_transportpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_transportpb_transportpb_proto_rawDescData)
	})
	return file_transportpb_transportpb_proto_rawDescData
}

var file_transportpb_transportpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transportpb_transportpb_proto_goTypes = []interface{}{
	(*Event)(nil),             // 0: transportpb.Event
	(*SendMessage)(nil),       // 1: transportpb.SendMessage
	(*MessageReceived)(nil),   // 2: transportpb.MessageReceived
	(*messagepb.Message)(nil), // 3: messagepb.Message
}
var file_transportpb_transportpb_proto_depIdxs = []int32{
	1, // 0: transportpb.Event.send_message:type_name -> transportpb.SendMessage
	2, // 1: transportpb.Event.message_received:type_name -> transportpb.MessageReceived
	3, // 2: transportpb.SendMessage.msg:type_name -> messagepb.Message
	3, // 3: transportpb.MessageReceived.msg:type_name -> messagepb.Message
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_transportpb_transportpb_proto_init() }
func file_transportpb_transportpb_proto_init() {
	if File_transportpb_transportpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transportpb_transportpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_transportpb_transportpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessage); i {
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
		file_transportpb_transportpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReceived); i {
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
	file_transportpb_transportpb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_SendMessage)(nil),
		(*Event_MessageReceived)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transportpb_transportpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transportpb_transportpb_proto_goTypes,
		DependencyIndexes: file_transportpb_transportpb_proto_depIdxs,
		MessageInfos:      file_transportpb_transportpb_proto_msgTypes,
	}.Build()
	File_transportpb_transportpb_proto = out.File
	file_transportpb_transportpb_proto_rawDesc = nil
	file_transportpb_transportpb_proto_goTypes = nil
	file_transportpb_transportpb_proto_depIdxs = nil
}
