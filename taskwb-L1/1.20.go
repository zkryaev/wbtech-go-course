package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// считывание
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()

	// разбиваем на слова
	str := string(text)
	strs := strings.Split(str, " ")

	// меняем порядок
	j := len(strs) - 1
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[j] = strs[j], strs[i]
	}

	// конкатенируем
	var result strings.Builder
	for i := range strs {
		result.WriteString(strs[i])
		if i != len(strs)-1 {
			result.WriteString(" ")
		}
	}
	// результат
	fmt.Printf("%s\n", result.String())
}
