// Code generated by protoc-gen-go. DO NOT EDIT.
// source: developer.proto

package developer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

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

type Developer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	IdCardNumber         string   `protobuf:"bytes,5,opt,name=id_card_number,json=idCardNumber,proto3" json:"id_card_number,omitempty"`
	RealName             string   `protobuf:"bytes,6,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	PrivateKey           string   `protobuf:"bytes,7,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicKey            string   `protobuf:"bytes,8,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AccessKey            string   `protobuf:"bytes,9,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	CreateAt             uint64   `protobuf:"varint,10,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt             uint64   `protobuf:"varint,11,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	Active               bool     `protobuf:"varint,12,opt,name=active,proto3" json:"active,omitempty"`
	Verified             bool     `protobuf:"varint,13,opt,name=verified,proto3" json:"verified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Developer) Reset()         { *m = Developer{} }
func (m *Developer) String() string { return proto.CompactTextString(m) }
func (*Developer) ProtoMessage()    {}
func (*Developer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{0}
}

func (m *Developer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Developer.Unmarshal(m, b)
}
func (m *Developer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Developer.Marshal(b, m, deterministic)
}
func (m *Developer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Developer.Merge(m, src)
}
func (m *Developer) XXX_Size() int {
	return xxx_messageInfo_Developer.Size(m)
}
func (m *Developer) XXX_DiscardUnknown() {
	xxx_messageInfo_Developer.DiscardUnknown(m)
}

var xxx_messageInfo_Developer proto.InternalMessageInfo

func (m *Developer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Developer) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Developer) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Developer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Developer) GetIdCardNumber() string {
	if m != nil {
		return m.IdCardNumber
	}
	return ""
}

func (m *Developer) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Developer) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *Developer) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Developer) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *Developer) GetCreateAt() uint64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *Developer) GetUpdateAt() uint64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func (m *Developer) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Developer) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	DeveloperId          string   `protobuf:"bytes,2,opt,name=developer_id,json=developerId,proto3" json:"developer_id,omitempty"`
	SecretKey            string   `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	AppId                string   `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	CryptoType           string   `protobuf:"bytes,5,opt,name=crypto_type,json=cryptoType,proto3" json:"crypto_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetDeveloperId() string {
	if m != nil {
		return m.DeveloperId
	}
	return ""
}

func (m *Session) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *Session) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Session) GetCryptoType() string {
	if m != nil {
		return m.CryptoType
	}
	return ""
}

type SessionRequest struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	CryptoType           string   `protobuf:"bytes,2,opt,name=crypto_type,json=cryptoType,proto3" json:"crypto_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{2}
}

func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (m *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(m, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *SessionRequest) GetCryptoType() string {
	if m != nil {
		return m.CryptoType
	}
	return ""
}

type AppBank struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeveloperId          string   `protobuf:"bytes,2,opt,name=developer_id,json=developerId,proto3" json:"developer_id,omitempty"`
	CryptoType           string   `protobuf:"bytes,3,opt,name=crypto_type,json=cryptoType,proto3" json:"crypto_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppBank) Reset()         { *m = AppBank{} }
func (m *AppBank) String() string { return proto.CompactTextString(m) }
func (*AppBank) ProtoMessage()    {}
func (*AppBank) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{3}
}

func (m *AppBank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppBank.Unmarshal(m, b)
}
func (m *AppBank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppBank.Marshal(b, m, deterministic)
}
func (m *AppBank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppBank.Merge(m, src)
}
func (m *AppBank) XXX_Size() int {
	return xxx_messageInfo_AppBank.Size(m)
}
func (m *AppBank) XXX_DiscardUnknown() {
	xxx_messageInfo_AppBank.DiscardUnknown(m)
}

var xxx_messageInfo_AppBank proto.InternalMessageInfo

func (m *AppBank) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AppBank) GetDeveloperId() string {
	if m != nil {
		return m.DeveloperId
	}
	return ""
}

func (m *AppBank) GetCryptoType() string {
	if m != nil {
		return m.CryptoType
	}
	return ""
}

type AppBankCurrency struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BankId               string   `protobuf:"bytes,2,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
	CoinType             string   `protobuf:"bytes,3,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	Balance              string   `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
	MainNet              string   `protobuf:"bytes,5,opt,name=main_net,json=mainNet,proto3" json:"main_net,omitempty"`
	GasRate              float64  `protobuf:"fixed64,6,opt,name=gas_rate,json=gasRate,proto3" json:"gas_rate,omitempty"`
	Gas                  uint32   `protobuf:"varint,7,opt,name=gas,proto3" json:"gas,omitempty"`
	AddressBalance       string   `protobuf:"bytes,8,opt,name=address_balance,json=addressBalance,proto3" json:"address_balance,omitempty"`
	Address              string   `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	Decimals             uint32   `protobuf:"varint,10,opt,name=decimals,proto3" json:"decimals,omitempty"`
	CryptoType           string   `protobuf:"bytes,11,opt,name=crypto_type,json=cryptoType,proto3" json:"crypto_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppBankCurrency) Reset()         { *m = AppBankCurrency{} }
func (m *AppBankCurrency) String() string { return proto.CompactTextString(m) }
func (*AppBankCurrency) ProtoMessage()    {}
func (*AppBankCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{4}
}

func (m *AppBankCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppBankCurrency.Unmarshal(m, b)
}
func (m *AppBankCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppBankCurrency.Marshal(b, m, deterministic)
}
func (m *AppBankCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppBankCurrency.Merge(m, src)
}
func (m *AppBankCurrency) XXX_Size() int {
	return xxx_messageInfo_AppBankCurrency.Size(m)
}
func (m *AppBankCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AppBankCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AppBankCurrency proto.InternalMessageInfo

func (m *AppBankCurrency) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AppBankCurrency) GetBankId() string {
	if m != nil {
		return m.BankId
	}
	return ""
}

func (m *AppBankCurrency) GetCoinType() string {
	if m != nil {
		return m.CoinType
	}
	return ""
}

func (m *AppBankCurrency) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *AppBankCurrency) GetMainNet() string {
	if m != nil {
		return m.MainNet
	}
	return ""
}

func (m *AppBankCurrency) GetGasRate() float64 {
	if m != nil {
		return m.GasRate
	}
	return 0
}

func (m *AppBankCurrency) GetGas() uint32 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *AppBankCurrency) GetAddressBalance() string {
	if m != nil {
		return m.AddressBalance
	}
	return ""
}

func (m *AppBankCurrency) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AppBankCurrency) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *AppBankCurrency) GetCryptoType() string {
	if m != nil {
		return m.CryptoType
	}
	return ""
}

type AllocateGas struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MainNet              string   `protobuf:"bytes,2,opt,name=main_net,json=mainNet,proto3" json:"main_net,omitempty"`
	GasType              string   `protobuf:"bytes,3,opt,name=gas_type,json=gasType,proto3" json:"gas_type,omitempty"`
	Decimal              uint32   `protobuf:"varint,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	SessionId            string   `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocateGas) Reset()         { *m = AllocateGas{} }
func (m *AllocateGas) String() string { return proto.CompactTextString(m) }
func (*AllocateGas) ProtoMessage()    {}
func (*AllocateGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{5}
}

func (m *AllocateGas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateGas.Unmarshal(m, b)
}
func (m *AllocateGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateGas.Marshal(b, m, deterministic)
}
func (m *AllocateGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateGas.Merge(m, src)
}
func (m *AllocateGas) XXX_Size() int {
	return xxx_messageInfo_AllocateGas.Size(m)
}
func (m *AllocateGas) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateGas.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateGas proto.InternalMessageInfo

func (m *AllocateGas) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AllocateGas) GetMainNet() string {
	if m != nil {
		return m.MainNet
	}
	return ""
}

func (m *AllocateGas) GetGasType() string {
	if m != nil {
		return m.GasType
	}
	return ""
}

func (m *AllocateGas) GetDecimal() uint32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *AllocateGas) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type UnAllocateGas struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MainNet              string   `protobuf:"bytes,2,opt,name=main_net,json=mainNet,proto3" json:"main_net,omitempty"`
	GasType              string   `protobuf:"bytes,3,opt,name=gas_type,json=gasType,proto3" json:"gas_type,omitempty"`
	Decimal              uint32   `protobuf:"varint,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	SessionId            string   `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnAllocateGas) Reset()         { *m = UnAllocateGas{} }
func (m *UnAllocateGas) String() string { return proto.CompactTextString(m) }
func (*UnAllocateGas) ProtoMessage()    {}
func (*UnAllocateGas) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{6}
}

func (m *UnAllocateGas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnAllocateGas.Unmarshal(m, b)
}
func (m *UnAllocateGas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnAllocateGas.Marshal(b, m, deterministic)
}
func (m *UnAllocateGas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnAllocateGas.Merge(m, src)
}
func (m *UnAllocateGas) XXX_Size() int {
	return xxx_messageInfo_UnAllocateGas.Size(m)
}
func (m *UnAllocateGas) XXX_DiscardUnknown() {
	xxx_messageInfo_UnAllocateGas.DiscardUnknown(m)
}

var xxx_messageInfo_UnAllocateGas proto.InternalMessageInfo

func (m *UnAllocateGas) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UnAllocateGas) GetMainNet() string {
	if m != nil {
		return m.MainNet
	}
	return ""
}

func (m *UnAllocateGas) GetGasType() string {
	if m != nil {
		return m.GasType
	}
	return ""
}

func (m *UnAllocateGas) GetDecimal() uint32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *UnAllocateGas) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type AppBankBalance struct {
	BankCurrencies       []*AppBankCurrency `protobuf:"bytes,1,rep,name=bank_currencies,json=bankCurrencies,proto3" json:"bank_currencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AppBankBalance) Reset()         { *m = AppBankBalance{} }
func (m *AppBankBalance) String() string { return proto.CompactTextString(m) }
func (*AppBankBalance) ProtoMessage()    {}
func (*AppBankBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{7}
}

func (m *AppBankBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppBankBalance.Unmarshal(m, b)
}
func (m *AppBankBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppBankBalance.Marshal(b, m, deterministic)
}
func (m *AppBankBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppBankBalance.Merge(m, src)
}
func (m *AppBankBalance) XXX_Size() int {
	return xxx_messageInfo_AppBankBalance.Size(m)
}
func (m *AppBankBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AppBankBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AppBankBalance proto.InternalMessageInfo

func (m *AppBankBalance) GetBankCurrencies() []*AppBankCurrency {
	if m != nil {
		return m.BankCurrencies
	}
	return nil
}

type BankTx struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BankId               string   `protobuf:"bytes,2,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
	TxType               string   `protobuf:"bytes,3,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	AppId                string   `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppUserId            string   `protobuf:"bytes,5,opt,name=app_user_id,json=appUserId,proto3" json:"app_user_id,omitempty"`
	CoinType             string   `protobuf:"bytes,6,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	Amount               string   `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
	MainNet              string   `protobuf:"bytes,8,opt,name=main_net,json=mainNet,proto3" json:"main_net,omitempty"`
	TradeAt              uint64   `protobuf:"varint,9,opt,name=trade_at,json=tradeAt,proto3" json:"trade_at,omitempty"`
	SessionId            string   `protobuf:"bytes,10,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	CurrencyId           string   `protobuf:"bytes,11,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	Decimals             uint32   `protobuf:"varint,12,opt,name=decimals,proto3" json:"decimals,omitempty"`
	TxStatus             uint32   `protobuf:"varint,13,opt,name=tx_status,json=txStatus,proto3" json:"tx_status,omitempty"`
	CryptoType           string   `protobuf:"bytes,14,opt,name=crypto_type,json=cryptoType,proto3" json:"crypto_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankTx) Reset()         { *m = BankTx{} }
func (m *BankTx) String() string { return proto.CompactTextString(m) }
func (*BankTx) ProtoMessage()    {}
func (*BankTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{8}
}

func (m *BankTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankTx.Unmarshal(m, b)
}
func (m *BankTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankTx.Marshal(b, m, deterministic)
}
func (m *BankTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankTx.Merge(m, src)
}
func (m *BankTx) XXX_Size() int {
	return xxx_messageInfo_BankTx.Size(m)
}
func (m *BankTx) XXX_DiscardUnknown() {
	xxx_messageInfo_BankTx.DiscardUnknown(m)
}

var xxx_messageInfo_BankTx proto.InternalMessageInfo

func (m *BankTx) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BankTx) GetBankId() string {
	if m != nil {
		return m.BankId
	}
	return ""
}

func (m *BankTx) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func (m *BankTx) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BankTx) GetAppUserId() string {
	if m != nil {
		return m.AppUserId
	}
	return ""
}

func (m *BankTx) GetCoinType() string {
	if m != nil {
		return m.CoinType
	}
	return ""
}

func (m *BankTx) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *BankTx) GetMainNet() string {
	if m != nil {
		return m.MainNet
	}
	return ""
}

func (m *BankTx) GetTradeAt() uint64 {
	if m != nil {
		return m.TradeAt
	}
	return 0
}

func (m *BankTx) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *BankTx) GetCurrencyId() string {
	if m != nil {
		return m.CurrencyId
	}
	return ""
}

func (m *BankTx) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *BankTx) GetTxStatus() uint32 {
	if m != nil {
		return m.TxStatus
	}
	return 0
}

func (m *BankTx) GetCryptoType() string {
	if m != nil {
		return m.CryptoType
	}
	return ""
}

type BankAddress struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	PrivateKey           string   `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	MainNet              string   `protobuf:"bytes,5,opt,name=main_net,json=mainNet,proto3" json:"main_net,omitempty"`
	CryptoType           string   `protobuf:"bytes,6,opt,name=crypto_type,json=cryptoType,proto3" json:"crypto_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankAddress) Reset()         { *m = BankAddress{} }
func (m *BankAddress) String() string { return proto.CompactTextString(m) }
func (*BankAddress) ProtoMessage()    {}
func (*BankAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_6300c5df2ba41f00, []int{9}
}

func (m *BankAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankAddress.Unmarshal(m, b)
}
func (m *BankAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankAddress.Marshal(b, m, deterministic)
}
func (m *BankAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankAddress.Merge(m, src)
}
func (m *BankAddress) XXX_Size() int {
	return xxx_messageInfo_BankAddress.Size(m)
}
func (m *BankAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_BankAddress.DiscardUnknown(m)
}

var xxx_messageInfo_BankAddress proto.InternalMessageInfo

func (m *BankAddress) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BankAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BankAddress) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BankAddress) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *BankAddress) GetMainNet() string {
	if m != nil {
		return m.MainNet
	}
	return ""
}

func (m *BankAddress) GetCryptoType() string {
	if m != nil {
		return m.CryptoType
	}
	return ""
}

func init() {
	proto.RegisterType((*Developer)(nil), "developer.Developer")
	proto.RegisterType((*Session)(nil), "developer.Session")
	proto.RegisterType((*SessionRequest)(nil), "developer.SessionRequest")
	proto.RegisterType((*AppBank)(nil), "developer.AppBank")
	proto.RegisterType((*AppBankCurrency)(nil), "developer.AppBankCurrency")
	proto.RegisterType((*AllocateGas)(nil), "developer.AllocateGas")
	proto.RegisterType((*UnAllocateGas)(nil), "developer.UnAllocateGas")
	proto.RegisterType((*AppBankBalance)(nil), "developer.AppBankBalance")
	proto.RegisterType((*BankTx)(nil), "developer.BankTx")
	proto.RegisterType((*BankAddress)(nil), "developer.BankAddress")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeveloperServiceClient is the client API for DeveloperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeveloperServiceClient interface {
	// developer init session
	InitSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*Session, error)
	// get balance
	GetBalance(ctx context.Context, in *Session, opts ...grpc.CallOption) (*AppBankBalance, error)
	// send coin from bank
	SendTxFromBank(ctx context.Context, in *BankTx, opts ...grpc.CallOption) (*BankTx, error)
}

type developerServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeveloperServiceClient(cc *grpc.ClientConn) DeveloperServiceClient {
	return &developerServiceClient{cc}
}

func (c *developerServiceClient) InitSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/developer.DeveloperService/initSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) GetBalance(ctx context.Context, in *Session, opts ...grpc.CallOption) (*AppBankBalance, error) {
	out := new(AppBankBalance)
	err := c.cc.Invoke(ctx, "/developer.DeveloperService/getBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) SendTxFromBank(ctx context.Context, in *BankTx, opts ...grpc.CallOption) (*BankTx, error) {
	out := new(BankTx)
	err := c.cc.Invoke(ctx, "/developer.DeveloperService/sendTxFromBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeveloperServiceServer is the server API for DeveloperService service.
type DeveloperServiceServer interface {
	// developer init session
	InitSession(context.Context, *SessionRequest) (*Session, error)
	// get balance
	GetBalance(context.Context, *Session) (*AppBankBalance, error)
	// send coin from bank
	SendTxFromBank(context.Context, *BankTx) (*BankTx, error)
}

func RegisterDeveloperServiceServer(s *grpc.Server, srv DeveloperServiceServer) {
	s.RegisterService(&_DeveloperService_serviceDesc, srv)
}

func _DeveloperService_InitSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).InitSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/developer.DeveloperService/InitSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).InitSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/developer.DeveloperService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).GetBalance(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_SendTxFromBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).SendTxFromBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/developer.DeveloperService/SendTxFromBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).SendTxFromBank(ctx, req.(*BankTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeveloperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "developer.DeveloperService",
	HandlerType: (*DeveloperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "initSession",
			Handler:    _DeveloperService_InitSession_Handler,
		},
		{
			MethodName: "getBalance",
			Handler:    _DeveloperService_GetBalance_Handler,
		},
		{
			MethodName: "sendTxFromBank",
			Handler:    _DeveloperService_SendTxFromBank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "developer.proto",
}

func init() { proto.RegisterFile("developer.proto", fileDescriptor_6300c5df2ba41f00) }

var fileDescriptor_6300c5df2ba41f00 = []byte{
	// 943 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xce, 0xd8, 0x89, 0xed, 0x29, 0xc7, 0xce, 0xd2, 0x5a, 0xd8, 0x89, 0x23, 0x20, 0x18, 0x24,
	0x72, 0x89, 0x2d, 0x2d, 0x17, 0x2e, 0x20, 0x39, 0x41, 0x20, 0x6b, 0xa5, 0x28, 0x9a, 0x24, 0x17,
	0x24, 0x34, 0x6a, 0x4f, 0xd7, 0x5a, 0xad, 0xcc, 0x4f, 0xd3, 0xd3, 0x36, 0xf6, 0x0b, 0x70, 0x42,
	0x3c, 0xc4, 0x1e, 0x79, 0x03, 0x9e, 0x81, 0xb7, 0xe1, 0x09, 0x50, 0xff, 0xcc, 0xc4, 0x9e, 0x09,
	0x68, 0x8f, 0xdc, 0x5c, 0xdf, 0x57, 0x53, 0xee, 0xfa, 0xea, 0x9b, 0xea, 0x81, 0x13, 0x86, 0x6b,
	0x4c, 0x72, 0x81, 0x72, 0x22, 0x64, 0xae, 0x72, 0xe2, 0x57, 0xc0, 0xe8, 0x6c, 0x99, 0xe7, 0xcb,
	0x04, 0xa7, 0x86, 0x58, 0xac, 0xde, 0x4e, 0x31, 0x15, 0x6a, 0x6b, 0xf3, 0xc6, 0x7f, 0xb7, 0xc0,
	0xff, 0xae, 0x4c, 0x25, 0x43, 0x68, 0x71, 0x16, 0x78, 0xe7, 0xde, 0x85, 0x1f, 0xb6, 0x38, 0x23,
	0xaf, 0xa0, 0xbb, 0x2a, 0x50, 0x46, 0x9c, 0x05, 0x2d, 0x03, 0x76, 0x74, 0x38, 0x67, 0xe4, 0x23,
	0xe8, 0xa4, 0xf9, 0x82, 0x27, 0x18, 0xb4, 0x2d, 0x6e, 0x23, 0xf2, 0x12, 0x8e, 0x30, 0xa5, 0x3c,
	0x09, 0x0e, 0x0d, 0x6c, 0x03, 0xf2, 0x05, 0x0c, 0x39, 0x8b, 0x62, 0x2a, 0x59, 0x94, 0xad, 0xd2,
	0x05, 0xca, 0xe0, 0xc8, 0xd0, 0xc7, 0x9c, 0x5d, 0x53, 0xc9, 0x6e, 0x0c, 0x46, 0xce, 0xc0, 0x97,
	0x48, 0x93, 0x28, 0xa3, 0x29, 0x06, 0x1d, 0x93, 0xd0, 0xd3, 0xc0, 0x0d, 0x4d, 0x91, 0x7c, 0x0a,
	0x7d, 0x21, 0xf9, 0x9a, 0x2a, 0x8c, 0x1e, 0x71, 0x1b, 0x74, 0x0d, 0x0d, 0x0e, 0x7a, 0x83, 0x5b,
	0xf2, 0x31, 0x80, 0x58, 0x2d, 0x12, 0x1e, 0x1b, 0xbe, 0x67, 0x78, 0xdf, 0x22, 0x8e, 0xa6, 0x71,
	0x8c, 0x45, 0x61, 0x68, 0xdf, 0xd2, 0x16, 0xd1, 0xf4, 0x19, 0xf8, 0xb1, 0x44, 0x5d, 0x9d, 0xaa,
	0x00, 0xce, 0xbd, 0x8b, 0xc3, 0xb0, 0x67, 0x81, 0x99, 0xd2, 0xe4, 0x4a, 0x30, 0x47, 0xf6, 0x2d,
	0x69, 0x81, 0x99, 0xd2, 0x4a, 0xd0, 0x58, 0xf1, 0x35, 0x06, 0xc7, 0xe7, 0xde, 0x45, 0x2f, 0x74,
	0x11, 0x19, 0x41, 0x6f, 0x8d, 0x92, 0xbf, 0xe5, 0xc8, 0x82, 0x81, 0x61, 0xaa, 0x78, 0xfc, 0xce,
	0x83, 0xee, 0x1d, 0x16, 0x05, 0xcf, 0x33, 0x7d, 0xb0, 0xc2, 0xfe, 0x8c, 0x2a, 0xe9, 0x7d, 0x87,
	0xcc, 0x19, 0xf9, 0x0c, 0x8e, 0xab, 0x49, 0x3e, 0x8d, 0xa1, 0x5f, 0x61, 0x73, 0x66, 0x2b, 0xc4,
	0x12, 0x95, 0x69, 0xad, 0x5d, 0x56, 0xd0, 0x88, 0x6e, 0xed, 0x43, 0xe8, 0x50, 0x21, 0xf4, 0xb3,
	0x6e, 0x26, 0x54, 0x88, 0x39, 0xd3, 0x82, 0xc6, 0x72, 0x2b, 0x54, 0x1e, 0xa9, 0xad, 0x40, 0x37,
	0x10, 0xb0, 0xd0, 0xfd, 0x56, 0xe0, 0xf8, 0x16, 0x86, 0xee, 0x8c, 0x21, 0xfe, 0xbc, 0xc2, 0x42,
	0xd5, 0x34, 0xf4, 0xea, 0x1a, 0xd6, 0x2a, 0xb6, 0x1a, 0x15, 0x7f, 0x82, 0xee, 0x4c, 0x88, 0x2b,
	0x9a, 0x3d, 0x36, 0x8c, 0xf6, 0x1e, 0x6d, 0xd6, 0xca, 0xb7, 0x1b, 0xe5, 0xff, 0x6c, 0xc1, 0x89,
	0xab, 0x7f, 0xbd, 0x92, 0x12, 0xb3, 0x78, 0xfb, 0x9c, 0xa1, 0x17, 0x34, 0x7b, 0xdc, 0x31, 0xb4,
	0x0e, 0xe7, 0xcc, 0x18, 0x20, 0xe7, 0xd9, 0x6e, 0xed, 0x9e, 0x06, 0x74, 0x65, 0x12, 0xe8, 0xa7,
	0x12, 0x9a, 0xc5, 0xe8, 0x34, 0x2c, 0x43, 0x72, 0x0a, 0xbd, 0x94, 0xf2, 0x2c, 0xca, 0x50, 0x39,
	0x09, 0xbb, 0x3a, 0xbe, 0x41, 0xa5, 0xa9, 0x25, 0x2d, 0x22, 0x49, 0x95, 0x75, 0xb3, 0x17, 0x76,
	0x97, 0xb4, 0x08, 0xa9, 0x42, 0xf2, 0x02, 0xda, 0x4b, 0x5a, 0x18, 0x13, 0x0f, 0x42, 0xfd, 0x93,
	0x7c, 0x09, 0x27, 0x94, 0x31, 0xa9, 0xb5, 0x2d, 0xff, 0xc9, 0x5a, 0x78, 0xe8, 0xe0, 0x2b, 0xf7,
	0x87, 0x01, 0x74, 0x1d, 0xe2, 0x4c, 0x5c, 0x86, 0xda, 0x70, 0x0c, 0x63, 0x9e, 0xd2, 0xa4, 0x30,
	0x0e, 0x1e, 0x84, 0x55, 0x5c, 0xd7, 0xae, 0xdf, 0xd0, 0xee, 0x37, 0x0f, 0xfa, 0xb3, 0x24, 0xc9,
	0x63, 0xaa, 0xf0, 0x07, 0x5a, 0x34, 0x74, 0xdb, 0xed, 0xb3, 0xf5, 0x6c, 0x9f, 0x3b, 0xc2, 0xe9,
	0x3e, 0x4b, 0xdd, 0xdc, 0x11, 0x8c, 0x6e, 0x83, 0xb0, 0x0c, 0x6b, 0xae, 0x3f, 0xaa, 0xb9, 0x7e,
	0xfc, 0xbb, 0x07, 0x83, 0x87, 0xec, 0x7f, 0x74, 0xa0, 0x07, 0x18, 0x3a, 0x6b, 0x95, 0x83, 0xb8,
	0x86, 0x13, 0xe3, 0xa4, 0xd8, 0x5a, 0x8d, 0x63, 0x11, 0x78, 0xe7, 0xed, 0x8b, 0xfe, 0xeb, 0xd1,
	0xe4, 0x69, 0x17, 0xd7, 0xec, 0x18, 0x0e, 0x17, 0x4f, 0x11, 0xc7, 0x62, 0xfc, 0x6b, 0x1b, 0x3a,
	0x3a, 0xe1, 0x7e, 0xf3, 0xfe, 0x4e, 0x7d, 0x05, 0x5d, 0xb5, 0xd9, 0xed, 0xae, 0xa3, 0x36, 0xa6,
	0xb9, 0x7f, 0x79, 0xd1, 0x3f, 0x81, 0xbe, 0x86, 0xcb, 0x3d, 0xee, 0x5a, 0xa3, 0x42, 0x3c, 0xd8,
	0x55, 0xbe, 0xe7, 0xfc, 0x4e, 0xcd, 0xf9, 0x7a, 0xbb, 0xa5, 0xf9, 0x2a, 0x53, 0x6e, 0xe3, 0xba,
	0x68, 0x4f, 0xfe, 0x5e, 0x43, 0x7e, 0x25, 0x29, 0x33, 0xcb, 0xd2, 0x37, 0xcb, 0xb2, 0x6b, 0xe2,
	0x99, 0xaa, 0x89, 0x0c, 0xf5, 0x5d, 0xa7, 0x5d, 0xea, 0x94, 0xd2, 0x7c, 0xe9, 0x52, 0x07, 0xcd,
	0xd9, 0x9e, 0xc5, 0x8f, 0x6b, 0x16, 0x3f, 0x03, 0x5f, 0x6d, 0xa2, 0x42, 0x51, 0xb5, 0x2a, 0xcc,
	0xc2, 0x1d, 0x84, 0x3d, 0xb5, 0xb9, 0x33, 0x71, 0xdd, 0xff, 0xc3, 0x86, 0xff, 0xff, 0xf0, 0xa0,
	0xaf, 0x07, 0x31, 0x73, 0x2f, 0x53, 0x7d, 0x1a, 0x3b, 0xaf, 0x5d, 0x6b, 0xff, 0xb5, 0x7b, 0x09,
	0x47, 0xf9, 0x2f, 0x19, 0x4a, 0x37, 0x0c, 0x1b, 0xd4, 0xaf, 0xab, 0xc3, 0xc6, 0x75, 0xf5, 0x1f,
	0x8b, 0xa3, 0x76, 0xd8, 0x4e, 0xfd, 0xb0, 0xaf, 0xff, 0xf2, 0xe0, 0x45, 0x75, 0x67, 0xdf, 0xa1,
	0x5c, 0xf3, 0x18, 0xc9, 0xb7, 0xd0, 0xe7, 0x19, 0x57, 0xe5, 0xb5, 0x72, 0xba, 0xe3, 0xc2, 0xfd,
	0x35, 0x3e, 0x22, 0x4d, 0x6a, 0x7c, 0x40, 0xbe, 0x01, 0x58, 0xa2, 0x2a, 0xdd, 0xfd, 0x4c, 0xce,
	0xe8, 0xb4, 0x69, 0x6c, 0x97, 0x3e, 0x3e, 0x20, 0x5f, 0xc3, 0xb0, 0xc0, 0x8c, 0xdd, 0x6f, 0xbe,
	0x97, 0x79, 0x6a, 0x56, 0xfc, 0x07, 0x3b, 0xe9, 0xd6, 0xe3, 0xa3, 0x26, 0x34, 0x3e, 0xb8, 0xca,
	0xe0, 0xf3, 0x38, 0x4f, 0x27, 0xf4, 0x11, 0xb7, 0x13, 0x2a, 0xf8, 0x84, 0x0a, 0xa1, 0x9d, 0x3e,
	0x59, 0x4a, 0x11, 0x3f, 0xe5, 0x5f, 0x0d, 0xab, 0x8e, 0x6f, 0xf5, 0x87, 0xcb, 0xad, 0xf7, 0xa3,
	0x79, 0xe4, 0x92, 0x0a, 0x71, 0xa9, 0xf3, 0x2f, 0x71, 0x43, 0x53, 0x91, 0x60, 0x31, 0xd5, 0x0f,
	0x4e, 0x4b, 0x78, 0x5a, 0x55, 0x78, 0xd7, 0x3a, 0x9c, 0xbd, 0xc1, 0xed, 0xa2, 0x63, 0x3e, 0x7c,
	0xbe, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x24, 0xee, 0x4e, 0xe9, 0x33, 0x09, 0x00, 0x00,
}
