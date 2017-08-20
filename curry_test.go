package functional

import (
	"testing"
)

func TestCurry(t *testing.T) {
	actual := Curry(
		func(a, b int) int {
			return a - b
		}).(func(int) func(int) int)(5)(3)
	if actual != 2 {
		t.Errorf("expected 2, got %d", actual)
	}
}

func TestUncurry(t *testing.T) {
	actual := Uncurry(
		func(a int) func(int) int {
			return func(b int) int {
				return a - b
			}
		}).(func(int, int) int)(5, 2)
	if actual != 3 {
		t.Errorf("expected 3, got %d", actual)
	}
}
