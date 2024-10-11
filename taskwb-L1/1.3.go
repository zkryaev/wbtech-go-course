package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	
	pows := make(chan int, len(nums))
	pows <- 0
	wg := sync.WaitGroup{}
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			a := <-pows
			pows <- a + (n * n)
		}(nums[i])
	}
	wg.Wait()

	close(pows)
	for v := range pows {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
