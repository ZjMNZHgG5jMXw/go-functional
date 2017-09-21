package functional

import (
	"testing"
)

func TestMakeFoldr(t *testing.T) {
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

func TestSliceAppendReverse(t *testing.T) {
	var foldIntSlice func(func(int, []int) []int, []int, []int) []int
	MakeFoldr(&foldIntSlice)

	var expected = []int{5, 4, 3, 2, 1}

	res := foldIntSlice(
		func(add int, slice []int) []int {
			return append(slice, add)
		},
		[]int{},
		[]int{1, 2, 3, 4, 5})

	for i, v := range expected {
		if v != expected[i] {
			t.Errorf("expected %#v, got %#v", expected, res)
			break
		}
	}
}

func TestMakeFoldl(t *testing.T) {
	var sumInt func(func(int, int) int, int, []int) int
	MakeFoldl(&sumInt)

	var expected = 15

	res := sumInt(
		func(a, b int) int { return a + b },
		0,
		[]int{1, 2, 3, 4, 5})

	if expected != res {
		t.Errorf("expected %#v, got %#v", expected, res)
	}
}

func TestSliceAppendIdentity(t *testing.T) {
	var foldIntSlice func(func([]int, int) []int, []int, []int) []int
	MakeFoldl(&foldIntSlice)

	var expected = []int{1, 2, 3, 4, 5}

	res := foldIntSlice(
		func(slice []int, add int) []int {
			return append(slice, add)
		},
		[]int{},
		expected)

	for i, v := range expected {
		if v != expected[i] {
			t.Errorf("expected %#v, got %#v", expected, res)
			break
		}
	}
}
