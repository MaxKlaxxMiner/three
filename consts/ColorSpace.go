package consts

type ColorSpace string

// Color space string identifiers, matching CSS Color Module Level 4 and WebGPU names where available.
const (
	NoColorSpace         ColorSpace = ""
	SRGBColorSpace       ColorSpace = "srgb"
	LinearSRGBColorSpace ColorSpace = "srgb-linear"

	LinearTransfer ColorSpace = "linear"
	SRGBTransfer   ColorSpace = "srgb"
)
