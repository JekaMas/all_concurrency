package lecture11

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var orderPricesAtomic = make(map[string]*uint32)

func TestGlobalMapAtomic(_ *testing.T) {
	orderPricesAtomic["Atlanta"] = new(uint32)
	atomic.StoreUint32(orderPricesAtomic["Atlanta"], 100)

	// start writers
	for i := 0; i < 10; i++ {
		go func() {
			timer := time.NewTicker(time.Microsecond)
			defer timer.Stop()
			for range timer.C {
				atomic.AddUint32(orderPricesAtomic["Atlanta"], 1)
			}
		}()
	}

	time.Sleep(time.Second)

	spew.Dump(orderPrices)

	fmt.Println("DONE")
}
