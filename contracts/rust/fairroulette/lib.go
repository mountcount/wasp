// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncForceReset, funcForceResetThunk)
	exports.AddFunc(FuncPayWinners, funcPayWinnersThunk)
	exports.AddFunc(FuncPlaceBet, funcPlaceBetThunk)
	exports.AddFunc(FuncPlayPeriod, funcPlayPeriodThunk)
	exports.AddView(ViewLastWinningNumber, viewLastWinningNumberThunk)
	exports.AddView(ViewRoundNumber, viewRoundNumberThunk)
	exports.AddView(ViewRoundStartedAt, viewRoundStartedAtThunk)
	exports.AddView(ViewRoundStatus, viewRoundStatusThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type ForceResetContext struct {
	State MutableFairRouletteState
}

func funcForceResetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairroulette.funcForceReset")
	// only SC creator can restart the round forcefully
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	f := &ForceResetContext{
		State: MutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcForceReset(ctx, f)
	ctx.Log("fairroulette.funcForceReset ok")
}

type PayWinnersContext struct {
	State MutableFairRouletteState
}

func funcPayWinnersThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairroulette.funcPayWinners")
	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.AccountID(), "no permission")

	f := &PayWinnersContext{
		State: MutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcPayWinners(ctx, f)
	ctx.Log("fairroulette.funcPayWinners ok")
}

type PlaceBetContext struct {
	Params ImmutablePlaceBetParams
	State  MutableFairRouletteState
}

func funcPlaceBetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairroulette.funcPlaceBet")
	f := &PlaceBetContext{
		Params: ImmutablePlaceBetParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Number().Exists(), "missing mandatory number")
	funcPlaceBet(ctx, f)
	ctx.Log("fairroulette.funcPlaceBet ok")
}

type PlayPeriodContext struct {
	Params ImmutablePlayPeriodParams
	State  MutableFairRouletteState
}

func funcPlayPeriodThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairroulette.funcPlayPeriod")
	// only SC creator can update the play period
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	f := &PlayPeriodContext{
		Params: ImmutablePlayPeriodParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.PlayPeriod().Exists(), "missing mandatory playPeriod")
	funcPlayPeriod(ctx, f)
	ctx.Log("fairroulette.funcPlayPeriod ok")
}

type LastWinningNumberContext struct {
	Results MutableLastWinningNumberResults
	State   ImmutableFairRouletteState
}

func viewLastWinningNumberThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("fairroulette.viewLastWinningNumber")
	f := &LastWinningNumberContext{
		Results: MutableLastWinningNumberResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewLastWinningNumber(ctx, f)
	ctx.Log("fairroulette.viewLastWinningNumber ok")
}

type RoundNumberContext struct {
	Results MutableRoundNumberResults
	State   ImmutableFairRouletteState
}

func viewRoundNumberThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("fairroulette.viewRoundNumber")
	f := &RoundNumberContext{
		Results: MutableRoundNumberResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewRoundNumber(ctx, f)
	ctx.Log("fairroulette.viewRoundNumber ok")
}

type RoundStartedAtContext struct {
	Results MutableRoundStartedAtResults
	State   ImmutableFairRouletteState
}

func viewRoundStartedAtThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("fairroulette.viewRoundStartedAt")
	f := &RoundStartedAtContext{
		Results: MutableRoundStartedAtResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewRoundStartedAt(ctx, f)
	ctx.Log("fairroulette.viewRoundStartedAt ok")
}

type RoundStatusContext struct {
	Results MutableRoundStatusResults
	State   ImmutableFairRouletteState
}

func viewRoundStatusThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("fairroulette.viewRoundStatus")
	f := &RoundStatusContext{
		Results: MutableRoundStatusResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableFairRouletteState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewRoundStatus(ctx, f)
	ctx.Log("fairroulette.viewRoundStatus ok")
}
