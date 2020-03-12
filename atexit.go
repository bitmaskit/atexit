package atexit

import (
	"os"
	"os/signal"
	"syscall"
)

var exitFuncs []func()

func Run(f func()) {
	exitFuncs = append(exitFuncs, f)
}

func callExitFuncs() {
	for i := len(exitFuncs) - 1; i >= 0; i-- {
		exitFuncs[i]()
	}
}

func trapSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		callExitFuncs()
		os.Exit(1)
	}()
}
