package array_test

import (
	"fmt"
	"github.com/devalexandre/gofn/array"
	"reflect"
	"testing"
)

//test array.Filter
func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Filter(a, func(x int) bool { return x%2 == 0 })
	if !reflect.DeepEqual(b, []int{2, 4}) {
		t.Error("Filter failed. Got", b, "Expected", []int{2, 4})
	}

	if len(a) == len(b) {
		t.Error("Filter failed. Got", len(a), "Expected", len(b))
	}

	fmt.Println(a)
}

func TestFilter0Loc(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Filter0Loc(a, func(x int) bool { return x%2 == 0 })
	if !reflect.DeepEqual(b, []int{2, 4}) {
		t.Error("Filter failed. Got", b, "Expected", []int{2, 4})
	}

	if len(a) == len(b) {
		t.Error("Filter failed. Got", len(a), "Expected", len(b))
	}

	fmt.Println(a)
}

//test array.Find
func TestFind(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b, _ := array.Find(a, func(x int) bool { return x%2 == 0 })
	if b != 2 {
		t.Error("Find failed. Got", b, "Expected", 2)
	}
}

//test array.Map
func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Map(a, func(x int) int { return x * 2 })
	if !reflect.DeepEqual(b, []int{2, 4, 6, 8, 10}) {
		t.Error("Map failed. Got", b, "Expected", []int{2, 4, 6, 8, 10})
	}
}

//test array.Reduce
func TestReduce(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Reduce(a, func(x, y int) int { return x + y })
	if b != 15 {
		t.Error("Reduce failed. Got", b, "Expected", 15)
	}
}

//test array.Some
func TestSome(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Some(a, func(x int) bool { return x%2 == 0 })
	if !b {
		t.Error("Some failed. Got", b, "Expected", true)
	}
}

//test array.Every
func TestEvery(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Every(a, func(x int) bool { return x%2 == 0 })
	if b {
		t.Error("Every failed. Got", b, "Expected", false)
	}
}

//test array.Any
func TestAny(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Any(a, func(x int) bool { return x%2 == 0 })
	if !b {
		t.Error("Any failed. Got", b, "Expected", true)
	}
}

//test array.FlatMap
func TestFlatMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.FlatMap(a, func(x int) []int { return []int{x, x * 2} })
	if !reflect.DeepEqual(b, []int{1, 2, 2, 4, 3, 6, 4, 8, 5, 10}) {
		t.Error("FlatMap failed. Got", b, "Expected", []int{1, 2, 2, 4, 3, 6, 4, 8, 5, 10})
	}
}

//test array.ForEach
func TestForEach(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	array.ForEach(a, func(x int) { x = x * 2 })
	if !reflect.DeepEqual(a, []int{1, 2, 3, 4, 5}) {
		t.Error("ForEach failed. Got", a, "Expected", []int{1, 2, 3, 4, 5})
	}
}

//test array.IndexOf
func TestIndexOf(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.IndexOf(a, 3)
	if b != 2 {
		t.Error("IndexOf failed. Got", b, "Expected", 2)
	}
}

//test array.Sum
func TestSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Sum(a)
	if b != 15 {
		t.Error("Sum failed. Got", b, "Expected", 15)
	}
}

//test array.Max
func TestMax(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Max(a)
	if b != 5 {
		t.Error("Max failed. Got", b, "Expected", 5)
	}
}

//test array.Min
func TestMin(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Min(a)
	if b != 1 {
		t.Error("Min failed. Got", b, "Expected", 1)
	}
}

//test array.Reverse
func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Reverse(a)
	if !reflect.DeepEqual(b, []int{5, 4, 3, 2, 1}) {
		t.Error("Reverse failed. Got", b, "Expected", []int{5, 4, 3, 2, 1})
	}
}

//test array.Product
func TestProduct(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Product(a)
	if b != 120 {
		t.Error("Product failed. Got", b, "Expected", 120)
	}
}

//test array.Equals
func TestEquals(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	c := array.Equals(a, b)
	if !c {
		t.Error("Equals failed. Got", c, "Expected", true)
	}
}

//test array.Contains
func TestContains(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Contains(a, 3)
	if !b {
		t.Error("Contains failed. Got", b, "Expected", true)
	}
}

//test array.Unique
func TestUnique(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	b := array.Unique(a)
	if !reflect.DeepEqual(b, []int{1, 2, 3, 4, 5}) {
		t.Error("Unique failed. Got", b, "Expected", []int{1, 2, 3, 4, 5})
	}
}

//test array.Union
func TestUnion(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	c := array.Union(a, b)
	if !reflect.DeepEqual(c, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Error("Union failed. Got", c, "Expected", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

//test array.Fill
func TestFill(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Fill(a, 0, 2, 10)
	if !reflect.DeepEqual(b, []int{10, 10, 3, 4, 5}) {
		t.Error("Fill failed. Got", b, "Expected", []int{10, 10, 3, 4, 5})
	}
}

//test array.Join
func TestJoin(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Join(a, "+")
	if b != "1+2+3+4+5" {
		t.Error("Join failed. Got", b, "Expected", "1+2+3+4+5")
	}
}

// test Pop
func TestPop(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	removed, newArray := array.Pop(a)
	if removed != 5 {
		t.Error("Pop failed. Got", removed, "Expected", 5)
	}
	if !reflect.DeepEqual(newArray, []int{1, 2, 3, 4}) {
		t.Error("Pop failed. Got", newArray, "Expected", []int{1, 2, 3, 4})
	}
}

//test Push
func TestPush(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Push(a, 6)
	if !reflect.DeepEqual(b, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Push failed. Got", b, "Expected", []int{1, 2, 3, 4, 5, 6})
	}
}

//test Shift
func TestShift(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	removed, newArray := array.Shift(a)
	if removed != 1 {
		t.Error("Shift failed. Got", removed, "Expected", 1)
	}
	if !reflect.DeepEqual(newArray, []int{2, 3, 4, 5}) {
		t.Error("Shift failed. Got", newArray, "Expected", []int{2, 3, 4, 5})
	}
}

// test Sort
func TestSort(t *testing.T) {
	a := []int{3, 2, 1, 5, 4}
	b := array.Sort(a, func(i, j int) bool { return a[i] < a[j] })

	if !reflect.DeepEqual(b, []int{1, 2, 3, 4, 5}) {
		t.Error("Sort failed. Got", b, "Expected", []int{1, 2, 3, 4, 5})
	}
}

//test sort with struct
func TestSortWithStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	a := []person{
		{"John", 30},
		{"Doe", 25},
		{"Jane", 20},
	}
	b := array.Sort(a, func(i, j int) bool { return a[i].age < a[j].age })

	if !reflect.DeepEqual(b, []person{
		{"Jane", 20},
		{"Doe", 25},
		{"John", 30},
	}) {
		t.Error("Sort failed. Got", b, "Expected", []person{
			{"Jane", 20},
			{"Doe", 25},
			{"John", 30},
		})
	}
}

//test Unshift

func TestUnshift(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Unshift(a, 0)
	if !reflect.DeepEqual(b, []int{0, 1, 2, 3, 4, 5}) {
		t.Error("Unshift failed. Got", b, "Expected", []int{0, 1, 2, 3, 4, 5})
	}
}

// test Shuffle
func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := array.Shuffle(a)
	if reflect.DeepEqual(a, b) {
		t.Error("Shuffle failed. Got", b, "Expected", []int{1, 2, 3, 4, 5})
	}
}

//GroupBy
func TestGroupBy(t *testing.T) {

	type Itens struct {
		Name        string
		Price       float64
		Description string
		Qty         int
	}

	itens := []Itens{
		{"Item 1", 10.0, "Description 1", 1},
		{"Item 2", 20.0, "Description 2", 2},
		{"Item 3", 30.0, "Description 3", 3},
		{"Item 4", 40.0, "Description 4", 10},
		{"Item 4", 40.0, "Description 4", 15},
		{"Item 4", 40.0, "Description 4", 25},
	}

	grouped := array.GroupBy(itens, func(item Itens) string {
		return fmt.Sprintf("%s - %v", item.Name, item.Price)
	})

	if len(grouped) != 4 {
		t.Error("GroupBy failed. Got", len(grouped), "Expected", 4)
	}
}
