//go:build !js

package utils

import (
	"github.com/MaxKlaxxMiner/three/utils/fake/js"
)

func CreateElementNS(name string) js.Value {
	return js.Value{}
}

func CreateCanvasElement() js.Value {
	return js.Value{}
}

func InstanceOf(value *js.Value, className string) bool {
	return false
}
