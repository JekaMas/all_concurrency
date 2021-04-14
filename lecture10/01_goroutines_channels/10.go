package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("Done!")
	}()

	fmt.Println(2)
	ch <- 1000
	ch <- 1000
	ch <- 1000
	ch <- 1000
	ch <- 1000
	ch <- 1000
	ch <- 1000
	close(ch)

	time.Sleep(time.Second)
}
