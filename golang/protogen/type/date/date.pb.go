// Code generated by protoc-gen-go.
// source: protobuf/type/date.proto
// DO NOT EDIT!

/*
Package date is a generated protocol buffer package.

It is generated from these files:
	protobuf/type/date.proto

It has these top-level messages:
	Date
*/
package date

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

// Date represents a whole calendar date.
type Date struct {
	// Year of date.
	Year int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	// Month of year.
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	// Day of month.
	Day int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*Date)(nil), "kudu.type.Date")
}

func init() { proto.RegisterFile("protobuf/type/date.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x49, 0x2c, 0x49, 0xd5, 0x03,
	0x0b, 0x09, 0x71, 0x66, 0x97, 0xa6, 0x94, 0xea, 0x81, 0x44, 0x95, 0x9c, 0xb8, 0x58, 0x5c, 0x12,
	0x4b, 0x52, 0x85, 0x84, 0xb8, 0x58, 0x2a, 0x53, 0x13, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xdc, 0xfc, 0xbc, 0x92, 0x0c, 0x09, 0x26, 0xb0, 0x20,
	0x84, 0x23, 0x24, 0xc0, 0xc5, 0x9c, 0x92, 0x58, 0x29, 0xc1, 0x0c, 0x16, 0x03, 0x31, 0x9d, 0xf4,
	0xa3, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0xf2, 0x52,
	0xf4, 0x41, 0xe6, 0xeb, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0xad, 0x4c, 0x4f, 0xcd,
	0x43, 0xb8, 0x22, 0x89, 0x0d, 0x2c, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xf4, 0x71,
	0x07, 0xa2, 0x00, 0x00, 0x00,
}
