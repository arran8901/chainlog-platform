// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contract/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetSmartContractRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryGetSmartContractRequest) Reset()         { *m = QueryGetSmartContractRequest{} }
func (m *QueryGetSmartContractRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSmartContractRequest) ProtoMessage()    {}
func (*QueryGetSmartContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{2}
}
func (m *QueryGetSmartContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSmartContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSmartContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSmartContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSmartContractRequest.Merge(m, src)
}
func (m *QueryGetSmartContractRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSmartContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSmartContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSmartContractRequest proto.InternalMessageInfo

func (m *QueryGetSmartContractRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryGetSmartContractResponse struct {
	SmartContract SmartContract `protobuf:"bytes,1,opt,name=smartContract,proto3" json:"smartContract"`
}

func (m *QueryGetSmartContractResponse) Reset()         { *m = QueryGetSmartContractResponse{} }
func (m *QueryGetSmartContractResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetSmartContractResponse) ProtoMessage()    {}
func (*QueryGetSmartContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{3}
}
func (m *QueryGetSmartContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSmartContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSmartContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSmartContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSmartContractResponse.Merge(m, src)
}
func (m *QueryGetSmartContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSmartContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSmartContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSmartContractResponse proto.InternalMessageInfo

func (m *QueryGetSmartContractResponse) GetSmartContract() SmartContract {
	if m != nil {
		return m.SmartContract
	}
	return SmartContract{}
}

type QueryAllSmartContractRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSmartContractRequest) Reset()         { *m = QueryAllSmartContractRequest{} }
func (m *QueryAllSmartContractRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllSmartContractRequest) ProtoMessage()    {}
func (*QueryAllSmartContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{4}
}
func (m *QueryAllSmartContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSmartContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSmartContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSmartContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSmartContractRequest.Merge(m, src)
}
func (m *QueryAllSmartContractRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSmartContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSmartContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSmartContractRequest proto.InternalMessageInfo

func (m *QueryAllSmartContractRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllSmartContractResponse struct {
	SmartContract []SmartContract     `protobuf:"bytes,1,rep,name=smartContract,proto3" json:"smartContract"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSmartContractResponse) Reset()         { *m = QueryAllSmartContractResponse{} }
func (m *QueryAllSmartContractResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllSmartContractResponse) ProtoMessage()    {}
func (*QueryAllSmartContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{5}
}
func (m *QueryAllSmartContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSmartContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSmartContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSmartContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSmartContractResponse.Merge(m, src)
}
func (m *QueryAllSmartContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSmartContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSmartContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSmartContractResponse proto.InternalMessageInfo

func (m *QueryAllSmartContractResponse) GetSmartContract() []SmartContract {
	if m != nil {
		return m.SmartContract
	}
	return nil
}

func (m *QueryAllSmartContractResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "arran8901.chainlogplatform.contract.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "arran8901.chainlogplatform.contract.QueryParamsResponse")
	proto.RegisterType((*QueryGetSmartContractRequest)(nil), "arran8901.chainlogplatform.contract.QueryGetSmartContractRequest")
	proto.RegisterType((*QueryGetSmartContractResponse)(nil), "arran8901.chainlogplatform.contract.QueryGetSmartContractResponse")
	proto.RegisterType((*QueryAllSmartContractRequest)(nil), "arran8901.chainlogplatform.contract.QueryAllSmartContractRequest")
	proto.RegisterType((*QueryAllSmartContractResponse)(nil), "arran8901.chainlogplatform.contract.QueryAllSmartContractResponse")
}

func init() { proto.RegisterFile("contract/query.proto", fileDescriptor_c26e9eb2398bbc97) }

var fileDescriptor_c26e9eb2398bbc97 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x33, 0xb5, 0x46, 0x1c, 0x29, 0xc8, 0x18, 0xa1, 0x84, 0x76, 0x95, 0x15, 0x54, 0xfc,
	0x33, 0x63, 0x62, 0xb1, 0x11, 0x44, 0x4c, 0x84, 0x16, 0xef, 0xda, 0x78, 0xe7, 0x85, 0x3a, 0xd9,
	0x4e, 0xb7, 0x0b, 0xbb, 0x33, 0xdb, 0x99, 0x89, 0x58, 0x44, 0x04, 0x9f, 0x40, 0xf0, 0x49, 0x7c,
	0x8b, 0x5c, 0x16, 0xbc, 0xd0, 0xab, 0x22, 0x89, 0xcf, 0xe0, 0xb5, 0x64, 0x66, 0x36, 0xcd, 0x1a,
	0x57, 0x37, 0xda, 0xbb, 0x9d, 0x99, 0x73, 0xbe, 0xf3, 0xfd, 0xf6, 0x9c, 0x19, 0x58, 0x0b, 0x04,
	0xd7, 0x92, 0x06, 0x9a, 0xec, 0xf7, 0x99, 0x3c, 0xc0, 0xa9, 0x14, 0x5a, 0xa0, 0x2b, 0x54, 0x4a,
	0xca, 0x5b, 0xf7, 0xef, 0x34, 0x70, 0xb0, 0x47, 0x23, 0x1e, 0x8b, 0x30, 0x8d, 0xa9, 0xde, 0x15,
	0x32, 0xc1, 0x59, 0x42, 0xbd, 0x16, 0x8a, 0x50, 0x98, 0x78, 0x32, 0xfe, 0xb2, 0xa9, 0xf5, 0x95,
	0x50, 0x88, 0x30, 0x66, 0x84, 0xa6, 0x11, 0xa1, 0x9c, 0x0b, 0x4d, 0x75, 0x24, 0xb8, 0x72, 0xa7,
	0x37, 0x02, 0xa1, 0x12, 0xa1, 0x48, 0x8f, 0x2a, 0x66, 0x2b, 0x92, 0x57, 0x8d, 0x1e, 0xd3, 0xb4,
	0x41, 0x52, 0x1a, 0x46, 0xdc, 0x04, 0xbb, 0xd8, 0x8b, 0x13, 0x6b, 0x29, 0x95, 0x34, 0xc9, 0x24,
	0x56, 0x27, 0xdb, 0x2a, 0xa1, 0x52, 0xbf, 0xc8, 0x96, 0xf6, 0xd8, 0xaf, 0x41, 0xb4, 0x3d, 0xd6,
	0xdd, 0x32, 0x39, 0x5d, 0xb6, 0xdf, 0x67, 0x4a, 0xfb, 0x2f, 0xe1, 0x85, 0xdc, 0xae, 0x4a, 0x05,
	0x57, 0x0c, 0x3d, 0x81, 0x55, 0xab, 0xbd, 0x0c, 0x2e, 0x83, 0xeb, 0xe7, 0x9a, 0x37, 0x71, 0x09,
	0x70, 0x6c, 0x45, 0x3a, 0x8b, 0x83, 0xa3, 0x4b, 0x95, 0xae, 0x13, 0xf0, 0x5b, 0x70, 0xc5, 0x54,
	0xd8, 0x64, 0xfa, 0xe9, 0xd8, 0xd7, 0x63, 0x17, 0xed, 0x1c, 0xa0, 0x65, 0x78, 0x86, 0xee, 0xec,
	0x48, 0xa6, 0x6c, 0xad, 0xb3, 0xdd, 0x6c, 0xe9, 0xbf, 0x83, 0xab, 0x05, 0x99, 0xce, 0xe5, 0x73,
	0xb8, 0xa4, 0xa6, 0x0f, 0x9c, 0xd9, 0x66, 0x29, 0xb3, 0x39, 0x49, 0xe7, 0x39, 0x2f, 0xe7, 0xef,
	0x3a, 0xeb, 0xed, 0x38, 0xfe, 0xad, 0xf5, 0x0d, 0x08, 0x8f, 0x9b, 0xe3, 0x8a, 0x5f, 0xc5, 0xb6,
	0x93, 0x78, 0xdc, 0x49, 0x6c, 0x67, 0xc7, 0x75, 0x12, 0x6f, 0xd1, 0x90, 0xb9, 0xdc, 0xee, 0x54,
	0xa6, 0x3f, 0x00, 0x8e, 0x74, 0xb6, 0x50, 0x31, 0xe9, 0xa9, 0x13, 0x24, 0x45, 0x9b, 0x39, 0x92,
	0x05, 0x43, 0x72, 0xed, 0xaf, 0x24, 0xd6, 0xdc, 0x34, 0x4a, 0xf3, 0xc7, 0x22, 0x3c, 0x6d, 0x50,
	0xd0, 0x27, 0x00, 0xab, 0x76, 0x20, 0xd0, 0x7a, 0x29, 0x9b, 0xb3, 0xd3, 0x59, 0x6f, 0xcd, 0x9f,
	0x68, 0x3d, 0xf9, 0x6b, 0xef, 0x3f, 0x7f, 0xff, 0xb8, 0x80, 0xd1, 0x2d, 0x32, 0x51, 0x20, 0x99,
	0xc2, 0xed, 0x4c, 0x82, 0xfc, 0x72, 0x91, 0xd0, 0x11, 0x80, 0x4b, 0xb9, 0xbf, 0x85, 0xda, 0xe5,
	0x1d, 0x14, 0x0c, 0x78, 0xbd, 0xf3, 0x3f, 0x12, 0x0e, 0x67, 0xc3, 0xe0, 0x3c, 0x42, 0x0f, 0xcb,
	0xe1, 0xe4, 0x1f, 0x00, 0xf2, 0xc6, 0xdd, 0xa8, 0xb7, 0xe8, 0x0b, 0x80, 0xe7, 0x73, 0x15, 0xda,
	0x71, 0x3c, 0x0f, 0x63, 0xc1, 0x4d, 0x98, 0x87, 0xb1, 0x68, 0xc6, 0xfd, 0x07, 0x86, 0xf1, 0x1e,
	0x5a, 0xfb, 0x17, 0xc6, 0xce, 0xf6, 0x60, 0xe8, 0x81, 0xc3, 0xa1, 0x07, 0xbe, 0x0d, 0x3d, 0xf0,
	0x61, 0xe4, 0x55, 0x0e, 0x47, 0x5e, 0xe5, 0xeb, 0xc8, 0xab, 0x3c, 0x5b, 0x0f, 0x23, 0xbd, 0xd7,
	0xef, 0xe1, 0x40, 0x24, 0x7f, 0x54, 0x7e, 0x7d, 0xac, 0xad, 0x0f, 0x52, 0xa6, 0x7a, 0x55, 0xf3,
	0x70, 0xde, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x81, 0x3a, 0xf3, 0x0b, 0x06, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a SmartContract by index.
	SmartContract(ctx context.Context, in *QueryGetSmartContractRequest, opts ...grpc.CallOption) (*QueryGetSmartContractResponse, error)
	// Queries a list of SmartContract items.
	SmartContractAll(ctx context.Context, in *QueryAllSmartContractRequest, opts ...grpc.CallOption) (*QueryAllSmartContractResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/arran8901.chainlogplatform.contract.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SmartContract(ctx context.Context, in *QueryGetSmartContractRequest, opts ...grpc.CallOption) (*QueryGetSmartContractResponse, error) {
	out := new(QueryGetSmartContractResponse)
	err := c.cc.Invoke(ctx, "/arran8901.chainlogplatform.contract.Query/SmartContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SmartContractAll(ctx context.Context, in *QueryAllSmartContractRequest, opts ...grpc.CallOption) (*QueryAllSmartContractResponse, error) {
	out := new(QueryAllSmartContractResponse)
	err := c.cc.Invoke(ctx, "/arran8901.chainlogplatform.contract.Query/SmartContractAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a SmartContract by index.
	SmartContract(context.Context, *QueryGetSmartContractRequest) (*QueryGetSmartContractResponse, error)
	// Queries a list of SmartContract items.
	SmartContractAll(context.Context, *QueryAllSmartContractRequest) (*QueryAllSmartContractResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) SmartContract(ctx context.Context, req *QueryGetSmartContractRequest) (*QueryGetSmartContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmartContract not implemented")
}
func (*UnimplementedQueryServer) SmartContractAll(ctx context.Context, req *QueryAllSmartContractRequest) (*QueryAllSmartContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmartContractAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arran8901.chainlogplatform.contract.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SmartContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSmartContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SmartContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arran8901.chainlogplatform.contract.Query/SmartContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SmartContract(ctx, req.(*QueryGetSmartContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SmartContractAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSmartContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SmartContractAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arran8901.chainlogplatform.contract.Query/SmartContractAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SmartContractAll(ctx, req.(*QueryAllSmartContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arran8901.chainlogplatform.contract.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SmartContract",
			Handler:    _Query_SmartContract_Handler,
		},
		{
			MethodName: "SmartContractAll",
			Handler:    _Query_SmartContractAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetSmartContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSmartContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSmartContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetSmartContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSmartContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSmartContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SmartContract.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllSmartContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSmartContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSmartContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllSmartContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSmartContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSmartContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SmartContract) > 0 {
		for iNdEx := len(m.SmartContract) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SmartContract[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetSmartContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetSmartContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SmartContract.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllSmartContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllSmartContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SmartContract) > 0 {
		for _, e := range m.SmartContract {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSmartContractRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSmartContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSmartContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSmartContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSmartContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSmartContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContract", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SmartContract.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllSmartContractRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllSmartContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSmartContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllSmartContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllSmartContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSmartContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContract", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SmartContract = append(m.SmartContract, SmartContract{})
			if err := m.SmartContract[len(m.SmartContract)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
