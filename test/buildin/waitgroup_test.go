package buildin

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(time.Millisecond * 200)
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		wg.Done()
	}()
	wg.Wait()
}
