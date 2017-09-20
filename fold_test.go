package functional

import (
	"testing"
)

func TestMakeFold(t *testing.T) {
	var sumInt func(func(int, int) int, int, []int) int
	MakeFoldr(&sumInt)

	var expected = 15

	res := sumInt(
		func(a, b int) int { return a + b },
		0,
		[]int{1, 2, 3, 4, 5})

	if expected != res {
		t.Errorf("expected %#v, got %#v", expected, res)
	}
}
