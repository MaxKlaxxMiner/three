//go:build js

package utils

import (
	"reflect"
	"syscall/js"
	"unsafe"
)

func CreateElementNS(name string) js.Value {
	return js.Global().Get("document").Call("createElementNS", "http://www.w3.org/1999/xhtml", name)
}

func CreateCanvasElement() js.Value {
	canvas := CreateElementNS("canvas")
	canvas.Get("style").Set("display", "block")
	return canvas
}

// --- JS Helper ---

var JsGlobal = js.Global()

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

var typeInt8Array = js.Global().Get("Int8Array")
var typeUint8Array = js.Global().Get("Uint8Array")
var typeUint8ClampedArray = js.Global().Get("Uint8ClampedArray")
var typeInt16Array = js.Global().Get("Int16Array")
var typeUint16Array = js.Global().Get("Uint16Array")
var typeInt32Array = js.Global().Get("Int32Array")
var typeUint32Array = js.Global().Get("Uint32Array")
var typeFloat32Array = js.Global().Get("Float32Array")
var typeFloat64Array = js.Global().Get("Float64Array")
var typeBigInt64Array = js.Global().Get("BigInt64Array")
var typeBigUint64Array = js.Global().Get("BigUint64Array")

func ConvertTypedArrayToFloat64Slice(value js.Value) []float64 {
	if !value.Truthy() {
		return nil
	}

	bytes := value.Get("byteLength").Int()
	tmp := make([]byte, bytes)
	if js.CopyBytesToGo(tmp, typeUint8Array.New(value.Get("buffer"))) != bytes {
		return nil
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&tmp))

	switch {
	case value.InstanceOf(typeFloat64Array):
		header.Len, header.Cap = bytes/8, bytes/8
		return *(*[]float64)(unsafe.Pointer(header))
	case value.InstanceOf(typeFloat32Array):
		header.Len, header.Cap = bytes/4, bytes/4
		array := *(*[]float32)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeInt8Array):
		array := *(*[]int8)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeUint8Array), value.InstanceOf(typeUint8ClampedArray):
		array := *(*[]uint8)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeInt16Array):
		header.Len, header.Cap = bytes/2, bytes/2
		array := *(*[]int16)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeUint16Array):
		header.Len, header.Cap = bytes/2, bytes/2
		array := *(*[]uint16)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeInt32Array):
		header.Len, header.Cap = bytes/4, bytes/4
		array := *(*[]int32)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeUint32Array):
		header.Len, header.Cap = bytes/4, bytes/4
		array := *(*[]uint32)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeBigInt64Array):
		header.Len, header.Cap = bytes/8, bytes/8
		array := *(*[]int64)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	case value.InstanceOf(typeBigUint64Array):
		header.Len, header.Cap = bytes/8, bytes/8
		array := *(*[]uint64)(unsafe.Pointer(header))
		output := make([]float64, len(array))
		for i := range array {
			output[i] = float64(array[i])
		}
		return output
	default:
		js.Global().Get("console").Call("log", "unknown typed array:", value.Type().String())
	}

	return nil
}
