package lecture11

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestErrorWaitGroupWithErrorsSimple(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// we want no more than 10 forecasts
	res := make(chan int, 10)

	var wg errgroup.Group
	go func() {
		for _, city := range cities {
			city := city

			wg.Go(func() error {
				// timeout
				timer := time.NewTimer(time.Duration(100*rand.Intn(50)) * time.Millisecond)
				defer timer.Stop()

				for {
					select {
					case <-ctx.Done():
						fmt.Printf("exit %q goroutine: %v\n", city, ctx.Err())
						return nil
					case <-timer.C:
						return fmt.Errorf("the forecast for %s is node with error: %q", city, "timeout")
					case res <- GetWeatherByTime("", time.Now()):
						// nothing to do
					}
				}
			})
		}
	}()

	// for test only
	time.Sleep(3 * time.Second)

	cancel()

	err := wg.Wait()
	if err != nil {
		// check for errors
		fmt.Println("error", err)
		// log all the errors or exit
	}

	close(res)

	i := 0
	for v := range res {
		fmt.Printf("%d: Temperature  is %d C at %v\n", i, v, time.Now())
		i++
	}

	fmt.Println("DONE")
}
