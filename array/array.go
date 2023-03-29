package array

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
