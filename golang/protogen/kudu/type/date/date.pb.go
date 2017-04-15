// Code generated by protoc-gen-go.
// source: type/date.proto
// DO NOT EDIT!

/*
Package date is a generated protocol buffer package.

It is generated from these files:
	type/date.proto

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

func init() { proto.RegisterFile("type/date.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xa9, 0x2c, 0x48,
	0xd5, 0x4f, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x2e, 0x4d,
	0x29, 0xd5, 0x03, 0x89, 0x2a, 0x39, 0x71, 0xb1, 0xb8, 0x24, 0x96, 0xa4, 0x0a, 0x09, 0x71, 0xb1,
	0x54, 0xa6, 0x26, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c,
	0xac, 0xb9, 0xf9, 0x79, 0x25, 0x19, 0x12, 0x4c, 0x60, 0x41, 0x08, 0x47, 0x48, 0x80, 0x8b, 0x39,
	0x25, 0xb1, 0x52, 0x82, 0x19, 0x2c, 0x06, 0x62, 0x3a, 0x99, 0x44, 0x19, 0xa5, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0xe5, 0xa5, 0xe8, 0x83, 0xcc, 0xd7, 0x4f, 0xcf,
	0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x07, 0x5b, 0x99, 0x9e, 0x9a, 0x07, 0x11, 0x84, 0x3b, 0x25, 0x89,
	0x0d, 0x2c, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x61, 0xfb, 0x54, 0x9e, 0x00, 0x00,
	0x00,
}