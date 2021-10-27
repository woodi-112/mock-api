package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func shutdownListener(signals chan os.Signal) chan bool {
	shutdown := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		signal := <-signals
		fmt.Printf("'%s' signal received from system", signal)
		shutdown <- true
	}()

	return shutdown
}
