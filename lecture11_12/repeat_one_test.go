package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestRepeatOneByTicker(_ *testing.T) {
	const Timeout = 5*time.Second

	city := "London"
	res := GetWeatherInfinite(city)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	fmt.Println("Started")

	select {
	case <-ticker.C:
		fmt.Printf("Temperature in %s is %d C at %v\n", city, <-res, time.Now())
	case <-time.After(Timeout):
		fmt.Println("Exit by timeout")
	}

	fmt.Println("DONE")
}
