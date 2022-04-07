// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

use crate::*;

#[derive(Clone)]
pub struct MapTokenIDToImmutableToken {
	pub(crate) proxy: Proxy,
}

impl MapTokenIDToImmutableToken {
    pub fn get_token(&self, key: &ScTokenID) -> ImmutableToken {
        ImmutableToken { proxy: self.proxy.key(&token_id_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct ArrayOfImmutableTokenID {
	pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableTokenID {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_token_id(&self, index: u32) -> ScImmutableTokenID {
        ScImmutableTokenID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct ImmutableTokenRegistryState {
	pub(crate) proxy: Proxy,
}

impl ImmutableTokenRegistryState {
    pub fn registry(&self) -> MapTokenIDToImmutableToken {
		MapTokenIDToImmutableToken { proxy: self.proxy.root(STATE_REGISTRY) }
	}

    pub fn token_list(&self) -> ArrayOfImmutableTokenID {
		ArrayOfImmutableTokenID { proxy: self.proxy.root(STATE_TOKEN_LIST) }
	}
}

#[derive(Clone)]
pub struct MapTokenIDToMutableToken {
	pub(crate) proxy: Proxy,
}

impl MapTokenIDToMutableToken {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_token(&self, key: &ScTokenID) -> MutableToken {
        MutableToken { proxy: self.proxy.key(&token_id_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct ArrayOfMutableTokenID {
	pub(crate) proxy: Proxy,
}

impl ArrayOfMutableTokenID {
	pub fn append_token_id(&self) -> ScMutableTokenID {
		ScMutableTokenID::new(self.proxy.append())
	}

	pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_token_id(&self, index: u32) -> ScMutableTokenID {
        ScMutableTokenID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct MutableTokenRegistryState {
	pub(crate) proxy: Proxy,
}

impl MutableTokenRegistryState {
    pub fn as_immutable(&self) -> ImmutableTokenRegistryState {
		ImmutableTokenRegistryState { proxy: self.proxy.root("") }
	}

    pub fn registry(&self) -> MapTokenIDToMutableToken {
		MapTokenIDToMutableToken { proxy: self.proxy.root(STATE_REGISTRY) }
	}

    pub fn token_list(&self) -> ArrayOfMutableTokenID {
		ArrayOfMutableTokenID { proxy: self.proxy.root(STATE_TOKEN_LIST) }
	}
}
