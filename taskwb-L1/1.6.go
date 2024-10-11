package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// С использованием:
	//
	// 1. Контекстов
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled")
			return
		}
	}(ctx)
	cancel()

	// 2. Каналов
	done := make(chan bool, 1)
	go func(done <-chan bool) {
		select {
		case <-done:
			fmt.Println("Channel is read")
			return
		}
	}(done)
	done <- true
	close(done)

	// 3. WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
	}()
	wg.Wait()
	fmt.Println("Counter of WaitGroup = 0")
}
