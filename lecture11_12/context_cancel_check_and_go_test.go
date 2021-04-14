package lecture11

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelCheckAndGo(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan int, 1)
	done := make(chan struct{}, len(cities))

	// get the first forecast and stop
	go func() {
		for _, city := range cities {
			go func(ctx context.Context, city string) {
				// early stop
				select {
				case <-ctx.Done():
					done <- struct{}{}
					fmt.Println("Early stop! The forecast of", city)
					return
				default:
				}

				// even longer processing
				GetWeatherWithContextAndRes(ctx, city, res, done)
			}(ctx, city)

			// processing
			time.Sleep(time.Second)
		}
	}()

	for v := range res {
		cancel()
		fmt.Printf("Temperature  is %d C at %v\n", v, time.Now())
		break
	}

	i := 0
	for range done {
		i++
		if i == len(cities) {
			break
		}
	}

	fmt.Println("DONE")
}
