package main

import "fmt"

func main() {
	var num int64
	var i int
	var state int
	fmt.Print("Число:")
	fmt.Scanln(&num)
	fmt.Print("Номер бита:")
	fmt.Scanln(&i)
	fmt.Print("Установить в (1 или 0):")
	fmt.Scanln(&state)

	fmt.Printf("До: %b = %d\n", num, num)
	switch state {
	case 1:
		num |= 1 << i
	case 0:
		num &= ^(1 << i)
	}
	fmt.Printf("После: %b = %d\n", num, num)
}
