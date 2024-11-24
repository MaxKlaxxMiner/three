package three

func (gl *GLContext) getExtension(name string) GLExtension {
	return GLExtension{gl.Call("getExtension", name), map[string]int{}}
}

func (gl *GLContext) getParameterInt(id int) int {
	return gl.Call("getParameter", id).Int()
}

func (gl *GLContext) Const(name string) int {
	if r, ok := gl.consts[name]; ok {
		return r
	}
	v := gl.Get(name)
	r := 0
	if v.Truthy() {
		r = v.Int()
	}
	gl.consts[name] = r
	return r
}

func (e *GLExtension) Const(name string) int {
	if r, ok := e.consts[name]; ok {
		return r
	}
	v := e.Get(name)
	r := 0
	if v.Truthy() {
		r = v.Int()
	}
	e.consts[name] = r
	return r
}
