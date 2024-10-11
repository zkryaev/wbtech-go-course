package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	TempSet := make(map[int][]float64)

	for _, v := range temps {
		key := int(v/10) * 10 // округление
		TempSet[key] = append(TempSet[key], v)
	}
	for k, v := range TempSet {
		fmt.Printf("%d:{", k)
		for i, t := range v {
			fmt.Printf("%.1f", t)
			if i != len(v)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("}")
	}
}
