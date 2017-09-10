package functional

import (
	"testing"
)

func TestMakeMap(t *testing.T) {
	var mapInt func(func(int) int, []int) []int
	MakeMap(&mapInt)

	var expected = []int{3, 4, 5, 6, 7}

	res := mapInt(
		func(a int) int { return a + 2 },
		[]int{1, 2, 3, 4, 5})

	for i, v := range expected {
		if v != res[i] {
			t.Errorf("expected %#v, got %#v", expected, res)
		}
	}
}
