// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/memif.api.json

/*
Package memif is a generated VPP binary API for 'memif' module.

It consists of:
	 10 messages
	  5 services
*/
package memif

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "memif"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x31b42e17
)

// MemifCreate represents VPP binary API message 'memif_create'.
type MemifCreate struct {
	Role       uint8
	Mode       uint8
	RxQueues   uint8
	TxQueues   uint8
	ID         uint32
	SocketID   uint32
	Secret     []byte `struc:"[24]byte"`
	RingSize   uint32
	BufferSize uint16
	HwAddr     []byte `struc:"[6]byte"`
}

func (*MemifCreate) GetMessageName() string {
	return "memif_create"
}
func (*MemifCreate) GetCrcString() string {
	return "6597cdb2"
}
func (*MemifCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifCreateReply represents VPP binary API message 'memif_create_reply'.
type MemifCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*MemifCreateReply) GetMessageName() string {
	return "memif_create_reply"
}
func (*MemifCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*MemifCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifDelete represents VPP binary API message 'memif_delete'.
type MemifDelete struct {
	SwIfIndex uint32
}

func (*MemifDelete) GetMessageName() string {
	return "memif_delete"
}
func (*MemifDelete) GetCrcString() string {
	return "529cb13f"
}
func (*MemifDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifDeleteReply represents VPP binary API message 'memif_delete_reply'.
type MemifDeleteReply struct {
	Retval int32
}

func (*MemifDeleteReply) GetMessageName() string {
	return "memif_delete_reply"
}
func (*MemifDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemifDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifDetails represents VPP binary API message 'memif_details'.
type MemifDetails struct {
	SwIfIndex   uint32
	IfName      []byte `struc:"[64]byte"`
	HwAddr      []byte `struc:"[6]byte"`
	ID          uint32
	Role        uint8
	Mode        uint8
	SocketID    uint32
	RingSize    uint32
	BufferSize  uint16
	AdminUpDown uint8
	LinkUpDown  uint8
}

func (*MemifDetails) GetMessageName() string {
	return "memif_details"
}
func (*MemifDetails) GetCrcString() string {
	return "4f5a3397"
}
func (*MemifDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifDump represents VPP binary API message 'memif_dump'.
type MemifDump struct{}

func (*MemifDump) GetMessageName() string {
	return "memif_dump"
}
func (*MemifDump) GetCrcString() string {
	return "51077d14"
}
func (*MemifDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifSocketFilenameAddDel represents VPP binary API message 'memif_socket_filename_add_del'.
type MemifSocketFilenameAddDel struct {
	IsAdd          uint8
	SocketID       uint32
	SocketFilename []byte `struc:"[128]byte"`
}

func (*MemifSocketFilenameAddDel) GetMessageName() string {
	return "memif_socket_filename_add_del"
}
func (*MemifSocketFilenameAddDel) GetCrcString() string {
	return "30e3929d"
}
func (*MemifSocketFilenameAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemifSocketFilenameAddDelReply represents VPP binary API message 'memif_socket_filename_add_del_reply'.
type MemifSocketFilenameAddDelReply struct {
	Retval int32
}

func (*MemifSocketFilenameAddDelReply) GetMessageName() string {
	return "memif_socket_filename_add_del_reply"
}
func (*MemifSocketFilenameAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemifSocketFilenameAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifSocketFilenameDetails represents VPP binary API message 'memif_socket_filename_details'.
type MemifSocketFilenameDetails struct {
	SocketID       uint32
	SocketFilename []byte `struc:"[128]byte"`
}

func (*MemifSocketFilenameDetails) GetMessageName() string {
	return "memif_socket_filename_details"
}
func (*MemifSocketFilenameDetails) GetCrcString() string {
	return "e347e32f"
}
func (*MemifSocketFilenameDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemifSocketFilenameDump represents VPP binary API message 'memif_socket_filename_dump'.
type MemifSocketFilenameDump struct{}

func (*MemifSocketFilenameDump) GetMessageName() string {
	return "memif_socket_filename_dump"
}
func (*MemifSocketFilenameDump) GetCrcString() string {
	return "51077d14"
}
func (*MemifSocketFilenameDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*MemifCreate)(nil), "memif.MemifCreate")
	api.RegisterMessage((*MemifCreateReply)(nil), "memif.MemifCreateReply")
	api.RegisterMessage((*MemifDelete)(nil), "memif.MemifDelete")
	api.RegisterMessage((*MemifDeleteReply)(nil), "memif.MemifDeleteReply")
	api.RegisterMessage((*MemifDetails)(nil), "memif.MemifDetails")
	api.RegisterMessage((*MemifDump)(nil), "memif.MemifDump")
	api.RegisterMessage((*MemifSocketFilenameAddDel)(nil), "memif.MemifSocketFilenameAddDel")
	api.RegisterMessage((*MemifSocketFilenameAddDelReply)(nil), "memif.MemifSocketFilenameAddDelReply")
	api.RegisterMessage((*MemifSocketFilenameDetails)(nil), "memif.MemifSocketFilenameDetails")
	api.RegisterMessage((*MemifSocketFilenameDump)(nil), "memif.MemifSocketFilenameDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MemifCreate)(nil),
		(*MemifCreateReply)(nil),
		(*MemifDelete)(nil),
		(*MemifDeleteReply)(nil),
		(*MemifDetails)(nil),
		(*MemifDump)(nil),
		(*MemifSocketFilenameAddDel)(nil),
		(*MemifSocketFilenameAddDelReply)(nil),
		(*MemifSocketFilenameDetails)(nil),
		(*MemifSocketFilenameDump)(nil),
	}
}

// RPCService represents RPC service API for memif module.
type RPCService interface {
	DumpMemif(ctx context.Context, in *MemifDump) (RPCService_DumpMemifClient, error)
	DumpMemifSocketFilename(ctx context.Context, in *MemifSocketFilenameDump) (RPCService_DumpMemifSocketFilenameClient, error)
	MemifCreate(ctx context.Context, in *MemifCreate) (*MemifCreateReply, error)
	MemifDelete(ctx context.Context, in *MemifDelete) (*MemifDeleteReply, error)
	MemifSocketFilenameAddDel(ctx context.Context, in *MemifSocketFilenameAddDel) (*MemifSocketFilenameAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpMemif(ctx context.Context, in *MemifDump) (RPCService_DumpMemifClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMemifClient{stream}
	return x, nil
}

type RPCService_DumpMemifClient interface {
	Recv() (*MemifDetails, error)
}

type serviceClient_DumpMemifClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMemifClient) Recv() (*MemifDetails, error) {
	m := new(MemifDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMemifSocketFilename(ctx context.Context, in *MemifSocketFilenameDump) (RPCService_DumpMemifSocketFilenameClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMemifSocketFilenameClient{stream}
	return x, nil
}

type RPCService_DumpMemifSocketFilenameClient interface {
	Recv() (*MemifSocketFilenameDetails, error)
}

type serviceClient_DumpMemifSocketFilenameClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMemifSocketFilenameClient) Recv() (*MemifSocketFilenameDetails, error) {
	m := new(MemifSocketFilenameDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) MemifCreate(ctx context.Context, in *MemifCreate) (*MemifCreateReply, error) {
	out := new(MemifCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MemifDelete(ctx context.Context, in *MemifDelete) (*MemifDeleteReply, error) {
	out := new(MemifDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MemifSocketFilenameAddDel(ctx context.Context, in *MemifSocketFilenameAddDel) (*MemifSocketFilenameAddDelReply, error) {
	out := new(MemifSocketFilenameAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
