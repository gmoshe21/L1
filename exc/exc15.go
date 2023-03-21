package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := strings.Repeat("годА", 1<<10)
	run := []rune(v)[:100]
	fmt.Println(string(run))
}

func main() {
	someFunc()
}