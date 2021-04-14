package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println(2)
	}()

	fmt.Println(2)
}
