package webgl

import "fmt"

func (gl *Context) getExtension(name string) GLExtension {
	return GLExtension{gl.Call("getExtension", name), map[string]int{}}
}

func (gl *Context) getParameterInt(id int) int {
	return gl.Call("getParameter", id).Int()
}

func (gl *Context) getParameterStr(id int) string {
	return gl.Call("getParameter", id).String()
}

func (gl *Context) Const(name string) int {
	if r, ok := gl.Consts[name]; ok {
		return r
	}
	v := gl.Get(name)
	r := 0
	if v.Truthy() {
		r = v.Int()
	} else {
		fmt.Println("warn const gl:", name)
	}
	gl.Consts[name] = r
	return r
}

func (e *GLExtension) Const(name string) int {
	if r, ok := e.Consts[name]; ok {
		return r
	}
	v := e.Get(name)
	r := 0
	if v.Truthy() {
		r = v.Int()
	} else {
		fmt.Println("warn const extension:", name)
	}
	e.Consts[name] = r
	return r
}
