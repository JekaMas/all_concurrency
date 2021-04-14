package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	fmt.Println(2)
	ch <- 1000
}
