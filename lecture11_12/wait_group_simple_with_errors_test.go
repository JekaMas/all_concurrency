package lecture11

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWaitGroupWithErrorsSimple(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(cities))

	// we want no more than 10 forecasts
	res := make(chan int, 10)
	errCh := make(chan error, len(cities))

	go func() {
		for _, city := range cities {
			go func(ctx context.Context, city string, errCh chan error) {
				defer wg.Done() // on return or panic

				// timeout
				timer := time.NewTimer(time.Duration(100*rand.Intn(50))*time.Millisecond)
				defer timer.Stop()

				for {
					select {
					case <-ctx.Done():
						fmt.Printf("exit %q goroutine: %v\n", city, ctx.Err())
						return
					case <-timer.C:
						errCh <- fmt.Errorf("the forecast for %s is node with error: %q", city, "timeout")
						return
					case res <- GetWeatherByTime("", time.Now()):
						// nothing to do
					}
				}
			}(ctx, city, errCh)
		}
	}()

	// for test only
	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()

	close(res)
	close(errCh)

	// check for errors
	for err := range errCh {
		fmt.Println("error", err)
		// log all the errors or exit
	}

	i := 0
	for v := range res {
		fmt.Printf("%d: Temperature  is %d C at %v\n", i, v, time.Now())
		i++
	}

	fmt.Println("DONE")
}
