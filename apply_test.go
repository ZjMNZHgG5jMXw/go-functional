package functional

import (
	"testing"
)

func TestApply(t *testing.T) {
	var (
		succ = func(i int) int { return i + 1 }
	)

	actual := Apply(succ, 4).(int)
	if actual != 5 {
		t.Errorf("expected 5, got %d", actual)
	}
}

func TestAssignApplied(t *testing.T) {
	var (
		succ   = func(i int) int { return i + 1 }
		actual int
	)

	AssignApplied(&actual, succ, 9)
	if actual != 10 {
		t.Errorf("expected 10, got %d", actual)
	}
}
