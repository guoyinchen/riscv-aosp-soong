// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontend.proto

package ninja_frontend

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status_Message_Level int32

const (
	Status_Message_INFO    Status_Message_Level = 0
	Status_Message_WARNING Status_Message_Level = 1
	Status_Message_ERROR   Status_Message_Level = 2
	Status_Message_DEBUG   Status_Message_Level = 3
)

var Status_Message_Level_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
	3: "DEBUG",
}

var Status_Message_Level_value = map[string]int32{
	"INFO":    0,
	"WARNING": 1,
	"ERROR":   2,
	"DEBUG":   3,
}

func (x Status_Message_Level) Enum() *Status_Message_Level {
	p := new(Status_Message_Level)
	*p = x
	return p
}

func (x Status_Message_Level) String() string {
	return proto.EnumName(Status_Message_Level_name, int32(x))
}

func (x *Status_Message_Level) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Status_Message_Level_value, data, "Status_Message_Level")
	if err != nil {
		return err
	}
	*x = Status_Message_Level(value)
	return nil
}

func (Status_Message_Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 5, 0}
}

type Status struct {
	TotalEdges           *Status_TotalEdges    `protobuf:"bytes,1,opt,name=total_edges,json=totalEdges" json:"total_edges,omitempty"`
	BuildStarted         *Status_BuildStarted  `protobuf:"bytes,2,opt,name=build_started,json=buildStarted" json:"build_started,omitempty"`
	BuildFinished        *Status_BuildFinished `protobuf:"bytes,3,opt,name=build_finished,json=buildFinished" json:"build_finished,omitempty"`
	EdgeStarted          *Status_EdgeStarted   `protobuf:"bytes,4,opt,name=edge_started,json=edgeStarted" json:"edge_started,omitempty"`
	EdgeFinished         *Status_EdgeFinished  `protobuf:"bytes,5,opt,name=edge_finished,json=edgeFinished" json:"edge_finished,omitempty"`
	Message              *Status_Message       `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetTotalEdges() *Status_TotalEdges {
	if m != nil {
		return m.TotalEdges
	}
	return nil
}

func (m *Status) GetBuildStarted() *Status_BuildStarted {
	if m != nil {
		return m.BuildStarted
	}
	return nil
}

func (m *Status) GetBuildFinished() *Status_BuildFinished {
	if m != nil {
		return m.BuildFinished
	}
	return nil
}

func (m *Status) GetEdgeStarted() *Status_EdgeStarted {
	if m != nil {
		return m.EdgeStarted
	}
	return nil
}

func (m *Status) GetEdgeFinished() *Status_EdgeFinished {
	if m != nil {
		return m.EdgeFinished
	}
	return nil
}

func (m *Status) GetMessage() *Status_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Status_TotalEdges struct {
	// New value for total edges in the build.
	TotalEdges           *uint32  `protobuf:"varint,1,opt,name=total_edges,json=totalEdges" json:"total_edges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_TotalEdges) Reset()         { *m = Status_TotalEdges{} }
func (m *Status_TotalEdges) String() string { return proto.CompactTextString(m) }
func (*Status_TotalEdges) ProtoMessage()    {}
func (*Status_TotalEdges) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 0}
}

func (m *Status_TotalEdges) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_TotalEdges.Unmarshal(m, b)
}
func (m *Status_TotalEdges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_TotalEdges.Marshal(b, m, deterministic)
}
func (m *Status_TotalEdges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_TotalEdges.Merge(m, src)
}
func (m *Status_TotalEdges) XXX_Size() int {
	return xxx_messageInfo_Status_TotalEdges.Size(m)
}
func (m *Status_TotalEdges) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_TotalEdges.DiscardUnknown(m)
}

var xxx_messageInfo_Status_TotalEdges proto.InternalMessageInfo

func (m *Status_TotalEdges) GetTotalEdges() uint32 {
	if m != nil && m.TotalEdges != nil {
		return *m.TotalEdges
	}
	return 0
}

type Status_BuildStarted struct {
	// Number of jobs Ninja will run in parallel.
	Parallelism *uint32 `protobuf:"varint,1,opt,name=parallelism" json:"parallelism,omitempty"`
	// Verbose value passed to ninja.
	Verbose              *bool    `protobuf:"varint,2,opt,name=verbose" json:"verbose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_BuildStarted) Reset()         { *m = Status_BuildStarted{} }
func (m *Status_BuildStarted) String() string { return proto.CompactTextString(m) }
func (*Status_BuildStarted) ProtoMessage()    {}
func (*Status_BuildStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 1}
}

func (m *Status_BuildStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_BuildStarted.Unmarshal(m, b)
}
func (m *Status_BuildStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_BuildStarted.Marshal(b, m, deterministic)
}
func (m *Status_BuildStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_BuildStarted.Merge(m, src)
}
func (m *Status_BuildStarted) XXX_Size() int {
	return xxx_messageInfo_Status_BuildStarted.Size(m)
}
func (m *Status_BuildStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_BuildStarted.DiscardUnknown(m)
}

var xxx_messageInfo_Status_BuildStarted proto.InternalMessageInfo

func (m *Status_BuildStarted) GetParallelism() uint32 {
	if m != nil && m.Parallelism != nil {
		return *m.Parallelism
	}
	return 0
}

func (m *Status_BuildStarted) GetVerbose() bool {
	if m != nil && m.Verbose != nil {
		return *m.Verbose
	}
	return false
}

type Status_BuildFinished struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_BuildFinished) Reset()         { *m = Status_BuildFinished{} }
func (m *Status_BuildFinished) String() string { return proto.CompactTextString(m) }
func (*Status_BuildFinished) ProtoMessage()    {}
func (*Status_BuildFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 2}
}

func (m *Status_BuildFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_BuildFinished.Unmarshal(m, b)
}
func (m *Status_BuildFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_BuildFinished.Marshal(b, m, deterministic)
}
func (m *Status_BuildFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_BuildFinished.Merge(m, src)
}
func (m *Status_BuildFinished) XXX_Size() int {
	return xxx_messageInfo_Status_BuildFinished.Size(m)
}
func (m *Status_BuildFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_BuildFinished.DiscardUnknown(m)
}

var xxx_messageInfo_Status_BuildFinished proto.InternalMessageInfo

type Status_EdgeStarted struct {
	// Edge identification number, unique to a Ninja run.
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Edge start time in milliseconds since Ninja started.
	StartTime *uint32 `protobuf:"varint,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// List of edge inputs.
	Inputs []string `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty"`
	// List of edge outputs.
	Outputs []string `protobuf:"bytes,4,rep,name=outputs" json:"outputs,omitempty"`
	// Description field from the edge.
	Desc *string `protobuf:"bytes,5,opt,name=desc" json:"desc,omitempty"`
	// Command field from the edge.
	Command *string `protobuf:"bytes,6,opt,name=command" json:"command,omitempty"`
	// Edge uses console.
	Console              *bool    `protobuf:"varint,7,opt,name=console" json:"console,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_EdgeStarted) Reset()         { *m = Status_EdgeStarted{} }
func (m *Status_EdgeStarted) String() string { return proto.CompactTextString(m) }
func (*Status_EdgeStarted) ProtoMessage()    {}
func (*Status_EdgeStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 3}
}

func (m *Status_EdgeStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_EdgeStarted.Unmarshal(m, b)
}
func (m *Status_EdgeStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_EdgeStarted.Marshal(b, m, deterministic)
}
func (m *Status_EdgeStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_EdgeStarted.Merge(m, src)
}
func (m *Status_EdgeStarted) XXX_Size() int {
	return xxx_messageInfo_Status_EdgeStarted.Size(m)
}
func (m *Status_EdgeStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_EdgeStarted.DiscardUnknown(m)
}

var xxx_messageInfo_Status_EdgeStarted proto.InternalMessageInfo

func (m *Status_EdgeStarted) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Status_EdgeStarted) GetStartTime() uint32 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *Status_EdgeStarted) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Status_EdgeStarted) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Status_EdgeStarted) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func (m *Status_EdgeStarted) GetCommand() string {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ""
}

func (m *Status_EdgeStarted) GetConsole() bool {
	if m != nil && m.Console != nil {
		return *m.Console
	}
	return false
}

type Status_EdgeFinished struct {
	// Edge identification number, unique to a Ninja run.
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Edge end time in milliseconds since Ninja started.
	EndTime *uint32 `protobuf:"varint,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Exit status (0 for success).
	Status *int32 `protobuf:"zigzag32,3,opt,name=status" json:"status,omitempty"`
	// Edge output, may contain ANSI codes.
	Output *string `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
	// Number of milliseconds spent executing in user mode
	UserTime *uint32 `protobuf:"varint,5,opt,name=user_time,json=userTime" json:"user_time,omitempty"`
	// Number of milliseconds spent executing in kernel mode
	SystemTime           *uint32  `protobuf:"varint,6,opt,name=system_time,json=systemTime" json:"system_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_EdgeFinished) Reset()         { *m = Status_EdgeFinished{} }
func (m *Status_EdgeFinished) String() string { return proto.CompactTextString(m) }
func (*Status_EdgeFinished) ProtoMessage()    {}
func (*Status_EdgeFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 4}
}

func (m *Status_EdgeFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_EdgeFinished.Unmarshal(m, b)
}
func (m *Status_EdgeFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_EdgeFinished.Marshal(b, m, deterministic)
}
func (m *Status_EdgeFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_EdgeFinished.Merge(m, src)
}
func (m *Status_EdgeFinished) XXX_Size() int {
	return xxx_messageInfo_Status_EdgeFinished.Size(m)
}
func (m *Status_EdgeFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_EdgeFinished.DiscardUnknown(m)
}

var xxx_messageInfo_Status_EdgeFinished proto.InternalMessageInfo

func (m *Status_EdgeFinished) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Status_EdgeFinished) GetEndTime() uint32 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *Status_EdgeFinished) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *Status_EdgeFinished) GetOutput() string {
	if m != nil && m.Output != nil {
		return *m.Output
	}
	return ""
}

func (m *Status_EdgeFinished) GetUserTime() uint32 {
	if m != nil && m.UserTime != nil {
		return *m.UserTime
	}
	return 0
}

func (m *Status_EdgeFinished) GetSystemTime() uint32 {
	if m != nil && m.SystemTime != nil {
		return *m.SystemTime
	}
	return 0
}

type Status_Message struct {
	// Message priority level (DEBUG, INFO, WARNING, ERROR).
	Level *Status_Message_Level `protobuf:"varint,1,opt,name=level,enum=ninja.Status_Message_Level,def=0" json:"level,omitempty"`
	// Info/warning/error message from Ninja.
	Message              *string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_Message) Reset()         { *m = Status_Message{} }
func (m *Status_Message) String() string { return proto.CompactTextString(m) }
func (*Status_Message) ProtoMessage()    {}
func (*Status_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0, 5}
}

func (m *Status_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_Message.Unmarshal(m, b)
}
func (m *Status_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_Message.Marshal(b, m, deterministic)
}
func (m *Status_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_Message.Merge(m, src)
}
func (m *Status_Message) XXX_Size() int {
	return xxx_messageInfo_Status_Message.Size(m)
}
func (m *Status_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Status_Message proto.InternalMessageInfo

const Default_Status_Message_Level Status_Message_Level = Status_Message_INFO

func (m *Status_Message) GetLevel() Status_Message_Level {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return Default_Status_Message_Level
}

func (m *Status_Message) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("ninja.Status_Message_Level", Status_Message_Level_name, Status_Message_Level_value)
	proto.RegisterType((*Status)(nil), "ninja.Status")
	proto.RegisterType((*Status_TotalEdges)(nil), "ninja.Status.TotalEdges")
	proto.RegisterType((*Status_BuildStarted)(nil), "ninja.Status.BuildStarted")
	proto.RegisterType((*Status_BuildFinished)(nil), "ninja.Status.BuildFinished")
	proto.RegisterType((*Status_EdgeStarted)(nil), "ninja.Status.EdgeStarted")
	proto.RegisterType((*Status_EdgeFinished)(nil), "ninja.Status.EdgeFinished")
	proto.RegisterType((*Status_Message)(nil), "ninja.Status.Message")
}

func init() {
	proto.RegisterFile("frontend.proto", fileDescriptor_eca3873955a29cfe)
}

var fileDescriptor_eca3873955a29cfe = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xbd, 0x4e, 0xe2, 0x38, 0x1e, 0x27, 0xb9, 0x61, 0x25, 0x90, 0xeb, 0x0a, 0x35, 0xea, 0x53,
	0x5f, 0x08, 0x12, 0x42, 0x42, 0x20, 0x24, 0x44, 0x44, 0x5a, 0x8a, 0x20, 0x95, 0xb6, 0x45, 0x48,
	0xbc, 0x44, 0x4e, 0x77, 0x5a, 0x8c, 0xec, 0x75, 0xe4, 0xdd, 0x54, 0xe2, 0x37, 0xf8, 0x09, 0xfe,
	0x80, 0xaf, 0xe3, 0x01, 0xed, 0xec, 0xda, 0x75, 0x68, 0xdf, 0x7c, 0x76, 0xce, 0x9c, 0x39, 0x7b,
	0x76, 0x64, 0x18, 0x5f, 0x55, 0xa5, 0xd4, 0x28, 0xc5, 0x6c, 0x53, 0x95, 0xba, 0x64, 0xbe, 0xcc,
	0xe4, 0xf7, 0xf4, 0xf0, 0x4f, 0x00, 0xfd, 0x73, 0x9d, 0xea, 0xad, 0x62, 0x2f, 0x21, 0xd2, 0xa5,
	0x4e, 0xf3, 0x15, 0x8a, 0x6b, 0x54, 0xb1, 0x37, 0xf5, 0x8e, 0xa2, 0x67, 0xf1, 0x8c, 0x78, 0x33,
	0xcb, 0x99, 0x5d, 0x18, 0xc2, 0xc2, 0xd4, 0x39, 0xe8, 0xe6, 0x9b, 0xbd, 0x81, 0xd1, 0x7a, 0x9b,
	0xe5, 0x62, 0xa5, 0x74, 0x5a, 0x69, 0x14, 0x71, 0x87, 0x9a, 0x93, 0xdd, 0xe6, 0xb9, 0xa1, 0x9c,
	0x5b, 0x06, 0x1f, 0xae, 0x5b, 0x88, 0xcd, 0x61, 0x6c, 0x05, 0xae, 0x32, 0x99, 0xa9, 0x6f, 0x28,
	0xe2, 0x2e, 0x29, 0xec, 0xdf, 0xa3, 0x70, 0xec, 0x28, 0xdc, 0xce, 0xac, 0x21, 0x7b, 0x0d, 0x43,
	0xe3, 0xbc, 0xf1, 0xd0, 0x23, 0x85, 0xbd, 0x5d, 0x05, 0xe3, 0xb7, 0xb6, 0x10, 0xe1, 0x2d, 0x30,
	0x57, 0xa0, 0xee, 0xc6, 0x80, 0x7f, 0xdf, 0x15, 0x4c, 0x7b, 0x33, 0x9f, 0xc6, 0x35, 0xe3, 0x9f,
	0x42, 0x50, 0xa0, 0x52, 0xe9, 0x35, 0xc6, 0x7d, 0x6a, 0x7d, 0xb8, 0xdb, 0xfa, 0xc9, 0x16, 0x79,
	0xcd, 0x4a, 0x9e, 0x00, 0xdc, 0xc6, 0xc9, 0x0e, 0xee, 0xa6, 0x3f, 0x6a, 0x67, 0x9c, 0x7c, 0x80,
	0x61, 0x3b, 0x40, 0x36, 0x85, 0x68, 0x93, 0x56, 0x69, 0x9e, 0x63, 0x9e, 0xa9, 0xc2, 0x35, 0xb4,
	0x8f, 0x58, 0x0c, 0xc1, 0x0d, 0x56, 0xeb, 0x52, 0x21, 0xbd, 0xc7, 0x80, 0xd7, 0x30, 0xf9, 0x1f,
	0x46, 0x3b, 0x51, 0x26, 0xbf, 0x3d, 0x88, 0x5a, 0xd1, 0xb0, 0x31, 0x74, 0x32, 0xe1, 0x34, 0x3b,
	0x99, 0x60, 0x8f, 0x01, 0x28, 0xd6, 0x95, 0xce, 0x0a, 0xab, 0x36, 0xe2, 0x21, 0x9d, 0x5c, 0x64,
	0x05, 0xb2, 0x47, 0xd0, 0xcf, 0xe4, 0x66, 0xab, 0x55, 0xdc, 0x9d, 0x76, 0x8f, 0x42, 0xee, 0x90,
	0x71, 0x50, 0x6e, 0x35, 0x15, 0x7a, 0x54, 0xa8, 0x21, 0x63, 0xd0, 0x13, 0xa8, 0x2e, 0x29, 0xe5,
	0x90, 0xd3, 0xb7, 0x61, 0x5f, 0x96, 0x45, 0x91, 0x4a, 0x41, 0x09, 0x86, 0xbc, 0x86, 0xb6, 0x22,
	0x55, 0x99, 0x63, 0x1c, 0xd8, 0x9b, 0x38, 0x98, 0xfc, 0xf2, 0x60, 0xd8, 0x7e, 0x94, 0x3b, 0xce,
	0xf7, 0x60, 0x80, 0x52, 0xb4, 0x7d, 0x07, 0x28, 0x45, 0xed, 0x5a, 0xd1, 0xdb, 0xd0, 0xb2, 0x3d,
	0xe0, 0x0e, 0x99, 0x73, 0x6b, 0x93, 0x56, 0x28, 0xe4, 0x0e, 0xb1, 0x7d, 0x08, 0xb7, 0x0a, 0x2b,
	0xab, 0xe5, 0x93, 0xd6, 0xc0, 0x1c, 0x90, 0xd8, 0x01, 0x44, 0xea, 0x87, 0xd2, 0x58, 0xd8, 0x72,
	0xdf, 0xbe, 0x9f, 0x3d, 0x32, 0x84, 0xe4, 0xa7, 0x07, 0x81, 0xdb, 0x01, 0xf6, 0x02, 0xfc, 0x1c,
	0x6f, 0x30, 0x27, 0x9f, 0xe3, 0x7f, 0xb7, 0xdc, 0xb1, 0x66, 0x1f, 0x0d, 0xe5, 0x55, 0xef, 0x74,
	0x79, 0x7c, 0xc6, 0x2d, 0xdf, 0x04, 0x51, 0x2f, 0x59, 0xc7, 0x46, 0xe4, 0xe0, 0xe1, 0x73, 0xf0,
	0x89, 0xcf, 0x06, 0x40, 0x1d, 0x93, 0xff, 0x58, 0x04, 0xc1, 0x97, 0xb7, 0x7c, 0x79, 0xba, 0x3c,
	0x99, 0x78, 0x2c, 0x04, 0x7f, 0xc1, 0xf9, 0x19, 0x9f, 0x74, 0xcc, 0xe7, 0xbb, 0xc5, 0xfc, 0xf3,
	0xc9, 0xa4, 0x3b, 0x67, 0xef, 0xbb, 0x5f, 0xc7, 0x34, 0x7c, 0x55, 0xff, 0x1f, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0x93, 0x48, 0xcf, 0x2a, 0x04, 0x00, 0x00,
}
