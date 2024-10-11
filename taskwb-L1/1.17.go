package main

import "fmt"

func BinarySearch(arr []int, target int) (int, bool) {
	l, r := 0, len(arr)-1

	for r >= l {
		m := l + (r-l)/2
		if arr[m] == target {
			return m, true
		}
		if arr[m] > target {
			r = m - 1
		}
		if arr[m] < target {
			l = m + 1
		}
	}
	return -1, false
}

func main() {
	arr := []int{1, 3, 7, 8, 9, 12, 15}
	target := 15
	fmt.Println(BinarySearch(arr, target))

}
