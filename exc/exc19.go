package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Scan(&str)
	ru := []rune(str)

	for i, j := 0, len(ru) - 1; i < j; i, j = i + 1, j - 1 {
		ru[i], ru[j] = ru[j], ru[i]
	}

	fmt.Println(string(ru))
}