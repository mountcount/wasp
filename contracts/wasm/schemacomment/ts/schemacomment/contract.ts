// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class StringMapOfStringArrayAppendCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncStringMapOfStringArrayAppend);
	params: sc.MutableStringMapOfStringArrayAppendParams = new sc.MutableStringMapOfStringArrayAppendParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableStringMapOfStringArrayAppendResults = new sc.ImmutableStringMapOfStringArrayAppendResults(wasmlib.ScView.nilProxy);
}

export class StringMapOfStringArrayAppendContext {
	events: sc.SchemaCommentEvents = new sc.SchemaCommentEvents();
	params: sc.ImmutableStringMapOfStringArrayAppendParams = new sc.ImmutableStringMapOfStringArrayAppendParams(wasmlib.paramsProxy());
	results: sc.MutableStringMapOfStringArrayAppendResults = new sc.MutableStringMapOfStringArrayAppendResults(wasmlib.ScView.nilProxy);
	state: sc.MutableSchemaCommentState = new sc.MutableSchemaCommentState(wasmlib.ScState.proxy());
}

export class StringMapOfStringArrayLengthCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewStringMapOfStringArrayLength);
	params: sc.MutableStringMapOfStringArrayLengthParams = new sc.MutableStringMapOfStringArrayLengthParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableStringMapOfStringArrayLengthResults = new sc.ImmutableStringMapOfStringArrayLengthResults(wasmlib.ScView.nilProxy);
}

export class StringMapOfStringArrayLengthContext {
	params: sc.ImmutableStringMapOfStringArrayLengthParams = new sc.ImmutableStringMapOfStringArrayLengthParams(wasmlib.paramsProxy());
	results: sc.MutableStringMapOfStringArrayLengthResults = new sc.MutableStringMapOfStringArrayLengthResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableSchemaCommentState = new sc.ImmutableSchemaCommentState(wasmlib.ScState.proxy());
}

export class ScFuncs {
	static stringMapOfStringArrayAppend(_ctx: wasmlib.ScFuncCallContext): StringMapOfStringArrayAppendCall {
		const f = new StringMapOfStringArrayAppendCall();
		f.params = new sc.MutableStringMapOfStringArrayAppendParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableStringMapOfStringArrayAppendResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static stringMapOfStringArrayLength(_ctx: wasmlib.ScViewCallContext): StringMapOfStringArrayLengthCall {
		const f = new StringMapOfStringArrayLengthCall();
		f.params = new sc.MutableStringMapOfStringArrayLengthParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableStringMapOfStringArrayLengthResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}
}
