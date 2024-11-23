package util

func If[T any](cmp bool, okVal, elseVal T) T {
	if cmp {
		return okVal
	} else {
		return elseVal
	}
}

func IfFunc[T any](cmp bool, okVal, elseVal func() T) T {
	if cmp {
		return okVal()
	} else {
		return elseVal()
	}
}

func NotNullOrDefault[T any](val *T, defaultVal T) T {
	if val == nil {
		return defaultVal
	}
	return *val
}
