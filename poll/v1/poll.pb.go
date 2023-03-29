// Copyright 2019, Intergral GmbH
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: deepproto/proto/poll/v1/poll.proto

package v1

import (
	v1 "github.com/intergral/go-deep-proto/resource/v1"
	v11 "github.com/intergral/go-deep-proto/tracepoint/v1"
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

type ResponseType int32

const (
	ResponseType_NO_CHANGE ResponseType = 0 // This is sent when the 'currentHash' from the request is the same as the response. So the client should do nothing.
	ResponseType_UPDATE    ResponseType = 1 // This is sent when the client should process the response to update the config.
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "NO_CHANGE",
		1: "UPDATE",
	}
	ResponseType_value = map[string]int32{
		"NO_CHANGE": 0,
		"UPDATE":    1,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_deepproto_proto_poll_v1_poll_proto_enumTypes[0].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_deepproto_proto_poll_v1_poll_proto_enumTypes[0]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_deepproto_proto_poll_v1_poll_proto_rawDescGZIP(), []int{0}
}

type PollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          int64        `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`                                     //time message was sent, acts as message id (useful for tracing)
	CurrentHash string       `protobuf:"bytes,2,opt,name=current_hash,json=currentHash,proto3" json:"current_hash,omitempty"` //some id that represents the clients current config, or null if no current config
	Resource    *v1.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`                          //this is the attributes that describe the resource requesting a config (e.g. service.name: shop-service)
}

func (x *PollRequest) Reset() {
	*x = PollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deepproto_proto_poll_v1_poll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollRequest) ProtoMessage() {}

func (x *PollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deepproto_proto_poll_v1_poll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollRequest.ProtoReflect.Descriptor instead.
func (*PollRequest) Descriptor() ([]byte, []int) {
	return file_deepproto_proto_poll_v1_poll_proto_rawDescGZIP(), []int{0}
}

func (x *PollRequest) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *PollRequest) GetCurrentHash() string {
	if x != nil {
		return x.CurrentHash
	}
	return ""
}

func (x *PollRequest) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type PollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           int64                   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`                                                                                   //time message was sent, acts as message id (useful for tracing)
	CurrentHash  string                  `protobuf:"bytes,2,opt,name=current_hash,json=currentHash,proto3" json:"current_hash,omitempty"`                                               //some id that represents the clients current config, or null if no current config
	Response     []*v11.TracePointConfig `protobuf:"bytes,3,rep,name=response,proto3" json:"response,omitempty"`                                                                        // this would be the list of dynamic configs that are to be installed. This should be the full set, not a partial or delta. Can be null if the response type is 'no-change'
	ResponseType ResponseType            `protobuf:"varint,4,opt,name=response_type,json=responseType,proto3,enum=deepproto.proto.poll.v1.ResponseType" json:"response_type,omitempty"` // This indicates if the config has changed or not. if 'no-change' then 'response' will be null/empty
}

func (x *PollResponse) Reset() {
	*x = PollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deepproto_proto_poll_v1_poll_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollResponse) ProtoMessage() {}

func (x *PollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deepproto_proto_poll_v1_poll_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollResponse.ProtoReflect.Descriptor instead.
func (*PollResponse) Descriptor() ([]byte, []int) {
	return file_deepproto_proto_poll_v1_poll_proto_rawDescGZIP(), []int{1}
}

func (x *PollResponse) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *PollResponse) GetCurrentHash() string {
	if x != nil {
		return x.CurrentHash
	}
	return ""
}

func (x *PollResponse) GetResponse() []*v11.TracePointConfig {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *PollResponse) GetResponseType() ResponseType {
	if x != nil {
		return x.ResponseType
	}
	return ResponseType_NO_CHANGE
}

var File_deepproto_proto_poll_v1_poll_proto protoreflect.FileDescriptor

var file_deepproto_proto_poll_v1_poll_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x64,
	0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x50, 0x6f,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x64, 0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0xda, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x4b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x29, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x32, 0x63, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x12, 0x24, 0x2e,
	0x64, 0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x50, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x67, 0x72, 0x61, 0x6c, 0x2e, 0x64, 0x65,
	0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x67, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x65, 0x70,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deepproto_proto_poll_v1_poll_proto_rawDescOnce sync.Once
	file_deepproto_proto_poll_v1_poll_proto_rawDescData = file_deepproto_proto_poll_v1_poll_proto_rawDesc
)

func file_deepproto_proto_poll_v1_poll_proto_rawDescGZIP() []byte {
	file_deepproto_proto_poll_v1_poll_proto_rawDescOnce.Do(func() {
		file_deepproto_proto_poll_v1_poll_proto_rawDescData = protoimpl.X.CompressGZIP(file_deepproto_proto_poll_v1_poll_proto_rawDescData)
	})
	return file_deepproto_proto_poll_v1_poll_proto_rawDescData
}

var file_deepproto_proto_poll_v1_poll_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deepproto_proto_poll_v1_poll_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_deepproto_proto_poll_v1_poll_proto_goTypes = []interface{}{
	(ResponseType)(0),            // 0: deepproto.proto.poll.v1.ResponseType
	(*PollRequest)(nil),          // 1: deepproto.proto.poll.v1.PollRequest
	(*PollResponse)(nil),         // 2: deepproto.proto.poll.v1.PollResponse
	(*v1.Resource)(nil),          // 3: deepproto.proto.resource.v1.Resource
	(*v11.TracePointConfig)(nil), // 4: deepproto.proto.tracepoint.v1.TracePointConfig
}
var file_deepproto_proto_poll_v1_poll_proto_depIdxs = []int32{
	3, // 0: deepproto.proto.poll.v1.PollRequest.resource:type_name -> deepproto.proto.resource.v1.Resource
	4, // 1: deepproto.proto.poll.v1.PollResponse.response:type_name -> deepproto.proto.tracepoint.v1.TracePointConfig
	0, // 2: deepproto.proto.poll.v1.PollResponse.response_type:type_name -> deepproto.proto.poll.v1.ResponseType
	1, // 3: deepproto.proto.poll.v1.PollConfig.poll:input_type -> deepproto.proto.poll.v1.PollRequest
	2, // 4: deepproto.proto.poll.v1.PollConfig.poll:output_type -> deepproto.proto.poll.v1.PollResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_deepproto_proto_poll_v1_poll_proto_init() }
func file_deepproto_proto_poll_v1_poll_proto_init() {
	if File_deepproto_proto_poll_v1_poll_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deepproto_proto_poll_v1_poll_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollRequest); i {
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
		file_deepproto_proto_poll_v1_poll_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollResponse); i {
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
			RawDescriptor: file_deepproto_proto_poll_v1_poll_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deepproto_proto_poll_v1_poll_proto_goTypes,
		DependencyIndexes: file_deepproto_proto_poll_v1_poll_proto_depIdxs,
		EnumInfos:         file_deepproto_proto_poll_v1_poll_proto_enumTypes,
		MessageInfos:      file_deepproto_proto_poll_v1_poll_proto_msgTypes,
	}.Build()
	File_deepproto_proto_poll_v1_poll_proto = out.File
	file_deepproto_proto_poll_v1_poll_proto_rawDesc = nil
	file_deepproto_proto_poll_v1_poll_proto_goTypes = nil
	file_deepproto_proto_poll_v1_poll_proto_depIdxs = nil
}
