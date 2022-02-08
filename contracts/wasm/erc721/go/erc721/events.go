// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:gocritic
package erc721

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type Erc721Events struct {
}

func (e Erc721Events) Approval(approved wasmtypes.ScAgentID, owner wasmtypes.ScAgentID, tokenID wasmtypes.ScHash) {
	evt := wasmlib.NewEventEncoder("erc721.approval")
	evt.Encode(wasmtypes.AgentIDToString(approved))
	evt.Encode(wasmtypes.AgentIDToString(owner))
	evt.Encode(wasmtypes.HashToString(tokenID))
	evt.Emit()
}

func (e Erc721Events) ApprovalForAll(approval bool, operator wasmtypes.ScAgentID, owner wasmtypes.ScAgentID) {
	evt := wasmlib.NewEventEncoder("erc721.approvalForAll")
	evt.Encode(wasmtypes.BoolToString(approval))
	evt.Encode(wasmtypes.AgentIDToString(operator))
	evt.Encode(wasmtypes.AgentIDToString(owner))
	evt.Emit()
}

func (e Erc721Events) Init(name string, symbol string) {
	evt := wasmlib.NewEventEncoder("erc721.init")
	evt.Encode(wasmtypes.StringToString(name))
	evt.Encode(wasmtypes.StringToString(symbol))
	evt.Emit()
}

func (e Erc721Events) Mint(balance uint64, owner wasmtypes.ScAgentID, tokenID wasmtypes.ScHash) {
	evt := wasmlib.NewEventEncoder("erc721.mint")
	evt.Encode(wasmtypes.Uint64ToString(balance))
	evt.Encode(wasmtypes.AgentIDToString(owner))
	evt.Encode(wasmtypes.HashToString(tokenID))
	evt.Emit()
}

func (e Erc721Events) Transfer(from wasmtypes.ScAgentID, to wasmtypes.ScAgentID, tokenID wasmtypes.ScHash) {
	evt := wasmlib.NewEventEncoder("erc721.transfer")
	evt.Encode(wasmtypes.AgentIDToString(from))
	evt.Encode(wasmtypes.AgentIDToString(to))
	evt.Encode(wasmtypes.HashToString(tokenID))
	evt.Emit()
}
