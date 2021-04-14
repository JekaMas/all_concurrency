package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutTimeAndExitManyByOneGoroutines(_ *testing.T) {
	exitCh := make(chan struct{}, 1)

	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}

	for _, city := range cities {
		go RunForecast(city, exitCh)
	}

	// no something until we want to stop the forecasts
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	i := 0
	for range ticker.C {
		// close one forecast randomly
		exitCh <- struct{}{}
		i++

		if i == len(cities) {
			fmt.Println("all goroutines are closed")
			break
		}
	}

	fmt.Println("DONE")
}
