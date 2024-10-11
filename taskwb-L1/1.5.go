package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var seconds int
	fmt.Print("Введите кол-во секунд:")
	fmt.Scanln(&seconds)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)

	msgs := make(chan int, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case val := <-msgs:
				fmt.Println(val) // чтобы отслеживать процесс работы
			}
		}
	}(ctx)

	for {
		select {
		case msgs <- rand.Int() % 10000:
		case <-ctx.Done():
			close(msgs)
			wg.Wait()
			fmt.Println("Timeout :)")
			return
		}
	}
}
