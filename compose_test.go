package functional

import (
	"testing"
)

func TestCompose(t *testing.T) {
	var (
		sum = func(a, b int) int { return a + b }
		dbl = func(a int) int { return 2 * a }
	)

	actual := Compose(dbl, dbl, dbl, dbl, sum).(func(int, int) int)(1, 2)
	if actual != 48 {
		t.Errorf("expected 48, got %d", actual)
	}
}

func TestAssignComposed(t *testing.T) {
	var (
		sum      = func(a, b int) int { return a + b }
		dbl      = func(a int) int { return 2 * a }
		composed func(int, int) int
	)

	AssignComposed(&composed, dbl, dbl, dbl, dbl, sum)
	actual := composed(3, 4)
	if actual != 112 {
		t.Errorf("expected 112, got %d", actual)
	}
}

func TestAssignWithNilErrorPassing(t *testing.T) {
	var (
		sum      = func(a, b int) (int, error) { return a + b, nil }
		dbl      = func(a int) int { return 2 * a }
		composed func(int, int) (int, error)
	)

	AssignComposed(&composed, dbl, dbl, dbl, dbl, sum)
	actual, err := composed(3, 4)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if actual != 112 {
		t.Errorf("expected 112, got %d", actual)
	}
}

type ErrTest struct{}

func (ErrTest) Error() string { return "test error" }

func TestAssignWithNonNilErrorPassing(t *testing.T) {
	var (
		errTest  = ErrTest{}
		sum      = func(a, b int) (int, error) { return a + b, errTest }
		dbl      = func(a int) int { return 2 * a }
		composed func(int, int) (int, error)
	)

	AssignComposed(&composed, dbl, dbl, dbl, dbl, sum)
	_, err := composed(3, 4)
	if err != errTest {
		t.Errorf("expected test error, got %s", err)
	}
}
