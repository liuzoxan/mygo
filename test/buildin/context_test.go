package buildin

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("goroutine2 ctx done")
				return
			default:
				log.Println("goroutine2")
				time.Sleep(time.Millisecond * 100)
			}
		}
	}(ctx)
	time.Sleep(time.Millisecond * 500)
	log.Println("goroutine1 cancel")
	cancel()
	time.Sleep(time.Millisecond * 300)
}

func TestContextCancel2(t *testing.T) {
	var f = func(ctx context.Context, name string) {
		for {
			select {
			case <- ctx.Done():
				log.Println(name, " done")
				return
			default:
				log.Println(name)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx, "routine1")
	go f(ctx, "routine2")
	time.Sleep(time.Millisecond * 500)
	cancel()
	time.Sleep(time.Millisecond * 300)
}

func TestContextError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				return
			}
		}
	}(ctx)
	time.Sleep(time.Millisecond * 500)
	cancel()
	log.Printf("context error: %v", ctx.Err())
}

func TestContextDeadline(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond * 300))
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				log.Println("routine done,	", ctx.Err())
				return
			}
		}
	}(ctx)
	time.Sleep(time.Millisecond * 500)
}

func TestContextTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond * 300)
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				log.Printf("routine done, %v", ctx.Err())
				return
			}
		}
	}(ctx)
	time.Sleep(time.Millisecond * 500)
}

func TestContextValue(t *testing.T) {
	ctx, _ := context.WithTimeout(context.WithValue(context.Background(), "name", "leon"), time.Millisecond * 300)
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				log.Printf("routine done, %v %v", ctx.Value("name"), ctx.Err())
				return
			}
		}
	}(ctx)
	time.Sleep(time.Millisecond * 500)
}
