// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/doggo.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/doggo.proto

It has these top-level messages:
	Doggo
	GetByIdRequest
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Doggo struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Breed  string `protobuf:"bytes,3,opt,name=breed" json:"breed,omitempty"`
	IsGood bool   `protobuf:"varint,4,opt,name=isGood" json:"isGood,omitempty"`
}

func (m *Doggo) Reset()                    { *m = Doggo{} }
func (m *Doggo) String() string            { return proto.CompactTextString(m) }
func (*Doggo) ProtoMessage()               {}
func (*Doggo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Doggo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Doggo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Doggo) GetBreed() string {
	if m != nil {
		return m.Breed
	}
	return ""
}

func (m *Doggo) GetIsGood() bool {
	if m != nil {
		return m.IsGood
	}
	return false
}

type GetByIdRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetByIdRequest) Reset()                    { *m = GetByIdRequest{} }
func (m *GetByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetByIdRequest) ProtoMessage()               {}
func (*GetByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetByIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Doggo)(nil), "Doggo")
	proto.RegisterType((*GetByIdRequest)(nil), "GetByIdRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DoggoService service

type DoggoServiceClient interface {
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Doggo, error)
}

type doggoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDoggoServiceClient(cc *grpc.ClientConn) DoggoServiceClient {
	return &doggoServiceClient{cc}
}

func (c *doggoServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Doggo, error) {
	out := new(Doggo)
	err := grpc.Invoke(ctx, "/DoggoService/GetById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DoggoService service

type DoggoServiceServer interface {
	GetById(context.Context, *GetByIdRequest) (*Doggo, error)
}

func RegisterDoggoServiceServer(s *grpc.Server, srv DoggoServiceServer) {
	s.RegisterService(&_DoggoService_serviceDesc, srv)
}

func _DoggoService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoggoServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DoggoService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoggoServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DoggoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DoggoService",
	HandlerType: (*DoggoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _DoggoService_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/doggo.proto",
}

func init() { proto.RegisterFile("pb/doggo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0x4f,
	0xc9, 0x4f, 0x4f, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe4, 0x62, 0x75, 0x01,
	0x71, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32,
	0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xa4, 0xa2, 0xd4, 0xd4, 0x14, 0x09, 0x66, 0xb0, 0x20, 0x84,
	0x23, 0x24, 0xc6, 0xc5, 0x96, 0x59, 0xec, 0x9e, 0x9f, 0x9f, 0x22, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1,
	0x11, 0x04, 0xe5, 0x29, 0x29, 0x70, 0xf1, 0xb9, 0xa7, 0x96, 0x38, 0x55, 0x7a, 0xa6, 0x04, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0xa0, 0xdb, 0x61, 0x64, 0xc4, 0xc5, 0x03, 0xb6, 0x3c, 0x38, 0xb5,
	0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x89, 0x8b, 0x1d, 0xaa, 0x43, 0x88, 0x5f, 0x0f, 0x55, 0xaf,
	0x14, 0x9b, 0x1e, 0x58, 0xa9, 0x13, 0x5f, 0x14, 0x4f, 0x41, 0x76, 0x3a, 0xc4, 0x0f, 0xfa, 0x05,
	0x49, 0x49, 0x6c, 0x60, 0x7f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0x1a, 0x5e, 0x2a,
	0xd9, 0x00, 0x00, 0x00,
}
