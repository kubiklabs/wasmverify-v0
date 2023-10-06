// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifier/verifier/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	return fileDescriptor_1619b52dd49aed89, []int{0}
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
	return fileDescriptor_1619b52dd49aed89, []int{1}
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

type QueryContractInfoRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryContractInfoRequest) Reset()         { *m = QueryContractInfoRequest{} }
func (m *QueryContractInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractInfoRequest) ProtoMessage()    {}
func (*QueryContractInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1619b52dd49aed89, []int{2}
}
func (m *QueryContractInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractInfoRequest.Merge(m, src)
}
func (m *QueryContractInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractInfoRequest proto.InternalMessageInfo

func (m *QueryContractInfoRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryContractInfoResponse struct {
	Contractinfo *ContractInfo `protobuf:"bytes,1,opt,name=contractinfo,proto3" json:"contractinfo,omitempty"`
}

func (m *QueryContractInfoResponse) Reset()         { *m = QueryContractInfoResponse{} }
func (m *QueryContractInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractInfoResponse) ProtoMessage()    {}
func (*QueryContractInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1619b52dd49aed89, []int{3}
}
func (m *QueryContractInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractInfoResponse.Merge(m, src)
}
func (m *QueryContractInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractInfoResponse proto.InternalMessageInfo

func (m *QueryContractInfoResponse) GetContractinfo() *ContractInfo {
	if m != nil {
		return m.Contractinfo
	}
	return nil
}

type QueryCurrentPendingContractRequest struct {
}

func (m *QueryCurrentPendingContractRequest) Reset()         { *m = QueryCurrentPendingContractRequest{} }
func (m *QueryCurrentPendingContractRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPendingContractRequest) ProtoMessage()    {}
func (*QueryCurrentPendingContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1619b52dd49aed89, []int{4}
}
func (m *QueryCurrentPendingContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentPendingContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentPendingContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentPendingContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPendingContractRequest.Merge(m, src)
}
func (m *QueryCurrentPendingContractRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentPendingContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPendingContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPendingContractRequest proto.InternalMessageInfo

type QueryCurrentPendingContractResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryCurrentPendingContractResponse) Reset()         { *m = QueryCurrentPendingContractResponse{} }
func (m *QueryCurrentPendingContractResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPendingContractResponse) ProtoMessage()    {}
func (*QueryCurrentPendingContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1619b52dd49aed89, []int{5}
}
func (m *QueryCurrentPendingContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentPendingContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentPendingContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentPendingContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPendingContractResponse.Merge(m, src)
}
func (m *QueryCurrentPendingContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentPendingContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPendingContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPendingContractResponse proto.InternalMessageInfo

func (m *QueryCurrentPendingContractResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "verifier.verifier.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "verifier.verifier.QueryParamsResponse")
	proto.RegisterType((*QueryContractInfoRequest)(nil), "verifier.verifier.QueryContractInfoRequest")
	proto.RegisterType((*QueryContractInfoResponse)(nil), "verifier.verifier.QueryContractInfoResponse")
	proto.RegisterType((*QueryCurrentPendingContractRequest)(nil), "verifier.verifier.QueryCurrentPendingContractRequest")
	proto.RegisterType((*QueryCurrentPendingContractResponse)(nil), "verifier.verifier.QueryCurrentPendingContractResponse")
}

func init() { proto.RegisterFile("verifier/verifier/query.proto", fileDescriptor_1619b52dd49aed89) }

var fileDescriptor_1619b52dd49aed89 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8e, 0x12, 0x31,
	0x18, 0xc7, 0x67, 0x70, 0xe5, 0x50, 0x37, 0x26, 0xd6, 0x8d, 0xd9, 0x19, 0x75, 0x56, 0x47, 0x51,
	0xb3, 0xeb, 0x4e, 0x03, 0x04, 0xbd, 0xc3, 0xc9, 0x8b, 0x41, 0x8e, 0x5e, 0xb0, 0x30, 0x65, 0xd2,
	0x44, 0xda, 0xa1, 0x2d, 0x44, 0x34, 0x5e, 0x7c, 0x02, 0x13, 0x2f, 0x3e, 0x8c, 0x0f, 0xc0, 0x91,
	0xc4, 0x8b, 0x17, 0x8d, 0x01, 0x1f, 0xc4, 0xd0, 0x76, 0x50, 0x9c, 0x01, 0xdc, 0x5b, 0x33, 0xdf,
	0xff, 0xff, 0xfd, 0x7f, 0xdf, 0xd7, 0x0e, 0xb8, 0x3d, 0x21, 0x82, 0x0e, 0x28, 0x11, 0x68, 0x7d,
	0x18, 0x8d, 0x89, 0x98, 0x46, 0xa9, 0xe0, 0x8a, 0xc3, 0x6b, 0xd9, 0xd7, 0x28, 0x3b, 0xf8, 0x47,
	0x09, 0x4f, 0xb8, 0xae, 0xa2, 0xd5, 0xc9, 0x08, 0xfd, 0x5b, 0x09, 0xe7, 0xc9, 0x6b, 0x82, 0x70,
	0x4a, 0x11, 0x66, 0x8c, 0x2b, 0xac, 0x28, 0x67, 0xd2, 0x56, 0x4f, 0xfb, 0x5c, 0x0e, 0xb9, 0x44,
	0x3d, 0x2c, 0x89, 0xe9, 0x8f, 0x26, 0xd5, 0x1e, 0x51, 0xb8, 0x8a, 0x52, 0x9c, 0x50, 0xa6, 0xc5,
	0x56, 0x1b, 0xe4, 0x89, 0x52, 0x2c, 0xf0, 0x30, 0xeb, 0x55, 0xc9, 0xd7, 0xfb, 0x9c, 0x29, 0x81,
	0xfb, 0xaa, 0x4b, 0xd9, 0xc0, 0x02, 0x85, 0x47, 0x00, 0xbe, 0x58, 0x05, 0xb5, 0xb5, 0xb7, 0x43,
	0x46, 0x63, 0x22, 0x55, 0xf8, 0x1c, 0x5c, 0xdf, 0xf8, 0x2a, 0x53, 0xce, 0x24, 0x81, 0x4f, 0x41,
	0xd9, 0x64, 0x1c, 0xbb, 0x77, 0xdc, 0x47, 0x57, 0x6a, 0x5e, 0x94, 0x9b, 0x3b, 0x32, 0x96, 0xe6,
	0xc1, 0xec, 0xc7, 0x89, 0xd3, 0xb1, 0xf2, 0xf0, 0x14, 0x1c, 0xeb, 0x7e, 0x2d, 0x4b, 0xf0, 0x8c,
	0x0d, 0xb8, 0xcd, 0x82, 0x57, 0x41, 0x89, 0xc6, 0xba, 0xe1, 0x41, 0xa7, 0x44, 0xe3, 0xf0, 0x15,
	0xf0, 0x0a, 0xb4, 0x96, 0xa0, 0x05, 0x0e, 0xb3, 0x29, 0x56, 0x43, 0x58, 0x8e, 0x93, 0x02, 0x8e,
	0x0d, 0xfb, 0x86, 0x29, 0xbc, 0x0f, 0x42, 0x93, 0x30, 0x16, 0x82, 0x30, 0xd5, 0x26, 0x2c, 0xa6,
	0x2c, 0xc9, 0x0c, 0xd9, 0x0e, 0x1a, 0xe0, 0xde, 0x4e, 0x95, 0x25, 0xfa, 0x07, 0xbf, 0xf6, 0xfd,
	0x12, 0xb8, 0xac, 0x7d, 0xf0, 0x2d, 0x28, 0x9b, 0x65, 0xc0, 0x4a, 0x01, 0x5f, 0x7e, 0xeb, 0xfe,
	0x83, 0x7d, 0x32, 0x13, 0x19, 0xde, 0xfd, 0xf0, 0xf5, 0xd7, 0xa7, 0xd2, 0x4d, 0xe8, 0xa1, 0x6d,
	0x6f, 0x00, 0x7e, 0x76, 0xc1, 0xe1, 0xdf, 0x1b, 0x80, 0x67, 0xdb, 0x7a, 0x17, 0x5c, 0x89, 0xff,
	0xf8, 0xff, 0xc4, 0x16, 0xe7, 0x5c, 0xe3, 0x3c, 0x84, 0x15, 0xb4, 0xe7, 0xc9, 0xa1, 0x77, 0x34,
	0x7e, 0x0f, 0xbf, 0xb8, 0xe0, 0x46, 0xf1, 0x4e, 0x61, 0x63, 0x6b, 0xee, 0xae, 0x9b, 0xf2, 0x9f,
	0x5c, 0xd4, 0x66, 0xc1, 0xeb, 0x1a, 0xfc, 0x1c, 0x9e, 0x15, 0x81, 0x1b, 0x6b, 0x37, 0x35, 0xde,
	0x6e, 0x36, 0x48, 0xb3, 0x3e, 0x5b, 0x04, 0xee, 0x7c, 0x11, 0xb8, 0x3f, 0x17, 0x81, 0xfb, 0x71,
	0x19, 0x38, 0xf3, 0x65, 0xe0, 0x7c, 0x5b, 0x06, 0xce, 0x4b, 0x6f, 0x6d, 0x7e, 0xf3, 0xa7, 0x8f,
	0x9a, 0xa6, 0x44, 0xf6, 0xca, 0xfa, 0x67, 0xab, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x69, 0x94,
	0x73, 0xb2, 0x47, 0x04, 0x00, 0x00,
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
	// Queries a list of ContractInfo items.
	ContractInfo(ctx context.Context, in *QueryContractInfoRequest, opts ...grpc.CallOption) (*QueryContractInfoResponse, error)
	// Queries a list of CurrentPendingContract items.
	CurrentPendingContract(ctx context.Context, in *QueryCurrentPendingContractRequest, opts ...grpc.CallOption) (*QueryCurrentPendingContractResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/verifier.verifier.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractInfo(ctx context.Context, in *QueryContractInfoRequest, opts ...grpc.CallOption) (*QueryContractInfoResponse, error) {
	out := new(QueryContractInfoResponse)
	err := c.cc.Invoke(ctx, "/verifier.verifier.Query/ContractInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CurrentPendingContract(ctx context.Context, in *QueryCurrentPendingContractRequest, opts ...grpc.CallOption) (*QueryCurrentPendingContractResponse, error) {
	out := new(QueryCurrentPendingContractResponse)
	err := c.cc.Invoke(ctx, "/verifier.verifier.Query/CurrentPendingContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ContractInfo items.
	ContractInfo(context.Context, *QueryContractInfoRequest) (*QueryContractInfoResponse, error)
	// Queries a list of CurrentPendingContract items.
	CurrentPendingContract(context.Context, *QueryCurrentPendingContractRequest) (*QueryCurrentPendingContractResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ContractInfo(ctx context.Context, req *QueryContractInfoRequest) (*QueryContractInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractInfo not implemented")
}
func (*UnimplementedQueryServer) CurrentPendingContract(ctx context.Context, req *QueryCurrentPendingContractRequest) (*QueryCurrentPendingContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentPendingContract not implemented")
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
		FullMethod: "/verifier.verifier.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verifier.verifier.Query/ContractInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractInfo(ctx, req.(*QueryContractInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CurrentPendingContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentPendingContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentPendingContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verifier.verifier.Query/CurrentPendingContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentPendingContract(ctx, req.(*QueryCurrentPendingContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "verifier.verifier.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ContractInfo",
			Handler:    _Query_ContractInfo_Handler,
		},
		{
			MethodName: "CurrentPendingContract",
			Handler:    _Query_CurrentPendingContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verifier/verifier/query.proto",
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

func (m *QueryContractInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Contractinfo != nil {
		{
			size, err := m.Contractinfo.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryCurrentPendingContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentPendingContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentPendingContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCurrentPendingContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentPendingContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentPendingContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
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

func (m *QueryContractInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryContractInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Contractinfo != nil {
		l = m.Contractinfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrentPendingContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCurrentPendingContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
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
func (m *QueryContractInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryContractInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryContractInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contractinfo", wireType)
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
			if m.Contractinfo == nil {
				m.Contractinfo = &ContractInfo{}
			}
			if err := m.Contractinfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCurrentPendingContractRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentPendingContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentPendingContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryCurrentPendingContractResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentPendingContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentPendingContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
