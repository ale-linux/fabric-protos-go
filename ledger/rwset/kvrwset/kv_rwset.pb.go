// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/rwset/kvrwset/kv_rwset.proto

package kvrwset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// KVRWSet encapsulates the read-write set for a chaincode that operates upon a KV or Document data model
// This structure is used for both the public data and the private data
type KVRWSet struct {
	Reads                []*KVRead          `protobuf:"bytes,1,rep,name=reads,proto3" json:"reads,omitempty"`
	RangeQueriesInfo     []*RangeQueryInfo  `protobuf:"bytes,2,rep,name=range_queries_info,json=rangeQueriesInfo,proto3" json:"range_queries_info,omitempty"`
	Writes               []*KVWrite         `protobuf:"bytes,3,rep,name=writes,proto3" json:"writes,omitempty"`
	MetadataWrites       []*KVMetadataWrite `protobuf:"bytes,4,rep,name=metadata_writes,json=metadataWrites,proto3" json:"metadata_writes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *KVRWSet) Reset()         { *m = KVRWSet{} }
func (m *KVRWSet) String() string { return proto.CompactTextString(m) }
func (*KVRWSet) ProtoMessage()    {}
func (*KVRWSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{0}
}

func (m *KVRWSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVRWSet.Unmarshal(m, b)
}
func (m *KVRWSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVRWSet.Marshal(b, m, deterministic)
}
func (m *KVRWSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVRWSet.Merge(m, src)
}
func (m *KVRWSet) XXX_Size() int {
	return xxx_messageInfo_KVRWSet.Size(m)
}
func (m *KVRWSet) XXX_DiscardUnknown() {
	xxx_messageInfo_KVRWSet.DiscardUnknown(m)
}

var xxx_messageInfo_KVRWSet proto.InternalMessageInfo

func (m *KVRWSet) GetReads() []*KVRead {
	if m != nil {
		return m.Reads
	}
	return nil
}

func (m *KVRWSet) GetRangeQueriesInfo() []*RangeQueryInfo {
	if m != nil {
		return m.RangeQueriesInfo
	}
	return nil
}

func (m *KVRWSet) GetWrites() []*KVWrite {
	if m != nil {
		return m.Writes
	}
	return nil
}

func (m *KVRWSet) GetMetadataWrites() []*KVMetadataWrite {
	if m != nil {
		return m.MetadataWrites
	}
	return nil
}

// HashedRWSet encapsulates hashed representation of a private read-write set for KV or Document data model
type HashedRWSet struct {
	HashedReads          []*KVReadHash          `protobuf:"bytes,1,rep,name=hashed_reads,json=hashedReads,proto3" json:"hashed_reads,omitempty"`
	HashedWrites         []*KVWriteHash         `protobuf:"bytes,2,rep,name=hashed_writes,json=hashedWrites,proto3" json:"hashed_writes,omitempty"`
	MetadataWrites       []*KVMetadataWriteHash `protobuf:"bytes,3,rep,name=metadata_writes,json=metadataWrites,proto3" json:"metadata_writes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HashedRWSet) Reset()         { *m = HashedRWSet{} }
func (m *HashedRWSet) String() string { return proto.CompactTextString(m) }
func (*HashedRWSet) ProtoMessage()    {}
func (*HashedRWSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{1}
}

func (m *HashedRWSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashedRWSet.Unmarshal(m, b)
}
func (m *HashedRWSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashedRWSet.Marshal(b, m, deterministic)
}
func (m *HashedRWSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashedRWSet.Merge(m, src)
}
func (m *HashedRWSet) XXX_Size() int {
	return xxx_messageInfo_HashedRWSet.Size(m)
}
func (m *HashedRWSet) XXX_DiscardUnknown() {
	xxx_messageInfo_HashedRWSet.DiscardUnknown(m)
}

var xxx_messageInfo_HashedRWSet proto.InternalMessageInfo

func (m *HashedRWSet) GetHashedReads() []*KVReadHash {
	if m != nil {
		return m.HashedReads
	}
	return nil
}

func (m *HashedRWSet) GetHashedWrites() []*KVWriteHash {
	if m != nil {
		return m.HashedWrites
	}
	return nil
}

func (m *HashedRWSet) GetMetadataWrites() []*KVMetadataWriteHash {
	if m != nil {
		return m.MetadataWrites
	}
	return nil
}

// KVRead captures a read operation performed during transaction simulation
// A 'nil' version indicates a non-existing key read by the transaction
type KVRead struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Version              *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	SpeculativeVersion   uint64   `protobuf:"varint,3,opt,name=speculative_version,json=speculativeVersion,proto3" json:"speculative_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVRead) Reset()         { *m = KVRead{} }
func (m *KVRead) String() string { return proto.CompactTextString(m) }
func (*KVRead) ProtoMessage()    {}
func (*KVRead) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{2}
}

func (m *KVRead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVRead.Unmarshal(m, b)
}
func (m *KVRead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVRead.Marshal(b, m, deterministic)
}
func (m *KVRead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVRead.Merge(m, src)
}
func (m *KVRead) XXX_Size() int {
	return xxx_messageInfo_KVRead.Size(m)
}
func (m *KVRead) XXX_DiscardUnknown() {
	xxx_messageInfo_KVRead.DiscardUnknown(m)
}

var xxx_messageInfo_KVRead proto.InternalMessageInfo

func (m *KVRead) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVRead) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *KVRead) GetSpeculativeVersion() uint64 {
	if m != nil {
		return m.SpeculativeVersion
	}
	return 0
}

// KVWrite captures a write (update/delete) operation performed during transaction simulation
type KVWrite struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	IsDelete             bool     `protobuf:"varint,2,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	SpeculativeVersion   uint64   `protobuf:"varint,4,opt,name=speculative_version,json=speculativeVersion,proto3" json:"speculative_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVWrite) Reset()         { *m = KVWrite{} }
func (m *KVWrite) String() string { return proto.CompactTextString(m) }
func (*KVWrite) ProtoMessage()    {}
func (*KVWrite) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{3}
}

func (m *KVWrite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVWrite.Unmarshal(m, b)
}
func (m *KVWrite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVWrite.Marshal(b, m, deterministic)
}
func (m *KVWrite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVWrite.Merge(m, src)
}
func (m *KVWrite) XXX_Size() int {
	return xxx_messageInfo_KVWrite.Size(m)
}
func (m *KVWrite) XXX_DiscardUnknown() {
	xxx_messageInfo_KVWrite.DiscardUnknown(m)
}

var xxx_messageInfo_KVWrite proto.InternalMessageInfo

func (m *KVWrite) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVWrite) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func (m *KVWrite) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KVWrite) GetSpeculativeVersion() uint64 {
	if m != nil {
		return m.SpeculativeVersion
	}
	return 0
}

// KVMetadataWrite captures all the entries in the metadata associated with a key
type KVMetadataWrite struct {
	Key                  string             `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Entries              []*KVMetadataEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *KVMetadataWrite) Reset()         { *m = KVMetadataWrite{} }
func (m *KVMetadataWrite) String() string { return proto.CompactTextString(m) }
func (*KVMetadataWrite) ProtoMessage()    {}
func (*KVMetadataWrite) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{4}
}

func (m *KVMetadataWrite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVMetadataWrite.Unmarshal(m, b)
}
func (m *KVMetadataWrite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVMetadataWrite.Marshal(b, m, deterministic)
}
func (m *KVMetadataWrite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVMetadataWrite.Merge(m, src)
}
func (m *KVMetadataWrite) XXX_Size() int {
	return xxx_messageInfo_KVMetadataWrite.Size(m)
}
func (m *KVMetadataWrite) XXX_DiscardUnknown() {
	xxx_messageInfo_KVMetadataWrite.DiscardUnknown(m)
}

var xxx_messageInfo_KVMetadataWrite proto.InternalMessageInfo

func (m *KVMetadataWrite) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVMetadataWrite) GetEntries() []*KVMetadataEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// KVReadHash is similar to the KVRead in spirit. However, it captures the hash of the key instead of the key itself
// version is kept as is for now. However, if the version also needs to be privacy-protected, it would need to be the
// hash of the version and hence of 'bytes' type
type KVReadHash struct {
	KeyHash              []byte   `protobuf:"bytes,1,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	Version              *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVReadHash) Reset()         { *m = KVReadHash{} }
func (m *KVReadHash) String() string { return proto.CompactTextString(m) }
func (*KVReadHash) ProtoMessage()    {}
func (*KVReadHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{5}
}

func (m *KVReadHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVReadHash.Unmarshal(m, b)
}
func (m *KVReadHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVReadHash.Marshal(b, m, deterministic)
}
func (m *KVReadHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVReadHash.Merge(m, src)
}
func (m *KVReadHash) XXX_Size() int {
	return xxx_messageInfo_KVReadHash.Size(m)
}
func (m *KVReadHash) XXX_DiscardUnknown() {
	xxx_messageInfo_KVReadHash.DiscardUnknown(m)
}

var xxx_messageInfo_KVReadHash proto.InternalMessageInfo

func (m *KVReadHash) GetKeyHash() []byte {
	if m != nil {
		return m.KeyHash
	}
	return nil
}

func (m *KVReadHash) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

// KVWriteHash is similar to the KVWrite. It captures a write (update/delete) operation performed during transaction simulation
type KVWriteHash struct {
	KeyHash              []byte   `protobuf:"bytes,1,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	IsDelete             bool     `protobuf:"varint,2,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	ValueHash            []byte   `protobuf:"bytes,3,opt,name=value_hash,json=valueHash,proto3" json:"value_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVWriteHash) Reset()         { *m = KVWriteHash{} }
func (m *KVWriteHash) String() string { return proto.CompactTextString(m) }
func (*KVWriteHash) ProtoMessage()    {}
func (*KVWriteHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{6}
}

func (m *KVWriteHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVWriteHash.Unmarshal(m, b)
}
func (m *KVWriteHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVWriteHash.Marshal(b, m, deterministic)
}
func (m *KVWriteHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVWriteHash.Merge(m, src)
}
func (m *KVWriteHash) XXX_Size() int {
	return xxx_messageInfo_KVWriteHash.Size(m)
}
func (m *KVWriteHash) XXX_DiscardUnknown() {
	xxx_messageInfo_KVWriteHash.DiscardUnknown(m)
}

var xxx_messageInfo_KVWriteHash proto.InternalMessageInfo

func (m *KVWriteHash) GetKeyHash() []byte {
	if m != nil {
		return m.KeyHash
	}
	return nil
}

func (m *KVWriteHash) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func (m *KVWriteHash) GetValueHash() []byte {
	if m != nil {
		return m.ValueHash
	}
	return nil
}

// KVMetadataWriteHash captures all the upserts to the metadata associated with a key hash
type KVMetadataWriteHash struct {
	KeyHash              []byte             `protobuf:"bytes,1,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	Entries              []*KVMetadataEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *KVMetadataWriteHash) Reset()         { *m = KVMetadataWriteHash{} }
func (m *KVMetadataWriteHash) String() string { return proto.CompactTextString(m) }
func (*KVMetadataWriteHash) ProtoMessage()    {}
func (*KVMetadataWriteHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{7}
}

func (m *KVMetadataWriteHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVMetadataWriteHash.Unmarshal(m, b)
}
func (m *KVMetadataWriteHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVMetadataWriteHash.Marshal(b, m, deterministic)
}
func (m *KVMetadataWriteHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVMetadataWriteHash.Merge(m, src)
}
func (m *KVMetadataWriteHash) XXX_Size() int {
	return xxx_messageInfo_KVMetadataWriteHash.Size(m)
}
func (m *KVMetadataWriteHash) XXX_DiscardUnknown() {
	xxx_messageInfo_KVMetadataWriteHash.DiscardUnknown(m)
}

var xxx_messageInfo_KVMetadataWriteHash proto.InternalMessageInfo

func (m *KVMetadataWriteHash) GetKeyHash() []byte {
	if m != nil {
		return m.KeyHash
	}
	return nil
}

func (m *KVMetadataWriteHash) GetEntries() []*KVMetadataEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// KVMetadataEntry captures a 'name'ed entry in the metadata of a key/key-hash.
type KVMetadataEntry struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVMetadataEntry) Reset()         { *m = KVMetadataEntry{} }
func (m *KVMetadataEntry) String() string { return proto.CompactTextString(m) }
func (*KVMetadataEntry) ProtoMessage()    {}
func (*KVMetadataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{8}
}

func (m *KVMetadataEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVMetadataEntry.Unmarshal(m, b)
}
func (m *KVMetadataEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVMetadataEntry.Marshal(b, m, deterministic)
}
func (m *KVMetadataEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVMetadataEntry.Merge(m, src)
}
func (m *KVMetadataEntry) XXX_Size() int {
	return xxx_messageInfo_KVMetadataEntry.Size(m)
}
func (m *KVMetadataEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_KVMetadataEntry.DiscardUnknown(m)
}

var xxx_messageInfo_KVMetadataEntry proto.InternalMessageInfo

func (m *KVMetadataEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KVMetadataEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Version encapsulates the version of a Key
// A version of a committed key is maintained as the height of the transaction that committed the key.
// The height is represenetd as a tuple <blockNum, txNum> where the txNum is the position of the transaction
// (starting with 0) within block
type Version struct {
	BlockNum             uint64   `protobuf:"varint,1,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	TxNum                uint64   `protobuf:"varint,2,opt,name=tx_num,json=txNum,proto3" json:"tx_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{9}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *Version) GetTxNum() uint64 {
	if m != nil {
		return m.TxNum
	}
	return 0
}

// RangeQueryInfo encapsulates the details of a range query performed by a transaction during simulation.
// This helps protect transactions from phantom reads by varifying during validation whether any new items
// got committed within the given range between transaction simuation and validation
// (in addition to regular checks for updates/deletes of the existing items).
// readInfo field contains either the KVReads (for the items read by the range query) or a merkle-tree hash
// if the KVReads exceeds a pre-configured numbers
type RangeQueryInfo struct {
	StartKey     string `protobuf:"bytes,1,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey       string `protobuf:"bytes,2,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
	ItrExhausted bool   `protobuf:"varint,3,opt,name=itr_exhausted,json=itrExhausted,proto3" json:"itr_exhausted,omitempty"`
	// Types that are valid to be assigned to ReadsInfo:
	//	*RangeQueryInfo_RawReads
	//	*RangeQueryInfo_ReadsMerkleHashes
	ReadsInfo            isRangeQueryInfo_ReadsInfo `protobuf_oneof:"reads_info"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RangeQueryInfo) Reset()         { *m = RangeQueryInfo{} }
func (m *RangeQueryInfo) String() string { return proto.CompactTextString(m) }
func (*RangeQueryInfo) ProtoMessage()    {}
func (*RangeQueryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{10}
}

func (m *RangeQueryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RangeQueryInfo.Unmarshal(m, b)
}
func (m *RangeQueryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RangeQueryInfo.Marshal(b, m, deterministic)
}
func (m *RangeQueryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeQueryInfo.Merge(m, src)
}
func (m *RangeQueryInfo) XXX_Size() int {
	return xxx_messageInfo_RangeQueryInfo.Size(m)
}
func (m *RangeQueryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeQueryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RangeQueryInfo proto.InternalMessageInfo

func (m *RangeQueryInfo) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *RangeQueryInfo) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *RangeQueryInfo) GetItrExhausted() bool {
	if m != nil {
		return m.ItrExhausted
	}
	return false
}

type isRangeQueryInfo_ReadsInfo interface {
	isRangeQueryInfo_ReadsInfo()
}

type RangeQueryInfo_RawReads struct {
	RawReads *QueryReads `protobuf:"bytes,4,opt,name=raw_reads,json=rawReads,proto3,oneof"`
}

type RangeQueryInfo_ReadsMerkleHashes struct {
	ReadsMerkleHashes *QueryReadsMerkleSummary `protobuf:"bytes,5,opt,name=reads_merkle_hashes,json=readsMerkleHashes,proto3,oneof"`
}

func (*RangeQueryInfo_RawReads) isRangeQueryInfo_ReadsInfo() {}

func (*RangeQueryInfo_ReadsMerkleHashes) isRangeQueryInfo_ReadsInfo() {}

func (m *RangeQueryInfo) GetReadsInfo() isRangeQueryInfo_ReadsInfo {
	if m != nil {
		return m.ReadsInfo
	}
	return nil
}

func (m *RangeQueryInfo) GetRawReads() *QueryReads {
	if x, ok := m.GetReadsInfo().(*RangeQueryInfo_RawReads); ok {
		return x.RawReads
	}
	return nil
}

func (m *RangeQueryInfo) GetReadsMerkleHashes() *QueryReadsMerkleSummary {
	if x, ok := m.GetReadsInfo().(*RangeQueryInfo_ReadsMerkleHashes); ok {
		return x.ReadsMerkleHashes
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RangeQueryInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RangeQueryInfo_RawReads)(nil),
		(*RangeQueryInfo_ReadsMerkleHashes)(nil),
	}
}

// QueryReads encapsulates the KVReads for the items read by a transaction as a result of a query execution
type QueryReads struct {
	KvReads              []*KVRead `protobuf:"bytes,1,rep,name=kv_reads,json=kvReads,proto3" json:"kv_reads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryReads) Reset()         { *m = QueryReads{} }
func (m *QueryReads) String() string { return proto.CompactTextString(m) }
func (*QueryReads) ProtoMessage()    {}
func (*QueryReads) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{11}
}

func (m *QueryReads) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReads.Unmarshal(m, b)
}
func (m *QueryReads) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReads.Marshal(b, m, deterministic)
}
func (m *QueryReads) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReads.Merge(m, src)
}
func (m *QueryReads) XXX_Size() int {
	return xxx_messageInfo_QueryReads.Size(m)
}
func (m *QueryReads) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReads.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReads proto.InternalMessageInfo

func (m *QueryReads) GetKvReads() []*KVRead {
	if m != nil {
		return m.KvReads
	}
	return nil
}

// QueryReadsMerkleSummary encapsulates the Merkle-tree hashes for the QueryReads
// This allows to reduce the size of RWSet in the presence of query results
// by storing certain hashes instead of actual results.
// maxDegree field refers to the maximum number of children in the tree at any level
// maxLevel field contains the lowest level which has lesser nodes than maxDegree (starting from leaf level)
type QueryReadsMerkleSummary struct {
	MaxDegree            uint32   `protobuf:"varint,1,opt,name=max_degree,json=maxDegree,proto3" json:"max_degree,omitempty"`
	MaxLevel             uint32   `protobuf:"varint,2,opt,name=max_level,json=maxLevel,proto3" json:"max_level,omitempty"`
	MaxLevelHashes       [][]byte `protobuf:"bytes,3,rep,name=max_level_hashes,json=maxLevelHashes,proto3" json:"max_level_hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReadsMerkleSummary) Reset()         { *m = QueryReadsMerkleSummary{} }
func (m *QueryReadsMerkleSummary) String() string { return proto.CompactTextString(m) }
func (*QueryReadsMerkleSummary) ProtoMessage()    {}
func (*QueryReadsMerkleSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee5d686eab23a142, []int{12}
}

func (m *QueryReadsMerkleSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReadsMerkleSummary.Unmarshal(m, b)
}
func (m *QueryReadsMerkleSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReadsMerkleSummary.Marshal(b, m, deterministic)
}
func (m *QueryReadsMerkleSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReadsMerkleSummary.Merge(m, src)
}
func (m *QueryReadsMerkleSummary) XXX_Size() int {
	return xxx_messageInfo_QueryReadsMerkleSummary.Size(m)
}
func (m *QueryReadsMerkleSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReadsMerkleSummary.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReadsMerkleSummary proto.InternalMessageInfo

func (m *QueryReadsMerkleSummary) GetMaxDegree() uint32 {
	if m != nil {
		return m.MaxDegree
	}
	return 0
}

func (m *QueryReadsMerkleSummary) GetMaxLevel() uint32 {
	if m != nil {
		return m.MaxLevel
	}
	return 0
}

func (m *QueryReadsMerkleSummary) GetMaxLevelHashes() [][]byte {
	if m != nil {
		return m.MaxLevelHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*KVRWSet)(nil), "kvrwset.KVRWSet")
	proto.RegisterType((*HashedRWSet)(nil), "kvrwset.HashedRWSet")
	proto.RegisterType((*KVRead)(nil), "kvrwset.KVRead")
	proto.RegisterType((*KVWrite)(nil), "kvrwset.KVWrite")
	proto.RegisterType((*KVMetadataWrite)(nil), "kvrwset.KVMetadataWrite")
	proto.RegisterType((*KVReadHash)(nil), "kvrwset.KVReadHash")
	proto.RegisterType((*KVWriteHash)(nil), "kvrwset.KVWriteHash")
	proto.RegisterType((*KVMetadataWriteHash)(nil), "kvrwset.KVMetadataWriteHash")
	proto.RegisterType((*KVMetadataEntry)(nil), "kvrwset.KVMetadataEntry")
	proto.RegisterType((*Version)(nil), "kvrwset.Version")
	proto.RegisterType((*RangeQueryInfo)(nil), "kvrwset.RangeQueryInfo")
	proto.RegisterType((*QueryReads)(nil), "kvrwset.QueryReads")
	proto.RegisterType((*QueryReadsMerkleSummary)(nil), "kvrwset.QueryReadsMerkleSummary")
}

func init() {
	proto.RegisterFile("ledger/rwset/kvrwset/kv_rwset.proto", fileDescriptor_ee5d686eab23a142)
}

var fileDescriptor_ee5d686eab23a142 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x6f, 0x6b, 0x13, 0x4d,
	0x10, 0x7f, 0xf2, 0xff, 0x32, 0x49, 0xda, 0x3c, 0x9b, 0x3e, 0x34, 0x0f, 0x2a, 0x94, 0x2b, 0x42,
	0x11, 0x9a, 0x40, 0x05, 0x51, 0xd4, 0x17, 0x4a, 0x0b, 0x95, 0x5a, 0xc1, 0x2d, 0xb4, 0xe0, 0x9b,
	0x63, 0x93, 0xdb, 0x26, 0x47, 0xee, 0x72, 0x75, 0x6f, 0x2f, 0x7f, 0x5e, 0x89, 0x9f, 0xce, 0x2f,
	0xe2, 0x07, 0x71, 0x77, 0x76, 0xd3, 0x5c, 0x63, 0x12, 0xd0, 0x57, 0xcd, 0xcc, 0x6f, 0x7e, 0xb3,
	0x33, 0xbf, 0x99, 0xce, 0xc1, 0x61, 0xc8, 0xfd, 0x01, 0x17, 0x5d, 0x31, 0x4d, 0xb8, 0xec, 0x8e,
	0x26, 0x8b, 0xbf, 0x1e, 0xfe, 0xe8, 0xdc, 0x89, 0x58, 0xc6, 0xa4, 0x62, 0xfd, 0xee, 0xcf, 0x1c,
	0x54, 0x2e, 0xae, 0xe9, 0xcd, 0x15, 0x97, 0xe4, 0x29, 0x94, 0x04, 0x67, 0x7e, 0xd2, 0xce, 0x1d,
	0x14, 0x8e, 0x6a, 0x27, 0xbb, 0x1d, 0x1b, 0xd4, 0x51, 0x01, 0xca, 0x4f, 0x0d, 0x4a, 0xce, 0x80,
	0x08, 0x36, 0x1e, 0x70, 0xef, 0x6b, 0xca, 0x45, 0xc0, 0x13, 0x2f, 0x18, 0xdf, 0xc6, 0xed, 0x3c,
	0x72, 0xf6, 0xef, 0x39, 0x54, 0x87, 0x7c, 0x56, 0x11, 0xf3, 0x0f, 0x0a, 0xa6, 0x4d, 0xb1, 0xb0,
	0x15, 0x43, 0x7b, 0xc8, 0x11, 0x94, 0xa7, 0x22, 0x90, 0x3c, 0x69, 0x17, 0x90, 0xda, 0xcc, 0x3c,
	0x77, 0xa3, 0x01, 0x6a, 0x71, 0xf2, 0x0e, 0x76, 0x23, 0x2e, 0x99, 0xcf, 0x24, 0xf3, 0x2c, 0xa5,
	0x88, 0x94, 0x76, 0x86, 0x72, 0x69, 0x23, 0x0c, 0x75, 0x27, 0xca, 0x9a, 0x89, 0xfb, 0x23, 0x07,
	0xb5, 0x73, 0x96, 0x0c, 0xb9, 0x6f, 0x5a, 0x7d, 0x01, 0xf5, 0x21, 0x9a, 0x5e, 0xb6, 0xe3, 0xd6,
	0x4a, 0xc7, 0x9a, 0x41, 0x6b, 0x26, 0x90, 0x62, 0xef, 0xaf, 0xa0, 0x61, 0x79, 0xb6, 0x10, 0xd3,
	0xf6, 0xde, 0x6a, 0xed, 0xc8, 0xb4, 0x4f, 0x98, 0x12, 0x94, 0x6c, 0xbf, 0x75, 0x61, 0x1a, 0x7f,
	0xbc, 0xa9, 0x0b, 0x4c, 0xb2, 0xda, 0xc9, 0x14, 0xca, 0xa6, 0x38, 0xd2, 0x84, 0xc2, 0x88, 0xcf,
	0x55, 0xe9, 0xb9, 0xa3, 0x2a, 0xd5, 0x3f, 0xc9, 0x33, 0xa8, 0x4c, 0xb8, 0x48, 0x82, 0x78, 0xac,
	0xea, 0xca, 0x3d, 0xd0, 0xf4, 0xda, 0xf8, 0xe9, 0x22, 0x80, 0x74, 0xa1, 0x95, 0xdc, 0xf1, 0x7e,
	0x1a, 0x32, 0x19, 0x4c, 0xb8, 0xb7, 0xe0, 0x15, 0x14, 0xaf, 0x48, 0x49, 0x06, 0xb2, 0x4c, 0xf7,
	0x9b, 0x5e, 0x14, 0x2c, 0x62, 0xcd, 0xcb, 0x8f, 0xa0, 0x1a, 0x24, 0x9e, 0xcf, 0x43, 0x2e, 0x39,
	0xbe, 0xed, 0x50, 0x27, 0x48, 0x4e, 0xd1, 0x26, 0x7b, 0x50, 0x9a, 0xb0, 0x30, 0xe5, 0x98, 0xbc,
	0x4e, 0x8d, 0xb1, 0xa9, 0x80, 0xe2, 0xc6, 0x02, 0x6e, 0x60, 0x77, 0x45, 0xa0, 0x35, 0x85, 0x9c,
	0x40, 0x85, 0x8f, 0xa5, 0x5e, 0x32, 0x3b, 0x9a, 0x75, 0x3b, 0x72, 0xa6, 0x22, 0xe6, 0x74, 0x11,
	0xe8, 0x5e, 0x01, 0x2c, 0xe7, 0x4d, 0xfe, 0x07, 0x47, 0x25, 0xf2, 0xf4, 0xec, 0x30, 0x71, 0x9d,
	0x56, 0x94, 0x8d, 0xd0, 0x1f, 0xe8, 0xeb, 0xfa, 0x50, 0xcb, 0xec, 0xc2, 0xb6, 0xac, 0x5b, 0xb5,
	0x7b, 0x02, 0x80, 0x72, 0x19, 0xa6, 0x11, 0xb0, 0x8a, 0x1e, 0xcd, 0x55, 0xaf, 0xb4, 0xd6, 0x2c,
	0xcd, 0xb6, 0xd7, 0xfe, 0x46, 0xa0, 0xd7, 0x59, 0xe5, 0x11, 0x23, 0x04, 0x8a, 0x63, 0x16, 0x71,
	0x2b, 0x3d, 0xfe, 0x5e, 0xce, 0x39, 0x9f, 0x99, 0xb3, 0xfb, 0x16, 0x2a, 0x56, 0x1c, 0xdd, 0x69,
	0x2f, 0x8c, 0xfb, 0x23, 0x6f, 0x9c, 0x46, 0xc8, 0x2c, 0x52, 0x07, 0x1d, 0x9f, 0xd2, 0x88, 0xfc,
	0x07, 0x65, 0x39, 0x43, 0x24, 0x8f, 0x48, 0x49, 0xce, 0x94, 0xdb, 0xfd, 0x9e, 0x87, 0x9d, 0x87,
	0xb7, 0x44, 0xa7, 0x49, 0x24, 0x13, 0xd2, 0x5b, 0xce, 0xde, 0x41, 0xc7, 0x85, 0x5a, 0x80, 0x7d,
	0xdd, 0x9f, 0x8f, 0x50, 0x1e, 0xa1, 0xb2, 0x32, 0x35, 0x70, 0x08, 0x8d, 0x40, 0x0a, 0x8f, 0xcf,
	0x86, 0x2c, 0x4d, 0x24, 0xf7, 0x51, 0x4c, 0x87, 0xd6, 0x95, 0xf3, 0x6c, 0xe1, 0x53, 0xea, 0x54,
	0x05, 0x9b, 0xda, 0xa3, 0x50, 0xc4, 0x19, 0x2f, 0x8f, 0x02, 0x56, 0x80, 0x77, 0xe0, 0xfc, 0x1f,
	0xea, 0xa8, 0x38, 0x73, 0x13, 0x28, 0xb4, 0x30, 0xde, 0x8b, 0xb8, 0x18, 0x85, 0x66, 0x52, 0x4a,
	0xdd, 0x12, 0xb2, 0x0f, 0xd6, 0xb0, 0x2f, 0x31, 0xee, 0x2a, 0x8d, 0x22, 0x26, 0xe6, 0x2a, 0xd5,
	0xbf, 0x62, 0xe9, 0xc5, 0x23, 0x95, 0xbc, 0xaf, 0x03, 0x98, 0x9c, 0xfa, 0xb6, 0xba, 0x2f, 0x01,
	0x96, 0x6c, 0xb5, 0x85, 0x8e, 0xbe, 0xe6, 0xdb, 0x2e, 0xb5, 0x3a, 0xef, 0x18, 0xab, 0xfe, 0x69,
	0xf7, 0x37, 0xbc, 0xab, 0x37, 0x2b, 0x62, 0x33, 0xb5, 0x77, 0x03, 0xc1, 0xcd, 0x1c, 0x1b, 0xb4,
	0xaa, 0x3c, 0xa7, 0xe8, 0xd0, 0x22, 0x6b, 0x38, 0xe4, 0x13, 0x1e, 0xa2, 0x92, 0x0d, 0xea, 0x28,
	0xc7, 0x47, 0x6d, 0xab, 0xdb, 0xdd, 0xbc, 0x07, 0x17, 0xfd, 0xea, 0x63, 0x56, 0x57, 0xe7, 0xca,
	0xc6, 0xd8, 0x46, 0x04, 0x9c, 0xc4, 0x62, 0xd0, 0x19, 0xce, 0xef, 0xb8, 0x30, 0x1f, 0xa6, 0xce,
	0x2d, 0xeb, 0x89, 0xa0, 0x6f, 0x3e, 0x44, 0x49, 0xc7, 0x3a, 0x4d, 0xf9, 0xb6, 0x8d, 0x2f, 0x6f,
	0x06, 0x81, 0x1c, 0xa6, 0xbd, 0x4e, 0x3f, 0x8e, 0xba, 0x19, 0x6a, 0xd7, 0x50, 0x8f, 0x0d, 0xf5,
	0x78, 0x10, 0x77, 0xd7, 0x7d, 0xeb, 0x7a, 0x65, 0xc4, 0x9f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x49, 0x30, 0x65, 0x0a, 0x07, 0x00, 0x00,
}
