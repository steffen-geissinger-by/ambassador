// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/route/route.proto

package envoy_api_v2_route

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("envoy/api/v2/route/route.proto", fileDescriptor_7dc2895fb75d9d41) }

var fileDescriptor_7dc2895fb75d9d41 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd2, 0x2f, 0xca, 0x2f, 0x2d, 0x49, 0x85, 0x90,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x60, 0x79, 0xbd, 0xc4, 0x82, 0x4c, 0xbd, 0x32,
	0x23, 0x3d, 0xb0, 0x8c, 0x94, 0x6c, 0x69, 0x4a, 0x41, 0xa2, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49,
	0x62, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0x7e, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0x44, 0x8b, 0x94,
	0x26, 0x2e, 0x23, 0xe3, 0x93, 0xf3, 0x73, 0x0b, 0xf2, 0xf3, 0x52, 0xf3, 0x4a, 0xa0, 0x4a, 0x9d,
	0x6c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46, 0x2e, 0x85,
	0xcc, 0x7c, 0x3d, 0xb0, 0xde, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x4c, 0x9b, 0x9d, 0xb8, 0x82,
	0x40, 0x54, 0x00, 0x48, 0x6f, 0x00, 0x63, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x18, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0x15, 0x43, 0xb9, 0xc6, 0x00, 0x00, 0x00,
}
