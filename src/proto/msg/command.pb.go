// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: command.proto

package msg

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

type CMD int32

const (
	CMD_KEEP_ALIVE                 CMD = 0
	CMD_RESP                       CMD = 1
	CMD_GATE                       CMD = 2
	CMD_SERVER_STATE               CMD = 3
	CMD_SERVER_LIST_CHANGED        CMD = 4
	CMD_SERVER_RESET               CMD = 5
	CMD_CONNECT_GATE_REQ           CMD = 6
	CMD_SESSION_CERT_NTF           CMD = 7
	CMD_CCHAT_CHAT                 CMD = 8
	CMD_DISPATCH                   CMD = 9
	CMD_LD_CREATE_PLAYER_REQ       CMD = 10
	CMD_GD_LOAD_PLAYER_REQ         CMD = 11
	CMD_GD_RELEASE_PLAYER_NTF      CMD = 12
	CMD_GD_CREATE_ROLE_REQ         CMD = 13
	CMD_GD_LOAD_ROLE_REQ           CMD = 14
	CMD_RELOAD_ETC_REQ             CMD = 15
	CMD_UPDATE_ETC_NTF             CMD = 16
	CMD_MOVE_NTF                   CMD = 17
	CMD_NEAR_ENTITIES_NTF          CMD = 18
	CMD_ENTITY_DATA_REQ            CMD = 19
	CMD_ROLE_ENTER                 CMD = 20
	CMD_UPDATE_PLAYER_ADDRESS_NTF  CMD = 21
	CMD_REMOVE_PLAYER_ADDRESS_NTF  CMD = 22
	CMD_UPDATE_MAP_ADDRESS_NTF     CMD = 23
	CMD_REMOVE_MAP_ADDRESS_NTF     CMD = 24
	CMD_GET_MAP_ADDRESS_REQ        CMD = 25
	CMD_LOAD_STATIC_MAP_REQ        CMD = 26
	CMD_LOAD_DYNAMIC_MAP_REQ       CMD = 27
	CMD_UPDATE_STATIC_MAP_UUID_NTF CMD = 28
	CMD_GET_STATIC_MAP_UUID_REQ    CMD = 29
	CMD_SAVE_STATIC_MAP_REQ        CMD = 30
	CMD_SAVE_DYNAMIC_MAP_REQ       CMD = 31
	CMD_SAVE_ROLE                  CMD = 32
	CMD_GET_MAP_ID                 CMD = 33
	CMD_MOVE_INF                   CMD = 34
	CMD_FORWARD_PLAYER_MESSAGE     CMD = 35
	CMD_NODE_START_NTF             CMD = 1000
	CMD_ETC_SYNC_NTF               CMD = 1001
	CMD_NODES_INFO_NTF             CMD = 1002
)

var CMD_name = map[int32]string{
	0:    "KEEP_ALIVE",
	1:    "RESP",
	2:    "GATE",
	3:    "SERVER_STATE",
	4:    "SERVER_LIST_CHANGED",
	5:    "SERVER_RESET",
	6:    "CONNECT_GATE_REQ",
	7:    "SESSION_CERT_NTF",
	8:    "CCHAT_CHAT",
	9:    "DISPATCH",
	10:   "LD_CREATE_PLAYER_REQ",
	11:   "GD_LOAD_PLAYER_REQ",
	12:   "GD_RELEASE_PLAYER_NTF",
	13:   "GD_CREATE_ROLE_REQ",
	14:   "GD_LOAD_ROLE_REQ",
	15:   "RELOAD_ETC_REQ",
	16:   "UPDATE_ETC_NTF",
	17:   "MOVE_NTF",
	18:   "NEAR_ENTITIES_NTF",
	19:   "ENTITY_DATA_REQ",
	20:   "ROLE_ENTER",
	21:   "UPDATE_PLAYER_ADDRESS_NTF",
	22:   "REMOVE_PLAYER_ADDRESS_NTF",
	23:   "UPDATE_MAP_ADDRESS_NTF",
	24:   "REMOVE_MAP_ADDRESS_NTF",
	25:   "GET_MAP_ADDRESS_REQ",
	26:   "LOAD_STATIC_MAP_REQ",
	27:   "LOAD_DYNAMIC_MAP_REQ",
	28:   "UPDATE_STATIC_MAP_UUID_NTF",
	29:   "GET_STATIC_MAP_UUID_REQ",
	30:   "SAVE_STATIC_MAP_REQ",
	31:   "SAVE_DYNAMIC_MAP_REQ",
	32:   "SAVE_ROLE",
	33:   "GET_MAP_ID",
	34:   "MOVE_INF",
	35:   "FORWARD_PLAYER_MESSAGE",
	1000: "NODE_START_NTF",
	1001: "ETC_SYNC_NTF",
	1002: "NODES_INFO_NTF",
}

var CMD_value = map[string]int32{
	"KEEP_ALIVE":                 0,
	"RESP":                       1,
	"GATE":                       2,
	"SERVER_STATE":               3,
	"SERVER_LIST_CHANGED":        4,
	"SERVER_RESET":               5,
	"CONNECT_GATE_REQ":           6,
	"SESSION_CERT_NTF":           7,
	"CCHAT_CHAT":                 8,
	"DISPATCH":                   9,
	"LD_CREATE_PLAYER_REQ":       10,
	"GD_LOAD_PLAYER_REQ":         11,
	"GD_RELEASE_PLAYER_NTF":      12,
	"GD_CREATE_ROLE_REQ":         13,
	"GD_LOAD_ROLE_REQ":           14,
	"RELOAD_ETC_REQ":             15,
	"UPDATE_ETC_NTF":             16,
	"MOVE_NTF":                   17,
	"NEAR_ENTITIES_NTF":          18,
	"ENTITY_DATA_REQ":            19,
	"ROLE_ENTER":                 20,
	"UPDATE_PLAYER_ADDRESS_NTF":  21,
	"REMOVE_PLAYER_ADDRESS_NTF":  22,
	"UPDATE_MAP_ADDRESS_NTF":     23,
	"REMOVE_MAP_ADDRESS_NTF":     24,
	"GET_MAP_ADDRESS_REQ":        25,
	"LOAD_STATIC_MAP_REQ":        26,
	"LOAD_DYNAMIC_MAP_REQ":       27,
	"UPDATE_STATIC_MAP_UUID_NTF": 28,
	"GET_STATIC_MAP_UUID_REQ":    29,
	"SAVE_STATIC_MAP_REQ":        30,
	"SAVE_DYNAMIC_MAP_REQ":       31,
	"SAVE_ROLE":                  32,
	"GET_MAP_ID":                 33,
	"MOVE_INF":                   34,
	"FORWARD_PLAYER_MESSAGE":     35,
	"NODE_START_NTF":             1000,
	"ETC_SYNC_NTF":               1001,
	"NODES_INFO_NTF":             1002,
}

func (x CMD) String() string {
	return proto.EnumName(CMD_name, int32(x))
}

func (CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_213c0bb044472049, []int{0}
}

func init() {
	proto.RegisterEnum("CMD", CMD_name, CMD_value)
}

func init() { proto.RegisterFile("command.proto", fileDescriptor_213c0bb044472049) }

var fileDescriptor_213c0bb044472049 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x72, 0x52, 0x31,
	0x14, 0x2e, 0xb6, 0x16, 0x7a, 0x04, 0x1a, 0xc2, 0x5f, 0xa1, 0xf6, 0xfa, 0xb7, 0x71, 0x5c, 0xc8,
	0xc2, 0x27, 0x88, 0x37, 0x87, 0x4b, 0xc6, 0x4b, 0xee, 0x6d, 0x12, 0x70, 0x70, 0x93, 0xd1, 0xda,
	0x71, 0x85, 0x38, 0xd4, 0xf1, 0x59, 0x7c, 0x24, 0x97, 0x6e, 0xdc, 0x3b, 0xb8, 0x51, 0x9f, 0xc2,
	0x39, 0xe1, 0xc2, 0x50, 0xc6, 0xdd, 0xe5, 0xfb, 0xf2, 0xfd, 0x24, 0xe7, 0x00, 0xb5, 0xab, 0xc5,
	0x7c, 0xfe, 0xf6, 0xe3, 0xfb, 0xe7, 0x9f, 0x96, 0x8b, 0xcf, 0x8b, 0x67, 0x3f, 0x8e, 0xe1, 0x30,
	0x1e, 0x4b, 0x5e, 0x07, 0x78, 0x85, 0x98, 0x7b, 0x91, 0xaa, 0x29, 0xb2, 0x03, 0x5e, 0x81, 0x23,
	0x83, 0x36, 0x67, 0x25, 0xfa, 0x4a, 0x84, 0x43, 0x76, 0x87, 0x33, 0xa8, 0x5a, 0x34, 0x53, 0x34,
	0xde, 0x3a, 0x42, 0x0e, 0x79, 0x17, 0x9a, 0x05, 0x92, 0x2a, 0xeb, 0x7c, 0x3c, 0x12, 0x3a, 0x41,
	0xc9, 0x8e, 0x76, 0x8e, 0x1a, 0xb4, 0xe8, 0xd8, 0x5d, 0xde, 0x02, 0x16, 0x67, 0x5a, 0x63, 0xec,
	0x3c, 0xd9, 0x79, 0x83, 0x97, 0xec, 0x98, 0x50, 0x8b, 0xd6, 0xaa, 0x4c, 0xfb, 0x18, 0x8d, 0xf3,
	0xda, 0x0d, 0x59, 0x99, 0xca, 0xc4, 0xf1, 0x48, 0x04, 0x43, 0xc7, 0x2a, 0xbc, 0x0a, 0x15, 0xa9,
	0x6c, 0x2e, 0x5c, 0x3c, 0x62, 0x27, 0xfc, 0x0c, 0x5a, 0xa9, 0xf4, 0xb1, 0x41, 0xb2, 0xc9, 0x53,
	0x31, 0x0b, 0x29, 0x97, 0x0c, 0x78, 0x07, 0x78, 0x22, 0x7d, 0x9a, 0x09, 0xb9, 0x8b, 0xdf, 0xe3,
	0x3d, 0x68, 0x27, 0xd2, 0x1b, 0x4c, 0x51, 0xd8, 0xad, 0x84, 0xa2, 0xaa, 0x85, 0xa4, 0x30, 0x33,
	0x59, 0xba, 0x2e, 0x56, 0xa3, 0x62, 0x1b, 0xab, 0x2d, 0x5a, 0xe7, 0x1c, 0xea, 0x06, 0x03, 0x88,
	0x2e, 0x0e, 0xd8, 0x29, 0x61, 0x93, 0x5c, 0x92, 0x9c, 0x30, 0x72, 0x65, 0x54, 0x78, 0x9c, 0x4d,
	0x31, 0xfc, 0x6a, 0xf0, 0x36, 0x34, 0x34, 0x0a, 0xe3, 0x51, 0x3b, 0xe5, 0x14, 0xda, 0x00, 0x73,
	0xde, 0x84, 0xd3, 0x80, 0xcc, 0xbc, 0x14, 0x4e, 0x04, 0xb7, 0x26, 0x5d, 0x3d, 0xe4, 0xa1, 0x76,
	0x68, 0x58, 0x8b, 0x5f, 0x40, 0xaf, 0x70, 0x2f, 0x6a, 0x0b, 0x29, 0x0d, 0xda, 0xb5, 0x47, 0x9b,
	0x68, 0x83, 0x21, 0xea, 0x3f, 0x74, 0x87, 0xf7, 0xa1, 0x53, 0xa8, 0xc7, 0x22, 0xbf, 0xc5, 0x75,
	0x89, 0x2b, 0xa4, 0xfb, 0xdc, 0x19, 0xcd, 0x35, 0x41, 0x77, 0x8b, 0xa0, 0x7a, 0x3d, 0x22, 0xc2,
	0xf5, 0x69, 0x01, 0x54, 0x1c, 0x0e, 0x10, 0xd1, 0x0f, 0x43, 0x21, 0x42, 0xce, 0xb4, 0x18, 0xef,
	0x30, 0xe7, 0x3c, 0x82, 0x7e, 0xd1, 0x61, 0x47, 0x34, 0x99, 0x28, 0x19, 0xb2, 0xee, 0xf3, 0x73,
	0xe8, 0x52, 0xd6, 0x3e, 0x49, 0xe2, 0x8b, 0xb0, 0x60, 0x62, 0x8a, 0xfb, 0x79, 0x11, 0xe5, 0x05,
	0x62, 0x3f, 0xef, 0x01, 0xaf, 0xc1, 0x49, 0x60, 0xe8, 0x19, 0xd9, 0x43, 0x7a, 0xd0, 0xcd, 0x55,
	0x94, 0x64, 0x8f, 0xb6, 0xa3, 0x51, 0x7a, 0xc8, 0x1e, 0xd3, 0x23, 0x0c, 0x33, 0xf3, 0x5a, 0x98,
	0xed, 0xc6, 0x8c, 0xd1, 0x5a, 0x91, 0x20, 0x7b, 0xc2, 0x9b, 0x50, 0xd7, 0x99, 0x0c, 0xd9, 0xc5,
	0x66, 0xfe, 0x2e, 0xf3, 0x06, 0x54, 0x69, 0xcc, 0x76, 0xa6, 0xd7, 0xb3, 0xfe, 0x53, 0xde, 0x9c,
	0xb3, 0x64, 0x99, 0x05, 0xf0, 0x6f, 0xf9, 0xe5, 0xd3, 0x6f, 0xab, 0xa8, 0xf4, 0x7d, 0x15, 0x95,
	0x7e, 0xae, 0xa2, 0xd2, 0xd7, 0x5f, 0xd1, 0xc1, 0x9b, 0x8e, 0xd2, 0xf6, 0x7a, 0xf9, 0xe5, 0x7a,
	0x39, 0xb8, 0x59, 0x5e, 0x0d, 0xc2, 0xbf, 0x6f, 0x30, 0xbf, 0xf9, 0xf0, 0xee, 0x38, 0x7c, 0xbe,
	0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x25, 0xe0, 0x10, 0xee, 0x99, 0x03, 0x00, 0x00,
}
