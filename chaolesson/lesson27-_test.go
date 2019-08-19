package chaolesson

import (
	"log"
	"sync"
	"testing"
	"time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

}

func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if ret, OK := <-ch; OK {
				log.Println(ret)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

func TestData(t *testing.T) {
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	wg.Wait()
}

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {
	ch := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancelled(ch) {
					break
				}
				time.Sleep(time.Microsecond * 100)
			}
			t.Logf("%d cancelled", i)
		}(i, ch)
	}
	cancel_2(ch)
	time.Sleep(time.Second)
}
