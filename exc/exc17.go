package main

import (
	"fmt"
)

func BinSearch(nums []int, val int) int {

	prev := 0
	next := len(nums) - 1

	for prev <= next {
		tmp := (prev + next) / 2
		if nums[tmp] < val {
			prev = tmp + 1
		} else {
			next = tmp - 1
		}
	}

	if prev == len(nums) || nums[prev] != val {
		return -1
	} else {
		return prev
	}

}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(BinSearch(nums, 10))

}