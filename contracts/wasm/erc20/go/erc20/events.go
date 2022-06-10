// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:gocritic
package erc20

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type Erc20Events struct{}

func (e Erc20Events) Approval(
	amount uint64,
	owner wasmtypes.ScAgentID,
	spender wasmtypes.ScAgentID,
) {
	evt := wasmlib.NewEventEncoder("erc20.approval")
	evt.Encode(wasmtypes.Uint64ToString(amount))
	evt.Encode(wasmtypes.AgentIDToString(owner))
	evt.Encode(wasmtypes.AgentIDToString(spender))
	evt.Emit()
}

func (e Erc20Events) Transfer(
	amount uint64,
	from wasmtypes.ScAgentID,
	to wasmtypes.ScAgentID,
) {
	evt := wasmlib.NewEventEncoder("erc20.transfer")
	evt.Encode(wasmtypes.Uint64ToString(amount))
	evt.Encode(wasmtypes.AgentIDToString(from))
	evt.Encode(wasmtypes.AgentIDToString(to))
	evt.Emit()
}
