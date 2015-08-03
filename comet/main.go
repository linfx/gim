package main

import (
	"log"
	//"flag"
	//"fmt"
	//"os"
	//"github.com/dtxlink/gim"
	"github.com/dtxlink/gim/signal"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//	configPtr := flag.String("config", "", "config file")
	//	flag.Usage = usage
	//	flag.Parse()

	log.Println("gim start up")
	StartTCP()
	signal.HandleSignal(signal.InitSignal())
}
