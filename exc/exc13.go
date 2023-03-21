package main

import "fmt"

func main() {
	i, j := 5, 10
	i, j = j, i
	fmt.Println(i, j)
}