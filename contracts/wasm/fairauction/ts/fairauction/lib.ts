// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

const exportMap: wasmlib.ScExportMap = {
	names: [
		sc.FuncFinalizeAuction,
		sc.FuncPlaceBid,
		sc.FuncSetOwnerMargin,
		sc.FuncStartAuction,
		sc.ViewGetInfo,
	],
	funcs: [
		funcFinalizeAuctionThunk,
		funcPlaceBidThunk,
		funcSetOwnerMarginThunk,
		funcStartAuctionThunk,
	],
	views: [
		viewGetInfoThunk,
	],
};

export function on_call(index: i32): void {
	wasmlib.ScExports.call(index, exportMap);
}

export function on_load(): void {
	wasmlib.ScExports.export(exportMap);
}

function funcFinalizeAuctionThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcFinalizeAuction");
	let f = new sc.FinalizeAuctionContext();

	// only SC itself can invoke this function
	ctx.require(ctx.caller().equals(ctx.accountID()), "no permission");

	ctx.require(f.params.token().exists(), "missing mandatory token");
	sc.funcFinalizeAuction(ctx, f);
	ctx.log("fairauction.funcFinalizeAuction ok");
}

function funcPlaceBidThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcPlaceBid");
	let f = new sc.PlaceBidContext();
	ctx.require(f.params.token().exists(), "missing mandatory token");
	sc.funcPlaceBid(ctx, f);
	ctx.log("fairauction.funcPlaceBid ok");
}

function funcSetOwnerMarginThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcSetOwnerMargin");
	let f = new sc.SetOwnerMarginContext();

	// only SC creator can set owner margin
	ctx.require(ctx.caller().equals(ctx.contractCreator()), "no permission");

	ctx.require(f.params.ownerMargin().exists(), "missing mandatory ownerMargin");
	sc.funcSetOwnerMargin(ctx, f);
	ctx.log("fairauction.funcSetOwnerMargin ok");
}

function funcStartAuctionThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcStartAuction");
	let f = new sc.StartAuctionContext();
	ctx.require(f.params.minimumBid().exists(), "missing mandatory minimumBid");
	ctx.require(f.params.token().exists(), "missing mandatory token");
	sc.funcStartAuction(ctx, f);
	ctx.log("fairauction.funcStartAuction ok");
}

function viewGetInfoThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("fairauction.viewGetInfo");
	let f = new sc.GetInfoContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetInfoResults(results.asProxy());
	ctx.require(f.params.token().exists(), "missing mandatory token");
	sc.viewGetInfo(ctx, f);
	ctx.results(results);
	ctx.log("fairauction.viewGetInfo ok");
}
