package lecture11

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type forecasterShared struct {
	//sync.RWMutex // мы можем сделать общую блокировку на все данные структуры

	city   string
	cityMu sync.RWMutex

	state  *uint32
	closed chan struct{}
}

func newForecasterShared(city string) *forecasterShared {
	return &forecasterShared{
		city:   city,
		state:  new(uint32),
		closed: make(chan struct{}),
	}
}

func (f *forecasterShared) isStarted() bool {
	return f.isState(StartState)
}

func (f *forecasterShared) isStopped() bool {
	return f.isState(StopState)
}

func (f *forecasterShared) isState(state uint32) bool {
	return atomic.LoadUint32(f.state) == state
}

func (f *forecasterShared) setStartState() {
	f.setState(StartState)
}

func (f *forecasterShared) setStopState() {
	f.setState(StopState)
}

func (f *forecasterShared) setState(state uint32) {
	atomic.StoreUint32(f.state, state)
}

func (f *forecasterShared) getCity() string {
	f.cityMu.RLock()
	defer f.cityMu.RUnlock()

	return f.city
}

func (f *forecasterShared) setCity(city string) {
	f.cityMu.Lock()
	defer f.cityMu.Unlock()

	f.city = city
}

// start without waiting
func (f *forecasterShared) start() {
	if !atomic.CompareAndSwapUint32(f.state, StopState, StartState) {
		return
	}

	res, exit := GetWeather(f.city)
	f.closed = exit

	go func() {
		for {
			select {
			case temp := <-res:
				fmt.Printf("Temperature in %s is %d C at %v\n", f.getCity(), temp, time.Now())
			case <-f.closed:
				fmt.Println("Exit")
				return
			}
		}
	}()

	fmt.Println("Started")
}

// stop without waiting
func (f *forecasterShared) stop() {
	if !atomic.CompareAndSwapUint32(f.state, StartState, StopState) {
		return
	}

	close(f.closed)

	fmt.Println("Stopped")
}

func TestSharedMemory(_ *testing.T) {
	city := "London"
	cast := newForecasterShared(city)

	cast.start()

	time.Sleep(5 * time.Second)

	cast.setCity("Atlanta")

	time.Sleep(5 * time.Second)

	cast.stop()

	fmt.Println("DONE")
}
