package array_test

import (
	"fmt"
	"github.com/devalexandre/gofn/array"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("test Filter", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Filter(func(x int) bool { return x%2 == 0 })

		if !reflect.DeepEqual(b, array.Array[int]{2, 4}) {
			t.Error("Filter failed. Got", b, "Expected", array.Array[int]{2, 4})
		}

		if len(a) == len(b) {
			t.Error("Filter failed. Got", len(a), "Expected", len(b))
		}
	})

	t.Run("test Find", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Find(func(x int) bool { return x%2 == 0 })
		if b != 2 {
			t.Error("Find failed. Got", b, "Expected", 2)
		}
	})

	t.Run("test Map", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Map(func(x int) int { return x * 2 })
		if b[0] != 2 {
			t.Error("Map failed. Got", b, "Expected", array.Array[int]{2, 4, 6, 8, 10})
		}
	})

	t.Run("test flatMap", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.FlatMap(func(x int) []int { return []int{x * 2} })

		if b[0] != 2 {
			t.Error("flatMap failed. Got", b, "Expected", array.Array[int]{2, 4, 6, 8, 10})
		}

	})

	t.Run("test Reduce", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Reduce(func(x, y int) int { return x + y })
		if b != 15 {
			t.Error("Reduce failed. Got", b, "Expected", 15)
		}
	})

	t.Run("test Some", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Some(func(x int) bool { return x%2 == 0 })
		if !b {
			t.Error("Some failed. Got", b, "Expected", true)
		}
	})

	t.Run("test Every", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Every(func(x int) bool { return x%2 == 0 })
		if b {
			t.Error("Every failed. Got", b, "Expected", false)
		}
	})

	t.Run("test Any", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Any(func(x int) bool { return x%2 == 0 })
		if !b {
			t.Error("Any failed. Got", b, "Expected", true)
		}
	})

	t.Run("test Length", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Length()
		if b != 5 {
			t.Error("Length failed. Got", b, "Expected", 5)
		}
	})

	t.Run("test Reverse", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Reverse()
		if !reflect.DeepEqual(b, array.Array[int]{5, 4, 3, 2, 1}) {
			t.Error("Reverse failed. Got", b, "Expected", array.Array[int]{5, 4, 3, 2, 1})
		}
	})

	//Shuffle
	t.Run("test Shuffle", func(t *testing.T) {
		a := array.Array[int]{1, 2, 3, 4, 5}
		b := a.Shuffle()
		if reflect.DeepEqual(b, array.Array[int]{1, 2, 3, 4, 5}) {
			t.Error("Shuffle failed. Got", b, "Expected", array.Array[int]{5, 4, 3, 2, 1})
		}
	})

	t.Run("test Sort", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		a := array.Array[person]{
			{"John", 30},
			{"Doe", 25},
			{"Jane", 20},
		}

		b := a.Sort(func(i, j int) bool { return a[i].age < a[j].age })

		if b[0].age != 20 {
			t.Error("Sort failed. Got", b, "Expected", array.Array[person]{
				{"Jane", 20},
				{"Doe", 25},
				{"John", 30},
			})
		}
	})

	t.Run("test GroupBy", func(t *testing.T) {

		type Itens struct {
			Name        string
			Price       float64
			Description string
			Qty         int
		}

		itens := array.Array[Itens]{
			{"Item 1", 10.0, "Description 1", 1},
			{"Item 2", 20.0, "Description 2", 2},
			{"Item 3", 30.0, "Description 3", 3},
			{"Item 4", 40.0, "Description 4", 10},
			{"Item 4", 40.0, "Description 4", 15},
			{"Item 4", 40.0, "Description 4", 25},
		}

		grouped := itens.GroupBy(func(item Itens) string {
			return fmt.Sprintf("%s - %v", item.Name, item.Price)
		})

		if len(grouped) != 4 {
			t.Error("GroupBy failed. Got", len(grouped), "Expected", 4)
		}
	})
}
