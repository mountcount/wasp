// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package schemacomment

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "schemacomment"
	ScDescription = "demonstration of dumping comments in yaml file to schema definition"
	HScName       = wasmtypes.ScHname(0xc8a14753)
)

const (
	ParamArr   = "arr"
	ParamName  = "name"
	ParamValue = "value"
)

const (
	ResultLength = "length"
)

const (
	StateTestState1 = "TestState1"
	StateTestState2 = "TestState2"
)

const (
	FuncTestFunc1 = "testFunc1"
	ViewTestView1 = "testView1"
)

const (
	HFuncTestFunc1 = wasmtypes.ScHname(0xf73370b1)
	HViewTestView1 = wasmtypes.ScHname(0xf0f3e6e8)
)
