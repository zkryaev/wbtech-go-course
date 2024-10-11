package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func main() {
	cnt := &Counter{val: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(cnt.val)
}
