package array

/* Chain
* Example:
*   a := []int{1, 2, 3, 4, 5}
*   b := Array(a).Reverse().Shuffle().Value()
*   fmt.Println(b) // [5 4 3 2 1]

 */

type Array[T any] struct {
	value []T
}

func NewArray[T any](a []T) *Array[T] {
	return &Array[T]{value: a}
}

func (a *Array[T]) Value() []T {
	return a.value
}

func (a *Array[T]) Length() int {
	return len(a.value)
}

func (a *Array[T]) Filter(f func(T) bool) *Array[T] {
	return &Array[T]{value: Filter(a.value, f)}
}

func (a *Array[T]) Find(f func(T) bool) *T {
	return Find(&a.value, f)
}

func (a *Array[T]) Map(f func(T) T) *Array[T] {
	return &Array[T]{value: Map(a.value, f)}
}

func (a *Array[T]) FlatMap(f func(T) []T) *Array[T] {
	return &Array[T]{value: FlatMap(a.value, f)}
}

func (a *Array[T]) Reduce(f func(T, T) T) T {
	return Reduce(a.value, f)
}

func (a *Array[T]) Reverse() *Array[T] {
	return &Array[T]{value: Reverse(a.value)}
}

func (a *Array[T]) Any(f func(T) bool) bool {
	return Any(a.value, f)
}

func (a *Array[T]) Some(f func(T) bool) bool {
	return Some(a.value, f)
}

func (a *Array[T]) Every(f func(T) bool) bool {
	return Every(a.value, f)
}

func (a *Array[T]) Shuffle() *Array[T] {
	return &Array[T]{value: Shuffle(a.value)}
}

func (a *Array[T]) Sort(f func(T, T) bool) *Array[T] {

	for i := 0; i < len(a.value)-1; i++ {
		for j := i + 1; j < len(a.value); j++ {
			if f(a.value[i], a.value[j]) {
				a.value[i], a.value[j] = a.value[j], a.value[i]
			}
		}
	}

	return a
}
