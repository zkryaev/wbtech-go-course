package main

import "fmt"

func WhatsType(input interface{}) {
	switch input.(type) {
	case string:
		fmt.Println("<- string")
	case bool:
		fmt.Println("<- bool")
	case chan int:
		fmt.Println("<- chan int")
	case int:
		fmt.Println("<- int")
	}
}

func main() {
	n := 5
	str := "Hello"
	is_bool := true
	ch := make(chan int, 1)

	WhatsType(n)       // int
	WhatsType(ch)      // chan int
	WhatsType(str)     // string
	WhatsType(is_bool) // bool
}
