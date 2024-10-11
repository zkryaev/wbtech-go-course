package main

import "fmt"

func main() {

	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range strs {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}
