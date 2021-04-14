package lecture11

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutTimeAndExitManyGoroutines(_ *testing.T) {
	exitCh := make(chan struct{})

	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}

	for _, city := range cities {
		go RunForecast(city, exitCh)
	}

	// no something until we want to stop the forecasts
	time.Sleep(5 * time.Second)

	close(exitCh)

	// at this point it's possible that some of goroutines are NOT closed
	fmt.Println("all goroutines are closed")
	fmt.Println("DONE")
}
