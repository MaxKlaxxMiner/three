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

func InstanceOf(value *js.Value, className string) bool {
	if value == nil || !value.Truthy() {
		return false
	}
	class := js.Global().Get(className)
	if !class.Truthy() {
		return false
	}
	return value.InstanceOf(class)
}

func JsNull() js.Value {
	return js.Null()
}
