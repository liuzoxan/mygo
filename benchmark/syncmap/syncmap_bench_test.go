package syncmap

import (
	"math/rand"
	"sync"
	"testing"
)

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100000; i++ {
					hm.Set(rand.Intn(1000), i*i)
					hm.Set(rand.Intn(1000), i*i)
					hm.Del(rand.Intn(1000))
				}
				wg.Done()
			}()
		}
		for i := 0; i < 80; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100000; i++ {
					hm.Get(i)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncmap(b *testing.B) {
	b.Run("mutex syncmap 1", func(b *testing.B) {
		hm := NewMutexSyncMap(1, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("mutex syncmap 100", func(b *testing.B) {
		hm := NewMutexSyncMap(100, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("mutex syncmap 200", func(b *testing.B) {
		hm := NewMutexSyncMap(200, func(key interface{}) int {
			return key.(int)
		})
		benchmarkMap(b, hm)
	})

	b.Run("std syncmap", func(b *testing.B) {
		hm := NewStdSyncMap()
		benchmarkMap(b, hm)
	})
}

/*
goos: darwin
goarch: amd64
pkg: github.com/lzxbill7/mygo/benchmark/syncmap
BenchmarkSyncmap/mutex_syncmap_1-4         	       1	1124292778 ns/op
BenchmarkSyncmap/mutex_syncmap_100-4       	       3	 379571967 ns/op
BenchmarkSyncmap/mutex_syncmap_200-4       	       3	 375369513 ns/op
BenchmarkSyncmap/std_syncmap-4             	       5	 299662216 ns/op
PASS
*/