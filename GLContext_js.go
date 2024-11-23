//go:build js

package three

import "syscall/js"

type GLContext struct {
	js.Value
	consts map[string]int
}

type GLExtension struct {
	js.Value
	consts map[string]int
}
