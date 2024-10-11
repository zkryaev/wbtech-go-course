package main

import "fmt"

func partition(arr []int, l, r int) int {
	pivot := arr[r]

	swap := l - 1
	for i := l; i <= r-1; i++ {
		if arr[i] <= pivot {
			swap++
			arr[i], arr[swap] = arr[swap], arr[i]
		}
	}
	arr[swap+1], arr[r] = arr[r], arr[swap+1]
	return swap + 1
}

func quicksort(arr []int, l, r int) {
	if l < r {
		m := partition(arr, l, r)
		quicksort(arr, l, m-1)
		quicksort(arr, m+1, r)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println(arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
