// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/proto/fleetspeak_server/services.proto

package fleetspeak_server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A ServiceConfig describes how the server should configure a 'service', which is
// a module that sends and processes messages.
type ServiceConfig struct {
	// The name that the service will be known as. Primary use is to address
	// messages to the service. The service names 'server' and 'client' are
	// reserved.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the factory which will be used to generate the service.
	Factory string `protobuf:"bytes,2,opt,name=factory,proto3" json:"factory,omitempty"`
	// The maximum number of simultaneous calls to the service's ProcessMessage
	// method. If unset, defaults to 100.
	MaxParallelism uint32 `protobuf:"varint,3,opt,name=max_parallelism,json=maxParallelism,proto3" json:"max_parallelism,omitempty"`
	// Additional configuration information for the factory to use when setting up
	// the service. The allowed type depends upon the factory.
	Config               *any.Any `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceConfig) Reset()         { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()    {}
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_9d67e677149c20ee, []int{0}
}
func (m *ServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceConfig.Unmarshal(m, b)
}
func (m *ServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceConfig.Marshal(b, m, deterministic)
}
func (dst *ServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceConfig.Merge(dst, src)
}
func (m *ServiceConfig) XXX_Size() int {
	return xxx_messageInfo_ServiceConfig.Size(m)
}
func (m *ServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceConfig proto.InternalMessageInfo

func (m *ServiceConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceConfig) GetFactory() string {
	if m != nil {
		return m.Factory
	}
	return ""
}

func (m *ServiceConfig) GetMaxParallelism() uint32 {
	if m != nil {
		return m.MaxParallelism
	}
	return 0
}

func (m *ServiceConfig) GetConfig() *any.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceConfig)(nil), "fleetspeak.server.ServiceConfig")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/proto/fleetspeak_server/services.proto", fileDescriptor_services_9d67e677149c20ee)
}

var fileDescriptor_services_9d67e677149c20ee = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x4f, 0x4a, 0xc7, 0x30,
	0x10, 0x85, 0x89, 0x96, 0x8a, 0x91, 0x2a, 0x06, 0x17, 0xd1, 0x55, 0x71, 0x63, 0x17, 0x92, 0x80,
	0x6e, 0xdd, 0x88, 0x17, 0x90, 0x7a, 0x80, 0x32, 0x0d, 0x93, 0x52, 0xcc, 0x9f, 0x92, 0x54, 0x69,
	0xef, 0xe1, 0x81, 0xe5, 0x97, 0xb4, 0x74, 0x95, 0xcc, 0x37, 0x1f, 0x6f, 0x1e, 0x7d, 0xd3, 0x06,
	0x71, 0x8e, 0x13, 0xc2, 0xb7, 0x8c, 0x41, 0xc9, 0x88, 0xe1, 0x17, 0x83, 0x9c, 0x82, 0x9f, 0xbd,
	0x3c, 0x76, 0xdd, 0xc6, 0x4f, 0xcf, 0xa8, 0x30, 0x8a, 0x24, 0xb0, 0xdb, 0xc3, 0x10, 0xd9, 0x78,
	0xb8, 0x1f, 0xbc, 0x1f, 0x0c, 0xe6, 0x84, 0xfe, 0x47, 0x4b, 0x70, 0x6b, 0xb6, 0x1f, 0xff, 0x08,
	0xad, 0xbe, 0x72, 0xc0, 0x87, 0x77, 0x7a, 0x1c, 0x18, 0xa3, 0x85, 0x03, 0x8b, 0x9c, 0xd4, 0xa4,
	0xb9, 0x6c, 0xd3, 0x9f, 0x71, 0x7a, 0xa1, 0x41, 0xcd, 0x3e, 0xac, 0xfc, 0x2c, 0xe1, 0x7d, 0x64,
	0x4f, 0xf4, 0xc6, 0xc2, 0xd2, 0x4d, 0x10, 0xc0, 0x18, 0x34, 0x63, 0xb4, 0xfc, 0xbc, 0x26, 0x4d,
	0xd5, 0x5e, 0x5b, 0x58, 0x3e, 0x0f, 0xca, 0x9e, 0x69, 0xa9, 0xd2, 0x01, 0x5e, 0xd4, 0xa4, 0xb9,
	0x7a, 0xb9, 0x13, 0xb9, 0x94, 0xd8, 0x4b, 0x89, 0x77, 0xb7, 0xb6, 0x9b, 0xd3, 0x97, 0x89, 0xbe,
	0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x40, 0x8c, 0x74, 0x3c, 0x0b, 0x01, 0x00, 0x00,
}
