// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: database.proto

package etc

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

type Database struct {
	IP                   string   `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	ConnMaxLifetime      int64    `protobuf:"varint,4,opt,name=ConnMaxLifetime,proto3" json:"ConnMaxLifetime,omitempty"`
	MaxOpenConns         int32    `protobuf:"varint,5,opt,name=MaxOpenConns,proto3" json:"MaxOpenConns,omitempty"`
	MaxIdleConns         int32    `protobuf:"varint,6,opt,name=MaxIdleConns,proto3" json:"MaxIdleConns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Database) Reset()         { *m = Database{} }
func (m *Database) String() string { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()    {}
func (*Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_b90fe3356ea5df07, []int{0}
}
func (m *Database) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Database.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Database.Merge(m, src)
}
func (m *Database) XXX_Size() int {
	return m.Size()
}
func (m *Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Database proto.InternalMessageInfo

func (m *Database) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Database) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Database) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Database) GetConnMaxLifetime() int64 {
	if m != nil {
		return m.ConnMaxLifetime
	}
	return 0
}

func (m *Database) GetMaxOpenConns() int32 {
	if m != nil {
		return m.MaxOpenConns
	}
	return 0
}

func (m *Database) GetMaxIdleConns() int32 {
	if m != nil {
		return m.MaxIdleConns
	}
	return 0
}

func init() {
	proto.RegisterType((*Database)(nil), "Database")
}

func init() { proto.RegisterFile("database.proto", fileDescriptor_b90fe3356ea5df07) }

var fileDescriptor_b90fe3356ea5df07 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x2c, 0x49,
	0x4c, 0x4a, 0x2c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x3a, 0xc2, 0xc8, 0xc5, 0xe1,
	0x02, 0x15, 0x12, 0xe2, 0xe3, 0x62, 0xf2, 0x0c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0xf2, 0x0c, 0x10, 0x92, 0xe2, 0xe2, 0x08, 0x2d, 0x4e, 0x2d, 0xf2, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x02, 0x8b, 0xc2, 0xf9, 0x20, 0xb9, 0x80, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x66,
	0x88, 0x1c, 0x8c, 0x2f, 0xa4, 0xc1, 0xc5, 0xef, 0x9c, 0x9f, 0x97, 0xe7, 0x9b, 0x58, 0xe1, 0x93,
	0x99, 0x96, 0x5a, 0x92, 0x99, 0x9b, 0x2a, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x2e, 0x2c,
	0xa4, 0xc4, 0xc5, 0xe3, 0x9b, 0x58, 0xe1, 0x5f, 0x90, 0x9a, 0x07, 0x92, 0x29, 0x96, 0x60, 0x55,
	0x60, 0xd4, 0x60, 0x0d, 0x42, 0x11, 0x83, 0xaa, 0xf1, 0x4c, 0xc9, 0x49, 0x85, 0xa8, 0x61, 0x83,
	0xab, 0x81, 0x8b, 0x39, 0x69, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x89, 0x79, 0xfa, 0x05, 0xa7, 0x16, 0x95, 0xa5, 0x16,
	0xe9, 0x17, 0x17, 0x25, 0xeb, 0x83, 0x3d, 0xab, 0x9f, 0x5a, 0x92, 0x9c, 0xc4, 0x06, 0x66, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x39, 0x16, 0x0a, 0x21, 0x09, 0x01, 0x00, 0x00,
}

func (m *Database) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Database) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Database) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxIdleConns != 0 {
		i = encodeVarintDatabase(dAtA, i, uint64(m.MaxIdleConns))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxOpenConns != 0 {
		i = encodeVarintDatabase(dAtA, i, uint64(m.MaxOpenConns))
		i--
		dAtA[i] = 0x28
	}
	if m.ConnMaxLifetime != 0 {
		i = encodeVarintDatabase(dAtA, i, uint64(m.ConnMaxLifetime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintDatabase(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UserName) > 0 {
		i -= len(m.UserName)
		copy(dAtA[i:], m.UserName)
		i = encodeVarintDatabase(dAtA, i, uint64(len(m.UserName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IP) > 0 {
		i -= len(m.IP)
		copy(dAtA[i:], m.IP)
		i = encodeVarintDatabase(dAtA, i, uint64(len(m.IP)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDatabase(dAtA []byte, offset int, v uint64) int {
	offset -= sovDatabase(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Database) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovDatabase(uint64(l))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovDatabase(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovDatabase(uint64(l))
	}
	if m.ConnMaxLifetime != 0 {
		n += 1 + sovDatabase(uint64(m.ConnMaxLifetime))
	}
	if m.MaxOpenConns != 0 {
		n += 1 + sovDatabase(uint64(m.MaxOpenConns))
	}
	if m.MaxIdleConns != 0 {
		n += 1 + sovDatabase(uint64(m.MaxIdleConns))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDatabase(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDatabase(x uint64) (n int) {
	return sovDatabase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Database) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatabase
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
			return fmt.Errorf("proto: Database: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Database: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabase
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
				return ErrInvalidLengthDatabase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatabase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabase
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
				return ErrInvalidLengthDatabase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatabase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabase
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
				return ErrInvalidLengthDatabase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatabase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnMaxLifetime", wireType)
			}
			m.ConnMaxLifetime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnMaxLifetime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOpenConns", wireType)
			}
			m.MaxOpenConns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxOpenConns |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxIdleConns", wireType)
			}
			m.MaxIdleConns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatabase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxIdleConns |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDatabase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatabase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDatabase
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
func skipDatabase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDatabase
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
					return 0, ErrIntOverflowDatabase
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
					return 0, ErrIntOverflowDatabase
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
				return 0, ErrInvalidLengthDatabase
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDatabase
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDatabase
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDatabase        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDatabase          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDatabase = fmt.Errorf("proto: unexpected end of group")
)
