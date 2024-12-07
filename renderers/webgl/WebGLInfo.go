package webgl

import "strconv"

type Info struct {
	gl     Context
	Memory struct {
		Geometries int
		Textures   int
	}
	Render struct {
		Frame     int
		Calls     int
		Triangles int
		Points    int
		Lines     int
	}
	AutoReset bool
}

func NewWebGLInfo(gl Context) *Info {
	return &Info{gl: gl, AutoReset: true}
}

func (i *Info) Update(count int, mode int32, instanceCount int) {
	i.Render.Calls++

	switch mode {
	case i.gl.TRIANGLES:
		i.Render.Triangles += instanceCount * (count / 3)
	case i.gl.LINES:
		i.Render.Lines += instanceCount * (count / 2)
	case i.gl.LINE_STRIP:
		i.Render.Lines += instanceCount * (count - 1)
	case i.gl.LINE_LOOP:
		i.Render.Lines += instanceCount * count
	case i.gl.POINTS:
		i.Render.Points += instanceCount * count
	default:
		panic("THREE.WebGLInfo: Unknown draw mode:" + strconv.Itoa(int(mode)))
	}
}

func (i *Info) Reset() {
	i.Render.Calls = 0
	i.Render.Triangles = 0
	i.Render.Points = 0
	i.Render.Lines = 0
}
