package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func work(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	mainCh := make(chan int)

	if len(os.Args) < 2 {
		fmt.Println("Error")
	}

	workerCnt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		i := 0
		for {
			if i == 5000 {
				i = 0
			}
			mainCh <- i
			time.Sleep(time.Millisecond * 500)
			i++
		}
	}()

	for i := 0; i < workerCnt; i++ {
		go work(mainCh)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
	close(mainCh)
}
