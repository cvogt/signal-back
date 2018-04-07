// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signal/Backups.proto

/*
Package signal is a generated protocol buffer package.

It is generated from these files:
	signal/Backups.proto

It has these top-level messages:
	SqlStatement
	SharedPreference
	Attachment
	Avatar
	DatabaseVersion
	Header
	BackupFrame
*/
package signal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SqlStatement struct {
	Statement        *string                      `protobuf:"bytes,1,opt,name=statement" json:"statement,omitempty"`
	Parameters       []*SqlStatement_SqlParameter `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *SqlStatement) Reset()                    { *m = SqlStatement{} }
func (m *SqlStatement) String() string            { return proto.CompactTextString(m) }
func (*SqlStatement) ProtoMessage()               {}
func (*SqlStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SqlStatement) GetStatement() string {
	if m != nil && m.Statement != nil {
		return *m.Statement
	}
	return ""
}

func (m *SqlStatement) GetParameters() []*SqlStatement_SqlParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type SqlStatement_SqlParameter struct {
	StringParamter   *string  `protobuf:"bytes,1,opt,name=stringParamter" json:"stringParamter,omitempty"`
	IntegerParameter *uint64  `protobuf:"varint,2,opt,name=integerParameter" json:"integerParameter,omitempty"`
	DoubleParameter  *float64 `protobuf:"fixed64,3,opt,name=doubleParameter" json:"doubleParameter,omitempty"`
	BlobParameter    []byte   `protobuf:"bytes,4,opt,name=blobParameter" json:"blobParameter,omitempty"`
	Nullparameter    *bool    `protobuf:"varint,5,opt,name=nullparameter" json:"nullparameter,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SqlStatement_SqlParameter) Reset()                    { *m = SqlStatement_SqlParameter{} }
func (m *SqlStatement_SqlParameter) String() string            { return proto.CompactTextString(m) }
func (*SqlStatement_SqlParameter) ProtoMessage()               {}
func (*SqlStatement_SqlParameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *SqlStatement_SqlParameter) GetStringParamter() string {
	if m != nil && m.StringParamter != nil {
		return *m.StringParamter
	}
	return ""
}

func (m *SqlStatement_SqlParameter) GetIntegerParameter() uint64 {
	if m != nil && m.IntegerParameter != nil {
		return *m.IntegerParameter
	}
	return 0
}

func (m *SqlStatement_SqlParameter) GetDoubleParameter() float64 {
	if m != nil && m.DoubleParameter != nil {
		return *m.DoubleParameter
	}
	return 0
}

func (m *SqlStatement_SqlParameter) GetBlobParameter() []byte {
	if m != nil {
		return m.BlobParameter
	}
	return nil
}

func (m *SqlStatement_SqlParameter) GetNullparameter() bool {
	if m != nil && m.Nullparameter != nil {
		return *m.Nullparameter
	}
	return false
}

type SharedPreference struct {
	File             *string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Key              *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SharedPreference) Reset()                    { *m = SharedPreference{} }
func (m *SharedPreference) String() string            { return proto.CompactTextString(m) }
func (*SharedPreference) ProtoMessage()               {}
func (*SharedPreference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SharedPreference) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

func (m *SharedPreference) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *SharedPreference) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Attachment struct {
	RowId            *uint64 `protobuf:"varint,1,opt,name=rowId" json:"rowId,omitempty"`
	AttachmentId     *uint64 `protobuf:"varint,2,opt,name=attachmentId" json:"attachmentId,omitempty"`
	Length           *uint32 `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Attachment) Reset()                    { *m = Attachment{} }
func (m *Attachment) String() string            { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()               {}
func (*Attachment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Attachment) GetRowId() uint64 {
	if m != nil && m.RowId != nil {
		return *m.RowId
	}
	return 0
}

func (m *Attachment) GetAttachmentId() uint64 {
	if m != nil && m.AttachmentId != nil {
		return *m.AttachmentId
	}
	return 0
}

func (m *Attachment) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type Avatar struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Length           *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Avatar) Reset()                    { *m = Avatar{} }
func (m *Avatar) String() string            { return proto.CompactTextString(m) }
func (*Avatar) ProtoMessage()               {}
func (*Avatar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Avatar) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Avatar) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type DatabaseVersion struct {
	Version          *uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DatabaseVersion) Reset()                    { *m = DatabaseVersion{} }
func (m *DatabaseVersion) String() string            { return proto.CompactTextString(m) }
func (*DatabaseVersion) ProtoMessage()               {}
func (*DatabaseVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DatabaseVersion) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type Header struct {
	Iv               []byte `protobuf:"bytes,1,opt,name=iv" json:"iv,omitempty"`
	Salt             []byte `protobuf:"bytes,2,opt,name=salt" json:"salt,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Header) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

func (m *Header) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

type BackupFrame struct {
	Header           *Header           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Statement        *SqlStatement     `protobuf:"bytes,2,opt,name=statement" json:"statement,omitempty"`
	Preference       *SharedPreference `protobuf:"bytes,3,opt,name=preference" json:"preference,omitempty"`
	Attachment       *Attachment       `protobuf:"bytes,4,opt,name=attachment" json:"attachment,omitempty"`
	Version          *DatabaseVersion  `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	End              *bool             `protobuf:"varint,6,opt,name=end" json:"end,omitempty"`
	Avatar           *Avatar           `protobuf:"bytes,7,opt,name=avatar" json:"avatar,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *BackupFrame) Reset()                    { *m = BackupFrame{} }
func (m *BackupFrame) String() string            { return proto.CompactTextString(m) }
func (*BackupFrame) ProtoMessage()               {}
func (*BackupFrame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BackupFrame) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BackupFrame) GetStatement() *SqlStatement {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *BackupFrame) GetPreference() *SharedPreference {
	if m != nil {
		return m.Preference
	}
	return nil
}

func (m *BackupFrame) GetAttachment() *Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *BackupFrame) GetVersion() *DatabaseVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *BackupFrame) GetEnd() bool {
	if m != nil && m.End != nil {
		return *m.End
	}
	return false
}

func (m *BackupFrame) GetAvatar() *Avatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func init() {
	proto.RegisterType((*SqlStatement)(nil), "signal.SqlStatement")
	proto.RegisterType((*SqlStatement_SqlParameter)(nil), "signal.SqlStatement.SqlParameter")
	proto.RegisterType((*SharedPreference)(nil), "signal.SharedPreference")
	proto.RegisterType((*Attachment)(nil), "signal.Attachment")
	proto.RegisterType((*Avatar)(nil), "signal.Avatar")
	proto.RegisterType((*DatabaseVersion)(nil), "signal.DatabaseVersion")
	proto.RegisterType((*Header)(nil), "signal.Header")
	proto.RegisterType((*BackupFrame)(nil), "signal.BackupFrame")
}

func init() { proto.RegisterFile("signal/Backups.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0x69, 0x9b, 0xb5, 0xb7, 0x69, 0xb7, 0x0c, 0x45, 0x83, 0xf8, 0x90, 0x0d, 0xb2, 0x04,
	0x95, 0xc8, 0x06, 0x1f, 0x7c, 0x6d, 0x11, 0x71, 0x5f, 0xa4, 0x4c, 0xc1, 0x47, 0x61, 0xda, 0xdc,
	0x4d, 0xc2, 0xa6, 0x49, 0x9d, 0x99, 0x54, 0xfc, 0x7f, 0x3e, 0xf9, 0x0b, 0xfc, 0x39, 0x32, 0x93,
	0xef, 0xba, 0x6f, 0xf7, 0x9e, 0x39, 0xe7, 0xce, 0xe4, 0xdc, 0x43, 0x60, 0x25, 0xd2, 0x38, 0x67,
	0xd9, 0xfb, 0x0d, 0x3b, 0x3c, 0x96, 0x27, 0x11, 0x9c, 0x78, 0x21, 0x0b, 0x62, 0x55, 0xa8, 0xf7,
	0xc7, 0x04, 0x7b, 0xf7, 0x23, 0xdb, 0x49, 0x26, 0xf1, 0x88, 0xb9, 0x24, 0xaf, 0x60, 0x2a, 0x9a,
	0xc6, 0x31, 0x5c, 0xc3, 0x9f, 0xd2, 0x0e, 0x20, 0x6b, 0x80, 0x13, 0xe3, 0xec, 0x88, 0x12, 0xb9,
	0x70, 0x4c, 0x77, 0xe4, 0xcf, 0xc2, 0x9b, 0xa0, 0x9a, 0x15, 0xf4, 0xe7, 0xa8, 0x66, 0xdb, 0x30,
	0x69, 0x4f, 0xf4, 0xf2, 0xaf, 0xa1, 0x6f, 0x6c, 0x0f, 0xc9, 0x2d, 0x2c, 0x84, 0xe4, 0x69, 0x1e,
	0x6b, 0x48, 0x22, 0xaf, 0xaf, 0xbd, 0x40, 0xc9, 0x1b, 0x58, 0xa6, 0xb9, 0xc4, 0x18, 0x79, 0xab,
	0x75, 0x4c, 0xd7, 0xf0, 0xc7, 0xf4, 0x3f, 0x9c, 0xf8, 0x70, 0x1d, 0x15, 0xe5, 0x3e, 0xc3, 0x8e,
	0x3a, 0x72, 0x0d, 0xdf, 0xa0, 0x97, 0x30, 0x79, 0x0d, 0xf3, 0x7d, 0x56, 0xec, 0x3b, 0xde, 0xd8,
	0x35, 0x7c, 0x9b, 0x0e, 0x41, 0xc5, 0xca, 0xcb, 0x2c, 0x6b, 0x3f, 0xc3, 0x99, 0xb8, 0x86, 0xff,
	0x8c, 0x0e, 0x41, 0xef, 0x2b, 0x2c, 0x77, 0x09, 0xe3, 0x18, 0x6d, 0x39, 0x3e, 0x20, 0xc7, 0xfc,
	0x80, 0x84, 0xc0, 0xf8, 0x21, 0xcd, 0xb0, 0xfe, 0x26, 0x5d, 0x93, 0x25, 0x8c, 0x1e, 0xf1, 0x97,
	0x7e, 0xfc, 0x94, 0xaa, 0x92, 0xac, 0x60, 0x72, 0x66, 0x59, 0x89, 0xfa, 0x95, 0x53, 0x5a, 0x35,
	0xde, 0x77, 0x80, 0xb5, 0x94, 0xec, 0x90, 0x68, 0xef, 0x57, 0x30, 0xe1, 0xc5, 0xcf, 0xfb, 0x48,
	0x8f, 0x1a, 0xd3, 0xaa, 0x21, 0x1e, 0xd8, 0xac, 0xe5, 0xdc, 0x47, 0xb5, 0x23, 0x03, 0x8c, 0x3c,
	0x07, 0x2b, 0xc3, 0x3c, 0x96, 0x89, 0x1e, 0x3f, 0xa7, 0x75, 0xe7, 0x7d, 0x00, 0x6b, 0x7d, 0x66,
	0x92, 0x71, 0xf5, 0xca, 0x9c, 0x1d, 0xdb, 0x57, 0xaa, 0xba, 0xa7, 0x32, 0x07, 0xaa, 0xb7, 0x70,
	0xfd, 0x89, 0x49, 0xb6, 0x67, 0x02, 0xbf, 0x21, 0x17, 0x69, 0x91, 0x13, 0x07, 0xae, 0xce, 0x55,
	0xa9, 0x27, 0xcc, 0x69, 0xd3, 0x7a, 0xef, 0xc0, 0xfa, 0x82, 0x2c, 0x42, 0x4e, 0x16, 0x60, 0xa6,
	0x67, 0x7d, 0x6c, 0x53, 0x33, 0x3d, 0xab, 0x2b, 0x05, 0xcb, 0xa4, 0x1e, 0x6e, 0x53, 0x5d, 0x7b,
	0xbf, 0x4d, 0x98, 0x55, 0x39, 0xfd, 0xac, 0x3c, 0x25, 0xb7, 0x60, 0x25, 0x5a, 0xad, 0x75, 0xb3,
	0x70, 0xd1, 0x44, 0xad, 0x9a, 0x49, 0xeb, 0x53, 0x12, 0xf6, 0x43, 0x6b, 0x6a, 0xea, 0xea, 0xa9,
	0x54, 0xf6, 0xa3, 0xfc, 0x11, 0xe0, 0xd4, 0xae, 0x49, 0x1b, 0x33, 0x0b, 0x9d, 0x56, 0x74, 0xb1,
	0x46, 0xda, 0xe3, 0x92, 0x10, 0xa0, 0xb3, 0x57, 0xe7, 0x65, 0x16, 0x92, 0x46, 0xd9, 0x2d, 0x8c,
	0xf6, 0x58, 0xe4, 0xae, 0x73, 0x68, 0xa2, 0x05, 0x2f, 0x1a, 0xc1, 0x85, 0x97, 0xad, 0x75, 0x2a,
	0x25, 0x98, 0x47, 0x8e, 0xa5, 0x93, 0xa6, 0x4a, 0x65, 0x07, 0xd3, 0xfb, 0x72, 0xae, 0x86, 0x76,
	0x54, 0x5b, 0xa4, 0xf5, 0xe9, 0xe6, 0x0e, 0x6e, 0x0a, 0x1e, 0x07, 0x32, 0x29, 0xca, 0x38, 0x91,
	0x07, 0x9e, 0x1e, 0x31, 0x10, 0x78, 0x28, 0x39, 0x8a, 0xa3, 0x08, 0xf6, 0xda, 0xe0, 0x8d, 0x5d,
	0x19, 0xbd, 0x55, 0xbf, 0x03, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x19, 0x28, 0x16, 0x8b, 0x26,
	0x04, 0x00, 0x00,
}