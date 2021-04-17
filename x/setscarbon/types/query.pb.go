// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: setscarbon/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("setscarbon/query.proto", fileDescriptor_61d1c47f1b3ac04d) }

var fileDescriptor_61d1c47f1b3ac04d = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x2d, 0x29,
	0x4e, 0x4e, 0x2c, 0x4a, 0xca, 0xcf, 0xd3, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x52, 0xf2, 0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0x0a, 0xc8, 0x34, 0x32, 0x34, 0x32,
	0xd1, 0x43, 0x28, 0x42, 0x62, 0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16,
	0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x4c, 0x90,
	0xd2, 0x4a, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0x85, 0x18, 0xad, 0x5f,
	0x66, 0x98, 0x94, 0x5a, 0x92, 0x68, 0xa8, 0x5f, 0x90, 0x98, 0x9e, 0x99, 0x07, 0x56, 0x0c, 0x51,
	0x6b, 0xc4, 0xce, 0xc5, 0x1a, 0x08, 0x52, 0xe1, 0x14, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x66, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x30, 0xb7, 0xe9, 0x82, 0x1c, 0xa7, 0x8f, 0xe4, 0x83, 0x0a, 0x64, 0x4e, 0x49, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0xd8, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x8c, 0x93, 0x96,
	0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MasterPi2124.setscarbon.setscarbon.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "setscarbon/query.proto",
}
