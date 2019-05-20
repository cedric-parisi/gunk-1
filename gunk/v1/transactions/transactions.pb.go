// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command-line-arguments/all.proto

package transactions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import types "github.com/openbank/gunk/gunk/v1/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Type defines the type of a transaction.
type Type int32

const (
	Type__ Type = 0
	// Type_Credit is the value for a credit transaction.
	Type_Credit Type = 1
	// Type_Debit is the value for a debit transaction.
	Type_Debit Type = 2
)

var Type_name = map[int32]string{
	0: "_",
	1: "Credit",
	2: "Debit",
}
var Type_value = map[string]int32{
	"_":      0,
	"Credit": 1,
	"Debit":  2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{0}
}

// Status defines the status of a transaction.
type Status int32

const (
	Status__ Status = 0
	// Status_Success is the value for a successful transaction.
	Status_Success Status = 1
	// Status_Pending is the value for a pending transaction.
	Status_Pending Status = 2
	// Status_Rejected is the value for a rejected transaction.
	Status_Rejected Status = 3
)

var Status_name = map[int32]string{
	0: "_",
	1: "Success",
	2: "Pending",
	3: "Rejected",
}
var Status_value = map[string]int32{
	"_":        0,
	"Success":  1,
	"Pending":  2,
	"Rejected": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{1}
}

// Transaction holds all details about a transaction.
type Transaction struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID string `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	// SourceAccountID is the identifier of the account emitting the transaction.
	SourceAccountID string `protobuf:"bytes,2,opt,name=SourceAccountID,json=source_account_id,proto3" json:"SourceAccountID,omitempty"`
	// DestinationAccountID is the identifier of the account receiving the transaction.
	DestinationAccountID string `protobuf:"bytes,3,opt,name=DestinationAccountID,json=destination_account_id,proto3" json:"DestinationAccountID,omitempty"`
	// Date is the date of the transaction.
	Date string `protobuf:"bytes,4,opt,name=Date,json=date,proto3" json:"Date,omitempty"`
	// Type is the type of transaction.
	Type Type `protobuf:"varint,5,opt,name=Type,json=type,proto3,enum=transactions.Type" json:"Type,omitempty"`
	// Status is the status of the transaction.
	Status Status `protobuf:"varint,6,opt,name=Status,json=status,proto3,enum=transactions.Status" json:"Status,omitempty"`
	// Amount holds the amount value and currency of the transaction.
	Amount *types.Amount `protobuf:"bytes,7,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Descriptor ...
	Descriptor_ string `protobuf:"bytes,8,opt,name=Descriptor,json=descriptor,proto3" json:"Descriptor,omitempty"`
	// UserID is the identifier of the issuer of the transaction.
	UserID string `protobuf:"bytes,9,opt,name=UserID,json=user_id,proto3" json:"UserID,omitempty"`
	// Remarks is an informational note about the transaction.
	Remarks              string   `protobuf:"bytes,10,opt,name=Remarks,json=remarks,proto3" json:"Remarks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *Transaction) GetSourceAccountID() string {
	if m != nil {
		return m.SourceAccountID
	}
	return ""
}

func (m *Transaction) GetDestinationAccountID() string {
	if m != nil {
		return m.DestinationAccountID
	}
	return ""
}

func (m *Transaction) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Transaction) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type__
}

func (m *Transaction) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status__
}

func (m *Transaction) GetAmount() *types.Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Transaction) GetDescriptor_() string {
	if m != nil {
		return m.Descriptor_
	}
	return ""
}

func (m *Transaction) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Transaction) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

// GetTransactionRequest is the request envelope to get an transaction by its identifier.
type GetTransactionRequest struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID        string   `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionRequest) Reset()         { *m = GetTransactionRequest{} }
func (m *GetTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionRequest) ProtoMessage()    {}
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{1}
}
func (m *GetTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionRequest.Unmarshal(m, b)
}
func (m *GetTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionRequest.Marshal(b, m, deterministic)
}
func (dst *GetTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionRequest.Merge(dst, src)
}
func (m *GetTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionRequest.Size(m)
}
func (m *GetTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionRequest proto.InternalMessageInfo

func (m *GetTransactionRequest) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

// GetTransactionsRequest is the request envelope to get a list of transactions.
type GetTransactionsRequest struct {
	// NextStartingIndex is a cursor for pagination. It's a TransactionID that defines
	// the current place in the total list of Transaction.
	NextStartingIndex    string   `protobuf:"bytes,1,opt,name=NextStartingIndex,json=next_starting_index,proto3" json:"NextStartingIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionsRequest) Reset()         { *m = GetTransactionsRequest{} }
func (m *GetTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsRequest) ProtoMessage()    {}
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{2}
}
func (m *GetTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsRequest.Unmarshal(m, b)
}
func (m *GetTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsRequest.Marshal(b, m, deterministic)
}
func (dst *GetTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsRequest.Merge(dst, src)
}
func (m *GetTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsRequest.Size(m)
}
func (m *GetTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsRequest proto.InternalMessageInfo

func (m *GetTransactionsRequest) GetNextStartingIndex() string {
	if m != nil {
		return m.NextStartingIndex
	}
	return ""
}

// GetTransactionsResponse wraps the list of transactions.
type GetTransactionsResponse struct {
	// Result is a list containing up to 20 transactions.
	Result []*Transaction `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
	// HasMore indicates if there are more transactions available.
	HasMore              bool     `protobuf:"varint,2,opt,name=HasMore,json=has_more,proto3" json:"HasMore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionsResponse) Reset()         { *m = GetTransactionsResponse{} }
func (m *GetTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsResponse) ProtoMessage()    {}
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{3}
}
func (m *GetTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsResponse.Unmarshal(m, b)
}
func (m *GetTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsResponse.Marshal(b, m, deterministic)
}
func (dst *GetTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsResponse.Merge(dst, src)
}
func (m *GetTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsResponse.Size(m)
}
func (m *GetTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsResponse proto.InternalMessageInfo

func (m *GetTransactionsResponse) GetResult() []*Transaction {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetTransactionsResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

// CreateTransactionRequest wraps all required field for a transaction creation.
type CreateTransactionRequest struct {
	// SourceAccountID is the identifier of the account emitting the transaction.
	SourceAccountID string `protobuf:"bytes,1,opt,name=SourceAccountID,json=source_account_id,proto3" json:"SourceAccountID,omitempty"`
	// DestinationAccountID is the identifier of the account receiving the transaction.
	DestinationAccountID string `protobuf:"bytes,2,opt,name=DestinationAccountID,json=destination_account_id,proto3" json:"DestinationAccountID,omitempty"`
	// Amount holds the amount value and currency of the transaction.
	Amount *types.Amount `protobuf:"bytes,3,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Remarks is an informational note about the transaction.
	Remarks              string   `protobuf:"bytes,4,opt,name=Remarks,json=remarks,proto3" json:"Remarks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransactionRequest) Reset()         { *m = CreateTransactionRequest{} }
func (m *CreateTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTransactionRequest) ProtoMessage()    {}
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{4}
}
func (m *CreateTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransactionRequest.Unmarshal(m, b)
}
func (m *CreateTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransactionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransactionRequest.Merge(dst, src)
}
func (m *CreateTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTransactionRequest.Size(m)
}
func (m *CreateTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransactionRequest proto.InternalMessageInfo

func (m *CreateTransactionRequest) GetSourceAccountID() string {
	if m != nil {
		return m.SourceAccountID
	}
	return ""
}

func (m *CreateTransactionRequest) GetDestinationAccountID() string {
	if m != nil {
		return m.DestinationAccountID
	}
	return ""
}

func (m *CreateTransactionRequest) GetAmount() *types.Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CreateTransactionRequest) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

// CreateTransactionResponse is the response envelope for a transaction creation.
type CreateTransactionResponse struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID        string   `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransactionResponse) Reset()         { *m = CreateTransactionResponse{} }
func (m *CreateTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTransactionResponse) ProtoMessage()    {}
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_ef02d29ff80d86de, []int{5}
}
func (m *CreateTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransactionResponse.Unmarshal(m, b)
}
func (m *CreateTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransactionResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransactionResponse.Merge(dst, src)
}
func (m *CreateTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTransactionResponse.Size(m)
}
func (m *CreateTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransactionResponse proto.InternalMessageInfo

func (m *CreateTransactionResponse) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "transactions.Transaction")
	proto.RegisterType((*GetTransactionRequest)(nil), "transactions.GetTransactionRequest")
	proto.RegisterType((*GetTransactionsRequest)(nil), "transactions.GetTransactionsRequest")
	proto.RegisterType((*GetTransactionsResponse)(nil), "transactions.GetTransactionsResponse")
	proto.RegisterType((*CreateTransactionRequest)(nil), "transactions.CreateTransactionRequest")
	proto.RegisterType((*CreateTransactionResponse)(nil), "transactions.CreateTransactionResponse")
	proto.RegisterEnum("transactions.Type", Type_name, Type_value)
	proto.RegisterEnum("transactions.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// GetTransaction retrieves the detail of a transaction, selected by its id.
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	// GetTransactions returns a list containing up to 20 transactions.
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	// CreateTransaction creates a new transaction and returns its id.
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	// GetTransaction retrieves the detail of a transaction, selected by its id.
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	// GetTransactions returns a list containing up to 20 transactions.
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	// CreateTransaction creates a new transaction and returns its id.
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransaction",
			Handler:    _TransactionService_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _TransactionService_GetTransactions_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command-line-arguments/all.proto",
}

func init() {
	proto.RegisterFile("command-line-arguments/all.proto", fileDescriptor_all_ef02d29ff80d86de)
}

var fileDescriptor_all_ef02d29ff80d86de = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcf, 0x6b, 0x24, 0xc5,
	0x17, 0xaf, 0xea, 0xc9, 0xce, 0x4c, 0x2a, 0xdf, 0x24, 0x93, 0xfa, 0x26, 0xb1, 0x77, 0x10, 0x29,
	0x66, 0x77, 0xdd, 0x10, 0x36, 0x33, 0xc9, 0x28, 0x22, 0x01, 0xc1, 0xec, 0x0e, 0xb8, 0xc1, 0x1f,
	0x84, 0x99, 0x55, 0x50, 0x84, 0xa1, 0xa6, 0xfb, 0x6d, 0xa7, 0xdc, 0x9e, 0xaa, 0xb6, 0xaa, 0x3a,
	0x9b, 0x28, 0x82, 0xec, 0xc5, 0x1c, 0x97, 0x78, 0xf2, 0xe6, 0xc1, 0x9b, 0xff, 0x80, 0xe0, 0x3f,
	0x21, 0x78, 0xf1, 0xe8, 0xc1, 0xbf, 0x40, 0x10, 0xc4, 0x93, 0x74, 0xf7, 0x4c, 0xa6, 0x3b, 0x33,
	0xd9, 0x2c, 0xeb, 0x65, 0x86, 0x7a, 0xf5, 0xf9, 0x3c, 0xde, 0x7b, 0x9f, 0xcf, 0xeb, 0x6e, 0xc2,
	0x3c, 0x35, 0x1c, 0x72, 0xe9, 0x6f, 0x85, 0x42, 0xc2, 0x16, 0xd7, 0x41, 0x3c, 0x04, 0x69, 0x4d,
	0x8b, 0x87, 0x61, 0x33, 0xd2, 0xca, 0x2a, 0xfa, 0x3f, 0xab, 0xb9, 0x34, 0xdc, 0xb3, 0x42, 0x49,
	0x53, 0x7f, 0x39, 0x50, 0x2a, 0x08, 0xa1, 0xc5, 0x23, 0xd1, 0xe2, 0x52, 0x2a, 0xcb, 0xd3, 0x78,
	0x86, 0xad, 0xdf, 0x49, 0xff, 0xbc, 0xad, 0x00, 0xe4, 0x96, 0x79, 0xcc, 0x83, 0x00, 0x74, 0x4b,
	0x45, 0x29, 0x62, 0x06, 0x7a, 0x3b, 0x10, 0xf6, 0x30, 0x1e, 0x34, 0x3d, 0x35, 0x6c, 0xa9, 0x08,
	0xe4, 0x80, 0xcb, 0x47, 0xad, 0x20, 0x1e, 0xff, 0x1c, 0xed, 0xb4, 0xec, 0x49, 0x04, 0xb9, 0x5a,
	0x1a, 0xff, 0x94, 0xc8, 0xc2, 0x83, 0x49, 0x39, 0x74, 0x87, 0x2c, 0xe6, 0x8e, 0xfb, 0x1d, 0x17,
	0x33, 0xbc, 0x31, 0x7f, 0x97, 0x54, 0x91, 0x8b, 0x36, 0xd0, 0x36, 0x3a, 0x40, 0xdd, 0xa5, 0x5c,
	0xf9, 0x7d, 0xe1, 0xd3, 0x37, 0xc8, 0x72, 0x4f, 0xc5, 0xda, 0x83, 0x3d, 0xcf, 0x53, 0xb1, 0xb4,
	0xfb, 0x1d, 0xd7, 0x99, 0x22, 0xad, 0x98, 0x14, 0xd2, 0xe7, 0x19, 0x26, 0xe1, 0xbd, 0x4d, 0x56,
	0x3b, 0x60, 0xac, 0x90, 0x69, 0x0b, 0x13, 0x72, 0x69, 0x8a, 0xbc, 0xee, 0x4f, 0x70, 0xf9, 0x0c,
	0xaf, 0x90, 0xb9, 0x0e, 0xb7, 0xe0, 0xce, 0x4d, 0x31, 0xe6, 0x7c, 0x6e, 0x81, 0xb6, 0xc9, 0xdc,
	0x83, 0x93, 0x08, 0xdc, 0x6b, 0x0c, 0x6f, 0x2c, 0xb5, 0x69, 0x33, 0x3f, 0xf7, 0x66, 0x72, 0x53,
	0xe4, 0x24, 0xb3, 0xa1, 0x6f, 0x92, 0x72, 0xcf, 0x72, 0x1b, 0x1b, 0xb7, 0x9c, 0xb2, 0x56, 0x8b,
	0xac, 0xec, 0xae, 0xc0, 0x2b, 0x9b, 0x34, 0x46, 0x77, 0x48, 0x79, 0x6f, 0x98, 0x94, 0xe6, 0x56,
	0x18, 0xde, 0x58, 0x68, 0x2f, 0x36, 0xd3, 0x61, 0x37, 0xb3, 0x60, 0x91, 0xc2, 0xd3, 0x18, 0xdd,
	0x24, 0xa4, 0x03, 0xc6, 0xd3, 0x22, 0xb2, 0x4a, 0xbb, 0xd5, 0xa9, 0x36, 0x88, 0x7f, 0x7e, 0x4b,
	0x6f, 0x90, 0xf2, 0x87, 0x06, 0xf4, 0x7e, 0xc7, 0x9d, 0x9f, 0xc2, 0x55, 0x62, 0x03, 0x3a, 0x99,
	0xc8, 0x4d, 0x52, 0xe9, 0xc2, 0x90, 0xeb, 0x47, 0xc6, 0x25, 0xd3, 0x28, 0x9d, 0x5d, 0xed, 0x96,
	0xab, 0xa8, 0x86, 0x5c, 0xd4, 0xe8, 0x92, 0xb5, 0x77, 0xc0, 0xe6, 0xf4, 0xee, 0xc2, 0xe7, 0x31,
	0x18, 0xfb, 0x02, 0x2e, 0x38, 0xcf, 0xf9, 0x29, 0x59, 0x2f, 0xe6, 0x34, 0xe3, 0xa4, 0xbb, 0x64,
	0xe5, 0x03, 0x38, 0xb6, 0x3d, 0xcb, 0xb5, 0x15, 0x32, 0xd8, 0x97, 0x3e, 0x1c, 0xcf, 0x48, 0xfc,
	0x7f, 0x09, 0xc7, 0xb6, 0x6f, 0x46, 0xa8, 0xbe, 0x48, 0x60, 0xe7, 0xd9, 0xbf, 0xc1, 0xe4, 0xa5,
	0xa9, 0xf4, 0x26, 0x52, 0xd2, 0x00, 0x7d, 0x8b, 0x94, 0xbb, 0x60, 0xe2, 0xd0, 0xba, 0x98, 0x95,
	0x36, 0x16, 0xda, 0xd7, 0x2f, 0xe8, 0x3d, 0x39, 0x14, 0xb5, 0xd0, 0x29, 0x89, 0xde, 0x22, 0x95,
	0xfb, 0xdc, 0xbc, 0xaf, 0x34, 0xa4, 0xf6, 0xad, 0x16, 0x40, 0xd5, 0x43, 0x6e, 0xfa, 0x43, 0xa5,
	0xe1, 0xbc, 0x92, 0xbf, 0x30, 0x71, 0xef, 0x69, 0xe0, 0x16, 0x66, 0xcc, 0x6f, 0xc6, 0x4a, 0xe0,
	0xff, 0xb2, 0x12, 0xce, 0x73, 0xaf, 0xc4, 0xc4, 0x84, 0xa5, 0xe7, 0x35, 0x61, 0xce, 0x33, 0x73,
	0x57, 0x7b, 0xe6, 0x23, 0x72, 0x7d, 0x46, 0xdb, 0x23, 0x09, 0x5e, 0xdc, 0x37, 0x9b, 0xaf, 0x67,
	0xbb, 0x4a, 0xe7, 0x09, 0xee, 0xd7, 0x50, 0xdd, 0xa9, 0x22, 0xba, 0x44, 0xca, 0xf7, 0x34, 0xf8,
	0xc2, 0xd6, 0x70, 0x7a, 0x5e, 0x24, 0xd7, 0x3a, 0x30, 0x10, 0xb6, 0xe6, 0x24, 0xc7, 0xba, 0xe3,
	0xa2, 0xcd, 0x77, 0xc7, 0xdb, 0x9a, 0xe7, 0x2d, 0x93, 0x4a, 0x2f, 0xf6, 0x3c, 0x30, 0x66, 0x44,
	0x5c, 0x26, 0x95, 0x03, 0x90, 0xbe, 0x90, 0x41, 0x46, 0xa5, 0x35, 0x52, 0xed, 0xc2, 0x67, 0xe0,
	0x59, 0xf0, 0x6b, 0xa5, 0x71, 0xb2, 0xf6, 0x0f, 0x15, 0x42, 0x73, 0xe5, 0xf7, 0x40, 0x1f, 0x09,
	0x0f, 0xe8, 0x13, 0x87, 0x2c, 0x15, 0x3d, 0x47, 0x6f, 0x14, 0xad, 0x35, 0x73, 0x89, 0xea, 0x97,
	0xfb, 0xaf, 0xf1, 0x23, 0x3e, 0xdb, 0xfb, 0xa2, 0xf8, 0xe0, 0x5d, 0xef, 0x82, 0xd5, 0x02, 0x8e,
	0x80, 0x71, 0x96, 0xe3, 0xd5, 0x7b, 0xe3, 0xb8, 0x61, 0x3c, 0x0c, 0x99, 0xcf, 0x2d, 0x67, 0x0f,
	0xb5, 0x1a, 0x16, 0x61, 0x77, 0x98, 0x81, 0x30, 0xed, 0x87, 0x0d, 0x4e, 0x98, 0x3d, 0x04, 0x56,
	0x1c, 0x33, 0x3b, 0x51, 0x31, 0x33, 0x71, 0x14, 0x85, 0x02, 0xfc, 0xe6, 0x93, 0x5f, 0xff, 0xf8,
	0xd6, 0x69, 0x50, 0x96, 0xbe, 0x10, 0x72, 0x35, 0xb6, 0xbe, 0x2c, 0xa8, 0xf7, 0xd5, 0xa9, 0x83,
	0x9e, 0x3a, 0xa9, 0x70, 0xf4, 0xfb, 0x12, 0x59, 0xbe, 0xb0, 0x78, 0xf4, 0xe6, 0xb3, 0xa6, 0x30,
	0x5e, 0xfb, 0xfa, 0xad, 0x2b, 0x50, 0x99, 0x75, 0x1a, 0x3f, 0x3b, 0x67, 0x7b, 0xdf, 0x39, 0xc5,
	0x99, 0xac, 0xbd, 0x27, 0x8c, 0x4d, 0xdb, 0x2e, 0xbc, 0x31, 0xff, 0xc4, 0x5d, 0xb0, 0xb1, 0x96,
	0x86, 0x71, 0x16, 0x26, 0x08, 0x4f, 0x49, 0xcb, 0x85, 0x14, 0x32, 0x60, 0x71, 0xc4, 0xac, 0x62,
	0xed, 0xed, 0x02, 0xa3, 0xc9, 0x3e, 0x56, 0x31, 0xf3, 0xb8, 0x64, 0x11, 0x0f, 0x92, 0xa5, 0x01,
	0x66, 0x0f, 0xb5, 0x8a, 0x83, 0xc3, 0x02, 0x2c, 0x19, 0x5d, 0x3a, 0xa3, 0x93, 0x24, 0xd3, 0x8c,
	0xc7, 0x10, 0x13, 0x32, 0x19, 0xa4, 0x66, 0x26, 0x1e, 0x98, 0xa4, 0x3d, 0x69, 0x0d, 0xf3, 0x78,
	0x18, 0x9a, 0xe6, 0x4c, 0xf8, 0xa8, 0x2e, 0x33, 0x4b, 0x10, 0xf5, 0x30, 0x8d, 0x86, 0xdc, 0xd8,
	0x4b, 0xae, 0xbc, 0x58, 0x6b, 0x90, 0x36, 0xa9, 0x19, 0x32, 0xc9, 0x28, 0xad, 0x5d, 0x94, 0x2c,
	0x27, 0xd1, 0xef, 0x98, 0xac, 0x4c, 0xad, 0x26, 0x7d, 0xb5, 0x38, 0xfe, 0xcb, 0x1e, 0x59, 0xf5,
	0xdb, 0x57, 0xe2, 0x46, 0x42, 0x1d, 0x9d, 0xed, 0x75, 0x8b, 0x32, 0xad, 0x66, 0xe0, 0x0b, 0xc6,
	0xdd, 0xca, 0xa2, 0x89, 0x48, 0x12, 0x1e, 0xe7, 0xaf, 0x18, 0x97, 0x3e, 0xd3, 0x23, 0x09, 0x85,
	0x35, 0x4c, 0x8c, 0x2c, 0xb9, 0xb6, 0x8b, 0x37, 0x1b, 0xcf, 0x68, 0xb1, 0x5e, 0x3a, 0x75, 0xd0,
	0xdd, 0xdb, 0xf7, 0xf1, 0x01, 0xfa, 0xa4, 0xf0, 0x11, 0xf5, 0x35, 0x46, 0xa7, 0x18, 0x3d, 0xc5,
	0xe8, 0x27, 0x8c, 0x7e, 0xc3, 0xe8, 0x6f, 0x8c, 0x7e, 0x71, 0xd0, 0xa0, 0x9c, 0x7e, 0xe2, 0xbc,
	0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xcd, 0x5d, 0xe7, 0x92, 0x09, 0x00, 0x00,
}