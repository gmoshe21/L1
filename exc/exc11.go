package main

import "fmt"

func main() {
	data1 := map[string]int{"Harry": 1, "Oliver": 1, "Jack": 1, "Charlie": 1,}
	data2 := map[string]int{"Riley ": 1, "Alfie ": 1, "Oliver": 1, "Jack": 1}

	var result []string

	for valData1 := range data1 {
		check := data2[valData1]
		if check != 0 {
			result = append(result, valData1)
		}
	}
	fmt.Println(result)
}