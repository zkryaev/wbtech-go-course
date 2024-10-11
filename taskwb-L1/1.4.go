package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var n int
	fmt.Print("Число воркеров:")
	fmt.Scanln(&n)

	msgs := make(chan int, n) // канал из которого воркеры будут читать
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background()) // для завершения работы всех воркеров
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case num := <-msgs:
					time.Sleep(60 * time.Millisecond) // чтобы не спамить в stdout
					fmt.Println(num)
				}
			}
		}(ctx)
	}
	for {
		select {
		case msgs <- rand.Int():
			continue
		case <-exit:
			cancel()
			close(msgs)
			wg.Done()
			fmt.Println("Main done")
			return
		}
	}
}
