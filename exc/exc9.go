package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	num := make([]int, 21, 21)
	for i := 0; i < 21; i++ {
		num[i] = i
	}
	go func(){
		for i := 0; i < len(num); i++ {
			ch1 <- num[i]
		}
		close(ch1)
	}()

	go func() {
		for val := range ch1 {
			ch2 <- val * val 
		}
		close(ch2)
	}()

	for val := range ch2 {
		fmt.Println(val)
	}
}