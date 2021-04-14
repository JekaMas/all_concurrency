package lecture11

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// go test ./lecture11/... -run=TestExitSignal -v
// Действие по сигналу из OS. Например мы ходим дообработать все пришедшие уже запросы, но перестать получать новые, затем, как будут обработаны все текузщие запросы, то в штатном режиме закрыть БД, логи, соединенния, только затем выйти.
func TestExitSignal(_ *testing.T) {
	sigs := make(chan os.Signal, 1)
	exitCh := make(chan struct{}, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	cities := []string{"London", "Moscow", "Berlin", "Madrid", "Osaka", "Tokyo", "Bangkok", "Pattaya"}
	for _, city := range cities{
		go RunForecast(city, exitCh)
	}

	go func() {
		sig := <-sigs
		fmt.Println("Got a signal! Processing...", sig)
		fmt.Println("Closing goroutines")
		close(exitCh)
	}()

	fmt.Println("awaiting signal")
	<-exitCh
	fmt.Println("DONE")
}