// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/query/generated/proto/admin/placement.proto

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

package admin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import placementpb "github.com/m3db/m3/src/cluster/generated/proto/placementpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlacementInitRequest struct {
	Instances         []*placementpb.Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	NumShards         int32                   `protobuf:"varint,2,opt,name=num_shards,json=numShards,proto3" json:"num_shards,omitempty"`
	ReplicationFactor int32                   `protobuf:"varint,3,opt,name=replication_factor,json=replicationFactor,proto3" json:"replication_factor,omitempty"`
}

func (m *PlacementInitRequest) Reset()                    { *m = PlacementInitRequest{} }
func (m *PlacementInitRequest) String() string            { return proto.CompactTextString(m) }
func (*PlacementInitRequest) ProtoMessage()               {}
func (*PlacementInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlacement, []int{0} }

func (m *PlacementInitRequest) GetInstances() []*placementpb.Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *PlacementInitRequest) GetNumShards() int32 {
	if m != nil {
		return m.NumShards
	}
	return 0
}

func (m *PlacementInitRequest) GetReplicationFactor() int32 {
	if m != nil {
		return m.ReplicationFactor
	}
	return 0
}

type PlacementGetResponse struct {
	Placement *placementpb.Placement `protobuf:"bytes,1,opt,name=placement" json:"placement,omitempty"`
	Version   int32                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *PlacementGetResponse) Reset()                    { *m = PlacementGetResponse{} }
func (m *PlacementGetResponse) String() string            { return proto.CompactTextString(m) }
func (*PlacementGetResponse) ProtoMessage()               {}
func (*PlacementGetResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlacement, []int{1} }

func (m *PlacementGetResponse) GetPlacement() *placementpb.Placement {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *PlacementGetResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type PlacementAddRequest struct {
	Instances []*placementpb.Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	// By default add requests will only succeed if all instances in the placement
	// are AVAILABLE for all their shards. force overrides that.
	Force bool `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (m *PlacementAddRequest) Reset()                    { *m = PlacementAddRequest{} }
func (m *PlacementAddRequest) String() string            { return proto.CompactTextString(m) }
func (*PlacementAddRequest) ProtoMessage()               {}
func (*PlacementAddRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlacement, []int{2} }

func (m *PlacementAddRequest) GetInstances() []*placementpb.Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *PlacementAddRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type PlacementReplaceRequest struct {
	LeavingInstanceIDs []string                `protobuf:"bytes,1,rep,name=leavingInstanceIDs" json:"leavingInstanceIDs,omitempty"`
	Candidates         []*placementpb.Instance `protobuf:"bytes,2,rep,name=candidates" json:"candidates,omitempty"`
	Force              bool                    `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (m *PlacementReplaceRequest) Reset()                    { *m = PlacementReplaceRequest{} }
func (m *PlacementReplaceRequest) String() string            { return proto.CompactTextString(m) }
func (*PlacementReplaceRequest) ProtoMessage()               {}
func (*PlacementReplaceRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlacement, []int{3} }

func (m *PlacementReplaceRequest) GetLeavingInstanceIDs() []string {
	if m != nil {
		return m.LeavingInstanceIDs
	}
	return nil
}

func (m *PlacementReplaceRequest) GetCandidates() []*placementpb.Instance {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func (m *PlacementReplaceRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

type PlacementSetRequest struct {
	Placement *placementpb.Placement `protobuf:"bytes,1,opt,name=placement" json:"placement,omitempty"`
	Version   int32                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Confirm must be set, otherwise just a dry run is executed.
	Confirm bool `protobuf:"varint,3,opt,name=confirm,proto3" json:"confirm,omitempty"`
}

func (m *PlacementSetRequest) Reset()                    { *m = PlacementSetRequest{} }
func (m *PlacementSetRequest) String() string            { return proto.CompactTextString(m) }
func (*PlacementSetRequest) ProtoMessage()               {}
func (*PlacementSetRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlacement, []int{4} }

func (m *PlacementSetRequest) GetPlacement() *placementpb.Placement {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *PlacementSetRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PlacementSetRequest) GetConfirm() bool {
	if m != nil {
		return m.Confirm
	}
	return false
}

type PlacementSetResponse struct {
	Placement *placementpb.Placement `protobuf:"bytes,1,opt,name=placement" json:"placement,omitempty"`
	Version   int32                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	DryRun    bool                   `protobuf:"varint,3,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
}

func (m *PlacementSetResponse) Reset()                    { *m = PlacementSetResponse{} }
func (m *PlacementSetResponse) String() string            { return proto.CompactTextString(m) }
func (*PlacementSetResponse) ProtoMessage()               {}
func (*PlacementSetResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlacement, []int{5} }

func (m *PlacementSetResponse) GetPlacement() *placementpb.Placement {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *PlacementSetResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PlacementSetResponse) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func init() {
	proto.RegisterType((*PlacementInitRequest)(nil), "admin.PlacementInitRequest")
	proto.RegisterType((*PlacementGetResponse)(nil), "admin.PlacementGetResponse")
	proto.RegisterType((*PlacementAddRequest)(nil), "admin.PlacementAddRequest")
	proto.RegisterType((*PlacementReplaceRequest)(nil), "admin.PlacementReplaceRequest")
	proto.RegisterType((*PlacementSetRequest)(nil), "admin.PlacementSetRequest")
	proto.RegisterType((*PlacementSetResponse)(nil), "admin.PlacementSetResponse")
}
func (m *PlacementInitRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementInitRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, msg := range m.Instances {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPlacement(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.NumShards != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.NumShards))
	}
	if m.ReplicationFactor != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.ReplicationFactor))
	}
	return i, nil
}

func (m *PlacementGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Placement != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.Placement.Size()))
		n1, err := m.Placement.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func (m *PlacementAddRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementAddRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, msg := range m.Instances {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPlacement(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Force {
		dAtA[i] = 0x10
		i++
		if m.Force {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *PlacementReplaceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementReplaceRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LeavingInstanceIDs) > 0 {
		for _, s := range m.LeavingInstanceIDs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Candidates) > 0 {
		for _, msg := range m.Candidates {
			dAtA[i] = 0x12
			i++
			i = encodeVarintPlacement(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Force {
		dAtA[i] = 0x18
		i++
		if m.Force {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *PlacementSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementSetRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Placement != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.Placement.Size()))
		n2, err := m.Placement.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.Version))
	}
	if m.Confirm {
		dAtA[i] = 0x18
		i++
		if m.Confirm {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *PlacementSetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlacementSetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Placement != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.Placement.Size()))
		n3, err := m.Placement.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPlacement(dAtA, i, uint64(m.Version))
	}
	if m.DryRun {
		dAtA[i] = 0x18
		i++
		if m.DryRun {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintPlacement(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PlacementInitRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, e := range m.Instances {
			l = e.Size()
			n += 1 + l + sovPlacement(uint64(l))
		}
	}
	if m.NumShards != 0 {
		n += 1 + sovPlacement(uint64(m.NumShards))
	}
	if m.ReplicationFactor != 0 {
		n += 1 + sovPlacement(uint64(m.ReplicationFactor))
	}
	return n
}

func (m *PlacementGetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Placement != nil {
		l = m.Placement.Size()
		n += 1 + l + sovPlacement(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovPlacement(uint64(m.Version))
	}
	return n
}

func (m *PlacementAddRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Instances) > 0 {
		for _, e := range m.Instances {
			l = e.Size()
			n += 1 + l + sovPlacement(uint64(l))
		}
	}
	if m.Force {
		n += 2
	}
	return n
}

func (m *PlacementReplaceRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.LeavingInstanceIDs) > 0 {
		for _, s := range m.LeavingInstanceIDs {
			l = len(s)
			n += 1 + l + sovPlacement(uint64(l))
		}
	}
	if len(m.Candidates) > 0 {
		for _, e := range m.Candidates {
			l = e.Size()
			n += 1 + l + sovPlacement(uint64(l))
		}
	}
	if m.Force {
		n += 2
	}
	return n
}

func (m *PlacementSetRequest) Size() (n int) {
	var l int
	_ = l
	if m.Placement != nil {
		l = m.Placement.Size()
		n += 1 + l + sovPlacement(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovPlacement(uint64(m.Version))
	}
	if m.Confirm {
		n += 2
	}
	return n
}

func (m *PlacementSetResponse) Size() (n int) {
	var l int
	_ = l
	if m.Placement != nil {
		l = m.Placement.Size()
		n += 1 + l + sovPlacement(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovPlacement(uint64(m.Version))
	}
	if m.DryRun {
		n += 2
	}
	return n
}

func sovPlacement(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlacement(x uint64) (n int) {
	return sovPlacement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlacementInitRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlacement
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
			return fmt.Errorf("proto: PlacementInitRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementInitRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
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
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instances = append(m.Instances, &placementpb.Instance{})
			if err := m.Instances[len(m.Instances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumShards", wireType)
			}
			m.NumShards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumShards |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			m.ReplicationFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicationFactor |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlacement
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
func (m *PlacementGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlacement
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
			return fmt.Errorf("proto: PlacementGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Placement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
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
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Placement == nil {
				m.Placement = &placementpb.Placement{}
			}
			if err := m.Placement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlacement
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
func (m *PlacementAddRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlacement
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
			return fmt.Errorf("proto: PlacementAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
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
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instances = append(m.Instances, &placementpb.Instance{})
			if err := m.Instances[len(m.Instances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Force", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Force = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlacement
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
func (m *PlacementReplaceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlacement
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
			return fmt.Errorf("proto: PlacementReplaceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementReplaceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeavingInstanceIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeavingInstanceIDs = append(m.LeavingInstanceIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Candidates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
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
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Candidates = append(m.Candidates, &placementpb.Instance{})
			if err := m.Candidates[len(m.Candidates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Force", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Force = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlacement
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
func (m *PlacementSetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlacement
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
			return fmt.Errorf("proto: PlacementSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Placement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
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
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Placement == nil {
				m.Placement = &placementpb.Placement{}
			}
			if err := m.Placement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirm", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Confirm = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlacement
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
func (m *PlacementSetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlacement
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
			return fmt.Errorf("proto: PlacementSetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlacementSetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Placement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
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
				return ErrInvalidLengthPlacement
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Placement == nil {
				m.Placement = &placementpb.Placement{}
			}
			if err := m.Placement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DryRun", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DryRun = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlacement
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
func skipPlacement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlacement
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
					return 0, ErrIntOverflowPlacement
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
					return 0, ErrIntOverflowPlacement
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
				return 0, ErrInvalidLengthPlacement
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlacement
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
				next, err := skipPlacement(dAtA[start:])
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
	ErrInvalidLengthPlacement = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlacement   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/query/generated/proto/admin/placement.proto", fileDescriptorPlacement)
}

var fileDescriptorPlacement = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x6e, 0xd4, 0x30,
	0x14, 0xc6, 0x31, 0xa3, 0x29, 0x8c, 0xbb, 0x01, 0x53, 0x4a, 0x84, 0x44, 0x34, 0xca, 0x6a, 0x36,
	0xc4, 0x52, 0x03, 0x07, 0xa0, 0x42, 0xa0, 0x61, 0x85, 0x3c, 0x07, 0x28, 0x8e, 0xfd, 0x32, 0xb5,
	0x14, 0xdb, 0xa9, 0xed, 0x54, 0xea, 0x06, 0xae, 0xc0, 0x0a, 0x89, 0x1b, 0xb1, 0xe4, 0x08, 0x68,
	0xb8, 0x08, 0xaa, 0x27, 0xff, 0xf8, 0xbb, 0x01, 0x96, 0xef, 0x7d, 0xcf, 0xdf, 0xf7, 0xcb, 0x7b,
	0x0a, 0x3e, 0xdd, 0xaa, 0x70, 0xde, 0x96, 0xb9, 0xb0, 0x9a, 0xea, 0x42, 0x96, 0x54, 0x17, 0xd4,
	0x3b, 0x41, 0x2f, 0x5a, 0x70, 0x57, 0x74, 0x0b, 0x06, 0x1c, 0x0f, 0x20, 0x69, 0xe3, 0x6c, 0xb0,
	0x94, 0x4b, 0xad, 0x0c, 0x6d, 0x6a, 0x2e, 0x40, 0x83, 0x09, 0x79, 0xec, 0x92, 0x79, 0x6c, 0x3f,
	0x7c, 0xf5, 0x1b, 0x2b, 0x51, 0xb7, 0x3e, 0x80, 0xfb, 0xc9, 0x6c, 0xb0, 0x69, 0xca, 0x1f, 0x2d,
	0xb3, 0x8f, 0x08, 0x1f, 0xbd, 0xee, 0x7b, 0x6b, 0xa3, 0x02, 0x83, 0x8b, 0x16, 0x7c, 0x20, 0x05,
	0x5e, 0x28, 0xe3, 0x03, 0x37, 0x02, 0x7c, 0x82, 0x96, 0xb3, 0xd5, 0xe1, 0xc9, 0xfd, 0x7c, 0xe2,
	0x94, 0xaf, 0x3b, 0x95, 0x8d, 0x73, 0xe4, 0x11, 0xc6, 0xa6, 0xd5, 0x67, 0xfe, 0x9c, 0x3b, 0xe9,
	0x93, 0x9b, 0x4b, 0xb4, 0x9a, 0xb3, 0x85, 0x69, 0xf5, 0x26, 0x36, 0xc8, 0x63, 0x4c, 0x1c, 0x34,
	0xb5, 0x12, 0x3c, 0x28, 0x6b, 0xce, 0x2a, 0x2e, 0x82, 0x75, 0xc9, 0x2c, 0x8e, 0xdd, 0x9d, 0x28,
	0x2f, 0xa2, 0x90, 0x55, 0x13, 0xb4, 0x97, 0x10, 0x18, 0xf8, 0xc6, 0x1a, 0x0f, 0xe4, 0x09, 0x5e,
	0x0c, 0x20, 0x09, 0x5a, 0xa2, 0xd5, 0xe1, 0xc9, 0xf1, 0x77, 0x68, 0xc3, 0x2b, 0x36, 0x0e, 0x92,
	0x04, 0xdf, 0xba, 0x04, 0xe7, 0x95, 0x35, 0x1d, 0x58, 0x5f, 0x66, 0x6f, 0xf0, 0xbd, 0xe1, 0xc5,
	0x33, 0x29, 0xff, 0x6a, 0x03, 0x47, 0x78, 0x5e, 0x59, 0x27, 0x20, 0x66, 0xdc, 0x66, 0xfb, 0x22,
	0xfb, 0x80, 0xf0, 0x83, 0x11, 0x0a, 0xa2, 0x49, 0x1f, 0x93, 0x63, 0x52, 0x03, 0xbf, 0x54, 0x66,
	0xdb, 0xfb, 0xad, 0x9f, 0xef, 0xf3, 0x16, 0xec, 0x17, 0x0a, 0x79, 0x8a, 0xb1, 0xe0, 0x46, 0x2a,
	0xc9, 0x03, 0x5c, 0xef, 0xf8, 0x0f, 0x5c, 0x93, 0xc1, 0x11, 0x6c, 0x36, 0x05, 0x7b, 0x37, 0xf9,
	0xf4, 0x0d, 0x0c, 0xc7, 0xff, 0xc7, 0x1b, 0xbe, 0x56, 0x84, 0x35, 0x95, 0x72, 0xba, 0x8b, 0xef,
	0xcb, 0xec, 0xed, 0xe4, 0xc6, 0x9b, 0xff, 0x77, 0x63, 0x72, 0x8c, 0x0f, 0xa4, 0xbb, 0x62, 0xad,
	0xe9, 0x00, 0xba, 0xea, 0xf4, 0xce, 0xa7, 0x5d, 0x8a, 0x3e, 0xef, 0x52, 0xf4, 0x65, 0x97, 0xa2,
	0xf7, 0x5f, 0xd3, 0x1b, 0xe5, 0x41, 0xfc, 0x31, 0x8a, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x6f, 0x96, 0x58, 0xb1, 0x03, 0x00, 0x00,
}