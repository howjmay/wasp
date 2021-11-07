// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use crate::*;

pub const SC_NAME:        &str = "root";
pub const SC_DESCRIPTION: &str = "Core root contract";
pub const HSC_NAME:       ScHname = ScHname(0xcebf5908);

pub(crate) const PARAM_DEPLOYER: &str = "dp";
pub(crate) const PARAM_DESCRIPTION: &str = "ds";
pub(crate) const PARAM_HNAME: &str = "hn";
pub(crate) const PARAM_NAME: &str = "nm";
pub(crate) const PARAM_PROGRAM_HASH: &str = "ph";

pub(crate) const RESULT_CONTRACT_FOUND: &str = "cf";
pub(crate) const RESULT_CONTRACT_REC_DATA: &str = "dt";
pub(crate) const RESULT_CONTRACT_REGISTRY: &str = "r";

pub(crate) const FUNC_DEPLOY_CONTRACT:  &str = "deployContract";
pub(crate) const FUNC_GRANT_DEPLOY_PERMISSION:  &str = "grantDeployPermission";
pub(crate) const FUNC_REVOKE_DEPLOY_PERMISSION:  &str = "revokeDeployPermission";
pub(crate) const VIEW_FIND_CONTRACT:  &str = "findContract";
pub(crate) const VIEW_GET_CONTRACT_RECORDS:  &str = "getContractRecords";

pub(crate) const HFUNC_DEPLOY_CONTRACT: ScHname = ScHname(0x28232c27);
pub(crate) const HFUNC_GRANT_DEPLOY_PERMISSION: ScHname = ScHname(0xf440263a);
pub(crate) const HFUNC_REVOKE_DEPLOY_PERMISSION: ScHname = ScHname(0x850744f1);
pub(crate) const HVIEW_FIND_CONTRACT: ScHname = ScHname(0xc145ca00);
pub(crate) const HVIEW_GET_CONTRACT_RECORDS: ScHname = ScHname(0x078b3ef3);
