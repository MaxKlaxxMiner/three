package tests

import "testing"

func unitPanic(t *testing.T, name, expectedPanic string, run func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s Expected panic: '%s', but none occurred", name, expectedPanic)
		} else {
			if r != expectedPanic {
				t.Errorf("%s Expected panic with message: '%s', but got '%v'", name, expectedPanic, r)
			}
		}
	}()

	run()
}
