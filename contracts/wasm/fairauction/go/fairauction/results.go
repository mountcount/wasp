// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableGetInfoResults struct {
	proxy wasmtypes.Proxy
}

// nr of bidders
// token of tokens for sale
// issuer of start_auction transaction
// deposit by auction owner to cover the SC fees
// auction description
// auction duration in minutes
// the current highest bid amount
// the current highest bidder
// minimum bid amount
// number of tokens for sale
// auction owner's margin in promilles
// timestamp when auction started
func (s ImmutableGetInfoResults) Bidders() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultBidders))
}

func (s ImmutableGetInfoResults) Creator() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ResultCreator))
}

func (s ImmutableGetInfoResults) Deposit() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultDeposit))
}

func (s ImmutableGetInfoResults) Description() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ResultDescription))
}

func (s ImmutableGetInfoResults) Duration() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultDuration))
}

func (s ImmutableGetInfoResults) HighestBid() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultHighestBid))
}

func (s ImmutableGetInfoResults) HighestBidder() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ResultHighestBidder))
}

func (s ImmutableGetInfoResults) MinimumBid() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultMinimumBid))
}

func (s ImmutableGetInfoResults) NumTokens() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultNumTokens))
}

func (s ImmutableGetInfoResults) OwnerMargin() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultOwnerMargin))
}

func (s ImmutableGetInfoResults) Token() wasmtypes.ScImmutableTokenID {
	return wasmtypes.NewScImmutableTokenID(s.proxy.Root(ResultToken))
}

func (s ImmutableGetInfoResults) WhenStarted() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultWhenStarted))
}

type MutableGetInfoResults struct {
	proxy wasmtypes.Proxy
}

// nr of bidders
// token of tokens for sale
// issuer of start_auction transaction
// deposit by auction owner to cover the SC fees
// auction description
// auction duration in minutes
// the current highest bid amount
// the current highest bidder
// minimum bid amount
// number of tokens for sale
// auction owner's margin in promilles
// timestamp when auction started
func (s MutableGetInfoResults) Bidders() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultBidders))
}

func (s MutableGetInfoResults) Creator() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ResultCreator))
}

func (s MutableGetInfoResults) Deposit() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultDeposit))
}

func (s MutableGetInfoResults) Description() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ResultDescription))
}

func (s MutableGetInfoResults) Duration() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultDuration))
}

func (s MutableGetInfoResults) HighestBid() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultHighestBid))
}

func (s MutableGetInfoResults) HighestBidder() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ResultHighestBidder))
}

func (s MutableGetInfoResults) MinimumBid() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultMinimumBid))
}

func (s MutableGetInfoResults) NumTokens() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultNumTokens))
}

func (s MutableGetInfoResults) OwnerMargin() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultOwnerMargin))
}

func (s MutableGetInfoResults) Token() wasmtypes.ScMutableTokenID {
	return wasmtypes.NewScMutableTokenID(s.proxy.Root(ResultToken))
}

func (s MutableGetInfoResults) WhenStarted() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultWhenStarted))
}
