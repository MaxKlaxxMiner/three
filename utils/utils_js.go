//go:build js

package utils

import "syscall/js"

func CreateElementNS(name string) js.Value {
	return js.Global().Get("document").Call("createElementNS", "http://www.w3.org/1999/xhtml", name)
}

func CreateCanvasElement() js.Value {
	canvas := CreateElementNS("canvas")
	canvas.Get("style").Set("display", "block")
	return canvas
}
