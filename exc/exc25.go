package main

import (
	"fmt"
	"time"
)

func Sleep(sl int64) {
	start := time.Now()
	end := start.Add(time.Millisecond*time.Duration(sl))

	for time.Now().UnixNano() < end.UnixNano() {

	}
}

func main() {
	var sec int64
	fmt.Scan(&sec)
	
	Sleep(sec)
}