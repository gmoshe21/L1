package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

type Counter struct {
	count int
}

func (c *Counter) Inc() {
	mu.Lock()
	c.count++
	mu.Unlock()
}

func (c *Counter) Get() int{
	return c.count
}

func r1(num int, data *Counter, wg *sync.WaitGroup) {
	for i := 0; i < num; i++ {
		data.Inc()
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	var data Counter
	var wg sync.WaitGroup

	wg.Add(2)
	go r1(2, &data, &wg)
	go r1(1, &data, &wg)
	wg.Wait()

	fmt.Println(data.Get())
}