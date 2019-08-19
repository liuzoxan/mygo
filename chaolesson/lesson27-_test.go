package chaolesson

import (
	"log"
	"sync"
	"testing"
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
