package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(123456789000222)
	b := big.NewInt(222000987654321)
	
	sum := big.NewInt(0).Add(a, b)
	diff := big.NewInt(0).Sub(a, b)
	mul := big.NewInt(0).Mul(a, b)
	div := big.NewInt(0).Div(a, b)

	fmt.Println("a + b", sum)
	fmt.Println("a - b", diff)
	fmt.Println("a * b", mul)
	fmt.Println("a / b", div)
}