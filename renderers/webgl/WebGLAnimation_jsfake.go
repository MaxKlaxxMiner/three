//go:build !js

package webgl

import "github.com/MaxKlaxxMiner/three/internal/js"

type Animation struct {
	context       js.Value
	isAnimating   bool
	animationLoop func(time float64, frame int)
	requestId     int
	frame         int
}

func NewWebGLAnimation() *Animation {
	return &Animation{context: js.Global()}
}

func (a *Animation) onAnimationFrame(time float64, frame int) {
	a.animationLoop(time, frame)

	a.requestId = a.context.Call("requestAnimationFrame", js.FuncOf(func(_ js.Value, args []js.Value) any {
		a.frame++
		a.onAnimationFrame(args[0].Float(), a.frame)
		return nil
	})).Int()
}

func (a *Animation) Start() {
	if a.isAnimating {
		return
	}
	if a.animationLoop == nil {
		return
	}
	a.requestId = a.context.Call("requestAnimationFrame", js.FuncOf(func(_ js.Value, args []js.Value) any {
		a.frame++
		a.onAnimationFrame(args[0].Float(), a.frame)
		return nil
	})).Int()

	a.isAnimating = true
}

func (a *Animation) Stop() {
	a.context.Call("cancelAnimationFrame", a.requestId)

	a.isAnimating = false
}

func (a *Animation) SetAnimationLoop(callback func(time float64, frame int)) {
	a.animationLoop = callback
}

func (a *Animation) SetContext(value js.Value) {
	a.context = value
}
