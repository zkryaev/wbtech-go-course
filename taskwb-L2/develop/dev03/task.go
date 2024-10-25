package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	column := flag.Int("k", 1, "указать колонку для сортировки (начиная с 1)")
	numeric := flag.Bool("n", false, "сортировка по числовому значению")
	reverse := flag.Bool("r", false, "обратный порядок сортировки")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Использование: sortutil [опции] <входной_файл> <выходной_файл>")
		os.Exit(1)
	}

	inputFile := flag.Args()[0]
	outputFile := flag.Args()[1]

	// Чтение строк из файла
	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// Удаление дубликатов, если указан флаг -u
	if *unique {
		lines = removeDuplicates(lines)
	}

	// Сортировка строк
	sort.SliceStable(lines, func(i, j int) bool {
		col1 := getColumn(lines[i], *column)
		col2 := getColumn(lines[j], *column)

		// Если числовая сортировка
		if *numeric {
			num1, err1 := strconv.ParseFloat(col1, 64)
			num2, err2 := strconv.ParseFloat(col2, 64)

			if err1 == nil && err2 == nil {
				if *reverse {
					return num1 > num2
				}
				return num1 < num2
			}
		}

		// Стандартное сравнение строк
		if *reverse {
			return col1 > col2
		}
		return col1 < col2
	})

	// Запись результата в выходной файл
	if err := writeLines(outputFile, lines); err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		os.Exit(1)
	}
}

// Функция для чтения строк из файла
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Функция для записи строк в файл
func writeLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

// Функция для получения конкретной колонки из строки
func getColumn(line string, column int) string {
	fields := strings.Fields(line)
	if column > 0 && column <= len(fields) {
		return fields[column-1]
	}
	return line
}

// Функция для удаления дубликатов
func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	var uniqueLines []string
	for _, line := range lines {
		if !seen[line] {
			uniqueLines = append(uniqueLines, line)
			seen[line] = true
		}
	}
	return uniqueLines
}
