package array

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

/*
 * Array functions
 * Example:
 *   a := []int{1, 2, 3, 4, 5}
 *   b := Filter(a, func(x int) bool { return x%2 == 0 })
 *   fmt.Println(b) // [2 4]
 */
func Filter[T any](a []T, f func(T) bool) []T {
	var b []T
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func Find[T any](a *[]T, f func(T) bool) *T {
	for _, x := range *a {
		if f(x) {
			return &x
		}
	}
	return nil
}

/* Map
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Map(a, func(x int) int { return x * 2 })
*   fmt.Println(b) // [2 4 6 8 10]
 */

func Map[T, U any](a []T, f func(T) U) []U {
	var b []U
	for _, x := range a {
		b = append(b, f(x))
	}
	return b
}

/* FlatMap
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := FlatMap(a, func(x int) []int { return []int{x, x * 2} })
*   fmt.Println(b) // [1 2 2 4 3 6 4 8 5 10]
 */

func FlatMap[T, U any](a []T, f func(T) []U) []U {
	var b []U
	for _, x := range a {
		b = append(b, f(x)...)
	}
	return b
}

/* Reduce
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Reduce(a, func(x, y int) int { return x + y })
*   fmt.Println(b) // 15
 */

func Reduce[T any](a []T, f func(T, T) T) T {
	if len(a) == 0 {
		panic("empty array")
	}
	x := a[0]
	for _, y := range a[1:] {
		x = f(x, y)
	}
	return x
}

/* Any
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Any(a, func(x int) bool { return x%2 == 0 })
*   fmt.Println(b) // true
 */

func Any[T any](input []T, f func(T) bool) bool {
	for _, x := range input {
		if f(x) {
			return true
		}
	}
	return false
}

/* Some
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Some(a, func(x int) bool { return x%2 == 0 })
*   fmt.Println(b) // true
 */

func Some[T any](a []T, f func(T) bool) bool {
	for _, x := range a {
		if f(x) {
			return true
		}
	}
	return false
}

/* Every
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Every(a, func(x int) bool { return x%2 == 0 })
*   fmt.Println(b) // false
 */
func Every[T any](a []T, f func(T) bool) bool {
	for _, x := range a {
		if !f(x) {
			return false
		}
	}
	return true
}

/* Sum array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Sum(a)
*   fmt.Println(b) // 15
 */

func Sum[T Number](a []T) T {
	return Reduce(a, func(x, y T) T { return x + y })
}

/* Product array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Product(a)
*   fmt.Println(b) // 120
 */
func Product[T Number](a []T) T {
	return Reduce(a, func(x, y T) T { return x * y })
}

/* Min array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Min(a)
*   fmt.Println(b) // 1
 */
func Min[T Number](a []T) T {
	return Reduce(a, func(x, y T) T {
		if x < y {
			return x
		}
		return y
	})
}

/* Max array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Max(a)
*   fmt.Println(b) // 5
 */

func Max[T Number](a []T) T {
	return Reduce(a, func(x, y T) T {
		if x > y {
			return x
		}
		return y
	})
}

/* ForEach
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   ForEach(a, func(x int) { fmt.Println(x) })
 */
func ForEach[T any](a []T, f func(T)) {
	for _, x := range a {
		f(x)
	}
}

/* IndexOf
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := IndexOf(a, 3)
*   fmt.Println(b) // 2
 */
func IndexOf[T comparable](a []T, x T) int {
	for i, y := range a {
		if x == y {
			return i
		}
	}
	return -1
}

/* Contains
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Contains(a, 3)
*   fmt.Println(b) // true
 */
func Contains[T comparable](a []T, x T) bool {
	return IndexOf(a, x) != -1
}

/* Equals
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := []int{1, 2, 3, 4, 5}
*   c := Equals(a, b)
*   fmt.Println(c) // true

 */

func Equals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

/* Reverse
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Reverse(a)
*   fmt.Println(b) // [5 4 3 2 1]
 */
func Reverse[T any](a []T) []T {
	b := make([]T, len(a))
	for i, x := range a {
		b[len(a)-1-i] = x
	}
	return b
}

/* Shuffle
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Shuffle(a)
*   fmt.Println(b) // [2 5 4 3 1]
 */
func Shuffle[T any](a []T) []T {
	b := make([]T, len(a))
	for i, x := range rand.Perm(len(a)) {
		b[x] = a[i]
	}
	return b
}

/* Unique
* Example:
*   a := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
*   b := Unique(a)
*   fmt.Println(b) // [1 2 3 4 5]
 */

func Unique[T comparable](a []T) []T {
	b := make([]T, 0, len(a))
	for _, x := range a {
		if !Contains(b, x) {
			b = append(b, x)
		}
	}
	return b
}

/* Union
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := []int{6, 7, 8, 9, 10}
*   c := Union(a, b)
*   fmt.Println(c) // [1 2 3 4 5 6 7 8 9 10]
 */

func Union[T comparable](a, b []T) []T {
	return append(a, b...)
}

/* Fill
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Fill(a, 0, 3, 10)
*   fmt.Println(b) // [10 10 10 4 5]
 */

func Fill[T any](a []T, start, end int, x T) []T {
	b := make([]T, len(a))
	copy(b, a)
	for i := start; i < end; i++ {
		b[i] = x
	}
	return b
}

/* Join
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Join(a, "+")
*   fmt.Println(b) // 1+2+3+4+5
 */

func Join[T any](a []T, sep string) string {
	b := make([]string, len(a))
	for i, x := range a {
		b[i] = fmt.Sprint(x)
	}
	return strings.Join(b, sep)
}

/* Pop removes and returns the last element of the slice
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b, c := Pop(a)
*   fmt.Println(b, c) // 5 [1 2 3 4]
 */
func Pop[T any](a []T) (T, []T) {
	return a[len(a)-1], a[:len(a)-1]
}

/* Push adds one or more elements to the end of the slice
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Push(a, 6, 7, 8, 9, 10)
*   fmt.Println(b) // [1 2 3 4 5 6 7 8 9 10]
 */

func Push[T any](a []T, x ...T) []T {
	return append(a, x...)
}

/* Shift removes and returns the first element of the slice
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b, c := Shift(a)
*   fmt.Println(b, c) // 1 [2 3 4 5]
 */

func Shift[T any](a []T) (T, []T) {
	return a[0], a[1:]
}

/* Sort sorts the slice in increasing order
* Example:
*   a := []int{5, 4, 3, 2, 1}
*   b := Sort(a,func(i, j int) bool { return a[i] < a[j] })
*   fmt.Println(b) // [1 2 3 4 5]
 */

func Sort[T any](a []T, less func(i, j int) bool) []T {
	b := make([]T, len(a))
	copy(b, a)
	sort.Slice(b, less)
	return b
}

/* Unshift adds one or more elements to the beginning of the slice
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Unshift(a, 6, 7, 8, 9, 10)
*   fmt.Println(b) // [6 7 8 9 10 1 2 3 4 5]
 */

func Unshift[T any](a []T, x ...T) []T {
	return append(x, a...)
}

/* group w rows by key
* Example:
* type Itens struct {
* 		Name        string
* 		Price       float64
* 		Description string
* 		Qty         int
* 	}
*
* 	itens := []Itens{
* 		{"Item 1", 10.0, "Description 1", 1},
* 		{"Item 2", 20.0, "Description 2", 2},
* 		{"Item 3", 30.0, "Description 3", 3},
*  		{"Item 4", 40.0, "Description 4", 10},
*		{"Item 4", 40.0, "Description 4", 15},
*		{"Item 4", 40.0, "Description 4", 25},
*	}
*
*	grouped := array.GroupBy(itens, func(item Itens) string {
*		return fmt.Sprintf("%s - %v", item.Name, item.Price)
*	})
 */

func GroupBy[T any, K comparable](w []T, key func(T) K) map[K][]T {
	m := make(map[K][]T)
	for _, x := range w {
		if _, ok := m[key(x)]; !ok {
			m[key(x)] = make([]T, 0)
		}
	}
	return m
}
