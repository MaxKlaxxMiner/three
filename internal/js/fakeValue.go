package js

type Value struct{}
type Func struct {
	Value
}

func FuncOf(fn func(this Value, args []Value) any) Func {
	return Func{}
}

func Global() Value {
	return Value{}
}

func Null() Value {
	return Value{}
}

func Undefined() Value {
	return Value{}
}

func (v Value) IsNull() bool {
	return true
}

func (v Value) Bool() bool {
	return false
}

func (v Value) Int() int {
	return 0
}

func (v Value) String() string {
	return ""
}

func (v Value) Call(m string, args ...any) Value {
	return Value{}
}

func (v Value) Get(p string) Value {
	return Value{}
}

func (v Value) Set(p string, x any) {
}

func (v Value) IsUndefined() bool {
	return true
}

func (v Value) Truthy() bool {
	return false
}

func (v Value) InstanceOf(t Value) bool {
	return false
}

func (v Value) New(args ...any) Value {
	return Value{}
}
