package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagramSets(&words)

	for key, group := range *anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}

func findAnagramSets(words *[]string) *map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range *words {
		wordLower := strings.ToLower(word)
		sortedWord := sortString(wordLower)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], wordLower)
	}

	result := make(map[string][]string)
	for _, group := range anagramMap {
		if len(group) > 1 {
			group = uniqueAndSort(group)
			result[group[0]] = group
		}
	}

	return &result
}

func sortString(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func uniqueAndSort(words []string) []string {
	wordSet := make(map[string]struct{})
	for _, word := range words {
		wordSet[word] = struct{}{}
	}

	uniqueWords := make([]string, 0, len(wordSet))
	for word := range wordSet {
		uniqueWords = append(uniqueWords, word)
	}

	sort.Strings(uniqueWords)
	return uniqueWords
}
