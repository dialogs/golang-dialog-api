// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: peers.proto

package dialog

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type PeerType int32

const (
	PEERTYPE_UNKNOWN          PeerType = 0
	PEERTYPE_PRIVATE          PeerType = 1
	PEERTYPE_GROUP            PeerType = 2
	PEERTYPE_ENCRYPTEDPRIVATE PeerType = 3
	PEERTYPE_SIP              PeerType = 4
)

var PeerType_name = map[int32]string{
	0: "PEERTYPE_UNKNOWN",
	1: "PEERTYPE_PRIVATE",
	2: "PEERTYPE_GROUP",
	3: "PEERTYPE_ENCRYPTEDPRIVATE",
	4: "PEERTYPE_SIP",
}

var PeerType_value = map[string]int32{
	"PEERTYPE_UNKNOWN":          0,
	"PEERTYPE_PRIVATE":          1,
	"PEERTYPE_GROUP":            2,
	"PEERTYPE_ENCRYPTEDPRIVATE": 3,
	"PEERTYPE_SIP":              4,
}

func (PeerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b029f9e0d26cab5, []int{0}
}

// Peer
type Peer struct {
	Type  PeerType           `protobuf:"varint,1,opt,name=type,proto3,enum=dialog.PeerType" json:"type,omitempty"`
	Id    int32              `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	StrId *types.StringValue `protobuf:"bytes,3,opt,name=str_id,json=strId,proto3" json:"str_id,omitempty"`
}

func (m *Peer) Reset()      { *m = Peer{} }
func (*Peer) ProtoMessage() {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b029f9e0d26cab5, []int{0}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetType() PeerType {
	if m != nil {
		return m.Type
	}
	return PEERTYPE_UNKNOWN
}

func (m *Peer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Peer) GetStrId() *types.StringValue {
	if m != nil {
		return m.StrId
	}
	return nil
}

// Out peer with access hash
type OutPeer struct {
	Type       PeerType           `protobuf:"varint,1,opt,name=type,proto3,enum=dialog.PeerType" json:"type,omitempty"`
	Id         int32              `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AccessHash int64              `protobuf:"varint,3,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty"`
	StrId      *types.StringValue `protobuf:"bytes,4,opt,name=str_id,json=strId,proto3" json:"str_id,omitempty"`
}

func (m *OutPeer) Reset()      { *m = OutPeer{} }
func (*OutPeer) ProtoMessage() {}
func (*OutPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b029f9e0d26cab5, []int{1}
}
func (m *OutPeer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutPeer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutPeer.Merge(m, src)
}
func (m *OutPeer) XXX_Size() int {
	return m.Size()
}
func (m *OutPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_OutPeer.DiscardUnknown(m)
}

var xxx_messageInfo_OutPeer proto.InternalMessageInfo

func (m *OutPeer) GetType() PeerType {
	if m != nil {
		return m.Type
	}
	return PEERTYPE_UNKNOWN
}

func (m *OutPeer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OutPeer) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

func (m *OutPeer) GetStrId() *types.StringValue {
	if m != nil {
		return m.StrId
	}
	return nil
}

// User's out peer
type UserOutPeer struct {
	Uid        int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	AccessHash int64 `protobuf:"varint,2,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty"`
}

func (m *UserOutPeer) Reset()      { *m = UserOutPeer{} }
func (*UserOutPeer) ProtoMessage() {}
func (*UserOutPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b029f9e0d26cab5, []int{2}
}
func (m *UserOutPeer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserOutPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserOutPeer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserOutPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOutPeer.Merge(m, src)
}
func (m *UserOutPeer) XXX_Size() int {
	return m.Size()
}
func (m *UserOutPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOutPeer.DiscardUnknown(m)
}

var xxx_messageInfo_UserOutPeer proto.InternalMessageInfo

func (m *UserOutPeer) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserOutPeer) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

// Group's out peer
type GroupOutPeer struct {
	GroupId    int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	AccessHash int64 `protobuf:"varint,2,opt,name=access_hash,json=accessHash,proto3" json:"access_hash,omitempty"`
}

func (m *GroupOutPeer) Reset()      { *m = GroupOutPeer{} }
func (*GroupOutPeer) ProtoMessage() {}
func (*GroupOutPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b029f9e0d26cab5, []int{3}
}
func (m *GroupOutPeer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupOutPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupOutPeer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupOutPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupOutPeer.Merge(m, src)
}
func (m *GroupOutPeer) XXX_Size() int {
	return m.Size()
}
func (m *GroupOutPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupOutPeer.DiscardUnknown(m)
}

var xxx_messageInfo_GroupOutPeer proto.InternalMessageInfo

func (m *GroupOutPeer) GetGroupId() int32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupOutPeer) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

func init() {
	proto.RegisterEnum("dialog.PeerType", PeerType_name, PeerType_value)
	proto.RegisterType((*Peer)(nil), "dialog.Peer")
	proto.RegisterType((*OutPeer)(nil), "dialog.OutPeer")
	proto.RegisterType((*UserOutPeer)(nil), "dialog.UserOutPeer")
	proto.RegisterType((*GroupOutPeer)(nil), "dialog.GroupOutPeer")
}

func init() { proto.RegisterFile("peers.proto", fileDescriptor_9b029f9e0d26cab5) }

var fileDescriptor_9b029f9e0d26cab5 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x1c, 0xc4, 0xbd, 0x4e, 0x9a, 0xe4, 0xfb, 0x27, 0x5f, 0x65, 0xac, 0x82, 0x42, 0x45, 0x97, 0x28,
	0xa7, 0x08, 0x09, 0xa7, 0x6a, 0x1f, 0x00, 0x11, 0xb0, 0x4a, 0x84, 0x94, 0x58, 0x6e, 0x52, 0x54,
	0x24, 0x14, 0x6d, 0xbc, 0xdb, 0xcd, 0x4a, 0x26, 0x5e, 0xed, 0xda, 0x45, 0x15, 0x17, 0xce, 0x3d,
	0xf1, 0x04, 0x9c, 0x79, 0x14, 0x8e, 0x39, 0xf6, 0x48, 0x9c, 0x0b, 0xea, 0xa9, 0x8f, 0x80, 0xe2,
	0xd4, 0x11, 0xb4, 0x02, 0x09, 0x24, 0x4e, 0x96, 0x67, 0xc6, 0x9e, 0xdf, 0xd8, 0x5a, 0xa8, 0x4a,
	0xc6, 0x94, 0x76, 0xa4, 0x8a, 0xe2, 0xc8, 0x2e, 0x51, 0x41, 0xc2, 0x88, 0x6f, 0x63, 0x1e, 0x45,
	0x3c, 0x64, 0xed, 0x4c, 0x1d, 0x27, 0x27, 0xed, 0x77, 0x8a, 0x48, 0xb9, 0xce, 0x6d, 0xdf, 0xa1,
	0xec, 0x44, 0x4c, 0x45, 0x2c, 0xa2, 0x69, 0x2e, 0xdd, 0xd5, 0x01, 0x09, 0x89, 0x1c, 0xb7, 0xaf,
	0xaf, 0x2b, 0xb9, 0xf9, 0x09, 0x41, 0xd1, 0x63, 0x4c, 0xd9, 0xfb, 0x50, 0x8c, 0xcf, 0x24, 0xab,
	0xa3, 0x06, 0x6a, 0x6d, 0xee, 0x59, 0xce, 0xaa, 0xc9, 0x59, 0x7a, 0x83, 0x33, 0xc9, 0x3a, 0xff,
	0x9f, 0x5f, 0xee, 0xfe, 0x07, 0xe5, 0x53, 0xa1, 0xc5, 0x38, 0x64, 0x7e, 0x16, 0xb6, 0x77, 0xc0,
	0x14, 0xb4, 0x6e, 0x36, 0x50, 0x6b, 0xe3, 0x66, 0xc0, 0x14, 0xd4, 0xee, 0x40, 0x49, 0xc7, 0x6a,
	0x24, 0x68, 0xbd, 0xd0, 0x40, 0xad, 0xea, 0xde, 0x03, 0x67, 0xc5, 0xed, 0xe4, 0xdc, 0xce, 0x61,
	0xac, 0xc4, 0x94, 0x1f, 0x91, 0x30, 0xb9, 0xd5, 0xb0, 0xa1, 0x63, 0xd5, 0xa5, 0xcd, 0x19, 0x82,
	0x72, 0x3f, 0x89, 0xff, 0x19, 0xe3, 0x63, 0xa8, 0x92, 0x20, 0x60, 0x5a, 0x8f, 0x26, 0x44, 0x4f,
	0x32, 0xd0, 0x42, 0xa7, 0x76, 0x7e, 0xb9, 0x5b, 0x81, 0x12, 0x25, 0x53, 0xce, 0x94, 0x0f, 0xab,
	0xc0, 0x0b, 0xa2, 0x27, 0x3f, 0x4c, 0x2a, 0xfe, 0xf5, 0xa4, 0x37, 0x50, 0x1d, 0x6a, 0xa6, 0xf2,
	0x55, 0x0f, 0xa1, 0x90, 0x08, 0x9a, 0x8d, 0xba, 0x45, 0xb8, 0x74, 0x6e, 0x22, 0x9a, 0xbf, 0x47,
	0x6c, 0x72, 0xa8, 0x1d, 0xa8, 0x28, 0x91, 0xf9, 0xfb, 0x5b, 0x50, 0xe1, 0xcb, 0xfb, 0xd1, 0xaf,
	0x4a, 0xca, 0x99, 0xdd, 0xfd, 0xd3, 0xa2, 0x47, 0xef, 0xa1, 0x92, 0x7f, 0x7a, 0x7b, 0x0b, 0x2c,
	0xcf, 0x75, 0xfd, 0xc1, 0xb1, 0xe7, 0x8e, 0x86, 0xbd, 0x97, 0xbd, 0xfe, 0xab, 0x9e, 0x65, 0xfc,
	0xa4, 0x7a, 0x7e, 0xf7, 0xe8, 0xe9, 0xc0, 0xb5, 0x90, 0x6d, 0xc3, 0xe6, 0x5a, 0x3d, 0xf0, 0xfb,
	0x43, 0xcf, 0x32, 0xed, 0x1d, 0xb8, 0xbf, 0xd6, 0xdc, 0xde, 0x33, 0xff, 0xd8, 0x1b, 0xb8, 0xcf,
	0xf3, 0x47, 0x0a, 0xb6, 0x05, 0xb5, 0xb5, 0x7d, 0xd8, 0xf5, 0xac, 0x62, 0x67, 0x98, 0x3e, 0xb9,
	0x07, 0x5b, 0xe2, 0xad, 0x43, 0x43, 0xee, 0x70, 0x25, 0x03, 0x47, 0x33, 0x75, 0x2a, 0x02, 0xa6,
	0x67, 0x73, 0x6c, 0x5c, 0xcc, 0xb1, 0x71, 0x35, 0xc7, 0xe8, 0x43, 0x8a, 0xd1, 0xe7, 0x14, 0xa3,
	0x2f, 0x29, 0x46, 0xb3, 0x14, 0xa3, 0xaf, 0x29, 0x46, 0xdf, 0x52, 0x6c, 0x5c, 0xa5, 0x18, 0x7d,
	0x5c, 0x60, 0x63, 0xb6, 0xc0, 0xc6, 0xc5, 0x02, 0x1b, 0xaf, 0xaf, 0x4f, 0xd6, 0xb8, 0x94, 0xfd,
	0xc7, 0xfd, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0x02, 0x40, 0xca, 0x77, 0x03, 0x00, 0x00,
}

func (x PeerType) String() string {
	s, ok := PeerType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Peer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Peer)
	if !ok {
		that2, ok := that.(Peer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if !this.StrId.Equal(that1.StrId) {
		return false
	}
	return true
}
func (this *OutPeer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OutPeer)
	if !ok {
		that2, ok := that.(OutPeer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.AccessHash != that1.AccessHash {
		return false
	}
	if !this.StrId.Equal(that1.StrId) {
		return false
	}
	return true
}
func (this *UserOutPeer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserOutPeer)
	if !ok {
		that2, ok := that.(UserOutPeer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Uid != that1.Uid {
		return false
	}
	if this.AccessHash != that1.AccessHash {
		return false
	}
	return true
}
func (this *GroupOutPeer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GroupOutPeer)
	if !ok {
		that2, ok := that.(GroupOutPeer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GroupId != that1.GroupId {
		return false
	}
	if this.AccessHash != that1.AccessHash {
		return false
	}
	return true
}
func (this *Peer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&dialog.Peer{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	if this.StrId != nil {
		s = append(s, "StrId: "+fmt.Sprintf("%#v", this.StrId)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *OutPeer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&dialog.OutPeer{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	if this.StrId != nil {
		s = append(s, "StrId: "+fmt.Sprintf("%#v", this.StrId)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UserOutPeer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dialog.UserOutPeer{")
	s = append(s, "Uid: "+fmt.Sprintf("%#v", this.Uid)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GroupOutPeer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dialog.GroupOutPeer{")
	s = append(s, "GroupId: "+fmt.Sprintf("%#v", this.GroupId)+",\n")
	s = append(s, "AccessHash: "+fmt.Sprintf("%#v", this.AccessHash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPeers(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StrId != nil {
		{
			size, err := m.StrId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPeers(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutPeer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutPeer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutPeer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StrId != nil {
		{
			size, err := m.StrId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPeers(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.AccessHash != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.AccessHash))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserOutPeer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserOutPeer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserOutPeer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccessHash != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.AccessHash))
		i--
		dAtA[i] = 0x10
	}
	if m.Uid != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GroupOutPeer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupOutPeer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupOutPeer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccessHash != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.AccessHash))
		i--
		dAtA[i] = 0x10
	}
	if m.GroupId != 0 {
		i = encodeVarintPeers(dAtA, i, uint64(m.GroupId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPeers(dAtA []byte, offset int, v uint64) int {
	offset -= sovPeers(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPeers(uint64(m.Type))
	}
	if m.Id != 0 {
		n += 1 + sovPeers(uint64(m.Id))
	}
	if m.StrId != nil {
		l = m.StrId.Size()
		n += 1 + l + sovPeers(uint64(l))
	}
	return n
}

func (m *OutPeer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPeers(uint64(m.Type))
	}
	if m.Id != 0 {
		n += 1 + sovPeers(uint64(m.Id))
	}
	if m.AccessHash != 0 {
		n += 1 + sovPeers(uint64(m.AccessHash))
	}
	if m.StrId != nil {
		l = m.StrId.Size()
		n += 1 + l + sovPeers(uint64(l))
	}
	return n
}

func (m *UserOutPeer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovPeers(uint64(m.Uid))
	}
	if m.AccessHash != 0 {
		n += 1 + sovPeers(uint64(m.AccessHash))
	}
	return n
}

func (m *GroupOutPeer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupId != 0 {
		n += 1 + sovPeers(uint64(m.GroupId))
	}
	if m.AccessHash != 0 {
		n += 1 + sovPeers(uint64(m.AccessHash))
	}
	return n
}

func sovPeers(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPeers(x uint64) (n int) {
	return sovPeers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Peer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Peer{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`StrId:` + strings.Replace(fmt.Sprintf("%v", this.StrId), "StringValue", "types.StringValue", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *OutPeer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OutPeer{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`AccessHash:` + fmt.Sprintf("%v", this.AccessHash) + `,`,
		`StrId:` + strings.Replace(fmt.Sprintf("%v", this.StrId), "StringValue", "types.StringValue", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UserOutPeer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UserOutPeer{`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`AccessHash:` + fmt.Sprintf("%v", this.AccessHash) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GroupOutPeer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GroupOutPeer{`,
		`GroupId:` + fmt.Sprintf("%v", this.GroupId) + `,`,
		`AccessHash:` + fmt.Sprintf("%v", this.AccessHash) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPeers(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPeers
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PeerType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
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
				return ErrInvalidLengthPeers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPeers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StrId == nil {
				m.StrId = &types.StringValue{}
			}
			if err := m.StrId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPeers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPeers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPeers
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
func (m *OutPeer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPeers
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
			return fmt.Errorf("proto: OutPeer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutPeer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PeerType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
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
				return ErrInvalidLengthPeers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPeers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StrId == nil {
				m.StrId = &types.StringValue{}
			}
			if err := m.StrId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPeers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPeers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPeers
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
func (m *UserOutPeer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPeers
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
			return fmt.Errorf("proto: UserOutPeer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserOutPeer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPeers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPeers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPeers
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
func (m *GroupOutPeer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPeers
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
			return fmt.Errorf("proto: GroupOutPeer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupOutPeer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessHash", wireType)
			}
			m.AccessHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessHash |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPeers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPeers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPeers
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
func skipPeers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPeers
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
					return 0, ErrIntOverflowPeers
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
					return 0, ErrIntOverflowPeers
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
				return 0, ErrInvalidLengthPeers
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPeers
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPeers
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPeers        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPeers          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPeers = fmt.Errorf("proto: unexpected end of group")
)
