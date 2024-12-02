//go:build js

package renderers

import (
	"github.com/MaxKlaxxMiner/three/utils"
	"syscall/js"
)

type localValues struct {
	canvas js.Value
}

type WebGLRendererParams struct {
	Canvas *js.Value
}

func (r *WebGLRenderer) initParameters(parameters WebGLRendererParams) {
	if parameters.Canvas != nil {
		r.canvas = *parameters.Canvas
	} else {
		r.canvas = utils.CreateCanvasElement()
	}
	// 			context = null, todo
	// 			depth = true, todo
	// 			stencil = false, todo
	// 			alpha = false, todo
	// 			antialias = false, todo
	// 			premultipliedAlpha = true, todo
	// 			preserveDrawingBuffer = false, todo
	// 			powerPreference = 'default', todo
	// 			failIfMajorPerformanceCaveat = false, todo
	// 			reverseDepthBuffer = false, todo
}
