//go:build !js

package webgl

import "github.com/MaxKlaxxMiner/three/internal/js"

type RendererParams struct {
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
