// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/client_state_pb2/client_state.proto

package client_state_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import client_list_control_pb2 "github.com/hyperledger/sawtooth-sdk-go/protobuf/client_list_control_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientStateListResponse_Status int32

const (
	ClientStateListResponse_STATUS_UNSET    ClientStateListResponse_Status = 0
	ClientStateListResponse_OK              ClientStateListResponse_Status = 1
	ClientStateListResponse_INTERNAL_ERROR  ClientStateListResponse_Status = 2
	ClientStateListResponse_NOT_READY       ClientStateListResponse_Status = 3
	ClientStateListResponse_NO_ROOT         ClientStateListResponse_Status = 4
	ClientStateListResponse_NO_RESOURCE     ClientStateListResponse_Status = 5
	ClientStateListResponse_INVALID_PAGING  ClientStateListResponse_Status = 6
	ClientStateListResponse_INVALID_SORT    ClientStateListResponse_Status = 7
	ClientStateListResponse_INVALID_ADDRESS ClientStateListResponse_Status = 8
	ClientStateListResponse_INVALID_ROOT    ClientStateListResponse_Status = 9
)

var ClientStateListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
	8: "INVALID_ADDRESS",
	9: "INVALID_ROOT",
}
var ClientStateListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":    0,
	"OK":              1,
	"INTERNAL_ERROR":  2,
	"NOT_READY":       3,
	"NO_ROOT":         4,
	"NO_RESOURCE":     5,
	"INVALID_PAGING":  6,
	"INVALID_SORT":    7,
	"INVALID_ADDRESS": 8,
	"INVALID_ROOT":    9,
}

func (x ClientStateListResponse_Status) String() string {
	return proto.EnumName(ClientStateListResponse_Status_name, int32(x))
}
func (ClientStateListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{1, 0}
}

type ClientStateGetResponse_Status int32

const (
	ClientStateGetResponse_STATUS_UNSET    ClientStateGetResponse_Status = 0
	ClientStateGetResponse_OK              ClientStateGetResponse_Status = 1
	ClientStateGetResponse_INTERNAL_ERROR  ClientStateGetResponse_Status = 2
	ClientStateGetResponse_NOT_READY       ClientStateGetResponse_Status = 3
	ClientStateGetResponse_NO_ROOT         ClientStateGetResponse_Status = 4
	ClientStateGetResponse_NO_RESOURCE     ClientStateGetResponse_Status = 5
	ClientStateGetResponse_INVALID_ADDRESS ClientStateGetResponse_Status = 6
	ClientStateGetResponse_INVALID_ROOT    ClientStateGetResponse_Status = 7
)

var ClientStateGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_ADDRESS",
	7: "INVALID_ROOT",
}
var ClientStateGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":    0,
	"OK":              1,
	"INTERNAL_ERROR":  2,
	"NOT_READY":       3,
	"NO_ROOT":         4,
	"NO_RESOURCE":     5,
	"INVALID_ADDRESS": 6,
	"INVALID_ROOT":    7,
}

func (x ClientStateGetResponse_Status) String() string {
	return proto.EnumName(ClientStateGetResponse_Status_name, int32(x))
}
func (ClientStateGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{3, 0}
}

// A request to list every entry in global state. Defaults to the most current
// tree, but can fetch older state by specifying a state root. Results can be
// further filtered by specifying a subtree with a partial address.
type ClientStateListRequest struct {
	StateRoot            string                                        `protobuf:"bytes,1,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Address              string                                        `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Paging               *client_list_control_pb2.ClientPagingControls `protobuf:"bytes,4,opt,name=paging,proto3" json:"paging,omitempty"`
	Sorting              []*client_list_control_pb2.ClientSortControls `protobuf:"bytes,5,rep,name=sorting,proto3" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *ClientStateListRequest) Reset()         { *m = ClientStateListRequest{} }
func (m *ClientStateListRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStateListRequest) ProtoMessage()    {}
func (*ClientStateListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{0}
}
func (m *ClientStateListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStateListRequest.Unmarshal(m, b)
}
func (m *ClientStateListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStateListRequest.Marshal(b, m, deterministic)
}
func (dst *ClientStateListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStateListRequest.Merge(dst, src)
}
func (m *ClientStateListRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStateListRequest.Size(m)
}
func (m *ClientStateListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStateListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStateListRequest proto.InternalMessageInfo

func (m *ClientStateListRequest) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *ClientStateListRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClientStateListRequest) GetPaging() *client_list_control_pb2.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientStateListRequest) GetSorting() []*client_list_control_pb2.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

type ClientStateListResponse struct {
	Status               ClientStateListResponse_Status                `protobuf:"varint,1,opt,name=status,proto3,enum=ClientStateListResponse_Status" json:"status,omitempty"`
	Entries              []*ClientStateListResponse_Entry              `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	StateRoot            string                                        `protobuf:"bytes,3,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Paging               *client_list_control_pb2.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *ClientStateListResponse) Reset()         { *m = ClientStateListResponse{} }
func (m *ClientStateListResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStateListResponse) ProtoMessage()    {}
func (*ClientStateListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{1}
}
func (m *ClientStateListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStateListResponse.Unmarshal(m, b)
}
func (m *ClientStateListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStateListResponse.Marshal(b, m, deterministic)
}
func (dst *ClientStateListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStateListResponse.Merge(dst, src)
}
func (m *ClientStateListResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStateListResponse.Size(m)
}
func (m *ClientStateListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStateListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStateListResponse proto.InternalMessageInfo

func (m *ClientStateListResponse) GetStatus() ClientStateListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientStateListResponse_STATUS_UNSET
}

func (m *ClientStateListResponse) GetEntries() []*ClientStateListResponse_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ClientStateListResponse) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *ClientStateListResponse) GetPaging() *client_list_control_pb2.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// An entry in the State
type ClientStateListResponse_Entry struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStateListResponse_Entry) Reset()         { *m = ClientStateListResponse_Entry{} }
func (m *ClientStateListResponse_Entry) String() string { return proto.CompactTextString(m) }
func (*ClientStateListResponse_Entry) ProtoMessage()    {}
func (*ClientStateListResponse_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{1, 0}
}
func (m *ClientStateListResponse_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStateListResponse_Entry.Unmarshal(m, b)
}
func (m *ClientStateListResponse_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStateListResponse_Entry.Marshal(b, m, deterministic)
}
func (dst *ClientStateListResponse_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStateListResponse_Entry.Merge(dst, src)
}
func (m *ClientStateListResponse_Entry) XXX_Size() int {
	return xxx_messageInfo_ClientStateListResponse_Entry.Size(m)
}
func (m *ClientStateListResponse_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStateListResponse_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStateListResponse_Entry proto.InternalMessageInfo

func (m *ClientStateListResponse_Entry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClientStateListResponse_Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A request from a client for a particular entry in global state.
// Like State List, it defaults to the newest state, but a state root
// can be used to specify older data. Unlike State List the request must be
// provided with a full address that corresponds to a single entry.
type ClientStateGetRequest struct {
	StateRoot            string   `protobuf:"bytes,1,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStateGetRequest) Reset()         { *m = ClientStateGetRequest{} }
func (m *ClientStateGetRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStateGetRequest) ProtoMessage()    {}
func (*ClientStateGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{2}
}
func (m *ClientStateGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStateGetRequest.Unmarshal(m, b)
}
func (m *ClientStateGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStateGetRequest.Marshal(b, m, deterministic)
}
func (dst *ClientStateGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStateGetRequest.Merge(dst, src)
}
func (m *ClientStateGetRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStateGetRequest.Size(m)
}
func (m *ClientStateGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStateGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStateGetRequest proto.InternalMessageInfo

func (m *ClientStateGetRequest) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *ClientStateGetRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// The response to a State Get Request from the client. Sends back just
// the data stored at the entry, not the address. Also sends back the
// head block id used to facilitate further requests.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the state_root specified was not found
//   * NO_RESOURCE - the address specified doesn't exist
//   * INVALID_ADDRESS - address isn't a valid, i.e. it's a subtree (truncated)
type ClientStateGetResponse struct {
	Status               ClientStateGetResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ClientStateGetResponse_Status" json:"status,omitempty"`
	Value                []byte                        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	StateRoot            string                        `protobuf:"bytes,3,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClientStateGetResponse) Reset()         { *m = ClientStateGetResponse{} }
func (m *ClientStateGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStateGetResponse) ProtoMessage()    {}
func (*ClientStateGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_state_9e777f6161b22e0e, []int{3}
}
func (m *ClientStateGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStateGetResponse.Unmarshal(m, b)
}
func (m *ClientStateGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStateGetResponse.Marshal(b, m, deterministic)
}
func (dst *ClientStateGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStateGetResponse.Merge(dst, src)
}
func (m *ClientStateGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStateGetResponse.Size(m)
}
func (m *ClientStateGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStateGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStateGetResponse proto.InternalMessageInfo

func (m *ClientStateGetResponse) GetStatus() ClientStateGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientStateGetResponse_STATUS_UNSET
}

func (m *ClientStateGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ClientStateGetResponse) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientStateListRequest)(nil), "ClientStateListRequest")
	proto.RegisterType((*ClientStateListResponse)(nil), "ClientStateListResponse")
	proto.RegisterType((*ClientStateListResponse_Entry)(nil), "ClientStateListResponse.Entry")
	proto.RegisterType((*ClientStateGetRequest)(nil), "ClientStateGetRequest")
	proto.RegisterType((*ClientStateGetResponse)(nil), "ClientStateGetResponse")
	proto.RegisterEnum("ClientStateListResponse_Status", ClientStateListResponse_Status_name, ClientStateListResponse_Status_value)
	proto.RegisterEnum("ClientStateGetResponse_Status", ClientStateGetResponse_Status_name, ClientStateGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/hyperledger/sawtooth-sdk-go/protobuf/client_state_pb2/client_state.proto", fileDescriptor_client_state_9e777f6161b22e0e)
}

var fileDescriptor_client_state_9e777f6161b22e0e = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xb2, 0x4d, 0xec, 0xe9, 0xba, 0x3b, 0xcc, 0x5a, 0x0d, 0x0b, 0x6a, 0xc9, 0x55,
	0x41, 0x37, 0x42, 0xc5, 0x3f, 0x78, 0x17, 0xdb, 0x50, 0x8a, 0x25, 0x29, 0x33, 0xa9, 0xa0, 0x37,
	0x21, 0xdd, 0x8e, 0x6b, 0xb0, 0x64, 0x6a, 0x66, 0xaa, 0xf8, 0x0a, 0x5e, 0xfa, 0x22, 0xfa, 0x10,
	0x3e, 0x98, 0x64, 0x92, 0x48, 0x9a, 0x75, 0xf5, 0x42, 0xf0, 0x2e, 0xf3, 0xcd, 0x6f, 0xbe, 0x73,
	0xf2, 0xcd, 0x61, 0xe0, 0xc1, 0x36, 0xe7, 0x92, 0xaf, 0x76, 0x6f, 0x1f, 0x9e, 0x6f, 0x52, 0x96,
	0xc9, 0x58, 0xc8, 0x44, 0xb2, 0x78, 0xbb, 0x1a, 0xed, 0x09, 0xae, 0xc2, 0x4e, 0x9f, 0xb7, 0xe9,
	0x4d, 0x2a, 0x64, 0x7c, 0xce, 0x33, 0x99, 0xf3, 0x4d, 0xf3, 0x50, 0x53, 0x2f, 0xcf, 0x3a, 0xdf,
	0x34, 0xb8, 0x35, 0x56, 0xbb, 0xb4, 0x70, 0x9c, 0xa7, 0x42, 0x12, 0xf6, 0x61, 0xc7, 0x84, 0xc4,
	0x77, 0x00, 0xca, 0xb2, 0x39, 0xe7, 0xd2, 0xd6, 0x06, 0xda, 0xb0, 0x4b, 0xba, 0x4a, 0x21, 0x9c,
	0x4b, 0x6c, 0x83, 0x95, 0xac, 0xd7, 0x39, 0x13, 0xc2, 0x36, 0xd4, 0x5e, 0xbd, 0xc4, 0x67, 0x60,
	0x6e, 0x93, 0x8b, 0x34, 0xbb, 0xb0, 0x0f, 0x06, 0xda, 0xb0, 0x37, 0xea, 0xbb, 0x65, 0x85, 0x85,
	0x12, 0xc7, 0x65, 0x7d, 0x41, 0x2a, 0x08, 0x9f, 0x81, 0x25, 0x78, 0x2e, 0x0b, 0xbe, 0x33, 0x30,
	0x86, 0xbd, 0xd1, 0x49, 0xc5, 0x53, 0x9e, 0xcb, 0x5f, 0x74, 0xcd, 0x38, 0x3f, 0x0c, 0xb8, 0x7d,
	0xa9, 0x63, 0xb1, 0xe5, 0x99, 0x60, 0xf8, 0x29, 0x98, 0x45, 0x83, 0x3b, 0xa1, 0xda, 0x3d, 0x1a,
	0xdd, 0x73, 0xaf, 0x20, 0x5d, 0xaa, 0x30, 0x52, 0xe1, 0xf8, 0x19, 0x58, 0x2c, 0x93, 0x79, 0xca,
	0x84, 0xad, 0xab, 0x1e, 0xee, 0x5e, 0x79, 0xd2, 0xcf, 0x64, 0xfe, 0x99, 0xd4, 0x78, 0x2b, 0x25,
	0xa3, 0x9d, 0xd2, 0x9f, 0xb3, 0xa8, 0x4d, 0xeb, 0x2c, 0x4e, 0x1f, 0x43, 0x47, 0xf9, 0x37, 0xd3,
	0xd5, 0xf6, 0xd3, 0xc5, 0x70, 0xb0, 0x4e, 0x64, 0x62, 0xeb, 0x03, 0x6d, 0x78, 0x48, 0xd4, 0xb7,
	0xf3, 0x5d, 0x03, 0xb3, 0xfc, 0x23, 0x8c, 0xe0, 0x90, 0x46, 0x5e, 0xb4, 0xa4, 0xf1, 0x32, 0xa0,
	0x7e, 0x84, 0xae, 0x61, 0x13, 0xf4, 0xf0, 0x25, 0xd2, 0x30, 0x86, 0xa3, 0x59, 0x10, 0xf9, 0x24,
	0xf0, 0xe6, 0xb1, 0x4f, 0x48, 0x48, 0x90, 0x8e, 0x6f, 0x40, 0x37, 0x08, 0xa3, 0x98, 0xf8, 0xde,
	0xe4, 0x35, 0x32, 0x70, 0x0f, 0xac, 0x20, 0x8c, 0x49, 0x18, 0x46, 0xe8, 0x00, 0x1f, 0x43, 0xaf,
	0x58, 0xf8, 0x34, 0x5c, 0x92, 0xb1, 0x8f, 0x3a, 0xa5, 0xc1, 0x2b, 0x6f, 0x3e, 0x9b, 0xc4, 0x0b,
	0x6f, 0x3a, 0x0b, 0xa6, 0xc8, 0x2c, 0xca, 0xd5, 0x1a, 0x0d, 0x49, 0x84, 0x2c, 0x7c, 0x02, 0xc7,
	0xb5, 0xe2, 0x4d, 0x26, 0xc4, 0xa7, 0x14, 0x5d, 0x6f, 0x62, 0xca, 0xbd, 0xeb, 0x2c, 0xa0, 0xdf,
	0x48, 0x78, 0xca, 0xfe, 0x79, 0xec, 0x9c, 0xaf, 0xfa, 0xde, 0x28, 0x2b, 0xcb, 0x6a, 0x2e, 0x9e,
	0xb4, 0xe6, 0x62, 0xef, 0x76, 0x1b, 0x60, 0x7b, 0x2c, 0x6e, 0x42, 0xe7, 0x63, 0xb2, 0xd9, 0xb1,
	0x2a, 0xec, 0x72, 0xf1, 0x97, 0x2b, 0x77, 0xbe, 0xfc, 0x97, 0xcb, 0xf8, 0x4d, 0xcc, 0xe6, 0xa5,
	0x98, 0xad, 0x17, 0xf7, 0xa1, 0x2f, 0x92, 0x4f, 0x92, 0x73, 0xf9, 0xce, 0x15, 0xeb, 0xf7, 0x6e,
	0xfd, 0x54, 0x2c, 0xb4, 0x37, 0xa8, 0xfd, 0xb6, 0xac, 0x4c, 0xb5, 0xfb, 0xe8, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x38, 0x15, 0x51, 0x2a, 0x7f, 0x04, 0x00, 0x00,
}