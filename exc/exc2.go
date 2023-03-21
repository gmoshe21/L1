package main

import (
	"fmt"
	"sync"
)

type exc2 struct {
	mass [5]int
	mu sync.Mutex
	wg sync.WaitGroup
}

func (d *exc2)run(ind int) {
	defer d.wg.Done()

	d.mu.Lock()
	fmt.Println(d.mass[ind] * d.mass[ind])
	d.mu.Unlock()
}

func main() {
	var data exc2
	data.mass = [5]int{2,4,6,8,10}
	for i := 0; i < len(data.mass); i++ {
		data.wg.Add(1)
		go data.run(i)
	}
	data.wg.Wait()
}