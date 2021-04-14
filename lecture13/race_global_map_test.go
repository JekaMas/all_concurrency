package lecture11

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var orderPrices = make(map[string]int)

func TestGlobalMapReadWrite(_ *testing.T) {
	orderPrices["Atlanta"] = 100

	// start writers
	for i := 0; i < 2; i++ {
		go func() {
			timer := time.NewTicker(500*time.Millisecond)
			defer timer.Stop()
			for range timer.C {
				orderPrices["Atlanta"]++
			}
		}()
	}

	time.Sleep(time.Second)

	spew.Dump(orderPrices)

	fmt.Println("DONE")
}

func TestGlobalMapDoubleWrite(_ *testing.T) {
	orderPrices["Atlanta"] = 100

	// start writers
	for i := 0; i < 2; i++ {
		go func() {
			timer := time.NewTicker(500*time.Millisecond)
			defer timer.Stop()
			for range timer.C {
				orderPrices["Atlanta"] = 1
			}
		}()
	}

	time.Sleep(time.Second)

	spew.Dump(orderPrices)

	fmt.Println("DONE")
}

var idPrices = make(map[int]int)

func TestGlobalMapDoubleWriteDifferentKeys(_ *testing.T) {
	idPrices[1] = 1

	// start writers
	for i := 0; i < 2; i++ {
		go func() {
			timer := time.NewTicker(500*time.Millisecond)
			defer timer.Stop()
			for range timer.C {
				k := rand.Int()
				fmt.Println("writing to", k)
				idPrices[k] = 1
			}
		}()
	}

	time.Sleep(time.Second)

	spew.Dump(orderPrices)

	fmt.Println("DONE")
}


func TestGlobalMap100Goroutines(_ *testing.T) {
	orderPrices["Atlanta"] = 100

	// start writers
	for i := 0; i < 100; i++ {
		go func() {
			timer := time.NewTicker(time.Microsecond)
			defer timer.Stop()
			for range timer.C {
				orderPrices["Atlanta"]++
			}
		}()
	}

	time.Sleep(time.Second)

	spew.Dump(orderPrices)

	fmt.Println("DONE")
}
