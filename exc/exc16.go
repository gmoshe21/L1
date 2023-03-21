package main

import (
	"fmt"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func rec(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = rec(arr, low, p-1)
		arr = rec(arr, p+1, high)
	}
	return arr
}

func quickSort(arr []int) []int {
	return rec(arr, 0, len(arr)-1)
}

func main() {
	nums := []int{4,2,6,8,3,4,6,8,1,4,0,9,7,4,2,4,6,89}

	fmt.Println(quickSort(nums))
}