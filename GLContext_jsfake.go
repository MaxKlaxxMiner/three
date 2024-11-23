//go:build !js

package three

import (
	"github.com/MaxKlaxxMiner/three/internal/fake/js"
)

type GLContext struct {
	js.Value
}

type GLExtension struct {
	js.Value
}
