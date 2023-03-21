package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var i int

	fmt.Print("ind = ")
	fmt.Scan(&i)

	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println(nums)
}