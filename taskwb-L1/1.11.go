package main

import "fmt"

func main() {
	set1 := []int{1, 3, 4, 5, 7, 9}
	set2 := []int{2, 3, 5, 7, 10}
	mp := make(map[int]struct{})

	for _, v := range set2 {
		mp[v] = struct{}{}
	}

	var ans []int
	for _, v := range set1 {
		if _, ok := mp[v]; ok {
			ans = append(ans, v)
		}
	}
	fmt.Println(ans)

}
