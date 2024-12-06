//go:build js

package renderers

import "syscall/js"

type GLContext struct {
	js.Value
	consts map[string]int
}

type GLExtension struct {
	js.Value
	consts map[string]int
}
