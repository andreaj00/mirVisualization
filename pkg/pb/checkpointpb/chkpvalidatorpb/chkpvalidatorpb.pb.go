// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: checkpointpb/chkpvalidatorpb/chkpvalidatorpb.proto

package chkpvalidatorpb

import (
	checkpointpb "github.com/filecoin-project/mir/pkg/pb/checkpointpb"
	contextstorepb "github.com/filecoin-project/mir/pkg/pb/contextstorepb"
	dslpb "github.com/filecoin-project/mir/pkg/pb/dslpb"
	_ "github.com/filecoin-project/mir/pkg/pb/mir"
	trantorpb "github.com/filecoin-project/mir/pkg/pb/trantorpb"
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
	//	*Event_ValidateCheckpoint
	//	*Event_CheckpointValidated
	Type isEvent_Type `protobuf_oneof:"type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[0]
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
	return file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetValidateCheckpoint() *ValidateCheckpoint {
	if x, ok := x.GetType().(*Event_ValidateCheckpoint); ok {
		return x.ValidateCheckpoint
	}
	return nil
}

func (x *Event) GetCheckpointValidated() *CheckpointValidated {
	if x, ok := x.GetType().(*Event_CheckpointValidated); ok {
		return x.CheckpointValidated
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_ValidateCheckpoint struct {
	ValidateCheckpoint *ValidateCheckpoint `protobuf:"bytes,1,opt,name=validate_checkpoint,json=validateCheckpoint,proto3,oneof"`
}

type Event_CheckpointValidated struct {
	CheckpointValidated *CheckpointValidated `protobuf:"bytes,2,opt,name=checkpoint_validated,json=checkpointValidated,proto3,oneof"`
}

func (*Event_ValidateCheckpoint) isEvent_Type() {}

func (*Event_CheckpointValidated) isEvent_Type() {}

type ValidateCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checkpoint  *checkpointpb.StableCheckpoint `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
	EpochNr     uint64                         `protobuf:"varint,2,opt,name=epoch_nr,json=epochNr,proto3" json:"epoch_nr,omitempty"`
	Memberships []*trantorpb.Membership        `protobuf:"bytes,3,rep,name=memberships,proto3" json:"memberships,omitempty"`
	Origin      *ValidateChkpOrigin            `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *ValidateCheckpoint) Reset() {
	*x = ValidateCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCheckpoint) ProtoMessage() {}

func (x *ValidateCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCheckpoint.ProtoReflect.Descriptor instead.
func (*ValidateCheckpoint) Descriptor() ([]byte, []int) {
	return file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateCheckpoint) GetCheckpoint() *checkpointpb.StableCheckpoint {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

func (x *ValidateCheckpoint) GetEpochNr() uint64 {
	if x != nil {
		return x.EpochNr
	}
	return 0
}

func (x *ValidateCheckpoint) GetMemberships() []*trantorpb.Membership {
	if x != nil {
		return x.Memberships
	}
	return nil
}

func (x *ValidateCheckpoint) GetOrigin() *ValidateChkpOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type CheckpointValidated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string              `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Origin *ValidateChkpOrigin `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *CheckpointValidated) Reset() {
	*x = CheckpointValidated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointValidated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointValidated) ProtoMessage() {}

func (x *CheckpointValidated) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointValidated.ProtoReflect.Descriptor instead.
func (*CheckpointValidated) Descriptor() ([]byte, []int) {
	return file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescGZIP(), []int{2}
}

func (x *CheckpointValidated) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CheckpointValidated) GetOrigin() *ValidateChkpOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type ValidateChkpOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// Types that are assignable to Type:
	//	*ValidateChkpOrigin_ContextStore
	//	*ValidateChkpOrigin_Dsl
	Type isValidateChkpOrigin_Type `protobuf_oneof:"type"`
}

func (x *ValidateChkpOrigin) Reset() {
	*x = ValidateChkpOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateChkpOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateChkpOrigin) ProtoMessage() {}

func (x *ValidateChkpOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateChkpOrigin.ProtoReflect.Descriptor instead.
func (*ValidateChkpOrigin) Descriptor() ([]byte, []int) {
	return file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateChkpOrigin) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (m *ValidateChkpOrigin) GetType() isValidateChkpOrigin_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ValidateChkpOrigin) GetContextStore() *contextstorepb.Origin {
	if x, ok := x.GetType().(*ValidateChkpOrigin_ContextStore); ok {
		return x.ContextStore
	}
	return nil
}

func (x *ValidateChkpOrigin) GetDsl() *dslpb.Origin {
	if x, ok := x.GetType().(*ValidateChkpOrigin_Dsl); ok {
		return x.Dsl
	}
	return nil
}

type isValidateChkpOrigin_Type interface {
	isValidateChkpOrigin_Type()
}

type ValidateChkpOrigin_ContextStore struct {
	ContextStore *contextstorepb.Origin `protobuf:"bytes,2,opt,name=context_store,json=contextStore,proto3,oneof"`
}

type ValidateChkpOrigin_Dsl struct {
	Dsl *dslpb.Origin `protobuf:"bytes,3,opt,name=dsl,proto3,oneof"`
}

func (*ValidateChkpOrigin_ContextStore) isValidateChkpOrigin_Type() {}

func (*ValidateChkpOrigin_Dsl) isValidateChkpOrigin_Type() {}

var File_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto protoreflect.FileDescriptor

var file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x63,
	0x68, 0x6b, 0x70, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63,
	0x68, 0x6b, 0x70, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x6b, 0x70, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x1a, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x64, 0x73, 0x6c, 0x70,
	0x62, 0x2f, 0x64, 0x73, 0x6c, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74,
	0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72,
	0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x56, 0x0a, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x68, 0x6b, 0x70, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x59, 0x0a, 0x14, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x6b, 0x70, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x13, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x3a, 0x04, 0x90, 0xa6, 0x1d, 0x01, 0x42, 0x0c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0xb0, 0x02, 0x0a, 0x12, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x3e, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x58, 0x0a, 0x08, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x6e, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x3d, 0x82, 0xa6, 0x1d, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x74,
	0x6f, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x72,
	0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x72, 0x12, 0x37, 0x0a, 0x0b, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x6b, 0x70, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x6b,
	0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x7f, 0x0a, 0x13, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0x82, 0xa6, 0x1d, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x6b, 0x70, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x6b, 0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x04, 0xa0, 0xa6, 0x1d, 0x01, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0xd4, 0x01, 0x0a,
	0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x6b, 0x70, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x4e, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x36, 0x82, 0xa6, 0x1d, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x64, 0x73, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x64, 0x73, 0x6c, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x48, 0x00,
	0x52, 0x03, 0x64, 0x73, 0x6c, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x6b, 0x70, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescOnce sync.Once
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescData = file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDesc
)

func file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescGZIP() []byte {
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescOnce.Do(func() {
		file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescData)
	})
	return file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDescData
}

var file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_goTypes = []interface{}{
	(*Event)(nil),                         // 0: chkpvalidatorpb.Event
	(*ValidateCheckpoint)(nil),            // 1: chkpvalidatorpb.ValidateCheckpoint
	(*CheckpointValidated)(nil),           // 2: chkpvalidatorpb.CheckpointValidated
	(*ValidateChkpOrigin)(nil),            // 3: chkpvalidatorpb.ValidateChkpOrigin
	(*checkpointpb.StableCheckpoint)(nil), // 4: checkpointpb.StableCheckpoint
	(*trantorpb.Membership)(nil),          // 5: trantorpb.Membership
	(*contextstorepb.Origin)(nil),         // 6: contextstorepb.Origin
	(*dslpb.Origin)(nil),                  // 7: dslpb.Origin
}
var file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_depIdxs = []int32{
	1, // 0: chkpvalidatorpb.Event.validate_checkpoint:type_name -> chkpvalidatorpb.ValidateCheckpoint
	2, // 1: chkpvalidatorpb.Event.checkpoint_validated:type_name -> chkpvalidatorpb.CheckpointValidated
	4, // 2: chkpvalidatorpb.ValidateCheckpoint.checkpoint:type_name -> checkpointpb.StableCheckpoint
	5, // 3: chkpvalidatorpb.ValidateCheckpoint.memberships:type_name -> trantorpb.Membership
	3, // 4: chkpvalidatorpb.ValidateCheckpoint.origin:type_name -> chkpvalidatorpb.ValidateChkpOrigin
	3, // 5: chkpvalidatorpb.CheckpointValidated.origin:type_name -> chkpvalidatorpb.ValidateChkpOrigin
	6, // 6: chkpvalidatorpb.ValidateChkpOrigin.context_store:type_name -> contextstorepb.Origin
	7, // 7: chkpvalidatorpb.ValidateChkpOrigin.dsl:type_name -> dslpb.Origin
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_init() }
func file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_init() {
	if File_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCheckpoint); i {
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
		file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointValidated); i {
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
		file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateChkpOrigin); i {
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
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_ValidateCheckpoint)(nil),
		(*Event_CheckpointValidated)(nil),
	}
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ValidateChkpOrigin_ContextStore)(nil),
		(*ValidateChkpOrigin_Dsl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_goTypes,
		DependencyIndexes: file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_depIdxs,
		MessageInfos:      file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_msgTypes,
	}.Build()
	File_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto = out.File
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_rawDesc = nil
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_goTypes = nil
	file_checkpointpb_chkpvalidatorpb_chkpvalidatorpb_proto_depIdxs = nil
}
