// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use crate::*;

pub const SC_NAME:        &str = "governance";
pub const SC_DESCRIPTION: &str = "Core governance contract";
pub const HSC_NAME:       ScHname = ScHname(0x17cf909f);

pub const PARAM_CHAIN_OWNER:              &str = "oi";
pub const PARAM_FEE_COLOR:                &str = "fc";
pub const PARAM_HNAME:                    &str = "hn";
pub const PARAM_MAX_BLOB_SIZE:            &str = "bs";
pub const PARAM_MAX_EVENT_SIZE:           &str = "es";
pub const PARAM_MAX_EVENTS_PER_REQ:       &str = "ne";
pub const PARAM_OWNER_FEE:                &str = "of";
pub const PARAM_STATE_CONTROLLER_ADDRESS: &str = "S";
pub const PARAM_VALIDATOR_FEE:            &str = "vf";

pub const RESULT_ALLOWED_STATE_CONTROLLER_ADDRESSES: &str = "a";
pub const RESULT_CHAIN_ID:                           &str = "c";
pub const RESULT_CHAIN_OWNER_ID:                     &str = "o";
pub const RESULT_DEFAULT_OWNER_FEE:                  &str = "do";
pub const RESULT_DEFAULT_VALIDATOR_FEE:              &str = "dv";
pub const RESULT_DESCRIPTION:                        &str = "d";
pub const RESULT_FEE_COLOR:                          &str = "f";
pub const RESULT_MAX_BLOB_SIZE:                      &str = "mb";
pub const RESULT_MAX_EVENT_SIZE:                     &str = "me";
pub const RESULT_MAX_EVENTS_PER_REQ:                 &str = "mr";
pub const RESULT_OWNER_FEE:                          &str = "of";
pub const RESULT_VALIDATOR_FEE:                      &str = "vf";

pub const FUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS:    &str = "addAllowedStateControllerAddress";
pub const FUNC_CLAIM_CHAIN_OWNERSHIP:                   &str = "claimChainOwnership";
pub const FUNC_DELEGATE_CHAIN_OWNERSHIP:                &str = "delegateChainOwnership";
pub const FUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS: &str = "removeAllowedStateControllerAddress";
pub const FUNC_ROTATE_STATE_CONTROLLER:                 &str = "rotateStateController";
pub const FUNC_SET_CHAIN_INFO:                          &str = "setChainInfo";
pub const FUNC_SET_CONTRACT_FEE:                        &str = "setContractFee";
pub const FUNC_SET_DEFAULT_FEE:                         &str = "setDefaultFee";
pub const VIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES:  &str = "getAllowedStateControllerAddresses";
pub const VIEW_GET_CHAIN_INFO:                          &str = "getChainInfo";
pub const VIEW_GET_FEE_INFO:                            &str = "getFeeInfo";
pub const VIEW_GET_MAX_BLOB_SIZE:                       &str = "getMaxBlobSize";

pub const HFUNC_ADD_ALLOWED_STATE_CONTROLLER_ADDRESS:    ScHname = ScHname(0x9469d567);
pub const HFUNC_CLAIM_CHAIN_OWNERSHIP:                   ScHname = ScHname(0x03ff0fc0);
pub const HFUNC_DELEGATE_CHAIN_OWNERSHIP:                ScHname = ScHname(0x93ecb6ad);
pub const HFUNC_REMOVE_ALLOWED_STATE_CONTROLLER_ADDRESS: ScHname = ScHname(0x31f69447);
pub const HFUNC_ROTATE_STATE_CONTROLLER:                 ScHname = ScHname(0x244d1038);
pub const HFUNC_SET_CHAIN_INFO:                          ScHname = ScHname(0x702f5d2b);
pub const HFUNC_SET_CONTRACT_FEE:                        ScHname = ScHname(0x8421a42b);
pub const HFUNC_SET_DEFAULT_FEE:                         ScHname = ScHname(0x3310ecd0);
pub const HVIEW_GET_ALLOWED_STATE_CONTROLLER_ADDRESSES:  ScHname = ScHname(0xf3505183);
pub const HVIEW_GET_CHAIN_INFO:                          ScHname = ScHname(0x434477e2);
pub const HVIEW_GET_FEE_INFO:                            ScHname = ScHname(0x9fe54b48);
pub const HVIEW_GET_MAX_BLOB_SIZE:                       ScHname = ScHname(0xe1db3d28);

// @formatter:on
