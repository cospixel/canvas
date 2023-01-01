// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canvas/canvas/point.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Point struct {
	X     uint64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y     uint64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ad7cf72f044f755, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Point.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return m.Size()
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetX() uint64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() uint64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Point) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func init() {
	proto.RegisterType((*Point)(nil), "canvas.canvas.Point")
}

func init() { proto.RegisterFile("canvas/canvas/point.proto", fileDescriptor_6ad7cf72f044f755) }

var fileDescriptor_6ad7cf72f044f755 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4e, 0xcc, 0x2b,
	0x4b, 0x2c, 0xd6, 0x87, 0x52, 0x05, 0xf9, 0x99, 0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xbc, 0x10, 0x31, 0x3d, 0x08, 0xa5, 0x64, 0xc9, 0xc5, 0x1a, 0x00, 0x92, 0x15, 0xe2, 0xe1,
	0x62, 0xac, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xac, 0x00, 0xf1, 0x2a, 0x25, 0x98,
	0x20, 0xbc, 0x4a, 0x21, 0x11, 0x2e, 0xd6, 0xe4, 0xfc, 0x9c, 0xfc, 0x22, 0x09, 0x66, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x08, 0xc7, 0x49, 0xff, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x44, 0xa1, 0xf6, 0x56, 0xc0, 0x1c, 0x50, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76,
	0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xb0, 0xaf, 0x2b, 0x9e, 0x00, 0x00, 0x00,
}

func (m *Point) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Point) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Point) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Color) > 0 {
		i -= len(m.Color)
		copy(dAtA[i:], m.Color)
		i = encodeVarintPoint(dAtA, i, uint64(len(m.Color)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Y != 0 {
		i = encodeVarintPoint(dAtA, i, uint64(m.Y))
		i--
		dAtA[i] = 0x10
	}
	if m.X != 0 {
		i = encodeVarintPoint(dAtA, i, uint64(m.X))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Point) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 1 + sovPoint(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovPoint(uint64(m.Y))
	}
	l = len(m.Color)
	if l > 0 {
		n += 1 + l + sovPoint(uint64(l))
	}
	return n
}

func sovPoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoint(x uint64) (n int) {
	return sovPoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Point) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoint
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
			return fmt.Errorf("proto: Point: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Point: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Color", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoint
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
				return ErrInvalidLengthPoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Color = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoint
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
func skipPoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoint
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
					return 0, ErrIntOverflowPoint
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
					return 0, ErrIntOverflowPoint
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
				return 0, ErrInvalidLengthPoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoint = fmt.Errorf("proto: unexpected end of group")
)