// Code generated by protoc-gen-go.
// source: param.proto
// DO NOT EDIT!

package Report

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Param_ParaType int32

const (
	Param_Null   Param_ParaType = 0
	Param_UINT8  Param_ParaType = 1
	Param_UINT16 Param_ParaType = 2
	Param_UINT32 Param_ParaType = 3
	Param_UINT64 Param_ParaType = 4
	Param_FLOAT  Param_ParaType = 16
	Param_DOUBLE Param_ParaType = 17
	Param_STRING Param_ParaType = 32
	Param_BYTES  Param_ParaType = 33
)

var Param_ParaType_name = map[int32]string{
	0:  "Null",
	1:  "UINT8",
	2:  "UINT16",
	3:  "UINT32",
	4:  "UINT64",
	16: "FLOAT",
	17: "DOUBLE",
	32: "STRING",
	33: "BYTES",
}
var Param_ParaType_value = map[string]int32{
	"Null":   0,
	"UINT8":  1,
	"UINT16": 2,
	"UINT32": 3,
	"UINT64": 4,
	"FLOAT":  16,
	"DOUBLE": 17,
	"STRING": 32,
	"BYTES":  33,
}

func (x Param_ParaType) String() string {
	return proto.EnumName(Param_ParaType_name, int32(x))
}
func (Param_ParaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type Param struct {
	Type    Param_ParaType `protobuf:"varint,1,opt,name=type,enum=Report.Param_ParaType" json:"type,omitempty"`
	Npara   uint64         `protobuf:"varint,2,opt,name=npara" json:"npara,omitempty"`
	Dpara   float64        `protobuf:"fixed64,3,opt,name=dpara" json:"dpara,omitempty"`
	Strpara string         `protobuf:"bytes,4,opt,name=strpara" json:"strpara,omitempty"`
	Bpara   []byte         `protobuf:"bytes,5,opt,name=bpara,proto3" json:"bpara,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Param) GetType() Param_ParaType {
	if m != nil {
		return m.Type
	}
	return Param_Null
}

func (m *Param) GetNpara() uint64 {
	if m != nil {
		return m.Npara
	}
	return 0
}

func (m *Param) GetDpara() float64 {
	if m != nil {
		return m.Dpara
	}
	return 0
}

func (m *Param) GetStrpara() string {
	if m != nil {
		return m.Strpara
	}
	return ""
}

func (m *Param) GetBpara() []byte {
	if m != nil {
		return m.Bpara
	}
	return nil
}

func init() {
	proto.RegisterType((*Param)(nil), "Report.Param")
	proto.RegisterEnum("Report.Param_ParaType", Param_ParaType_name, Param_ParaType_value)
}

func init() { proto.RegisterFile("param.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x48, 0x2c, 0x4a,
	0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x0b, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x51,
	0xba, 0xc7, 0xc8, 0xc5, 0x1a, 0x00, 0x12, 0x17, 0x52, 0xe1, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd3, 0x83, 0x28, 0xd0, 0x03, 0x4b, 0x82, 0xc9, 0x90,
	0xca, 0x82, 0x54, 0x21, 0x5e, 0x2e, 0xd6, 0x3c, 0x90, 0x39, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x2c,
	0x20, 0x6e, 0x0a, 0x98, 0xcb, 0xac, 0xc0, 0xa8, 0xc1, 0x28, 0xc4, 0xcf, 0xc5, 0x5e, 0x5c, 0x52,
	0x04, 0x16, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x04, 0xc9, 0x27, 0x81, 0xb9, 0xac, 0x0a, 0x8c, 0x1a,
	0x3c, 0x4a, 0x85, 0x5c, 0x1c, 0x70, 0x93, 0x38, 0xb8, 0x58, 0xfc, 0x4a, 0x73, 0x72, 0x04, 0x18,
	0x84, 0x38, 0xb9, 0x58, 0x43, 0x3d, 0xfd, 0x42, 0x2c, 0x04, 0x18, 0x85, 0xb8, 0xb8, 0xd8, 0x40,
	0x4c, 0x43, 0x33, 0x01, 0x26, 0x18, 0xdb, 0xd8, 0x48, 0x80, 0x19, 0xc6, 0x36, 0x33, 0x11, 0x60,
	0x01, 0x29, 0x77, 0xf3, 0xf1, 0x77, 0x0c, 0x11, 0x10, 0x00, 0x09, 0xbb, 0xf8, 0x87, 0x3a, 0xf9,
	0xb8, 0x0a, 0x08, 0x82, 0xd8, 0xc1, 0x21, 0x41, 0x9e, 0x7e, 0xee, 0x02, 0x0a, 0x20, 0x25, 0x4e,
	0x91, 0x21, 0xae, 0xc1, 0x02, 0x8a, 0x49, 0x6c, 0x60, 0xff, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x49, 0xf6, 0x91, 0xf1, 0xfe, 0x00, 0x00, 0x00,
}