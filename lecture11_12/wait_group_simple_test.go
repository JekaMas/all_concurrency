package lecture11

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroupSimple(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(cities))

	// we want no more than 10 forecasts
	res := make(chan int, 10)
	go func() {
		for _, city := range cities {
			go func(ctx context.Context, city string) {
				defer wg.Done() // on return or panic

				for {
					select {
					case <-ctx.Done():
						fmt.Printf("exit %q goroutine\n", city)
						return
					case res <- GetWeatherByTime("", time.Now()):
						// nothing to do
					}
				}
			}(ctx, city)
		}
	}()

	// for test only
	time.Sleep(3 * time.Second)

	cancel()

	wg.Wait()
	close(res)

	i := 0
	for v := range res {
		fmt.Printf("%d: Temperature  is %d C at %v\n", i, v, time.Now())
		i++
	}

	fmt.Println("DONE")
}
