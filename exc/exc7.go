package main

import (
	"fmt"
	"sync"
)

var m map[string]int
var mu sync.Mutex
var wg sync.WaitGroup

func add(key []string, val []int) {
	defer wg.Done()
	for i := 0; i < len(key); i++ {
		mu.Lock()
		m[key[i]] = val[i]
		mu.Unlock()
	}
}

func main() {
	m = make(map[string]int)

	key := [][]string{{"aa", "dd", "gg", "kk"},{"ae", "gd", "hg", "kf"}}
	val := [][]int{{1,2,3,4},{6,7,8,9}}
	
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go add(key[i], val[i])
	}
	wg.Wait()
	fmt.Println(m)
}