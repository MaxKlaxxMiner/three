//go:build !js

package renderers

import (
	"github.com/MaxKlaxxMiner/three/internal/js"
	"github.com/MaxKlaxxMiner/three/utils"
)

type localJsValues struct {
	canvas  js.Value
	context js.Value
}

type GlobalJsValues struct {
	DomElement js.Value
}

type WebGLRendererParams struct {
	Canvas                       *js.Value
	Context                      *js.Value
	Depth                        *bool
	Stencil                      *bool
	Alpha                        *bool
	Antialias                    *bool
	PremultipliedAlpha           *bool
	PreserveDrawingBuffer        *bool
	PowerPreference              string
	FailIfMajorPerformanceCaveat *bool
	ReverseDepthBuffer           *bool
}

func (r *WebGLRenderer) initParameters(parameters WebGLRendererParams) {
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
