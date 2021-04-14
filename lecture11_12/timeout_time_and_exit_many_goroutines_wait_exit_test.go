package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutTimeAndExitManyWaitExitGoroutines(_ *testing.T) {
	exitCh := make(chan struct{})

	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	done := make(chan struct{}, len(cities))

	for _, city := range cities{
		go RunForecastWithDone(city, exitCh, done)
	}

	// no something until we want to stop the forecasts
	time.Sleep(5*time.Second)

	close(exitCh)

	i:= 0
	for range done {
		i++

		if i == len(cities) {
			break
		}
	}

	fmt.Println("all goroutines are closed")
	fmt.Println("DONE")
}

