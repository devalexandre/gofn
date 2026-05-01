package array

/* Chain
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Array(a).Reverse().Shuffle().Value()
*   fmt.Println(b) // [5 4 3 2 1]

 */

type Array[T any] []T

func (a Array[T]) Length() int {
	return len(a)
}

func (a Array[T]) Filter(f func(T) bool) Array[T] {
	return Filter(a, f)
}

func (a Array[T]) Find(f func(T) bool) T {
	return Find(a, f)
}

func (a Array[T]) Map(f func(T) T) Array[T] {
	return Map(a, f)
}

func (a Array[T]) FlatMap(f func(T) []T) Array[T] {
	return FlatMap(a, f)
}

func (a Array[T]) Reduce(f func(T, T) T) T {
	return Reduce(a, f)
}

func (a Array[T]) Reverse() Array[T] {
	return Reverse(a)
}

func (a Array[T]) Any(f func(T) bool) bool {
	return Any(a, f)
}

func (a Array[T]) Some(f func(T) bool) bool {
	return Some(a, f)
}

func (a Array[T]) Every(f func(T) bool) bool {
	return Every(a, f)
}

func (a Array[T]) Shuffle() Array[T] {
	return Shuffle(a)
}

func (a Array[T]) Sort(f func(i, j int) bool) Array[T] {
	return Sort(a, f)
}

func (a Array[T]) GroupBy(f func(T) string) map[string][]T {
	return GroupBy(a, f)
}

func (a Array[T]) GroupSumBy(key func(T) string, value func(T) float64) map[string]float64 {
	return GroupSumBy(a, key, value)
}

func (a Array[T]) GroupSumByWhere(where func(T) bool, key func(T) string, value func(T) float64) map[string]float64 {
	return GroupSumByWhere(a, where, key, value)
}

func (a Array[T]) GroupCountBy(key func(T) string) map[string]int {
	return GroupCountBy(a, key)
}

func (a Array[T]) GroupStatsBy(key func(T) string, value func(T) float64) map[string]GroupStats[float64] {
	return GroupStatsBy(a, key, value)
}

func (a Array[T]) DistinctBy(key func(T) string) Array[T] {
	return DistinctBy(a, key)
}

func (a Array[T]) IndexBy(key func(T) string) map[string]T {
	return IndexBy(a, key)
}

func (a Array[T]) Partition(f func(T) bool) (Array[T], Array[T]) {
	matched, unmatched := Partition(a, f)
	return matched, unmatched
}

func (a Array[T]) SortByString(key func(T) string) Array[T] {
	return SortBy(a, key)
}

func (a Array[T]) SortByFloat64(key func(T) float64) Array[T] {
	return SortBy(a, key)
}

func (a Array[T]) Take(n int) Array[T] {
	return Take(a, n)
}

func (a Array[T]) Skip(n int) Array[T] {
	return Skip(a, n)
}

func (a Array[T]) Chunk(size int) []Array[T] {
	chunks := Chunk(a, size)
	result := make([]Array[T], len(chunks))
	for i, chunk := range chunks {
		result[i] = chunk
	}

	return result
}
