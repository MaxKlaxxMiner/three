//go:build js

package webgl

import "syscall/js"

type Context struct {
	js.Value
	Consts map[string]int
}

type GLExtension struct {
	js.Value
	Consts map[string]int
}
