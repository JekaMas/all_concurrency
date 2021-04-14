package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestDoSomethingBuffered1Action(_ *testing.T) {
	city := "London"
	const length = 5
	res := make(chan int, length)
	GetWeatherInfiniteWithRes(city, res)

	fmt.Println("Started")

	buffer := make([]string, length)
	i := 0
	for {
		for temp := range res {
			buffer[i] = fmt.Sprintf("Temperature in %s is %d C at %v\n", city, temp, time.Now())
			i++

			if i == length {
				// todo output
				fmt.Print(buffer)
				i = 0
			}
		}

		// without any sleep or delay it'll eat CPU
	}

	fmt.Println("DONE")
}
