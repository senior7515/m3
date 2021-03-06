// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/metrics/generated/proto/policypb/policy.proto

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
	Package policypb is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/metrics/generated/proto/policypb/policy.proto

	It has these top-level messages:
		Resolution
		Retention
		StoragePolicy
		Policy
*/
package policypb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import aggregationpb "github.com/m3db/m3/src/metrics/generated/proto/aggregationpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DropPolicy int32

const (
	DropPolicy_NONE               DropPolicy = 0
	DropPolicy_DROP_MUST          DropPolicy = 1
	DropPolicy_DROP_IF_ONLY_MATCH DropPolicy = 2
)

var DropPolicy_name = map[int32]string{
	0: "NONE",
	1: "DROP_MUST",
	2: "DROP_IF_ONLY_MATCH",
}
var DropPolicy_value = map[string]int32{
	"NONE":               0,
	"DROP_MUST":          1,
	"DROP_IF_ONLY_MATCH": 2,
}

func (x DropPolicy) String() string {
	return proto.EnumName(DropPolicy_name, int32(x))
}
func (DropPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{0} }

type Resolution struct {
	WindowSize int64 `protobuf:"varint,1,opt,name=window_size,json=windowSize,proto3" json:"window_size,omitempty"`
	Precision  int64 `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (m *Resolution) Reset()                    { *m = Resolution{} }
func (m *Resolution) String() string            { return proto.CompactTextString(m) }
func (*Resolution) ProtoMessage()               {}
func (*Resolution) Descriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{0} }

func (m *Resolution) GetWindowSize() int64 {
	if m != nil {
		return m.WindowSize
	}
	return 0
}

func (m *Resolution) GetPrecision() int64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

type Retention struct {
	Period int64 `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
}

func (m *Retention) Reset()                    { *m = Retention{} }
func (m *Retention) String() string            { return proto.CompactTextString(m) }
func (*Retention) ProtoMessage()               {}
func (*Retention) Descriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{1} }

func (m *Retention) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type StoragePolicy struct {
	Resolution Resolution `protobuf:"bytes,1,opt,name=resolution" json:"resolution"`
	Retention  Retention  `protobuf:"bytes,2,opt,name=retention" json:"retention"`
}

func (m *StoragePolicy) Reset()                    { *m = StoragePolicy{} }
func (m *StoragePolicy) String() string            { return proto.CompactTextString(m) }
func (*StoragePolicy) ProtoMessage()               {}
func (*StoragePolicy) Descriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{2} }

func (m *StoragePolicy) GetResolution() Resolution {
	if m != nil {
		return m.Resolution
	}
	return Resolution{}
}

func (m *StoragePolicy) GetRetention() Retention {
	if m != nil {
		return m.Retention
	}
	return Retention{}
}

type Policy struct {
	StoragePolicy    *StoragePolicy                  `protobuf:"bytes,1,opt,name=storage_policy,json=storagePolicy" json:"storage_policy,omitempty"`
	AggregationTypes []aggregationpb.AggregationType `protobuf:"varint,2,rep,packed,name=aggregation_types,json=aggregationTypes,enum=aggregationpb.AggregationType" json:"aggregation_types,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptorPolicy, []int{3} }

func (m *Policy) GetStoragePolicy() *StoragePolicy {
	if m != nil {
		return m.StoragePolicy
	}
	return nil
}

func (m *Policy) GetAggregationTypes() []aggregationpb.AggregationType {
	if m != nil {
		return m.AggregationTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*Resolution)(nil), "policypb.Resolution")
	proto.RegisterType((*Retention)(nil), "policypb.Retention")
	proto.RegisterType((*StoragePolicy)(nil), "policypb.StoragePolicy")
	proto.RegisterType((*Policy)(nil), "policypb.Policy")
	proto.RegisterEnum("policypb.DropPolicy", DropPolicy_name, DropPolicy_value)
}
func (m *Resolution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resolution) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WindowSize != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPolicy(dAtA, i, uint64(m.WindowSize))
	}
	if m.Precision != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPolicy(dAtA, i, uint64(m.Precision))
	}
	return i, nil
}

func (m *Retention) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Retention) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Period != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPolicy(dAtA, i, uint64(m.Period))
	}
	return i, nil
}

func (m *StoragePolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoragePolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintPolicy(dAtA, i, uint64(m.Resolution.Size()))
	n1, err := m.Resolution.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintPolicy(dAtA, i, uint64(m.Retention.Size()))
	n2, err := m.Retention.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StoragePolicy != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPolicy(dAtA, i, uint64(m.StoragePolicy.Size()))
		n3, err := m.StoragePolicy.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.AggregationTypes) > 0 {
		dAtA5 := make([]byte, len(m.AggregationTypes)*10)
		var j4 int
		for _, num := range m.AggregationTypes {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintPolicy(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	return i, nil
}

func encodeVarintPolicy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Resolution) Size() (n int) {
	var l int
	_ = l
	if m.WindowSize != 0 {
		n += 1 + sovPolicy(uint64(m.WindowSize))
	}
	if m.Precision != 0 {
		n += 1 + sovPolicy(uint64(m.Precision))
	}
	return n
}

func (m *Retention) Size() (n int) {
	var l int
	_ = l
	if m.Period != 0 {
		n += 1 + sovPolicy(uint64(m.Period))
	}
	return n
}

func (m *StoragePolicy) Size() (n int) {
	var l int
	_ = l
	l = m.Resolution.Size()
	n += 1 + l + sovPolicy(uint64(l))
	l = m.Retention.Size()
	n += 1 + l + sovPolicy(uint64(l))
	return n
}

func (m *Policy) Size() (n int) {
	var l int
	_ = l
	if m.StoragePolicy != nil {
		l = m.StoragePolicy.Size()
		n += 1 + l + sovPolicy(uint64(l))
	}
	if len(m.AggregationTypes) > 0 {
		l = 0
		for _, e := range m.AggregationTypes {
			l += sovPolicy(uint64(e))
		}
		n += 1 + sovPolicy(uint64(l)) + l
	}
	return n
}

func sovPolicy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPolicy(x uint64) (n int) {
	return sovPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resolution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Resolution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resolution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowSize", wireType)
			}
			m.WindowSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Precision", wireType)
			}
			m.Precision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Precision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolicy
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
func (m *Retention) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Retention: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Retention: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolicy
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
func (m *StoragePolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoragePolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoragePolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resolution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resolution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retention", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Retention.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolicy
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
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StoragePolicy == nil {
				m.StoragePolicy = &StoragePolicy{}
			}
			if err := m.StoragePolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v aggregationpb.AggregationType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPolicy
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (aggregationpb.AggregationType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AggregationTypes = append(m.AggregationTypes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPolicy
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPolicy
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v aggregationpb.AggregationType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPolicy
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (aggregationpb.AggregationType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AggregationTypes = append(m.AggregationTypes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationTypes", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolicy
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
func skipPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPolicy
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPolicy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPolicy
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPolicy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPolicy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicy   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/metrics/generated/proto/policypb/policy.proto", fileDescriptorPolicy)
}

var fileDescriptorPolicy = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0xcd, 0xa4, 0x25, 0x34, 0x37, 0xa4, 0xac, 0xa3, 0xd4, 0x50, 0x64, 0x5b, 0xd6, 0x97, 0x22,
	0xb8, 0x0b, 0xc9, 0x43, 0x41, 0x50, 0x68, 0x8d, 0x62, 0xa9, 0xdd, 0x94, 0x4d, 0x7c, 0xd0, 0x97,
	0x65, 0x3f, 0xc6, 0x71, 0xa0, 0xbb, 0x33, 0xcc, 0x4c, 0x28, 0xe9, 0xb3, 0x3f, 0xc0, 0x17, 0xff,
	0x53, 0x1f, 0xfd, 0x05, 0x22, 0xf1, 0x8f, 0x48, 0x66, 0x27, 0xee, 0xee, 0xa3, 0xbe, 0xdd, 0x73,
	0xee, 0x3d, 0xe7, 0x1e, 0xee, 0x0c, 0x4c, 0x29, 0xd3, 0x5f, 0x96, 0xa9, 0x9f, 0xf1, 0x22, 0x28,
	0x26, 0x79, 0x1a, 0x14, 0x93, 0x40, 0xc9, 0x2c, 0x28, 0x88, 0x96, 0x2c, 0x53, 0x01, 0x25, 0x25,
	0x91, 0x89, 0x26, 0x79, 0x20, 0x24, 0xd7, 0x3c, 0x10, 0xfc, 0x86, 0x65, 0x2b, 0x91, 0xda, 0xc2,
	0x37, 0x2c, 0xde, 0xdb, 0xd2, 0x87, 0xcf, 0x1b, 0x7e, 0x94, 0x53, 0x5e, 0xc9, 0xd2, 0xe5, 0x67,
	0x83, 0x2a, 0x8f, 0x4d, 0x55, 0x09, 0x0f, 0xc3, 0x7f, 0x5c, 0x9f, 0x50, 0x2a, 0x09, 0x4d, 0x34,
	0xe3, 0xa5, 0x48, 0x9b, 0xa8, 0xf2, 0xf3, 0x2e, 0x01, 0x22, 0xa2, 0xf8, 0xcd, 0x72, 0xc3, 0xe1,
	0x23, 0x18, 0xdc, 0xb2, 0x32, 0xe7, 0xb7, 0xb1, 0x62, 0x77, 0x64, 0x84, 0x8e, 0xd1, 0xc9, 0x4e,
	0x04, 0x15, 0x35, 0x67, 0x77, 0x04, 0x3f, 0x81, 0xbe, 0x90, 0x24, 0x63, 0x8a, 0xf1, 0x72, 0xd4,
	0x35, 0xed, 0x9a, 0xf0, 0x9e, 0x42, 0x3f, 0x22, 0x9a, 0x94, 0xc6, 0xeb, 0x00, 0x7a, 0x82, 0x48,
	0xc6, 0x73, 0x6b, 0x63, 0x91, 0xf7, 0x15, 0xc1, 0x70, 0xae, 0xb9, 0x4c, 0x28, 0xb9, 0x36, 0x47,
	0xc0, 0x2f, 0x00, 0xe4, 0xdf, 0x0c, 0x66, 0x7a, 0x30, 0x7e, 0xe4, 0x6f, 0x2f, 0xe4, 0xd7, 0xf9,
	0xce, 0x77, 0xef, 0x7f, 0x1e, 0x75, 0xa2, 0xc6, 0x34, 0x3e, 0x85, 0xbe, 0xdc, 0xae, 0x34, 0x81,
	0x06, 0xe3, 0x87, 0x4d, 0xa9, 0x6d, 0x59, 0x65, 0x3d, 0xeb, 0x7d, 0x47, 0xd0, 0xb3, 0xfb, 0x5f,
	0xc1, 0xbe, 0xaa, 0x02, 0xc5, 0x95, 0xd2, 0x66, 0x78, 0x5c, 0x1b, 0xb5, 0x02, 0x47, 0x43, 0xd5,
	0xca, 0x7f, 0x09, 0x0f, 0x1a, 0x87, 0x8d, 0xf5, 0x4a, 0x10, 0x35, 0xea, 0x1e, 0xef, 0x9c, 0xec,
	0x8f, 0x5d, 0xbf, 0xf5, 0x00, 0xfe, 0x59, 0x8d, 0x16, 0x2b, 0x41, 0x22, 0x27, 0x69, 0x13, 0xea,
	0xd9, 0x4b, 0x80, 0xa9, 0xe4, 0xc2, 0x5a, 0xef, 0xc1, 0x6e, 0x38, 0x0b, 0xdf, 0x38, 0x1d, 0x3c,
	0x84, 0xfe, 0x34, 0x9a, 0x5d, 0xc7, 0x57, 0x1f, 0xe6, 0x0b, 0x07, 0xe1, 0x03, 0xc0, 0x06, 0x5e,
	0xbc, 0x8d, 0x67, 0xe1, 0xfb, 0x8f, 0xf1, 0xd5, 0xd9, 0xe2, 0xf5, 0x3b, 0xa7, 0x7b, 0x7e, 0x71,
	0xbf, 0x76, 0xd1, 0x8f, 0xb5, 0x8b, 0x7e, 0xad, 0x5d, 0xf4, 0xed, 0xb7, 0xdb, 0xf9, 0x74, 0xfa,
	0x9f, 0x1f, 0x36, 0xed, 0x19, 0x3c, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x16, 0x63, 0x9d, 0x5c,
	0xf2, 0x02, 0x00, 0x00,
}
