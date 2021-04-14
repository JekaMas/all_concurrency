package lecture11

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// GetWeather gets a temperature forecast from remote API by the town name and time
func GetWeatherByTime(_ string, _ time.Time) int {
	res := int(rand.Int31n(45))
	switch {
	case res%2 == 0:
		time.Sleep(500 * time.Millisecond)
	case res%3 == 0:
		time.Sleep(800 * time.Millisecond)
	case res%5 == 0:
		time.Sleep(3 * time.Second)
	case res%10 == 0:
		time.Sleep(time.Minute)
	default:
		time.Sleep(300 * time.Millisecond)
	}
	return res
}

// GetWeather gets a forecast from remote API by the town name
func GetWeatherInfinite(_ string) chan int {
	res := make(chan int, 1)

	go func() {
		for {
			res <- GetWeatherByTime("", time.Now())
		}
	}()

	return res
}

func GetWeatherInfiniteWithRes(_ string, res chan int) {
	go func() {
		for {
			res <- GetWeatherByTime("", time.Now())
		}
	}()
}

func GetWeather(_ string) (chan int, chan struct{}) {
	res := make(chan int, 1)

	go func() {
		for {
			res <- GetWeatherByTime("", time.Now())
		}
	}()

	exitCh := make(chan struct{})
	return res, exitCh
}

func GetWeatherWithExit(town string, exitCh chan struct{}) chan int {
	res := make(chan int, 1)

	go func() {
		for {
			select {
			case <-exitCh:
				fmt.Printf("exit %q goroutine\n", town)
				return
			case res <- GetWeatherByTime("", time.Now()):
				// nothing to do
			}
		}
	}()

	return res
}

func RunForecast(city string, exitCh chan struct{}) {
	const Timeout = time.Second

	res := GetWeatherWithExit(city, exitCh)

loop:
	for {
		select {
		case temperature := <-res:
			fmt.Printf("Temperature in %s is %d C at %v\n", city, temperature, time.Now())
		case v, ok := <-exitCh:
			fmt.Println("Exit by command. Forecast for", city, v, ok)
			return
		case <-time.After(Timeout):
			fmt.Println("Exit by timeout. Forecast for", city)
			break loop
		}
	}
}

func RunForecastWithDone(city string, exitCh chan struct{}, done chan<- struct{}) {
	/*
		// if panic is possible
		defer func() {
			done <- struct{}{}
		}()
	*/

	const Timeout = time.Second

	res := GetWeatherWithExit(city, exitCh)

loop:
	for {
		select {
		case temperature := <-res:
			fmt.Printf("Temperature in %s is %d C at %v\n", city, temperature, time.Now())
		case <-exitCh:
			fmt.Println("Exit by command. Forecast for", city)
			break loop
		case <-time.After(Timeout):
			fmt.Println("Exit by timeout. Forecast for", city)
			break loop
		}
	}

	done <- struct{}{}
}

func RunForecastWithContext(ctx context.Context, city string, done chan<- struct{}) {
	const Timeout = time.Second

	res := GetWeatherWithContext(ctx, city)

loop:
	for {
		select {
		case temperature := <-res:
			fmt.Printf("Temperature in %s is %d C at %v\n", city, temperature, time.Now())
		case <-ctx.Done():
			fmt.Println("Exit by command. Forecast for", city, ctx.Err())
			break loop
		case <-time.After(Timeout):
			fmt.Println("Exit by timeout. Forecast for", city)
			break loop
		}
	}

	done <- struct{}{}
}

func GetWeatherWithContext(ctx context.Context, town string) chan int {
	res := make(chan int, 1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("exit %q goroutine\n", town)
				return
			case res <- GetWeatherByTime("", time.Now()):
				// nothing to do
			}
		}
	}()

	return res
}

func GetWeatherWithContextAndRes(ctx context.Context, town string, res chan int, done chan struct{}) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("exit %q goroutine\n", town)
				done <- struct{}{}
				return
			case res <- GetWeatherByTime("", time.Now()):
				// nothing to do
			}
		}
	}()
}
