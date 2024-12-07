package webgl

import "fmt"

func (gl *Context) GetExtension(name string) Extension {
	return Extension{gl.Call("getExtension", name), map[string]ConstValue{}}
}

func (gl *Context) GetParameterInt(id ConstValue) int {
	return gl.Call("getParameter", int(id)).Int()
}

func (gl *Context) GetParameterStr(id ConstValue) string {
	return gl.Call("getParameter", int(id)).String()
}

type ConstValue int32

func (gl *Context) Const(name string) ConstValue {
	v := gl.Get(name)
	var r ConstValue
	if v.Truthy() {
		r = ConstValue(v.Int())
	} else {
		fmt.Println("warn const gl:", name)
	}
	return r
}

func (e *Extension) Const(name string) ConstValue {
	if r, ok := e.Consts[name]; ok {
		return r
	}
	v := e.Get(name)
	var r ConstValue
	if v.Truthy() {
		r = ConstValue(v.Int())
	} else {
		fmt.Println("warn const extension:", name)
	}
	e.Consts[name] = r
	return r
}
