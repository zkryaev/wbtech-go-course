package main

import (
	"fmt"
	"time"
)

func MineSleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	MineSleep(3 * time.Second)
	fmt.Println("Wake up samurai =_=")
}
