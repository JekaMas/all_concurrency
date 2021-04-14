package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println(2)
	}()

	fmt.Println(2)

	time.Sleep(time.Millisecond)
}
