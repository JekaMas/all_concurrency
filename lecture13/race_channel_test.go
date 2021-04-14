package lecture11

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroupRaceChannel(_ *testing.T) {
	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := new(sync.WaitGroup)
	wg.Add(len(cities))

	// we want no more than 10 forecasts
	res := make(chan int, 10)
	go func() {
		res = runForecastsWithWG(ctx, wg, cities, res)
	}()

	// for test only
	time.Sleep(1 * time.Second)

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
