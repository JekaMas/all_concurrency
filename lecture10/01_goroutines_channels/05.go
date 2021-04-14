package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()

	fmt.Println(2)
	ch <- 1000
}
