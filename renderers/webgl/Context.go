package webgl

import (
	"fmt"
	"github.com/MaxKlaxxMiner/three/utils"
)

func (gl *Context) GetExtension(name string) Extension {
	return Extension{gl.Call("getExtension", name), map[string]int32{}}
}

func (gl *Context) GetParameterInt(id int32) int {
	return gl.Call("getParameter", id).Int()
}

func (gl *Context) GetParameterStr(id int32) string {
	return gl.Call("getParameter", id).String()
}

func (gl *Context) GetParameterFloat64Array(id int32) []float64 {
	return utils.ConvertTypedArrayToFloat64Slice(gl.Call("getParameter", id))
}

func (gl *Context) Const(name string) int32 {
	v := gl.Get(name)
	var r int32
	if !v.IsUndefined() {
		r = int32(v.Int())
	} else {
		fmt.Println("warn const gl:", name)
	}
	return r
}

func (e *Extension) Const(name string) int32 {
	if r, ok := e.Consts[name]; ok {
		return r
	}
	v := e.Get(name)
	var r int32
	if v.Truthy() {
		r = int32(v.Int())
	} else {
		fmt.Println("warn const extension:", name)
	}
	e.Consts[name] = r
	return r
}
