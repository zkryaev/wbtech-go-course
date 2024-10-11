package main

import (
	"fmt"
	"strings"
)

func IsAllCharsUnique(s string) bool {
	set := make(map[rune]struct{})
	s = strings.ToLower(s)
	r := []rune(s)
	for _, sym := range r {
		if _, ok := set[sym]; !ok {
			set[sym] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range strs {
		fmt.Printf("%s - %t\n", str, IsAllCharsUnique(str))
	}
}
