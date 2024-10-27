// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.0
// source: logservice/logservice.proto

package logservice

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

type CreateLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App       string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	ProcessId string `protobuf:"bytes,4,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	Type      string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Status    string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	User      string `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateLogRequest) Reset() {
	*x = CreateLogRequest{}
	mi := &file_logservice_logservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogRequest) ProtoMessage() {}

func (x *CreateLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_logservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogRequest.ProtoReflect.Descriptor instead.
func (*CreateLogRequest) Descriptor() ([]byte, []int) {
	return file_logservice_logservice_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLogRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *CreateLogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLogRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CreateLogRequest) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *CreateLogRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateLogRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateLogRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type CreateLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLogResponse) Reset() {
	*x = CreateLogResponse{}
	mi := &file_logservice_logservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogResponse) ProtoMessage() {}

func (x *CreateLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_logservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogResponse.ProtoReflect.Descriptor instead.
func (*CreateLogResponse) Descriptor() ([]byte, []int) {
	return file_logservice_logservice_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLogResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLogRequest) Reset() {
	*x = GetLogRequest{}
	mi := &file_logservice_logservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogRequest) ProtoMessage() {}

func (x *GetLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_logservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogRequest.ProtoReflect.Descriptor instead.
func (*GetLogRequest) Descriptor() ([]byte, []int) {
	return file_logservice_logservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	App       string `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Data      string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ProcessId string `protobuf:"bytes,5,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	Type      string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Status    string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	User      string `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetLogResponse) Reset() {
	*x = GetLogResponse{}
	mi := &file_logservice_logservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogResponse) ProtoMessage() {}

func (x *GetLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logservice_logservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogResponse.ProtoReflect.Descriptor instead.
func (*GetLogResponse) Descriptor() ([]byte, []int) {
	return file_logservice_logservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLogResponse) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *GetLogResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLogResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GetLogResponse) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *GetLogResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetLogResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetLogResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_logservice_logservice_proto protoreflect.FileDescriptor

var file_logservice_logservice_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xb9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x5f, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x2e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x79, 0x75,
	0x76, 0x65, 0x64, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logservice_logservice_proto_rawDescOnce sync.Once
	file_logservice_logservice_proto_rawDescData = file_logservice_logservice_proto_rawDesc
)

func file_logservice_logservice_proto_rawDescGZIP() []byte {
	file_logservice_logservice_proto_rawDescOnce.Do(func() {
		file_logservice_logservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_logservice_logservice_proto_rawDescData)
	})
	return file_logservice_logservice_proto_rawDescData
}

var file_logservice_logservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logservice_logservice_proto_goTypes = []any{
	(*CreateLogRequest)(nil),  // 0: CreateLogRequest
	(*CreateLogResponse)(nil), // 1: CreateLogResponse
	(*GetLogRequest)(nil),     // 2: GetLogRequest
	(*GetLogResponse)(nil),    // 3: GetLogResponse
}
var file_logservice_logservice_proto_depIdxs = []int32{
	0, // 0: Log.Add:input_type -> CreateLogRequest
	2, // 1: Log.Get:input_type -> GetLogRequest
	1, // 2: Log.Add:output_type -> CreateLogResponse
	3, // 3: Log.Get:output_type -> GetLogResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logservice_logservice_proto_init() }
func file_logservice_logservice_proto_init() {
	if File_logservice_logservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logservice_logservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logservice_logservice_proto_goTypes,
		DependencyIndexes: file_logservice_logservice_proto_depIdxs,
		MessageInfos:      file_logservice_logservice_proto_msgTypes,
	}.Build()
	File_logservice_logservice_proto = out.File
	file_logservice_logservice_proto_rawDesc = nil
	file_logservice_logservice_proto_goTypes = nil
	file_logservice_logservice_proto_depIdxs = nil
}
