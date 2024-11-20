//go:build !js

package utils

import (
	"github.com/MaxKlaxxMiner/three/internal/fake/js"
)

var JsGlobal = js.Global()

func CreateElementNS(_ string) js.Value {
	return js.Value{}
}

func CreateCanvasElement() js.Value {
	return js.Value{}
}

func InstanceOf(_ *js.Value, _ string) bool {
	return false
}

func JsNull() js.Value {
	return js.Null()
}

func JsUndefined() js.Value {
	return js.Undefined()
}

type JsValue js.Value
type JsValueSlice []js.Value

func FuncOf(_ func(this JsValue, _ JsValueSlice) any) js.Func {
	return js.Func{}
}

func (j JsValue) AsJsValue() js.Value {
	return js.Value(j)
}
