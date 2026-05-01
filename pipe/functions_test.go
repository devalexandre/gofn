package pipe

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	f := Filter[int](func(n int) bool { return n%2 == 0 })
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{2, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMap(t *testing.T) {
	f := Map[int, int](func(n int) int { return n * 2 })
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReduce(t *testing.T) {
	f := Reduce[int](func(a, b int) int { return a + b })
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 15
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSum(t *testing.T) {
	f := Sum[int]()
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 15
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestProduct(t *testing.T) {
	f := Product[int]()
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 120
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMin(t *testing.T) {
	f := Min[int]()
	result, err := f([]int{5, 2, 3, 4, 1})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 1
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMax(t *testing.T) {
	f := Max[int]()
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 5
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestUnique(t *testing.T) {
	f := Unique[int]()
	result, err := f([]int{1, 2, 2, 3, 4, 5, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestReverse(t *testing.T) {
	f := Reverse[int]()
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestJoin(t *testing.T) {
	f := Join[int]("-")
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := "1-2-3-4-5"
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestContains(t *testing.T) {
	element := 3
	f := Contains[int](element)
	result, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result {
		t.Errorf("Expected to find element %v in slice", element)
	}
}

func TestIndexOf(t *testing.T) {
	element := 4
	f := IndexOf[int](element)
	result, err := f([]int{1, 2, 3, 4, 5})
	expected := 3
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected index %d, got %d for element %v", expected, result, element)
	}
}

func TestForEach(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	sum := 0
	action := func(n int) { sum += n }
	f := ForEach[int](action)
	_, err := f(a)
	expectedSum := 15
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, sum)
	}
}

func TestPop(t *testing.T) {
	f := Pop[int]()
	lastElement, remainingSlice, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedLastElement := 5
	expectedRemaining := []int{1, 2, 3, 4}
	if lastElement != expectedLastElement || !reflect.DeepEqual(remainingSlice, expectedRemaining) {
		t.Errorf("Expected last element %d and remaining slice %v, got last element %d and remaining slice %v", expectedLastElement, expectedRemaining, lastElement, remainingSlice)
	}
}

func TestPush(t *testing.T) {
	elementsToAdd := []int{6, 7}
	f := Push(elementsToAdd...)
	resultSlice, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(resultSlice, expectedSlice) {
		t.Errorf("Expected slice %v, got %v", expectedSlice, resultSlice)
	}
}

func TestShift(t *testing.T) {
	f := Shift[int]()
	firstElement, remainingSlice, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedFirstElement := 1
	expectedRemaining := []int{2, 3, 4, 5}
	if firstElement != expectedFirstElement || !reflect.DeepEqual(remainingSlice, expectedRemaining) {
		t.Errorf("Expected first element %d and remaining slice %v, got first element %d and remaining slice %v", expectedFirstElement, expectedRemaining, firstElement, remainingSlice)
	}
}

func TestSort(t *testing.T) {
	lessFunc := func(i, j int) bool { return i < j }
	f := Sort[int](lessFunc)
	sortedSlice, err := f([]int{5, 3, 4, 1, 2})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedSlice := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(sortedSlice, expectedSlice) {
		t.Errorf("Expected sorted slice %v, got %v", expectedSlice, sortedSlice)
	}
}

func TestUnshift(t *testing.T) {
	elementsToAdd := []int{0, -1}
	f := Unshift[int](elementsToAdd...)
	resultSlice, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedSlice := []int{0, -1, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(resultSlice, expectedSlice) {
		t.Errorf("Expected slice %v, got %v", expectedSlice, resultSlice)
	}
}

func TestGroupByFunc(t *testing.T) {
	type Item struct {
		Name        string
		Price       float64
		Description string
		Qty         int
	}

	items := []Item{
		{"Item 1", 10.0, "Description 1", 1},
		{"Item 2", 20.0, "Description 2", 2},
		{"Item 3", 30.0, "Description 3", 3},
		{"Item 4", 40.0, "Description 4", 10},
		{"Item 4", 40.0, "Description 4", 15},
		{"Item 4", 40.0, "Description 4", 25},
	}

	groupByFunc := GroupBy(func(item Item) string {
		return fmt.Sprintf("%s - %.1f", item.Name, item.Price)
	})

	groupedItems, err := groupByFunc(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedGroupCount := 4
	if len(groupedItems) != expectedGroupCount {
		t.Errorf("Expected %d groups, got %d", expectedGroupCount, len(groupedItems))
	}

}

func TestGroupSumByFunc(t *testing.T) {
	type Item struct {
		Name        string
		Price       float64
		Description string
		Qty         int
	}

	items := []Item{
		{"Item 1", 10.0, "Description 1", 1},
		{"Item 2", 20.0, "Description 2", 2},
		{"Item 3", 30.0, "Description 3", 3},
		{"Item 4", 40.0, "Description 4", 10},
		{"Item 4", 40.0, "Description 4", 15},
		{"Item 4", 40.0, "Description 4", 25},
	}

	groupSumByFunc := GroupSumBy(
		func(item Item) string { return item.Name },
		func(item Item) int { return item.Qty },
	)

	sumByName, err := groupSumByFunc(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sumByName["Item 4"] != 50 {
		t.Errorf("Expected %d, got %d", 50, sumByName["Item 4"])
	}
}

func TestRowHelperFuncs(t *testing.T) {
	type Summary struct {
		TotalPrice float64
		TotalQty   int
	}
	type Item struct {
		Name        string
		Price       float64
		Description string
		Qty         int
	}

	items := []Item{
		{"Item 1", 10.0, "Description 1", 1},
		{"Item 2", 20.0, "Description 2", 2},
		{"Item 3", 30.0, "Description 3", 3},
		{"Item 4", 40.0, "Description 4", 10},
		{"Item 4", 40.0, "Description 4", 15},
		{"Item 4", 40.0, "Description 4", 25},
	}

	sumWhere, err := GroupSumByWhere(
		func(item Item) bool { return item.Qty > 3 },
		func(item Item) string { return item.Name },
		func(item Item) float64 { return item.Price },
	)(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sumWhere["Item 4"] != 120 {
		t.Errorf("Expected %v, got %v", 120.0, sumWhere["Item 4"])
	}

	count, err := GroupCountBy(func(item Item) string { return item.Name })(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if count["Item 4"] != 3 {
		t.Errorf("Expected %d, got %d", 3, count["Item 4"])
	}

	summary, err := GroupReduceBy(
		func(item Item) string { return item.Name },
		func(acc Summary, item Item) Summary {
			acc.TotalPrice += item.Price
			acc.TotalQty += item.Qty
			return acc
		},
	)(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if summary["Item 4"] != (Summary{TotalPrice: 120, TotalQty: 50}) {
		t.Errorf("Expected %v, got %v", Summary{TotalPrice: 120, TotalQty: 50}, summary["Item 4"])
	}

	stats, err := GroupStatsBy(
		func(item Item) string { return item.Name },
		func(item Item) float64 { return item.Price },
	)(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if stats["Item 4"].Count != 3 || stats["Item 4"].Avg != 40 {
		t.Errorf("Expected Item 4 stats count 3 avg 40, got %v", stats["Item 4"])
	}

	distinct, err := DistinctBy(func(item Item) string { return item.Name })(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(distinct) != 4 || distinct[3].Qty != 10 {
		t.Errorf("Expected first item for each name, got %v", distinct)
	}

	indexed, err := IndexBy(func(item Item) string { return item.Name })(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if indexed["Item 4"].Qty != 25 {
		t.Errorf("Expected last item for duplicated key, got %v", indexed["Item 4"])
	}

	highQty, lowQty, err := Partition(func(item Item) bool { return item.Qty > 3 })(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(highQty) != 3 || len(lowQty) != 3 {
		t.Errorf("Expected partition sizes 3 and 3, got %d and %d", len(highQty), len(lowQty))
	}

	sorted, err := SortBy(func(item Item) int { return item.Qty })(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sorted[0].Qty != 1 || sorted[len(sorted)-1].Qty != 25 {
		t.Errorf("Expected order by Qty, got %v", sorted)
	}

	taken, err := Take[Item](2)(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(taken) != 2 || taken[1].Name != "Item 2" {
		t.Errorf("Expected first two items, got %v", taken)
	}

	skipped, err := Skip[Item](3)(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(skipped) != 3 || skipped[0].Qty != 10 {
		t.Errorf("Expected last three items, got %v", skipped)
	}

	chunks, err := Chunk[Item](2)(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(chunks) != 3 || len(chunks[2]) != 2 || chunks[2][1].Qty != 25 {
		t.Errorf("Expected three chunks with two items, got %v", chunks)
	}
}
