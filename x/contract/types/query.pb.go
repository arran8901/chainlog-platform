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

type QueryContractCodeRequest struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (m *QueryContractCodeRequest) Reset()         { *m = QueryContractCodeRequest{} }
func (m *QueryContractCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractCodeRequest) ProtoMessage()    {}
func (*QueryContractCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{2}
}
func (m *QueryContractCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractCodeRequest.Merge(m, src)
}
func (m *QueryContractCodeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractCodeRequest proto.InternalMessageInfo

func (m *QueryContractCodeRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

type QueryContractCodeResponse struct {
	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	DynamicKb string `protobuf:"bytes,2,opt,name=dynamicKb,proto3" json:"dynamicKb,omitempty"`
}

func (m *QueryContractCodeResponse) Reset()         { *m = QueryContractCodeResponse{} }
func (m *QueryContractCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractCodeResponse) ProtoMessage()    {}
func (*QueryContractCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{3}
}
func (m *QueryContractCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractCodeResponse.Merge(m, src)
}
func (m *QueryContractCodeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractCodeResponse proto.InternalMessageInfo

func (m *QueryContractCodeResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *QueryContractCodeResponse) GetDynamicKb() string {
	if m != nil {
		return m.DynamicKb
	}
	return ""
}

type QueryAllContractsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllContractsRequest) Reset()         { *m = QueryAllContractsRequest{} }
func (m *QueryAllContractsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllContractsRequest) ProtoMessage()    {}
func (*QueryAllContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{4}
}
func (m *QueryAllContractsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllContractsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllContractsRequest.Merge(m, src)
}
func (m *QueryAllContractsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllContractsRequest proto.InternalMessageInfo

func (m *QueryAllContractsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllContractsResponse struct {
	Contracts  []SmartContractWithAddress `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts"`
	Pagination *query.PageResponse        `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllContractsResponse) Reset()         { *m = QueryAllContractsResponse{} }
func (m *QueryAllContractsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllContractsResponse) ProtoMessage()    {}
func (*QueryAllContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26e9eb2398bbc97, []int{5}
}
func (m *QueryAllContractsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllContractsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllContractsResponse.Merge(m, src)
}
func (m *QueryAllContractsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllContractsResponse proto.InternalMessageInfo

func (m *QueryAllContractsResponse) GetContracts() []SmartContractWithAddress {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *QueryAllContractsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "arran8901.chainlogplatform.contract.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "arran8901.chainlogplatform.contract.QueryParamsResponse")
	proto.RegisterType((*QueryContractCodeRequest)(nil), "arran8901.chainlogplatform.contract.QueryContractCodeRequest")
	proto.RegisterType((*QueryContractCodeResponse)(nil), "arran8901.chainlogplatform.contract.QueryContractCodeResponse")
	proto.RegisterType((*QueryAllContractsRequest)(nil), "arran8901.chainlogplatform.contract.QueryAllContractsRequest")
	proto.RegisterType((*QueryAllContractsResponse)(nil), "arran8901.chainlogplatform.contract.QueryAllContractsResponse")
}

func init() { proto.RegisterFile("contract/query.proto", fileDescriptor_c26e9eb2398bbc97) }

var fileDescriptor_c26e9eb2398bbc97 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0x35, 0x06, 0x32, 0x2d, 0x08, 0x63, 0x84, 0x18, 0xea, 0x2a, 0x2b, 0x68, 0xf0,
	0x63, 0xc6, 0xa4, 0x4a, 0x2b, 0x52, 0xa1, 0xad, 0x5a, 0x44, 0x0a, 0x6d, 0xbc, 0x10, 0xbc, 0xa9,
	0xb3, 0xbb, 0xe3, 0x66, 0x61, 0x77, 0x67, 0xbb, 0x33, 0x11, 0x83, 0x78, 0xe3, 0x13, 0x08, 0x3e,
	0x89, 0x4f, 0x61, 0xf1, 0xaa, 0xe0, 0x8d, 0x17, 0x22, 0x92, 0x08, 0xbe, 0x86, 0xec, 0x7c, 0xe4,
	0xcb, 0xa0, 0x1b, 0xef, 0x66, 0x67, 0xe6, 0xfc, 0xcf, 0xff, 0x77, 0xce, 0xd9, 0x01, 0x35, 0x8f,
	0x25, 0x22, 0x23, 0x9e, 0xc0, 0x47, 0x3d, 0x9a, 0xf5, 0x51, 0x9a, 0x31, 0xc1, 0xe0, 0x65, 0x92,
	0x65, 0x24, 0xd9, 0xb8, 0x7b, 0xab, 0x85, 0xbc, 0x2e, 0x09, 0x93, 0x88, 0x05, 0x69, 0x44, 0xc4,
	0x4b, 0x96, 0xc5, 0xc8, 0x04, 0x34, 0x6a, 0x01, 0x0b, 0x98, 0xbc, 0x8f, 0xf3, 0x95, 0x0a, 0x6d,
	0xac, 0x06, 0x8c, 0x05, 0x11, 0xc5, 0x24, 0x0d, 0x31, 0x49, 0x12, 0x26, 0x88, 0x08, 0x59, 0xc2,
	0xf5, 0xe9, 0x35, 0x8f, 0xf1, 0x98, 0x71, 0xec, 0x12, 0x4e, 0x55, 0x46, 0xfc, 0xaa, 0xe5, 0x52,
	0x41, 0x5a, 0x38, 0x25, 0x41, 0x98, 0xc8, 0xcb, 0xfa, 0xee, 0xb9, 0x91, 0xb5, 0x94, 0x64, 0x24,
	0x36, 0x12, 0x17, 0x46, 0xdb, 0x3c, 0x26, 0x99, 0x38, 0x34, 0x9f, 0xea, 0xd8, 0xa9, 0x01, 0x78,
	0x90, 0xeb, 0xee, 0xcb, 0x98, 0x0e, 0x3d, 0xea, 0x51, 0x2e, 0x9c, 0x17, 0xe0, 0xec, 0xd4, 0x2e,
	0x4f, 0x59, 0xc2, 0x29, 0x7c, 0x0c, 0x2a, 0x4a, 0xbb, 0x6e, 0x5d, 0xb2, 0x9a, 0xcb, 0xed, 0xeb,
	0xa8, 0x00, 0x38, 0x52, 0x22, 0xdb, 0xe5, 0xe3, 0xef, 0x17, 0x4b, 0x1d, 0x2d, 0xe0, 0x3c, 0x00,
	0x75, 0x99, 0x61, 0x47, 0xdf, 0xda, 0x61, 0x3e, 0xd5, 0xd9, 0x61, 0x13, 0x9c, 0x31, 0xc1, 0x5b,
	0xbe, 0x9f, 0x51, 0xae, 0xf2, 0x55, 0x3b, 0xb3, 0xdb, 0xce, 0x1e, 0x38, 0x3f, 0x47, 0x45, 0xbb,
	0x85, 0xa0, 0xec, 0x31, 0x9f, 0xea, 0x58, 0xb9, 0x86, 0xab, 0xa0, 0xea, 0xf7, 0x13, 0x12, 0x87,
	0xde, 0x13, 0xb7, 0xbe, 0x24, 0x0f, 0xc6, 0x1b, 0x8e, 0xab, 0x4d, 0x6d, 0x45, 0x91, 0x51, 0x34,
	0x25, 0x81, 0x8f, 0x00, 0x18, 0x97, 0x5c, 0xf3, 0x5f, 0x41, 0xaa, 0x3f, 0x28, 0xef, 0x0f, 0x52,
	0x13, 0xa1, 0xfb, 0x83, 0xf6, 0x49, 0x60, 0x80, 0x3a, 0x13, 0x91, 0xce, 0x27, 0x4b, 0x7b, 0x9e,
	0x4e, 0xa2, 0x3d, 0x13, 0x50, 0x35, 0x8c, 0x39, 0xf4, 0xa9, 0xe6, 0x72, 0x7b, 0xb3, 0x50, 0x91,
	0x9f, 0xe6, 0xcd, 0x35, 0x7a, 0xcf, 0x42, 0xd1, 0xd5, 0x25, 0xd2, 0x65, 0x1f, 0xab, 0xc2, 0xdd,
	0x29, 0x90, 0x25, 0x09, 0x72, 0xf5, 0x9f, 0x20, 0xca, 0xdf, 0x24, 0x49, 0xfb, 0x57, 0x19, 0x9c,
	0x96, 0x24, 0xf0, 0xa3, 0x05, 0x2a, 0xaa, 0xcb, 0x70, 0xbd, 0x90, 0xdb, 0x3f, 0x47, 0xae, 0xb1,
	0xb1, 0x78, 0xa0, 0xf2, 0xe4, 0xdc, 0x7e, 0xf7, 0xe5, 0xe7, 0x87, 0x25, 0x04, 0x6f, 0xe0, 0x91,
	0x02, 0x36, 0x0a, 0x37, 0x8d, 0x04, 0x9e, 0xf9, 0x3b, 0xe0, 0x37, 0x0b, 0xac, 0x4c, 0x8e, 0x0d,
	0xdc, 0x2c, 0x6e, 0x60, 0xce, 0xd0, 0x36, 0xee, 0xff, 0x6f, 0xb8, 0xa6, 0xd8, 0x93, 0x14, 0xbb,
	0xf0, 0x61, 0x31, 0x0a, 0xb3, 0x38, 0xcc, 0xc7, 0x1a, 0xbf, 0x99, 0xf9, 0x31, 0xde, 0xc2, 0xcf,
	0x16, 0x58, 0x99, 0x9c, 0xb0, 0x45, 0xf0, 0xe6, 0x8c, 0xff, 0x22, 0x78, 0xf3, 0x06, 0xdb, 0xb9,
	0x27, 0xf1, 0xee, 0xc0, 0xb5, 0x62, 0x78, 0x24, 0x8a, 0x46, 0x2f, 0x15, 0xdf, 0x3e, 0x38, 0x1e,
	0xd8, 0xd6, 0xc9, 0xc0, 0xb6, 0x7e, 0x0c, 0x6c, 0xeb, 0xfd, 0xd0, 0x2e, 0x9d, 0x0c, 0xed, 0xd2,
	0xd7, 0xa1, 0x5d, 0x7a, 0xbe, 0x1e, 0x84, 0xa2, 0xdb, 0x73, 0x91, 0xc7, 0xe2, 0xbf, 0x0a, 0xbf,
	0x1e, 0x4b, 0x8b, 0x7e, 0x4a, 0xb9, 0x5b, 0x91, 0xcf, 0xdf, 0xda, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x77, 0x64, 0xae, 0x39, 0xd1, 0x05, 0x00, 0x00,
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
	// Queries a list of ContractCode items.
	ContractCode(ctx context.Context, in *QueryContractCodeRequest, opts ...grpc.CallOption) (*QueryContractCodeResponse, error)
	// Queries a list of AllContracts items.
	AllContracts(ctx context.Context, in *QueryAllContractsRequest, opts ...grpc.CallOption) (*QueryAllContractsResponse, error)
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

func (c *queryClient) ContractCode(ctx context.Context, in *QueryContractCodeRequest, opts ...grpc.CallOption) (*QueryContractCodeResponse, error) {
	out := new(QueryContractCodeResponse)
	err := c.cc.Invoke(ctx, "/arran8901.chainlogplatform.contract.Query/ContractCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllContracts(ctx context.Context, in *QueryAllContractsRequest, opts ...grpc.CallOption) (*QueryAllContractsResponse, error) {
	out := new(QueryAllContractsResponse)
	err := c.cc.Invoke(ctx, "/arran8901.chainlogplatform.contract.Query/AllContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ContractCode items.
	ContractCode(context.Context, *QueryContractCodeRequest) (*QueryContractCodeResponse, error)
	// Queries a list of AllContracts items.
	AllContracts(context.Context, *QueryAllContractsRequest) (*QueryAllContractsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ContractCode(ctx context.Context, req *QueryContractCodeRequest) (*QueryContractCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCode not implemented")
}
func (*UnimplementedQueryServer) AllContracts(ctx context.Context, req *QueryAllContractsRequest) (*QueryAllContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllContracts not implemented")
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

func _Query_ContractCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arran8901.chainlogplatform.contract.Query/ContractCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractCode(ctx, req.(*QueryContractCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arran8901.chainlogplatform.contract.Query/AllContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllContracts(ctx, req.(*QueryAllContractsRequest))
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
			MethodName: "ContractCode",
			Handler:    _Query_ContractCode_Handler,
		},
		{
			MethodName: "AllContracts",
			Handler:    _Query_AllContracts_Handler,
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

func (m *QueryContractCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DynamicKb) > 0 {
		i -= len(m.DynamicKb)
		copy(dAtA[i:], m.DynamicKb)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DynamicKb)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllContractsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllContractsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllContractsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllContractsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllContractsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllContractsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contracts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryContractCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.DynamicKb)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllContractsRequest) Size() (n int) {
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

func (m *QueryAllContractsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for _, e := range m.Contracts {
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
func (m *QueryContractCodeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryContractCodeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
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
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicKb", wireType)
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
			m.DynamicKb = string(dAtA[iNdEx:postIndex])
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
func (m *QueryAllContractsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllContractsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllContractsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllContractsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllContractsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllContractsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
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
			m.Contracts = append(m.Contracts, SmartContractWithAddress{})
			if err := m.Contracts[len(m.Contracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
