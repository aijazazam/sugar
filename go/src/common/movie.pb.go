// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: movie.proto

package common

import (
	encoding_binary "encoding/binary"
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

type Movie struct {
	MovieId     uint32  `protobuf:"varint,1,opt,name=MovieId,json=movieId,proto3" json:"MovieId,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Genere      string  `protobuf:"bytes,3,opt,name=Genere,json=genere,proto3" json:"Genere,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	Director    string  `protobuf:"bytes,5,opt,name=Director,json=director,proto3" json:"Director,omitempty"`
	Actors      string  `protobuf:"bytes,6,opt,name=Actors,json=actors,proto3" json:"Actors,omitempty"`
	Year        uint32  `protobuf:"varint,7,opt,name=Year,json=year,proto3" json:"Year,omitempty"`
	Duration    uint32  `protobuf:"varint,8,opt,name=Duration,json=duration,proto3" json:"Duration,omitempty"`
	Rating      float32 `protobuf:"fixed32,9,opt,name=Rating,json=rating,proto3" json:"Rating,omitempty"`
	Votes       uint32  `protobuf:"varint,10,opt,name=Votes,json=votes,proto3" json:"Votes,omitempty"`
	Revenue     float32 `protobuf:"fixed32,11,opt,name=Revenue,json=revenue,proto3" json:"Revenue,omitempty"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{0}
}
func (m *Movie) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return m.Size()
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetMovieId() uint32 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *Movie) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Movie) GetGenere() string {
	if m != nil {
		return m.Genere
	}
	return ""
}

func (m *Movie) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Movie) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

func (m *Movie) GetActors() string {
	if m != nil {
		return m.Actors
	}
	return ""
}

func (m *Movie) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Movie) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Movie) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Movie) GetVotes() uint32 {
	if m != nil {
		return m.Votes
	}
	return 0
}

func (m *Movie) GetRevenue() float32 {
	if m != nil {
		return m.Revenue
	}
	return 0
}

func init() {
	proto.RegisterType((*Movie)(nil), "common.Movie")
}

func init() { proto.RegisterFile("movie.proto", fileDescriptor_fde087a4194eda75) }

var fileDescriptor_fde087a4194eda75 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x9b, 0xda, 0xa6, 0xbd, 0x14, 0x97, 0x20, 0xf2, 0x70, 0x08, 0xc5, 0xa9, 0x93, 0x8b,
	0x9f, 0x40, 0x39, 0x10, 0x07, 0x1d, 0x3a, 0x08, 0x8e, 0xb1, 0x7d, 0x1c, 0x1d, 0x92, 0x1c, 0x69,
	0xae, 0xe0, 0xb7, 0x70, 0xf7, 0x0b, 0x39, 0xde, 0xe8, 0x28, 0xed, 0x17, 0x91, 0xbc, 0x1e, 0xdc,
	0xf6, 0x7e, 0xff, 0x97, 0xfc, 0xf2, 0xf2, 0x44, 0x65, 0xdc, 0x34, 0xe0, 0xdd, 0xde, 0xbb, 0xe0,
	0x24, 0xef, 0x9c, 0x31, 0xce, 0xde, 0x7e, 0xa7, 0x22, 0x7f, 0x89, 0xb9, 0x04, 0x51, 0x50, 0xf1,
	0xdc, 0x03, 0xab, 0x59, 0x73, 0xd9, 0x16, 0x66, 0x45, 0x29, 0x45, 0xf6, 0xaa, 0x0d, 0x42, 0x5a,
	0xb3, 0x66, 0xd3, 0x66, 0x56, 0x1b, 0x94, 0xd7, 0x82, 0x3f, 0xa1, 0x45, 0x8f, 0x70, 0x41, 0x29,
	0xdf, 0x11, 0xc9, 0x5a, 0x54, 0x5b, 0x1c, 0x3b, 0x3f, 0xec, 0xc3, 0xe0, 0x2c, 0x64, 0xd4, 0xac,
	0xfa, 0x73, 0x24, 0x6f, 0x44, 0xb9, 0x1d, 0x3c, 0x76, 0xc1, 0x79, 0xc8, 0xa9, 0x5d, 0xf6, 0x27,
	0x8e, 0xd6, 0x87, 0x58, 0x8c, 0xc0, 0x57, 0xab, 0x26, 0x8a, 0x13, 0xbc, 0xa3, 0xf6, 0x50, 0xd0,
	0x60, 0xd9, 0x27, 0x6a, 0x4f, 0x9e, 0x83, 0xd7, 0xf4, 0x4c, 0x49, 0x79, 0xd9, 0x9f, 0x38, 0x7a,
	0x5a, 0x1d, 0x06, 0xbb, 0x83, 0x4d, 0xcd, 0x9a, 0xb4, 0xe5, 0x9e, 0x48, 0x5e, 0x89, 0xfc, 0xcd,
	0x05, 0x1c, 0x41, 0xd0, 0x85, 0x7c, 0x8a, 0x10, 0x7f, 0xde, 0xe2, 0x84, 0xf6, 0x80, 0x50, 0xd1,
	0xf1, 0xc2, 0xaf, 0xf8, 0x08, 0x3f, 0xb3, 0x62, 0xc7, 0x59, 0xb1, 0xbf, 0x59, 0xb1, 0xaf, 0x45,
	0x25, 0xc7, 0x45, 0x25, 0xbf, 0x8b, 0x4a, 0x3e, 0x38, 0xad, 0xf1, 0xfe, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0x17, 0xde, 0x4f, 0xc8, 0x55, 0x01, 0x00, 0x00,
}

func (m *Movie) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Movie) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Movie) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Revenue != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Revenue))))
		i--
		dAtA[i] = 0x5d
	}
	if m.Votes != 0 {
		i = encodeVarintMovie(dAtA, i, uint64(m.Votes))
		i--
		dAtA[i] = 0x50
	}
	if m.Rating != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Rating))))
		i--
		dAtA[i] = 0x4d
	}
	if m.Duration != 0 {
		i = encodeVarintMovie(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x40
	}
	if m.Year != 0 {
		i = encodeVarintMovie(dAtA, i, uint64(m.Year))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Actors) > 0 {
		i -= len(m.Actors)
		copy(dAtA[i:], m.Actors)
		i = encodeVarintMovie(dAtA, i, uint64(len(m.Actors)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Director) > 0 {
		i -= len(m.Director)
		copy(dAtA[i:], m.Director)
		i = encodeVarintMovie(dAtA, i, uint64(len(m.Director)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMovie(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Genere) > 0 {
		i -= len(m.Genere)
		copy(dAtA[i:], m.Genere)
		i = encodeVarintMovie(dAtA, i, uint64(len(m.Genere)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMovie(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.MovieId != 0 {
		i = encodeVarintMovie(dAtA, i, uint64(m.MovieId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMovie(dAtA []byte, offset int, v uint64) int {
	offset -= sovMovie(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Movie) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MovieId != 0 {
		n += 1 + sovMovie(uint64(m.MovieId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMovie(uint64(l))
	}
	l = len(m.Genere)
	if l > 0 {
		n += 1 + l + sovMovie(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMovie(uint64(l))
	}
	l = len(m.Director)
	if l > 0 {
		n += 1 + l + sovMovie(uint64(l))
	}
	l = len(m.Actors)
	if l > 0 {
		n += 1 + l + sovMovie(uint64(l))
	}
	if m.Year != 0 {
		n += 1 + sovMovie(uint64(m.Year))
	}
	if m.Duration != 0 {
		n += 1 + sovMovie(uint64(m.Duration))
	}
	if m.Rating != 0 {
		n += 5
	}
	if m.Votes != 0 {
		n += 1 + sovMovie(uint64(m.Votes))
	}
	if m.Revenue != 0 {
		n += 5
	}
	return n
}

func sovMovie(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMovie(x uint64) (n int) {
	return sovMovie(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Movie) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMovie
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
			return fmt.Errorf("proto: Movie: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Movie: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MovieId", wireType)
			}
			m.MovieId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MovieId |= uint32(b&0x7F) << shift
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
					return ErrIntOverflowMovie
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
				return ErrInvalidLengthMovie
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMovie
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Genere", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
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
				return ErrInvalidLengthMovie
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMovie
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Genere = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
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
				return ErrInvalidLengthMovie
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMovie
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Director", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
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
				return ErrInvalidLengthMovie
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMovie
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Director = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
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
				return ErrInvalidLengthMovie
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMovie
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actors = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Year", wireType)
			}
			m.Year = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Year |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rating", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Rating = float32(math.Float32frombits(v))
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			m.Votes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMovie
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Votes |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revenue", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Revenue = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMovie(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMovie
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMovie
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
func skipMovie(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMovie
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
					return 0, ErrIntOverflowMovie
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
					return 0, ErrIntOverflowMovie
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
				return 0, ErrInvalidLengthMovie
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMovie
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMovie
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMovie        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMovie          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMovie = fmt.Errorf("proto: unexpected end of group")
)
