package slice

import "testing"

var N = 300000

func BenchmarkSlice(b *testing.B) {
	b.Run("cap = 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var arr []int
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = 0.1 * len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, N/10)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = 0.5 * len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, N/2)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, N)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("cap = 10 * len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			arr := make([]int, 0, 10*N)
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
	b.Run("reuse", func(b *testing.B) {
		var arr []int
		for i := 0; i < b.N; i++ {
			arr = arr[:0]
			for i := 0; i < N; i++ {
				arr = append(arr, i)
			}
		}
	})
}

/*
goos: darwin
goarch: amd64
pkg: github.com/lzxbill7/mygo/benchmark/slice
BenchmarkSlice/cap_=_0-4         	     500	   2979148 ns/op
BenchmarkSlice/cap_=_0.1_*_len-4 	     500	   2881691 ns/op
BenchmarkSlice/cap_=_0.5_*_len-4 	     500	   2087178 ns/op
BenchmarkSlice/cap_=_len-4       	    2000	    745361 ns/op
BenchmarkSlice/cap_=_10_*_len-4  	     300	   3985154 ns/op
BenchmarkSlice/reuse-4           	    5000	    254029 ns/op
PASS
*/