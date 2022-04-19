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
pub struct ImmutableArrayOfArraysAppendParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableArrayOfArraysAppendParams {
    pub fn index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn value(&self) -> ArrayOfImmutableString {
		ArrayOfImmutableString { proxy: self.proxy.root(PARAM_VALUE) }
	}
}

#[derive(Clone)]
pub struct MutableArrayOfArraysAppendParams {
	pub(crate) proxy: Proxy,
}

impl MutableArrayOfArraysAppendParams {
    pub fn index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn value(&self) -> ArrayOfMutableString {
		ArrayOfMutableString { proxy: self.proxy.root(PARAM_VALUE) }
	}
}

#[derive(Clone)]
pub struct ImmutableArrayOfArraysSetParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableArrayOfArraysSetParams {
    pub fn index0(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX0))
	}

    pub fn index1(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX1))
	}

    pub fn value(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableArrayOfArraysSetParams {
	pub(crate) proxy: Proxy,
}

impl MutableArrayOfArraysSetParams {
    pub fn index0(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX0))
	}

    pub fn index1(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX1))
	}

    pub fn value(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableArrayOfMapsSetParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableArrayOfMapsSetParams {
    pub fn index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn key(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_KEY))
	}

    pub fn value(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableArrayOfMapsSetParams {
	pub(crate) proxy: Proxy,
}

impl MutableArrayOfMapsSetParams {
    pub fn index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn key(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_KEY))
	}

    pub fn value(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfArraysAppendParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfArraysAppendParams {
    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn value(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableMapOfArraysAppendParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfArraysAppendParams {
    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn value(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfArraysClearParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfArraysClearParams {
    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableMapOfArraysClearParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfArraysClearParams {
    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfArraysSetParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfArraysSetParams {
    pub fn index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn value(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableMapOfArraysSetParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfArraysSetParams {
    pub fn index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn value(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfMapsClearParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfMapsClearParams {
    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableMapOfMapsClearParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfMapsClearParams {
    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfMapsSetParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfMapsSetParams {
    pub fn key(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_KEY))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn value(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct MutableMapOfMapsSetParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfMapsSetParams {
    pub fn key(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_KEY))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn value(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_VALUE))
	}
}

#[derive(Clone)]
pub struct MapStringToImmutableBytes {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableBytes {
    pub fn get_bytes(&self, key: &str) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableParamTypesParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableParamTypesParams {
    pub fn address(&self) -> ScImmutableAddress {
		ScImmutableAddress::new(self.proxy.root(PARAM_ADDRESS))
	}

    pub fn agent_id(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
	}

    pub fn bool(&self) -> ScImmutableBool {
		ScImmutableBool::new(self.proxy.root(PARAM_BOOL))
	}

    pub fn bytes(&self) -> ScImmutableBytes {
		ScImmutableBytes::new(self.proxy.root(PARAM_BYTES))
	}

    pub fn chain_id(&self) -> ScImmutableChainID {
		ScImmutableChainID::new(self.proxy.root(PARAM_CHAIN_ID))
	}

    pub fn color(&self) -> ScImmutableColor {
		ScImmutableColor::new(self.proxy.root(PARAM_COLOR))
	}

    pub fn hash(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_HASH))
	}

    pub fn hname(&self) -> ScImmutableHname {
		ScImmutableHname::new(self.proxy.root(PARAM_HNAME))
	}

    pub fn int16(&self) -> ScImmutableInt16 {
		ScImmutableInt16::new(self.proxy.root(PARAM_INT16))
	}

    pub fn int32(&self) -> ScImmutableInt32 {
		ScImmutableInt32::new(self.proxy.root(PARAM_INT32))
	}

    pub fn int64(&self) -> ScImmutableInt64 {
		ScImmutableInt64::new(self.proxy.root(PARAM_INT64))
	}

    pub fn int8(&self) -> ScImmutableInt8 {
		ScImmutableInt8::new(self.proxy.root(PARAM_INT8))
	}

    pub fn param(&self) -> MapStringToImmutableBytes {
		MapStringToImmutableBytes { proxy: self.proxy.clone() }
	}

    pub fn request_id(&self) -> ScImmutableRequestID {
		ScImmutableRequestID::new(self.proxy.root(PARAM_REQUEST_ID))
	}

    pub fn string(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_STRING))
	}

    pub fn uint16(&self) -> ScImmutableUint16 {
		ScImmutableUint16::new(self.proxy.root(PARAM_UINT16))
	}

    pub fn uint32(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_UINT32))
	}

    pub fn uint64(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_UINT64))
	}

    pub fn uint8(&self) -> ScImmutableUint8 {
		ScImmutableUint8::new(self.proxy.root(PARAM_UINT8))
	}
}

#[derive(Clone)]
pub struct MapStringToMutableBytes {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableBytes {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_bytes(&self, key: &str) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableParamTypesParams {
	pub(crate) proxy: Proxy,
}

impl MutableParamTypesParams {
    pub fn address(&self) -> ScMutableAddress {
		ScMutableAddress::new(self.proxy.root(PARAM_ADDRESS))
	}

    pub fn agent_id(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
	}

    pub fn bool(&self) -> ScMutableBool {
		ScMutableBool::new(self.proxy.root(PARAM_BOOL))
	}

    pub fn bytes(&self) -> ScMutableBytes {
		ScMutableBytes::new(self.proxy.root(PARAM_BYTES))
	}

    pub fn chain_id(&self) -> ScMutableChainID {
		ScMutableChainID::new(self.proxy.root(PARAM_CHAIN_ID))
	}

    pub fn color(&self) -> ScMutableColor {
		ScMutableColor::new(self.proxy.root(PARAM_COLOR))
	}

    pub fn hash(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_HASH))
	}

    pub fn hname(&self) -> ScMutableHname {
		ScMutableHname::new(self.proxy.root(PARAM_HNAME))
	}

    pub fn int16(&self) -> ScMutableInt16 {
		ScMutableInt16::new(self.proxy.root(PARAM_INT16))
	}

    pub fn int32(&self) -> ScMutableInt32 {
		ScMutableInt32::new(self.proxy.root(PARAM_INT32))
	}

    pub fn int64(&self) -> ScMutableInt64 {
		ScMutableInt64::new(self.proxy.root(PARAM_INT64))
	}

    pub fn int8(&self) -> ScMutableInt8 {
		ScMutableInt8::new(self.proxy.root(PARAM_INT8))
	}

    pub fn param(&self) -> MapStringToMutableBytes {
		MapStringToMutableBytes { proxy: self.proxy.clone() }
	}

    pub fn request_id(&self) -> ScMutableRequestID {
		ScMutableRequestID::new(self.proxy.root(PARAM_REQUEST_ID))
	}

    pub fn string(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_STRING))
	}

    pub fn uint16(&self) -> ScMutableUint16 {
		ScMutableUint16::new(self.proxy.root(PARAM_UINT16))
	}

    pub fn uint32(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_UINT32))
	}

    pub fn uint64(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_UINT64))
	}

    pub fn uint8(&self) -> ScMutableUint8 {
		ScMutableUint8::new(self.proxy.root(PARAM_UINT8))
	}
}

#[derive(Clone)]
pub struct ImmutableTriggerEventParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableTriggerEventParams {
    pub fn address(&self) -> ScImmutableAddress {
		ScImmutableAddress::new(self.proxy.root(PARAM_ADDRESS))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableTriggerEventParams {
	pub(crate) proxy: Proxy,
}

impl MutableTriggerEventParams {
    pub fn address(&self) -> ScMutableAddress {
		ScMutableAddress::new(self.proxy.root(PARAM_ADDRESS))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct ImmutableArrayOfArraysValueParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableArrayOfArraysValueParams {
    pub fn index0(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX0))
	}

    pub fn index1(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX1))
	}
}

#[derive(Clone)]
pub struct MutableArrayOfArraysValueParams {
	pub(crate) proxy: Proxy,
}

impl MutableArrayOfArraysValueParams {
    pub fn index0(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX0))
	}

    pub fn index1(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX1))
	}
}

#[derive(Clone)]
pub struct ImmutableArrayOfMapsValueParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableArrayOfMapsValueParams {
    pub fn index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn key(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_KEY))
	}
}

#[derive(Clone)]
pub struct MutableArrayOfMapsValueParams {
	pub(crate) proxy: Proxy,
}

impl MutableArrayOfMapsValueParams {
    pub fn index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn key(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_KEY))
	}
}

#[derive(Clone)]
pub struct ImmutableBlockRecordParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableBlockRecordParams {
    pub fn block_index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_BLOCK_INDEX))
	}

    pub fn record_index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_RECORD_INDEX))
	}
}

#[derive(Clone)]
pub struct MutableBlockRecordParams {
	pub(crate) proxy: Proxy,
}

impl MutableBlockRecordParams {
    pub fn block_index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_BLOCK_INDEX))
	}

    pub fn record_index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_RECORD_INDEX))
	}
}

#[derive(Clone)]
pub struct ImmutableBlockRecordsParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableBlockRecordsParams {
    pub fn block_index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_BLOCK_INDEX))
	}
}

#[derive(Clone)]
pub struct MutableBlockRecordsParams {
	pub(crate) proxy: Proxy,
}

impl MutableBlockRecordsParams {
    pub fn block_index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_BLOCK_INDEX))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfArraysLengthParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfArraysLengthParams {
    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableMapOfArraysLengthParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfArraysLengthParams {
    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfArraysValueParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfArraysValueParams {
    pub fn index(&self) -> ScImmutableUint32 {
		ScImmutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableMapOfArraysValueParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfArraysValueParams {
    pub fn index(&self) -> ScMutableUint32 {
		ScMutableUint32::new(self.proxy.root(PARAM_INDEX))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct ImmutableMapOfMapsValueParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableMapOfMapsValueParams {
    pub fn key(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_KEY))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableMapOfMapsValueParams {
	pub(crate) proxy: Proxy,
}

impl MutableMapOfMapsValueParams {
    pub fn key(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_KEY))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}
