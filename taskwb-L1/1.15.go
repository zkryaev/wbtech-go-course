package main

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "a"
	}
	return v
}

var justString string

func someFunc() {
	// Проблема:
	// 'justString' указывает на тот же underlying массив, что и 'v'
	// Т.к justString глобальная переменная, то garbage collector не очищает эту область, т.к
	// на нее указывает justString
	//
	// Решение:
	// 1. Сделать justString локальной переменной и возвращать ее из функции
	//
	// 2. Выделить новую область памяти
	v := createHugeString(1 << 10)
	//
	// Используя приведение к типу []byte/[]rune создает новый массив с другим типом,
	// который мы далее приводим к string
	justString = string([]byte(v[:100]))
	// Это работает потому, что строка иммутабельна, а массив байт - да.
	// Поэтому Go выполняет эту работу за нас и выделяет память под слайс байтов и копирует содержимое строки.
}

func main() {
	someFunc()
}
