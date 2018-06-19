// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/plugins/proto/plugins/config.proto

/*
Package fleetspeak_plugins is a generated protocol buffer package.

It is generated from these files:
	fleetspeak/src/server/plugins/proto/plugins/config.proto

It has these top-level messages:
	Plugin
	Config
	HttpsConfig
*/
package fleetspeak_plugins

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Plugin message describes how to instantiate a FS component from a go plugin
// file, as defined by go's plugin library: https://golang.org/pkg/plugin/
type Plugin struct {
	// The full path to the plugin file containing this component.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The name of a factory function exported by the file. It will be passed the
	// config string and must be one of the *Factory types defined in the
	// Fleetspeak server.plugins package.
	FactoryName string `protobuf:"bytes,2,opt,name=factory_name,json=factoryName" json:"factory_name,omitempty"`
	Config      string `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *Plugin) Reset()                    { *m = Plugin{} }
func (m *Plugin) String() string            { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()               {}
func (*Plugin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Plugin) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Plugin) GetFactoryName() string {
	if m != nil {
		return m.FactoryName
	}
	return ""
}

func (m *Plugin) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type Config struct {
	Datastore      *Plugin   `protobuf:"bytes,1,opt,name=datastore" json:"datastore,omitempty"`
	ServiceFactory []*Plugin `protobuf:"bytes,2,rep,name=service_factory,json=serviceFactory" json:"service_factory,omitempty"`
	Communicator   []*Plugin `protobuf:"bytes,3,rep,name=communicator" json:"communicator,omitempty"`
	StatsCollector *Plugin   `protobuf:"bytes,4,opt,name=stats_collector,json=statsCollector" json:"stats_collector,omitempty"`
	Authorizer     *Plugin   `protobuf:"bytes,5,opt,name=authorizer" json:"authorizer,omitempty"`
	Notifier       *Plugin   `protobuf:"bytes,6,opt,name=notifier" json:"notifier,omitempty"`
	Listener       *Plugin   `protobuf:"bytes,7,opt,name=listener" json:"listener,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetDatastore() *Plugin {
	if m != nil {
		return m.Datastore
	}
	return nil
}

func (m *Config) GetServiceFactory() []*Plugin {
	if m != nil {
		return m.ServiceFactory
	}
	return nil
}

func (m *Config) GetCommunicator() []*Plugin {
	if m != nil {
		return m.Communicator
	}
	return nil
}

func (m *Config) GetStatsCollector() *Plugin {
	if m != nil {
		return m.StatsCollector
	}
	return nil
}

func (m *Config) GetAuthorizer() *Plugin {
	if m != nil {
		return m.Authorizer
	}
	return nil
}

func (m *Config) GetNotifier() *Plugin {
	if m != nil {
		return m.Notifier
	}
	return nil
}

func (m *Config) GetListener() *Plugin {
	if m != nil {
		return m.Listener
	}
	return nil
}

type HttpsConfig struct {
	// The bind address to listen for connections on, e.g. ":443" or "localhost:1234".
	ListenAddress string `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress" json:"listen_address,omitempty"`
	// A certificate chain which identifies the server to clients. x509 format.
	Certificate string `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
	// The private key used to identify the server. Must match the first entry in
	// certificates. x509 format.
	Key string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	// If set includes support for streaming connections to clients.
	Streaming bool `protobuf:"varint,4,opt,name=streaming" json:"streaming,omitempty"`
}

func (m *HttpsConfig) Reset()                    { *m = HttpsConfig{} }
func (m *HttpsConfig) String() string            { return proto.CompactTextString(m) }
func (*HttpsConfig) ProtoMessage()               {}
func (*HttpsConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HttpsConfig) GetListenAddress() string {
	if m != nil {
		return m.ListenAddress
	}
	return ""
}

func (m *HttpsConfig) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *HttpsConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HttpsConfig) GetStreaming() bool {
	if m != nil {
		return m.Streaming
	}
	return false
}

func init() {
	proto.RegisterType((*Plugin)(nil), "fleetspeak.plugins.Plugin")
	proto.RegisterType((*Config)(nil), "fleetspeak.plugins.Config")
	proto.RegisterType((*HttpsConfig)(nil), "fleetspeak.plugins.HttpsConfig")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/plugins/proto/plugins/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x69, 0xd3, 0x27, 0x4f, 0x3b, 0xe9, 0xd3, 0x47, 0xf6, 0x20, 0x8b, 0x78, 0xa8, 0x05,
	0xc1, 0x53, 0x02, 0x0a, 0x52, 0x3c, 0x08, 0x52, 0x10, 0x4f, 0x22, 0xbd, 0x78, 0x0c, 0xeb, 0x76,
	0xd2, 0x2e, 0x4d, 0xb2, 0x61, 0x77, 0x2a, 0xd4, 0x0f, 0xe0, 0x67, 0xf2, 0xe3, 0x49, 0x36, 0xdb,
	0x17, 0xf1, 0x90, 0xdb, 0xec, 0x7f, 0xfe, 0xbf, 0x79, 0x49, 0x06, 0xa6, 0x59, 0x8e, 0x48, 0xb6,
	0x42, 0xb1, 0x4e, 0xac, 0x91, 0x89, 0x45, 0xf3, 0x8e, 0x26, 0xa9, 0xf2, 0xcd, 0x52, 0x95, 0x36,
	0xa9, 0x8c, 0x26, 0xbd, 0x7f, 0x49, 0x5d, 0x66, 0x6a, 0x19, 0x3b, 0x91, 0xb1, 0x03, 0x19, 0x7b,
	0xc3, 0xe4, 0x15, 0xc2, 0x17, 0x17, 0x32, 0x06, 0xbd, 0x4a, 0xd0, 0x8a, 0x77, 0xc6, 0x9d, 0xab,
	0xc1, 0xdc, 0xc5, 0xec, 0x02, 0x86, 0x99, 0x90, 0xa4, 0xcd, 0x36, 0x2d, 0x45, 0x81, 0xbc, 0xeb,
	0x72, 0x91, 0xd7, 0x9e, 0x45, 0x81, 0xec, 0x14, 0xc2, 0xa6, 0x09, 0x0f, 0x5c, 0xd2, 0xbf, 0x26,
	0x5f, 0x01, 0x84, 0x33, 0x17, 0xb2, 0x29, 0x0c, 0x16, 0x82, 0x84, 0x25, 0x6d, 0xd0, 0x95, 0x8f,
	0xae, 0xcf, 0xe2, 0xdf, 0xb3, 0xc4, 0xcd, 0x20, 0xf3, 0x83, 0x99, 0xcd, 0xe0, 0x7f, 0xbd, 0x9e,
	0x92, 0x98, 0xfa, 0x9e, 0xbc, 0x3b, 0x0e, 0x5a, 0xf8, 0x91, 0x47, 0x1e, 0x1b, 0x82, 0xdd, 0xc3,
	0x50, 0xea, 0xa2, 0xd8, 0x94, 0x4a, 0x0a, 0xd2, 0x86, 0x07, 0xad, 0x15, 0x7e, 0xf8, 0xdd, 0x10,
	0x24, 0xc8, 0xa6, 0x52, 0xe7, 0x39, 0xd6, 0x35, 0x79, 0xaf, 0x75, 0x89, 0x91, 0x43, 0x66, 0x3b,
	0x82, 0xdd, 0x01, 0x88, 0x0d, 0xad, 0xb4, 0x51, 0x1f, 0x68, 0xf8, 0x9f, 0x56, 0xfe, 0xc8, 0xcd,
	0x6e, 0xa1, 0x5f, 0x6a, 0x52, 0x99, 0x42, 0xc3, 0xc3, 0x56, 0x72, 0xef, 0xad, 0xb9, 0x5c, 0x59,
	0xc2, 0x12, 0x0d, 0xff, 0xdb, 0xce, 0xed, 0xbc, 0x93, 0xcf, 0x0e, 0x44, 0x4f, 0x44, 0x95, 0xf5,
	0xff, 0xef, 0x12, 0x46, 0x4d, 0x2e, 0x15, 0x8b, 0x85, 0x41, 0x6b, 0xfd, 0x8d, 0xfc, 0x6b, 0xd4,
	0x87, 0x46, 0x64, 0x63, 0x88, 0x24, 0x9a, 0xba, 0xb7, 0x14, 0xb4, 0xbf, 0x95, 0x23, 0x89, 0x9d,
	0x40, 0xb0, 0xc6, 0xad, 0x3f, 0x94, 0x3a, 0x64, 0xe7, 0x30, 0xb0, 0x64, 0x50, 0x14, 0xaa, 0x5c,
	0xba, 0xaf, 0xda, 0x9f, 0x1f, 0x84, 0xb7, 0xd0, 0xdd, 0xed, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xe4, 0x09, 0xa8, 0xf3, 0x02, 0x00, 0x00,
}