package main

import (
//"os"
//"os/signal"
//"syscall"
)

func main() {
	s := NewComet()
	go s.Listener()

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//<-c
	HandleSignal(InitSignal())
}
