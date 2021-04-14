package lecture11

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type order struct {
	total    *big.Int
	products []string
}

func increasePrice(order order, price int64) order {
	order.total.Add(order.total, big.NewInt(price))

	return order
}

func TestCopyObjectRace(_ *testing.T) {
	// create value
	order := order{big.NewInt(100), []string{"apple"}}

	go func() {
		timer := time.NewTicker(50 * time.Millisecond)
		defer timer.Stop()
		for range timer.C {
			// pass by value
			increasePrice(order, 10)
		}
	}()

	time.Sleep(time.Second)

	spew.Dump(order)

	fmt.Println("DONE")
}
