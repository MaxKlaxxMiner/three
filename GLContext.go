package three

func (gl *GLContext) getExtension(name string) GLExtension {
	return GLExtension{gl.Call("getExtension", name)}
}
