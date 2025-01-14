// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/abf.api.json

/*
Package abf is a generated VPP binary API for 'abf' module.

It consists of:
	  4 types
	 10 messages
	  5 services
*/
package abf

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
	ModuleName = "abf"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x3abf2f80
)

// AbfItfAttach represents VPP binary API type 'abf_itf_attach'.
type AbfItfAttach struct {
	PolicyID  uint32
	SwIfIndex uint32
	Priority  uint32
	IsIPv6    uint8
}

func (*AbfItfAttach) GetTypeName() string {
	return "abf_itf_attach"
}
func (*AbfItfAttach) GetCrcString() string {
	return "aa3ea7fe"
}

// AbfPolicy represents VPP binary API type 'abf_policy'.
type AbfPolicy struct {
	PolicyID uint32
	ACLIndex uint32
	NPaths   uint8 `struc:"sizeof=Paths"`
	Paths    []FibPath
}

func (*AbfPolicy) GetTypeName() string {
	return "abf_policy"
}
func (*AbfPolicy) GetCrcString() string {
	return "252c563e"
}

// FibMplsLabel represents VPP binary API type 'fib_mpls_label'.
type FibMplsLabel struct {
	IsUniform uint8
	Label     uint32
	TTL       uint8
	Exp       uint8
}

func (*FibMplsLabel) GetTypeName() string {
	return "fib_mpls_label"
}
func (*FibMplsLabel) GetCrcString() string {
	return "c93bf35c"
}

// FibPath represents VPP binary API type 'fib_path'.
type FibPath struct {
	SwIfIndex         uint32
	TableID           uint32
	Weight            uint8
	Preference        uint8
	IsLocal           uint8
	IsDrop            uint8
	IsUDPEncap        uint8
	IsUnreach         uint8
	IsProhibit        uint8
	IsResolveHost     uint8
	IsResolveAttached uint8
	IsDvr             uint8
	IsSourceLookup    uint8
	IsInterfaceRx     uint8
	Afi               uint8
	NextHop           []byte `struc:"[16]byte"`
	NextHopID         uint32
	RpfID             uint32
	ViaLabel          uint32
	NLabels           uint8 `struc:"sizeof=LabelStack"` // MANUALLY FIXED
	LabelStack        []FibMplsLabel
}

func (*FibPath) GetTypeName() string {
	return "fib_path"
}
func (*FibPath) GetCrcString() string {
	return "ba7a81f0"
}

// AbfItfAttachAddDel represents VPP binary API message 'abf_itf_attach_add_del'.
type AbfItfAttachAddDel struct {
	IsAdd  uint8
	Attach AbfItfAttach
}

func (*AbfItfAttachAddDel) GetMessageName() string {
	return "abf_itf_attach_add_del"
}
func (*AbfItfAttachAddDel) GetCrcString() string {
	return "b0b50ab1"
}
func (*AbfItfAttachAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AbfItfAttachAddDelReply represents VPP binary API message 'abf_itf_attach_add_del_reply'.
type AbfItfAttachAddDelReply struct {
	Retval int32
}

func (*AbfItfAttachAddDelReply) GetMessageName() string {
	return "abf_itf_attach_add_del_reply"
}
func (*AbfItfAttachAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AbfItfAttachAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AbfItfAttachDetails represents VPP binary API message 'abf_itf_attach_details'.
type AbfItfAttachDetails struct {
	Attach AbfItfAttach
}

func (*AbfItfAttachDetails) GetMessageName() string {
	return "abf_itf_attach_details"
}
func (*AbfItfAttachDetails) GetCrcString() string {
	return "999006a7"
}
func (*AbfItfAttachDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AbfItfAttachDump represents VPP binary API message 'abf_itf_attach_dump'.
type AbfItfAttachDump struct{}

func (*AbfItfAttachDump) GetMessageName() string {
	return "abf_itf_attach_dump"
}
func (*AbfItfAttachDump) GetCrcString() string {
	return "51077d14"
}
func (*AbfItfAttachDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AbfPluginGetVersion represents VPP binary API message 'abf_plugin_get_version'.
type AbfPluginGetVersion struct{}

func (*AbfPluginGetVersion) GetMessageName() string {
	return "abf_plugin_get_version"
}
func (*AbfPluginGetVersion) GetCrcString() string {
	return "51077d14"
}
func (*AbfPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AbfPluginGetVersionReply represents VPP binary API message 'abf_plugin_get_version_reply'.
type AbfPluginGetVersionReply struct {
	Major uint32
	Minor uint32
}

func (*AbfPluginGetVersionReply) GetMessageName() string {
	return "abf_plugin_get_version_reply"
}
func (*AbfPluginGetVersionReply) GetCrcString() string {
	return "9b32cf86"
}
func (*AbfPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AbfPolicyAddDel represents VPP binary API message 'abf_policy_add_del'.
type AbfPolicyAddDel struct {
	IsAdd  uint8
	Policy AbfPolicy
}

func (*AbfPolicyAddDel) GetMessageName() string {
	return "abf_policy_add_del"
}
func (*AbfPolicyAddDel) GetCrcString() string {
	return "40432d41"
}
func (*AbfPolicyAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AbfPolicyAddDelReply represents VPP binary API message 'abf_policy_add_del_reply'.
type AbfPolicyAddDelReply struct {
	Retval int32
}

func (*AbfPolicyAddDelReply) GetMessageName() string {
	return "abf_policy_add_del_reply"
}
func (*AbfPolicyAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AbfPolicyAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AbfPolicyDetails represents VPP binary API message 'abf_policy_details'.
type AbfPolicyDetails struct {
	Policy AbfPolicy
}

func (*AbfPolicyDetails) GetMessageName() string {
	return "abf_policy_details"
}
func (*AbfPolicyDetails) GetCrcString() string {
	return "ca17332a"
}
func (*AbfPolicyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AbfPolicyDump represents VPP binary API message 'abf_policy_dump'.
type AbfPolicyDump struct{}

func (*AbfPolicyDump) GetMessageName() string {
	return "abf_policy_dump"
}
func (*AbfPolicyDump) GetCrcString() string {
	return "51077d14"
}
func (*AbfPolicyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*AbfItfAttachAddDel)(nil), "abf.AbfItfAttachAddDel")
	api.RegisterMessage((*AbfItfAttachAddDelReply)(nil), "abf.AbfItfAttachAddDelReply")
	api.RegisterMessage((*AbfItfAttachDetails)(nil), "abf.AbfItfAttachDetails")
	api.RegisterMessage((*AbfItfAttachDump)(nil), "abf.AbfItfAttachDump")
	api.RegisterMessage((*AbfPluginGetVersion)(nil), "abf.AbfPluginGetVersion")
	api.RegisterMessage((*AbfPluginGetVersionReply)(nil), "abf.AbfPluginGetVersionReply")
	api.RegisterMessage((*AbfPolicyAddDel)(nil), "abf.AbfPolicyAddDel")
	api.RegisterMessage((*AbfPolicyAddDelReply)(nil), "abf.AbfPolicyAddDelReply")
	api.RegisterMessage((*AbfPolicyDetails)(nil), "abf.AbfPolicyDetails")
	api.RegisterMessage((*AbfPolicyDump)(nil), "abf.AbfPolicyDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AbfItfAttachAddDel)(nil),
		(*AbfItfAttachAddDelReply)(nil),
		(*AbfItfAttachDetails)(nil),
		(*AbfItfAttachDump)(nil),
		(*AbfPluginGetVersion)(nil),
		(*AbfPluginGetVersionReply)(nil),
		(*AbfPolicyAddDel)(nil),
		(*AbfPolicyAddDelReply)(nil),
		(*AbfPolicyDetails)(nil),
		(*AbfPolicyDump)(nil),
	}
}

// RPCService represents RPC service API for abf module.
type RPCService interface {
	DumpAbfItfAttach(ctx context.Context, in *AbfItfAttachDump) (RPCService_DumpAbfItfAttachClient, error)
	DumpAbfPolicy(ctx context.Context, in *AbfPolicyDump) (RPCService_DumpAbfPolicyClient, error)
	AbfItfAttachAddDel(ctx context.Context, in *AbfItfAttachAddDel) (*AbfItfAttachAddDelReply, error)
	AbfPluginGetVersion(ctx context.Context, in *AbfPluginGetVersion) (*AbfPluginGetVersionReply, error)
	AbfPolicyAddDel(ctx context.Context, in *AbfPolicyAddDel) (*AbfPolicyAddDelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpAbfItfAttach(ctx context.Context, in *AbfItfAttachDump) (RPCService_DumpAbfItfAttachClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpAbfItfAttachClient{stream}
	return x, nil
}

type RPCService_DumpAbfItfAttachClient interface {
	Recv() (*AbfItfAttachDetails, error)
}

type serviceClient_DumpAbfItfAttachClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpAbfItfAttachClient) Recv() (*AbfItfAttachDetails, error) {
	m := new(AbfItfAttachDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpAbfPolicy(ctx context.Context, in *AbfPolicyDump) (RPCService_DumpAbfPolicyClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpAbfPolicyClient{stream}
	return x, nil
}

type RPCService_DumpAbfPolicyClient interface {
	Recv() (*AbfPolicyDetails, error)
}

type serviceClient_DumpAbfPolicyClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpAbfPolicyClient) Recv() (*AbfPolicyDetails, error) {
	m := new(AbfPolicyDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) AbfItfAttachAddDel(ctx context.Context, in *AbfItfAttachAddDel) (*AbfItfAttachAddDelReply, error) {
	out := new(AbfItfAttachAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AbfPluginGetVersion(ctx context.Context, in *AbfPluginGetVersion) (*AbfPluginGetVersionReply, error) {
	out := new(AbfPluginGetVersionReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AbfPolicyAddDel(ctx context.Context, in *AbfPolicyAddDel) (*AbfPolicyAddDelReply, error) {
	out := new(AbfPolicyAddDelReply)
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
