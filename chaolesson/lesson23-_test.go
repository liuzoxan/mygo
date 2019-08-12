package chaolesson

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "done"
}

func otherService() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("something else done")
}

func asyncService() chan string {
	c := make(chan string)
	go func() {
		ret:= service()
		c <- ret
	}()
	return c
}

func TestAsyncService(t *testing.T) {
	c := asyncService()
	otherService()
	t.Log(<-c)
}