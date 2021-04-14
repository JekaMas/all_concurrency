package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutTimerTime(_ *testing.T) {
	const Timeout = time.Second

	city := "London"
	res := GetWeatherInfinite(city)
	timer := time.NewTimer(Timeout)
	defer timer.Stop()

loop:
	for {
		select {
		case temperature := <-res:
			fmt.Printf("Temperature in %s is %d C at %v\n", city, temperature, time.Now())
		case <-timer.C:
			fmt.Println("Exit by timeout")
			break loop
		}
	}

	fmt.Println("DONE")
}
