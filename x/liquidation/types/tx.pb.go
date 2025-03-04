// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: petri/liquidation/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgLiquidateVaultRequest struct {
	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	AppId   uint64 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	VaultId uint64 `protobuf:"varint,3,opt,name=vault_id,json=vaultId,proto3" json:"vault_id,omitempty" yaml:"vault_id"`
}

func (m *MsgLiquidateVaultRequest) Reset()         { *m = MsgLiquidateVaultRequest{} }
func (m *MsgLiquidateVaultRequest) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidateVaultRequest) ProtoMessage()    {}
func (*MsgLiquidateVaultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d2beddde56424c, []int{0}
}
func (m *MsgLiquidateVaultRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidateVaultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidateVaultRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidateVaultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidateVaultRequest.Merge(m, src)
}
func (m *MsgLiquidateVaultRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidateVaultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidateVaultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidateVaultRequest proto.InternalMessageInfo

type MsgLiquidateVaultResponse struct {
}

func (m *MsgLiquidateVaultResponse) Reset()         { *m = MsgLiquidateVaultResponse{} }
func (m *MsgLiquidateVaultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidateVaultResponse) ProtoMessage()    {}
func (*MsgLiquidateVaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d2beddde56424c, []int{1}
}
func (m *MsgLiquidateVaultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidateVaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidateVaultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidateVaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidateVaultResponse.Merge(m, src)
}
func (m *MsgLiquidateVaultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidateVaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidateVaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidateVaultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLiquidateVaultRequest)(nil), "petri.liquidation.v1beta1.MsgLiquidateVaultRequest")
	proto.RegisterType((*MsgLiquidateVaultResponse)(nil), "petri.liquidation.v1beta1.MsgLiquidateVaultResponse")
}

func init() {
	proto.RegisterFile("petri/liquidation/v1beta1/tx.proto", fileDescriptor_46d2beddde56424c)
}

var fileDescriptor_46d2beddde56424c = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0xcf, 0x4d,
	0x49, 0xad, 0xd0, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f,
	0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x82, 0x28, 0xd2, 0x43, 0x52, 0xa4, 0x07, 0x55, 0x24, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x56, 0xa6, 0x0f, 0x62, 0x41, 0x74, 0x28, 0xad, 0x65, 0xe4, 0x92, 0xf0, 0x2d, 0x4e, 0xf7, 0x81,
	0x6a, 0x48, 0x0d, 0x4b, 0x2c, 0xcd, 0x29, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52,
	0xe6, 0x62, 0x49, 0x2b, 0xca, 0xcf, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xe2, 0xff, 0x74,
	0x4f, 0x9e, 0xbb, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x24, 0xaa, 0x14, 0x04, 0x96, 0x14, 0x32,
	0xe4, 0x62, 0x4b, 0x2c, 0x28, 0x88, 0xcf, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x71, 0x92,
	0x7a, 0x74, 0x4f, 0x9e, 0xd5, 0xb1, 0xa0, 0xc0, 0x33, 0xe5, 0xd3, 0x3d, 0x79, 0x5e, 0x88, 0x7a,
	0x88, 0x02, 0xa5, 0x20, 0xd6, 0x44, 0x90, 0xb8, 0x90, 0x25, 0x17, 0x47, 0x19, 0xc8, 0x1e, 0x90,
	0x26, 0x66, 0xb0, 0x26, 0xb9, 0x47, 0xf7, 0xe4, 0xd9, 0xc1, 0x76, 0x83, 0xb5, 0xf1, 0x43, 0xb4,
	0xc1, 0x14, 0x29, 0x05, 0xb1, 0x97, 0x41, 0xe4, 0x94, 0xa4, 0xb9, 0x24, 0xb1, 0x38, 0xb7, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xa8, 0x83, 0x91, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0xa8, 0x81, 0x91,
	0x4b, 0x10, 0x43, 0x95, 0x90, 0x89, 0x1e, 0xee, 0xd0, 0xd1, 0xc3, 0x15, 0x06, 0x52, 0xa6, 0x24,
	0xea, 0x82, 0x38, 0xc5, 0x29, 0xfc, 0xc4, 0x43, 0x39, 0x86, 0x15, 0x8f, 0xe4, 0x18, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x09,
	0x64, 0xbc, 0x3e, 0xc4, 0x0a, 0xdd, 0xfc, 0xb4, 0xb4, 0xcc, 0xe4, 0xcc, 0xc4, 0x1c, 0x28, 0x5f,
	0x1f, 0x35, 0xb6, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xf1, 0x66, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x43, 0x73, 0x26, 0x9a, 0x10, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	MsgLiquidateVault(ctx context.Context, in *MsgLiquidateVaultRequest, opts ...grpc.CallOption) (*MsgLiquidateVaultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MsgLiquidateVault(ctx context.Context, in *MsgLiquidateVaultRequest, opts ...grpc.CallOption) (*MsgLiquidateVaultResponse, error) {
	out := new(MsgLiquidateVaultResponse)
	err := c.cc.Invoke(ctx, "/petri.liquidation.v1beta1.Msg/MsgLiquidateVault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MsgLiquidateVault(context.Context, *MsgLiquidateVaultRequest) (*MsgLiquidateVaultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MsgLiquidateVault(ctx context.Context, req *MsgLiquidateVaultRequest) (*MsgLiquidateVaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgLiquidateVault not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MsgLiquidateVault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidateVaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgLiquidateVault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petri.liquidation.v1beta1.Msg/MsgLiquidateVault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgLiquidateVault(ctx, req.(*MsgLiquidateVaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "petri.liquidation.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgLiquidateVault",
			Handler:    _Msg_MsgLiquidateVault_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petri/liquidation/v1beta1/tx.proto",
}

func (m *MsgLiquidateVaultRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidateVaultRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidateVaultRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VaultId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VaultId))
		i--
		dAtA[i] = 0x18
	}
	if m.AppId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLiquidateVaultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidateVaultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidateVaultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLiquidateVaultRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AppId != 0 {
		n += 1 + sovTx(uint64(m.AppId))
	}
	if m.VaultId != 0 {
		n += 1 + sovTx(uint64(m.VaultId))
	}
	return n
}

func (m *MsgLiquidateVaultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLiquidateVaultRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLiquidateVaultRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidateVaultRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultId", wireType)
			}
			m.VaultId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VaultId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLiquidateVaultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLiquidateVaultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidateVaultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
