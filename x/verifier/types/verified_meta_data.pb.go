// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifier/verifier/verified_meta_data.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type VerifiedMetaData struct {
	CodeId        uint64 `protobuf:"varint,1,opt,name=codeId,proto3" json:"codeId,omitempty"`
	ApplicationId uint64 `protobuf:"varint,2,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
}

func (m *VerifiedMetaData) Reset()         { *m = VerifiedMetaData{} }
func (m *VerifiedMetaData) String() string { return proto.CompactTextString(m) }
func (*VerifiedMetaData) ProtoMessage()    {}
func (*VerifiedMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3fe56b3ecf36778, []int{0}
}
func (m *VerifiedMetaData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifiedMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifiedMetaData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifiedMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifiedMetaData.Merge(m, src)
}
func (m *VerifiedMetaData) XXX_Size() int {
	return m.Size()
}
func (m *VerifiedMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifiedMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_VerifiedMetaData proto.InternalMessageInfo

func (m *VerifiedMetaData) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *VerifiedMetaData) GetApplicationId() uint64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func init() {
	proto.RegisterType((*VerifiedMetaData)(nil), "verifier.verifier.VerifiedMetaData")
}

func init() {
	proto.RegisterFile("verifier/verifier/verified_meta_data.proto", fileDescriptor_e3fe56b3ecf36778)
}

var fileDescriptor_e3fe56b3ecf36778 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x4b, 0x2d, 0xca,
	0x4c, 0xcb, 0x4c, 0x2d, 0xd2, 0x47, 0x67, 0xa4, 0xc4, 0xe7, 0xa6, 0x96, 0x24, 0xc6, 0xa7, 0x24,
	0x96, 0x24, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0xc2, 0x94, 0xe8, 0xc1, 0x18, 0x4a,
	0x01, 0x5c, 0x02, 0x61, 0x50, 0xe5, 0xbe, 0xa9, 0x25, 0x89, 0x2e, 0x89, 0x25, 0x89, 0x42, 0x62,
	0x5c, 0x6c, 0xc9, 0xf9, 0x29, 0xa9, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50,
	0x9e, 0x90, 0x0a, 0x17, 0x6f, 0x62, 0x41, 0x41, 0x4e, 0x66, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x9e,
	0x67, 0x8a, 0x04, 0x13, 0x58, 0x1a, 0x55, 0xd0, 0xc9, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0x24, 0xe1, 0x2e, 0xac, 0x40, 0x38, 0xb6, 0xa4, 0xb2, 0x20, 0xb5, 0x38,
	0x89, 0x0d, 0xec, 0x40, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x52, 0xa6, 0x3b, 0xce,
	0x00, 0x00, 0x00,
}

func (m *VerifiedMetaData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifiedMetaData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifiedMetaData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ApplicationId != 0 {
		i = encodeVarintVerifiedMetaData(dAtA, i, uint64(m.ApplicationId))
		i--
		dAtA[i] = 0x10
	}
	if m.CodeId != 0 {
		i = encodeVarintVerifiedMetaData(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVerifiedMetaData(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerifiedMetaData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VerifiedMetaData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovVerifiedMetaData(uint64(m.CodeId))
	}
	if m.ApplicationId != 0 {
		n += 1 + sovVerifiedMetaData(uint64(m.ApplicationId))
	}
	return n
}

func sovVerifiedMetaData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerifiedMetaData(x uint64) (n int) {
	return sovVerifiedMetaData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifiedMetaData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifiedMetaData
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
			return fmt.Errorf("proto: VerifiedMetaData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifiedMetaData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiedMetaData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationId", wireType)
			}
			m.ApplicationId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiedMetaData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApplicationId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVerifiedMetaData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerifiedMetaData
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
func skipVerifiedMetaData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifiedMetaData
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
					return 0, ErrIntOverflowVerifiedMetaData
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
					return 0, ErrIntOverflowVerifiedMetaData
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
				return 0, ErrInvalidLengthVerifiedMetaData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerifiedMetaData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerifiedMetaData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerifiedMetaData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifiedMetaData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerifiedMetaData = fmt.Errorf("proto: unexpected end of group")
)
