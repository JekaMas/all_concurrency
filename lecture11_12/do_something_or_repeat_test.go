package lecture11

import (
	"fmt"
	"testing"
	"time"
)

// сложный селект, который что-то делает, а если ничего не происходит, то Ticker
func TestDoSomethingOrRepeatAction(_ *testing.T) {
	city := "London"
	res := GetWeatherInfinite(city)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	fmt.Println("Started")

	for {
		select {
		case temp := <-res:
			fmt.Printf("Temperature in %s is %d C at %v\n", city, temp, time.Now())
		case <-ticker.C:
			fmt.Println("Nothing happens. Run GC, flush data to the disk etc")
		}
	}

	fmt.Println("DONE")
}
