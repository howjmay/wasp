// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ApproveCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncApprove);
	params: sc.MutableApproveParams = new sc.MutableApproveParams();
}

export class ApproveContext {
	events:  sc.Erc20Events = new sc.Erc20Events();
	params: sc.ImmutableApproveParams = new sc.ImmutableApproveParams();
	state: sc.MutableErc20State = new sc.MutableErc20State();
}

export class InitCall {
	func: wasmlib.ScInitFunc = new wasmlib.ScInitFunc(sc.HScName, sc.HFuncInit);
	params: sc.MutableInitParams = new sc.MutableInitParams();
}

export class InitContext {
	events:  sc.Erc20Events = new sc.Erc20Events();
	params: sc.ImmutableInitParams = new sc.ImmutableInitParams();
	state: sc.MutableErc20State = new sc.MutableErc20State();
}

export class TransferCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTransfer);
	params: sc.MutableTransferParams = new sc.MutableTransferParams();
}

export class TransferContext {
	events:  sc.Erc20Events = new sc.Erc20Events();
	params: sc.ImmutableTransferParams = new sc.ImmutableTransferParams();
	state: sc.MutableErc20State = new sc.MutableErc20State();
}

export class TransferFromCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTransferFrom);
	params: sc.MutableTransferFromParams = new sc.MutableTransferFromParams();
}

export class TransferFromContext {
	events:  sc.Erc20Events = new sc.Erc20Events();
	params: sc.ImmutableTransferFromParams = new sc.ImmutableTransferFromParams();
	state: sc.MutableErc20State = new sc.MutableErc20State();
}

export class AllowanceCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewAllowance);
	params: sc.MutableAllowanceParams = new sc.MutableAllowanceParams();
	results: sc.ImmutableAllowanceResults = new sc.ImmutableAllowanceResults();
}

export class AllowanceContext {
	params: sc.ImmutableAllowanceParams = new sc.ImmutableAllowanceParams();
	results: sc.MutableAllowanceResults = new sc.MutableAllowanceResults();
	state: sc.ImmutableErc20State = new sc.ImmutableErc20State();
}

export class BalanceOfCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewBalanceOf);
	params: sc.MutableBalanceOfParams = new sc.MutableBalanceOfParams();
	results: sc.ImmutableBalanceOfResults = new sc.ImmutableBalanceOfResults();
}

export class BalanceOfContext {
	params: sc.ImmutableBalanceOfParams = new sc.ImmutableBalanceOfParams();
	results: sc.MutableBalanceOfResults = new sc.MutableBalanceOfResults();
	state: sc.ImmutableErc20State = new sc.ImmutableErc20State();
}

export class TotalSupplyCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewTotalSupply);
	results: sc.ImmutableTotalSupplyResults = new sc.ImmutableTotalSupplyResults();
}

export class TotalSupplyContext {
	results: sc.MutableTotalSupplyResults = new sc.MutableTotalSupplyResults();
	state: sc.ImmutableErc20State = new sc.ImmutableErc20State();
}

export class ScFuncs {
    static approve(ctx: wasmlib.ScFuncCallContext): ApproveCall {
        let f = new ApproveCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static init(ctx: wasmlib.ScFuncCallContext): InitCall {
        let f = new InitCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static transfer(ctx: wasmlib.ScFuncCallContext): TransferCall {
        let f = new TransferCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static transferFrom(ctx: wasmlib.ScFuncCallContext): TransferFromCall {
        let f = new TransferFromCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static allowance(ctx: wasmlib.ScViewCallContext): AllowanceCall {
        let f = new AllowanceCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static balanceOf(ctx: wasmlib.ScViewCallContext): BalanceOfCall {
        let f = new BalanceOfCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static totalSupply(ctx: wasmlib.ScViewCallContext): TotalSupplyCall {
        let f = new TotalSupplyCall();
        f.func.setPtrs(null, f.results);
        return f;
    }
}
