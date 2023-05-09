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
// source: lorawan-stack/api/enums.proto

package ttnpb

import (
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
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

type DownlinkPathConstraint int32

const (
	// Indicates that the gateway can be selected for downlink without constraints by the Network Server.
	DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NONE DownlinkPathConstraint = 0
	// Indicates that the gateway can be selected for downlink only if no other or better gateway can be selected.
	DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER DownlinkPathConstraint = 1
	// Indicates that this gateway will never be selected for downlink, even if that results in no available downlink path.
	DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NEVER DownlinkPathConstraint = 2
)

// Enum value maps for DownlinkPathConstraint.
var (
	DownlinkPathConstraint_name = map[int32]string{
		0: "DOWNLINK_PATH_CONSTRAINT_NONE",
		1: "DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER",
		2: "DOWNLINK_PATH_CONSTRAINT_NEVER",
	}
	DownlinkPathConstraint_value = map[string]int32{
		"DOWNLINK_PATH_CONSTRAINT_NONE":         0,
		"DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER": 1,
		"DOWNLINK_PATH_CONSTRAINT_NEVER":        2,
	}
)

func (x DownlinkPathConstraint) Enum() *DownlinkPathConstraint {
	p := new(DownlinkPathConstraint)
	*p = x
	return p
}

func (x DownlinkPathConstraint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownlinkPathConstraint) Descriptor() protoreflect.EnumDescriptor {
	return file_lorawan_stack_api_enums_proto_enumTypes[0].Descriptor()
}

func (DownlinkPathConstraint) Type() protoreflect.EnumType {
	return &file_lorawan_stack_api_enums_proto_enumTypes[0]
}

func (x DownlinkPathConstraint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownlinkPathConstraint.Descriptor instead.
func (DownlinkPathConstraint) EnumDescriptor() ([]byte, []int) {
	return file_lorawan_stack_api_enums_proto_rawDescGZIP(), []int{0}
}

// State enum defines states that an entity can be in.
type State int32

const (
	// Denotes that the entity has been requested and is pending review by an admin.
	State_STATE_REQUESTED State = 0
	// Denotes that the entity has been reviewed and approved by an admin.
	State_STATE_APPROVED State = 1
	// Denotes that the entity has been reviewed and rejected by an admin.
	State_STATE_REJECTED State = 2
	// Denotes that the entity has been flagged and is pending review by an admin.
	State_STATE_FLAGGED State = 3
	// Denotes that the entity has been reviewed and suspended by an admin.
	State_STATE_SUSPENDED State = 4
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_REQUESTED",
		1: "STATE_APPROVED",
		2: "STATE_REJECTED",
		3: "STATE_FLAGGED",
		4: "STATE_SUSPENDED",
	}
	State_value = map[string]int32{
		"STATE_REQUESTED": 0,
		"STATE_APPROVED":  1,
		"STATE_REJECTED":  2,
		"STATE_FLAGGED":   3,
		"STATE_SUSPENDED": 4,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_lorawan_stack_api_enums_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_lorawan_stack_api_enums_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_lorawan_stack_api_enums_proto_rawDescGZIP(), []int{1}
}

type ClusterRole int32

const (
	ClusterRole_NONE                         ClusterRole = 0
	ClusterRole_ENTITY_REGISTRY              ClusterRole = 1
	ClusterRole_ACCESS                       ClusterRole = 2
	ClusterRole_GATEWAY_SERVER               ClusterRole = 3
	ClusterRole_NETWORK_SERVER               ClusterRole = 4
	ClusterRole_APPLICATION_SERVER           ClusterRole = 5
	ClusterRole_JOIN_SERVER                  ClusterRole = 6
	ClusterRole_CRYPTO_SERVER                ClusterRole = 7
	ClusterRole_DEVICE_TEMPLATE_CONVERTER    ClusterRole = 8
	ClusterRole_DEVICE_CLAIMING_SERVER       ClusterRole = 9
	ClusterRole_GATEWAY_CONFIGURATION_SERVER ClusterRole = 10
	ClusterRole_QR_CODE_GENERATOR            ClusterRole = 11
	ClusterRole_PACKET_BROKER_AGENT          ClusterRole = 12
	ClusterRole_DEVICE_REPOSITORY            ClusterRole = 13
)

// Enum value maps for ClusterRole.
var (
	ClusterRole_name = map[int32]string{
		0:  "NONE",
		1:  "ENTITY_REGISTRY",
		2:  "ACCESS",
		3:  "GATEWAY_SERVER",
		4:  "NETWORK_SERVER",
		5:  "APPLICATION_SERVER",
		6:  "JOIN_SERVER",
		7:  "CRYPTO_SERVER",
		8:  "DEVICE_TEMPLATE_CONVERTER",
		9:  "DEVICE_CLAIMING_SERVER",
		10: "GATEWAY_CONFIGURATION_SERVER",
		11: "QR_CODE_GENERATOR",
		12: "PACKET_BROKER_AGENT",
		13: "DEVICE_REPOSITORY",
	}
	ClusterRole_value = map[string]int32{
		"NONE":                         0,
		"ENTITY_REGISTRY":              1,
		"ACCESS":                       2,
		"GATEWAY_SERVER":               3,
		"NETWORK_SERVER":               4,
		"APPLICATION_SERVER":           5,
		"JOIN_SERVER":                  6,
		"CRYPTO_SERVER":                7,
		"DEVICE_TEMPLATE_CONVERTER":    8,
		"DEVICE_CLAIMING_SERVER":       9,
		"GATEWAY_CONFIGURATION_SERVER": 10,
		"QR_CODE_GENERATOR":            11,
		"PACKET_BROKER_AGENT":          12,
		"DEVICE_REPOSITORY":            13,
	}
)

func (x ClusterRole) Enum() *ClusterRole {
	p := new(ClusterRole)
	*p = x
	return p
}

func (x ClusterRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterRole) Descriptor() protoreflect.EnumDescriptor {
	return file_lorawan_stack_api_enums_proto_enumTypes[2].Descriptor()
}

func (ClusterRole) Type() protoreflect.EnumType {
	return &file_lorawan_stack_api_enums_proto_enumTypes[2]
}

func (x ClusterRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterRole.Descriptor instead.
func (ClusterRole) EnumDescriptor() ([]byte, []int) {
	return file_lorawan_stack_api_enums_proto_rawDescGZIP(), []int{2}
}

var File_lorawan_stack_api_enums_proto protoreflect.FileDescriptor

var file_lorawan_stack_api_enums_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6a, 0x73,
	0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xac, 0x01, 0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x50, 0x61, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x1d, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x41, 0x54, 0x48,
	0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50,
	0x41, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x54, 0x5f, 0x50,
	0x52, 0x45, 0x46, 0x45, 0x52, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x22, 0x0a,
	0x1e, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x5f, 0x43,
	0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10,
	0x02, 0x1a, 0x20, 0xea, 0xaa, 0x19, 0x1c, 0x18, 0x01, 0x2a, 0x18, 0x44, 0x4f, 0x57, 0x4e, 0x4c,
	0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41,
	0x49, 0x4e, 0x54, 0x2a, 0x7b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x46, 0x4c, 0x41, 0x47, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0x04, 0x1a, 0x0d, 0xea, 0xaa, 0x19, 0x09, 0x18, 0x01, 0x2a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x2a, 0xc0, 0x02, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x47,
	0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4a,
	0x4f, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x52, 0x59, 0x50, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x07, 0x12,
	0x1d, 0x0a, 0x19, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41,
	0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x45, 0x52, 0x10, 0x08, 0x12, 0x1a,
	0x0a, 0x16, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x09, 0x12, 0x20, 0x0a, 0x1c, 0x47, 0x41,
	0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11,
	0x51, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x42, 0x52,
	0x4f, 0x4b, 0x45, 0x52, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f, 0x52,
	0x59, 0x10, 0x0d, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lorawan_stack_api_enums_proto_rawDescOnce sync.Once
	file_lorawan_stack_api_enums_proto_rawDescData = file_lorawan_stack_api_enums_proto_rawDesc
)

func file_lorawan_stack_api_enums_proto_rawDescGZIP() []byte {
	file_lorawan_stack_api_enums_proto_rawDescOnce.Do(func() {
		file_lorawan_stack_api_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_lorawan_stack_api_enums_proto_rawDescData)
	})
	return file_lorawan_stack_api_enums_proto_rawDescData
}

var file_lorawan_stack_api_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_lorawan_stack_api_enums_proto_goTypes = []interface{}{
	(DownlinkPathConstraint)(0), // 0: ttn.lorawan.v3.DownlinkPathConstraint
	(State)(0),                  // 1: ttn.lorawan.v3.State
	(ClusterRole)(0),            // 2: ttn.lorawan.v3.ClusterRole
}
var file_lorawan_stack_api_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lorawan_stack_api_enums_proto_init() }
func file_lorawan_stack_api_enums_proto_init() {
	if File_lorawan_stack_api_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lorawan_stack_api_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lorawan_stack_api_enums_proto_goTypes,
		DependencyIndexes: file_lorawan_stack_api_enums_proto_depIdxs,
		EnumInfos:         file_lorawan_stack_api_enums_proto_enumTypes,
	}.Build()
	File_lorawan_stack_api_enums_proto = out.File
	file_lorawan_stack_api_enums_proto_rawDesc = nil
	file_lorawan_stack_api_enums_proto_goTypes = nil
	file_lorawan_stack_api_enums_proto_depIdxs = nil
}