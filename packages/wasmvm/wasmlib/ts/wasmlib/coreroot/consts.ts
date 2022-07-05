// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "root";
export const ScDescription = "Root Contract";
export const HScName       = new wasmtypes.ScHname(0xcebf5908);

export const ParamCloseFunc                = "bcc";
export const ParamDeployPermissionsEnabled = "de";
export const ParamDeployer                 = "dp";
export const ParamDescription              = "ds";
export const ParamHname                    = "hn";
export const ParamName                     = "nm";
export const ParamOpenFunc                 = "bco";
export const ParamProgramHash              = "ph";

export const ResultContractFound    = "cf";
export const ResultContractRecData  = "dt";
export const ResultContractRegistry = "r";

export const FuncDeployContract           = "deployContract";
export const FuncGrantDeployPermission    = "grantDeployPermission";
export const FuncRequireDeployPermissions = "requireDeployPermissions";
export const FuncRevokeDeployPermission   = "revokeDeployPermission";
export const FuncSubscribeBlockContext    = "subscribeBlockContext";
export const ViewFindContract             = "findContract";
export const ViewGetContractRecords       = "getContractRecords";

export const HFuncDeployContract           = new wasmtypes.ScHname(0x28232c27);
export const HFuncGrantDeployPermission    = new wasmtypes.ScHname(0xf440263a);
export const HFuncRequireDeployPermissions = new wasmtypes.ScHname(0xefff8d83);
export const HFuncRevokeDeployPermission   = new wasmtypes.ScHname(0x850744f1);
export const HFuncSubscribeBlockContext    = new wasmtypes.ScHname(0xf2f8a06d);
export const HViewFindContract             = new wasmtypes.ScHname(0xc145ca00);
export const HViewGetContractRecords       = new wasmtypes.ScHname(0x078b3ef3);
