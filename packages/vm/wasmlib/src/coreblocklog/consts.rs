// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use crate::*;

pub const SC_NAME:        &str = "blocklog";
pub const SC_DESCRIPTION: &str = "Core block log contract";
pub const HSC_NAME:       ScHname = ScHname(0xf538ef2b);

pub(crate) const PARAM_BLOCK_INDEX: &str = "n";
pub(crate) const PARAM_CONTRACT_HNAME: &str = "h";
pub(crate) const PARAM_FROM_BLOCK: &str = "f";
pub(crate) const PARAM_REQUEST_ID: &str = "u";
pub(crate) const PARAM_TO_BLOCK: &str = "t";

pub(crate) const RESULT_BLOCK_INDEX: &str = "n";
pub(crate) const RESULT_BLOCK_INFO: &str = "i";
pub(crate) const RESULT_EVENT: &str = "e";
pub(crate) const RESULT_GOVERNING_ADDRESS: &str = "g";
pub(crate) const RESULT_REQUEST_ID: &str = "u";
pub(crate) const RESULT_REQUEST_INDEX: &str = "r";
pub(crate) const RESULT_REQUEST_PROCESSED: &str = "p";
pub(crate) const RESULT_REQUEST_RECORD: &str = "d";
pub(crate) const RESULT_STATE_CONTROLLER_ADDRESS: &str = "s";

pub(crate) const VIEW_CONTROL_ADDRESSES:  &str = "controlAddresses";
pub(crate) const VIEW_GET_BLOCK_INFO:  &str = "getBlockInfo";
pub(crate) const VIEW_GET_EVENTS_FOR_BLOCK:  &str = "getEventsForBlock";
pub(crate) const VIEW_GET_EVENTS_FOR_CONTRACT:  &str = "getEventsForContract";
pub(crate) const VIEW_GET_EVENTS_FOR_REQUEST:  &str = "getEventsForRequest";
pub(crate) const VIEW_GET_LATEST_BLOCK_INFO:  &str = "getLatestBlockInfo";
pub(crate) const VIEW_GET_REQUEST_I_DS_FOR_BLOCK:  &str = "getRequestIDsForBlock";
pub(crate) const VIEW_GET_REQUEST_RECEIPT:  &str = "getRequestReceipt";
pub(crate) const VIEW_GET_REQUEST_RECEIPTS_FOR_BLOCK:  &str = "getRequestReceiptsForBlock";
pub(crate) const VIEW_IS_REQUEST_PROCESSED:  &str = "isRequestProcessed";

pub(crate) const HVIEW_CONTROL_ADDRESSES: ScHname = ScHname(0x796bd223);
pub(crate) const HVIEW_GET_BLOCK_INFO: ScHname = ScHname(0xbe89f9b3);
pub(crate) const HVIEW_GET_EVENTS_FOR_BLOCK: ScHname = ScHname(0x36232798);
pub(crate) const HVIEW_GET_EVENTS_FOR_CONTRACT: ScHname = ScHname(0x682a1922);
pub(crate) const HVIEW_GET_EVENTS_FOR_REQUEST: ScHname = ScHname(0x4f8d68e4);
pub(crate) const HVIEW_GET_LATEST_BLOCK_INFO: ScHname = ScHname(0x084a1760);
pub(crate) const HVIEW_GET_REQUEST_I_DS_FOR_BLOCK: ScHname = ScHname(0x5a20327a);
pub(crate) const HVIEW_GET_REQUEST_RECEIPT: ScHname = ScHname(0xb7f9534f);
pub(crate) const HVIEW_GET_REQUEST_RECEIPTS_FOR_BLOCK: ScHname = ScHname(0x77e3beef);
pub(crate) const HVIEW_IS_REQUEST_PROCESSED: ScHname = ScHname(0xd57d50a9);

// @formatter:on
