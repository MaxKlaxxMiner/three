//go:build js

package webgl

import "syscall/js"

type Context struct {
	js.Value
	ContextConsts
}

type Extension struct {
	js.Value
	Consts map[string]int32
}
