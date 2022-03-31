// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package schemacomment

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableStringMapOfStringArrayAppendResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArrayAppendResults) Length() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultLength))
}

type MutableStringMapOfStringArrayAppendResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArrayAppendResults) Length() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultLength))
}

type ImmutableStringMapOfStringArrayLengthResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableStringMapOfStringArrayLengthResults) Length() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ResultLength))
}

type MutableStringMapOfStringArrayLengthResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableStringMapOfStringArrayLengthResults) Length() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ResultLength))
}
