package lecture11

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextCancel(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	done := make(chan struct{}, len(cities))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, city := range cities {
		go RunForecastWithContext(ctx, city, done)
	}

	// no something until we want to stop the forecasts
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	i := 0
	for range ticker.C {
		i++
		if i == 3 {
			cancel()
			fmt.Println("closing goroutines")
			break
		}
	}

	i = 0
	for range done {
		i++
		if i == len(cities) {
			break
		}
	}

	fmt.Println("DONE")
}
