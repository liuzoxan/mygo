package chaolesson

import (
	"fmt"
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
	count := 0
	for i:=0; i<5000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second)
	t.Log(count)
}
