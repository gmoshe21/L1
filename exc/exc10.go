package main

import (
	"fmt"
)

func main() {
	data := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32)

	for _, val := range data {
		tmp := int(val) / 10
		mass , ok := result[tmp * 10]
		if ok == true {
			result[tmp * 10] = append(mass, val)
		} else {
			result[tmp * 10] = []float32{val}
		}
	}
	fmt.Println(result)
}