package main

import (
	"fmt"
	"sync"
	"time"
)

func rut1(quit <-chan struct{}) {
	for {
        select {
        case <-quit:
            return
        default:
            fmt.Println("rut1")
			time.Sleep(1 * time.Second)
        }
    }
}

func rut2(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("rut2")
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func rut3() {
	for {
		fmt.Println("rut3")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	quit := make(chan struct{})
	wg.Add(1)

	go rut1(quit)
	go rut2(&wg)
	go rut3()
	
	time.Sleep(1 * time.Second)
	close(quit)
	wg.Wait()
}