package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var i int
	fmt.Print("i=")
	fmt.Scanln(&i)

	slices.Delete(arr, i, i+1)
	fmt.Println(arr)

	tmp := arr[:i]
	tmp = append(tmp, arr[i+1:]...)
	// tmp = append(tmp[:i], tmp[i+1:]...)
	fmt.Println(tmp)
}
