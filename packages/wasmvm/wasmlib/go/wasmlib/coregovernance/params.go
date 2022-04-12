// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coregovernance

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableAddAllowedStateControllerAddressParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddAllowedStateControllerAddressParams) ChainOwner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamChainOwner))
}

func (s ImmutableAddAllowedStateControllerAddressParams) StateControllerAddress() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamStateControllerAddress))
}

type MutableAddAllowedStateControllerAddressParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddAllowedStateControllerAddressParams) ChainOwner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamChainOwner))
}

func (s MutableAddAllowedStateControllerAddressParams) StateControllerAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamStateControllerAddress))
}

type ImmutableDelegateChainOwnershipParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableDelegateChainOwnershipParams) ChainOwner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamChainOwner))
}

type MutableDelegateChainOwnershipParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableDelegateChainOwnershipParams) ChainOwner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamChainOwner))
}

type ImmutableRemoveAllowedStateControllerAddressParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableRemoveAllowedStateControllerAddressParams) StateControllerAddress() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamStateControllerAddress))
}

type MutableRemoveAllowedStateControllerAddressParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableRemoveAllowedStateControllerAddressParams) StateControllerAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamStateControllerAddress))
}

type ImmutableRotateStateControllerParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableRotateStateControllerParams) StateControllerAddress() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamStateControllerAddress))
}

type MutableRotateStateControllerParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableRotateStateControllerParams) StateControllerAddress() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamStateControllerAddress))
}

type ImmutableSetChainInfoParams struct {
	proxy wasmtypes.Proxy
}

// default no change
func (s ImmutableSetChainInfoParams) MaxBlobSize() wasmtypes.ScImmutableInt32 {
	return wasmtypes.NewScImmutableInt32(s.proxy.Root(ParamMaxBlobSize))
}

// default no change
func (s ImmutableSetChainInfoParams) MaxEventSize() wasmtypes.ScImmutableInt16 {
	return wasmtypes.NewScImmutableInt16(s.proxy.Root(ParamMaxEventSize))
}

// default no change
func (s ImmutableSetChainInfoParams) MaxEventsPerReq() wasmtypes.ScImmutableInt16 {
	return wasmtypes.NewScImmutableInt16(s.proxy.Root(ParamMaxEventsPerReq))
}

// default no change
func (s ImmutableSetChainInfoParams) OwnerFee() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamOwnerFee))
}

// default no change
func (s ImmutableSetChainInfoParams) ValidatorFee() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamValidatorFee))
}

type MutableSetChainInfoParams struct {
	proxy wasmtypes.Proxy
}

// default no change
func (s MutableSetChainInfoParams) MaxBlobSize() wasmtypes.ScMutableInt32 {
	return wasmtypes.NewScMutableInt32(s.proxy.Root(ParamMaxBlobSize))
}

// default no change
func (s MutableSetChainInfoParams) MaxEventSize() wasmtypes.ScMutableInt16 {
	return wasmtypes.NewScMutableInt16(s.proxy.Root(ParamMaxEventSize))
}

// default no change
func (s MutableSetChainInfoParams) MaxEventsPerReq() wasmtypes.ScMutableInt16 {
	return wasmtypes.NewScMutableInt16(s.proxy.Root(ParamMaxEventsPerReq))
}

// default no change
func (s MutableSetChainInfoParams) OwnerFee() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamOwnerFee))
}

// default no change
func (s MutableSetChainInfoParams) ValidatorFee() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamValidatorFee))
}

type ImmutableSetContractFeeParams struct {
	proxy wasmtypes.Proxy
}

// contract id
func (s ImmutableSetContractFeeParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

// default 0 (no fee)
func (s ImmutableSetContractFeeParams) OwnerFee() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamOwnerFee))
}

// default 0 (no fee)
func (s ImmutableSetContractFeeParams) ValidatorFee() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamValidatorFee))
}

type MutableSetContractFeeParams struct {
	proxy wasmtypes.Proxy
}

// contract id
func (s MutableSetContractFeeParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}

// default 0 (no fee)
func (s MutableSetContractFeeParams) OwnerFee() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamOwnerFee))
}

// default 0 (no fee)
func (s MutableSetContractFeeParams) ValidatorFee() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamValidatorFee))
}

type ImmutableSetDefaultFeeParams struct {
	proxy wasmtypes.Proxy
}

// default -1 (not set)
func (s ImmutableSetDefaultFeeParams) OwnerFee() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamOwnerFee))
}

// default -1 (not set)
func (s ImmutableSetDefaultFeeParams) ValidatorFee() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamValidatorFee))
}

type MutableSetDefaultFeeParams struct {
	proxy wasmtypes.Proxy
}

// default -1 (not set)
func (s MutableSetDefaultFeeParams) OwnerFee() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamOwnerFee))
}

// default -1 (not set)
func (s MutableSetDefaultFeeParams) ValidatorFee() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamValidatorFee))
}

type ImmutableGetFeeInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetFeeInfoParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

type MutableGetFeeInfoParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetFeeInfoParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}
