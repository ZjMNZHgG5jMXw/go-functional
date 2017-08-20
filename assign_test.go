package functional

import (
	"testing"
)

func TestMakeAssignment(t *testing.T) {
	var (
		sub       = func(a, b int) int { return a - b }
		assignSub func(*int, int, int)
		actual    int
	)
	MakeAssignment(&assignSub, sub)

	assignSub(&actual, 10, 3)
	if actual != 7 {
		t.Errorf("expected 7, got %d", actual)
	}
}

func TestMakeAssignmentInterface(t *testing.T) {
	var (
		sub       = func(a, b int) interface{} { return a - b }
		assignSub func(interface{}, int, int)
		actual    int
	)
	MakeAssignment(&assignSub, sub)

	assignSub(&actual, 5, 2)
	if actual != 3 {
		t.Errorf("expected 3, got %d", actual)
	}
}

func TestMakeAssignmentPointer(t *testing.T) {
	var (
		sub       = func(a, b int) *int { c := a - b; return &c }
		assignSub func(*int, int, int)
		actual    int
	)
	MakeAssignment(&assignSub, sub)

	assignSub(&actual, 21, 7)
	if actual != 14 {
		t.Errorf("expected 14, got %d", actual)
	}
}

func TestMakeAssignmentNilPointer(t *testing.T) {
	var (
		sub       = func(a, b int) *int { return nil }
		assignSub func(*int, int, int)
		actual    int
	)
	MakeAssignment(&assignSub, sub)

	assignSub(&actual, 21, 7)
	if actual != 0 {
		t.Errorf("expected 0, got %d", actual)
	}
}

func TestMakeAssignmentKeepPointer(t *testing.T) {
	var (
		sub       = func(a, b int) *int { c := a - b; return &c }
		assignSub func(**int, int, int)
		actual    *int
	)
	MakeAssignment(&assignSub, sub, KeepPointerMode)

	assignSub(&actual, 21, 15)
	if *actual != 6 {
		t.Errorf("expected 6, got %d", *actual)
	}
}

func TestMakeAssignmentKeepNilPointer(t *testing.T) {
	var (
		sub       = func(a, b int) *int { return nil }
		assignSub func(**int, int, int)
		actual    *int
	)
	MakeAssignment(&assignSub, sub, KeepPointerMode)
	actual = func(i int) *int { return &i }(3)

	assignSub(&actual, 21, 16)
	if actual != nil {
		t.Errorf("expected nil, got %p", actual)
	}
}

func TestMakeAssignmentVariadic(t *testing.T) {
	var (
		sum = func(as ...int) (total int) {
			for _, a := range as {
				total += a
			}
			return
		}
		assignSum func(*int, ...int)
		actual    int
	)
	MakeAssignment(&assignSum, sum)

	assignSum(&actual, 1, 2, 3, 4)
	if actual != 10 {
		t.Errorf("expected 10, got %d", actual)
	}
}

func TestMakeAssignmentVariadicInterface(t *testing.T) {
	var (
		sum = func(as ...interface{}) interface{} {
			var total int
			for _, a := range as {
				total += a.(int)
			}
			return total
		}
		assignSum func(interface{}, ...interface{})
		actual    interface{}
	)
	MakeAssignment(&assignSum, sum)

	assignSum(&actual, 5, 6, 7, 8)
	if actual.(int) != 26 {
		t.Errorf("expected 26, got %d", actual)
	}
}

func TestMakeAssignmentVariadicPointer(t *testing.T) {
	var (
		sum = func(as ...int) (total *int) {
			total = new(int)
			for _, a := range as {
				*total += a
			}
			return
		}
		assignSum func(*int, ...int)
		actual    int
	)
	MakeAssignment(&assignSum, sum)

	assignSum(&actual, 1, 2, 3, 4, 5, 6, 7, 8)
	if actual != 36 {
		t.Errorf("expected 36, got %d", actual)
	}
}
