// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: append_entries_request.proto

package protobuf

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AppendEntriesRequest struct {
	Term                 *uint64     `protobuf:"varint,1,req,name=Term" json:"Term,omitempty"`
	PrevLogIndex         *uint64     `protobuf:"varint,2,req,name=PrevLogIndex" json:"PrevLogIndex,omitempty"`
	PrevLogTerm          *uint64     `protobuf:"varint,3,req,name=PrevLogTerm" json:"PrevLogTerm,omitempty"`
	CommitIndex          *uint64     `protobuf:"varint,4,req,name=CommitIndex" json:"CommitIndex,omitempty"`
	LeaderName           *string     `protobuf:"bytes,5,req,name=LeaderName" json:"LeaderName,omitempty"`
	Entries              []*LogEntry `protobuf:"bytes,6,rep,name=Entries" json:"Entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppendEntriesRequest) Reset()      { *m = AppendEntriesRequest{} }
func (*AppendEntriesRequest) ProtoMessage() {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92c1c73991de9c02, []int{0}
}
func (m *AppendEntriesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppendEntriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppendEntriesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppendEntriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesRequest.Merge(m, src)
}
func (m *AppendEntriesRequest) XXX_Size() int {
	return m.Size()
}
func (m *AppendEntriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesRequest proto.InternalMessageInfo

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil && m.PrevLogIndex != nil {
		return *m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil && m.PrevLogTerm != nil {
		return *m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*AppendEntriesRequest)(nil), "protobuf.AppendEntriesRequest")
}

func init() { proto.RegisterFile("append_entries_request.proto", fileDescriptor_92c1c73991de9c02) }

var fileDescriptor_92c1c73991de9c02 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2c, 0x28, 0x48,
	0xcd, 0x4b, 0x89, 0x4f, 0xcd, 0x2b, 0x29, 0xca, 0x4c, 0x2d, 0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x52,
	0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9,
	0xfa, 0x30, 0x19, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0x1a, 0xa5, 0xf8, 0x73, 0xf2, 0xd3, 0xc1,
	0x66, 0x56, 0x42, 0x04, 0x94, 0x1e, 0x30, 0x72, 0x89, 0x38, 0x82, 0xad, 0x72, 0x85, 0xd8, 0x14,
	0x04, 0xb1, 0x48, 0x48, 0x88, 0x8b, 0x25, 0x24, 0xb5, 0x28, 0x57, 0x82, 0x51, 0x81, 0x49, 0x83,
	0x25, 0x08, 0xcc, 0x16, 0x52, 0xe2, 0xe2, 0x09, 0x28, 0x4a, 0x2d, 0xf3, 0xc9, 0x4f, 0xf7, 0xcc,
	0x4b, 0x49, 0xad, 0x90, 0x60, 0x02, 0xcb, 0xa1, 0x88, 0x09, 0x29, 0x70, 0x71, 0x43, 0xf9, 0x60,
	0xed, 0xcc, 0x60, 0x25, 0xc8, 0x42, 0x20, 0x15, 0xce, 0xf9, 0xb9, 0xb9, 0x99, 0x25, 0x10, 0x43,
	0x58, 0x20, 0x2a, 0x90, 0x84, 0x84, 0xe4, 0xb8, 0xb8, 0x7c, 0x52, 0x13, 0x53, 0x52, 0x8b, 0xfc,
	0x12, 0x73, 0x53, 0x25, 0x58, 0x15, 0x98, 0x34, 0x38, 0x83, 0x90, 0x44, 0x84, 0x74, 0xb8, 0xd8,
	0xa1, 0xae, 0x95, 0x60, 0x53, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd2, 0x83, 0x79, 0x5b, 0xcf, 0x27,
	0x3f, 0x1d, 0x24, 0x57, 0x19, 0x04, 0x53, 0xe2, 0x64, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x83, 0x87,
	0x72, 0x8c, 0x1f, 0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5,
	0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x0f, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x00, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x38, 0x93, 0x74, 0xd3, 0x7d, 0x01, 0x00, 0x00,
}

func (this *AppendEntriesRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AppendEntriesRequest)
	if !ok {
		that2, ok := that.(AppendEntriesRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AppendEntriesRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AppendEntriesRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AppendEntriesRequest but is not nil && this == nil")
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return fmt.Errorf("Term this(%v) Not Equal that(%v)", *this.Term, *that1.Term)
		}
	} else if this.Term != nil {
		return fmt.Errorf("this.Term == nil && that.Term != nil")
	} else if that1.Term != nil {
		return fmt.Errorf("Term this(%v) Not Equal that(%v)", this.Term, that1.Term)
	}
	if this.PrevLogIndex != nil && that1.PrevLogIndex != nil {
		if *this.PrevLogIndex != *that1.PrevLogIndex {
			return fmt.Errorf("PrevLogIndex this(%v) Not Equal that(%v)", *this.PrevLogIndex, *that1.PrevLogIndex)
		}
	} else if this.PrevLogIndex != nil {
		return fmt.Errorf("this.PrevLogIndex == nil && that.PrevLogIndex != nil")
	} else if that1.PrevLogIndex != nil {
		return fmt.Errorf("PrevLogIndex this(%v) Not Equal that(%v)", this.PrevLogIndex, that1.PrevLogIndex)
	}
	if this.PrevLogTerm != nil && that1.PrevLogTerm != nil {
		if *this.PrevLogTerm != *that1.PrevLogTerm {
			return fmt.Errorf("PrevLogTerm this(%v) Not Equal that(%v)", *this.PrevLogTerm, *that1.PrevLogTerm)
		}
	} else if this.PrevLogTerm != nil {
		return fmt.Errorf("this.PrevLogTerm == nil && that.PrevLogTerm != nil")
	} else if that1.PrevLogTerm != nil {
		return fmt.Errorf("PrevLogTerm this(%v) Not Equal that(%v)", this.PrevLogTerm, that1.PrevLogTerm)
	}
	if this.CommitIndex != nil && that1.CommitIndex != nil {
		if *this.CommitIndex != *that1.CommitIndex {
			return fmt.Errorf("CommitIndex this(%v) Not Equal that(%v)", *this.CommitIndex, *that1.CommitIndex)
		}
	} else if this.CommitIndex != nil {
		return fmt.Errorf("this.CommitIndex == nil && that.CommitIndex != nil")
	} else if that1.CommitIndex != nil {
		return fmt.Errorf("CommitIndex this(%v) Not Equal that(%v)", this.CommitIndex, that1.CommitIndex)
	}
	if this.LeaderName != nil && that1.LeaderName != nil {
		if *this.LeaderName != *that1.LeaderName {
			return fmt.Errorf("LeaderName this(%v) Not Equal that(%v)", *this.LeaderName, *that1.LeaderName)
		}
	} else if this.LeaderName != nil {
		return fmt.Errorf("this.LeaderName == nil && that.LeaderName != nil")
	} else if that1.LeaderName != nil {
		return fmt.Errorf("LeaderName this(%v) Not Equal that(%v)", this.LeaderName, that1.LeaderName)
	}
	if len(this.Entries) != len(that1.Entries) {
		return fmt.Errorf("Entries this(%v) Not Equal that(%v)", len(this.Entries), len(that1.Entries))
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(that1.Entries[i]) {
			return fmt.Errorf("Entries this[%v](%v) Not Equal that[%v](%v)", i, this.Entries[i], i, that1.Entries[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *AppendEntriesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AppendEntriesRequest)
	if !ok {
		that2, ok := that.(AppendEntriesRequest)
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
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return false
		}
	} else if this.Term != nil {
		return false
	} else if that1.Term != nil {
		return false
	}
	if this.PrevLogIndex != nil && that1.PrevLogIndex != nil {
		if *this.PrevLogIndex != *that1.PrevLogIndex {
			return false
		}
	} else if this.PrevLogIndex != nil {
		return false
	} else if that1.PrevLogIndex != nil {
		return false
	}
	if this.PrevLogTerm != nil && that1.PrevLogTerm != nil {
		if *this.PrevLogTerm != *that1.PrevLogTerm {
			return false
		}
	} else if this.PrevLogTerm != nil {
		return false
	} else if that1.PrevLogTerm != nil {
		return false
	}
	if this.CommitIndex != nil && that1.CommitIndex != nil {
		if *this.CommitIndex != *that1.CommitIndex {
			return false
		}
	} else if this.CommitIndex != nil {
		return false
	} else if that1.CommitIndex != nil {
		return false
	}
	if this.LeaderName != nil && that1.LeaderName != nil {
		if *this.LeaderName != *that1.LeaderName {
			return false
		}
	} else if this.LeaderName != nil {
		return false
	} else if that1.LeaderName != nil {
		return false
	}
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(that1.Entries[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AppendEntriesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&protobuf.AppendEntriesRequest{")
	if this.Term != nil {
		s = append(s, "Term: "+valueToGoStringAppendEntriesRequest(this.Term, "uint64")+",\n")
	}
	if this.PrevLogIndex != nil {
		s = append(s, "PrevLogIndex: "+valueToGoStringAppendEntriesRequest(this.PrevLogIndex, "uint64")+",\n")
	}
	if this.PrevLogTerm != nil {
		s = append(s, "PrevLogTerm: "+valueToGoStringAppendEntriesRequest(this.PrevLogTerm, "uint64")+",\n")
	}
	if this.CommitIndex != nil {
		s = append(s, "CommitIndex: "+valueToGoStringAppendEntriesRequest(this.CommitIndex, "uint64")+",\n")
	}
	if this.LeaderName != nil {
		s = append(s, "LeaderName: "+valueToGoStringAppendEntriesRequest(this.LeaderName, "string")+",\n")
	}
	if this.Entries != nil {
		s = append(s, "Entries: "+fmt.Sprintf("%#v", this.Entries)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAppendEntriesRequest(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AppendEntriesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppendEntriesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAppendEntriesRequest(dAtA, i, uint64(*m.Term))
	}
	if m.PrevLogIndex == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAppendEntriesRequest(dAtA, i, uint64(*m.PrevLogIndex))
	}
	if m.PrevLogTerm == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAppendEntriesRequest(dAtA, i, uint64(*m.PrevLogTerm))
	}
	if m.CommitIndex == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAppendEntriesRequest(dAtA, i, uint64(*m.CommitIndex))
	}
	if m.LeaderName == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAppendEntriesRequest(dAtA, i, uint64(len(*m.LeaderName)))
		i += copy(dAtA[i:], *m.LeaderName)
	}
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0x32
			i++
			i = encodeVarintAppendEntriesRequest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAppendEntriesRequest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAppendEntriesRequest(r randyAppendEntriesRequest, easy bool) *AppendEntriesRequest {
	this := &AppendEntriesRequest{}
	v1 := uint64(uint64(r.Uint32()))
	this.Term = &v1
	v2 := uint64(uint64(r.Uint32()))
	this.PrevLogIndex = &v2
	v3 := uint64(uint64(r.Uint32()))
	this.PrevLogTerm = &v3
	v4 := uint64(uint64(r.Uint32()))
	this.CommitIndex = &v4
	v5 := string(randStringAppendEntriesRequest(r))
	this.LeaderName = &v5
	if r.Intn(10) != 0 {
		v6 := r.Intn(5)
		this.Entries = make([]*LogEntry, v6)
		for i := 0; i < v6; i++ {
			this.Entries[i] = NewPopulatedLogEntry(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAppendEntriesRequest(r, 7)
	}
	return this
}

type randyAppendEntriesRequest interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAppendEntriesRequest(r randyAppendEntriesRequest) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAppendEntriesRequest(r randyAppendEntriesRequest) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneAppendEntriesRequest(r)
	}
	return string(tmps)
}
func randUnrecognizedAppendEntriesRequest(r randyAppendEntriesRequest, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAppendEntriesRequest(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAppendEntriesRequest(dAtA []byte, r randyAppendEntriesRequest, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAppendEntriesRequest(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateAppendEntriesRequest(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateAppendEntriesRequest(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAppendEntriesRequest(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAppendEntriesRequest(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAppendEntriesRequest(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAppendEntriesRequest(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AppendEntriesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Term != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.Term))
	}
	if m.PrevLogIndex != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.PrevLogIndex))
	}
	if m.PrevLogTerm != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.PrevLogTerm))
	}
	if m.CommitIndex != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.CommitIndex))
	}
	if m.LeaderName != nil {
		l = len(*m.LeaderName)
		n += 1 + l + sovAppendEntriesRequest(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovAppendEntriesRequest(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAppendEntriesRequest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAppendEntriesRequest(x uint64) (n int) {
	return sovAppendEntriesRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AppendEntriesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AppendEntriesRequest{`,
		`Term:` + valueToStringAppendEntriesRequest(this.Term) + `,`,
		`PrevLogIndex:` + valueToStringAppendEntriesRequest(this.PrevLogIndex) + `,`,
		`PrevLogTerm:` + valueToStringAppendEntriesRequest(this.PrevLogTerm) + `,`,
		`CommitIndex:` + valueToStringAppendEntriesRequest(this.CommitIndex) + `,`,
		`LeaderName:` + valueToStringAppendEntriesRequest(this.LeaderName) + `,`,
		`Entries:` + strings.Replace(fmt.Sprintf("%v", this.Entries), "LogEntry", "LogEntry", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAppendEntriesRequest(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AppendEntriesRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppendEntriesRequest
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
			return fmt.Errorf("proto: AppendEntriesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppendEntriesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppendEntriesRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Term = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevLogIndex", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppendEntriesRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevLogIndex = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevLogTerm", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppendEntriesRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevLogTerm = &v
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitIndex", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppendEntriesRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CommitIndex = &v
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppendEntriesRequest
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
				return ErrInvalidLengthAppendEntriesRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppendEntriesRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.LeaderName = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppendEntriesRequest
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
				return ErrInvalidLengthAppendEntriesRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAppendEntriesRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &LogEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppendEntriesRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAppendEntriesRequest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAppendEntriesRequest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAppendEntriesRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAppendEntriesRequest
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
					return 0, ErrIntOverflowAppendEntriesRequest
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
					return 0, ErrIntOverflowAppendEntriesRequest
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
				return 0, ErrInvalidLengthAppendEntriesRequest
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAppendEntriesRequest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAppendEntriesRequest
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
				next, err := skipAppendEntriesRequest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAppendEntriesRequest
				}
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
	ErrInvalidLengthAppendEntriesRequest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAppendEntriesRequest   = fmt.Errorf("proto: integer overflow")
)
