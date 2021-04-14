package lecture11

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextTimeout(_ *testing.T) {
		cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
		done := make(chan struct{}, len(cities))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for _, city := range cities {
			go RunForecastWithContext(ctx, city, done)
		}

		// no something until we want to stop the forecasts
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		i := 0
		for range done {
			i++
			if i == len(cities) {
				break
			}
		}

		fmt.Println("DONE")
}
