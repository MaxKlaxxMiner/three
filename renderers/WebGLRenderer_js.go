//go:build js

package renderers

import (
	"github.com/MaxKlaxxMiner/three/renderers/webgl"
	"github.com/MaxKlaxxMiner/three/utils"
	"syscall/js"
)

type localJsValues struct {
	canvas  js.Value
	context js.Value
}

type GlobalJsValues struct {
	DomElement js.Value
}

func (r *WebGLRenderer) initParameters(parameters webgl.RendererParams) {
	r.parameters = parameters
	if parameters.Canvas != nil {
		r.canvas = *parameters.Canvas
	} else {
		r.canvas = utils.CreateCanvasElement()
	}
	r.context = utils.NotNullOrDefault(parameters.Context, js.Null())
	r.depth = utils.NotNullOrDefault(parameters.Depth, true)
	r.stencil = utils.NotNullOrDefault(parameters.Stencil, false)
	r.alpha = utils.NotNullOrDefault(parameters.Alpha, false)
	r.antialias = utils.NotNullOrDefault(parameters.Antialias, false)
	r.premultipliedAlpha = utils.NotNullOrDefault(parameters.PremultipliedAlpha, true)
	r.preserveDrawingBuffer = utils.NotNullOrDefault(parameters.PreserveDrawingBuffer, false)
	r.powerPreference = utils.If(len(parameters.PowerPreference) != 0, parameters.PowerPreference, "default")
	r.failIfMajorPerformanceCaveat = utils.NotNullOrDefault(parameters.FailIfMajorPerformanceCaveat, false)
	r.reverseDepthBuffer = utils.NotNullOrDefault(parameters.ReverseDepthBuffer, false)
}
