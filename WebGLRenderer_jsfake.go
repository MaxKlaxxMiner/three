//go:build !js

package three

import (
	"github.com/MaxKlaxxMiner/three/utils/fake/js"
)

type WebGLRenderer struct {
	AutoClear        bool // Defines whether the renderer should automatically clear its output before rendering a frame. Default is true.
	AutoClearColor   bool // If autoClear is true, defines whether the renderer should clear the color buffer. Default is true.
	AutoClearDepth   bool // If autoClear is true, defines whether the renderer should clear the depth buffer. Default is true.
	AutoClearStencil bool // If autoClear is true, defines whether the renderer should clear the stencil buffer. Default is true.
	Debug            struct {
		// If it is true, defines whether material shader programs are checked for errors during compilation and linkage process.
		// It may be useful to disable this check in production for performance gain. It is strongly recommended to keep these checks enabled during development.
		// If the shader does not compile and link - it will not work and associated material will not render. Default is true.
		CheckShaderErrors bool
		OnShaderError     func() // Callback for custom error reporting. todo: params
	}
	// A canvas where the renderer draws its output. This is automatically created by the renderer in the constructor (if not provided already); you just need to add it to your page like so:
	// 	js.Global().Get("document").Get("body").Call("appendChild", renderer.DomElement)
	DomElement js.Value
	// Defines whether the renderer should sort objects. Default is true.
	// 	Note: Sorting is used to attempt to properly render objects that have some degree of transparency. By definition, sorting objects may not work in all cases. Depending on the needs of application, it may be necessary to turn off sorting and use other methods to deal with transparency rendering e.g. manually determining each object's rendering order.
	SortObjects bool

	renderParams     // internal Render Params
	renderProperties // internal Properties/Variables
}

type WebGLRendererParams struct {
	Canvas                       *js.Value // A Canvas where the renderer draws its output. This corresponds to the domElement property below. If not passed in here, a new canvas element will be created.
	Context                      *js.Value // This can be used to attach the renderer to an existing RenderingContext. Default is null.
	Depth                        *bool     // whether the drawing buffer has a Depth buffer of at least 16 bits. Default is true.
	Stencil                      *bool     // whether the drawing buffer has a Stencil buffer of at least 8 bits. Default is false.
	Alpha                        *bool     // controls the default clear Alpha value. When set to true, the value is 0. Otherwise it's 1. Default is false.
	Antialias                    *bool     // whether to perform antialiasing. Default is false.
	PremultipliedAlpha           *bool     // whether the renderer will assume that colors have premultiplied alpha. Default is true.
	PreserveDrawingBuffer        *bool     // whether to preserve the buffers until manually cleared or overwritten. Default is false.
	PowerPreference              string    // Provides a hint to the user agent indicating what configuration of GPU is suitable for this WebGL context. Can be "high-performance", "low-power" or "default". Default is "default". See WebGL spec for details.
	FailIfMajorPerformanceCaveat *bool     // whether the renderer creation will fail upon low performance is detected. Default is false. See WebGL spec for details.
	ReverseDepthBuffer           *bool     // whether to use a reverse depth buffer. Requires the EXT_clip_control extension. This is a more faster and accurate version than logarithmic depth buffer. Default is false.
}

type renderParams struct {
	depth                        bool
	stencil                      bool
	alpha                        bool
	antialias                    bool
	premultipliedAlpha           bool
	preserveDrawingBuffer        bool
	powerPreference              string
	failIfMajorPerformanceCaveat bool
	reverseDepthBuffer           bool
}

type renderProperties struct {
	canvas  js.Value
	context js.Value
}

func (p *WebGLRendererParams) getBaseRenderer() *WebGLRenderer {
	return new(WebGLRenderer)
}

// --- API ---

func (r *WebGLRenderer) GetContext() js.Value {
	return r.context
}
