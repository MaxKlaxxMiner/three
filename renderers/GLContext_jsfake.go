//go:build !js

package renderers

import (
	"github.com/MaxKlaxxMiner/three/internal/js"
)

type GLContext struct {
	js.Value
	consts map[string]int
}

type GLExtension struct {
	js.Value
	consts map[string]int
}
