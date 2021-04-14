package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestDoSomethingBufferedAction(_ *testing.T) {
	city := "London"
	res := make(chan int, 5)
	GetWeatherInfiniteWithRes(city, res)

	fmt.Println("Started")

	for {
		if len(res) == cap(res) {
			for temp := range res {
				fmt.Printf("Temperature in %s is %d C at %v\n", city, temp, time.Now())
			}
		}

		// without any sleep or delay it'll eat CPU
	}

	fmt.Println("DONE")
}
