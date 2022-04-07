// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

const exportMap: wasmlib.ScExportMap = {
	names: [
		sc.FuncCallOnChain,
		sc.FuncCheckContextFromFullEP,
		sc.FuncClaimAllowance,
		sc.FuncDoNothing,
		sc.FuncEstimateMinDust,
		sc.FuncIncCounter,
		sc.FuncInfiniteLoop,
		sc.FuncInit,
		sc.FuncPassTypesFull,
		sc.FuncPingAllowanceBack,
		sc.FuncRunRecursion,
		sc.FuncSendLargeRequest,
		sc.FuncSendNFTsBack,
		sc.FuncSendToAddress,
		sc.FuncSetInt,
		sc.FuncSpawn,
		sc.FuncSplitFunds,
		sc.FuncSplitFundsNativeTokens,
		sc.FuncTestBlockContext1,
		sc.FuncTestBlockContext2,
		sc.FuncTestCallPanicFullEP,
		sc.FuncTestCallPanicViewEPFromFull,
		sc.FuncTestChainOwnerIDFull,
		sc.FuncTestEventLogDeploy,
		sc.FuncTestEventLogEventData,
		sc.FuncTestEventLogGenericData,
		sc.FuncTestPanicFullEP,
		sc.FuncWithdrawFromChain,
		sc.ViewCheckContextFromViewEP,
		sc.ViewFibonacci,
		sc.ViewGetCounter,
		sc.ViewGetInt,
		sc.ViewGetStringValue,
		sc.ViewInfiniteLoopView,
		sc.ViewJustView,
		sc.ViewPassTypesView,
		sc.ViewTestCallPanicViewEPFromView,
		sc.ViewTestChainOwnerIDView,
		sc.ViewTestPanicViewEP,
		sc.ViewTestSandboxCall,
	],
	funcs: [
		funcCallOnChainThunk,
		funcCheckContextFromFullEPThunk,
		funcClaimAllowanceThunk,
		funcDoNothingThunk,
		funcEstimateMinDustThunk,
		funcIncCounterThunk,
		funcInfiniteLoopThunk,
		funcInitThunk,
		funcPassTypesFullThunk,
		funcPingAllowanceBackThunk,
		funcRunRecursionThunk,
		funcSendLargeRequestThunk,
		funcSendNFTsBackThunk,
		funcSendToAddressThunk,
		funcSetIntThunk,
		funcSpawnThunk,
		funcSplitFundsThunk,
		funcSplitFundsNativeTokensThunk,
		funcTestBlockContext1Thunk,
		funcTestBlockContext2Thunk,
		funcTestCallPanicFullEPThunk,
		funcTestCallPanicViewEPFromFullThunk,
		funcTestChainOwnerIDFullThunk,
		funcTestEventLogDeployThunk,
		funcTestEventLogEventDataThunk,
		funcTestEventLogGenericDataThunk,
		funcTestPanicFullEPThunk,
		funcWithdrawFromChainThunk,
	],
	views: [
		viewCheckContextFromViewEPThunk,
		viewFibonacciThunk,
		viewGetCounterThunk,
		viewGetIntThunk,
		viewGetStringValueThunk,
		viewInfiniteLoopViewThunk,
		viewJustViewThunk,
		viewPassTypesViewThunk,
		viewTestCallPanicViewEPFromViewThunk,
		viewTestChainOwnerIDViewThunk,
		viewTestPanicViewEPThunk,
		viewTestSandboxCallThunk,
	],
};

export function on_call(index: i32): void {
	wasmlib.ScExports.call(index, exportMap);
}

export function on_load(): void {
	wasmlib.ScExports.export(exportMap);
}

function funcCallOnChainThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcCallOnChain");
	let f = new sc.CallOnChainContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableCallOnChainResults(results.asProxy());
	ctx.require(f.params.intValue().exists(), "missing mandatory intValue");
	sc.funcCallOnChain(ctx, f);
	ctx.results(results);
	ctx.log("testcore.funcCallOnChain ok");
}

function funcCheckContextFromFullEPThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcCheckContextFromFullEP");
	let f = new sc.CheckContextFromFullEPContext();
	ctx.require(f.params.agentID().exists(), "missing mandatory agentID");
	ctx.require(f.params.caller().exists(), "missing mandatory caller");
	ctx.require(f.params.chainID().exists(), "missing mandatory chainID");
	ctx.require(f.params.chainOwnerID().exists(), "missing mandatory chainOwnerID");
	ctx.require(f.params.contractCreator().exists(), "missing mandatory contractCreator");
	sc.funcCheckContextFromFullEP(ctx, f);
	ctx.log("testcore.funcCheckContextFromFullEP ok");
}

function funcClaimAllowanceThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcClaimAllowance");
	let f = new sc.ClaimAllowanceContext();
	sc.funcClaimAllowance(ctx, f);
	ctx.log("testcore.funcClaimAllowance ok");
}

function funcDoNothingThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcDoNothing");
	let f = new sc.DoNothingContext();
	sc.funcDoNothing(ctx, f);
	ctx.log("testcore.funcDoNothing ok");
}

function funcEstimateMinDustThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcEstimateMinDust");
	let f = new sc.EstimateMinDustContext();
	sc.funcEstimateMinDust(ctx, f);
	ctx.log("testcore.funcEstimateMinDust ok");
}

function funcIncCounterThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcIncCounter");
	let f = new sc.IncCounterContext();
	sc.funcIncCounter(ctx, f);
	ctx.log("testcore.funcIncCounter ok");
}

function funcInfiniteLoopThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcInfiniteLoop");
	let f = new sc.InfiniteLoopContext();
	sc.funcInfiniteLoop(ctx, f);
	ctx.log("testcore.funcInfiniteLoop ok");
}

function funcInitThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcInit");
	let f = new sc.InitContext();
	sc.funcInit(ctx, f);
	ctx.log("testcore.funcInit ok");
}

function funcPassTypesFullThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcPassTypesFull");
	let f = new sc.PassTypesFullContext();
	ctx.require(f.params.address().exists(), "missing mandatory address");
	ctx.require(f.params.agentID().exists(), "missing mandatory agentID");
	ctx.require(f.params.chainID().exists(), "missing mandatory chainID");
	ctx.require(f.params.contractID().exists(), "missing mandatory contractID");
	ctx.require(f.params.hash().exists(), "missing mandatory hash");
	ctx.require(f.params.hname().exists(), "missing mandatory hname");
	ctx.require(f.params.hnameZero().exists(), "missing mandatory hnameZero");
	ctx.require(f.params.int64().exists(), "missing mandatory int64");
	ctx.require(f.params.int64Zero().exists(), "missing mandatory int64Zero");
	ctx.require(f.params.string().exists(), "missing mandatory string");
	ctx.require(f.params.stringZero().exists(), "missing mandatory stringZero");
	sc.funcPassTypesFull(ctx, f);
	ctx.log("testcore.funcPassTypesFull ok");
}

function funcPingAllowanceBackThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcPingAllowanceBack");
	let f = new sc.PingAllowanceBackContext();
	sc.funcPingAllowanceBack(ctx, f);
	ctx.log("testcore.funcPingAllowanceBack ok");
}

function funcRunRecursionThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcRunRecursion");
	let f = new sc.RunRecursionContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableRunRecursionResults(results.asProxy());
	ctx.require(f.params.intValue().exists(), "missing mandatory intValue");
	sc.funcRunRecursion(ctx, f);
	ctx.results(results);
	ctx.log("testcore.funcRunRecursion ok");
}

function funcSendLargeRequestThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSendLargeRequest");
	let f = new sc.SendLargeRequestContext();
	sc.funcSendLargeRequest(ctx, f);
	ctx.log("testcore.funcSendLargeRequest ok");
}

function funcSendNFTsBackThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSendNFTsBack");
	let f = new sc.SendNFTsBackContext();
	sc.funcSendNFTsBack(ctx, f);
	ctx.log("testcore.funcSendNFTsBack ok");
}

function funcSendToAddressThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSendToAddress");
	let f = new sc.SendToAddressContext();
	sc.funcSendToAddress(ctx, f);
	ctx.log("testcore.funcSendToAddress ok");
}

function funcSetIntThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSetInt");
	let f = new sc.SetIntContext();
	ctx.require(f.params.intValue().exists(), "missing mandatory intValue");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	sc.funcSetInt(ctx, f);
	ctx.log("testcore.funcSetInt ok");
}

function funcSpawnThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSpawn");
	let f = new sc.SpawnContext();
	ctx.require(f.params.progHash().exists(), "missing mandatory progHash");
	sc.funcSpawn(ctx, f);
	ctx.log("testcore.funcSpawn ok");
}

function funcSplitFundsThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSplitFunds");
	let f = new sc.SplitFundsContext();
	sc.funcSplitFunds(ctx, f);
	ctx.log("testcore.funcSplitFunds ok");
}

function funcSplitFundsNativeTokensThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcSplitFundsNativeTokens");
	let f = new sc.SplitFundsNativeTokensContext();
	sc.funcSplitFundsNativeTokens(ctx, f);
	ctx.log("testcore.funcSplitFundsNativeTokens ok");
}

function funcTestBlockContext1Thunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestBlockContext1");
	let f = new sc.TestBlockContext1Context();
	sc.funcTestBlockContext1(ctx, f);
	ctx.log("testcore.funcTestBlockContext1 ok");
}

function funcTestBlockContext2Thunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestBlockContext2");
	let f = new sc.TestBlockContext2Context();
	sc.funcTestBlockContext2(ctx, f);
	ctx.log("testcore.funcTestBlockContext2 ok");
}

function funcTestCallPanicFullEPThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestCallPanicFullEP");
	let f = new sc.TestCallPanicFullEPContext();
	sc.funcTestCallPanicFullEP(ctx, f);
	ctx.log("testcore.funcTestCallPanicFullEP ok");
}

function funcTestCallPanicViewEPFromFullThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestCallPanicViewEPFromFull");
	let f = new sc.TestCallPanicViewEPFromFullContext();
	sc.funcTestCallPanicViewEPFromFull(ctx, f);
	ctx.log("testcore.funcTestCallPanicViewEPFromFull ok");
}

function funcTestChainOwnerIDFullThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestChainOwnerIDFull");
	let f = new sc.TestChainOwnerIDFullContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableTestChainOwnerIDFullResults(results.asProxy());
	sc.funcTestChainOwnerIDFull(ctx, f);
	ctx.results(results);
	ctx.log("testcore.funcTestChainOwnerIDFull ok");
}

function funcTestEventLogDeployThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestEventLogDeploy");
	let f = new sc.TestEventLogDeployContext();
	sc.funcTestEventLogDeploy(ctx, f);
	ctx.log("testcore.funcTestEventLogDeploy ok");
}

function funcTestEventLogEventDataThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestEventLogEventData");
	let f = new sc.TestEventLogEventDataContext();
	sc.funcTestEventLogEventData(ctx, f);
	ctx.log("testcore.funcTestEventLogEventData ok");
}

function funcTestEventLogGenericDataThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestEventLogGenericData");
	let f = new sc.TestEventLogGenericDataContext();
	ctx.require(f.params.counter().exists(), "missing mandatory counter");
	sc.funcTestEventLogGenericData(ctx, f);
	ctx.log("testcore.funcTestEventLogGenericData ok");
}

function funcTestPanicFullEPThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcTestPanicFullEP");
	let f = new sc.TestPanicFullEPContext();
	sc.funcTestPanicFullEP(ctx, f);
	ctx.log("testcore.funcTestPanicFullEP ok");
}

function funcWithdrawFromChainThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("testcore.funcWithdrawFromChain");
	let f = new sc.WithdrawFromChainContext();
	ctx.require(f.params.chainID().exists(), "missing mandatory chainID");
	ctx.require(f.params.gasBudget().exists(), "missing mandatory gasBudget");
	ctx.require(f.params.iotasWithdrawal().exists(), "missing mandatory iotasWithdrawal");
	sc.funcWithdrawFromChain(ctx, f);
	ctx.log("testcore.funcWithdrawFromChain ok");
}

function viewCheckContextFromViewEPThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewCheckContextFromViewEP");
	let f = new sc.CheckContextFromViewEPContext();
	ctx.require(f.params.agentID().exists(), "missing mandatory agentID");
	ctx.require(f.params.chainID().exists(), "missing mandatory chainID");
	ctx.require(f.params.chainOwnerID().exists(), "missing mandatory chainOwnerID");
	ctx.require(f.params.contractCreator().exists(), "missing mandatory contractCreator");
	sc.viewCheckContextFromViewEP(ctx, f);
	ctx.log("testcore.viewCheckContextFromViewEP ok");
}

function viewFibonacciThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewFibonacci");
	let f = new sc.FibonacciContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableFibonacciResults(results.asProxy());
	ctx.require(f.params.intValue().exists(), "missing mandatory intValue");
	sc.viewFibonacci(ctx, f);
	ctx.results(results);
	ctx.log("testcore.viewFibonacci ok");
}

function viewGetCounterThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewGetCounter");
	let f = new sc.GetCounterContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetCounterResults(results.asProxy());
	sc.viewGetCounter(ctx, f);
	ctx.results(results);
	ctx.log("testcore.viewGetCounter ok");
}

function viewGetIntThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewGetInt");
	let f = new sc.GetIntContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetIntResults(results.asProxy());
	ctx.require(f.params.name().exists(), "missing mandatory name");
	sc.viewGetInt(ctx, f);
	ctx.results(results);
	ctx.log("testcore.viewGetInt ok");
}

function viewGetStringValueThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewGetStringValue");
	let f = new sc.GetStringValueContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetStringValueResults(results.asProxy());
	ctx.require(f.params.varName().exists(), "missing mandatory varName");
	sc.viewGetStringValue(ctx, f);
	ctx.results(results);
	ctx.log("testcore.viewGetStringValue ok");
}

function viewInfiniteLoopViewThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewInfiniteLoopView");
	let f = new sc.InfiniteLoopViewContext();
	sc.viewInfiniteLoopView(ctx, f);
	ctx.log("testcore.viewInfiniteLoopView ok");
}

function viewJustViewThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewJustView");
	let f = new sc.JustViewContext();
	sc.viewJustView(ctx, f);
	ctx.log("testcore.viewJustView ok");
}

function viewPassTypesViewThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewPassTypesView");
	let f = new sc.PassTypesViewContext();
	ctx.require(f.params.address().exists(), "missing mandatory address");
	ctx.require(f.params.agentID().exists(), "missing mandatory agentID");
	ctx.require(f.params.chainID().exists(), "missing mandatory chainID");
	ctx.require(f.params.contractID().exists(), "missing mandatory contractID");
	ctx.require(f.params.hash().exists(), "missing mandatory hash");
	ctx.require(f.params.hname().exists(), "missing mandatory hname");
	ctx.require(f.params.hnameZero().exists(), "missing mandatory hnameZero");
	ctx.require(f.params.int64().exists(), "missing mandatory int64");
	ctx.require(f.params.int64Zero().exists(), "missing mandatory int64Zero");
	ctx.require(f.params.string().exists(), "missing mandatory string");
	ctx.require(f.params.stringZero().exists(), "missing mandatory stringZero");
	sc.viewPassTypesView(ctx, f);
	ctx.log("testcore.viewPassTypesView ok");
}

function viewTestCallPanicViewEPFromViewThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewTestCallPanicViewEPFromView");
	let f = new sc.TestCallPanicViewEPFromViewContext();
	sc.viewTestCallPanicViewEPFromView(ctx, f);
	ctx.log("testcore.viewTestCallPanicViewEPFromView ok");
}

function viewTestChainOwnerIDViewThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewTestChainOwnerIDView");
	let f = new sc.TestChainOwnerIDViewContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableTestChainOwnerIDViewResults(results.asProxy());
	sc.viewTestChainOwnerIDView(ctx, f);
	ctx.results(results);
	ctx.log("testcore.viewTestChainOwnerIDView ok");
}

function viewTestPanicViewEPThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewTestPanicViewEP");
	let f = new sc.TestPanicViewEPContext();
	sc.viewTestPanicViewEP(ctx, f);
	ctx.log("testcore.viewTestPanicViewEP ok");
}

function viewTestSandboxCallThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("testcore.viewTestSandboxCall");
	let f = new sc.TestSandboxCallContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableTestSandboxCallResults(results.asProxy());
	sc.viewTestSandboxCall(ctx, f);
	ctx.results(results);
	ctx.log("testcore.viewTestSandboxCall ok");
}
