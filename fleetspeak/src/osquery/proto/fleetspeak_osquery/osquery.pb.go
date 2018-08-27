// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/osquery/proto/fleetspeak_osquery/osquery.proto

package fleetspeak_osquery

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

type CompressionType int32

const (
	CompressionType_UNCOMPRESSED CompressionType = 0
	CompressionType_ZCOMPRESSION CompressionType = 1
)

var CompressionType_name = map[int32]string{
	0: "UNCOMPRESSED",
	1: "ZCOMPRESSION",
}
var CompressionType_value = map[string]int32{
	"UNCOMPRESSED": 0,
	"ZCOMPRESSION": 1,
}

func (x CompressionType) String() string {
	return proto.EnumName(CompressionType_name, int32(x))
}
func (CompressionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{0}
}

type LoggedResult_Type int32

const (
	LoggedResult_UNKNOWN  LoggedResult_Type = 0
	LoggedResult_STRING   LoggedResult_Type = 1
	LoggedResult_SNAPSHOT LoggedResult_Type = 2
	LoggedResult_HEALTH   LoggedResult_Type = 3
	LoggedResult_INIT     LoggedResult_Type = 4
	LoggedResult_STATUS   LoggedResult_Type = 5
)

var LoggedResult_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "STRING",
	2: "SNAPSHOT",
	3: "HEALTH",
	4: "INIT",
	5: "STATUS",
}
var LoggedResult_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"STRING":   1,
	"SNAPSHOT": 2,
	"HEALTH":   3,
	"INIT":     4,
	"STATUS":   5,
}

func (x LoggedResult_Type) String() string {
	return proto.EnumName(LoggedResult_Type_name, int32(x))
}
func (LoggedResult_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{0, 0}
}

// A LoggedResult message contains data provided to an osquery Logger, containing
// e.g. query results from a snapshot query.
type LoggedResult struct {
	Type     LoggedResult_Type `protobuf:"varint,1,opt,name=type,proto3,enum=fleetspeak.osquery.LoggedResult_Type" json:"type,omitempty"`
	Compress CompressionType   `protobuf:"varint,2,opt,name=compress,proto3,enum=fleetspeak.osquery.CompressionType" json:"compress,omitempty"`
	// JSON payload, encoded according to the compress attribute.
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoggedResult) Reset()         { *m = LoggedResult{} }
func (m *LoggedResult) String() string { return proto.CompactTextString(m) }
func (*LoggedResult) ProtoMessage()    {}
func (*LoggedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{0}
}
func (m *LoggedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoggedResult.Unmarshal(m, b)
}
func (m *LoggedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoggedResult.Marshal(b, m, deterministic)
}
func (dst *LoggedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggedResult.Merge(dst, src)
}
func (m *LoggedResult) XXX_Size() int {
	return xxx_messageInfo_LoggedResult.Size(m)
}
func (m *LoggedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggedResult.DiscardUnknown(m)
}

var xxx_messageInfo_LoggedResult proto.InternalMessageInfo

func (m *LoggedResult) GetType() LoggedResult_Type {
	if m != nil {
		return m.Type
	}
	return LoggedResult_UNKNOWN
}

func (m *LoggedResult) GetCompress() CompressionType {
	if m != nil {
		return m.Compress
	}
	return CompressionType_UNCOMPRESSED
}

func (m *LoggedResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A Queries message contains queries to execute using osquery through the
// osquery distributed inteface.
type Queries struct {
	// map from query name to SQL text
	Queries map[string]string `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// map from query name to SQL text, if set the corresponding query will only
	// be executed if the discory query returns a result.
	Discovery            map[string]string `protobuf:"bytes,2,rep,name=discovery,proto3" json:"discovery,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Queries) Reset()         { *m = Queries{} }
func (m *Queries) String() string { return proto.CompactTextString(m) }
func (*Queries) ProtoMessage()    {}
func (*Queries) Descriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{1}
}
func (m *Queries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Queries.Unmarshal(m, b)
}
func (m *Queries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Queries.Marshal(b, m, deterministic)
}
func (dst *Queries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Queries.Merge(dst, src)
}
func (m *Queries) XXX_Size() int {
	return xxx_messageInfo_Queries.Size(m)
}
func (m *Queries) XXX_DiscardUnknown() {
	xxx_messageInfo_Queries.DiscardUnknown(m)
}

var xxx_messageInfo_Queries proto.InternalMessageInfo

func (m *Queries) GetQueries() map[string]string {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *Queries) GetDiscovery() map[string]string {
	if m != nil {
		return m.Discovery
	}
	return nil
}

type Row struct {
	Row                  map[string]string `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{2}
}
func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (dst *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(dst, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetRow() map[string]string {
	if m != nil {
		return m.Row
	}
	return nil
}

type Rows struct {
	Rows                 []*Row   `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rows) Reset()         { *m = Rows{} }
func (m *Rows) String() string { return proto.CompactTextString(m) }
func (*Rows) ProtoMessage()    {}
func (*Rows) Descriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{3}
}
func (m *Rows) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rows.Unmarshal(m, b)
}
func (m *Rows) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rows.Marshal(b, m, deterministic)
}
func (dst *Rows) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rows.Merge(dst, src)
}
func (m *Rows) XXX_Size() int {
	return xxx_messageInfo_Rows.Size(m)
}
func (m *Rows) XXX_DiscardUnknown() {
	xxx_messageInfo_Rows.DiscardUnknown(m)
}

var xxx_messageInfo_Rows proto.InternalMessageInfo

func (m *Rows) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

// A QueryResults message contains one or more results for a query.
type QueryResults struct {
	// The query that these results are for.
	QueryName string `protobuf:"bytes,1,opt,name=query_name,json=queryName,proto3" json:"query_name,omitempty"`
	// Status code for the execution of query_name, 0=OK.
	Status   int64           `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Compress CompressionType `protobuf:"varint,3,opt,name=compress,proto3,enum=fleetspeak.osquery.CompressionType" json:"compress,omitempty"`
	// A serialized Rows message, compressed according to the compress attribued.
	Rows                 []byte   `protobuf:"bytes,4,opt,name=Rows,proto3" json:"Rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResults) Reset()         { *m = QueryResults{} }
func (m *QueryResults) String() string { return proto.CompactTextString(m) }
func (*QueryResults) ProtoMessage()    {}
func (*QueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_osquery_55742b8371616e08, []int{4}
}
func (m *QueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResults.Unmarshal(m, b)
}
func (m *QueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResults.Marshal(b, m, deterministic)
}
func (dst *QueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResults.Merge(dst, src)
}
func (m *QueryResults) XXX_Size() int {
	return xxx_messageInfo_QueryResults.Size(m)
}
func (m *QueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResults proto.InternalMessageInfo

func (m *QueryResults) GetQueryName() string {
	if m != nil {
		return m.QueryName
	}
	return ""
}

func (m *QueryResults) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *QueryResults) GetCompress() CompressionType {
	if m != nil {
		return m.Compress
	}
	return CompressionType_UNCOMPRESSED
}

func (m *QueryResults) GetRows() []byte {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*LoggedResult)(nil), "fleetspeak.osquery.LoggedResult")
	proto.RegisterType((*Queries)(nil), "fleetspeak.osquery.Queries")
	proto.RegisterMapType((map[string]string)(nil), "fleetspeak.osquery.Queries.DiscoveryEntry")
	proto.RegisterMapType((map[string]string)(nil), "fleetspeak.osquery.Queries.QueriesEntry")
	proto.RegisterType((*Row)(nil), "fleetspeak.osquery.Row")
	proto.RegisterMapType((map[string]string)(nil), "fleetspeak.osquery.Row.RowEntry")
	proto.RegisterType((*Rows)(nil), "fleetspeak.osquery.Rows")
	proto.RegisterType((*QueryResults)(nil), "fleetspeak.osquery.QueryResults")
	proto.RegisterEnum("fleetspeak.osquery.CompressionType", CompressionType_name, CompressionType_value)
	proto.RegisterEnum("fleetspeak.osquery.LoggedResult_Type", LoggedResult_Type_name, LoggedResult_Type_value)
}

func init() {
	proto.RegisterFile("fleetspeak/src/osquery/proto/fleetspeak_osquery/osquery.proto", fileDescriptor_osquery_55742b8371616e08)
}

var fileDescriptor_osquery_55742b8371616e08 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x6d, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0x97, 0x87, 0xad, 0xed, 0x59, 0x98, 0x97, 0x83, 0x68, 0x28, 0x08, 0x25, 0x22, 0x94,
	0x09, 0x29, 0x74, 0x28, 0x3a, 0x14, 0xa9, 0x5b, 0xb1, 0xc5, 0x99, 0xcc, 0x9b, 0x14, 0xc1, 0x37,
	0x23, 0xb6, 0xd7, 0x51, 0xd6, 0xf6, 0xa6, 0xb9, 0xc9, 0x42, 0x3e, 0x86, 0x6f, 0xfd, 0x92, 0x7e,
	0x05, 0xc9, 0xbd, 0x89, 0xdd, 0xb4, 0x0a, 0xf3, 0x45, 0xb8, 0xe7, 0x9e, 0xf3, 0xfb, 0x9f, 0x73,
	0x72, 0x72, 0x02, 0xaf, 0xbf, 0x2e, 0x18, 0x4b, 0x45, 0xcc, 0xa2, 0xab, 0x9e, 0x48, 0xa6, 0x3d,
	0x2e, 0xd6, 0x19, 0x4b, 0x8a, 0x5e, 0x9c, 0xf0, 0x94, 0xf7, 0x36, 0xc1, 0x8b, 0x3a, 0x50, 0x9d,
	0xae, 0x04, 0x10, 0x37, 0x84, 0x5b, 0x45, 0x9c, 0x1f, 0x1a, 0x58, 0x67, 0xfc, 0xf2, 0x92, 0xcd,
	0x28, 0x13, 0xd9, 0x22, 0xc5, 0x97, 0x60, 0xa6, 0x45, 0xcc, 0x6c, 0xad, 0xa3, 0x75, 0x0f, 0xfa,
	0x4f, 0xdc, 0x3f, 0x35, 0xee, 0x4d, 0xde, 0x0d, 0x8b, 0x98, 0x51, 0x29, 0xc1, 0x37, 0xd0, 0x9c,
	0xf2, 0x65, 0x9c, 0x30, 0x21, 0x6c, 0x5d, 0xca, 0x1f, 0x6f, 0x93, 0x9f, 0x54, 0xcc, 0x9c, 0xaf,
	0xa4, 0xf8, 0x97, 0x08, 0x11, 0xcc, 0x59, 0x94, 0x46, 0xb6, 0xd1, 0xd1, 0xba, 0x16, 0x95, 0xb6,
	0xe3, 0x83, 0x59, 0x52, 0xb8, 0x0f, 0x8d, 0x89, 0xf7, 0xde, 0xf3, 0x3f, 0x79, 0x64, 0x07, 0x01,
	0xf6, 0x82, 0x90, 0x8e, 0xbd, 0x77, 0x44, 0x43, 0x0b, 0x9a, 0x81, 0x37, 0x38, 0x0f, 0x46, 0x7e,
	0x48, 0xf4, 0x32, 0x32, 0x1a, 0x0e, 0xce, 0xc2, 0x11, 0x31, 0xb0, 0x09, 0xe6, 0xd8, 0x1b, 0x87,
	0xc4, 0x54, 0xfc, 0x20, 0x9c, 0x04, 0x64, 0xd7, 0xf9, 0xa6, 0x43, 0xe3, 0x63, 0xc6, 0x92, 0x39,
	0x13, 0xf8, 0x16, 0x1a, 0x6b, 0x65, 0xda, 0x5a, 0xc7, 0xe8, 0xee, 0xf7, 0xbb, 0xdb, 0x1a, 0xae,
	0xe8, 0xfa, 0x1c, 0xae, 0xd2, 0xa4, 0xa0, 0xb5, 0x10, 0x47, 0xd0, 0x9a, 0xcd, 0xc5, 0x94, 0x5f,
	0xb3, 0xa4, 0xb0, 0x75, 0x99, 0xe5, 0xf0, 0x5f, 0x59, 0x4e, 0x6b, 0x58, 0xe5, 0xd9, 0x88, 0xdb,
	0xc7, 0x60, 0xdd, 0x2c, 0x81, 0x04, 0x8c, 0x2b, 0x56, 0xc8, 0x2f, 0xd1, 0xa2, 0xa5, 0x89, 0xf7,
	0x61, 0xf7, 0x3a, 0x5a, 0x64, 0x4c, 0x8e, 0xb7, 0x45, 0xd5, 0xe5, 0x58, 0x7f, 0xa1, 0xb5, 0x5f,
	0xc1, 0xc1, 0xed, 0xc4, 0x77, 0x51, 0x3b, 0x6b, 0x30, 0x28, 0xcf, 0xb1, 0x0f, 0x46, 0xc2, 0xf3,
	0x6a, 0x14, 0x9d, 0x6d, 0x2f, 0x41, 0x79, 0x5e, 0x3e, 0xaa, 0xf5, 0x12, 0x6e, 0x3f, 0x87, 0x66,
	0xed, 0xb8, 0x53, 0xc9, 0x23, 0x30, 0x29, 0xcf, 0x05, 0x3e, 0x05, 0x33, 0xe1, 0x79, 0x3d, 0xff,
	0x87, 0x7f, 0x29, 0x4a, 0x25, 0xe4, 0x7c, 0xd7, 0xd4, 0x88, 0x0a, 0xb5, 0x7c, 0x02, 0x1f, 0x01,
	0x48, 0xe6, 0x62, 0x15, 0x2d, 0x59, 0x55, 0xb8, 0x25, 0x3d, 0x5e, 0xb4, 0x64, 0xf8, 0x00, 0xf6,
	0x44, 0x1a, 0xa5, 0x99, 0xda, 0x47, 0x83, 0x56, 0xb7, 0x5b, 0x9b, 0x6a, 0xfc, 0xe7, 0xa6, 0x96,
	0xdd, 0xdb, 0xa6, 0xda, 0xd4, 0xd2, 0x3e, 0x7c, 0x06, 0xf7, 0x7e, 0x13, 0x20, 0x01, 0x6b, 0xe2,
	0x9d, 0xf8, 0x1f, 0xce, 0xe9, 0x30, 0x08, 0x86, 0xa7, 0x64, 0xa7, 0xf4, 0x7c, 0xae, 0x1d, 0x63,
	0xdf, 0x23, 0xda, 0x97, 0x3d, 0xf9, 0x73, 0x1e, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x92, 0x26,
	0x4e, 0x2a, 0xdd, 0x03, 0x00, 0x00,
}