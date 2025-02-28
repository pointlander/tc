// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"syscall/js"
)

func main() {
	js.Global().Set("draw", drawWrapper())
	select {}
}

func Draw(tree string) string {
	document := js.Global().Get("document")
	canvas := document.Call("getElementById", "myCanvas")
	context := canvas.Call("getContext", "2d")

	context.Call("beginPath")
	context.Call("arc", 256, 256, 256, 0, 2*math.Pi, true)
	context.Call("stroke")
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
