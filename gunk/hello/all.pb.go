// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: github.com/ghazninattarshah/tracing-with-jaeger/gunk/hello/all.proto

package hello

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServiceID is identifier that the message shall be sent to
type ServiceID int32

const (
	ServiceID_UNKNOWN_Service ServiceID = 0
	ServiceID_ONE             ServiceID = 1
	ServiceID_TWO             ServiceID = 2
	ServiceID_THREE           ServiceID = 3
)

// Enum value maps for ServiceID.
var (
	ServiceID_name = map[int32]string{
		0: "UNKNOWN_Service",
		1: "ONE",
		2: "TWO",
		3: "THREE",
	}
	ServiceID_value = map[string]int32{
		"UNKNOWN_Service": 0,
		"ONE":             1,
		"TWO":             2,
		"THREE":           3,
	}
)

func (x ServiceID) Enum() *ServiceID {
	p := new(ServiceID)
	*p = x
	return p
}

func (x ServiceID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceID) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_enumTypes[0].Descriptor()
}

func (ServiceID) Type() protoreflect.EnumType {
	return &file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_enumTypes[0]
}

func (x ServiceID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceID.Descriptor instead.
func (ServiceID) EnumDescriptor() ([]byte, []int) {
	return file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescGZIP(), []int{0}
}

// HelloRequest contains an hello request message.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Msg is a message from a client.
	Msg string `protobuf:"bytes,1,opt,name=Msg,json=msg,proto3" json:"msg,omitempty"`
	// ServiceID is the service that the hello will be sent to
	ServiceID ServiceID `protobuf:"varint,2,opt,name=ServiceID,json=service_id,proto3,enum=hello.ServiceID" json:"service_id,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *HelloRequest) GetServiceID() ServiceID {
	if x != nil {
		return x.ServiceID
	}
	return ServiceID_UNKNOWN_Service
}

// HelloResponse contains an hello response from specific service.
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Msg is a message from a service.
	Msg string `protobuf:"bytes,1,opt,name=Msg,json=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto protoreflect.FileDescriptor

var file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x68, 0x61,
	0x7a, 0x6e, 0x69, 0x6e, 0x61, 0x74, 0x74, 0x61, 0x72, 0x73, 0x68, 0x61, 0x68, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x6a, 0x61, 0x65, 0x67, 0x65,
	0x72, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3b, 0x0a, 0x09, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x35,
	0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x2a, 0x51, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x12, 0x17, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x0b, 0x0a, 0x03, 0x4f,
	0x4e, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x0b, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10,
	0x02, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x0d, 0x0a, 0x05, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0x03,
	0x1a, 0x02, 0x08, 0x00, 0x1a, 0x02, 0x18, 0x00, 0x32, 0xd6, 0x02, 0x0a, 0x0c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc0, 0x02, 0x0a, 0x08, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x84, 0x02, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92, 0x41, 0xe6, 0x01, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x0a, 0x08, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x35, 0x53, 0x65, 0x6e, 0x64, 0x20, 0x61, 0x20, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x20, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x1d, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x61, 0x79, 0x20, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4a, 0x43, 0x0a, 0x03,
	0x32, 0x30, 0x30, 0x12, 0x3c, 0x0a, 0x16, 0x41, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x12, 0x22, 0x0a,
	0x20, 0x1a, 0x1e, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4a, 0x31, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x2a, 0x0a, 0x28, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f,
	0x73, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02,
	0x00, 0x42, 0x9a, 0x02, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x68, 0x61, 0x7a, 0x6e, 0x69, 0x6e, 0x61, 0x74, 0x74, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x68, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2d, 0x77, 0x69,
	0x74, 0x68, 0x2d, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00,
	0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x92,
	0x41, 0xbb, 0x01, 0x0a, 0x03, 0x32, 0x2e, 0x30, 0x12, 0xb0, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x41, 0x50, 0x49, 0x12, 0x93,
	0x01, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x6d, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x20, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x64, 0x20, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x73,
	0x68, 0x61, 0x6c, 0x6c, 0x20, 0x70, 0x61, 0x73, 0x73, 0x20, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x20, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescOnce sync.Once
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescData = file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDesc
)

func file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescGZIP() []byte {
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescOnce.Do(func() {
		file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescData)
	})
	return file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDescData
}

var (
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes  = make([]protoimpl.MessageInfo, 2)
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_goTypes   = []interface{}{
		(ServiceID)(0),        // 0: hello.ServiceID
		(*HelloRequest)(nil),  // 1: hello.HelloRequest
		(*HelloResponse)(nil), // 2: hello.HelloResponse
	}
)

var file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_depIdxs = []int32{
	0, // 0: hello.HelloRequest.ServiceID:type_name -> hello.ServiceID
	1, // 1: hello.HelloService.SayHello:input_type -> hello.HelloRequest
	2, // 2: hello.HelloService.SayHello:output_type -> hello.HelloResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_init() }
func file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_init() {
	if File_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_goTypes,
		DependencyIndexes: file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_depIdxs,
		EnumInfos:         file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_enumTypes,
		MessageInfos:      file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_msgTypes,
	}.Build()
	File_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto = out.File
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_rawDesc = nil
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_goTypes = nil
	file_github_com_ghazninattarshah_tracing_with_jaeger_gunk_hello_all_proto_depIdxs = nil
}
