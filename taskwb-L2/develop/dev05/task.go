package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Опции для утилиты
type grepOptions struct {
	after      int  // -A
	before     int  // -B
	context    int  // -C
	count      bool // -c
	ignoreCase bool // -i
	invert     bool // -v
	fixed      bool // -F
	lineNum    bool // -n
}

// Функция для проверки совпадения строки с шаблоном
func matches(line, pattern string, opts *grepOptions) bool {
	if opts.ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if opts.fixed {
		return line == pattern
	}

	return strings.Contains(line, pattern)
}

// Функция для фильтрации строк с учетом опций
func grep(lines []string, pattern string, opts *grepOptions) {
	var matchLines []int

	for i, line := range lines {
		match := matches(line, pattern, opts)
		if opts.invert {
			match = !match
		}

		if match {
			matchLines = append(matchLines, i)
		}
	}

	// Печать количества совпадений
	if opts.count {
		fmt.Println(len(matchLines))
		return
	}

	// Печать совпадающих строк с учетом -A, -B, -C
	printLines(lines, matchLines, opts)
}

// Функция для вывода строк с учетом параметров контекста
func printLines(lines []string, matchLines []int, opts *grepOptions) {
	printed := make(map[int]bool)

	for _, lineIdx := range matchLines {
		start := max(0, lineIdx-opts.before-opts.context)
		end := min(len(lines)-1, lineIdx+opts.after+opts.context)

		for i := start; i <= end; i++ {
			if !printed[i] {
				if opts.lineNum {
					fmt.Printf("%d:", i+1)
				}
				fmt.Println(lines[i])
				printed[i] = true
			}
		}
	}
}

func main() {
	// Определение флагов
	opts := grepOptions{}
	flag.IntVar(&opts.after, "A", 0, "Print N lines after match")
	flag.IntVar(&opts.before, "B", 0, "Print N lines before match")
	flag.IntVar(&opts.context, "C", 0, "Print N lines before and after match")
	flag.BoolVar(&opts.count, "c", false, "Print the count of matching lines")
	flag.BoolVar(&opts.ignoreCase, "i", false, "Ignore case distinctions")
	flag.BoolVar(&opts.invert, "v", false, "Select non-matching lines")
	flag.BoolVar(&opts.fixed, "F", false, "Interpret pattern as a fixed string")
	flag.BoolVar(&opts.lineNum, "n", false, "Print line number with output lines")
	flag.Parse()

	// Получаем аргументы после флагов
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		return
	}

	pattern := args[0]
	var lines []string

	// Чтение строк из файла или stdin
	if len(args) > 1 {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	grep(lines, pattern, &opts)
}
