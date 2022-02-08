// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type ForcePayoutCall struct {
	Func    *wasmlib.ScFunc
}

type ForceResetCall struct {
	Func    *wasmlib.ScFunc
}

type PayWinnersCall struct {
	Func    *wasmlib.ScFunc
}

type PlaceBetCall struct {
	Func    *wasmlib.ScFunc
	Params  MutablePlaceBetParams
}

type PlayPeriodCall struct {
	Func    *wasmlib.ScFunc
	Params  MutablePlayPeriodParams
}

type LastWinningNumberCall struct {
	Func    *wasmlib.ScView
	Results ImmutableLastWinningNumberResults
}

type RoundNumberCall struct {
	Func    *wasmlib.ScView
	Results ImmutableRoundNumberResults
}

type RoundStartedAtCall struct {
	Func    *wasmlib.ScView
	Results ImmutableRoundStartedAtResults
}

type RoundStatusCall struct {
	Func    *wasmlib.ScView
	Results ImmutableRoundStatusResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) ForcePayout(ctx wasmlib.ScFuncCallContext) *ForcePayoutCall {
	return &ForcePayoutCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncForcePayout)}
}

func (sc Funcs) ForceReset(ctx wasmlib.ScFuncCallContext) *ForceResetCall {
	return &ForceResetCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncForceReset)}
}

func (sc Funcs) PayWinners(ctx wasmlib.ScFuncCallContext) *PayWinnersCall {
	return &PayWinnersCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncPayWinners)}
}

func (sc Funcs) PlaceBet(ctx wasmlib.ScFuncCallContext) *PlaceBetCall {
	f := &PlaceBetCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncPlaceBet)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) PlayPeriod(ctx wasmlib.ScFuncCallContext) *PlayPeriodCall {
	f := &PlayPeriodCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncPlayPeriod)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) LastWinningNumber(ctx wasmlib.ScViewCallContext) *LastWinningNumberCall {
	f := &LastWinningNumberCall{Func: wasmlib.NewScView(ctx, HScName, HViewLastWinningNumber)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) RoundNumber(ctx wasmlib.ScViewCallContext) *RoundNumberCall {
	f := &RoundNumberCall{Func: wasmlib.NewScView(ctx, HScName, HViewRoundNumber)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) RoundStartedAt(ctx wasmlib.ScViewCallContext) *RoundStartedAtCall {
	f := &RoundStartedAtCall{Func: wasmlib.NewScView(ctx, HScName, HViewRoundStartedAt)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) RoundStatus(ctx wasmlib.ScViewCallContext) *RoundStatusCall {
	f := &RoundStatusCall{Func: wasmlib.NewScView(ctx, HScName, HViewRoundStatus)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}
