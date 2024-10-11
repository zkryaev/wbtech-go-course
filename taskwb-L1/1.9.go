package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	input := make(chan int, 1)
	output := make(chan int, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case x := <-input:
				output <- x * x
			case <-ctx.Done():
				return
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case input <- rand.Intn(10000 + 1):
				time.Sleep(1 * time.Second)
			case <-ctx.Done():
				return
			}
		}
	}()
	for {
		select {
		case xx := <-output:
			fmt.Println(xx)
		case <-ctx.Done():
			wg.Wait()
			close(output)
			close(input)
			return
		}
	}
}
