package lecture11

import (
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func addProduct(order *order, product string) {
	order.products = append(order.products, product)
}

func TestSliceRace(_ *testing.T) {
	order := &order{}

	go func() {
		timer := time.NewTicker(50 * time.Millisecond)
		defer timer.Stop()
		for range timer.C {
			addProduct(order,"Atlanta")
		}
	}()

	time.Sleep(time.Second)

	spew.Dump(order)

	fmt.Println("DONE")
}
