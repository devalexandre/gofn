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
