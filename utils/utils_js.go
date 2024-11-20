//go:build js

package utils

import "syscall/js"

var JsGlobal = js.Global()

func CreateElementNS(name string) js.Value {
	return JsGlobal.Get("document").Call("createElementNS", "http://www.w3.org/1999/xhtml", name)
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
	class := JsGlobal.Get(className)
	if !class.Truthy() {
		return false
	}
	return value.InstanceOf(class)
}

func JsNull() js.Value {
	return js.Null()
}

func JsUndefined() js.Value {
	return js.Undefined()
}

type JsValue js.Value
type JsValueSlice []js.Value

func FuncOf(fn func(this JsValue, args JsValueSlice) any) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return fn(JsValue(this), JsValueSlice(args))
	})
}

func (j JsValue) AsJsValue() js.Value {
	return js.Value(j)
}
