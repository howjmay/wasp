// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use std::ptr;

use wasmlib::*;

use crate::consts::*;
use crate::results::*;

pub struct NowCall {
	pub func: ScFunc,
}

pub struct GetTimestampCall {
	pub func: ScView,
	pub results: ImmutableGetTimestampResults,
}

pub struct ScFuncs {
}

impl ScFuncs {
    pub fn now(_ctx: & dyn ScFuncCallContext) -> NowCall {
        NowCall {
            func: ScFunc::new(HSC_NAME, HFUNC_NOW),
        }
    }
    pub fn get_timestamp(_ctx: & dyn ScViewCallContext) -> GetTimestampCall {
        let mut f = GetTimestampCall {
            func: ScView::new(HSC_NAME, HVIEW_GET_TIMESTAMP),
            results: ImmutableGetTimestampResults { id: 0 },
        };
        f.func.set_ptrs(ptr::null_mut(), &mut f.results.id);
        f
    }
}
