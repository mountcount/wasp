// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableCallOnChainParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableCallOnChainParams) HnameContract() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHnameContract))
}

func (s ImmutableCallOnChainParams) HnameEP() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHnameEP))
}

func (s ImmutableCallOnChainParams) IntValue() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamIntValue))
}

type MutableCallOnChainParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableCallOnChainParams) HnameContract() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHnameContract))
}

func (s MutableCallOnChainParams) HnameEP() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHnameEP))
}

func (s MutableCallOnChainParams) IntValue() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamIntValue))
}

type ImmutableCheckContextFromFullEPParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableCheckContextFromFullEPParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutableCheckContextFromFullEPParams) Caller() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamCaller))
}

func (s ImmutableCheckContextFromFullEPParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutableCheckContextFromFullEPParams) ChainOwnerID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamChainOwnerID))
}

func (s ImmutableCheckContextFromFullEPParams) ContractCreator() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamContractCreator))
}

type MutableCheckContextFromFullEPParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableCheckContextFromFullEPParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutableCheckContextFromFullEPParams) Caller() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamCaller))
}

func (s MutableCheckContextFromFullEPParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutableCheckContextFromFullEPParams) ChainOwnerID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamChainOwnerID))
}

func (s MutableCheckContextFromFullEPParams) ContractCreator() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamContractCreator))
}

type ImmutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableInitParams) Fail() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamFail))
}

type MutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableInitParams) Fail() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamFail))
}

type ImmutablePassTypesFullParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablePassTypesFullParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutablePassTypesFullParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutablePassTypesFullParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutablePassTypesFullParams) ContractID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamContractID))
}

func (s ImmutablePassTypesFullParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

func (s ImmutablePassTypesFullParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

func (s ImmutablePassTypesFullParams) HnameZero() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHnameZero))
}

func (s ImmutablePassTypesFullParams) Int64() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64))
}

func (s ImmutablePassTypesFullParams) Int64Zero() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64Zero))
}

func (s ImmutablePassTypesFullParams) String() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamString))
}

func (s ImmutablePassTypesFullParams) StringZero() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamStringZero))
}

type MutablePassTypesFullParams struct {
	proxy wasmtypes.Proxy
}

func (s MutablePassTypesFullParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutablePassTypesFullParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutablePassTypesFullParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutablePassTypesFullParams) ContractID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamContractID))
}

func (s MutablePassTypesFullParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

func (s MutablePassTypesFullParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}

func (s MutablePassTypesFullParams) HnameZero() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHnameZero))
}

func (s MutablePassTypesFullParams) Int64() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64))
}

func (s MutablePassTypesFullParams) Int64Zero() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64Zero))
}

func (s MutablePassTypesFullParams) String() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamString))
}

func (s MutablePassTypesFullParams) StringZero() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamStringZero))
}

type ImmutableRunRecursionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableRunRecursionParams) IntValue() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamIntValue))
}

type MutableRunRecursionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableRunRecursionParams) IntValue() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamIntValue))
}

type ImmutableSetIntParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSetIntParams) IntValue() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamIntValue))
}

func (s ImmutableSetIntParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableSetIntParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSetIntParams) IntValue() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamIntValue))
}

func (s MutableSetIntParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableSpawnParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSpawnParams) ProgHash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamProgHash))
}

type MutableSpawnParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSpawnParams) ProgHash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamProgHash))
}

type ImmutableTestEventLogGenericDataParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTestEventLogGenericDataParams) Counter() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamCounter))
}

type MutableTestEventLogGenericDataParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTestEventLogGenericDataParams) Counter() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamCounter))
}

type ImmutableWithdrawFromChainParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableWithdrawFromChainParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutableWithdrawFromChainParams) GasBudget() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamGasBudget))
}

func (s ImmutableWithdrawFromChainParams) IotasWithdrawal() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamIotasWithdrawal))
}

type MutableWithdrawFromChainParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableWithdrawFromChainParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutableWithdrawFromChainParams) GasBudget() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamGasBudget))
}

func (s MutableWithdrawFromChainParams) IotasWithdrawal() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamIotasWithdrawal))
}

type ImmutableCheckContextFromViewEPParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableCheckContextFromViewEPParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutableCheckContextFromViewEPParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutableCheckContextFromViewEPParams) ChainOwnerID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamChainOwnerID))
}

func (s ImmutableCheckContextFromViewEPParams) ContractCreator() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamContractCreator))
}

type MutableCheckContextFromViewEPParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableCheckContextFromViewEPParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutableCheckContextFromViewEPParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutableCheckContextFromViewEPParams) ChainOwnerID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamChainOwnerID))
}

func (s MutableCheckContextFromViewEPParams) ContractCreator() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamContractCreator))
}

type ImmutableFibonacciParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableFibonacciParams) IntValue() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamIntValue))
}

type MutableFibonacciParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableFibonacciParams) IntValue() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamIntValue))
}

type ImmutableGetIntParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetIntParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableGetIntParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetIntParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableGetStringValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetStringValueParams) VarName() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamVarName))
}

type MutableGetStringValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetStringValueParams) VarName() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamVarName))
}

type ImmutablePassTypesViewParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablePassTypesViewParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutablePassTypesViewParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutablePassTypesViewParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutablePassTypesViewParams) ContractID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamContractID))
}

func (s ImmutablePassTypesViewParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

func (s ImmutablePassTypesViewParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

func (s ImmutablePassTypesViewParams) HnameZero() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHnameZero))
}

func (s ImmutablePassTypesViewParams) Int64() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64))
}

func (s ImmutablePassTypesViewParams) Int64Zero() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64Zero))
}

func (s ImmutablePassTypesViewParams) String() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamString))
}

func (s ImmutablePassTypesViewParams) StringZero() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamStringZero))
}

type MutablePassTypesViewParams struct {
	proxy wasmtypes.Proxy
}

func (s MutablePassTypesViewParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutablePassTypesViewParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutablePassTypesViewParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutablePassTypesViewParams) ContractID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamContractID))
}

func (s MutablePassTypesViewParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

func (s MutablePassTypesViewParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}

func (s MutablePassTypesViewParams) HnameZero() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHnameZero))
}

func (s MutablePassTypesViewParams) Int64() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64))
}

func (s MutablePassTypesViewParams) Int64Zero() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64Zero))
}

func (s MutablePassTypesViewParams) String() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamString))
}

func (s MutablePassTypesViewParams) StringZero() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamStringZero))
}
