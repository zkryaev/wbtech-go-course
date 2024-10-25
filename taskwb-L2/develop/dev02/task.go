package main

import (
	"errors"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}
	last := '0'
	stack := make([]rune, 0, 1)
	for _, r := range input {
		if unicode.IsDigit(r) {
			if last == '0' {
				return "", errors.New("Incorrect string")
			}
			num, _ := strconv.Atoi(string(r))
			for num-1 > 0 {
				stack = append(stack, last)
				num--
			}
			last = '0'
			continue
		} else {
			stack = append(stack, r)
			last = r
			continue
		}
	}
	return string(stack), nil
}

func main() {

}
