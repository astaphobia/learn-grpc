// Code generated by protoc-gen-go.
// source: customer.proto
// DO NOT EDIT!

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	customer.proto

It has these top-level messages:
	CustomerRequest
	CustomerResponse
	CustomerFilter
*/
package customer

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

// Request message for creating a new customer.
type CustomerRequest struct {
	Id        int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email     string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone     string                     `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Addresses []*CustomerRequest_Address `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *CustomerRequest) Reset()                    { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()               {}
func (*CustomerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CustomerRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CustomerRequest) GetAddresses() []*CustomerRequest_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type CustomerRequest_Address struct {
	Street            string `protobuf:"bytes,1,opt,name=street" json:"street,omitempty"`
	City              string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	State             string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Zip               string `protobuf:"bytes,4,opt,name=zip" json:"zip,omitempty"`
	IsShippingAddress bool   `protobuf:"varint,5,opt,name=isShippingAddress" json:"isShippingAddress,omitempty"`
}

func (m *CustomerRequest_Address) Reset()                    { *m = CustomerRequest_Address{} }
func (m *CustomerRequest_Address) String() string            { return proto.CompactTextString(m) }
func (*CustomerRequest_Address) ProtoMessage()               {}
func (*CustomerRequest_Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CustomerRequest_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *CustomerRequest_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CustomerRequest_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CustomerRequest_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *CustomerRequest_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

// Response message for creating a new customer.
type CustomerResponse struct {
	Id      int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *CustomerResponse) Reset()                    { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string            { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()               {}
func (*CustomerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CustomerResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CustomerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// Request message for filtering Customers.
type CustomerFilter struct {
	Keyword string `protobuf:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *CustomerFilter) Reset()                    { *m = CustomerFilter{} }
func (m *CustomerFilter) String() string            { return proto.CompactTextString(m) }
func (*CustomerFilter) ProtoMessage()               {}
func (*CustomerFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CustomerFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomerRequest)(nil), "customer.CustomerRequest")
	proto.RegisterType((*CustomerRequest_Address)(nil), "customer.CustomerRequest.Address")
	proto.RegisterType((*CustomerResponse)(nil), "customer.CustomerResponse")
	proto.RegisterType((*CustomerFilter)(nil), "customer.CustomerFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Customer service

type CustomerClient interface {
	// Get all Customers with a filter - A server-to-client streaming RPC.
	GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error)
	// Create a new Customer - A simple RPC.
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomers(ctx context.Context, in *CustomerFilter, opts ...grpc.CallOption) (Customer_GetCustomersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Customer_serviceDesc.Streams[0], c.cc, "/customer.Customer/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_GetCustomersClient interface {
	Recv() (*CustomerRequest, error)
	grpc.ClientStream
}

type customerGetCustomersClient struct {
	grpc.ClientStream
}

func (x *customerGetCustomersClient) Recv() (*CustomerRequest, error) {
	m := new(CustomerRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := grpc.Invoke(ctx, "/customer.Customer/CreateCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Customer service

type CustomerServer interface {
	// Get all Customers with a filter - A server-to-client streaming RPC.
	GetCustomers(*CustomerFilter, Customer_GetCustomersServer) error
	// Create a new Customer - A simple RPC.
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).GetCustomers(m, &customerGetCustomersServer{stream})
}

type Customer_GetCustomersServer interface {
	Send(*CustomerRequest) error
	grpc.ServerStream
}

type customerGetCustomersServer struct {
	grpc.ServerStream
}

func (x *customerGetCustomersServer) Send(m *CustomerRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _Customer_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "customer.proto",
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xef, 0x4a, 0xc3, 0x30,
	0x10, 0xc0, 0x97, 0x6e, 0xdd, 0x9f, 0x53, 0xea, 0x0c, 0x22, 0xb1, 0x9f, 0x6a, 0x3f, 0x15, 0x91,
	0x21, 0xf3, 0xab, 0x20, 0x32, 0x70, 0xf8, 0xb5, 0x3e, 0x41, 0x6d, 0x0f, 0x17, 0xdc, 0xda, 0x9a,
	0xcb, 0x90, 0xf9, 0x0a, 0xbe, 0x83, 0xcf, 0xe0, 0x23, 0x4a, 0xd2, 0x66, 0x03, 0xe7, 0xbe, 0xdd,
	0xef, 0x72, 0x77, 0xf9, 0xe5, 0x08, 0x04, 0xf9, 0x9a, 0x74, 0xb5, 0x42, 0x35, 0xa9, 0x55, 0xa5,
	0x2b, 0x3e, 0x74, 0x1c, 0xff, 0x78, 0x70, 0x32, 0x6b, 0x21, 0xc5, 0xf7, 0x35, 0x92, 0xe6, 0x01,
	0x78, 0xb2, 0x10, 0x2c, 0x62, 0x89, 0x9f, 0x7a, 0xb2, 0xe0, 0x1c, 0x7a, 0x65, 0xb6, 0x42, 0xe1,
	0x45, 0x2c, 0x19, 0xa5, 0x36, 0xe6, 0x67, 0xe0, 0xe3, 0x2a, 0x93, 0x4b, 0xd1, 0xb5, 0xc9, 0x06,
	0x4c, 0xb6, 0x5e, 0x54, 0x25, 0x8a, 0x5e, 0x93, 0xb5, 0xc0, 0xef, 0x61, 0x94, 0x15, 0x85, 0x42,
	0x22, 0x24, 0xe1, 0x47, 0xdd, 0xe4, 0x68, 0x7a, 0x39, 0xd9, 0x1a, 0xfd, 0xb9, 0x7d, 0xf2, 0xd0,
	0x94, 0xa6, 0xbb, 0x9e, 0xf0, 0x8b, 0xc1, 0xa0, 0x4d, 0xf3, 0x73, 0xe8, 0x93, 0x56, 0x88, 0xda,
	0x0a, 0x8e, 0xd2, 0x96, 0x8c, 0x64, 0x2e, 0xf5, 0xc6, 0x49, 0x9a, 0xd8, 0xe8, 0x90, 0xce, 0x34,
	0x3a, 0x49, 0x0b, 0x7c, 0x0c, 0xdd, 0x4f, 0x59, 0xb7, 0x8a, 0x26, 0xe4, 0xd7, 0x70, 0x2a, 0xe9,
	0x79, 0x21, 0xeb, 0x5a, 0x96, 0xaf, 0xed, 0x45, 0xc2, 0x8f, 0x58, 0x32, 0x4c, 0xf7, 0x0f, 0xe2,
	0x3b, 0x18, 0xef, 0x9c, 0xa9, 0xae, 0x4a, 0xc2, 0xbd, 0x95, 0x09, 0x18, 0xd0, 0x3a, 0xcf, 0xcd,
	0x1c, 0xcf, 0xce, 0x71, 0x18, 0x5f, 0x41, 0xe0, 0xba, 0x1f, 0xe5, 0x52, 0xa3, 0x32, 0xb5, 0x6f,
	0xb8, 0xf9, 0xa8, 0x54, 0xd1, 0x3e, 0xc9, 0xe1, 0xf4, 0x9b, 0xc1, 0xd0, 0x15, 0xf3, 0x39, 0x1c,
	0xcf, 0x51, 0x3b, 0x24, 0x2e, 0xf6, 0x57, 0xd8, 0x0c, 0x0c, 0x2f, 0x0e, 0x2e, 0x37, 0xee, 0xdc,
	0x30, 0xfe, 0x04, 0xc1, 0x4c, 0x61, 0xa6, 0x71, 0x3b, 0xfa, 0x70, 0x43, 0x18, 0xfe, 0x77, 0xd4,
	0x3c, 0x3a, 0xee, 0xbc, 0xf4, 0xed, 0x77, 0xba, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xde, 0x91,
	0xd3, 0x62, 0x60, 0x02, 0x00, 0x00,
}
