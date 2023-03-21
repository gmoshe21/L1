package main

import "fmt"


func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	data := make(map[string]int)

	for _, val := range str {
		check := data[val]
		if check == 0 {
			data[val] = 1
		} else {
			data[val] = data[val] + 1
		}
	}
	fmt.Println(data)
}