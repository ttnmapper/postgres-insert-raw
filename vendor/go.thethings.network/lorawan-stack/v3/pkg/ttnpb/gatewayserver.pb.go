// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: lorawan-stack/api/gatewayserver.proto

package ttnpb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GatewayUp may contain zero or more uplink messages and/or a status message for the gateway.
type GatewayUp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uplink messages received by the gateway.
	UplinkMessages []*UplinkMessage `protobuf:"bytes,1,rep,name=uplink_messages,json=uplinkMessages,proto3" json:"uplink_messages,omitempty"`
	// Gateway status produced by the gateway.
	GatewayStatus *GatewayStatus `protobuf:"bytes,2,opt,name=gateway_status,json=gatewayStatus,proto3" json:"gateway_status,omitempty"`
	// A Tx acknowledgment or error.
	TxAcknowledgment *TxAcknowledgment `protobuf:"bytes,3,opt,name=tx_acknowledgment,json=txAcknowledgment,proto3" json:"tx_acknowledgment,omitempty"`
}

func (x *GatewayUp) Reset() {
	*x = GatewayUp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayUp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayUp) ProtoMessage() {}

func (x *GatewayUp) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayUp.ProtoReflect.Descriptor instead.
func (*GatewayUp) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayUp) GetUplinkMessages() []*UplinkMessage {
	if x != nil {
		return x.UplinkMessages
	}
	return nil
}

func (x *GatewayUp) GetGatewayStatus() *GatewayStatus {
	if x != nil {
		return x.GatewayStatus
	}
	return nil
}

func (x *GatewayUp) GetTxAcknowledgment() *TxAcknowledgment {
	if x != nil {
		return x.TxAcknowledgment
	}
	return nil
}

// GatewayDown contains downlink messages for the gateway.
type GatewayDown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DownlinkMessage for the gateway.
	DownlinkMessage *DownlinkMessage `protobuf:"bytes,1,opt,name=downlink_message,json=downlinkMessage,proto3" json:"downlink_message,omitempty"`
}

func (x *GatewayDown) Reset() {
	*x = GatewayDown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayDown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayDown) ProtoMessage() {}

func (x *GatewayDown) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayDown.ProtoReflect.Descriptor instead.
func (*GatewayDown) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayDown) GetDownlinkMessage() *DownlinkMessage {
	if x != nil {
		return x.DownlinkMessage
	}
	return nil
}

type ScheduleDownlinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of time between the message has been scheduled and it will be transmitted by the gateway.
	Delay *durationpb.Duration `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// Downlink path chosen by the Gateway Server.
	DownlinkPath *DownlinkPath `protobuf:"bytes,2,opt,name=downlink_path,json=downlinkPath,proto3" json:"downlink_path,omitempty"`
	// Whether RX1 has been chosen for the downlink message.
	// Both RX1 and RX2 can be used for transmitting the same message by the same gateway.
	Rx1 bool `protobuf:"varint,3,opt,name=rx1,proto3" json:"rx1,omitempty"`
	// Whether RX2 has been chosen for the downlink message.
	// Both RX1 and RX2 can be used for transmitting the same message by the same gateway.
	Rx2 bool `protobuf:"varint,4,opt,name=rx2,proto3" json:"rx2,omitempty"`
}

func (x *ScheduleDownlinkResponse) Reset() {
	*x = ScheduleDownlinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleDownlinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleDownlinkResponse) ProtoMessage() {}

func (x *ScheduleDownlinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleDownlinkResponse.ProtoReflect.Descriptor instead.
func (*ScheduleDownlinkResponse) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleDownlinkResponse) GetDelay() *durationpb.Duration {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *ScheduleDownlinkResponse) GetDownlinkPath() *DownlinkPath {
	if x != nil {
		return x.DownlinkPath
	}
	return nil
}

func (x *ScheduleDownlinkResponse) GetRx1() bool {
	if x != nil {
		return x.Rx1
	}
	return false
}

func (x *ScheduleDownlinkResponse) GetRx2() bool {
	if x != nil {
		return x.Rx2
	}
	return false
}

type ScheduleDownlinkErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors per path when downlink scheduling failed.
	PathErrors []*ErrorDetails `protobuf:"bytes,1,rep,name=path_errors,json=pathErrors,proto3" json:"path_errors,omitempty"`
}

func (x *ScheduleDownlinkErrorDetails) Reset() {
	*x = ScheduleDownlinkErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleDownlinkErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleDownlinkErrorDetails) ProtoMessage() {}

func (x *ScheduleDownlinkErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleDownlinkErrorDetails.ProtoReflect.Descriptor instead.
func (*ScheduleDownlinkErrorDetails) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP(), []int{3}
}

func (x *ScheduleDownlinkErrorDetails) GetPathErrors() []*ErrorDetails {
	if x != nil {
		return x.PathErrors
	}
	return nil
}

type BatchGetGatewayConnectionStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayIds []*GatewayIdentifiers `protobuf:"bytes,1,rep,name=gateway_ids,json=gatewayIds,proto3" json:"gateway_ids,omitempty"`
	// The names of the gateway stats fields that should be returned.
	// This mask will be applied on each entry returned.
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (x *BatchGetGatewayConnectionStatsRequest) Reset() {
	*x = BatchGetGatewayConnectionStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetGatewayConnectionStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetGatewayConnectionStatsRequest) ProtoMessage() {}

func (x *BatchGetGatewayConnectionStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetGatewayConnectionStatsRequest.ProtoReflect.Descriptor instead.
func (*BatchGetGatewayConnectionStatsRequest) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP(), []int{4}
}

func (x *BatchGetGatewayConnectionStatsRequest) GetGatewayIds() []*GatewayIdentifiers {
	if x != nil {
		return x.GatewayIds
	}
	return nil
}

func (x *BatchGetGatewayConnectionStatsRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

type BatchGetGatewayConnectionStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The map key is the gateway identifier.
	Entries map[string]*GatewayConnectionStats `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BatchGetGatewayConnectionStatsResponse) Reset() {
	*x = BatchGetGatewayConnectionStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetGatewayConnectionStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetGatewayConnectionStatsResponse) ProtoMessage() {}

func (x *BatchGetGatewayConnectionStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gatewayserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetGatewayConnectionStatsResponse.ProtoReflect.Descriptor instead.
func (*BatchGetGatewayConnectionStatsResponse) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP(), []int{5}
}

func (x *BatchGetGatewayConnectionStatsResponse) GetEntries() map[string]*GatewayConnectionStats {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_lorawan_stack_api_gatewayserver_proto protoreflect.FileDescriptor

var file_lorawan_stack_api_gatewayserver_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x70, 0x12, 0x46, 0x0a,
	0x0f, 0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a, 0x11, 0x74,
	0x78, 0x5f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x78, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x74, 0x78, 0x41, 0x63, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x0b, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x4a, 0x0a, 0x10, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x18, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0xaa, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x41, 0x0a,
	0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x78, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x72,
	0x78, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x78, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x72, 0x78, 0x32, 0x22, 0x5d, 0x0a, 0x1c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x25, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a,
	0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x92, 0x01, 0x04, 0x08, 0x01,
	0x10, 0x64, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xeb, 0x01, 0x0a, 0x26, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x1a, 0x62, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xdf, 0x03, 0x0a, 0x05, 0x47, 0x74, 0x77, 0x47,
	0x73, 0x12, 0x49, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0x19, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x70, 0x1a, 0x1b, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x28, 0x01, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x97, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x51, 0x54, 0x54, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a,
	0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x4d, 0x51, 0x54, 0x54, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x67, 0x73,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x9b, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x4d, 0x51, 0x54, 0x54, 0x56, 0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x22, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x51, 0x54,
	0x54, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x67, 0x73, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x76, 0x32, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x65, 0x0a, 0x04, 0x4e, 0x73, 0x47,
	0x73, 0x12, 0x5d, 0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x28, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xde, 0x02, 0x0a, 0x02, 0x47, 0x73, 0x12, 0x9b, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x26, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x67, 0x73, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0xb9, 0x01, 0x0a, 0x1e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x35, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22,
	0x1d, 0x2f, 0x67, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x3a, 0x01,
	0x2a, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74,
	0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lorawan_stack_api_gatewayserver_proto_rawDescOnce sync.Once
	file_lorawan_stack_api_gatewayserver_proto_rawDescData = file_lorawan_stack_api_gatewayserver_proto_rawDesc
)

func file_lorawan_stack_api_gatewayserver_proto_rawDescGZIP() []byte {
	file_lorawan_stack_api_gatewayserver_proto_rawDescOnce.Do(func() {
		file_lorawan_stack_api_gatewayserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_lorawan_stack_api_gatewayserver_proto_rawDescData)
	})
	return file_lorawan_stack_api_gatewayserver_proto_rawDescData
}

var file_lorawan_stack_api_gatewayserver_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lorawan_stack_api_gatewayserver_proto_goTypes = []interface{}{
	(*GatewayUp)(nil),                              // 0: ttn.lorawan.v3.GatewayUp
	(*GatewayDown)(nil),                            // 1: ttn.lorawan.v3.GatewayDown
	(*ScheduleDownlinkResponse)(nil),               // 2: ttn.lorawan.v3.ScheduleDownlinkResponse
	(*ScheduleDownlinkErrorDetails)(nil),           // 3: ttn.lorawan.v3.ScheduleDownlinkErrorDetails
	(*BatchGetGatewayConnectionStatsRequest)(nil),  // 4: ttn.lorawan.v3.BatchGetGatewayConnectionStatsRequest
	(*BatchGetGatewayConnectionStatsResponse)(nil), // 5: ttn.lorawan.v3.BatchGetGatewayConnectionStatsResponse
	nil,                            // 6: ttn.lorawan.v3.BatchGetGatewayConnectionStatsResponse.EntriesEntry
	(*UplinkMessage)(nil),          // 7: ttn.lorawan.v3.UplinkMessage
	(*GatewayStatus)(nil),          // 8: ttn.lorawan.v3.GatewayStatus
	(*TxAcknowledgment)(nil),       // 9: ttn.lorawan.v3.TxAcknowledgment
	(*DownlinkMessage)(nil),        // 10: ttn.lorawan.v3.DownlinkMessage
	(*durationpb.Duration)(nil),    // 11: google.protobuf.Duration
	(*DownlinkPath)(nil),           // 12: ttn.lorawan.v3.DownlinkPath
	(*ErrorDetails)(nil),           // 13: ttn.lorawan.v3.ErrorDetails
	(*GatewayIdentifiers)(nil),     // 14: ttn.lorawan.v3.GatewayIdentifiers
	(*fieldmaskpb.FieldMask)(nil),  // 15: google.protobuf.FieldMask
	(*GatewayConnectionStats)(nil), // 16: ttn.lorawan.v3.GatewayConnectionStats
	(*emptypb.Empty)(nil),          // 17: google.protobuf.Empty
	(*ConcentratorConfig)(nil),     // 18: ttn.lorawan.v3.ConcentratorConfig
	(*MQTTConnectionInfo)(nil),     // 19: ttn.lorawan.v3.MQTTConnectionInfo
}
var file_lorawan_stack_api_gatewayserver_proto_depIdxs = []int32{
	7,  // 0: ttn.lorawan.v3.GatewayUp.uplink_messages:type_name -> ttn.lorawan.v3.UplinkMessage
	8,  // 1: ttn.lorawan.v3.GatewayUp.gateway_status:type_name -> ttn.lorawan.v3.GatewayStatus
	9,  // 2: ttn.lorawan.v3.GatewayUp.tx_acknowledgment:type_name -> ttn.lorawan.v3.TxAcknowledgment
	10, // 3: ttn.lorawan.v3.GatewayDown.downlink_message:type_name -> ttn.lorawan.v3.DownlinkMessage
	11, // 4: ttn.lorawan.v3.ScheduleDownlinkResponse.delay:type_name -> google.protobuf.Duration
	12, // 5: ttn.lorawan.v3.ScheduleDownlinkResponse.downlink_path:type_name -> ttn.lorawan.v3.DownlinkPath
	13, // 6: ttn.lorawan.v3.ScheduleDownlinkErrorDetails.path_errors:type_name -> ttn.lorawan.v3.ErrorDetails
	14, // 7: ttn.lorawan.v3.BatchGetGatewayConnectionStatsRequest.gateway_ids:type_name -> ttn.lorawan.v3.GatewayIdentifiers
	15, // 8: ttn.lorawan.v3.BatchGetGatewayConnectionStatsRequest.field_mask:type_name -> google.protobuf.FieldMask
	6,  // 9: ttn.lorawan.v3.BatchGetGatewayConnectionStatsResponse.entries:type_name -> ttn.lorawan.v3.BatchGetGatewayConnectionStatsResponse.EntriesEntry
	16, // 10: ttn.lorawan.v3.BatchGetGatewayConnectionStatsResponse.EntriesEntry.value:type_name -> ttn.lorawan.v3.GatewayConnectionStats
	0,  // 11: ttn.lorawan.v3.GtwGs.LinkGateway:input_type -> ttn.lorawan.v3.GatewayUp
	17, // 12: ttn.lorawan.v3.GtwGs.GetConcentratorConfig:input_type -> google.protobuf.Empty
	14, // 13: ttn.lorawan.v3.GtwGs.GetMQTTConnectionInfo:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	14, // 14: ttn.lorawan.v3.GtwGs.GetMQTTV2ConnectionInfo:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	10, // 15: ttn.lorawan.v3.NsGs.ScheduleDownlink:input_type -> ttn.lorawan.v3.DownlinkMessage
	14, // 16: ttn.lorawan.v3.Gs.GetGatewayConnectionStats:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	4,  // 17: ttn.lorawan.v3.Gs.BatchGetGatewayConnectionStats:input_type -> ttn.lorawan.v3.BatchGetGatewayConnectionStatsRequest
	1,  // 18: ttn.lorawan.v3.GtwGs.LinkGateway:output_type -> ttn.lorawan.v3.GatewayDown
	18, // 19: ttn.lorawan.v3.GtwGs.GetConcentratorConfig:output_type -> ttn.lorawan.v3.ConcentratorConfig
	19, // 20: ttn.lorawan.v3.GtwGs.GetMQTTConnectionInfo:output_type -> ttn.lorawan.v3.MQTTConnectionInfo
	19, // 21: ttn.lorawan.v3.GtwGs.GetMQTTV2ConnectionInfo:output_type -> ttn.lorawan.v3.MQTTConnectionInfo
	2,  // 22: ttn.lorawan.v3.NsGs.ScheduleDownlink:output_type -> ttn.lorawan.v3.ScheduleDownlinkResponse
	16, // 23: ttn.lorawan.v3.Gs.GetGatewayConnectionStats:output_type -> ttn.lorawan.v3.GatewayConnectionStats
	5,  // 24: ttn.lorawan.v3.Gs.BatchGetGatewayConnectionStats:output_type -> ttn.lorawan.v3.BatchGetGatewayConnectionStatsResponse
	18, // [18:25] is the sub-list for method output_type
	11, // [11:18] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_lorawan_stack_api_gatewayserver_proto_init() }
func file_lorawan_stack_api_gatewayserver_proto_init() {
	if File_lorawan_stack_api_gatewayserver_proto != nil {
		return
	}
	file_lorawan_stack_api_error_proto_init()
	file_lorawan_stack_api_gateway_proto_init()
	file_lorawan_stack_api_identifiers_proto_init()
	file_lorawan_stack_api_lorawan_proto_init()
	file_lorawan_stack_api_messages_proto_init()
	file_lorawan_stack_api_mqtt_proto_init()
	file_lorawan_stack_api_regional_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lorawan_stack_api_gatewayserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayUp); i {
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
		file_lorawan_stack_api_gatewayserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayDown); i {
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
		file_lorawan_stack_api_gatewayserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleDownlinkResponse); i {
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
		file_lorawan_stack_api_gatewayserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleDownlinkErrorDetails); i {
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
		file_lorawan_stack_api_gatewayserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetGatewayConnectionStatsRequest); i {
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
		file_lorawan_stack_api_gatewayserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetGatewayConnectionStatsResponse); i {
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
			RawDescriptor: file_lorawan_stack_api_gatewayserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_lorawan_stack_api_gatewayserver_proto_goTypes,
		DependencyIndexes: file_lorawan_stack_api_gatewayserver_proto_depIdxs,
		MessageInfos:      file_lorawan_stack_api_gatewayserver_proto_msgTypes,
	}.Build()
	File_lorawan_stack_api_gatewayserver_proto = out.File
	file_lorawan_stack_api_gatewayserver_proto_rawDesc = nil
	file_lorawan_stack_api_gatewayserver_proto_goTypes = nil
	file_lorawan_stack_api_gatewayserver_proto_depIdxs = nil
}