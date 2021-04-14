package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutTimeOnce(_ *testing.T) {
	const Timeout = time.Second

	city := "London"
	res := GetWeatherInfinite(city)

	select {
	case temperature := <-res:
		fmt.Printf("Temperature in %s is %d C at %v\n", city, temperature, time.Now())
	case <-time.After(Timeout):
		fmt.Println("Exit by timeout")
	}

	fmt.Println("DONE")
}
