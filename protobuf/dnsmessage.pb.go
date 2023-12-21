// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dnsmessage.proto

package dnsmessage

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

type PBDNSMessage_Type int32

const (
	PBDNSMessage_DNSQueryType            PBDNSMessage_Type = 1
	PBDNSMessage_DNSResponseType         PBDNSMessage_Type = 2
	PBDNSMessage_DNSOutgoingQueryType    PBDNSMessage_Type = 3
	PBDNSMessage_DNSIncomingResponseType PBDNSMessage_Type = 4
)

var PBDNSMessage_Type_name = map[int32]string{
	1: "DNSQueryType",
	2: "DNSResponseType",
	3: "DNSOutgoingQueryType",
	4: "DNSIncomingResponseType",
}

var PBDNSMessage_Type_value = map[string]int32{
	"DNSQueryType":            1,
	"DNSResponseType":         2,
	"DNSOutgoingQueryType":    3,
	"DNSIncomingResponseType": 4,
}

func (x PBDNSMessage_Type) Enum() *PBDNSMessage_Type {
	p := new(PBDNSMessage_Type)
	*p = x
	return p
}

func (x PBDNSMessage_Type) String() string {
	return proto.EnumName(PBDNSMessage_Type_name, int32(x))
}

func (x *PBDNSMessage_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBDNSMessage_Type_value, data, "PBDNSMessage_Type")
	if err != nil {
		return err
	}
	*x = PBDNSMessage_Type(value)
	return nil
}

func (PBDNSMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 0}
}

type PBDNSMessage_SocketFamily int32

const (
	PBDNSMessage_INET  PBDNSMessage_SocketFamily = 1
	PBDNSMessage_INET6 PBDNSMessage_SocketFamily = 2
)

var PBDNSMessage_SocketFamily_name = map[int32]string{
	1: "INET",
	2: "INET6",
}

var PBDNSMessage_SocketFamily_value = map[string]int32{
	"INET":  1,
	"INET6": 2,
}

func (x PBDNSMessage_SocketFamily) Enum() *PBDNSMessage_SocketFamily {
	p := new(PBDNSMessage_SocketFamily)
	*p = x
	return p
}

func (x PBDNSMessage_SocketFamily) String() string {
	return proto.EnumName(PBDNSMessage_SocketFamily_name, int32(x))
}

func (x *PBDNSMessage_SocketFamily) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBDNSMessage_SocketFamily_value, data, "PBDNSMessage_SocketFamily")
	if err != nil {
		return err
	}
	*x = PBDNSMessage_SocketFamily(value)
	return nil
}

func (PBDNSMessage_SocketFamily) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 1}
}

type PBDNSMessage_SocketProtocol int32

const (
	PBDNSMessage_UDP PBDNSMessage_SocketProtocol = 1
	PBDNSMessage_TCP PBDNSMessage_SocketProtocol = 2
)

var PBDNSMessage_SocketProtocol_name = map[int32]string{
	1: "UDP",
	2: "TCP",
}

var PBDNSMessage_SocketProtocol_value = map[string]int32{
	"UDP": 1,
	"TCP": 2,
}

func (x PBDNSMessage_SocketProtocol) Enum() *PBDNSMessage_SocketProtocol {
	p := new(PBDNSMessage_SocketProtocol)
	*p = x
	return p
}

func (x PBDNSMessage_SocketProtocol) String() string {
	return proto.EnumName(PBDNSMessage_SocketProtocol_name, int32(x))
}

func (x *PBDNSMessage_SocketProtocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBDNSMessage_SocketProtocol_value, data, "PBDNSMessage_SocketProtocol")
	if err != nil {
		return err
	}
	*x = PBDNSMessage_SocketProtocol(value)
	return nil
}

func (PBDNSMessage_SocketProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 2}
}

type PBDNSMessage_PolicyType int32

const (
	PBDNSMessage_UNKNOWN    PBDNSMessage_PolicyType = 1
	PBDNSMessage_QNAME      PBDNSMessage_PolicyType = 2
	PBDNSMessage_CLIENTIP   PBDNSMessage_PolicyType = 3
	PBDNSMessage_RESPONSEIP PBDNSMessage_PolicyType = 4
	PBDNSMessage_NSDNAME    PBDNSMessage_PolicyType = 5
	PBDNSMessage_NSIP       PBDNSMessage_PolicyType = 6
)

var PBDNSMessage_PolicyType_name = map[int32]string{
	1: "UNKNOWN",
	2: "QNAME",
	3: "CLIENTIP",
	4: "RESPONSEIP",
	5: "NSDNAME",
	6: "NSIP",
}

var PBDNSMessage_PolicyType_value = map[string]int32{
	"UNKNOWN":    1,
	"QNAME":      2,
	"CLIENTIP":   3,
	"RESPONSEIP": 4,
	"NSDNAME":    5,
	"NSIP":       6,
}

func (x PBDNSMessage_PolicyType) Enum() *PBDNSMessage_PolicyType {
	p := new(PBDNSMessage_PolicyType)
	*p = x
	return p
}

func (x PBDNSMessage_PolicyType) String() string {
	return proto.EnumName(PBDNSMessage_PolicyType_name, int32(x))
}

func (x *PBDNSMessage_PolicyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBDNSMessage_PolicyType_value, data, "PBDNSMessage_PolicyType")
	if err != nil {
		return err
	}
	*x = PBDNSMessage_PolicyType(value)
	return nil
}

func (PBDNSMessage_PolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 3}
}

type PBDNSMessage struct {
	Type                    *PBDNSMessage_Type           `protobuf:"varint,1,req,name=type,enum=PBDNSMessage_Type" json:"type,omitempty"`
	MessageId               []byte                       `protobuf:"bytes,2,opt,name=messageId" json:"messageId,omitempty"`
	ServerIdentity          []byte                       `protobuf:"bytes,3,opt,name=serverIdentity" json:"serverIdentity,omitempty"`
	SocketFamily            *PBDNSMessage_SocketFamily   `protobuf:"varint,4,opt,name=socketFamily,enum=PBDNSMessage_SocketFamily" json:"socketFamily,omitempty"`
	SocketProtocol          *PBDNSMessage_SocketProtocol `protobuf:"varint,5,opt,name=socketProtocol,enum=PBDNSMessage_SocketProtocol" json:"socketProtocol,omitempty"`
	From                    []byte                       `protobuf:"bytes,6,opt,name=from" json:"from,omitempty"`
	To                      []byte                       `protobuf:"bytes,7,opt,name=to" json:"to,omitempty"`
	InBytes                 *uint64                      `protobuf:"varint,8,opt,name=inBytes" json:"inBytes,omitempty"`
	TimeSec                 *uint32                      `protobuf:"varint,9,opt,name=timeSec" json:"timeSec,omitempty"`
	TimeUsec                *uint32                      `protobuf:"varint,10,opt,name=timeUsec" json:"timeUsec,omitempty"`
	Id                      *uint32                      `protobuf:"varint,11,opt,name=id" json:"id,omitempty"`
	Question                *PBDNSMessage_DNSQuestion    `protobuf:"bytes,12,opt,name=question" json:"question,omitempty"`
	Response                *PBDNSMessage_DNSResponse    `protobuf:"bytes,13,opt,name=response" json:"response,omitempty"`
	OriginalRequestorSubnet []byte                       `protobuf:"bytes,14,opt,name=originalRequestorSubnet" json:"originalRequestorSubnet,omitempty"`
	RequestorId             *string                      `protobuf:"bytes,15,opt,name=requestorId" json:"requestorId,omitempty"`
	InitialRequestId        []byte                       `protobuf:"bytes,16,opt,name=initialRequestId" json:"initialRequestId,omitempty"`
	DeviceId                []byte                       `protobuf:"bytes,17,opt,name=deviceId" json:"deviceId,omitempty"`
	NewlyObservedDomain     *bool                        `protobuf:"varint,18,opt,name=newlyObservedDomain" json:"newlyObservedDomain,omitempty"`
	DeviceName              *string                      `protobuf:"bytes,19,opt,name=deviceName" json:"deviceName,omitempty"`
	FromPort                *uint32                      `protobuf:"varint,20,opt,name=fromPort" json:"fromPort,omitempty"`
	ToPort                  *uint32                      `protobuf:"varint,21,opt,name=toPort" json:"toPort,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                     `json:"-"`
	XXX_unrecognized        []byte                       `json:"-"`
	XXX_sizecache           int32                        `json:"-"`
}

func (m *PBDNSMessage) Reset()         { *m = PBDNSMessage{} }
func (m *PBDNSMessage) String() string { return proto.CompactTextString(m) }
func (*PBDNSMessage) ProtoMessage()    {}
func (*PBDNSMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0}
}

func (m *PBDNSMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBDNSMessage.Unmarshal(m, b)
}
func (m *PBDNSMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBDNSMessage.Marshal(b, m, deterministic)
}
func (m *PBDNSMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBDNSMessage.Merge(m, src)
}
func (m *PBDNSMessage) XXX_Size() int {
	return xxx_messageInfo_PBDNSMessage.Size(m)
}
func (m *PBDNSMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PBDNSMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PBDNSMessage proto.InternalMessageInfo

func (m *PBDNSMessage) GetType() PBDNSMessage_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PBDNSMessage_DNSQueryType
}

func (m *PBDNSMessage) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *PBDNSMessage) GetServerIdentity() []byte {
	if m != nil {
		return m.ServerIdentity
	}
	return nil
}

func (m *PBDNSMessage) GetSocketFamily() PBDNSMessage_SocketFamily {
	if m != nil && m.SocketFamily != nil {
		return *m.SocketFamily
	}
	return PBDNSMessage_INET
}

func (m *PBDNSMessage) GetSocketProtocol() PBDNSMessage_SocketProtocol {
	if m != nil && m.SocketProtocol != nil {
		return *m.SocketProtocol
	}
	return PBDNSMessage_UDP
}

func (m *PBDNSMessage) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PBDNSMessage) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PBDNSMessage) GetInBytes() uint64 {
	if m != nil && m.InBytes != nil {
		return *m.InBytes
	}
	return 0
}

func (m *PBDNSMessage) GetTimeSec() uint32 {
	if m != nil && m.TimeSec != nil {
		return *m.TimeSec
	}
	return 0
}

func (m *PBDNSMessage) GetTimeUsec() uint32 {
	if m != nil && m.TimeUsec != nil {
		return *m.TimeUsec
	}
	return 0
}

func (m *PBDNSMessage) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PBDNSMessage) GetQuestion() *PBDNSMessage_DNSQuestion {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *PBDNSMessage) GetResponse() *PBDNSMessage_DNSResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PBDNSMessage) GetOriginalRequestorSubnet() []byte {
	if m != nil {
		return m.OriginalRequestorSubnet
	}
	return nil
}

func (m *PBDNSMessage) GetRequestorId() string {
	if m != nil && m.RequestorId != nil {
		return *m.RequestorId
	}
	return ""
}

func (m *PBDNSMessage) GetInitialRequestId() []byte {
	if m != nil {
		return m.InitialRequestId
	}
	return nil
}

func (m *PBDNSMessage) GetDeviceId() []byte {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *PBDNSMessage) GetNewlyObservedDomain() bool {
	if m != nil && m.NewlyObservedDomain != nil {
		return *m.NewlyObservedDomain
	}
	return false
}

func (m *PBDNSMessage) GetDeviceName() string {
	if m != nil && m.DeviceName != nil {
		return *m.DeviceName
	}
	return ""
}

func (m *PBDNSMessage) GetFromPort() uint32 {
	if m != nil && m.FromPort != nil {
		return *m.FromPort
	}
	return 0
}

func (m *PBDNSMessage) GetToPort() uint32 {
	if m != nil && m.ToPort != nil {
		return *m.ToPort
	}
	return 0
}

type PBDNSMessage_DNSQuestion struct {
	QName                *string  `protobuf:"bytes,1,opt,name=qName" json:"qName,omitempty"`
	QType                *uint32  `protobuf:"varint,2,opt,name=qType" json:"qType,omitempty"`
	QClass               *uint32  `protobuf:"varint,3,opt,name=qClass" json:"qClass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBDNSMessage_DNSQuestion) Reset()         { *m = PBDNSMessage_DNSQuestion{} }
func (m *PBDNSMessage_DNSQuestion) String() string { return proto.CompactTextString(m) }
func (*PBDNSMessage_DNSQuestion) ProtoMessage()    {}
func (*PBDNSMessage_DNSQuestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 0}
}

func (m *PBDNSMessage_DNSQuestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBDNSMessage_DNSQuestion.Unmarshal(m, b)
}
func (m *PBDNSMessage_DNSQuestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBDNSMessage_DNSQuestion.Marshal(b, m, deterministic)
}
func (m *PBDNSMessage_DNSQuestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBDNSMessage_DNSQuestion.Merge(m, src)
}
func (m *PBDNSMessage_DNSQuestion) XXX_Size() int {
	return xxx_messageInfo_PBDNSMessage_DNSQuestion.Size(m)
}
func (m *PBDNSMessage_DNSQuestion) XXX_DiscardUnknown() {
	xxx_messageInfo_PBDNSMessage_DNSQuestion.DiscardUnknown(m)
}

var xxx_messageInfo_PBDNSMessage_DNSQuestion proto.InternalMessageInfo

func (m *PBDNSMessage_DNSQuestion) GetQName() string {
	if m != nil && m.QName != nil {
		return *m.QName
	}
	return ""
}

func (m *PBDNSMessage_DNSQuestion) GetQType() uint32 {
	if m != nil && m.QType != nil {
		return *m.QType
	}
	return 0
}

func (m *PBDNSMessage_DNSQuestion) GetQClass() uint32 {
	if m != nil && m.QClass != nil {
		return *m.QClass
	}
	return 0
}

type PBDNSMessage_DNSResponse struct {
	Rcode                *uint32                           `protobuf:"varint,1,opt,name=rcode" json:"rcode,omitempty"`
	Rrs                  []*PBDNSMessage_DNSResponse_DNSRR `protobuf:"bytes,2,rep,name=rrs" json:"rrs,omitempty"`
	AppliedPolicy        *string                           `protobuf:"bytes,3,opt,name=appliedPolicy" json:"appliedPolicy,omitempty"`
	Tags                 []string                          `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	QueryTimeSec         *uint32                           `protobuf:"varint,5,opt,name=queryTimeSec" json:"queryTimeSec,omitempty"`
	QueryTimeUsec        *uint32                           `protobuf:"varint,6,opt,name=queryTimeUsec" json:"queryTimeUsec,omitempty"`
	AppliedPolicyType    *PBDNSMessage_PolicyType          `protobuf:"varint,7,opt,name=appliedPolicyType,enum=PBDNSMessage_PolicyType" json:"appliedPolicyType,omitempty"`
	AppliedPolicyTrigger *string                           `protobuf:"bytes,8,opt,name=appliedPolicyTrigger" json:"appliedPolicyTrigger,omitempty"`
	AppliedPolicyHit     *string                           `protobuf:"bytes,9,opt,name=appliedPolicyHit" json:"appliedPolicyHit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *PBDNSMessage_DNSResponse) Reset()         { *m = PBDNSMessage_DNSResponse{} }
func (m *PBDNSMessage_DNSResponse) String() string { return proto.CompactTextString(m) }
func (*PBDNSMessage_DNSResponse) ProtoMessage()    {}
func (*PBDNSMessage_DNSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 1}
}

func (m *PBDNSMessage_DNSResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBDNSMessage_DNSResponse.Unmarshal(m, b)
}
func (m *PBDNSMessage_DNSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBDNSMessage_DNSResponse.Marshal(b, m, deterministic)
}
func (m *PBDNSMessage_DNSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBDNSMessage_DNSResponse.Merge(m, src)
}
func (m *PBDNSMessage_DNSResponse) XXX_Size() int {
	return xxx_messageInfo_PBDNSMessage_DNSResponse.Size(m)
}
func (m *PBDNSMessage_DNSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PBDNSMessage_DNSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PBDNSMessage_DNSResponse proto.InternalMessageInfo

func (m *PBDNSMessage_DNSResponse) GetRcode() uint32 {
	if m != nil && m.Rcode != nil {
		return *m.Rcode
	}
	return 0
}

func (m *PBDNSMessage_DNSResponse) GetRrs() []*PBDNSMessage_DNSResponse_DNSRR {
	if m != nil {
		return m.Rrs
	}
	return nil
}

func (m *PBDNSMessage_DNSResponse) GetAppliedPolicy() string {
	if m != nil && m.AppliedPolicy != nil {
		return *m.AppliedPolicy
	}
	return ""
}

func (m *PBDNSMessage_DNSResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PBDNSMessage_DNSResponse) GetQueryTimeSec() uint32 {
	if m != nil && m.QueryTimeSec != nil {
		return *m.QueryTimeSec
	}
	return 0
}

func (m *PBDNSMessage_DNSResponse) GetQueryTimeUsec() uint32 {
	if m != nil && m.QueryTimeUsec != nil {
		return *m.QueryTimeUsec
	}
	return 0
}

func (m *PBDNSMessage_DNSResponse) GetAppliedPolicyType() PBDNSMessage_PolicyType {
	if m != nil && m.AppliedPolicyType != nil {
		return *m.AppliedPolicyType
	}
	return PBDNSMessage_UNKNOWN
}

func (m *PBDNSMessage_DNSResponse) GetAppliedPolicyTrigger() string {
	if m != nil && m.AppliedPolicyTrigger != nil {
		return *m.AppliedPolicyTrigger
	}
	return ""
}

func (m *PBDNSMessage_DNSResponse) GetAppliedPolicyHit() string {
	if m != nil && m.AppliedPolicyHit != nil {
		return *m.AppliedPolicyHit
	}
	return ""
}

// See exportTypes in https://docs.powerdns.com/recursor/lua-config/protobuf.html#protobufServer
// for the list of supported resource record types.
type PBDNSMessage_DNSResponse_DNSRR struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type                 *uint32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Class                *uint32  `protobuf:"varint,3,opt,name=class" json:"class,omitempty"`
	Ttl                  *uint32  `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
	Rdata                []byte   `protobuf:"bytes,5,opt,name=rdata" json:"rdata,omitempty"`
	Udr                  *bool    `protobuf:"varint,6,opt,name=udr" json:"udr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBDNSMessage_DNSResponse_DNSRR) Reset()         { *m = PBDNSMessage_DNSResponse_DNSRR{} }
func (m *PBDNSMessage_DNSResponse_DNSRR) String() string { return proto.CompactTextString(m) }
func (*PBDNSMessage_DNSResponse_DNSRR) ProtoMessage()    {}
func (*PBDNSMessage_DNSResponse_DNSRR) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{0, 1, 0}
}

func (m *PBDNSMessage_DNSResponse_DNSRR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBDNSMessage_DNSResponse_DNSRR.Unmarshal(m, b)
}
func (m *PBDNSMessage_DNSResponse_DNSRR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBDNSMessage_DNSResponse_DNSRR.Marshal(b, m, deterministic)
}
func (m *PBDNSMessage_DNSResponse_DNSRR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBDNSMessage_DNSResponse_DNSRR.Merge(m, src)
}
func (m *PBDNSMessage_DNSResponse_DNSRR) XXX_Size() int {
	return xxx_messageInfo_PBDNSMessage_DNSResponse_DNSRR.Size(m)
}
func (m *PBDNSMessage_DNSResponse_DNSRR) XXX_DiscardUnknown() {
	xxx_messageInfo_PBDNSMessage_DNSResponse_DNSRR.DiscardUnknown(m)
}

var xxx_messageInfo_PBDNSMessage_DNSResponse_DNSRR proto.InternalMessageInfo

func (m *PBDNSMessage_DNSResponse_DNSRR) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PBDNSMessage_DNSResponse_DNSRR) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PBDNSMessage_DNSResponse_DNSRR) GetClass() uint32 {
	if m != nil && m.Class != nil {
		return *m.Class
	}
	return 0
}

func (m *PBDNSMessage_DNSResponse_DNSRR) GetTtl() uint32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *PBDNSMessage_DNSResponse_DNSRR) GetRdata() []byte {
	if m != nil {
		return m.Rdata
	}
	return nil
}

func (m *PBDNSMessage_DNSResponse_DNSRR) GetUdr() bool {
	if m != nil && m.Udr != nil {
		return *m.Udr
	}
	return false
}

type PBDNSMessageList struct {
	Msg                  []*PBDNSMessage `protobuf:"bytes,1,rep,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PBDNSMessageList) Reset()         { *m = PBDNSMessageList{} }
func (m *PBDNSMessageList) String() string { return proto.CompactTextString(m) }
func (*PBDNSMessageList) ProtoMessage()    {}
func (*PBDNSMessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3136ceafbfed9e7, []int{1}
}

func (m *PBDNSMessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBDNSMessageList.Unmarshal(m, b)
}
func (m *PBDNSMessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBDNSMessageList.Marshal(b, m, deterministic)
}
func (m *PBDNSMessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBDNSMessageList.Merge(m, src)
}
func (m *PBDNSMessageList) XXX_Size() int {
	return xxx_messageInfo_PBDNSMessageList.Size(m)
}
func (m *PBDNSMessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_PBDNSMessageList.DiscardUnknown(m)
}

var xxx_messageInfo_PBDNSMessageList proto.InternalMessageInfo

func (m *PBDNSMessageList) GetMsg() []*PBDNSMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterEnum("PBDNSMessage_Type", PBDNSMessage_Type_name, PBDNSMessage_Type_value)
	proto.RegisterEnum("PBDNSMessage_SocketFamily", PBDNSMessage_SocketFamily_name, PBDNSMessage_SocketFamily_value)
	proto.RegisterEnum("PBDNSMessage_SocketProtocol", PBDNSMessage_SocketProtocol_name, PBDNSMessage_SocketProtocol_value)
	proto.RegisterEnum("PBDNSMessage_PolicyType", PBDNSMessage_PolicyType_name, PBDNSMessage_PolicyType_value)
	proto.RegisterType((*PBDNSMessage)(nil), "PBDNSMessage")
	proto.RegisterType((*PBDNSMessage_DNSQuestion)(nil), "PBDNSMessage.DNSQuestion")
	proto.RegisterType((*PBDNSMessage_DNSResponse)(nil), "PBDNSMessage.DNSResponse")
	proto.RegisterType((*PBDNSMessage_DNSResponse_DNSRR)(nil), "PBDNSMessage.DNSResponse.DNSRR")
	proto.RegisterType((*PBDNSMessageList)(nil), "PBDNSMessageList")
}

func init() {
	proto.RegisterFile("dnsmessage.proto", fileDescriptor_c3136ceafbfed9e7)
}

var fileDescriptor_c3136ceafbfed9e7 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x8f, 0xdb, 0xc4,
	0x17, 0x95, 0x3f, 0xb2, 0x49, 0x6e, 0xec, 0xd4, 0x9d, 0xcd, 0xef, 0xd7, 0x21, 0x54, 0xd4, 0x0a,
	0xa8, 0xb2, 0x78, 0x58, 0x41, 0x10, 0x88, 0x27, 0x24, 0xba, 0x49, 0x85, 0x45, 0xeb, 0xf5, 0x8e,
	0xb3, 0x42, 0x3c, 0xba, 0xf6, 0x60, 0x8d, 0x48, 0x3c, 0x59, 0x7b, 0x52, 0x94, 0x27, 0x84, 0xf8,
	0xc7, 0xd1, 0x5c, 0xe7, 0xc3, 0xee, 0xee, 0xbe, 0xdd, 0x73, 0xee, 0xb9, 0xc7, 0x9e, 0x7b, 0xef,
	0x0c, 0x78, 0x79, 0x59, 0x6f, 0x78, 0x5d, 0xa7, 0x05, 0xbf, 0xda, 0x56, 0x52, 0xc9, 0xd9, 0x3f,
	0x2e, 0x38, 0xf1, 0x9b, 0x45, 0x94, 0xbc, 0x6f, 0x68, 0xf2, 0x1a, 0x6c, 0xb5, 0xdf, 0x72, 0x6a,
	0xf8, 0x66, 0x30, 0x9e, 0x93, 0xab, 0x76, 0xf2, 0x6a, 0xb5, 0xdf, 0x72, 0x86, 0x79, 0xf2, 0x12,
	0x86, 0x07, 0xa7, 0x30, 0xa7, 0xa6, 0x6f, 0x04, 0x0e, 0x3b, 0x13, 0xe4, 0x35, 0x8c, 0x6b, 0x5e,
	0x7d, 0xe4, 0x55, 0x98, 0xf3, 0x52, 0x09, 0xb5, 0xa7, 0x16, 0x4a, 0x3e, 0x61, 0xc9, 0x4f, 0xe0,
	0xd4, 0x32, 0xfb, 0x93, 0xab, 0xb7, 0xe9, 0x46, 0xac, 0xf7, 0xd4, 0xf6, 0x8d, 0x60, 0x3c, 0x9f,
	0x76, 0xbf, 0x9a, 0xb4, 0x14, 0xac, 0xa3, 0x27, 0x0b, 0x18, 0x37, 0x38, 0xd6, 0xa7, 0xc9, 0xe4,
	0x9a, 0xf6, 0xd0, 0xe1, 0xe5, 0x63, 0x0e, 0x47, 0x0d, 0xfb, 0xa4, 0x86, 0x10, 0xb0, 0xff, 0xa8,
	0xe4, 0x86, 0x5e, 0xe0, 0x3f, 0x62, 0x4c, 0xc6, 0x60, 0x2a, 0x49, 0xfb, 0xc8, 0x98, 0x4a, 0x12,
	0x0a, 0x7d, 0x51, 0xbe, 0xd9, 0x2b, 0x5e, 0xd3, 0x81, 0x6f, 0x04, 0x36, 0x3b, 0x42, 0x9d, 0x51,
	0x62, 0xc3, 0x13, 0x9e, 0xd1, 0xa1, 0x6f, 0x04, 0x2e, 0x3b, 0x42, 0x32, 0x85, 0x81, 0x0e, 0xef,
	0x6a, 0x9e, 0x51, 0xc0, 0xd4, 0x09, 0x6b, 0x7f, 0x91, 0xd3, 0x11, 0xb2, 0xa6, 0xc8, 0xc9, 0xf7,
	0x30, 0xb8, 0xdf, 0xf1, 0x5a, 0x09, 0x59, 0x52, 0xc7, 0x37, 0x82, 0xd1, 0xfc, 0xb3, 0xee, 0x19,
	0x16, 0x51, 0x72, 0x7b, 0x10, 0xb0, 0x93, 0x54, 0x97, 0x55, 0xbc, 0xde, 0xca, 0xb2, 0xe6, 0xd4,
	0x7d, 0xa2, 0x8c, 0x1d, 0x04, 0xec, 0x24, 0x25, 0x3f, 0xc2, 0x0b, 0x59, 0x89, 0x42, 0x94, 0xe9,
	0x9a, 0x71, 0x34, 0x93, 0x55, 0xb2, 0xfb, 0x50, 0x72, 0x45, 0xc7, 0x78, 0xe4, 0xa7, 0xd2, 0xc4,
	0x87, 0x51, 0x75, 0xa4, 0xc2, 0x9c, 0x3e, 0xf3, 0x8d, 0x60, 0xc8, 0xda, 0x14, 0xf9, 0x1a, 0x3c,
	0x51, 0x0a, 0x25, 0x4e, 0xb5, 0x61, 0x4e, 0x3d, 0x34, 0x7d, 0xc0, 0xeb, 0x0e, 0xe5, 0xfc, 0xa3,
	0xc8, 0xf4, 0x12, 0x3d, 0x47, 0xcd, 0x09, 0x93, 0x6f, 0xe0, 0xb2, 0xe4, 0x7f, 0xad, 0xf7, 0x37,
	0x1f, 0x70, 0x69, 0xf2, 0x85, 0xdc, 0xa4, 0xa2, 0xa4, 0xc4, 0x37, 0x82, 0x01, 0x7b, 0x2c, 0x45,
	0xbe, 0x00, 0x68, 0xaa, 0xa3, 0x74, 0xc3, 0xe9, 0x25, 0xfe, 0x5a, 0x8b, 0xd1, 0x5f, 0xd3, 0xb3,
	0x8d, 0x65, 0xa5, 0xe8, 0xa4, 0x99, 0xc7, 0x11, 0x93, 0xff, 0xc3, 0x85, 0x92, 0x98, 0xf9, 0x1f,
	0x66, 0x0e, 0x68, 0x7a, 0x0b, 0xa3, 0x56, 0xe7, 0xc9, 0x04, 0x7a, 0xf7, 0xe8, 0x6e, 0xa0, 0x7b,
	0x03, 0x90, 0xd5, 0x77, 0x03, 0x2f, 0x82, 0xcb, 0x1a, 0xa0, 0x2d, 0xef, 0xaf, 0xd7, 0x69, 0x5d,
	0xe3, 0xf2, 0xbb, 0xec, 0x80, 0xa6, 0xff, 0xda, 0xe8, 0x79, 0x1c, 0x8b, 0xae, 0xae, 0x32, 0x99,
	0x37, 0x9e, 0x2e, 0x6b, 0x00, 0xf9, 0x16, 0xac, 0xaa, 0xaa, 0xa9, 0xe9, 0x5b, 0xc1, 0x68, 0xfe,
	0xea, 0xc9, 0xa1, 0x62, 0xcc, 0x98, 0xd6, 0x92, 0xaf, 0xc0, 0x4d, 0xb7, 0xdb, 0xb5, 0xe0, 0x79,
	0x2c, 0xd7, 0x22, 0x6b, 0x2e, 0xdd, 0x90, 0x75, 0x49, 0xbd, 0xed, 0x2a, 0x2d, 0x6a, 0x6a, 0xfb,
	0x56, 0x30, 0x64, 0x18, 0x93, 0x19, 0x38, 0xf7, 0x3b, 0x5e, 0xed, 0x57, 0x87, 0x45, 0xee, 0xe1,
	0x9f, 0x74, 0x38, 0xed, 0x7e, 0xc2, 0xb8, 0xd2, 0x17, 0x28, 0xea, 0x92, 0xe4, 0x2d, 0x3c, 0xef,
	0x7c, 0x0e, 0xdb, 0xd2, 0xc7, 0x4b, 0x49, 0xbb, 0x87, 0x38, 0xe7, 0xd9, 0xc3, 0x12, 0x32, 0x87,
	0x49, 0x97, 0xac, 0x44, 0x51, 0xf0, 0x0a, 0x2f, 0xdf, 0x90, 0x3d, 0x9a, 0xd3, 0x9b, 0xd7, 0xe1,
	0x7f, 0x11, 0x0a, 0xaf, 0xe4, 0x90, 0x3d, 0xe0, 0xa7, 0x7f, 0x43, 0x0f, 0x3b, 0xa7, 0xdb, 0x51,
	0x9e, 0x07, 0x8a, 0x31, 0xb6, 0xe8, 0x3c, 0xce, 0xe6, 0xc1, 0x9b, 0x40, 0x2f, 0x6b, 0x0d, 0xb3,
	0x01, 0xc4, 0x03, 0x4b, 0xa9, 0x35, 0xbe, 0x5b, 0x2e, 0xd3, 0x21, 0x4e, 0x33, 0x4f, 0x55, 0x8a,
	0x3d, 0x74, 0x58, 0x03, 0xb4, 0x6e, 0x97, 0x57, 0xd8, 0xb2, 0x01, 0xd3, 0xe1, 0x2c, 0x07, 0x1b,
	0x0f, 0xea, 0x81, 0xd3, 0x2c, 0x58, 0x85, 0x07, 0xf7, 0x0c, 0x72, 0x09, 0xcf, 0x5a, 0x03, 0x46,
	0xd2, 0x24, 0x14, 0x26, 0x8b, 0x28, 0xb9, 0xd9, 0xa9, 0x42, 0x8a, 0xb2, 0x38, 0xcb, 0x2d, 0xf2,
	0x39, 0xbc, 0x58, 0x44, 0x49, 0x58, 0x66, 0x72, 0x23, 0xca, 0xa2, 0x53, 0x66, 0xcf, 0xbe, 0x04,
	0xa7, 0xfd, 0x7c, 0x92, 0x01, 0xd8, 0x61, 0xb4, 0x5c, 0x79, 0x06, 0x19, 0x42, 0x4f, 0x47, 0x3f,
	0x78, 0xe6, 0x6c, 0x06, 0xe3, 0xee, 0x0b, 0x49, 0xfa, 0x60, 0xdd, 0x2d, 0x62, 0xcf, 0xd0, 0xc1,
	0xea, 0x3a, 0xf6, 0xcc, 0xd9, 0xef, 0x00, 0xad, 0xe9, 0x8c, 0xa0, 0x7f, 0x17, 0xfd, 0x1a, 0xdd,
	0xfc, 0x16, 0x35, 0x4e, 0xb7, 0xd1, 0xcf, 0xef, 0x97, 0x9e, 0x49, 0x1c, 0x18, 0x5c, 0xbf, 0x0b,
	0x97, 0xd1, 0x2a, 0x8c, 0x3d, 0x8b, 0x8c, 0x01, 0xd8, 0x32, 0x89, 0x6f, 0xa2, 0x64, 0x19, 0xc6,
	0x9e, 0xad, 0xab, 0xa2, 0x64, 0x81, 0xd2, 0x9e, 0xfe, 0x93, 0x28, 0x09, 0x63, 0xef, 0x62, 0xf6,
	0x1d, 0x78, 0xed, 0xc5, 0x78, 0x27, 0x6a, 0x45, 0x5e, 0x81, 0xb5, 0xa9, 0x0b, 0x6a, 0xe0, 0xf6,
	0xbb, 0x9d, 0xc5, 0x61, 0x3a, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x56, 0xce, 0x98,
	0xcb, 0x06, 0x00, 0x00,
}