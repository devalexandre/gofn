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
	f := Sort[int](lessFunc) // Especificando explicitamente que T é int.
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
	f := Unshift(elementsToAdd...)
	resultSlice, err := f([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedSlice := []int{-1, 0, 1, 2, 3, 4, 5}
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

	// Como o resultado é um slice de slices, a verificação exata de igualdade esperada é mais complicada.
	// Você precisará ajustar essa parte com base na lógica específica desejada para verificar os grupos.
	// Por exemplo, verificar a contagem de grupos e um ou mais elementos específicos.
	expectedGroupCount := 4 // Esperado 4 grupos distintos baseados no nome e preço.
	if len(groupedItems) != expectedGroupCount {
		t.Errorf("Expected %d groups, got %d", expectedGroupCount, len(groupedItems))
	}

}
