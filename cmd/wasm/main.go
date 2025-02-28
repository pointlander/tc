// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

package main

import (
	"math"
	"syscall/js"

	. "github.com/pointlander/tc/lib"
)

func main() {
	js.Global().Set("draw", drawWrapper())
	select {}
}

func Draw(tree string) string {
	_, t := Parse([]byte(tree))
	n := t.Label()
	polygon := t.Triangulation(n)
	segment := 2 * math.Pi / float64(len(polygon))
	offset := -math.Pi/2 + segment/2

	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "myCanvas")
	context := canvas.Call("getContext", "2d")

	context.Call("beginPath")
	context.Call("arc", 256, 256, 248, 0, 2*math.Pi, true)
	context.Call("stroke")

	for i := range polygon {
		for _, j := range polygon[i] {
			context.Call("beginPath")
			context.Call("moveTo", -248*math.Cos(offset+float64(i)*segment)+256, 248*math.Sin(offset+float64(i)*segment)+256)
			context.Call("lineTo", -248*math.Cos(offset+float64(j)*segment)+256, 248*math.Sin(offset+float64(j)*segment)+256)
			context.Call("stroke")
		}
	}
	for i := range polygon {
		context.Call("beginPath")
		context.Call("arc", -248*math.Cos(offset+float64(i)*segment)+256, 248*math.Sin(offset+float64(i)*segment)+256,
			8, 0, 2*math.Pi, true)
		context.Set("fillStyle", "blue")
		context.Call("fill")
	}
	return tree
}

func drawWrapper() js.Func {
	drawFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		return Draw(args[0].String())
	})
	return drawFunc
}
