package array_test

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/devalexandre/gofn/array"
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
		if len(b) != len(a) {
			t.Error("Shuffle failed. Got length", len(b), "Expected", len(a))
		}
		sorted := slices.Clone(b)
		slices.Sort(sorted)
		if !reflect.DeepEqual(sorted, a) {
			t.Error("Shuffle failed. Got", b, "Expected same elements as", a)
		}
		if !reflect.DeepEqual(a, array.Array[int]{1, 2, 3, 4, 5}) {
			t.Error("Shuffle mutated input. Got", a, "Expected", array.Array[int]{1, 2, 3, 4, 5})
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

		key := "Item 4 - 40"
		if len(grouped[key]) != 3 {
			t.Error("GroupBy failed. Got", grouped[key], "Expected 3 items")
		}
		if grouped[key][0].Qty != 10 || grouped[key][1].Qty != 15 || grouped[key][2].Qty != 25 {
			t.Error("GroupBy failed. Got", grouped[key], "Expected items in original order")
		}
	})

	t.Run("test GroupSumBy", func(t *testing.T) {

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

		priceByName := itens.
			Filter(func(item Itens) bool {
				return item.Qty > 3
			}).
			GroupSumBy(
				func(item Itens) string { return item.Name },
				func(item Itens) float64 { return item.Price },
			)

		if priceByName["Item 4"] != 120.0 {
			t.Error("GroupSumBy failed. Got", priceByName["Item 4"], "Expected", 120.0)
		}
	})

	t.Run("test row helper chain methods", func(t *testing.T) {

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

		sum := itens.GroupSumByWhere(
			func(item Itens) bool { return item.Qty > 3 },
			func(item Itens) string { return item.Name },
			func(item Itens) float64 { return item.Price },
		)
		if sum["Item 4"] != 120 {
			t.Error("GroupSumByWhere failed. Got", sum["Item 4"], "Expected", 120)
		}

		count := itens.GroupCountBy(func(item Itens) string { return item.Name })
		if count["Item 4"] != 3 {
			t.Error("GroupCountBy failed. Got", count["Item 4"], "Expected", 3)
		}

		stats := itens.GroupStatsBy(
			func(item Itens) string { return item.Name },
			func(item Itens) float64 { return item.Price },
		)
		if stats["Item 4"].Avg != 40 {
			t.Error("GroupStatsBy failed. Got", stats["Item 4"], "Expected avg", 40)
		}

		distinct := itens.DistinctBy(func(item Itens) string { return item.Name })
		if len(distinct) != 4 || distinct[3].Qty != 10 {
			t.Error("DistinctBy failed. Got", distinct, "Expected first item for each name")
		}

		indexed := itens.IndexBy(func(item Itens) string { return item.Name })
		if indexed["Item 4"].Qty != 25 {
			t.Error("IndexBy failed. Got", indexed["Item 4"], "Expected last item for duplicated key")
		}

		highQty, lowQty := itens.Partition(func(item Itens) bool { return item.Qty > 3 })
		if len(highQty) != 3 || len(lowQty) != 3 {
			t.Error("Partition failed. Got", len(highQty), len(lowQty), "Expected", 3, 3)
		}

		byName := itens.SortByString(func(item Itens) string { return item.Name })
		if byName[0].Name != "Item 1" || byName[len(byName)-1].Name != "Item 4" {
			t.Error("SortByString failed. Got", byName)
		}

		byPrice := itens.SortByFloat64(func(item Itens) float64 { return item.Price })
		if byPrice[0].Price != 10 || byPrice[len(byPrice)-1].Price != 40 {
			t.Error("SortByFloat64 failed. Got", byPrice)
		}

		if got := itens.Take(2); len(got) != 2 || got[1].Name != "Item 2" {
			t.Error("Take failed. Got", got, "Expected first two items")
		}
		if got := itens.Skip(3); len(got) != 3 || got[0].Qty != 10 {
			t.Error("Skip failed. Got", got, "Expected last three items")
		}
		chunks := itens.Chunk(2)
		if len(chunks) != 3 || len(chunks[2]) != 2 || chunks[2][1].Qty != 25 {
			t.Error("Chunk failed. Got", chunks, "Expected three chunks with two items")
		}
	})
}
