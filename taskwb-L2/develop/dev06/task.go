package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Функция для фильтрации и вывода строк по колонкам
func cut(fields []int, delimiter string, separated bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, delimiter)

		if separated && len(columns) < 2 {
			continue // Пропускаем строки без разделителя
		}

		var selected []string
		for _, field := range fields {
			if field-1 < len(columns) {
				selected = append(selected, columns[field-1])
			}
		}

		if len(selected) > 0 {
			fmt.Println(strings.Join(selected, delimiter))
		}
	}
}

func parseFields(fields string) ([]int, error) {
	parts := strings.Split(fields, ",")
	var result []int

	for _, part := range parts {
		var field int
		_, err := fmt.Sscanf(part, "%d", &field)
		if err != nil {
			return nil, fmt.Errorf("invalid field: %s", part)
		}
		result = append(result, field)
	}

	return result, nil
}

func main() {
	// Определение флагов
	fieldsFlag := flag.String("f", "", "Fields to select (e.g., 1,2,3)")
	delimiterFlag := flag.String("d", "\t", "Delimiter to use (default: TAB)")
	separatedFlag := flag.Bool("s", false, "Only print lines with the delimiter")
	flag.Parse()

	// Проверка наличия флага -f
	if *fieldsFlag == "" {
		fmt.Println("Usage: cut -f fields [-d delimiter] [-s]")
		return
	}

	// Разбор колонок для выбора
	fields, err := parseFields(*fieldsFlag)
	if err != nil {
		fmt.Printf("Error parsing fields: %v\n", err)
		return
	}

	// Выполнение утилиты cut
	cut(fields, *delimiterFlag, *separatedFlag)
}
