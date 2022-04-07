// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreroot

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableFindContractResults struct {
	proxy wasmtypes.Proxy
}

// encoded contract record
func (s ImmutableFindContractResults) ContractFound() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ResultContractFound))
}

// encoded contract record
func (s ImmutableFindContractResults) ContractRecData() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ResultContractRecData))
}

type MutableFindContractResults struct {
	proxy wasmtypes.Proxy
}

// encoded contract record
func (s MutableFindContractResults) ContractFound() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ResultContractFound))
}

// encoded contract record
func (s MutableFindContractResults) ContractRecData() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ResultContractRecData))
}

type MapHnameToImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapHnameToImmutableBytes) GetBytes(key wasmtypes.ScHname) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.proxy.Key(wasmtypes.HnameToBytes(key)))
}

type ImmutableGetContractRecordsResults struct {
	proxy wasmtypes.Proxy
}

// contract records
func (s ImmutableGetContractRecordsResults) ContractRegistry() MapHnameToImmutableBytes {
	return MapHnameToImmutableBytes{proxy: s.proxy.Root(ResultContractRegistry)}
}

type MapHnameToMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapHnameToMutableBytes) Clear() {
	m.proxy.ClearMap()
}

func (m MapHnameToMutableBytes) GetBytes(key wasmtypes.ScHname) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.proxy.Key(wasmtypes.HnameToBytes(key)))
}

type MutableGetContractRecordsResults struct {
	proxy wasmtypes.Proxy
}

// contract records
func (s MutableGetContractRecordsResults) ContractRegistry() MapHnameToMutableBytes {
	return MapHnameToMutableBytes{proxy: s.proxy.Root(ResultContractRegistry)}
}
