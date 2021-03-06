// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_CHAT               Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_REJECT       Message_MessageType = 5
	Message_ORDER_CANCEL       Message_MessageType = 6
	Message_ORDER_CONFIRMATION Message_MessageType = 7
	Message_ORDER_FULFILLMENT  Message_MessageType = 8
	Message_ORDER_COMPLETION   Message_MessageType = 9
	Message_DISPUTE_OPEN       Message_MessageType = 10
	Message_DISPUTE_UPDATE     Message_MessageType = 11
	Message_DISPUTE_CLOSE      Message_MessageType = 12
	Message_REFUND             Message_MessageType = 13
	Message_OFFLINE_ACK        Message_MessageType = 14
	Message_OFFLINE_RELAY      Message_MessageType = 15
	Message_MODERATOR_ADD      Message_MessageType = 16
	Message_MODERATOR_REMOVE   Message_MessageType = 17
	Message_STORE              Message_MessageType = 18
	Message_BLOCK              Message_MessageType = 19
	Message_ERROR              Message_MessageType = 500
)

var Message_MessageType_name = map[int32]string{
	0:   "PING",
	1:   "CHAT",
	2:   "FOLLOW",
	3:   "UNFOLLOW",
	4:   "ORDER",
	5:   "ORDER_REJECT",
	6:   "ORDER_CANCEL",
	7:   "ORDER_CONFIRMATION",
	8:   "ORDER_FULFILLMENT",
	9:   "ORDER_COMPLETION",
	10:  "DISPUTE_OPEN",
	11:  "DISPUTE_UPDATE",
	12:  "DISPUTE_CLOSE",
	13:  "REFUND",
	14:  "OFFLINE_ACK",
	15:  "OFFLINE_RELAY",
	16:  "MODERATOR_ADD",
	17:  "MODERATOR_REMOVE",
	18:  "STORE",
	19:  "BLOCK",
	500: "ERROR",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"CHAT":               1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_REJECT":       5,
	"ORDER_CANCEL":       6,
	"ORDER_CONFIRMATION": 7,
	"ORDER_FULFILLMENT":  8,
	"ORDER_COMPLETION":   9,
	"DISPUTE_OPEN":       10,
	"DISPUTE_UPDATE":     11,
	"DISPUTE_CLOSE":      12,
	"REFUND":             13,
	"OFFLINE_ACK":        14,
	"OFFLINE_RELAY":      15,
	"MODERATOR_ADD":      16,
	"MODERATOR_REMOVE":   17,
	"STORE":              18,
	"BLOCK":              19,
	"ERROR":              500,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Chat_Flag int32

const (
	Chat_MESSAGE Chat_Flag = 0
	Chat_TYPING  Chat_Flag = 1
	Chat_READ    Chat_Flag = 2
)

var Chat_Flag_name = map[int32]string{
	0: "MESSAGE",
	1: "TYPING",
	2: "READ",
}
var Chat_Flag_value = map[string]int32{
	"MESSAGE": 0,
	"TYPING":  1,
	"READ":    2,
}

func (x Chat_Flag) String() string {
	return proto.EnumName(Chat_Flag_name, int32(x))
}
func (Chat_Flag) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

type Message struct {
	MessageType Message_MessageType   `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf1.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	RequestId   int32                 `protobuf:"varint,3,opt,name=requestId" json:"requestId,omitempty"`
	IsResponse  bool                  `protobuf:"varint,4,opt,name=isResponse" json:"isResponse,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_PING
}

func (m *Message) GetPayload() *google_protobuf1.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Message) GetIsResponse() bool {
	if m != nil {
		return m.IsResponse
	}
	return false
}

type Envelope struct {
	Message   *Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Pubkey    []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Chat struct {
	MessageId string                     `protobuf:"bytes,1,opt,name=messageId" json:"messageId,omitempty"`
	Subject   string                     `protobuf:"bytes,2,opt,name=subject" json:"subject,omitempty"`
	Message   string                     `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Flag      Chat_Flag                  `protobuf:"varint,5,opt,name=flag,enum=Chat_Flag" json:"flag,omitempty"`
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Chat) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Chat) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Chat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Chat) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Chat) GetFlag() Chat_Flag {
	if m != nil {
		return m.Flag
	}
	return Chat_MESSAGE
}

type SignedData struct {
	SenderPubkey   []byte `protobuf:"bytes,1,opt,name=senderPubkey,proto3" json:"senderPubkey,omitempty"`
	SerializedData []byte `protobuf:"bytes,2,opt,name=serializedData,proto3" json:"serializedData,omitempty"`
	Signature      []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedData) Reset()                    { *m = SignedData{} }
func (m *SignedData) String() string            { return proto.CompactTextString(m) }
func (*SignedData) ProtoMessage()               {}
func (*SignedData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *SignedData) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *SignedData) GetSerializedData() []byte {
	if m != nil {
		return m.SerializedData
	}
	return nil
}

func (m *SignedData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignedData_Command struct {
	PeerID    string                     `protobuf:"bytes,1,opt,name=peerID" json:"peerID,omitempty"`
	Type      Message_MessageType        `protobuf:"varint,2,opt,name=type,enum=Message_MessageType" json:"type,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *SignedData_Command) Reset()                    { *m = SignedData_Command{} }
func (m *SignedData_Command) String() string            { return proto.CompactTextString(m) }
func (*SignedData_Command) ProtoMessage()               {}
func (*SignedData_Command) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3, 0} }

func (m *SignedData_Command) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *SignedData_Command) GetType() Message_MessageType {
	if m != nil {
		return m.Type
	}
	return Message_PING
}

func (m *SignedData_Command) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type CidList struct {
	Cids []string `protobuf:"bytes,1,rep,name=cids" json:"cids,omitempty"`
}

func (m *CidList) Reset()                    { *m = CidList{} }
func (m *CidList) String() string            { return proto.CompactTextString(m) }
func (*CidList) ProtoMessage()               {}
func (*CidList) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *CidList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type Block struct {
	RawData []byte `protobuf:"bytes,1,opt,name=rawData,proto3" json:"rawData,omitempty"`
	Cid     string `protobuf:"bytes,2,opt,name=cid" json:"cid,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *Block) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func (m *Block) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

type Error struct {
	Code         uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage" json:"errorMessage,omitempty"`
	OrderID      string `protobuf:"bytes,3,opt,name=orderID" json:"orderID,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *Error) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *Error) GetOrderID() string {
	if m != nil {
		return m.OrderID
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Chat)(nil), "Chat")
	proto.RegisterType((*SignedData)(nil), "SignedData")
	proto.RegisterType((*SignedData_Command)(nil), "SignedData.Command")
	proto.RegisterType((*CidList)(nil), "CidList")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("Chat_Flag", Chat_Flag_name, Chat_Flag_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x8e, 0xe3, 0x44,
	0x10, 0x5d, 0x27, 0xce, 0x38, 0x29, 0x27, 0xd9, 0x9e, 0x66, 0x58, 0x85, 0x11, 0x2c, 0x91, 0x0f,
	0x28, 0x5c, 0xbc, 0x52, 0x56, 0x42, 0x5c, 0x3d, 0x76, 0x7b, 0x31, 0xeb, 0xd8, 0x51, 0xc7, 0x01,
	0x0d, 0x97, 0xc8, 0x89, 0x7b, 0x83, 0xd9, 0xc4, 0x36, 0xb6, 0x03, 0x0a, 0x77, 0x3e, 0x81, 0xcf,
	0xe3, 0xc4, 0x2f, 0x70, 0x46, 0xa8, 0xdb, 0x6d, 0x32, 0xb3, 0x48, 0x2b, 0xed, 0xc9, 0x55, 0xaf,
	0x9e, 0xab, 0xab, 0x5f, 0xbd, 0x86, 0xd1, 0x91, 0x55, 0x55, 0xbc, 0x67, 0x66, 0x51, 0xe6, 0x75,
	0x7e, 0xfb, 0xc9, 0x3e, 0xcf, 0xf7, 0x07, 0xf6, 0x42, 0x64, 0xdb, 0xd3, 0x9b, 0x17, 0x71, 0x76,
	0x96, 0xa5, 0xcf, 0xdf, 0x2d, 0xd5, 0xe9, 0x91, 0x55, 0x75, 0x7c, 0x2c, 0x1a, 0x82, 0xf1, 0x87,
	0x0a, 0xda, 0xa2, 0xe9, 0x86, 0xbf, 0x02, 0x5d, 0x36, 0x8e, 0xce, 0x05, 0x9b, 0x28, 0x53, 0x65,
	0x36, 0x9e, 0xdf, 0x98, 0xb2, 0xdc, 0x7e, 0x79, 0x8d, 0x3e, 0x24, 0x62, 0x13, 0xb4, 0x22, 0x3e,
	0x1f, 0xf2, 0x38, 0x99, 0x74, 0xa6, 0xca, 0x4c, 0x9f, 0xdf, 0x98, 0xcd, 0xb1, 0x66, 0x7b, 0xac,
	0x69, 0x65, 0x67, 0xda, 0x92, 0xf0, 0xa7, 0x30, 0x28, 0xd9, 0xcf, 0x27, 0x56, 0xd5, 0x5e, 0x32,
	0xe9, 0x4e, 0x95, 0x59, 0x8f, 0x5e, 0x00, 0xfc, 0x1c, 0x20, 0xad, 0x28, 0xab, 0x8a, 0x3c, 0xab,
	0xd8, 0x44, 0x9d, 0x2a, 0xb3, 0x3e, 0x7d, 0x80, 0x18, 0x7f, 0x75, 0x40, 0x7f, 0x30, 0x0a, 0xee,
	0x83, 0xba, 0xf4, 0x82, 0x57, 0xe8, 0x09, 0x8f, 0xec, 0x6f, 0xac, 0x08, 0x29, 0x18, 0xe0, 0xca,
	0x0d, 0x7d, 0x3f, 0xfc, 0x1e, 0x75, 0xf0, 0x10, 0xfa, 0xeb, 0x40, 0x66, 0x5d, 0x3c, 0x80, 0x5e,
	0x48, 0x1d, 0x42, 0x91, 0x8a, 0x11, 0x0c, 0x45, 0xb8, 0xa1, 0xe4, 0x5b, 0x62, 0x47, 0xa8, 0x77,
	0x41, 0x6c, 0x2b, 0xb0, 0x89, 0x8f, 0xae, 0xf0, 0x33, 0xc0, 0x12, 0x09, 0x03, 0xd7, 0xa3, 0x0b,
	0x2b, 0xf2, 0xc2, 0x00, 0x69, 0xf8, 0x63, 0xb8, 0x6e, 0x70, 0x77, 0xed, 0xbb, 0x9e, 0xef, 0x2f,
	0x48, 0x10, 0xa1, 0x3e, 0xbe, 0x01, 0xd4, 0xd2, 0x17, 0x4b, 0x9f, 0x08, 0xf2, 0x80, 0xb7, 0x75,
	0xbc, 0xd5, 0x72, 0x1d, 0x91, 0x4d, 0xb8, 0x24, 0x01, 0x02, 0x8c, 0x61, 0xdc, 0x22, 0xeb, 0xa5,
	0x63, 0x45, 0x04, 0xe9, 0xf8, 0x1a, 0x46, 0x2d, 0x66, 0xfb, 0xe1, 0x8a, 0xa0, 0x21, 0xbf, 0x06,
	0x25, 0xee, 0x3a, 0x70, 0xd0, 0x08, 0x3f, 0x05, 0x3d, 0x74, 0x5d, 0xdf, 0x0b, 0xc8, 0xc6, 0xb2,
	0x5f, 0xa3, 0x31, 0xe7, 0xb7, 0x00, 0x25, 0xbe, 0x75, 0x8f, 0x9e, 0x72, 0x68, 0x11, 0x3a, 0x84,
	0x5a, 0x51, 0x48, 0x37, 0x96, 0xe3, 0x20, 0xc4, 0x27, 0xba, 0x40, 0x94, 0x2c, 0xc2, 0xef, 0x08,
	0xba, 0xe6, 0x2a, 0xac, 0xa2, 0x90, 0x12, 0x84, 0x79, 0x78, 0xe7, 0x87, 0xf6, 0x6b, 0xf4, 0x11,
	0x06, 0xe8, 0x11, 0x4a, 0x43, 0x8a, 0xfe, 0xee, 0x1a, 0x09, 0xf4, 0x49, 0xf6, 0x0b, 0x3b, 0xe4,
	0x05, 0xc3, 0x06, 0x68, 0x72, 0xdd, 0xc2, 0x13, 0xfa, 0xbc, 0xdf, 0x7a, 0x81, 0xb6, 0x05, 0xfc,
	0x0c, 0xae, 0x8a, 0xd3, 0xf6, 0x2d, 0x3b, 0x0b, 0x0b, 0x0c, 0xa9, 0xcc, 0xf8, 0xae, 0xab, 0x74,
	0x9f, 0xc5, 0xf5, 0xa9, 0x64, 0x62, 0xd7, 0x43, 0x7a, 0x01, 0x8c, 0x3f, 0x15, 0x50, 0xed, 0x1f,
	0xe3, 0x9a, 0xd3, 0x64, 0x27, 0x2f, 0x11, 0x87, 0x0c, 0xe8, 0x05, 0xc0, 0x13, 0xd0, 0xaa, 0xd3,
	0xf6, 0x27, 0xb6, 0xab, 0x45, 0xf7, 0x01, 0x6d, 0x53, 0x5e, 0x69, 0x47, 0xeb, 0x36, 0x95, 0x76,
	0xa0, 0xaf, 0x61, 0xf0, 0x9f, 0xd7, 0x85, 0x8b, 0xf4, 0xf9, 0xed, 0xff, 0x6c, 0x19, 0xb5, 0x0c,
	0x7a, 0x21, 0xe3, 0xe7, 0xa0, 0xbe, 0x39, 0xc4, 0xfb, 0x49, 0x4f, 0xf8, 0x1f, 0x4c, 0x3e, 0xa0,
	0xe9, 0x1e, 0xe2, 0x3d, 0x15, 0xb8, 0xf1, 0x25, 0xa8, 0x3c, 0xc3, 0x3a, 0x68, 0x0b, 0xb2, 0x5a,
	0x59, 0xaf, 0x08, 0x7a, 0xc2, 0x57, 0x15, 0xdd, 0x0b, 0x1f, 0x2a, 0xdc, 0x87, 0x94, 0x58, 0x0e,
	0xea, 0x18, 0xff, 0x28, 0x00, 0xab, 0x74, 0x9f, 0xb1, 0xc4, 0x89, 0xeb, 0x18, 0x1b, 0x30, 0xac,
	0x58, 0x96, 0xb0, 0x72, 0xd9, 0x48, 0xa5, 0x08, 0x3d, 0x1e, 0x61, 0xf8, 0x0b, 0x18, 0x57, 0xac,
	0x4c, 0xe3, 0x43, 0xfa, 0x5b, 0xf3, 0x97, 0x14, 0xf4, 0x1d, 0xf4, 0xfd, 0xc2, 0xde, 0xfe, 0xae,
	0x80, 0x66, 0xe7, 0xc7, 0x63, 0x9c, 0x25, 0x62, 0x35, 0x8c, 0x95, 0x9e, 0x23, 0x85, 0x95, 0x19,
	0x9e, 0x81, 0x5a, 0xf3, 0x77, 0xde, 0x79, 0xcf, 0x3b, 0x17, 0x8c, 0xc7, 0x5a, 0x76, 0x3f, 0x40,
	0x4b, 0xe3, 0x33, 0xd0, 0xec, 0x34, 0xf1, 0xd3, 0xaa, 0xc6, 0x18, 0xd4, 0x5d, 0x9a, 0x54, 0x13,
	0x65, 0xda, 0x9d, 0x0d, 0xa8, 0x88, 0x8d, 0x97, 0xd0, 0xbb, 0x3b, 0xe4, 0xbb, 0xb7, 0x7c, 0x8f,
	0x65, 0xfc, 0xab, 0xb8, 0x6e, 0x23, 0x4a, 0x9b, 0x62, 0x04, 0xdd, 0x5d, 0x9a, 0xc8, 0xbd, 0xf3,
	0xd0, 0xb8, 0x87, 0x1e, 0x29, 0xcb, 0xbc, 0x14, 0x1d, 0xf3, 0xa4, 0x31, 0xe5, 0x88, 0x8a, 0x98,
	0x4b, 0xcc, 0x78, 0x51, 0x5e, 0x42, 0xfe, 0xf7, 0x08, 0xe3, 0x87, 0xe5, 0x65, 0x22, 0x14, 0x91,
	0xa6, 0x91, 0xe9, 0x9d, 0xfa, 0x43, 0xa7, 0xd8, 0x6e, 0xaf, 0xc4, 0x9d, 0x5e, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x9c, 0x4d, 0x60, 0x67, 0x05, 0x00, 0x00,
}
