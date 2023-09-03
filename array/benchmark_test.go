package array

import (
	"math/rand"
)

//func BenchmarkFilter1(b *testing.B) {
//	data := make([]int, 500_000)
//	for i := range data {
//		data[i] = rand.Intn(100)
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Filter(data, func(x int) bool { return x%2 == 0 })
//	}
//}
//
//func BenchmarkFilter2(b *testing.B) {
//	data := make([]int, 500_000)
//	for i := range data {
//		data[i] = rand.Intn(100)
//	}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Filter0Loc(data, func(x int) bool { return x%2 == 0 })
//	}
//}
var data = make([]int, 500_000)

func init() {
	for i := range data {
		data[i] = rand.Intn(1000)
	}
}

//func BenchmarkFind(b *testing.B) {
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Find(data, func(x int) bool { return x%2 == 0 && x%3 == 0 && x%5 == 0 })
//	}
//}
//
//func BenchmarkFind2(b *testing.B) {
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Find2(data, func(x int) bool { return x%2 == 0 && x%3 == 0 && x%5 == 0 })
//	}
//}
//
//func BenchmarkMap(b *testing.B) {
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Map(data, func(x int) int { return x * 2 })
//	}
//}
//
//func BenchmarkMap2(b *testing.B) {
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		Map2(data, func(x int) int { return x * 2 })
//	}
//}
