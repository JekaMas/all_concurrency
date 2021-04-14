package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		time.Sleep(time.Second)

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println("Done!")
	}()

	fmt.Println(2)
	ch <- 1000
	ch <- 1100
	ch <- 1200
	ch <- 1300
	close(ch)

	fmt.Println(cap(ch), len(ch))
}
