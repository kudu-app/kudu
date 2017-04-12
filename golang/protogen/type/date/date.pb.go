// Code generated by protoc-gen-go.
// source: date.proto
// DO NOT EDIT!

/*
Package date is a generated protocol buffer package.

It is generated from these files:
	date.proto

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

func init() { proto.RegisterFile("date.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x2e, 0x4d, 0x29, 0xd5, 0x2b, 0xa9, 0x2c,
	0x48, 0x55, 0x72, 0xe2, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x15, 0x12, 0xe2, 0x62, 0xa9, 0x4c, 0x4d,
	0x2c, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x73, 0xf3,
	0xf3, 0x4a, 0x32, 0x24, 0x98, 0xc0, 0x82, 0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x4a, 0x62, 0xa5,
	0x04, 0x33, 0x58, 0x0c, 0xc4, 0x74, 0x32, 0x89, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xca, 0x4b, 0xd1, 0x07, 0x99, 0xaf, 0x9f, 0x9e, 0x9f, 0x93, 0x98,
	0x97, 0xae, 0x0f, 0xb6, 0x32, 0x3d, 0x35, 0x4f, 0x1f, 0x64, 0x9f, 0x3e, 0xc8, 0x15, 0xd6, 0x20,
	0x22, 0x89, 0x0d, 0x2c, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x2a, 0x45, 0xca, 0x99,
	0x00, 0x00, 0x00,
}