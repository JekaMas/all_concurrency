package lecture11

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestTimeoutTimeOrExit(_ *testing.T) {
	const Timeout = time.Second

	city := "London"
	res := GetWeatherInfinite(city)
	exitCh := make(chan struct{})
	counter := new(int32)

	go func(c *int32) {
		// stop forecast if we get 3 or more
		for {
			if atomic.LoadInt32(c) >= 3 {
				close(exitCh)
				return
			}
		}
	}(counter)

loop:
	for {
		select {
		case temperature := <-res:
			atomic.AddInt32(counter, 1)

			fmt.Printf("Temperature in %s is %d C at %v\n", city, temperature, time.Now())
		case <-exitCh:
			fmt.Println("Exit by command. Forecast for", city)
			return
		case <-time.After(Timeout):
			fmt.Println("Exit by timeout. Forecast for", city)
			break loop
		}
	}
}
