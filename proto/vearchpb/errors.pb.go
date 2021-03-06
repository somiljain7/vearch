// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errors.proto

package vearchpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type ErrorEnum int32

const (
	ErrorEnum_SUCCESS                              ErrorEnum = 0
	ErrorEnum_INTERNAL_ERROR                       ErrorEnum = 1
	ErrorEnum_NAME_OR_PASSWORD                     ErrorEnum = 2
	ErrorEnum_SYSBUSY                              ErrorEnum = 3
	ErrorEnum_PARAM_ERROR                          ErrorEnum = 4
	ErrorEnum_INVALID_CFG                          ErrorEnum = 5
	ErrorEnum_TIMEOUT                              ErrorEnum = 6
	ErrorEnum_SERVICE_UNAVAILABLE                  ErrorEnum = 7
	ErrorEnum_ZONE_NOT_EXISTS                      ErrorEnum = 8
	ErrorEnum_LOCAL_ZONE_OPS_FAILED                ErrorEnum = 9
	ErrorEnum_DUP_ZONE                             ErrorEnum = 10
	ErrorEnum_DUP_DB                               ErrorEnum = 11
	ErrorEnum_INVALID_ENGINE                       ErrorEnum = 12
	ErrorEnum_DB_NOTEXISTS                         ErrorEnum = 13
	ErrorEnum_DB_Not_Empty                         ErrorEnum = 14
	ErrorEnum_DUP_SPACE                            ErrorEnum = 15
	ErrorEnum_SPACE_NOTEXISTS                      ErrorEnum = 16
	ErrorEnum_PARTITION_HAS_TASK_NOW               ErrorEnum = 17
	ErrorEnum_REPLICA_NOT_EXISTS                   ErrorEnum = 18
	ErrorEnum_DUP_REPLICA                          ErrorEnum = 19
	ErrorEnum_PARTITION_REPLICA_LEADER_NOT_DELETE  ErrorEnum = 20
	ErrorEnum_PS_NOTEXISTS                         ErrorEnum = 21
	ErrorEnum_PS_Already_Exists                    ErrorEnum = 22
	ErrorEnum_LOCAL_SPACE_OPS_FAILED               ErrorEnum = 23
	ErrorEnum_Local_PS_Ops_Failed                  ErrorEnum = 24
	ErrorEnum_GENID_FAILED                         ErrorEnum = 25
	ErrorEnum_LOCALDB_OPTFAILED                    ErrorEnum = 26
	ErrorEnum_SPACE_SCHEMA_INVALID                 ErrorEnum = 27
	ErrorEnum_RPC_GET_CLIENT_FAILED                ErrorEnum = 28
	ErrorEnum_RPC_INVALID_RESP                     ErrorEnum = 29
	ErrorEnum_RPC_INVOKE_FAILED                    ErrorEnum = 30
	ErrorEnum_RPC_PARAM_ERROR                      ErrorEnum = 31
	ErrorEnum_METHOD_NOT_IMPLEMENT                 ErrorEnum = 32
	ErrorEnum_USER_NOT_EXISTS                      ErrorEnum = 33
	ErrorEnum_DUP_USER                             ErrorEnum = 34
	ErrorEnum_USER_OPS_FAILED                      ErrorEnum = 35
	ErrorEnum_AUTHENTICATION_FAILED                ErrorEnum = 36
	ErrorEnum_REGION_NOT_EXISTS                    ErrorEnum = 37
	ErrorEnum_MASTER_PS_CAN_NOT_SELECT             ErrorEnum = 38
	ErrorEnum_MASTER_PS_NOT_ENOUGH_SELECT          ErrorEnum = 39
	ErrorEnum_PARTITION_DUPLICATE                  ErrorEnum = 40
	ErrorEnum_PARTITION_NOT_EXIST                  ErrorEnum = 41
	ErrorEnum_PARTITION_NOT_LEADER                 ErrorEnum = 42
	ErrorEnum_PARTITION_NO_LEADER                  ErrorEnum = 43
	ErrorEnum_PARTITION_REQ_PARAM                  ErrorEnum = 44
	ErrorEnum_PARTITON_ENGINENAME_INVALID          ErrorEnum = 45
	ErrorEnum_UNKNOWN_PARTITION_RAFT_CMD_TYPE      ErrorEnum = 46
	ErrorEnum_MASTER_SERVER_IS_NOT_RUNNING         ErrorEnum = 47
	ErrorEnum_PARTITION_IS_INVALID                 ErrorEnum = 48
	ErrorEnum_PARTITION_IS_CLOSED                  ErrorEnum = 49
	ErrorEnum_DOCUMENT_NOT_EXIST                   ErrorEnum = 50
	ErrorEnum_DOCUMENT_EXIST                       ErrorEnum = 51
	ErrorEnum_DOCUMENT_MUST_HAS_SOURCE             ErrorEnum = 52
	ErrorEnum_PULL_OUT_VERSION_NOT_MATCH           ErrorEnum = 53
	ErrorEnum_FUNC_CAN_NOT_INVOKE_IN_FROZEN_ENGINE ErrorEnum = 54
	ErrorEnum_ROUTER_NO_PS_CLIENT                  ErrorEnum = 55
	ErrorEnum_ROUTER_CALL_PS_RPC_ERR               ErrorEnum = 56
	ErrorEnum_GAMMA_SEARCH_QUERY_NUM_LESS_0        ErrorEnum = 57
	ErrorEnum_GAMMA_SEARCH_NO_CREATE_INDEX         ErrorEnum = 58
	ErrorEnum_GAMMA_SEARCH_INDEX_QUERY_ERR         ErrorEnum = 59
	ErrorEnum_GAMMA_SEARCH_OTHER_ERR               ErrorEnum = 60
	ErrorEnum_Primary_IS_INVALID                   ErrorEnum = 61
	ErrorEnum_PARSING_RESULT_ERROR                 ErrorEnum = 62
	ErrorEnum_FORCE_MERGE_BUILD_INDEX_ERR          ErrorEnum = 63
	ErrorEnum_DELETE_BY_QUERY_SERACH_ERR           ErrorEnum = 64
	ErrorEnum_DELETE_BY_QUERY_SEARCH_ID_IS_0       ErrorEnum = 65
	ErrorEnum_Create_RpcClient_Failed              ErrorEnum = 70
	ErrorEnum_Call_RpcClient_Failed                ErrorEnum = 71
	ErrorEnum_RECOVER                              ErrorEnum = 100
)

var ErrorEnum_name = map[int32]string{
	0:   "SUCCESS",
	1:   "INTERNAL_ERROR",
	2:   "NAME_OR_PASSWORD",
	3:   "SYSBUSY",
	4:   "PARAM_ERROR",
	5:   "INVALID_CFG",
	6:   "TIMEOUT",
	7:   "SERVICE_UNAVAILABLE",
	8:   "ZONE_NOT_EXISTS",
	9:   "LOCAL_ZONE_OPS_FAILED",
	10:  "DUP_ZONE",
	11:  "DUP_DB",
	12:  "INVALID_ENGINE",
	13:  "DB_NOTEXISTS",
	14:  "DB_Not_Empty",
	15:  "DUP_SPACE",
	16:  "SPACE_NOTEXISTS",
	17:  "PARTITION_HAS_TASK_NOW",
	18:  "REPLICA_NOT_EXISTS",
	19:  "DUP_REPLICA",
	20:  "PARTITION_REPLICA_LEADER_NOT_DELETE",
	21:  "PS_NOTEXISTS",
	22:  "PS_Already_Exists",
	23:  "LOCAL_SPACE_OPS_FAILED",
	24:  "Local_PS_Ops_Failed",
	25:  "GENID_FAILED",
	26:  "LOCALDB_OPTFAILED",
	27:  "SPACE_SCHEMA_INVALID",
	28:  "RPC_GET_CLIENT_FAILED",
	29:  "RPC_INVALID_RESP",
	30:  "RPC_INVOKE_FAILED",
	31:  "RPC_PARAM_ERROR",
	32:  "METHOD_NOT_IMPLEMENT",
	33:  "USER_NOT_EXISTS",
	34:  "DUP_USER",
	35:  "USER_OPS_FAILED",
	36:  "AUTHENTICATION_FAILED",
	37:  "REGION_NOT_EXISTS",
	38:  "MASTER_PS_CAN_NOT_SELECT",
	39:  "MASTER_PS_NOT_ENOUGH_SELECT",
	40:  "PARTITION_DUPLICATE",
	41:  "PARTITION_NOT_EXIST",
	42:  "PARTITION_NOT_LEADER",
	43:  "PARTITION_NO_LEADER",
	44:  "PARTITION_REQ_PARAM",
	45:  "PARTITON_ENGINENAME_INVALID",
	46:  "UNKNOWN_PARTITION_RAFT_CMD_TYPE",
	47:  "MASTER_SERVER_IS_NOT_RUNNING",
	48:  "PARTITION_IS_INVALID",
	49:  "PARTITION_IS_CLOSED",
	50:  "DOCUMENT_NOT_EXIST",
	51:  "DOCUMENT_EXIST",
	52:  "DOCUMENT_MUST_HAS_SOURCE",
	53:  "PULL_OUT_VERSION_NOT_MATCH",
	54:  "FUNC_CAN_NOT_INVOKE_IN_FROZEN_ENGINE",
	55:  "ROUTER_NO_PS_CLIENT",
	56:  "ROUTER_CALL_PS_RPC_ERR",
	57:  "GAMMA_SEARCH_QUERY_NUM_LESS_0",
	58:  "GAMMA_SEARCH_NO_CREATE_INDEX",
	59:  "GAMMA_SEARCH_INDEX_QUERY_ERR",
	60:  "GAMMA_SEARCH_OTHER_ERR",
	61:  "Primary_IS_INVALID",
	62:  "PARSING_RESULT_ERROR",
	63:  "FORCE_MERGE_BUILD_INDEX_ERR",
	64:  "DELETE_BY_QUERY_SERACH_ERR",
	65:  "DELETE_BY_QUERY_SEARCH_ID_IS_0",
	70:  "Create_RpcClient_Failed",
	71:  "Call_RpcClient_Failed",
	100: "RECOVER",
}

var ErrorEnum_value = map[string]int32{
	"SUCCESS":                              0,
	"INTERNAL_ERROR":                       1,
	"NAME_OR_PASSWORD":                     2,
	"SYSBUSY":                              3,
	"PARAM_ERROR":                          4,
	"INVALID_CFG":                          5,
	"TIMEOUT":                              6,
	"SERVICE_UNAVAILABLE":                  7,
	"ZONE_NOT_EXISTS":                      8,
	"LOCAL_ZONE_OPS_FAILED":                9,
	"DUP_ZONE":                             10,
	"DUP_DB":                               11,
	"INVALID_ENGINE":                       12,
	"DB_NOTEXISTS":                         13,
	"DB_Not_Empty":                         14,
	"DUP_SPACE":                            15,
	"SPACE_NOTEXISTS":                      16,
	"PARTITION_HAS_TASK_NOW":               17,
	"REPLICA_NOT_EXISTS":                   18,
	"DUP_REPLICA":                          19,
	"PARTITION_REPLICA_LEADER_NOT_DELETE":  20,
	"PS_NOTEXISTS":                         21,
	"PS_Already_Exists":                    22,
	"LOCAL_SPACE_OPS_FAILED":               23,
	"Local_PS_Ops_Failed":                  24,
	"GENID_FAILED":                         25,
	"LOCALDB_OPTFAILED":                    26,
	"SPACE_SCHEMA_INVALID":                 27,
	"RPC_GET_CLIENT_FAILED":                28,
	"RPC_INVALID_RESP":                     29,
	"RPC_INVOKE_FAILED":                    30,
	"RPC_PARAM_ERROR":                      31,
	"METHOD_NOT_IMPLEMENT":                 32,
	"USER_NOT_EXISTS":                      33,
	"DUP_USER":                             34,
	"USER_OPS_FAILED":                      35,
	"AUTHENTICATION_FAILED":                36,
	"REGION_NOT_EXISTS":                    37,
	"MASTER_PS_CAN_NOT_SELECT":             38,
	"MASTER_PS_NOT_ENOUGH_SELECT":          39,
	"PARTITION_DUPLICATE":                  40,
	"PARTITION_NOT_EXIST":                  41,
	"PARTITION_NOT_LEADER":                 42,
	"PARTITION_NO_LEADER":                  43,
	"PARTITION_REQ_PARAM":                  44,
	"PARTITON_ENGINENAME_INVALID":          45,
	"UNKNOWN_PARTITION_RAFT_CMD_TYPE":      46,
	"MASTER_SERVER_IS_NOT_RUNNING":         47,
	"PARTITION_IS_INVALID":                 48,
	"PARTITION_IS_CLOSED":                  49,
	"DOCUMENT_NOT_EXIST":                   50,
	"DOCUMENT_EXIST":                       51,
	"DOCUMENT_MUST_HAS_SOURCE":             52,
	"PULL_OUT_VERSION_NOT_MATCH":           53,
	"FUNC_CAN_NOT_INVOKE_IN_FROZEN_ENGINE": 54,
	"ROUTER_NO_PS_CLIENT":                  55,
	"ROUTER_CALL_PS_RPC_ERR":               56,
	"GAMMA_SEARCH_QUERY_NUM_LESS_0":        57,
	"GAMMA_SEARCH_NO_CREATE_INDEX":         58,
	"GAMMA_SEARCH_INDEX_QUERY_ERR":         59,
	"GAMMA_SEARCH_OTHER_ERR":               60,
	"Primary_IS_INVALID":                   61,
	"PARSING_RESULT_ERROR":                 62,
	"FORCE_MERGE_BUILD_INDEX_ERR":          63,
	"DELETE_BY_QUERY_SERACH_ERR":           64,
	"DELETE_BY_QUERY_SEARCH_ID_IS_0":       65,
	"Create_RpcClient_Failed":              70,
	"Call_RpcClient_Failed":                71,
	"RECOVER":                              100,
}

func (x ErrorEnum) String() string {
	return proto.EnumName(ErrorEnum_name, int32(x))
}

func (ErrorEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}

type SearchResultCode int32

const (
	SearchResultCode_SEARCH_SUCCESS    SearchResultCode = 0
	SearchResultCode_INDEX_NOT_BUILDED SearchResultCode = 1
	SearchResultCode_SEARCH_ERROR      SearchResultCode = 2
)

var SearchResultCode_name = map[int32]string{
	0: "SEARCH_SUCCESS",
	1: "INDEX_NOT_BUILDED",
	2: "SEARCH_ERROR",
}

var SearchResultCode_value = map[string]int32{
	"SEARCH_SUCCESS":    0,
	"INDEX_NOT_BUILDED": 1,
	"SEARCH_ERROR":      2,
}

func (x SearchResultCode) String() string {
	return proto.EnumName(SearchResultCode_name, int32(x))
}

func (SearchResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{1}
}

type Error struct {
	Code                 ErrorEnum `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorEnum" json:"code,omitempty"`
	Msg                  string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Error) Reset()      { *m = Error{} }
func (*Error) ProtoMessage() {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ErrorEnum", ErrorEnum_name, ErrorEnum_value)
	proto.RegisterEnum("SearchResultCode", SearchResultCode_name, SearchResultCode_value)
	proto.RegisterType((*Error)(nil), "Error")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }

var fileDescriptor_24fe73c7f0ddb19c = []byte{
	// 1122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x55, 0xcb, 0x76, 0x13, 0x47,
	0x13, 0xd6, 0x70, 0x77, 0x63, 0x4c, 0xd3, 0x60, 0x30, 0x06, 0x06, 0x73, 0xf9, 0x7f, 0x1c, 0x12,
	0x0c, 0x81, 0xdc, 0xc8, 0x95, 0x56, 0x4f, 0x49, 0x9a, 0xc3, 0x4c, 0xf7, 0xd0, 0xdd, 0x63, 0x30,
	0x9b, 0x3e, 0xb2, 0xad, 0x18, 0x9f, 0x23, 0x23, 0x1f, 0x49, 0xce, 0x89, 0x77, 0x79, 0x8c, 0x3c,
	0x42, 0x36, 0xd9, 0x67, 0x99, 0x25, 0xcb, 0x2c, 0xb3, 0xc4, 0xca, 0x0b, 0x64, 0x99, 0x65, 0x4e,
	0xf5, 0xcc, 0xc8, 0x52, 0xbc, 0x9b, 0xa9, 0xaf, 0x2e, 0x5f, 0x7d, 0x55, 0x35, 0x43, 0x66, 0x3b,
	0xfd, 0x7e, 0xaf, 0x3f, 0x58, 0xd9, 0xed, 0xf7, 0x86, 0xbd, 0xc5, 0x07, 0x5b, 0xdb, 0xc3, 0x37,
	0x7b, 0xeb, 0x2b, 0x1b, 0xbd, 0x9d, 0x87, 0x5b, 0xbd, 0xad, 0xde, 0x43, 0x6f, 0x5e, 0xdf, 0xfb,
	0xde, 0xbf, 0xf9, 0x17, 0xff, 0x54, 0xb8, 0xdf, 0x7e, 0x4a, 0x4e, 0x02, 0x86, 0xb3, 0x90, 0x9c,
	0xd8, 0xe8, 0x6d, 0x76, 0x16, 0x82, 0xa5, 0x60, 0x79, 0xee, 0x31, 0x59, 0xf1, 0x56, 0x78, 0xbb,
	0xb7, 0xa3, 0xbd, 0x9d, 0x51, 0x72, 0x7c, 0x67, 0xb0, 0xb5, 0x70, 0x6c, 0x29, 0x58, 0x9e, 0xd1,
	0xf8, 0x78, 0xff, 0xd7, 0x73, 0x64, 0x66, 0xec, 0xc5, 0xce, 0x92, 0xd3, 0x26, 0x17, 0x02, 0x8c,
	0xa1, 0x35, 0xc6, 0xc8, 0x5c, 0x2c, 0x2d, 0x68, 0xc9, 0x13, 0x07, 0x5a, 0x2b, 0x4d, 0x03, 0x76,
	0x89, 0x50, 0xc9, 0x53, 0x70, 0x4a, 0xbb, 0x8c, 0x1b, 0xf3, 0x52, 0xe9, 0x88, 0x1e, 0xf3, 0x61,
	0x6b, 0xa6, 0x9e, 0x9b, 0x35, 0x7a, 0x9c, 0x9d, 0x27, 0x67, 0x33, 0xae, 0x79, 0x5a, 0xc6, 0x9c,
	0x40, 0x43, 0x2c, 0x57, 0x79, 0x12, 0x47, 0x4e, 0x34, 0x9a, 0xf4, 0x24, 0xba, 0xdb, 0x38, 0x05,
	0x95, 0x5b, 0x7a, 0x8a, 0x5d, 0x21, 0x17, 0x0d, 0xe8, 0xd5, 0x58, 0x80, 0xcb, 0x25, 0x5f, 0xe5,
	0x71, 0xc2, 0xeb, 0x09, 0xd0, 0xd3, 0xec, 0x22, 0x39, 0xff, 0x5a, 0x49, 0x70, 0x52, 0x59, 0x07,
	0xaf, 0x62, 0x63, 0x0d, 0x3d, 0xc3, 0xae, 0x92, 0xf9, 0x44, 0x09, 0x9e, 0x38, 0x0f, 0xa9, 0xcc,
	0xb8, 0x06, 0x8f, 0x13, 0x88, 0xe8, 0x0c, 0x9b, 0x25, 0x67, 0xa2, 0x3c, 0xf3, 0x00, 0x25, 0x8c,
	0x90, 0x53, 0xf8, 0x16, 0xd5, 0xe9, 0xd9, 0xa2, 0x91, 0x82, 0x00, 0xc8, 0x66, 0x2c, 0x81, 0xce,
	0x32, 0x4a, 0x66, 0xa3, 0x3a, 0xe6, 0x2e, 0x53, 0x9f, 0xab, 0x2c, 0xbd, 0xa1, 0x83, 0x9d, 0xdd,
	0xe1, 0x3e, 0x9d, 0x63, 0xe7, 0xc8, 0x0c, 0xe6, 0x30, 0x19, 0x17, 0x40, 0xcf, 0x23, 0x21, 0xff,
	0x38, 0x11, 0x45, 0xd9, 0x22, 0xb9, 0x9c, 0x71, 0x6d, 0x63, 0x1b, 0x2b, 0xe9, 0x5a, 0xdc, 0x38,
	0xcb, 0xcd, 0x73, 0x27, 0xd5, 0x4b, 0x7a, 0x81, 0x5d, 0x26, 0x4c, 0x43, 0x96, 0xc4, 0x82, 0x4f,
	0x36, 0xc1, 0x50, 0x10, 0xcc, 0x5b, 0x62, 0xf4, 0x22, 0xbb, 0x47, 0xee, 0x1c, 0x26, 0xa9, 0x42,
	0x12, 0xe0, 0x11, 0x68, 0x1f, 0x19, 0x41, 0x02, 0x16, 0xe8, 0x25, 0xe4, 0x98, 0x99, 0x89, 0xfa,
	0xf3, 0x6c, 0x9e, 0x5c, 0xc8, 0x8c, 0xe3, 0xdd, 0x7e, 0xa7, 0xbd, 0xb9, 0xef, 0xe0, 0xc7, 0xed,
	0xc1, 0x70, 0x40, 0x2f, 0x23, 0xad, 0x42, 0xa7, 0x82, 0xf1, 0x84, 0x50, 0x57, 0x50, 0xf1, 0xa4,
	0xb7, 0xd1, 0xee, 0xba, 0xcc, 0x38, 0xb5, 0x3b, 0x70, 0x8d, 0xf6, 0x76, 0xb7, 0xb3, 0x49, 0x17,
	0x30, 0x7b, 0x13, 0x64, 0x1c, 0x55, 0xae, 0x57, 0x31, 0xbb, 0x4f, 0x13, 0xd5, 0x9d, 0xca, 0x6c,
	0x69, 0x5e, 0x64, 0x0b, 0xe4, 0x52, 0x91, 0xd7, 0x88, 0x16, 0xa4, 0xdc, 0x95, 0xea, 0xd2, 0x6b,
	0x38, 0x1f, 0x9d, 0x09, 0xd7, 0x04, 0xeb, 0x44, 0x12, 0x83, 0xb4, 0x55, 0xae, 0xeb, 0xb8, 0x3a,
	0x08, 0x55, 0x93, 0xd0, 0x60, 0x32, 0x7a, 0x03, 0x2b, 0x94, 0x56, 0xf5, 0x1c, 0x2a, 0xe7, 0x10,
	0xb5, 0x46, 0xf3, 0xe4, 0x22, 0xdd, 0xc4, 0xb2, 0x29, 0xd8, 0x96, 0x8a, 0xbc, 0x28, 0x71, 0x9a,
	0x25, 0x90, 0x82, 0xb4, 0x74, 0x09, 0xdd, 0x73, 0x53, 0x8a, 0x55, 0x4a, 0x73, 0xab, 0x5a, 0x08,
	0x04, 0xe8, 0xed, 0xb1, 0xcb, 0x84, 0x14, 0x77, 0x90, 0x2e, 0xcf, 0x6d, 0x0b, 0xa4, 0x8d, 0x05,
	0xf7, 0xea, 0x97, 0xd0, 0x5d, 0x4f, 0x0c, 0x9a, 0x68, 0x9a, 0x48, 0xfa, 0x3f, 0x76, 0x9d, 0x2c,
	0xa4, 0xdc, 0x58, 0xd0, 0xa8, 0x9e, 0xe0, 0x05, 0x6a, 0x20, 0x01, 0x61, 0xe9, 0xff, 0xd9, 0x4d,
	0x72, 0xed, 0x10, 0xf5, 0x71, 0x52, 0xe5, 0xcd, 0x56, 0xe5, 0x70, 0x0f, 0xb5, 0x3f, 0x9c, 0x74,
	0x94, 0xfb, 0x49, 0x5b, 0xa0, 0xcb, 0xd3, 0xc0, 0xb8, 0x22, 0xfd, 0x00, 0x9b, 0x9e, 0x06, 0x8a,
	0xbd, 0xa0, 0xf7, 0xff, 0x1b, 0x52, 0x01, 0x1f, 0x4e, 0x03, 0x1a, 0x5e, 0x14, 0x32, 0xd2, 0x8f,
	0x90, 0x5e, 0x01, 0x28, 0x59, 0x5e, 0x82, 0x3f, 0xe6, 0x6a, 0x7c, 0x0f, 0xd8, 0x1d, 0x72, 0x33,
	0x97, 0xcf, 0xa5, 0x7a, 0x29, 0xdd, 0x44, 0x06, 0xde, 0xb0, 0x4e, 0xa4, 0x91, 0xb3, 0x6b, 0x19,
	0xd0, 0x15, 0xb6, 0x44, 0xae, 0x97, 0x4d, 0xe2, 0xe1, 0x82, 0x76, 0x71, 0xd1, 0xab, 0xce, 0xa5,
	0x8c, 0x65, 0x93, 0x3e, 0x9c, 0xe6, 0x1c, 0x9b, 0x71, 0x81, 0x47, 0xd3, 0xd4, 0x62, 0xe3, 0x44,
	0xa2, 0x0c, 0x44, 0xf4, 0x63, 0xbc, 0x95, 0x48, 0x89, 0x1c, 0xe7, 0x39, 0xd1, 0xfe, 0x63, 0xbc,
	0xdd, 0xb1, 0xbd, 0xb0, 0x3d, 0xc1, 0x19, 0x8c, 0x6d, 0x69, 0x6e, 0xac, 0xbf, 0x3b, 0xa3, 0x72,
	0x2d, 0x80, 0x7e, 0xc2, 0x42, 0xb2, 0x98, 0xe5, 0x49, 0xe2, 0x54, 0x6e, 0xdd, 0x2a, 0x68, 0x53,
	0xe9, 0x96, 0x72, 0x2b, 0x5a, 0xf4, 0x53, 0xb6, 0x4c, 0xee, 0x36, 0x72, 0x29, 0xc6, 0xc3, 0x2b,
	0x57, 0x2f, 0x96, 0xae, 0xa1, 0xd5, 0x6b, 0xa8, 0x94, 0xa1, 0x9f, 0x21, 0x59, 0xad, 0x72, 0xeb,
	0xf7, 0xca, 0x8f, 0xdb, 0x6f, 0x34, 0xfd, 0x1c, 0xaf, 0xab, 0x04, 0x04, 0x4f, 0x12, 0x84, 0x70,
	0x59, 0x41, 0x6b, 0xfa, 0x05, 0xbb, 0x45, 0x6e, 0x34, 0x79, 0x9a, 0x72, 0x67, 0x80, 0x6b, 0xd1,
	0x72, 0x2f, 0x72, 0xd0, 0x6b, 0x4e, 0xe6, 0xa9, 0x4b, 0xc0, 0x18, 0xf7, 0x88, 0x3e, 0x45, 0x01,
	0xa7, 0x5c, 0xa4, 0x72, 0x42, 0x03, 0xb7, 0x48, 0x22, 0x82, 0x57, 0xf4, 0xcb, 0x23, 0x1e, 0xde,
	0x5e, 0xa6, 0xc2, 0x32, 0x5f, 0x21, 0x85, 0x29, 0x0f, 0x65, 0x5b, 0xa0, 0x3d, 0xf6, 0x35, 0x6a,
	0x99, 0xf5, 0xb7, 0x77, 0xda, 0xfd, 0xfd, 0x49, 0xf1, 0xbf, 0x29, 0xc7, 0x62, 0x62, 0xd9, 0xc4,
	0xeb, 0xcb, 0x13, 0x5b, 0x5e, 0xd6, 0xb7, 0xb8, 0x18, 0x0d, 0xa5, 0x05, 0xb8, 0x14, 0x74, 0x13,
	0x5c, 0x3d, 0x8f, 0x93, 0xa8, 0x2c, 0x8a, 0x29, 0xbf, 0x43, 0x51, 0x8b, 0x8f, 0x90, 0xab, 0xaf,
	0x95, 0x3c, 0x0c, 0x68, 0x2e, 0x5a, 0x1e, 0x7f, 0xc6, 0x6e, 0x93, 0xf0, 0x28, 0x5e, 0x50, 0x8f,
	0x90, 0xc4, 0x23, 0xca, 0xd9, 0x35, 0x72, 0x45, 0xf4, 0x3b, 0xed, 0x61, 0xc7, 0xe9, 0xdd, 0x0d,
	0xd1, 0xdd, 0xee, 0xbc, 0x1d, 0x56, 0xdf, 0x9e, 0x06, 0x5e, 0xa2, 0x68, 0x77, 0xbb, 0x47, 0xa1,
	0x26, 0xfe, 0x2e, 0x34, 0x08, 0xb5, 0x0a, 0x9a, 0x6e, 0xde, 0x57, 0x84, 0x9a, 0x4e, 0xbb, 0xbf,
	0xf1, 0x46, 0x77, 0x06, 0x7b, 0xdd, 0xa1, 0xc0, 0xbf, 0x1a, 0x23, 0x73, 0x65, 0xb1, 0xc3, 0x9f,
	0xd7, 0x3c, 0xb9, 0x50, 0xf0, 0xc7, 0x11, 0xfb, 0x7e, 0x20, 0xa2, 0x01, 0x7e, 0xe2, 0x4a, 0xd7,
	0xa2, 0xf5, 0x63, 0xf5, 0x67, 0xef, 0x0e, 0xc2, 0xda, 0x9f, 0x07, 0x61, 0xed, 0xfd, 0x41, 0x58,
	0xfb, 0xfb, 0x20, 0xac, 0xfd, 0x73, 0x10, 0x06, 0x3f, 0x8d, 0xc2, 0xe0, 0x97, 0x51, 0x18, 0xfc,
	0x36, 0x0a, 0x6b, 0xbf, 0x8f, 0xc2, 0xda, 0xbb, 0x51, 0x18, 0xfc, 0x31, 0x0a, 0x83, 0xf7, 0xa3,
	0x30, 0xf8, 0xf9, 0xaf, 0xb0, 0xd6, 0x0a, 0x5e, 0x9f, 0xf9, 0xc1, 0xd3, 0xd8, 0x5d, 0x5f, 0x3f,
	0xe5, 0x7f, 0xc2, 0x4f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x33, 0xe3, 0x26, 0xc3, 0x07,
	0x00, 0x00,
}

func (this *Error) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Error)
	if !ok {
		that2, ok := that.(Error)
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
	if this.Code != that1.Code {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedError(r randyErrors, easy bool) *Error {
	this := &Error{}
	this.Code = ErrorEnum([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 70, 71, 100}[r.Intn(69)])
	this.Msg = string(randStringErrors(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedErrors(r, 3)
	}
	return this
}

type randyErrors interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneErrors(r randyErrors) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringErrors(r randyErrors) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneErrors(r)
	}
	return string(tmps)
}
func randUnrecognizedErrors(r randyErrors, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldErrors(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldErrors(dAtA []byte, r randyErrors, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateErrors(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateErrors(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Error{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Msg:` + fmt.Sprintf("%v", this.Msg) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringErrors(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= ErrorEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
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
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
