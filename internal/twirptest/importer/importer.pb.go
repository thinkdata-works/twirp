// Code generated by protoc-gen-go. DO NOT EDIT.
// source: importer.proto

// Test to make sure that importing other packages doesnt break

package importer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/thinkdata-works/twirp/internal/twirptest/importable"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("importer.proto", fileDescriptor_394990ff89e0d02b) }

var fileDescriptor_394990ff89e0d02b = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0x29, 0xcf, 0x2c,
	0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x03, 0x73, 0x4b, 0x52, 0x8b,
	0x4b, 0xf4, 0x60, 0x0a, 0xa5, 0xfc, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x4b, 0x32, 0x32, 0xf3, 0xb2, 0x53, 0x12, 0x4b, 0x12, 0x75, 0xcb, 0xf3, 0x8b, 0xb2, 0x8b,
	0xf5, 0xc1, 0xca, 0xf5, 0x61, 0xba, 0xf5, 0xe1, 0xba, 0xf5, 0x21, 0xba, 0x13, 0x93, 0x72, 0x52,
	0x91, 0x98, 0x10, 0x3b, 0x8d, 0x92, 0xb8, 0x58, 0x82, 0xcb, 0x92, 0x8d, 0x84, 0xa2, 0xb8, 0x58,
	0x82, 0x53, 0xf3, 0x52, 0x84, 0x34, 0xf4, 0x08, 0x38, 0x02, 0xac, 0xd7, 0xb7, 0x38, 0x5d, 0x8a,
	0x68, 0x95, 0x4e, 0x5c, 0x51, 0x1c, 0x30, 0x0f, 0x24, 0xb1, 0x81, 0xad, 0x35, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0xa9, 0x23, 0x27, 0xfc, 0x00, 0x00, 0x00,
}
