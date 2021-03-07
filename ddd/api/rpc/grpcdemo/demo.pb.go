// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: demo.proto

package democli

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Demo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Demo) Reset()         { *m = Demo{} }
func (m *Demo) String() string { return proto.CompactTextString(m) }
func (*Demo) ProtoMessage()    {}
func (*Demo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}
func (m *Demo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Demo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Demo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Demo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Demo.Merge(m, src)
}
func (m *Demo) XXX_Size() int {
	return m.Size()
}
func (m *Demo) XXX_DiscardUnknown() {
	xxx_messageInfo_Demo.DiscardUnknown(m)
}

var xxx_messageInfo_Demo proto.InternalMessageInfo

func (m *Demo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Demo) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Demo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetDemosReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDemosReq) Reset()         { *m = GetDemosReq{} }
func (m *GetDemosReq) String() string { return proto.CompactTextString(m) }
func (*GetDemosReq) ProtoMessage()    {}
func (*GetDemosReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}
func (m *GetDemosReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDemosReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDemosReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDemosReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDemosReq.Merge(m, src)
}
func (m *GetDemosReq) XXX_Size() int {
	return m.Size()
}
func (m *GetDemosReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDemosReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetDemosReq proto.InternalMessageInfo

func (m *GetDemosReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetDemosReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetDemosReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetDemosReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetDemosResp struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Demos                []*Demo  `protobuf:"bytes,2,rep,name=demos,proto3" json:"demos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDemosResp) Reset()         { *m = GetDemosResp{} }
func (m *GetDemosResp) String() string { return proto.CompactTextString(m) }
func (*GetDemosResp) ProtoMessage()    {}
func (*GetDemosResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}
func (m *GetDemosResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDemosResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDemosResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDemosResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDemosResp.Merge(m, src)
}
func (m *GetDemosResp) XXX_Size() int {
	return m.Size()
}
func (m *GetDemosResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDemosResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetDemosResp proto.InternalMessageInfo

func (m *GetDemosResp) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetDemosResp) GetDemos() []*Demo {
	if m != nil {
		return m.Demos
	}
	return nil
}

func init() {
	proto.RegisterType((*Demo)(nil), "democli.Demo")
	proto.RegisterType((*GetDemosReq)(nil), "democli.GetDemosReq")
	proto.RegisterType((*GetDemosResp)(nil), "democli.GetDemosResp")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0xb1, 0x93, 0x73, 0x32, 0x95, 0x3c, 0xb9,
	0x58, 0x5c, 0x52, 0x73, 0xf3, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x98, 0x83, 0x98, 0x32, 0x53, 0x84, 0x64, 0xb9, 0xb8, 0x92, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53,
	0xe2, 0x13, 0x4b, 0x24, 0x98, 0xc0, 0xe2, 0x9c, 0x50, 0x11, 0xc7, 0x12, 0x21, 0x21, 0x2e, 0x96,
	0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x29, 0x9e, 0x8b,
	0xdb, 0x3d, 0xb5, 0x04, 0x64, 0x5a, 0x71, 0x50, 0x6a, 0x21, 0x86, 0x89, 0x30, 0x2d, 0x4c, 0x08,
	0x2d, 0x42, 0x62, 0x5c, 0x6c, 0xf9, 0x69, 0x69, 0xc5, 0xa9, 0x25, 0x60, 0x83, 0x58, 0x83, 0xa0,
	0x3c, 0x21, 0x11, 0x2e, 0xd6, 0x9c, 0xcc, 0xdc, 0xcc, 0x12, 0x09, 0x16, 0xb0, 0x30, 0x84, 0xa3,
	0xe4, 0xc9, 0xc5, 0x83, 0xb0, 0xa0, 0xb8, 0x00, 0xa4, 0x2a, 0x39, 0xbf, 0x34, 0xaf, 0x04, 0x6a,
	0x09, 0x84, 0x23, 0xa4, 0xcc, 0xc5, 0x0a, 0xf2, 0x5c, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7,
	0x11, 0xaf, 0x1e, 0xd4, 0xab, 0x7a, 0x20, 0x8d, 0x41, 0x10, 0x39, 0x23, 0x37, 0x2e, 0x6e, 0x10,
	0x37, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x9c, 0x8b, 0x03, 0x66, 0xb2, 0x90, 0x08,
	0x5c, 0x03, 0x92, 0x6f, 0xa4, 0x44, 0xb1, 0x88, 0x16, 0x17, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81,
	0x03, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x63, 0x13, 0x95, 0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	GetDemos(ctx context.Context, in *GetDemosReq, opts ...grpc.CallOption) (*GetDemosResp, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) GetDemos(ctx context.Context, in *GetDemosReq, opts ...grpc.CallOption) (*GetDemosResp, error) {
	out := new(GetDemosResp)
	err := c.cc.Invoke(ctx, "/democli.DemoService/GetDemos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	GetDemos(context.Context, *GetDemosReq) (*GetDemosResp, error)
}

// UnimplementedDemoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (*UnimplementedDemoServiceServer) GetDemos(ctx context.Context, req *GetDemosReq) (*GetDemosResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDemos not implemented")
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_GetDemos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDemosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).GetDemos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/democli.DemoService/GetDemos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).GetDemos(ctx, req.(*GetDemosReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "democli.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDemos",
			Handler:    _DemoService_GetDemos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

func (m *Demo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Demo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Demo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDemo(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CreatedAt != 0 {
		i = encodeVarintDemo(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintDemo(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetDemosReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDemosReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDemosReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Limit != 0 {
		i = encodeVarintDemo(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x20
	}
	if m.Offset != 0 {
		i = encodeVarintDemo(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDemo(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintDemo(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetDemosResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDemosResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDemosResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Demos) > 0 {
		for iNdEx := len(m.Demos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Demos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDemo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Count != 0 {
		i = encodeVarintDemo(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDemo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDemo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Demo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDemo(uint64(m.Id))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovDemo(uint64(m.CreatedAt))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDemo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetDemosReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDemo(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDemo(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovDemo(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovDemo(uint64(m.Limit))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetDemosResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovDemo(uint64(m.Count))
	}
	if len(m.Demos) > 0 {
		for _, e := range m.Demos {
			l = e.Size()
			n += 1 + l + sovDemo(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDemo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDemo(x uint64) (n int) {
	return sovDemo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Demo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
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
			return fmt.Errorf("proto: Demo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Demo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
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
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetDemosReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
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
			return fmt.Errorf("proto: GetDemosReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDemosReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
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
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetDemosResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
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
			return fmt.Errorf("proto: GetDemosResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDemosResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Demos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
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
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Demos = append(m.Demos, &Demo{})
			if err := m.Demos[len(m.Demos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDemo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDemo
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
					return 0, ErrIntOverflowDemo
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
					return 0, ErrIntOverflowDemo
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
				return 0, ErrInvalidLengthDemo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDemo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDemo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDemo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDemo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDemo = fmt.Errorf("proto: unexpected end of group")
)
