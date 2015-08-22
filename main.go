package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/surge/glog"
	"github.com/surgemq/surgemq/service"
)

var (
	keepAlive        int
	connectTimeout   int
	ackTimeout       int
	timeoutRetries   int
	authenticator    string
	sessionsProvider string
	topicsProvider   string
	cpuprofile       string
)

func init() {
	flag.IntVar(&keepAlive, "keepalive", service.DefaultKeepAlive, "Keepalive (sec)")
	flag.IntVar(&connectTimeout, "connecttimeout", service.DefaultConnectTimeout, "Connect Timeout (sec)")
	flag.IntVar(&ackTimeout, "acktimeout", service.DefaultAckTimeout, "Ack Timeout (sec)")
	flag.IntVar(&timeoutRetries, "retries", service.DefaultTimeoutRetries, "Timeout Retries")
	flag.StringVar(&authenticator, "auth", service.DefaultAuthenticator, "Authenticator Type")
	flag.StringVar(&sessionsProvider, "sessions", service.DefaultSessionsProvider, "Session Provider Type")
	flag.StringVar(&topicsProvider, "topics", service.DefaultTopicsProvider, "Topics Provider Type")
	flag.StringVar(&cpuprofile, "cpuprofile", "", "CPU Profile Filename")
	flag.Parse()
}

func main() {
	svr := &service.Server{
		KeepAlive:        keepAlive,
		ConnectTimeout:   connectTimeout,
		AckTimeout:       ackTimeout,
		TimeoutRetries:   timeoutRetries,
		SessionsProvider: sessionsProvider,
		TopicsProvider:   topicsProvider,
	}

	var f *os.File
	var err error

	if cpuprofile != "" {
		f, err = os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, os.Kill)
	go func() {
		sig := <-sigchan
		glog.Errorf("Existing due to trapped signal; %v", sig)

		if f != nil {
			glog.Errorf("Stopping profile")
			pprof.StopCPUProfile()
			f.Close()
		}

		svr.Close()

		os.Exit(0)
	}()

	err = svr.ListenAndServe("tcp://:1883")
	if err != nil {
		glog.Errorf("surgemq/main: %v", err)
	}
}
