// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblob

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type MapStringToImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableBytes) GetBytes(key string) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableStoreBlobParams struct {
	proxy wasmtypes.Proxy
}

// set of named blobs
func (s ImmutableStoreBlobParams) Blobs() MapStringToImmutableBytes {
	//nolint:gosimple
	return MapStringToImmutableBytes{proxy: s.proxy}
}

type MapStringToMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableBytes) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableBytes) GetBytes(key string) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableStoreBlobParams struct {
	proxy wasmtypes.Proxy
}

// set of named blobs
func (s MutableStoreBlobParams) Blobs() MapStringToMutableBytes {
	//nolint:gosimple
	return MapStringToMutableBytes{proxy: s.proxy}
}

type ImmutableGetBlobFieldParams struct {
	proxy wasmtypes.Proxy
}

// blob name
// blob set
func (s ImmutableGetBlobFieldParams) Field() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamField))
}

func (s ImmutableGetBlobFieldParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

type MutableGetBlobFieldParams struct {
	proxy wasmtypes.Proxy
}

// blob name
// blob set
func (s MutableGetBlobFieldParams) Field() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamField))
}

func (s MutableGetBlobFieldParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

type ImmutableGetBlobInfoParams struct {
	proxy wasmtypes.Proxy
}

// blob set
func (s ImmutableGetBlobInfoParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

type MutableGetBlobInfoParams struct {
	proxy wasmtypes.Proxy
}

// blob set
func (s MutableGetBlobInfoParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}
