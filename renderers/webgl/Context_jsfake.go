//go:build !js

package webgl

import (
	"github.com/MaxKlaxxMiner/three/internal/js"
)

type Context struct {
	js.Value
	ContextConsts
}

type Extension struct {
	js.Value
	Consts map[string]int32
}
