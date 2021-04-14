package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestDoSomethingAction(_ *testing.T) {
	city := "London"
	res := make(chan int, 100)
	GetWeatherInfiniteWithRes(city, res)

	fmt.Println("Started")

	for temp := range res {
		fmt.Printf("Temperature in %s is %d C at %v\n", city, temp, time.Now())
	}

	fmt.Println("DONE")
}
