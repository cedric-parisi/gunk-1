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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{0}
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{1}
}

// Transaction holds all details about a transaction.
type Transaction struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID string `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	// SourceAccount is the account emitting the transaction.
	SourceAccount *BankAccountInfo `protobuf:"bytes,2,opt,name=SourceAccount,json=source_account_id,proto3" json:"SourceAccount,omitempty"`
	// DestinationAccount is the account receiving the transaction.
	DestinationAccount *BankAccountInfo `protobuf:"bytes,3,opt,name=DestinationAccount,json=destination_account_id,proto3" json:"DestinationAccount,omitempty"`
	// Date is the date of the transaction.
	Date string `protobuf:"bytes,4,opt,name=Date,json=date,proto3" json:"Date,omitempty"`
	// Type is the type of transaction.
	Type Type `protobuf:"varint,5,opt,name=Type,json=type,proto3,enum=transactions.Type" json:"Type,omitempty"`
	// Status is the status of the transaction.
	Status Status `protobuf:"varint,6,opt,name=Status,json=status,proto3,enum=transactions.Status" json:"Status,omitempty"`
	// Amount holds the amount value and currency of the transaction.
	Amount *types.Amount `protobuf:"bytes,7,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Description describes the transaction.
	Description string `protobuf:"bytes,8,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{0}
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

func (m *Transaction) GetSourceAccount() *BankAccountInfo {
	if m != nil {
		return m.SourceAccount
	}
	return nil
}

func (m *Transaction) GetDestinationAccount() *BankAccountInfo {
	if m != nil {
		return m.DestinationAccount
	}
	return nil
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

func (m *Transaction) GetDescription() string {
	if m != nil {
		return m.Description
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

// BankAccountInfo holds a lightweight account information.
type BankAccountInfo struct {
	// AccountID is the identifier of the account.
	AccountID string `protobuf:"bytes,1,opt,name=AccountID,json=account_id,proto3" json:"AccountID,omitempty"`
	// BankCode is code of the bank the account belongs to.
	BankCode types.BankCode `protobuf:"varint,2,opt,name=BankCode,json=bank_code,proto3,enum=types.BankCode" json:"BankCode,omitempty"`
	// OwnerName is the name of the owner of the account.
	OwnerName string `protobuf:"bytes,3,opt,name=OwnerName,json=owner_name,proto3" json:"OwnerName,omitempty"`
	// MajorType is the type of account.
	MajorType            types.MajorType `protobuf:"varint,4,opt,name=MajorType,json=major_type,proto3,enum=types.MajorType" json:"MajorType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BankAccountInfo) Reset()         { *m = BankAccountInfo{} }
func (m *BankAccountInfo) String() string { return proto.CompactTextString(m) }
func (*BankAccountInfo) ProtoMessage()    {}
func (*BankAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_all_07be4f7b4536dcbb, []int{1}
}
func (m *BankAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankAccountInfo.Unmarshal(m, b)
}
func (m *BankAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankAccountInfo.Marshal(b, m, deterministic)
}
func (dst *BankAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankAccountInfo.Merge(dst, src)
}
func (m *BankAccountInfo) XXX_Size() int {
	return xxx_messageInfo_BankAccountInfo.Size(m)
}
func (m *BankAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BankAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BankAccountInfo proto.InternalMessageInfo

func (m *BankAccountInfo) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *BankAccountInfo) GetBankCode() types.BankCode {
	if m != nil {
		return m.BankCode
	}
	return types.BankCode__
}

func (m *BankAccountInfo) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *BankAccountInfo) GetMajorType() types.MajorType {
	if m != nil {
		return m.MajorType
	}
	return types.MajorType__
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{2}
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{3}
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{4}
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{5}
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
	return fileDescriptor_all_07be4f7b4536dcbb, []int{6}
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
	proto.RegisterType((*BankAccountInfo)(nil), "transactions.BankAccountInfo")
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
	proto.RegisterFile("command-line-arguments/all.proto", fileDescriptor_all_07be4f7b4536dcbb)
}

var fileDescriptor_all_07be4f7b4536dcbb = []byte{
	// 1565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x5d, 0x6f, 0x1b, 0x4d,
	0x15, 0xde, 0xb5, 0x1d, 0xc7, 0x9e, 0x90, 0xc4, 0x99, 0x7e, 0xe0, 0xd7, 0x7a, 0x79, 0x35, 0xa4,
	0x2d, 0x0d, 0x51, 0xb3, 0xfe, 0x68, 0x41, 0x55, 0x50, 0x45, 0x9d, 0x04, 0xda, 0x94, 0x7e, 0x44,
	0x4e, 0xa9, 0xa0, 0x80, 0xac, 0xf1, 0xee, 0xb1, 0xbd, 0xcd, 0x7a, 0xc6, 0xcc, 0xcc, 0xe6, 0x03,
	0x84, 0x84, 0xb8, 0x80, 0x5e, 0x56, 0xa9, 0xb8, 0xe3, 0x8e, 0xff, 0x80, 0xfa, 0x0b, 0xb8, 0x46,
	0x42, 0x42, 0xdc, 0xc0, 0x15, 0x42, 0x02, 0x71, 0x03, 0x08, 0xa9, 0x97, 0x68, 0x66, 0xd7, 0xce,
	0xae, 0xed, 0xb4, 0x7d, 0x7b, 0xd3, 0x6a, 0xe6, 0x9c, 0xf3, 0xcc, 0x39, 0xcf, 0x79, 0xce, 0xc9,
	0x1a, 0x11, 0x97, 0x0f, 0x06, 0x94, 0x79, 0x1b, 0x81, 0xcf, 0x60, 0x83, 0x8a, 0x5e, 0x38, 0x00,
	0xa6, 0x64, 0x95, 0x06, 0x81, 0x33, 0x14, 0x5c, 0x71, 0xfc, 0x05, 0x25, 0x28, 0x93, 0xd4, 0x55,
	0x3e, 0x67, 0xb2, 0xf2, 0x69, 0x8f, 0xf3, 0x5e, 0x00, 0x55, 0x3a, 0xf4, 0xab, 0x94, 0x31, 0xae,
	0xa8, 0xb9, 0x8f, 0x7c, 0x2b, 0x37, 0xcc, 0x7f, 0xee, 0x46, 0x0f, 0xd8, 0x86, 0x3c, 0xa2, 0xbd,
	0x1e, 0x88, 0x2a, 0x1f, 0x1a, 0x8f, 0x19, 0xde, 0xb5, 0x9e, 0xaf, 0xfa, 0x61, 0xc7, 0x71, 0xf9,
	0xa0, 0xca, 0x87, 0xc0, 0x3a, 0x94, 0x1d, 0x54, 0x7b, 0xe1, 0xe8, 0x9f, 0xc3, 0x7a, 0x55, 0x9d,
	0x0c, 0x21, 0x91, 0xcb, 0xea, 0xef, 0x72, 0x68, 0xe1, 0xe9, 0x59, 0x3a, 0xb8, 0x8e, 0x16, 0x13,
	0xc7, 0xdd, 0x9d, 0xb2, 0x4d, 0xec, 0xb5, 0xe2, 0x16, 0x2a, 0x58, 0x65, 0x6b, 0xcd, 0xaa, 0x59,
	0x7b, 0x56, 0x6b, 0x29, 0x91, 0x7e, 0xdb, 0xf7, 0xf0, 0x3e, 0x5a, 0xdc, 0xe7, 0xa1, 0x70, 0xa1,
	0xe9, 0xba, 0x3c, 0x64, 0xaa, 0x9c, 0x21, 0xf6, 0xda, 0x42, 0xe3, 0x4b, 0x4e, 0xb2, 0x4c, 0x67,
	0x8b, 0xb2, 0x83, 0xd8, 0x61, 0x97, 0x75, 0x79, 0x0a, 0x71, 0x45, 0x9a, 0xf8, 0x36, 0x8d, 0xec,
	0x1a, 0xf4, 0x47, 0x08, 0xef, 0x80, 0x54, 0x3e, 0x33, 0xf5, 0x8d, 0x90, 0xb3, 0x9f, 0x17, 0xf9,
	0xb2, 0x77, 0x06, 0x92, 0x84, 0xff, 0x0c, 0xe5, 0x76, 0xa8, 0x82, 0x72, 0x6e, 0xaa, 0xba, 0x9c,
	0x47, 0x15, 0xe0, 0x06, 0xca, 0x3d, 0x3d, 0x19, 0x42, 0x79, 0x8e, 0xd8, 0x6b, 0x4b, 0x0d, 0x9c,
	0x7e, 0x50, 0x5b, 0xd2, 0x31, 0x9a, 0x55, 0x7c, 0x1b, 0xe5, 0xf7, 0x15, 0x55, 0xa1, 0x2c, 0xe7,
	0x4d, 0xd4, 0xc5, 0x74, 0x54, 0x64, 0x4b, 0xc5, 0xe5, 0xa5, 0xb9, 0xc3, 0x75, 0x94, 0x6f, 0x0e,
	0x4c, 0x81, 0xf3, 0xa6, 0xc0, 0x45, 0xc7, 0xb4, 0xc9, 0x89, 0x2e, 0xd3, 0x21, 0xd4, 0xdc, 0xe1,
	0x1b, 0x68, 0x61, 0x07, 0xa4, 0x2b, 0x7c, 0xa3, 0x85, 0x72, 0x61, 0xaa, 0x8e, 0x05, 0xef, 0xcc,
	0x8c, 0xaf, 0xa0, 0xfc, 0x77, 0x25, 0x88, 0xdd, 0x9d, 0x72, 0x71, 0xca, 0x71, 0x3e, 0x94, 0x20,
	0x34, 0x27, 0x57, 0xd1, 0x7c, 0x0b, 0x06, 0x54, 0x1c, 0xc8, 0x32, 0x9a, 0xf6, 0x12, 0x91, 0x69,
	0x33, 0x5f, 0xb0, 0x4a, 0x56, 0xd9, 0x5a, 0xfd, 0xbb, 0x8d, 0x96, 0x27, 0x98, 0xc7, 0x5f, 0x45,
	0xc5, 0xd1, 0x71, 0x96, 0x70, 0x50, 0xa2, 0x01, 0xb7, 0x51, 0x41, 0x47, 0x6f, 0x73, 0x0f, 0x8c,
	0x5e, 0x96, 0x1a, 0xcb, 0x71, 0xd1, 0xa3, 0xeb, 0x54, 0x68, 0x51, 0x2b, 0xb9, 0xed, 0x72, 0x0f,
	0xf4, 0x23, 0x4f, 0x8e, 0x18, 0x88, 0xc7, 0x74, 0x00, 0x46, 0x10, 0x13, 0x8f, 0x70, 0x6d, 0x6c,
	0x33, 0x3a, 0x00, 0xfc, 0x0d, 0x54, 0x7c, 0x44, 0x5f, 0x70, 0x61, 0x5a, 0x99, 0x33, 0xaf, 0x94,
	0xe2, 0x57, 0xc6, 0xf7, 0xe9, 0xe0, 0x81, 0xbe, 0x6e, 0x6b, 0x97, 0x71, 0xa1, 0x2d, 0x74, 0xe9,
	0x1e, 0xa8, 0xc4, 0x50, 0xb4, 0xe0, 0xc7, 0x21, 0x48, 0xf5, 0x11, 0xa3, 0x32, 0xc6, 0xfc, 0x21,
	0xba, 0x9c, 0xc6, 0x94, 0x23, 0xd0, 0x4d, 0xb4, 0xf2, 0x18, 0x8e, 0xd5, 0xbe, 0xa2, 0x42, 0xf9,
	0xac, 0xb7, 0xcb, 0x3c, 0x38, 0x9e, 0x01, 0x7c, 0x81, 0xc1, 0xb1, 0x6a, 0xcb, 0xd8, 0xab, 0xed,
	0x6b, 0xb7, 0x31, 0xfa, 0xaf, 0x6c, 0xf4, 0xc5, 0x29, 0x78, 0x39, 0xe4, 0x4c, 0x02, 0xbe, 0x83,
	0xf2, 0x2d, 0x90, 0x61, 0xa0, 0xca, 0x36, 0xc9, 0xae, 0x2d, 0x34, 0x3e, 0x99, 0x90, 0xf6, 0xd9,
	0x21, 0x2d, 0x3b, 0x61, 0x82, 0xf0, 0x35, 0x34, 0x7f, 0x9f, 0xca, 0x47, 0x5c, 0x44, 0x5d, 0x2b,
	0xa4, 0x9c, 0x0a, 0x7d, 0x2a, 0xdb, 0x03, 0x2e, 0xce, 0xb8, 0xfb, 0x9f, 0x8d, 0xca, 0xdb, 0x02,
	0xa8, 0x82, 0x19, 0xfc, 0x7d, 0x1d, 0x2d, 0xa7, 0xf6, 0xc6, 0x4c, 0x06, 0x67, 0xac, 0x86, 0xbb,
	0xe8, 0xe2, 0xf4, 0x6a, 0xd8, 0xdd, 0x31, 0x09, 0x15, 0x3f, 0x68, 0xfa, 0xcf, 0xe6, 0x2d, 0xfb,
	0xa1, 0xf3, 0x96, 0x18, 0x8e, 0xdc, 0xfb, 0x87, 0xe3, 0x19, 0xfa, 0x64, 0x46, 0xd9, 0x71, 0x0b,
	0x3e, 0x5e, 0x37, 0xeb, 0xb7, 0xa2, 0xb5, 0x84, 0x8b, 0xc8, 0x6e, 0x97, 0xac, 0x4a, 0xa6, 0x60,
	0xe1, 0x25, 0x94, 0xdf, 0x16, 0xe0, 0xf9, 0xaa, 0x64, 0x9b, 0xf3, 0x22, 0x9a, 0xdb, 0x81, 0x8e,
	0xaf, 0x4a, 0x19, 0x7d, 0xac, 0x64, 0xca, 0xd6, 0xfa, 0x77, 0x46, 0x8b, 0x29, 0x19, 0xb7, 0x8c,
	0xe6, 0xf7, 0x43, 0xd7, 0x05, 0x29, 0xe3, 0xc0, 0x65, 0x34, 0xbf, 0x07, 0xcc, 0xf3, 0x59, 0x2f,
	0x0a, 0xc5, 0x25, 0x54, 0x68, 0xc1, 0x0b, 0x70, 0x15, 0x78, 0xa5, 0xec, 0x08, 0xac, 0xf1, 0xeb,
	0x22, 0xc2, 0x89, 0xf4, 0xf7, 0x41, 0x1c, 0xfa, 0x2e, 0xe0, 0xff, 0x66, 0xd0, 0x52, 0x5a, 0x73,
	0xf8, 0x4a, 0x5a, 0x5a, 0x33, 0x87, 0xa8, 0x72, 0xbe, 0xfe, 0x56, 0x7f, 0x93, 0x39, 0x6d, 0xfe,
	0xd3, 0x4e, 0xff, 0x79, 0xba, 0xdc, 0x02, 0x25, 0x7c, 0x38, 0x04, 0x42, 0x49, 0x22, 0xb0, 0xb2,
	0x3f, 0xba, 0x97, 0x84, 0x06, 0x01, 0xf1, 0xa8, 0xa2, 0xa4, 0x2b, 0xf8, 0x20, 0xed, 0x76, 0x83,
	0x48, 0x08, 0x4c, 0x41, 0xa4, 0x73, 0x42, 0x54, 0x1f, 0x48, 0x9a, 0x67, 0x72, 0xc2, 0x43, 0x22,
	0xc3, 0xe1, 0x30, 0xf0, 0xc1, 0x73, 0x1e, 0x5c, 0x47, 0xd9, 0x46, 0xad, 0x86, 0x09, 0xfa, 0x2c,
	0xce, 0x96, 0xc0, 0x31, 0xb8, 0xa1, 0x0e, 0x97, 0x11, 0x75, 0xdd, 0x30, 0x08, 0x4e, 0x9c, 0x07,
	0x75, 0x94, 0xbd, 0x55, 0xbb, 0x85, 0xd7, 0xd1, 0x5a, 0x0b, 0x54, 0x28, 0x18, 0x78, 0xe4, 0xa8,
	0x0f, 0xcc, 0xbc, 0x20, 0x20, 0x52, 0x30, 0xf1, 0x25, 0x61, 0x5c, 0x91, 0x2e, 0x0f, 0x99, 0xe7,
	0x74, 0x30, 0x2a, 0xa1, 0xfc, 0x93, 0x66, 0xa8, 0xfa, 0x0d, 0x9c, 0x47, 0x39, 0x01, 0xd4, 0xfb,
	0xc5, 0x1f, 0xff, 0xf6, 0x3a, 0xb3, 0x8a, 0x89, 0xf9, 0x33, 0x9d, 0x20, 0xa5, 0xfa, 0xd3, 0x94,
	0x5c, 0x7e, 0xf6, 0x32, 0x63, 0xbd, 0xca, 0x18, 0xa5, 0xe0, 0xbf, 0x64, 0xd1, 0xf2, 0xc4, 0xa4,
	0xe3, 0xab, 0xef, 0xa2, 0x7d, 0xb4, 0x67, 0x2a, 0xd7, 0xde, 0xe3, 0x15, 0x69, 0x75, 0xf5, 0x6d,
	0xe6, 0xb4, 0xf9, 0xa7, 0x4c, 0xba, 0x07, 0x97, 0x1e, 0xfa, 0x52, 0x19, 0x9a, 0x53, 0xdf, 0x31,
	0xff, 0xb1, 0xa3, 0xe2, 0x25, 0xa1, 0x24, 0xd0, 0x1e, 0x2e, 0x67, 0x8a, 0xfa, 0xcc, 0x67, 0x3d,
	0x12, 0x0e, 0x89, 0xe2, 0xa4, 0x51, 0x4b, 0x45, 0x38, 0xe4, 0xfb, 0x3c, 0x24, 0x2e, 0x65, 0x64,
	0x48, 0x7b, 0x7a, 0x4a, 0x81, 0xa8, 0xbe, 0xe0, 0x61, 0xaf, 0x9f, 0x72, 0xd3, 0xad, 0x32, 0x3d,
	0x39, 0xd1, 0x48, 0x33, 0xf6, 0x1e, 0xf1, 0x99, 0x6e, 0x9c, 0x20, 0x32, 0xec, 0x48, 0x5d, 0x1e,
	0x53, 0x92, 0xb8, 0x34, 0x08, 0xa4, 0x33, 0xd3, 0x3d, 0xce, 0x4b, 0xce, 0x12, 0x00, 0xef, 0x9a,
	0xdb, 0x80, 0x4a, 0x75, 0x8e, 0xc9, 0x0d, 0x85, 0x00, 0xa6, 0x74, 0xce, 0xf0, 0xe1, 0x12, 0x39,
	0xb7, 0xdf, 0x18, 0x97, 0x26, 0xfb, 0x9d, 0xe8, 0xef, 0x6f, 0x33, 0x68, 0x65, 0x6a, 0x91, 0xe0,
	0xaf, 0xa4, 0x7b, 0x77, 0xde, 0x82, 0xad, 0x5c, 0x7f, 0xaf, 0x5f, 0xdc, 0xe5, 0x37, 0xf6, 0x69,
	0xf3, 0xf5, 0xc4, 0xa4, 0x5d, 0x8c, 0xdc, 0x27, 0xe6, 0x6c, 0x23, 0xba, 0xd5, 0x3d, 0x66, 0x70,
	0x94, 0x34, 0x11, 0xca, 0x3c, 0x22, 0x62, 0x05, 0xf8, 0x4a, 0x12, 0xdf, 0x73, 0x1e, 0xac, 0x6b,
	0x7a, 0xea, 0xf8, 0x0a, 0xfa, 0x72, 0x02, 0x9a, 0xb8, 0x06, 0x63, 0x92, 0xa1, 0x0b, 0x68, 0x65,
	0xcc, 0xd0, 0x3c, 0x9a, 0x3b, 0x12, 0xbe, 0x02, 0x43, 0xd1, 0xa5, 0x4d, 0x7b, 0x7d, 0xf5, 0x1d,
	0x2c, 0x55, 0xb2, 0x2f, 0x33, 0xd6, 0xd6, 0xbf, 0xe7, 0x4e, 0x9b, 0xbf, 0x9f, 0x43, 0xd9, 0x86,
	0x53, 0xc3, 0x3f, 0x40, 0xa5, 0xa4, 0x9a, 0x49, 0x73, 0x6f, 0x17, 0xdf, 0xdd, 0x13, 0xfc, 0xd0,
	0xf7, 0x40, 0xc6, 0xcf, 0xc7, 0xf9, 0x52, 0x8f, 0xf0, 0x21, 0x88, 0xe8, 0x2b, 0x9a, 0x70, 0x36,
	0x29, 0x8a, 0xf1, 0xfc, 0x3a, 0x8d, 0xb9, 0xba, 0x53, 0x73, 0x6a, 0xab, 0xd9, 0xea, 0x61, 0x7d,
	0xdd, 0xce, 0x34, 0x4a, 0x54, 0xef, 0x09, 0xd7, 0x44, 0x56, 0x5f, 0x48, 0xce, 0x36, 0xa7, 0x6e,
	0x5a, 0x0f, 0xf5, 0x62, 0xa8, 0xe3, 0x6f, 0xa1, 0xed, 0x59, 0x8b, 0x21, 0x12, 0x8c, 0xc7, 0x21,
	0xda, 0x0c, 0x29, 0x5d, 0xea, 0x8f, 0xb2, 0xeb, 0x26, 0x57, 0x0f, 0x98, 0xf2, 0x69, 0x20, 0x9d,
	0xd6, 0x9e, 0x46, 0xbb, 0x89, 0x77, 0xd1, 0xbd, 0x69, 0x34, 0xed, 0x7f, 0x06, 0xd5, 0xa7, 0x87,
	0x40, 0x86, 0x20, 0x06, 0xbe, 0x94, 0xba, 0x08, 0xc5, 0x09, 0x35, 0x5c, 0xa7, 0x56, 0x92, 0xd3,
	0xfa, 0x36, 0xca, 0x7e, 0xad, 0x56, 0xc3, 0xdf, 0x44, 0x77, 0xd2, 0x88, 0x94, 0x91, 0x90, 0xc1,
	0xf1, 0x30, 0xda, 0x95, 0x20, 0x04, 0x17, 0x84, 0xbb, 0x6e, 0x28, 0xc0, 0x1b, 0x71, 0x24, 0x41,
	0x1c, 0x82, 0x20, 0xd2, 0xf7, 0xc0, 0x69, 0xb5, 0x75, 0x66, 0x35, 0xfc, 0x3d, 0xf4, 0xec, 0xfc,
	0x3a, 0x3b, 0xdc, 0x3b, 0xd1, 0x4b, 0x70, 0x40, 0x83, 0x2e, 0x17, 0x03, 0xaa, 0x34, 0x34, 0x4f,
	0x24, 0x3d, 0xa0, 0xca, 0xed, 0x9b, 0x90, 0xf1, 0xcb, 0x71, 0xac, 0xf3, 0xfc, 0x5f, 0x36, 0xfa,
	0x87, 0x3d, 0x96, 0xc7, 0x5f, 0xed, 0x42, 0x16, 0xff, 0xd2, 0x6e, 0xc6, 0x25, 0xf1, 0xf4, 0xbe,
	0x18, 0x95, 0x27, 0xf5, 0x6b, 0x02, 0xa4, 0x12, 0xbe, 0x01, 0xd3, 0x24, 0x84, 0xaa, 0xaf, 0xe9,
	0x74, 0x8d, 0xfe, 0x34, 0x67, 0xd2, 0x21, 0x4f, 0xfb, 0x90, 0x34, 0x68, 0xbe, 0x86, 0x82, 0x1b,
	0xe8, 0x2e, 0x0f, 0x02, 0x7e, 0x14, 0xb1, 0xa6, 0x9f, 0xe6, 0xc2, 0xff, 0x49, 0xe4, 0xa1, 0xbf,
	0x50, 0x49, 0x37, 0xe0, 0x47, 0xce, 0x5a, 0xae, 0x51, 0xd0, 0xe2, 0xd4, 0x10, 0x9b, 0x45, 0x23,
	0x53, 0x7e, 0x00, 0x6c, 0x6b, 0x13, 0x55, 0xa2, 0x39, 0xc7, 0xf8, 0x9e, 0xa0, 0x7a, 0x27, 0x19,
	0xb5, 0x45, 0x7d, 0x40, 0x9f, 0xc6, 0x0a, 0xc7, 0x17, 0x62, 0xa3, 0x39, 0xc5, 0xd6, 0xfb, 0xf6,
	0x9e, 0xf5, 0x3c, 0xf5, 0x2b, 0xf1, 0xe7, 0xb6, 0xf5, 0xd2, 0xb6, 0x5e, 0xd9, 0xd6, 0x1b, 0xdb,
	0xfa, 0xb3, 0x6d, 0xbd, 0xb5, 0xad, 0x3f, 0x64, 0xac, 0x4e, 0xde, 0xfc, 0x86, 0xbb, 0xf9, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x32, 0x5d, 0x74, 0x73, 0x0e, 0x00, 0x00,
}
