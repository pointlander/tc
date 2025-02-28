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

func Draw(id, tree string) string {
	_, t := Parse([]byte(tree))
	n := t.Label()
	polygon := t.Triangulation(n)
	segment := 2 * math.Pi / float64(len(polygon))
	offset := -math.Pi/2 + segment/2

	document := js.Global().Get("document")
	canvas := document.Call("getElementById", id)
	width := float64(canvas.Get("width").Int())
	height := float64(canvas.Get("height").Int())
	radius := width / 2
	if height/2 < radius {
		radius = height / 2
	}
	context := canvas.Call("getContext", "2d")

	context.Call("beginPath")
	context.Call("arc", width/2, height/2, radius-8, 0, 2*math.Pi, true)
	context.Call("stroke")

	for i := range polygon {
		for _, j := range polygon[i] {
			context.Call("beginPath")
			context.Call("moveTo",
				-(radius-8)*math.Cos(offset+float64(i)*segment)+width/2,
				(radius-8)*math.Sin(offset+float64(i)*segment)+height/2)
			context.Call("lineTo",
				-(radius-8)*math.Cos(offset+float64(j)*segment)+width/2,
				(radius-8)*math.Sin(offset+float64(j)*segment)+height/2)
			context.Call("stroke")
		}
	}
	for i := range polygon {
		context.Call("beginPath")
		context.Call("arc",
			-(radius-8)*math.Cos(offset+float64(i)*segment)+width/2,
			(radius-8)*math.Sin(offset+float64(i)*segment)+height/2,
			8, 0, 2*math.Pi, true)
		context.Set("fillStyle", "blue")
		context.Call("fill")
	}
	return tree
}

func drawWrapper() js.Func {
	drawFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
		return Draw(args[0].String(), args[1].String())
	})
	return drawFunc
}
