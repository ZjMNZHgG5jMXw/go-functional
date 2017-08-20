package functional

import (
	"testing"
)

func TestFlip(t *testing.T) {
	actual := Flip(
		func(a int) func(int) int {
			return func(b int) int {
				return a - b
			}
		}).(func(int) func(int) int)(5)(11)

	if actual != 6 {
		t.Errorf("expected 6, got %d", actual)
	}
}

func TestAssignFlipped(t *testing.T) {
	var subFlipped func(int) func(int) int

	AssignFlipped(
		&subFlipped,
		func(a int) func(int) int {
			return func(b int) int {
				return a - b
			}
		})

	actual := subFlipped(5)(12)
	if actual != 7 {
		t.Errorf("expected 7, got %d", actual)
	}
}
