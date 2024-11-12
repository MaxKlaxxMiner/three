//go:build !js

package three

import (
	"github.com/MaxKlaxxMiner/three/utils/fake/js"
)

type GLContext struct {
	js.Value
}
