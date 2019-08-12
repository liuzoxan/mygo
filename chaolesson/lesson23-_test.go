package chaolesson

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestCounter(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			count++
			lock.Unlock()
			wg.Done()
		}()
	}
	// time.Sleep(time.Second * 5)
	wg.Wait()
	t.Log(count)
}
