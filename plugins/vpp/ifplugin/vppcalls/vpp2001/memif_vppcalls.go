//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp2001

import (
	interfaces "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp2001/memif"
)

// AddMemifInterface implements interface handler.
func (h *InterfaceVppHandler) AddMemifInterface(ifName string, memIface *interfaces.MemifLink, socketID uint32) (swIdx uint32, err error) {
	req := &memif.MemifCreate{
		ID:         memIface.Id,
		Mode:       setMemifMode(memIface.Mode),
		Secret:     memIface.Secret,
		SocketID:   socketID,
		BufferSize: uint16(memIface.BufferSize),
		RingSize:   memIface.RingSize,
		RxQueues:   uint8(memIface.RxQueues),
		TxQueues:   uint8(memIface.TxQueues),
		NoZeroCopy: memIface.NoZeroCopy,
	}
	if memIface.Master {
		req.Role = 0
	} else {
		req.Role = 1
	}
	// TODO: temporary fix, waiting for https://gerrit.fd.io/r/#/c/7266/
	if req.RxQueues == 0 {
		req.RxQueues = 1
	}
	if req.TxQueues == 0 {
		req.TxQueues = 1
	}
	reply := &memif.MemifCreateReply{}

	if err = h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, err
	}

	return uint32(reply.SwIfIndex), h.SetInterfaceTag(ifName, uint32(reply.SwIfIndex))
}

// DeleteMemifInterface implements interface handler.
func (h *InterfaceVppHandler) DeleteMemifInterface(ifName string, idx uint32) error {
	req := &memif.MemifDelete{
		SwIfIndex: memif.InterfaceIndex(idx),
	}
	reply := &memif.MemifDeleteReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return h.RemoveInterfaceTag(ifName, idx)
}

// RegisterMemifSocketFilename implements interface handler.
func (h *InterfaceVppHandler) RegisterMemifSocketFilename(filename []byte, id uint32) error {
	req := &memif.MemifSocketFilenameAddDel{
		SocketFilename: string(filename),
		SocketID:       id,
		IsAdd:          true, // sockets can be added only
	}
	reply := &memif.MemifSocketFilenameAddDelReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return nil
}

func setMemifMode(mode interfaces.MemifLink_MemifMode) memif.MemifMode {
	switch mode {
	case interfaces.MemifLink_IP:
		return memif.MEMIF_MODE_API_IP
	case interfaces.MemifLink_PUNT_INJECT:
		return memif.MEMIF_MODE_API_PUNT_INJECT
	default:
		return memif.MEMIF_MODE_API_ETHERNET
	}
}
