package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println(1)
		ch <- struct{}{}
	}()

	fmt.Println(2)
	<-ch
}
