package tests

import (
	"github.com/MaxKlaxxMiner/three/util"
	"testing"
)

func TestDummy(t *testing.T) {
	// --- goExtras ---
	_ = util.If(true, 1, 0)
	_ = util.If(false, 1, 0)
	_ = util.IfFunc(true, func() int { return 1 }, func() int { return 0 })
	_ = util.IfFunc(false, func() int { return 1 }, func() int { return 0 })
	_ = util.NotNullOrDefault((*int)(nil), 1)
	tmp := 1
	_ = util.NotNullOrDefault(&tmp, 1)

	// --- utils_jsfake ---
	_ = util.CreateElementNS("")
	_ = util.CreateCanvasElement()
	_ = util.InstanceOf(nil, "")
	_ = util.JsNull()
	_ = util.JsUndefined()
	_ = util.FuncOf(nil)
	_ = util.JsValue{}.AsJsValue()
}
