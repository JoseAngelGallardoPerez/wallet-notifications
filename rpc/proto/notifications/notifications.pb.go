// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/proto/notifications/notifications.proto

package notifications

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

type SettingsRequest struct {
	SettingNames         []string `protobuf:"bytes,1,rep,name=settingNames,proto3" json:"settingNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingsRequest) Reset()         { *m = SettingsRequest{} }
func (m *SettingsRequest) String() string { return proto.CompactTextString(m) }
func (*SettingsRequest) ProtoMessage()    {}
func (*SettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{0}
}

func (m *SettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsRequest.Unmarshal(m, b)
}
func (m *SettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsRequest.Marshal(b, m, deterministic)
}
func (m *SettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsRequest.Merge(m, src)
}
func (m *SettingsRequest) XXX_Size() int {
	return xxx_messageInfo_SettingsRequest.Size(m)
}
func (m *SettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsRequest proto.InternalMessageInfo

func (m *SettingsRequest) GetSettingNames() []string {
	if m != nil {
		return m.SettingNames
	}
	return nil
}

type SettingsResponse struct {
	Settings             []*Setting `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SettingsResponse) Reset()         { *m = SettingsResponse{} }
func (m *SettingsResponse) String() string { return proto.CompactTextString(m) }
func (*SettingsResponse) ProtoMessage()    {}
func (*SettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{1}
}

func (m *SettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsResponse.Unmarshal(m, b)
}
func (m *SettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsResponse.Marshal(b, m, deterministic)
}
func (m *SettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsResponse.Merge(m, src)
}
func (m *SettingsResponse) XXX_Size() int {
	return xxx_messageInfo_SettingsResponse.Size(m)
}
func (m *SettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsResponse proto.InternalMessageInfo

func (m *SettingsResponse) GetSettings() []*Setting {
	if m != nil {
		return m.Settings
	}
	return nil
}

type Setting struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{2}
}

func (m *Setting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting.Unmarshal(m, b)
}
func (m *Setting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting.Marshal(b, m, deterministic)
}
func (m *Setting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting.Merge(m, src)
}
func (m *Setting) XXX_Size() int {
	return xxx_messageInfo_Setting.Size(m)
}
func (m *Setting) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting.DiscardUnknown(m)
}

var xxx_messageInfo_Setting proto.InternalMessageInfo

func (m *Setting) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Setting) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UserSettingsRequest struct {
	NotificationName     string   `protobuf:"bytes,1,opt,name=notificationName,proto3" json:"notificationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSettingsRequest) Reset()         { *m = UserSettingsRequest{} }
func (m *UserSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*UserSettingsRequest) ProtoMessage()    {}
func (*UserSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{3}
}

func (m *UserSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettingsRequest.Unmarshal(m, b)
}
func (m *UserSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettingsRequest.Marshal(b, m, deterministic)
}
func (m *UserSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingsRequest.Merge(m, src)
}
func (m *UserSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_UserSettingsRequest.Size(m)
}
func (m *UserSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingsRequest proto.InternalMessageInfo

func (m *UserSettingsRequest) GetNotificationName() string {
	if m != nil {
		return m.NotificationName
	}
	return ""
}

type UserSettingsResponse struct {
	UserSettings         []*UsersSetting `protobuf:"bytes,1,rep,name=userSettings,proto3" json:"userSettings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserSettingsResponse) Reset()         { *m = UserSettingsResponse{} }
func (m *UserSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*UserSettingsResponse) ProtoMessage()    {}
func (*UserSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{4}
}

func (m *UserSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettingsResponse.Unmarshal(m, b)
}
func (m *UserSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettingsResponse.Marshal(b, m, deterministic)
}
func (m *UserSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingsResponse.Merge(m, src)
}
func (m *UserSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_UserSettingsResponse.Size(m)
}
func (m *UserSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingsResponse proto.InternalMessageInfo

func (m *UserSettingsResponse) GetUserSettings() []*UsersSetting {
	if m != nil {
		return m.UserSettings
	}
	return nil
}

type UsersSetting struct {
	NotificationName     string   `protobuf:"bytes,1,opt,name=notificationName,proto3" json:"notificationName,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersSetting) Reset()         { *m = UsersSetting{} }
func (m *UsersSetting) String() string { return proto.CompactTextString(m) }
func (*UsersSetting) ProtoMessage()    {}
func (*UsersSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{5}
}

func (m *UsersSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersSetting.Unmarshal(m, b)
}
func (m *UsersSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersSetting.Marshal(b, m, deterministic)
}
func (m *UsersSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersSetting.Merge(m, src)
}
func (m *UsersSetting) XXX_Size() int {
	return xxx_messageInfo_UsersSetting.Size(m)
}
func (m *UsersSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersSetting.DiscardUnknown(m)
}

var xxx_messageInfo_UsersSetting proto.InternalMessageInfo

func (m *UsersSetting) GetNotificationName() string {
	if m != nil {
		return m.NotificationName
	}
	return ""
}

func (m *UsersSetting) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type Request struct {
	To                   string        `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	EventName            string        `protobuf:"bytes,2,opt,name=eventName,proto3" json:"eventName,omitempty"`
	TemplateData         *TemplateData `protobuf:"bytes,3,opt,name=templateData,proto3" json:"templateData,omitempty"`
	Notifiers            []string      `protobuf:"bytes,4,rep,name=notifiers,proto3" json:"notifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{6}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Request) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Request) GetTemplateData() *TemplateData {
	if m != nil {
		return m.TemplateData
	}
	return nil
}

func (m *Request) GetNotifiers() []string {
	if m != nil {
		return m.Notifiers
	}
	return nil
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{7}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Details              string   `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{8}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Error) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type TemplateData struct {
	UserName                       string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	FirstName                      string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName                       string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	SiteName                       string   `protobuf:"bytes,4,opt,name=siteName,proto3" json:"siteName,omitempty"`
	SiteLoginUrl                   string   `protobuf:"bytes,5,opt,name=siteLoginUrl,proto3" json:"siteLoginUrl,omitempty"`
	Logo                           string   `protobuf:"bytes,6,opt,name=logo,proto3" json:"logo,omitempty"`
	OneTimeLoginUrl                string   `protobuf:"bytes,7,opt,name=oneTimeLoginUrl,proto3" json:"oneTimeLoginUrl,omitempty"`
	PrivateMessageRecipient        string   `protobuf:"bytes,8,opt,name=privateMessageRecipient,proto3" json:"privateMessageRecipient,omitempty"`
	PrivateMessageAuthor           string   `protobuf:"bytes,9,opt,name=privateMessageAuthor,proto3" json:"privateMessageAuthor,omitempty"`
	PrivateMessageUrl              string   `protobuf:"bytes,10,opt,name=privateMessageUrl,proto3" json:"privateMessageUrl,omitempty"`
	PrivateMessageRecipientEditUrl string   `protobuf:"bytes,11,opt,name=privateMessageRecipientEditUrl,proto3" json:"privateMessageRecipientEditUrl,omitempty"`
	Reason                         string   `protobuf:"bytes,12,opt,name=reason,proto3" json:"reason,omitempty"`
	Link                           string   `protobuf:"bytes,13,opt,name=link,proto3" json:"link,omitempty"`
	DocumentName                   string   `protobuf:"bytes,14,opt,name=documentName,proto3" json:"documentName,omitempty"`
	Tan                            string   `protobuf:"bytes,15,opt,name=tan,proto3" json:"tan,omitempty"`
	SiteUrl                        string   `protobuf:"bytes,16,opt,name=siteUrl,proto3" json:"siteUrl,omitempty"`
	Password                       string   `protobuf:"bytes,17,opt,name=password,proto3" json:"password,omitempty"`
	EntityType                     string   `protobuf:"bytes,18,opt,name=entityType,proto3" json:"entityType,omitempty"`
	EntityID                       uint64   `protobuf:"varint,19,opt,name=entityID,proto3" json:"entityID,omitempty"`
	MessageUnreadedCount           uint64   `protobuf:"varint,20,opt,name=messageUnreadedCount,proto3" json:"messageUnreadedCount,omitempty"`
	SenderID                       string   `protobuf:"bytes,21,opt,name=senderID,proto3" json:"senderID,omitempty"`
	VerificationLink               string   `protobuf:"bytes,22,opt,name=verificationLink,proto3" json:"verificationLink,omitempty"`
	AccountNumber                  string   `protobuf:"bytes,23,opt,name=accountNumber,proto3" json:"accountNumber,omitempty"`
	TransactionId                  uint64   `protobuf:"varint,24,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	ConfirmationCode               string   `protobuf:"bytes,25,opt,name=confirmationCode,proto3" json:"confirmationCode,omitempty"`
	SetPasswordConfirmationCode    string   `protobuf:"bytes,26,opt,name=setPasswordConfirmationCode,proto3" json:"setPasswordConfirmationCode,omitempty"`
	RequestId                      uint64   `protobuf:"varint,27,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Count                          uint64   `protobuf:"varint,28,opt,name=count,proto3" json:"count,omitempty"`
	InvoiceID                      string   `protobuf:"bytes,29,opt,name=invoiceID,proto3" json:"invoiceID,omitempty"`
	SupplierCompany                string   `protobuf:"bytes,30,opt,name=supplierCompany,proto3" json:"supplierCompany,omitempty"`
	FunderCompany                  string   `protobuf:"bytes,31,opt,name=funderCompany,proto3" json:"funderCompany,omitempty"`
	Date                           string   `protobuf:"bytes,32,opt,name=date,proto3" json:"date,omitempty"`
	PlatformAdmin                  string   `protobuf:"bytes,33,opt,name=platformAdmin,proto3" json:"platformAdmin,omitempty"`
	StaffFirstName                 string   `protobuf:"bytes,34,opt,name=staffFirstName,proto3" json:"staffFirstName,omitempty"`
	OwnerFirstName                 string   `protobuf:"bytes,35,opt,name=ownerFirstName,proto3" json:"ownerFirstName,omitempty"`
	OwnerLastName                  string   `protobuf:"bytes,36,opt,name=ownerLastName,proto3" json:"ownerLastName,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *TemplateData) Reset()         { *m = TemplateData{} }
func (m *TemplateData) String() string { return proto.CompactTextString(m) }
func (*TemplateData) ProtoMessage()    {}
func (*TemplateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_05cc8845bd7b2e40, []int{9}
}

func (m *TemplateData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateData.Unmarshal(m, b)
}
func (m *TemplateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateData.Marshal(b, m, deterministic)
}
func (m *TemplateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateData.Merge(m, src)
}
func (m *TemplateData) XXX_Size() int {
	return xxx_messageInfo_TemplateData.Size(m)
}
func (m *TemplateData) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateData.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateData proto.InternalMessageInfo

func (m *TemplateData) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *TemplateData) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *TemplateData) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *TemplateData) GetSiteName() string {
	if m != nil {
		return m.SiteName
	}
	return ""
}

func (m *TemplateData) GetSiteLoginUrl() string {
	if m != nil {
		return m.SiteLoginUrl
	}
	return ""
}

func (m *TemplateData) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *TemplateData) GetOneTimeLoginUrl() string {
	if m != nil {
		return m.OneTimeLoginUrl
	}
	return ""
}

func (m *TemplateData) GetPrivateMessageRecipient() string {
	if m != nil {
		return m.PrivateMessageRecipient
	}
	return ""
}

func (m *TemplateData) GetPrivateMessageAuthor() string {
	if m != nil {
		return m.PrivateMessageAuthor
	}
	return ""
}

func (m *TemplateData) GetPrivateMessageUrl() string {
	if m != nil {
		return m.PrivateMessageUrl
	}
	return ""
}

func (m *TemplateData) GetPrivateMessageRecipientEditUrl() string {
	if m != nil {
		return m.PrivateMessageRecipientEditUrl
	}
	return ""
}

func (m *TemplateData) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TemplateData) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *TemplateData) GetDocumentName() string {
	if m != nil {
		return m.DocumentName
	}
	return ""
}

func (m *TemplateData) GetTan() string {
	if m != nil {
		return m.Tan
	}
	return ""
}

func (m *TemplateData) GetSiteUrl() string {
	if m != nil {
		return m.SiteUrl
	}
	return ""
}

func (m *TemplateData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *TemplateData) GetEntityType() string {
	if m != nil {
		return m.EntityType
	}
	return ""
}

func (m *TemplateData) GetEntityID() uint64 {
	if m != nil {
		return m.EntityID
	}
	return 0
}

func (m *TemplateData) GetMessageUnreadedCount() uint64 {
	if m != nil {
		return m.MessageUnreadedCount
	}
	return 0
}

func (m *TemplateData) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *TemplateData) GetVerificationLink() string {
	if m != nil {
		return m.VerificationLink
	}
	return ""
}

func (m *TemplateData) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

func (m *TemplateData) GetTransactionId() uint64 {
	if m != nil {
		return m.TransactionId
	}
	return 0
}

func (m *TemplateData) GetConfirmationCode() string {
	if m != nil {
		return m.ConfirmationCode
	}
	return ""
}

func (m *TemplateData) GetSetPasswordConfirmationCode() string {
	if m != nil {
		return m.SetPasswordConfirmationCode
	}
	return ""
}

func (m *TemplateData) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *TemplateData) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TemplateData) GetInvoiceID() string {
	if m != nil {
		return m.InvoiceID
	}
	return ""
}

func (m *TemplateData) GetSupplierCompany() string {
	if m != nil {
		return m.SupplierCompany
	}
	return ""
}

func (m *TemplateData) GetFunderCompany() string {
	if m != nil {
		return m.FunderCompany
	}
	return ""
}

func (m *TemplateData) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TemplateData) GetPlatformAdmin() string {
	if m != nil {
		return m.PlatformAdmin
	}
	return ""
}

func (m *TemplateData) GetStaffFirstName() string {
	if m != nil {
		return m.StaffFirstName
	}
	return ""
}

func (m *TemplateData) GetOwnerFirstName() string {
	if m != nil {
		return m.OwnerFirstName
	}
	return ""
}

func (m *TemplateData) GetOwnerLastName() string {
	if m != nil {
		return m.OwnerLastName
	}
	return ""
}

func init() {
	proto.RegisterType((*SettingsRequest)(nil), "velmie.wallet.notifications.SettingsRequest")
	proto.RegisterType((*SettingsResponse)(nil), "velmie.wallet.notifications.SettingsResponse")
	proto.RegisterType((*Setting)(nil), "velmie.wallet.notifications.Setting")
	proto.RegisterType((*UserSettingsRequest)(nil), "velmie.wallet.notifications.UserSettingsRequest")
	proto.RegisterType((*UserSettingsResponse)(nil), "velmie.wallet.notifications.UserSettingsResponse")
	proto.RegisterType((*UsersSetting)(nil), "velmie.wallet.notifications.UsersSetting")
	proto.RegisterType((*Request)(nil), "velmie.wallet.notifications.Request")
	proto.RegisterType((*Response)(nil), "velmie.wallet.notifications.Response")
	proto.RegisterType((*Error)(nil), "velmie.wallet.notifications.Error")
	proto.RegisterType((*TemplateData)(nil), "velmie.wallet.notifications.templateData")
}

func init() {
	proto.RegisterFile("rpc/proto/notifications/notifications.proto", fileDescriptor_05cc8845bd7b2e40)
}

var fileDescriptor_05cc8845bd7b2e40 = []byte{
	// 947 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x6d, 0x6f, 0x1c, 0x35,
	0x10, 0x56, 0xde, 0x93, 0xc9, 0xcb, 0xa5, 0x4e, 0x68, 0x4d, 0x52, 0x42, 0x59, 0x02, 0x2a, 0x50,
	0xae, 0x90, 0x0a, 0xc1, 0x47, 0x42, 0xd2, 0x42, 0xa5, 0x34, 0x42, 0x47, 0x2a, 0x24, 0xc4, 0x17,
	0x77, 0xd7, 0x97, 0x1a, 0x76, 0xed, 0xc5, 0xf6, 0x5e, 0x94, 0xbf, 0x83, 0xf8, 0x45, 0xfc, 0xa2,
	0x7a, 0x6c, 0xef, 0xdd, 0xfa, 0xda, 0x5e, 0xda, 0x6f, 0x3b, 0xcf, 0xcc, 0x33, 0xe3, 0x19, 0xcf,
	0x8c, 0x17, 0xbe, 0xd2, 0x75, 0xfe, 0xb0, 0xd6, 0xca, 0xaa, 0x87, 0x52, 0x59, 0x31, 0x14, 0x39,
	0xb3, 0x42, 0x49, 0x93, 0x4a, 0x7d, 0x6f, 0x41, 0xf6, 0x47, 0xbc, 0xac, 0x04, 0xef, 0x5f, 0xb1,
	0xb2, 0xe4, 0xb6, 0x9f, 0x98, 0x64, 0xdf, 0x41, 0xef, 0x37, 0x6e, 0xad, 0x90, 0x97, 0x66, 0xc0,
	0xff, 0x69, 0xb8, 0xb1, 0x24, 0x83, 0x0d, 0x13, 0xa0, 0x73, 0x56, 0x71, 0x43, 0xe7, 0xee, 0x2d,
	0xdc, 0x5f, 0x1b, 0x24, 0x58, 0x76, 0x01, 0xdb, 0x13, 0x9a, 0xa9, 0x9d, 0x27, 0x4e, 0x7e, 0x84,
	0xd5, 0x68, 0x13, 0x38, 0xeb, 0x47, 0x87, 0xfd, 0x19, 0xa1, 0xfb, 0xd1, 0xc1, 0x60, 0xcc, 0xca,
	0x1e, 0xc1, 0x4a, 0x04, 0x09, 0x81, 0x45, 0xe9, 0x22, 0x39, 0x47, 0x73, 0x2e, 0xb8, 0xff, 0x26,
	0xbb, 0xb0, 0x34, 0x62, 0x65, 0xc3, 0xe9, 0xbc, 0x07, 0x83, 0x90, 0x1d, 0xc3, 0xce, 0x73, 0xc3,
	0xf5, 0x74, 0x16, 0x5f, 0xc2, 0x76, 0x37, 0xdc, 0xf9, 0xc4, 0xd9, 0x6b, 0x78, 0xc6, 0x61, 0x37,
	0x75, 0x11, 0x33, 0x7a, 0x06, 0x1b, 0x4d, 0x07, 0x8f, 0x59, 0x7d, 0x31, 0x33, 0x2b, 0x74, 0x64,
	0xda, 0xd4, 0x12, 0x7a, 0x76, 0x06, 0x1b, 0x5d, 0xed, 0xfb, 0x1c, 0x91, 0x6c, 0xc3, 0x42, 0x23,
	0x8a, 0x98, 0x39, 0x7e, 0x66, 0xff, 0xcd, 0xc1, 0x4a, 0x9b, 0xec, 0x16, 0xcc, 0x5b, 0x15, 0xb9,
	0xee, 0x8b, 0xdc, 0x85, 0x35, 0x3e, 0xe2, 0xd2, 0x7a, 0x97, 0x81, 0x33, 0x01, 0x30, 0x2d, 0xcb,
	0xab, 0xba, 0x64, 0x96, 0x9f, 0x32, 0xcb, 0xe8, 0x82, 0x33, 0xb8, 0x29, 0xad, 0x2e, 0x61, 0x90,
	0xd0, 0x31, 0x58, 0xb0, 0x75, 0xa9, 0xd1, 0x45, 0xdf, 0x2c, 0x13, 0x20, 0xfb, 0x13, 0x56, 0xc7,
	0xf5, 0xbc, 0x0d, 0xcb, 0xc6, 0x32, 0xdb, 0x98, 0x78, 0xd4, 0x28, 0x91, 0x1f, 0x60, 0x89, 0x6b,
	0xad, 0xb4, 0x3f, 0xea, 0xfa, 0x51, 0x36, 0xf3, 0x24, 0x8f, 0xd1, 0x72, 0x10, 0x08, 0xd9, 0xf7,
	0xb0, 0xe4, 0x65, 0xec, 0x0d, 0x2b, 0x6c, 0xd9, 0x16, 0x30, 0x08, 0x84, 0xc2, 0x4a, 0xc1, 0x2d,
	0x13, 0xa5, 0x89, 0x55, 0x68, 0xc5, 0xec, 0x5f, 0x48, 0x8b, 0x40, 0xf6, 0x60, 0x15, 0x2f, 0xab,
	0x73, 0x09, 0x63, 0x19, 0x33, 0x1c, 0x0a, 0x6d, 0x92, 0x72, 0x8e, 0x01, 0x64, 0x96, 0x2c, 0x2a,
	0x17, 0x02, 0xb3, 0x95, 0x51, 0x67, 0x84, 0xe5, 0x5e, 0xb7, 0x18, 0x74, 0xad, 0xec, 0xe7, 0xcc,
	0x7d, 0x9f, 0xa9, 0x4b, 0x21, 0x9f, 0xeb, 0x92, 0x2e, 0x79, 0x7d, 0x82, 0xe1, 0x18, 0x94, 0xea,
	0x52, 0xd1, 0xe5, 0x30, 0x06, 0xf8, 0x4d, 0xee, 0x43, 0x4f, 0x49, 0x7e, 0x21, 0xaa, 0x09, 0x75,
	0xc5, 0xab, 0xa7, 0x61, 0x57, 0xd7, 0x3b, 0xb5, 0x16, 0x23, 0x97, 0xe2, 0x33, 0x6e, 0x0c, 0xbb,
	0xe4, 0x03, 0x9e, 0x8b, 0x5a, 0xb8, 0x3e, 0xa0, 0xab, 0x9e, 0xf1, 0x36, 0x35, 0x39, 0x82, 0xdd,
	0x54, 0x75, 0xdc, 0xd8, 0x97, 0xee, 0x82, 0xd6, 0x3c, 0xed, 0x8d, 0x3a, 0xf2, 0x00, 0x6e, 0xa5,
	0x38, 0x9e, 0x0c, 0x3c, 0xe1, 0x75, 0x05, 0x79, 0x02, 0x07, 0x6f, 0x09, 0xfe, 0xb8, 0x10, 0x16,
	0xa9, 0xeb, 0x9e, 0x7a, 0x83, 0x15, 0xf6, 0x94, 0xe6, 0xcc, 0x28, 0x49, 0x37, 0x42, 0x4f, 0x05,
	0xc9, 0x57, 0x4e, 0xc8, 0xbf, 0xe9, 0x66, 0xac, 0x9c, 0xfb, 0xc6, 0x8a, 0x17, 0x2a, 0x6f, 0xaa,
	0x76, 0x32, 0xb6, 0x42, 0xc5, 0xbb, 0x18, 0x0e, 0x9a, 0x65, 0x92, 0xf6, 0xc2, 0xa0, 0xb9, 0x4f,
	0x6c, 0x22, 0xbc, 0x13, 0x3c, 0xd2, 0x76, 0x68, 0xa2, 0x28, 0xe2, 0xed, 0xd6, 0xcc, 0x98, 0x2b,
	0xa5, 0x0b, 0x7a, 0x2b, 0xdc, 0x6e, 0x2b, 0x93, 0x03, 0x00, 0xe7, 0x52, 0xd8, 0xeb, 0x8b, 0xeb,
	0x9a, 0x53, 0xe2, 0xb5, 0x1d, 0x04, 0xb9, 0x41, 0x7a, 0x7a, 0x4a, 0x77, 0x9c, 0x76, 0x71, 0x30,
	0x96, 0xb1, 0xfa, 0x55, 0xac, 0x94, 0x74, 0xe9, 0x14, 0xbc, 0x38, 0x51, 0x8d, 0xbb, 0xb4, 0x5d,
	0x6f, 0xf7, 0x46, 0x9d, 0xef, 0x34, 0x2e, 0x0b, 0xae, 0x9d, 0xbf, 0x0f, 0x62, 0xa7, 0x45, 0x19,
	0x17, 0xcd, 0x88, 0xeb, 0xf1, 0x0c, 0x9d, 0x61, 0x5d, 0x6e, 0x87, 0x45, 0x33, 0x8d, 0x93, 0x43,
	0xd8, 0x64, 0x79, 0x8e, 0x2e, 0xcf, 0x9b, 0xea, 0x05, 0xd7, 0xf4, 0x8e, 0x37, 0x4c, 0x41, 0xb4,
	0xb2, 0x9a, 0x49, 0xc3, 0x72, 0x24, 0x3e, 0x2d, 0x28, 0xf5, 0x47, 0x4b, 0x41, 0x8c, 0x9b, 0x2b,
	0xe9, 0x26, 0xa5, 0xf2, 0xfe, 0x4f, 0x54, 0xc1, 0xe9, 0x87, 0x21, 0xee, 0x34, 0xee, 0x5e, 0x8f,
	0x7d, 0xf7, 0x0e, 0xfc, 0x1a, 0xcb, 0x77, 0x32, 0x4d, 0xdb, 0xf3, 0xb4, 0x59, 0x26, 0x38, 0xa5,
	0x3a, 0xec, 0x43, 0x77, 0x9e, 0x7d, 0x7f, 0x9e, 0x09, 0x80, 0x0b, 0xc2, 0x27, 0x40, 0xef, 0x7a,
	0x4d, 0x10, 0x90, 0x23, 0xe4, 0x48, 0x89, 0x9c, 0xbb, 0xb2, 0x7d, 0x14, 0x26, 0x7b, 0x0c, 0xe0,
	0xa4, 0x99, 0xa6, 0xae, 0x4b, 0xb7, 0xc8, 0x4e, 0x54, 0x55, 0x33, 0x79, 0x4d, 0x0f, 0xc2, 0xa4,
	0x4d, 0xc1, 0x58, 0x8f, 0x61, 0x83, 0xd5, 0x6e, 0xed, 0x3e, 0x0e, 0x55, 0x4b, 0x40, 0xec, 0xc9,
	0xc2, 0xb5, 0x32, 0xbd, 0x17, 0x7a, 0x12, 0xbf, 0x91, 0x89, 0x4b, 0x68, 0xa8, 0x74, 0x75, 0x5c,
	0x54, 0x42, 0xd2, 0x4f, 0x02, 0x33, 0x01, 0xc9, 0xe7, 0xb0, 0xe5, 0x76, 0xe5, 0x70, 0xf8, 0x64,
	0xbc, 0x86, 0x32, 0x6f, 0x36, 0x85, 0xa2, 0x9d, 0xba, 0x92, 0x5c, 0x4f, 0xec, 0x3e, 0x0d, 0x76,
	0x29, 0x8a, 0x51, 0x3d, 0x72, 0xd6, 0x2e, 0xae, 0xc3, 0x10, 0x35, 0x01, 0x8f, 0xfe, 0x9f, 0x87,
	0x9d, 0xf3, 0xce, 0xf2, 0xfd, 0x85, 0xc9, 0xa2, 0x74, 0xb7, 0xff, 0x3b, 0xac, 0x9e, 0x0a, 0x53,
	0x33, 0x9b, 0xbf, 0x24, 0xb3, 0xdf, 0xf8, 0xf8, 0x40, 0xed, 0x7d, 0x76, 0x83, 0x55, 0x7c, 0x20,
	0xfe, 0x82, 0xf5, 0x9f, 0xb9, 0x6d, 0x1f, 0x4c, 0xf2, 0xe0, 0x5d, 0xfe, 0x1f, 0xda, 0x17, 0x7f,
	0xef, 0xeb, 0x77, 0xb4, 0x8e, 0xb1, 0x46, 0xd0, 0x73, 0xb1, 0xba, 0xef, 0x3e, 0xf9, 0xe6, 0xc6,
	0x97, 0x7d, 0x3a, 0xe6, 0xb7, 0xef, 0xc1, 0x08, 0x71, 0x7f, 0xea, 0xfd, 0xb1, 0x99, 0x58, 0xbd,
	0x58, 0xf6, 0xbf, 0x69, 0x8f, 0x5e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x39, 0x8e, 0xc3, 0x28, 0xd5,
	0x09, 0x00, 0x00,
}
