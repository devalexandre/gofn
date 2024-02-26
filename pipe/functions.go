package pipe

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/devalexandre/gofn/array"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Filter adapts the filter function for pipeline use.
func Filter[T any](f func(T) bool) func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		if len(a) == 0 {
			return nil, fmt.Errorf("empty slice")
		}
		return array.Filter(a, f), nil
	}
}

// Map adapts the map function for pipeline use.
func Map[T any, U any](f func(T) U) func([]T) ([]U, error) {
	return func(a []T) ([]U, error) {
		return array.Map(a, f), nil
	}
}

// Reduce adapts the reduce function for pipeline use.
func Reduce[T any](f func(T, T) T) func([]T) (T, error) {
	return func(a []T) (T, error) {
		if len(a) == 0 {
			var zero T // Retorna um valor zero do tipo T se o slice estiver vazio.
			return zero, nil
		}
		return array.Reduce(a, f), nil
	}
}

// Sum adapts the sum function for pipeline use.
func Sum[T Number]() func([]T) (T, error) {
	return func(a []T) (T, error) {
		return array.Sum(a), nil
	}
}

// Product adapts the product function for pipeline use.
func Product[T Number]() func([]T) (T, error) {
	return func(a []T) (T, error) {
		return array.Product(a), nil
	}
}

func Min[T Number]() func([]T) (T, error) {
	return func(a []T) (T, error) {
		if len(a) == 0 {
			return *new(T), fmt.Errorf("slice is empty, cannot determine Min")
		}
		return array.Min(a), nil
	}
}

func Max[T Number]() func([]T) (T, error) {
	return func(a []T) (T, error) {
		if len(a) == 0 {
			return *new(T), fmt.Errorf("slice is empty, cannot determine Max")
		}
		return array.Max(a), nil
	}
}

// Unique adapts the unique function for pipeline use.
func Unique[T comparable]() func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		return array.Unique(a), nil
	}
}

// Reverse adapts the reverse function for pipeline use.
func Reverse[T any]() func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		return array.Reverse(a), nil
	}
}

// Shuffle adapts the shuffle function for pipeline use.
func Shuffle[T any]() func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		rand.Seed(time.Now().UnixNano())
		return array.Shuffle(a), nil
	}
}

// Join adapts the join function for pipeline use.
func Join[T any](sep string) func([]T) (string, error) {
	return func(a []T) (string, error) {
		return array.Join(a, sep), nil
	}
}

// Contains adapts the contains function for pipeline use.
func Contains[T comparable](element T) func([]T) (bool, error) {
	return func(a []T) (bool, error) {
		return array.Contains(a, element), nil
	}
}

// IndexOf adapts the indexOf function for pipeline use.
func IndexOf[T comparable](element T) func([]T) (int, error) {
	return func(a []T) (int, error) {
		return array.IndexOf(a, element), nil
	}
}

// ForEach adapts the forEach function for pipeline use.
// Nota: Embora ForEach não retorne um valor, para manter a assinatura consistente,
// retornaremos a própria slice e nil para o erro.
func ForEach[T any](action func(T)) func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		array.ForEach(a, action)
		return a, nil // Retorna a mesma slice para manter a cadeia
	}
}

// Pop adapts the pop function for pipeline use.
func Pop[T any]() func([]T) (T, []T, error) {
	return func(a []T) (T, []T, error) {
		if len(a) == 0 {
			var zero T
			return zero, nil, fmt.Errorf("empty slice, cannot Pop")
		}
		value, remaining := array.Pop(a)
		return value, remaining, nil
	}
}

// Push adapts the push function for pipeline use.
func Push[T any](elements ...T) func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		return array.Push(a, elements...), nil
	}
}

// Shift adapts the shift function for pipeline use.
func Shift[T any]() func([]T) (T, []T, error) {
	return func(a []T) (T, []T, error) {
		if len(a) == 0 {
			var zero T
			return zero, nil, fmt.Errorf("empty slice, cannot Shift")
		}
		value, remaining := array.Shift(a)
		return value, remaining, nil
	}
}

// Sort adapts the sort function for pipeline use.
// Sort adapts the sort function for pipeline use, requiring a comparison function.
func Sort[T any](less func(i, j T) bool) func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		if len(a) == 0 {
			return nil, fmt.Errorf("empty slice, cannot Sort")
		}
		b := make([]T, len(a))
		copy(b, a)
		// Ajusta a chamada para sort.Slice para usar uma função de comparação que opera sobre os elementos da slice.
		sort.Slice(b, func(i, j int) bool {
			return less(b[i], b[j])
		})
		return b, nil
	}
}

// Unshift adapts the unshift function for pipeline use.
func Unshift[T any](elements ...T) func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		return array.Unshift(a, elements...), nil
	}
}

func GroupBy[T any, K comparable](f func(T) K) func([]T) ([]T, error) {
	return func(a []T) ([]T, error) {
		if len(a) == 0 {
			return nil, fmt.Errorf("the input slice is empty")
		}
		seen := make(map[K]bool)
		var result []T
		for _, item := range a {
			key := f(item)
			if !seen[key] {
				seen[key] = true
				result = append(result, item)
			}
		}
		return result, nil
	}
}
