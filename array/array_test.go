package array_test

import (
	"github.com/devalexandre/gofn/array"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("test Filter", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Filter(func(x int) bool { return x%2 == 0 }).Value()
		if !reflect.DeepEqual(b, []int{2, 4}) {
			t.Error("Filter failed. Got", b, "Expected", []int{2, 4})
		}
	})

	t.Run("test Find", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Find(func(x int) bool { return x%2 == 0 })
		if *b != 2 {
			t.Error("Find failed. Got", b, "Expected", 2)
		}
	})

	t.Run("test Map", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Map(func(x int) int { return x * 2 }).Value()
		if !reflect.DeepEqual(b, []int{2, 4, 6, 8, 10}) {
			t.Error("Map failed. Got", b, "Expected", []int{2, 4, 6, 8, 10})
		}
	})

	t.Run("test flatMap", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).FlatMap(func(x int) []int { return []int{x * 2} }).Value()
		if !reflect.DeepEqual(b, []int{2, 4, 6, 8, 10}) {
			t.Error("FlatMap failed. Got", b, "Expected", []int{2, 4, 6, 8, 10})
		}

	})

	t.Run("test Reduce", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Reduce(func(x, y int) int { return x + y })
		if b != 15 {
			t.Error("Reduce failed. Got", b, "Expected", 15)
		}
	})

	t.Run("test Some", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Some(func(x int) bool { return x%2 == 0 })
		if !b {
			t.Error("Some failed. Got", b, "Expected", true)
		}
	})

	t.Run("test Every", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Every(func(x int) bool { return x%2 == 0 })
		if b {
			t.Error("Every failed. Got", b, "Expected", false)
		}
	})

	t.Run("test Any", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := array.NewArray(a).Any(func(x int) bool { return x%2 == 0 })
		if !b {
			t.Error("Any failed. Got", b, "Expected", true)
		}
	})

}
