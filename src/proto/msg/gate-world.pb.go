// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gate-world.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetMapIDReq struct {
	MapUUID              string   `protobuf:"bytes,1,opt,name=MapUUID,proto3" json:"MapUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMapIDReq) Reset()         { *m = GetMapIDReq{} }
func (m *GetMapIDReq) String() string { return proto.CompactTextString(m) }
func (*GetMapIDReq) ProtoMessage()    {}
func (*GetMapIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_697686bfa75b3a2b, []int{0}
}
func (m *GetMapIDReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMapIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMapIDReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMapIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMapIDReq.Merge(m, src)
}
func (m *GetMapIDReq) XXX_Size() int {
	return m.Size()
}
func (m *GetMapIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMapIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetMapIDReq proto.InternalMessageInfo

func (m *GetMapIDReq) GetMapUUID() string {
	if m != nil {
		return m.MapUUID
	}
	return ""
}

type GetMapIDResp struct {
	MapID                int32    `protobuf:"varint,1,opt,name=MapID,proto3" json:"MapID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMapIDResp) Reset()         { *m = GetMapIDResp{} }
func (m *GetMapIDResp) String() string { return proto.CompactTextString(m) }
func (*GetMapIDResp) ProtoMessage()    {}
func (*GetMapIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_697686bfa75b3a2b, []int{1}
}
func (m *GetMapIDResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMapIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMapIDResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMapIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMapIDResp.Merge(m, src)
}
func (m *GetMapIDResp) XXX_Size() int {
	return m.Size()
}
func (m *GetMapIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMapIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetMapIDResp proto.InternalMessageInfo

func (m *GetMapIDResp) GetMapID() int32 {
	if m != nil {
		return m.MapID
	}
	return 0
}

func init() {
	proto.RegisterType((*GetMapIDReq)(nil), "GetMapIDReq")
	proto.RegisterType((*GetMapIDResp)(nil), "GetMapIDResp")
}

func init() { proto.RegisterFile("gate-world.proto", fileDescriptor_697686bfa75b3a2b) }

var fileDescriptor_697686bfa75b3a2b = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0x2c, 0x49,
	0xd5, 0x2d, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe7, 0xe2,
	0x76, 0x4f, 0x2d, 0xf1, 0x4d, 0x2c, 0xf0, 0x74, 0x09, 0x4a, 0x2d, 0x14, 0x92, 0xe0, 0x62, 0xf7,
	0x4d, 0x2c, 0x08, 0x0d, 0xf5, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95,
	0x54, 0xb8, 0x78, 0x10, 0x0a, 0x8b, 0x0b, 0x84, 0x44, 0xb8, 0x58, 0xc1, 0x1c, 0xb0, 0x3a, 0xd6,
	0x20, 0x08, 0xc7, 0x49, 0xe3, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0xcc, 0xd3, 0x2f, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48,
	0xbf, 0xb8, 0x28, 0x59, 0x1f, 0x6c, 0xa9, 0x7e, 0x6e, 0x71, 0x7a, 0x12, 0x1b, 0x98, 0x69, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x84, 0x10, 0xae, 0x93, 0x00, 0x00, 0x00,
}

func (m *GetMapIDReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMapIDReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMapIDReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MapUUID) > 0 {
		i -= len(m.MapUUID)
		copy(dAtA[i:], m.MapUUID)
		i = encodeVarintGateWorld(dAtA, i, uint64(len(m.MapUUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetMapIDResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMapIDResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMapIDResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MapID != 0 {
		i = encodeVarintGateWorld(dAtA, i, uint64(m.MapID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGateWorld(dAtA []byte, offset int, v uint64) int {
	offset -= sovGateWorld(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetMapIDReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MapUUID)
	if l > 0 {
		n += 1 + l + sovGateWorld(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetMapIDResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MapID != 0 {
		n += 1 + sovGateWorld(uint64(m.MapID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGateWorld(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGateWorld(x uint64) (n int) {
	return sovGateWorld(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetMapIDReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateWorld
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
			return fmt.Errorf("proto: GetMapIDReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMapIDReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateWorld
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
				return ErrInvalidLengthGateWorld
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGateWorld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MapUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGateWorld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateWorld
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGateWorld
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
func (m *GetMapIDResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGateWorld
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
			return fmt.Errorf("proto: GetMapIDResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMapIDResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapID", wireType)
			}
			m.MapID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGateWorld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MapID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGateWorld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGateWorld
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGateWorld
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
func skipGateWorld(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGateWorld
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
					return 0, ErrIntOverflowGateWorld
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
					return 0, ErrIntOverflowGateWorld
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
				return 0, ErrInvalidLengthGateWorld
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGateWorld
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGateWorld
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGateWorld        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGateWorld          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGateWorld = fmt.Errorf("proto: unexpected end of group")
)
