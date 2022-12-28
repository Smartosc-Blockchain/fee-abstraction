// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fachain/interchainquery/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type QueryPendingQueriesRequest struct {
}

func (m *QueryPendingQueriesRequest) Reset()         { *m = QueryPendingQueriesRequest{} }
func (m *QueryPendingQueriesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPendingQueriesRequest) ProtoMessage()    {}
func (*QueryPendingQueriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d0d5c872396e5b, []int{0}
}
func (m *QueryPendingQueriesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPendingQueriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPendingQueriesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPendingQueriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPendingQueriesRequest.Merge(m, src)
}
func (m *QueryPendingQueriesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPendingQueriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPendingQueriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPendingQueriesRequest proto.InternalMessageInfo

type QueryPendingQueriesResponse struct {
	PendingQueries []Query `protobuf:"bytes,1,rep,name=pending_queries,json=pendingQueries,proto3" json:"pending_queries"`
}

func (m *QueryPendingQueriesResponse) Reset()         { *m = QueryPendingQueriesResponse{} }
func (m *QueryPendingQueriesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPendingQueriesResponse) ProtoMessage()    {}
func (*QueryPendingQueriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d0d5c872396e5b, []int{1}
}
func (m *QueryPendingQueriesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPendingQueriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPendingQueriesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPendingQueriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPendingQueriesResponse.Merge(m, src)
}
func (m *QueryPendingQueriesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPendingQueriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPendingQueriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPendingQueriesResponse proto.InternalMessageInfo

func (m *QueryPendingQueriesResponse) GetPendingQueries() []Query {
	if m != nil {
		return m.PendingQueries
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryPendingQueriesRequest)(nil), "fachain.interchainquery.v1.QueryPendingQueriesRequest")
	proto.RegisterType((*QueryPendingQueriesResponse)(nil), "fachain.interchainquery.v1.QueryPendingQueriesResponse")
}

func init() {
	proto.RegisterFile("fachain/interchainquery/v1/query.proto", fileDescriptor_37d0d5c872396e5b)
}

var fileDescriptor_37d0d5c872396e5b = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbd, 0x4e, 0x32, 0x41,
	0x14, 0x86, 0x77, 0xbf, 0xcf, 0x58, 0xac, 0x06, 0x93, 0x8d, 0x85, 0x59, 0xc9, 0xaa, 0x14, 0x86,
	0x86, 0x9d, 0x00, 0x51, 0x0b, 0x0b, 0x13, 0x6a, 0x0b, 0x84, 0x4a, 0x1b, 0x33, 0x2c, 0x87, 0x61,
	0x22, 0xcc, 0x19, 0x66, 0x06, 0x22, 0xad, 0x57, 0x60, 0xe2, 0xfd, 0x58, 0xd3, 0x49, 0x62, 0x63,
	0x65, 0x0c, 0x78, 0x21, 0x86, 0x19, 0x2d, 0xf8, 0x51, 0x63, 0xf7, 0x64, 0xdf, 0xb3, 0xcf, 0xc9,
	0x99, 0x37, 0x38, 0x6c, 0xd1, 0xb4, 0x4d, 0xb9, 0x20, 0x5c, 0x18, 0x50, 0x16, 0x7b, 0x7d, 0x50,
	0x43, 0x32, 0x28, 0x12, 0x0b, 0x89, 0x54, 0x68, 0x30, 0x8c, 0x3e, 0xe7, 0x92, 0x85, 0xb9, 0x64,
	0x50, 0x8c, 0xf2, 0x3f, 0x38, 0x18, 0x08, 0xd0, 0x5c, 0x3b, 0x4b, 0x94, 0x65, 0x88, 0xac, 0x03,
	0x84, 0x4a, 0x4e, 0xa8, 0x10, 0x68, 0xa8, 0xe1, 0x28, 0xbe, 0xd2, 0x6d, 0x86, 0x0c, 0x2d, 0x92,
	0x19, 0xb9, 0xaf, 0xb9, 0x6c, 0x10, 0x5d, 0xcc, 0x6c, 0x55, 0x10, 0x4d, 0x2e, 0xd8, 0x8c, 0x39,
	0xe8, 0x1a, 0xf4, 0xfa, 0xa0, 0x4d, 0x0e, 0x83, 0xdd, 0x95, 0xa9, 0x96, 0x28, 0x34, 0x84, 0xd5,
	0x60, 0x4b, 0xba, 0xe4, 0xba, 0xe7, 0xa2, 0x1d, 0x7f, 0xff, 0x7f, 0x7e, 0xa3, 0x74, 0x90, 0x7c,
	0x7f, 0x50, 0x62, 0x8d, 0x95, 0xb5, 0xd1, 0xeb, 0x9e, 0x57, 0xcb, 0xc8, 0x39, 0x73, 0xe9, 0xc9,
	0x0f, 0x36, 0x6d, 0x5e, 0x07, 0x35, 0xe0, 0x29, 0x84, 0x8f, 0x7e, 0x90, 0x99, 0xdf, 0x1e, 0x1e,
	0xff, 0x2a, 0x5f, 0x79, 0x4c, 0x74, 0xf2, 0xe7, 0xff, 0xdc, 0x99, 0xb9, 0xd3, 0xbb, 0xe7, 0xf7,
	0x87, 0x7f, 0x47, 0x61, 0x99, 0xd4, 0x8d, 0xe2, 0x4d, 0x28, 0x9c, 0xd3, 0x86, 0x26, 0xda, 0xf2,
	0x52, 0x2b, 0x0b, 0x0f, 0x52, 0xb9, 0x1c, 0x4d, 0x62, 0x7f, 0x3c, 0x89, 0xfd, 0xb7, 0x49, 0xec,
	0xdf, 0x4f, 0x63, 0x6f, 0x3c, 0x8d, 0xbd, 0x97, 0x69, 0xec, 0x5d, 0x9d, 0x31, 0x6e, 0xda, 0xfd,
	0x46, 0x92, 0x62, 0x97, 0xd4, 0xbb, 0x54, 0x19, 0xd4, 0x69, 0xa1, 0xd2, 0xc1, 0xf4, 0xc6, 0xf5,
	0xdd, 0xa2, 0x05, 0x07, 0xb7, 0x4b, 0x4b, 0xcc, 0x50, 0x82, 0x6e, 0xac, 0xdb, 0x0a, 0xcb, 0x1f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xea, 0x36, 0x32, 0x66, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	PendingQueries(ctx context.Context, in *QueryPendingQueriesRequest, opts ...grpc.CallOption) (*QueryPendingQueriesResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) PendingQueries(ctx context.Context, in *QueryPendingQueriesRequest, opts ...grpc.CallOption) (*QueryPendingQueriesResponse, error) {
	out := new(QueryPendingQueriesResponse)
	err := c.cc.Invoke(ctx, "/fachain.interchainquery.v1.QueryService/PendingQueries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	PendingQueries(context.Context, *QueryPendingQueriesRequest) (*QueryPendingQueriesResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) PendingQueries(ctx context.Context, req *QueryPendingQueriesRequest) (*QueryPendingQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingQueries not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_PendingQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPendingQueriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PendingQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fachain.interchainquery.v1.QueryService/PendingQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PendingQueries(ctx, req.(*QueryPendingQueriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fachain.interchainquery.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PendingQueries",
			Handler:    _QueryService_PendingQueries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fachain/interchainquery/v1/query.proto",
}

func (m *QueryPendingQueriesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPendingQueriesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPendingQueriesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryPendingQueriesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPendingQueriesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPendingQueriesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PendingQueries) > 0 {
		for iNdEx := len(m.PendingQueries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingQueries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryPendingQueriesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryPendingQueriesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PendingQueries) > 0 {
		for _, e := range m.PendingQueries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryPendingQueriesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryPendingQueriesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPendingQueriesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryPendingQueriesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryPendingQueriesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPendingQueriesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingQueries", wireType)
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
			m.PendingQueries = append(m.PendingQueries, Query{})
			if err := m.PendingQueries[len(m.PendingQueries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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