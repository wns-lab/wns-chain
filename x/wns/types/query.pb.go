// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wns/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type QueryMetaDataRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *QueryMetaDataRequest) Reset()         { *m = QueryMetaDataRequest{} }
func (m *QueryMetaDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMetaDataRequest) ProtoMessage()    {}
func (*QueryMetaDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c79f8ca2ab11f3be, []int{0}
}
func (m *QueryMetaDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMetaDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMetaDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMetaDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMetaDataRequest.Merge(m, src)
}
func (m *QueryMetaDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMetaDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMetaDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMetaDataRequest proto.InternalMessageInfo

func (m *QueryMetaDataRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type QueryMetaDataResponse struct {
	Metadata *MetaData `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *QueryMetaDataResponse) Reset()         { *m = QueryMetaDataResponse{} }
func (m *QueryMetaDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMetaDataResponse) ProtoMessage()    {}
func (*QueryMetaDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c79f8ca2ab11f3be, []int{1}
}
func (m *QueryMetaDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMetaDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMetaDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMetaDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMetaDataResponse.Merge(m, src)
}
func (m *QueryMetaDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMetaDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMetaDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMetaDataResponse proto.InternalMessageInfo

func (m *QueryMetaDataResponse) GetMetadata() *MetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryMetaDataRequest)(nil), "wns.QueryMetaDataRequest")
	proto.RegisterType((*QueryMetaDataResponse)(nil), "wns.QueryMetaDataResponse")
}

func init() { proto.RegisterFile("wns/query.proto", fileDescriptor_c79f8ca2ab11f3be) }

var fileDescriptor_c79f8ca2ab11f3be = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x6e, 0xfc, 0xa3, 0xae, 0x14, 0x61, 0xa9, 0x60, 0x53, 0x59, 0xb4, 0xa7, 0x2a, 0x98, 0x85,
	0xfa, 0x06, 0xc1, 0x6b, 0x0f, 0xf6, 0xe8, 0x6d, 0xd2, 0x8e, 0x69, 0xa0, 0xd9, 0x4d, 0xbb, 0x13,
	0x62, 0x11, 0x2f, 0x3e, 0x81, 0xe0, 0x4b, 0x79, 0x2c, 0x78, 0xf1, 0x28, 0x89, 0x0f, 0x22, 0xbb,
	0x4d, 0x7b, 0x90, 0x9e, 0x76, 0x66, 0xbf, 0x5f, 0x86, 0x9d, 0x16, 0xca, 0xc8, 0x79, 0x8e, 0x8b,
	0x65, 0x90, 0x2d, 0x34, 0x69, 0xbe, 0x5f, 0x28, 0xe3, 0xb7, 0x63, 0x1d, 0x6b, 0xb7, 0x4b, 0x3b,
	0xad, 0x21, 0xbf, 0x65, 0xb9, 0x85, 0x32, 0xf5, 0xda, 0x8d, 0xb5, 0x8e, 0x67, 0x28, 0xdd, 0x16,
	0xe5, 0x4f, 0x12, 0xd3, 0x8c, 0x6a, 0x1b, 0xff, 0xa2, 0x06, 0x21, 0x4b, 0x24, 0x28, 0xa5, 0x09,
	0x28, 0xd1, 0x1b, 0x69, 0xef, 0x86, 0xb5, 0x1f, 0x6c, 0xe6, 0x10, 0x09, 0xee, 0x81, 0x60, 0x84,
	0xf3, 0x1c, 0x0d, 0x71, 0xce, 0x0e, 0x14, 0xa4, 0x78, 0xee, 0x5d, 0x7a, 0xfd, 0xe3, 0x91, 0x9b,
	0x7b, 0x21, 0x3b, 0xfb, 0xc7, 0x35, 0x99, 0x56, 0x06, 0xf9, 0x35, 0x6b, 0xa6, 0x48, 0x30, 0x01,
	0x02, 0x27, 0x38, 0x19, 0xb4, 0x02, 0xdb, 0x6e, 0x4b, 0xdc, 0xc2, 0x83, 0x94, 0x1d, 0x3a, 0x0f,
	0x3e, 0x61, 0xcd, 0x61, 0xfd, 0xc9, 0x3b, 0x8e, 0xbd, 0xab, 0x87, 0xef, 0xef, 0x82, 0xd6, 0xb1,
	0xbd, 0xab, 0xb7, 0xaf, 0xdf, 0x8f, 0xbd, 0x2e, 0xef, 0xd8, 0x53, 0xc8, 0x08, 0x09, 0xe4, 0x26,
	0x47, 0xbe, 0xd8, 0xc6, 0xaf, 0x61, 0xf8, 0x59, 0x0a, 0x6f, 0x55, 0x0a, 0xef, 0xa7, 0x14, 0xde,
	0x7b, 0x25, 0x1a, 0xab, 0x4a, 0x34, 0xbe, 0x2b, 0xd1, 0x78, 0xec, 0xc7, 0x09, 0x4d, 0xf3, 0x28,
	0x18, 0xeb, 0xd4, 0xca, 0x6f, 0x67, 0x10, 0xb9, 0x77, 0x3c, 0x85, 0x44, 0xc9, 0x67, 0x67, 0x49,
	0xcb, 0x0c, 0x4d, 0x74, 0xe4, 0x2e, 0x75, 0xf7, 0x17, 0x00, 0x00, 0xff, 0xff, 0x74, 0x97, 0x10,
	0x0c, 0xa1, 0x01, 0x00, 0x00,
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
	Metadata(ctx context.Context, in *QueryMetaDataRequest, opts ...grpc.CallOption) (*QueryMetaDataResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Metadata(ctx context.Context, in *QueryMetaDataRequest, opts ...grpc.CallOption) (*QueryMetaDataResponse, error) {
	out := new(QueryMetaDataResponse)
	err := c.cc.Invoke(ctx, "/wns.Query/Metadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Metadata(context.Context, *QueryMetaDataRequest) (*QueryMetaDataResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Metadata(ctx context.Context, req *QueryMetaDataRequest) (*QueryMetaDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metadata not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Metadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetaDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Metadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wns.Query/Metadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Metadata(ctx, req.(*QueryMetaDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wns.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Metadata",
			Handler:    _Query_Metadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wns/query.proto",
}

func (m *QueryMetaDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMetaDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMetaDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryMetaDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMetaDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMetaDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryMetaDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryMetaDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
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
func (m *QueryMetaDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMetaDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMetaDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
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
func (m *QueryMetaDataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMetaDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMetaDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if m.Metadata == nil {
				m.Metadata = &MetaData{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
