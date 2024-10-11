package main

import (
	"fmt"
	"math/big"
)

// Для проверки работы
// 1,048,576
// 1,500,000

func main() {
	var a, b string
	fmt.Print("A:")
	fmt.Scanln(&a)
	fmt.Print("B:")
	fmt.Scanln(&b)

	A := new(big.Int)
	B := new(big.Int)

	A.SetString(a, 10)
	B.SetString(b, 10)

	A.Add(A, B)
	fmt.Println("A + B = ", A, "= A")
	A.Sub(A, B)
	fmt.Println("A - B = ", A, "= A")
	A.Mul(A, B)
	fmt.Println("A * B = ", A, "= A")
	A.Div(A, B)
	fmt.Println("A / B = ", A, "= A")
}
