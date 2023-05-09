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
// source: lorawan-stack/api/networkserver.proto

package ttnpb

import (
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Response of GenerateDevAddr.
type GenerateDevAddrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevAddr []byte `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3" json:"dev_addr,omitempty"`
}

func (x *GenerateDevAddrResponse) Reset() {
	*x = GenerateDevAddrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDevAddrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDevAddrResponse) ProtoMessage() {}

func (x *GenerateDevAddrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDevAddrResponse.ProtoReflect.Descriptor instead.
func (*GenerateDevAddrResponse) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_networkserver_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateDevAddrResponse) GetDevAddr() []byte {
	if x != nil {
		return x.DevAddr
	}
	return nil
}

// Request of GetDefaultMACSettings.
type GetDefaultMACSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrequencyPlanId   string     `protobuf:"bytes,1,opt,name=frequency_plan_id,json=frequencyPlanId,proto3" json:"frequency_plan_id,omitempty"`
	LorawanPhyVersion PHYVersion `protobuf:"varint,2,opt,name=lorawan_phy_version,json=lorawanPhyVersion,proto3,enum=ttn.lorawan.v3.PHYVersion" json:"lorawan_phy_version,omitempty"`
}

func (x *GetDefaultMACSettingsRequest) Reset() {
	*x = GetDefaultMACSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultMACSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultMACSettingsRequest) ProtoMessage() {}

func (x *GetDefaultMACSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultMACSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetDefaultMACSettingsRequest) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_networkserver_proto_rawDescGZIP(), []int{1}
}

func (x *GetDefaultMACSettingsRequest) GetFrequencyPlanId() string {
	if x != nil {
		return x.FrequencyPlanId
	}
	return ""
}

func (x *GetDefaultMACSettingsRequest) GetLorawanPhyVersion() PHYVersion {
	if x != nil {
		return x.LorawanPhyVersion
	}
	return PHYVersion_PHY_UNKNOWN
}

type GetNetIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetId []byte `protobuf:"bytes,1,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
}

func (x *GetNetIDResponse) Reset() {
	*x = GetNetIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetIDResponse) ProtoMessage() {}

func (x *GetNetIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetIDResponse.ProtoReflect.Descriptor instead.
func (*GetNetIDResponse) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_networkserver_proto_rawDescGZIP(), []int{2}
}

func (x *GetNetIDResponse) GetNetId() []byte {
	if x != nil {
		return x.NetId
	}
	return nil
}

type GetDeviceAdressPrefixesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevAddrPrefixes [][]byte `protobuf:"bytes,1,rep,name=dev_addr_prefixes,json=devAddrPrefixes,proto3" json:"dev_addr_prefixes,omitempty"`
}

func (x *GetDeviceAdressPrefixesResponse) Reset() {
	*x = GetDeviceAdressPrefixesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceAdressPrefixesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceAdressPrefixesResponse) ProtoMessage() {}

func (x *GetDeviceAdressPrefixesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_networkserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceAdressPrefixesResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceAdressPrefixesResponse) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_networkserver_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceAdressPrefixesResponse) GetDevAddrPrefixes() [][]byte {
	if x != nil {
		return x.DevAddrPrefixes
	}
	return nil
}

var File_lorawan_stack_api_networkserver_proto protoreflect.FileDescriptor

var file_lorawan_stack_api_networkserver_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64,
	0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a,
	0x73, 0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xc8, 0x01, 0x0a, 0x08, 0x64, 0x65, 0x76,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0xac, 0x01, 0xfa, 0x42,
	0x06, 0x7a, 0x04, 0x68, 0x04, 0x70, 0x01, 0xea, 0xaa, 0x19, 0x82, 0x01, 0x0a, 0x3f, 0x67, 0x6f,
	0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3f, 0x67,
	0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55,
	0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x34, 0x42, 0x79, 0x74, 0x65, 0x73, 0x92, 0x41,
	0x19, 0x4a, 0x0a, 0x22, 0x32, 0x36, 0x30, 0x30, 0x41, 0x42, 0x43, 0x44, 0x22, 0x9a, 0x02, 0x01,
	0x07, 0xa2, 0x02, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x64, 0x65, 0x76, 0x41,
	0x64, 0x64, 0x72, 0x22, 0xb3, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4d, 0x41, 0x43, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x11, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x0f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x13, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x5f, 0x70, 0x68, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x48, 0x59, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x11, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x50, 0x68, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a,
	0x08, 0xf2, 0xaa, 0x19, 0x04, 0x08, 0x00, 0x10, 0x01, 0x22, 0xd7, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xc2,
	0x01, 0x0a, 0x06, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0xaa, 0x01, 0xfa, 0x42, 0x06, 0x7a, 0x04, 0x68, 0x03, 0x70, 0x01, 0xea, 0xaa, 0x19, 0x82, 0x01,
	0x0a, 0x3f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x48, 0x45, 0x58, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x3f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x33, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x92, 0x41, 0x17, 0x4a, 0x08, 0x22, 0x30, 0x30, 0x30, 0x30, 0x31, 0x33, 0x22, 0x9a,
	0x02, 0x01, 0x07, 0xa2, 0x02, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x6e, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xea, 0x01, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x42, 0xbd, 0x01, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x7a, 0x02,
	0x68, 0x05, 0xea, 0xaa, 0x19, 0x98, 0x01, 0x0a, 0x49, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c,
	0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x44, 0x65, 0x76,
	0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x92,
	0x41, 0x11, 0x4a, 0x0f, 0x5b, 0x22, 0x32, 0x36, 0x30, 0x30, 0x41, 0x42, 0x30, 0x30, 0x2f, 0x32,
	0x34, 0x22, 0x5d, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x32, 0xfe, 0x03, 0x0a, 0x02, 0x4e, 0x73, 0x12, 0x68, 0x0a, 0x0f, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x76,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x12, 0xae, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4d, 0x41, 0x43, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x41, 0x43, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x41, 0x43, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x44, 0x12, 0x42, 0x2f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x6d, 0x61, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x7b, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x5f, 0x70, 0x68, 0x79, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x58, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74,
	0x49, 0x44, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x6e, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x12, 0x82, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15,
	0x2f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x65, 0x73, 0x32, 0x90, 0x02, 0x0a, 0x04, 0x41, 0x73, 0x4e, 0x73, 0x12, 0x54,
	0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x75, 0x73, 0x68, 0x12, 0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5f, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x74,
	0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e,
	0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x1a, 0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x32, 0xa8, 0x01, 0x0a, 0x04, 0x47, 0x73, 0x4e,
	0x73, 0x12, 0x45, 0x0a, 0x0c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x1d, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x59, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x78, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x54, 0x78, 0x41, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0xbc, 0x06, 0x0a, 0x13, 0x4e, 0x73, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0xb2, 0x01, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x65, 0x12, 0x63, 0x2f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x86, 0x02, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x64,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45,
	0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xbe, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0xb7, 0x01, 0x1a, 0x63, 0x2f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x69, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x4d, 0x22, 0x48, 0x2f, 0x6e,
	0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xce, 0x01, 0x0a, 0x14, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x2b, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x6e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x68, 0x32, 0x63, 0x2f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x95, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x2a, 0x45, 0x2f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lorawan_stack_api_networkserver_proto_rawDescOnce sync.Once
	file_lorawan_stack_api_networkserver_proto_rawDescData = file_lorawan_stack_api_networkserver_proto_rawDesc
)

func file_lorawan_stack_api_networkserver_proto_rawDescGZIP() []byte {
	file_lorawan_stack_api_networkserver_proto_rawDescOnce.Do(func() {
		file_lorawan_stack_api_networkserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_lorawan_stack_api_networkserver_proto_rawDescData)
	})
	return file_lorawan_stack_api_networkserver_proto_rawDescData
}

var file_lorawan_stack_api_networkserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lorawan_stack_api_networkserver_proto_goTypes = []interface{}{
	(*GenerateDevAddrResponse)(nil),         // 0: ttn.lorawan.v3.GenerateDevAddrResponse
	(*GetDefaultMACSettingsRequest)(nil),    // 1: ttn.lorawan.v3.GetDefaultMACSettingsRequest
	(*GetNetIDResponse)(nil),                // 2: ttn.lorawan.v3.GetNetIDResponse
	(*GetDeviceAdressPrefixesResponse)(nil), // 3: ttn.lorawan.v3.GetDeviceAdressPrefixesResponse
	(PHYVersion)(0),                         // 4: ttn.lorawan.v3.PHYVersion
	(*emptypb.Empty)(nil),                   // 5: google.protobuf.Empty
	(*DownlinkQueueRequest)(nil),            // 6: ttn.lorawan.v3.DownlinkQueueRequest
	(*EndDeviceIdentifiers)(nil),            // 7: ttn.lorawan.v3.EndDeviceIdentifiers
	(*UplinkMessage)(nil),                   // 8: ttn.lorawan.v3.UplinkMessage
	(*GatewayTxAcknowledgment)(nil),         // 9: ttn.lorawan.v3.GatewayTxAcknowledgment
	(*GetEndDeviceRequest)(nil),             // 10: ttn.lorawan.v3.GetEndDeviceRequest
	(*SetEndDeviceRequest)(nil),             // 11: ttn.lorawan.v3.SetEndDeviceRequest
	(*ResetAndGetEndDeviceRequest)(nil),     // 12: ttn.lorawan.v3.ResetAndGetEndDeviceRequest
	(*MACSettings)(nil),                     // 13: ttn.lorawan.v3.MACSettings
	(*ApplicationDownlinks)(nil),            // 14: ttn.lorawan.v3.ApplicationDownlinks
	(*EndDevice)(nil),                       // 15: ttn.lorawan.v3.EndDevice
}
var file_lorawan_stack_api_networkserver_proto_depIdxs = []int32{
	4,  // 0: ttn.lorawan.v3.GetDefaultMACSettingsRequest.lorawan_phy_version:type_name -> ttn.lorawan.v3.PHYVersion
	5,  // 1: ttn.lorawan.v3.Ns.GenerateDevAddr:input_type -> google.protobuf.Empty
	1,  // 2: ttn.lorawan.v3.Ns.GetDefaultMACSettings:input_type -> ttn.lorawan.v3.GetDefaultMACSettingsRequest
	5,  // 3: ttn.lorawan.v3.Ns.GetNetID:input_type -> google.protobuf.Empty
	5,  // 4: ttn.lorawan.v3.Ns.GetDeviceAddressPrefixes:input_type -> google.protobuf.Empty
	6,  // 5: ttn.lorawan.v3.AsNs.DownlinkQueueReplace:input_type -> ttn.lorawan.v3.DownlinkQueueRequest
	6,  // 6: ttn.lorawan.v3.AsNs.DownlinkQueuePush:input_type -> ttn.lorawan.v3.DownlinkQueueRequest
	7,  // 7: ttn.lorawan.v3.AsNs.DownlinkQueueList:input_type -> ttn.lorawan.v3.EndDeviceIdentifiers
	8,  // 8: ttn.lorawan.v3.GsNs.HandleUplink:input_type -> ttn.lorawan.v3.UplinkMessage
	9,  // 9: ttn.lorawan.v3.GsNs.ReportTxAcknowledgment:input_type -> ttn.lorawan.v3.GatewayTxAcknowledgment
	10, // 10: ttn.lorawan.v3.NsEndDeviceRegistry.Get:input_type -> ttn.lorawan.v3.GetEndDeviceRequest
	11, // 11: ttn.lorawan.v3.NsEndDeviceRegistry.Set:input_type -> ttn.lorawan.v3.SetEndDeviceRequest
	12, // 12: ttn.lorawan.v3.NsEndDeviceRegistry.ResetFactoryDefaults:input_type -> ttn.lorawan.v3.ResetAndGetEndDeviceRequest
	7,  // 13: ttn.lorawan.v3.NsEndDeviceRegistry.Delete:input_type -> ttn.lorawan.v3.EndDeviceIdentifiers
	0,  // 14: ttn.lorawan.v3.Ns.GenerateDevAddr:output_type -> ttn.lorawan.v3.GenerateDevAddrResponse
	13, // 15: ttn.lorawan.v3.Ns.GetDefaultMACSettings:output_type -> ttn.lorawan.v3.MACSettings
	2,  // 16: ttn.lorawan.v3.Ns.GetNetID:output_type -> ttn.lorawan.v3.GetNetIDResponse
	3,  // 17: ttn.lorawan.v3.Ns.GetDeviceAddressPrefixes:output_type -> ttn.lorawan.v3.GetDeviceAdressPrefixesResponse
	5,  // 18: ttn.lorawan.v3.AsNs.DownlinkQueueReplace:output_type -> google.protobuf.Empty
	5,  // 19: ttn.lorawan.v3.AsNs.DownlinkQueuePush:output_type -> google.protobuf.Empty
	14, // 20: ttn.lorawan.v3.AsNs.DownlinkQueueList:output_type -> ttn.lorawan.v3.ApplicationDownlinks
	5,  // 21: ttn.lorawan.v3.GsNs.HandleUplink:output_type -> google.protobuf.Empty
	5,  // 22: ttn.lorawan.v3.GsNs.ReportTxAcknowledgment:output_type -> google.protobuf.Empty
	15, // 23: ttn.lorawan.v3.NsEndDeviceRegistry.Get:output_type -> ttn.lorawan.v3.EndDevice
	15, // 24: ttn.lorawan.v3.NsEndDeviceRegistry.Set:output_type -> ttn.lorawan.v3.EndDevice
	15, // 25: ttn.lorawan.v3.NsEndDeviceRegistry.ResetFactoryDefaults:output_type -> ttn.lorawan.v3.EndDevice
	5,  // 26: ttn.lorawan.v3.NsEndDeviceRegistry.Delete:output_type -> google.protobuf.Empty
	14, // [14:27] is the sub-list for method output_type
	1,  // [1:14] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_lorawan_stack_api_networkserver_proto_init() }
func file_lorawan_stack_api_networkserver_proto_init() {
	if File_lorawan_stack_api_networkserver_proto != nil {
		return
	}
	file_lorawan_stack_api_end_device_proto_init()
	file_lorawan_stack_api_identifiers_proto_init()
	file_lorawan_stack_api_messages_proto_init()
	file_lorawan_stack_api_lorawan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lorawan_stack_api_networkserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDevAddrResponse); i {
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
		file_lorawan_stack_api_networkserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultMACSettingsRequest); i {
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
		file_lorawan_stack_api_networkserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetIDResponse); i {
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
		file_lorawan_stack_api_networkserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceAdressPrefixesResponse); i {
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
			RawDescriptor: file_lorawan_stack_api_networkserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_lorawan_stack_api_networkserver_proto_goTypes,
		DependencyIndexes: file_lorawan_stack_api_networkserver_proto_depIdxs,
		MessageInfos:      file_lorawan_stack_api_networkserver_proto_msgTypes,
	}.Build()
	File_lorawan_stack_api_networkserver_proto = out.File
	file_lorawan_stack_api_networkserver_proto_rawDesc = nil
	file_lorawan_stack_api_networkserver_proto_goTypes = nil
	file_lorawan_stack_api_networkserver_proto_depIdxs = nil
}