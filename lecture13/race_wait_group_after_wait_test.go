package lecture11

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroupAfterWait(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := new(sync.WaitGroup)
	wg.Add(len(cities))

	// we want no more than 10 forecasts
	res := make(chan int, 10)
	go func() {
		runForecastsWithWG(ctx, wg, cities, res)
	}()

	// for test only
	time.Sleep(1 * time.Second)

	// additional tasks
	go func() {
		for i := 0; i < 100; i++ {
			i++
			wg.Add(1)
			// "Do something else"
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}
	}()

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

func runForecastsWithWG(ctx context.Context, wg *sync.WaitGroup, cities []string, res chan int) chan int {
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
	return res
}
