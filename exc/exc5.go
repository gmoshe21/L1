package main

import (
	"fmt"
	"time"
)

var t int = 5

func run(ch chan string) {
	for {
		fmt.Println("run:", <-ch)
	}
}

func main() {
	start := time.Now()
	ch := make(chan string)
	
	var tmp string
	var duration time.Duration
	
	go run(ch)
	
	for {
		fmt.Scan(&tmp)
		ch <- tmp
		duration = time.Since(start)
		if duration.Seconds() > float64(t) {
			return 
		}
	}
}