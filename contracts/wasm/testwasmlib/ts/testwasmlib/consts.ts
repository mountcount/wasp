// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "testwasmlib";
export const ScDescription = "Exercise several aspects of WasmLib";
export const HScName       = new wasmtypes.ScHname(0x89703a45);

export const ParamAddress     = "address";
export const ParamAgentID     = "agentID";
export const ParamBlockIndex  = "blockIndex";
export const ParamBool        = "bool";
export const ParamBytes       = "bytes";
export const ParamChainID     = "chainID";
export const ParamHash        = "hash";
export const ParamHname       = "hname";
export const ParamIndex       = "index";
export const ParamIndex0      = "index0";
export const ParamIndex1      = "index1";
export const ParamInt16       = "int16";
export const ParamInt32       = "int32";
export const ParamInt64       = "int64";
export const ParamInt8        = "int8";
export const ParamKey         = "key";
export const ParamKeyAddr     = "keyAddr";
export const ParamName        = "name";
export const ParamNameAddr    = "nameAddr";
export const ParamNftID       = "nftID";
export const ParamParam       = "this";
export const ParamRecordIndex = "recordIndex";
export const ParamRequestID   = "requestID";
export const ParamString      = "string";
export const ParamTokenID     = "tokenID";
export const ParamUint16      = "uint16";
export const ParamUint32      = "uint32";
export const ParamUint64      = "uint64";
export const ParamUint8       = "uint8";
export const ParamValue       = "value";
export const ParamValueAddr   = "valueAddr";

export const ResultCount     = "count";
export const ResultIotas     = "iotas";
export const ResultLength    = "length";
export const ResultRandom    = "random";
export const ResultRecord    = "record";
export const ResultValue     = "value";
export const ResultValueAddr = "valueAddr";

export const StateAddressMapOfAddressArray = "addressMapOfAddressArray";
export const StateAddressMapOfAddressMap   = "addressMapOfAddressMap";
export const StateArrayOfAddressArray      = "arrayOfAddressArray";
export const StateArrayOfAddressMap        = "arrayOfAddressMap";
export const StateArrayOfStringArray       = "arrayOfStringArray";
export const StateArrayOfStringMap         = "arrayOfStringMap";
export const StateLatLong                  = "latLong";
export const StateRandom                   = "random";
export const StateStringMapOfStringArray   = "stringMapOfStringArray";
export const StateStringMapOfStringMap     = "stringMapOfStringMap";

export const FuncAddressMapOfAddressArrayAppend = "addressMapOfAddressArrayAppend";
export const FuncAddressMapOfAddressArrayClear  = "addressMapOfAddressArrayClear";
export const FuncAddressMapOfAddressArraySet    = "addressMapOfAddressArraySet";
export const FuncAddressMapOfAddressMapClear    = "addressMapOfAddressMapClear";
export const FuncAddressMapOfAddressMapSet      = "addressMapOfAddressMapSet";
export const FuncArrayOfAddressArrayAppend      = "arrayOfAddressArrayAppend";
export const FuncArrayOfAddressArrayClear       = "arrayOfAddressArrayClear";
export const FuncArrayOfAddressArraySet         = "arrayOfAddressArraySet";
export const FuncArrayOfAddressMapClear         = "arrayOfAddressMapClear";
export const FuncArrayOfAddressMapSet           = "arrayOfAddressMapSet";
export const FuncArrayOfStringArrayAppend       = "arrayOfStringArrayAppend";
export const FuncArrayOfStringArrayClear        = "arrayOfStringArrayClear";
export const FuncArrayOfStringArraySet          = "arrayOfStringArraySet";
export const FuncArrayOfStringMapClear          = "arrayOfStringMapClear";
export const FuncArrayOfStringMapSet            = "arrayOfStringMapSet";
export const FuncParamTypes                     = "paramTypes";
export const FuncRandom                         = "random";
export const FuncStringMapOfStringArrayAppend   = "stringMapOfStringArrayAppend";
export const FuncStringMapOfStringArrayClear    = "stringMapOfStringArrayClear";
export const FuncStringMapOfStringArraySet      = "stringMapOfStringArraySet";
export const FuncStringMapOfStringMapClear      = "stringMapOfStringMapClear";
export const FuncStringMapOfStringMapSet        = "stringMapOfStringMapSet";
export const FuncTakeAllowance                  = "takeAllowance";
export const FuncTakeBalance                    = "takeBalance";
export const FuncTriggerEvent                   = "triggerEvent";
export const ViewAddressMapOfAddressArrayLength = "addressMapOfAddressArrayLength";
export const ViewAddressMapOfAddressArrayValue  = "addressMapOfAddressArrayValue";
export const ViewAddressMapOfAddressMapValue    = "addressMapOfAddressMapValue";
export const ViewArrayOfAddressArrayLength      = "arrayOfAddressArrayLength";
export const ViewArrayOfAddressArrayValue       = "arrayOfAddressArrayValue";
export const ViewArrayOfAddressMapValue         = "arrayOfAddressMapValue";
export const ViewArrayOfStringArrayLength       = "arrayOfStringArrayLength";
export const ViewArrayOfStringArrayValue        = "arrayOfStringArrayValue";
export const ViewArrayOfStringMapValue          = "arrayOfStringMapValue";
export const ViewBlockRecord                    = "blockRecord";
export const ViewBlockRecords                   = "blockRecords";
export const ViewGetRandom                      = "getRandom";
export const ViewIotaBalance                    = "iotaBalance";
export const ViewStringMapOfStringArrayLength   = "stringMapOfStringArrayLength";
export const ViewStringMapOfStringArrayValue    = "stringMapOfStringArrayValue";
export const ViewStringMapOfStringMapValue      = "stringMapOfStringMapValue";

export const HFuncAddressMapOfAddressArrayAppend = new wasmtypes.ScHname(0x328fea8f);
export const HFuncAddressMapOfAddressArrayClear  = new wasmtypes.ScHname(0x771b8d2d);
export const HFuncAddressMapOfAddressArraySet    = new wasmtypes.ScHname(0x9b51ab5e);
export const HFuncAddressMapOfAddressMapClear    = new wasmtypes.ScHname(0x2a5316b1);
export const HFuncAddressMapOfAddressMapSet      = new wasmtypes.ScHname(0x7736a024);
export const HFuncArrayOfAddressArrayAppend      = new wasmtypes.ScHname(0xaf6e6b46);
export const HFuncArrayOfAddressArrayClear       = new wasmtypes.ScHname(0x8967252b);
export const HFuncArrayOfAddressArraySet         = new wasmtypes.ScHname(0xc84c8afa);
export const HFuncArrayOfAddressMapClear         = new wasmtypes.ScHname(0x18720c38);
export const HFuncArrayOfAddressMapSet           = new wasmtypes.ScHname(0xd02edb28);
export const HFuncArrayOfStringArrayAppend       = new wasmtypes.ScHname(0x2f37355a);
export const HFuncArrayOfStringArrayClear        = new wasmtypes.ScHname(0xcec430af);
export const HFuncArrayOfStringArraySet          = new wasmtypes.ScHname(0xa74fa352);
export const HFuncArrayOfStringMapClear          = new wasmtypes.ScHname(0x4ed3a6d7);
export const HFuncArrayOfStringMapSet            = new wasmtypes.ScHname(0xbe6a8ae7);
export const HFuncParamTypes                     = new wasmtypes.ScHname(0x6921c4cd);
export const HFuncRandom                         = new wasmtypes.ScHname(0xe86c97ca);
export const HFuncStringMapOfStringArrayAppend   = new wasmtypes.ScHname(0x414f806d);
export const HFuncStringMapOfStringArrayClear    = new wasmtypes.ScHname(0xe706a05f);
export const HFuncStringMapOfStringArraySet      = new wasmtypes.ScHname(0x5acf713b);
export const HFuncStringMapOfStringMapClear      = new wasmtypes.ScHname(0xf8edb8cb);
export const HFuncStringMapOfStringMapSet        = new wasmtypes.ScHname(0xa28e5cbb);
export const HFuncTakeAllowance                  = new wasmtypes.ScHname(0x91e7bd00);
export const HFuncTakeBalance                    = new wasmtypes.ScHname(0x8ad1cb27);
export const HFuncTriggerEvent                   = new wasmtypes.ScHname(0xd5438ac6);
export const HViewAddressMapOfAddressArrayLength = new wasmtypes.ScHname(0x30ecea55);
export const HViewAddressMapOfAddressArrayValue  = new wasmtypes.ScHname(0x2775d926);
export const HViewAddressMapOfAddressMapValue    = new wasmtypes.ScHname(0x3e49b429);
export const HViewArrayOfAddressArrayLength      = new wasmtypes.ScHname(0xc22488ab);
export const HViewArrayOfAddressArrayValue       = new wasmtypes.ScHname(0xb67ef99e);
export const HViewArrayOfAddressMapValue         = new wasmtypes.ScHname(0xc9ae7aec);
export const HViewArrayOfStringArrayLength       = new wasmtypes.ScHname(0x3eb0d035);
export const HViewArrayOfStringArrayValue        = new wasmtypes.ScHname(0xf2b0f6c8);
export const HViewArrayOfStringMapValue          = new wasmtypes.ScHname(0x8b35543c);
export const HViewBlockRecord                    = new wasmtypes.ScHname(0xad13b2f8);
export const HViewBlockRecords                   = new wasmtypes.ScHname(0x16e249ea);
export const HViewGetRandom                      = new wasmtypes.ScHname(0x46263045);
export const HViewIotaBalance                    = new wasmtypes.ScHname(0x9d3920bd);
export const HViewStringMapOfStringArrayLength   = new wasmtypes.ScHname(0x9f433699);
export const HViewStringMapOfStringArrayValue    = new wasmtypes.ScHname(0xb263e298);
export const HViewStringMapOfStringMapValue      = new wasmtypes.ScHname(0xb93f2535);
