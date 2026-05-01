package array

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type GroupStats[V Number] struct {
	Count int
	Sum   V
	Min   V
	Max   V
	Avg   float64
}

func Filter[T any](a []T, f func(T) bool) []T {
	y := make([]T, 0, len(a))
	for _, x := range a {
		if f(x) {
			y = append(y, x)
		}
	}
	return y
}

// with side effects
func Filter0Loc[T any](a []T, f func(T) bool) []T {
	i := 0
	for _, x := range a {
		if f(x) {
			a[i] = x
			i++
		}
	}
	return a[:i]
}

func Find[T any](a []T, f func(T) bool) T {
	for _, x := range a {
		if f(x) {
			return x
		}
	}
	var zero T
	return zero
}

/* Map
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Map(a, func(x int) int { return x * 2 })
*   fmt.Println(b) // [2 4 6 8 10]
 */

func Map[T, U any](a []T, f func(T) U) []U {
	b := make([]U, len(a))
	for i, x := range a {
		b[i] = f(x)
	}

	return b
}

// with side effects
func Map0Loc[T any](a []T, f func(T) T) {
	for i, x := range a {
		a[i] = f(x)
	}
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
	return Any(a, f)
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
	if len(a) == 0 {
		panic("empty array")
	}
	total := a[0]
	for _, x := range a[1:] {
		total += x
	}
	return total
}

/* Product array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Product(a)
*   fmt.Println(b) // 120
 */
func Product[T Number](a []T) T {
	if len(a) == 0 {
		panic("empty array")
	}
	total := a[0]
	for _, x := range a[1:] {
		total *= x
	}
	return total
}

/* Min array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Min(a)
*   fmt.Println(b) // 1
 */
func Min[T Number](a []T) T {
	if len(a) == 0 {
		panic("empty array")
	}
	minimum := a[0]
	for _, x := range a[1:] {
		if minimum < x {
			continue
		}
		minimum = x
	}
	return minimum
}

/* Max array of numbers
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Max(a)
*   fmt.Println(b) // 5
 */

func Max[T Number](a []T) T {
	if len(a) == 0 {
		panic("empty array")
	}
	maximum := a[0]
	for _, x := range a[1:] {
		if maximum > x {
			continue
		}
		maximum = x
	}
	return maximum
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
	return slices.Index(a, x)
}

/* Contains
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Contains(a, 3)
*   fmt.Println(b) // true
 */
func Contains[T comparable](a []T, x T) bool {
	return slices.Contains(a, x)
}

/* Equals
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := []int{1, 2, 3, 4, 5}
*   c := Equals(a, b)
*   fmt.Println(c) // true

 */

func Equals[T comparable](a, b []T) bool {
	return slices.Equal(a, b)
}

/* Reverse
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Reverse(a)
*   fmt.Println(b) // [5 4 3 2 1]
 */
func Reverse[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	slices.Reverse(b)
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
	copy(b, a)
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
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
	seen := make(map[T]struct{}, len(a))
	for _, x := range a {
		if _, ok := seen[x]; ok {
			continue
		}
		seen[x] = struct{}{}
		b = append(b, x)
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

func Union[T any](a, b []T) []T {
	if len(a)+len(b) == 0 {
		return append(a, b...)
	}

	c := make([]T, 0, len(a)+len(b))
	c = append(c, a...)
	c = append(c, b...)
	return c
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
		k := key(x)
		m[k] = append(m[k], x)
	}

	return m
}

/* GroupSumBy groups rows by key and sums a numeric value for each group.
* Example:
*   totalByName := GroupSumBy(items,
*       func(item Item) string { return item.Name },
*       func(item Item) float64 { return item.Price },
*   )
 */
func GroupSumBy[T any, K comparable, V Number](w []T, key func(T) K, value func(T) V) map[K]V {
	m := make(map[K]V)
	for _, x := range w {
		m[key(x)] += value(x)
	}

	return m
}

func GroupSumByWhere[T any, K comparable, V Number](w []T, where func(T) bool, key func(T) K, value func(T) V) map[K]V {
	m := make(map[K]V)
	for _, x := range w {
		if !where(x) {
			continue
		}
		m[key(x)] += value(x)
	}

	return m
}

func GroupCountBy[T any, K comparable](w []T, key func(T) K) map[K]int {
	m := make(map[K]int)
	for _, x := range w {
		m[key(x)]++
	}

	return m
}

func GroupReduceBy[T any, K comparable, A any](w []T, key func(T) K, reduce func(A, T) A) map[K]A {
	m := make(map[K]A)
	for _, x := range w {
		k := key(x)
		m[k] = reduce(m[k], x)
	}

	return m
}

func GroupStatsBy[T any, K comparable, V Number](w []T, key func(T) K, value func(T) V) map[K]GroupStats[V] {
	m := make(map[K]GroupStats[V])
	for _, x := range w {
		k := key(x)
		v := value(x)
		stats, ok := m[k]
		if !ok {
			m[k] = GroupStats[V]{
				Count: 1,
				Sum:   v,
				Min:   v,
				Max:   v,
				Avg:   float64(v),
			}
			continue
		}

		stats.Count++
		stats.Sum += v
		if v < stats.Min {
			stats.Min = v
		}
		if v > stats.Max {
			stats.Max = v
		}
		stats.Avg = float64(stats.Sum) / float64(stats.Count)
		m[k] = stats
	}

	return m
}

func DistinctBy[T any, K comparable](w []T, key func(T) K) []T {
	result := make([]T, 0, len(w))
	seen := make(map[K]struct{}, len(w))
	for _, x := range w {
		k := key(x)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		result = append(result, x)
	}

	return result
}

func IndexBy[T any, K comparable](w []T, key func(T) K) map[K]T {
	m := make(map[K]T, len(w))
	for _, x := range w {
		m[key(x)] = x
	}

	return m
}

func Partition[T any](w []T, f func(T) bool) ([]T, []T) {
	matched := make([]T, 0, len(w))
	unmatched := make([]T, 0, len(w))
	for _, x := range w {
		if f(x) {
			matched = append(matched, x)
			continue
		}
		unmatched = append(unmatched, x)
	}

	return matched, unmatched
}

func SortBy[T any, K cmp.Ordered](w []T, key func(T) K) []T {
	result := slices.Clone(w)
	slices.SortStableFunc(result, func(a, b T) int {
		return cmp.Compare(key(a), key(b))
	})

	return result
}

func Take[T any](w []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n > len(w) {
		n = len(w)
	}

	return slices.Clone(w[:n])
}

func Skip[T any](w []T, n int) []T {
	if n <= 0 {
		return slices.Clone(w)
	}
	if n >= len(w) {
		return []T{}
	}

	return slices.Clone(w[n:])
}

func Chunk[T any](w []T, size int) [][]T {
	if size <= 0 {
		panic("chunk size must be positive")
	}

	chunks := make([][]T, 0, (len(w)+size-1)/size)
	for start := 0; start < len(w); start += size {
		end := start + size
		if end > len(w) {
			end = len(w)
		}
		chunks = append(chunks, slices.Clone(w[start:end]))
	}

	return chunks
}
