package main

import "fmt"

// Строки для тестирования в unicode
// 福喜愛永愛
// Горутиновый

func main() {
	var s string
	fmt.Scanln(&s)
	rs := []rune(s)

	j := len(rs) - 1
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[j] = rs[j], rs[i]
		j--
	}
	fmt.Println(string(rs))
}
