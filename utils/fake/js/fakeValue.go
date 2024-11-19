//go:build !js

package js

type Value struct{}

func (v Value) IsNull() bool {
	return true
}

func (v Value) Bool() bool {
	return true
}

func (v Value) Call(m string, args ...any) Value {
	return Value{}
}

func (v Value) Get(p string) Value {
	return Value{}
}

func (v Value) IsUndefined() bool {
	return true
}
