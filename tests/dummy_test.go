package tests

import (
	"github.com/MaxKlaxxMiner/three/utils"
	"testing"
)

func TestDummy(t *testing.T) {
	// --- goExtras ---
	_ = utils.If(true, 1, 0)
	_ = utils.If(false, 1, 0)
	_ = utils.IfFunc(true, func() int { return 1 }, func() int { return 0 })
	_ = utils.IfFunc(false, func() int { return 1 }, func() int { return 0 })
	_ = utils.NotNullOrDefault((*int)(nil), 1)
	tmp := 1
	_ = utils.NotNullOrDefault(&tmp, 1)

	// --- utils_jsfake ---
	_ = utils.CreateElementNS("")
	_ = utils.CreateCanvasElement()
	_ = utils.InstanceOf(nil, "")
	_ = utils.JsNull()
	_ = utils.JsUndefined()
	_ = utils.FuncOf(nil)
	_ = utils.JsValue{}.AsJsValue()
}
