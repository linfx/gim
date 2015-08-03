package main

import (
	//"fmt"
	"log"
	"net"
)

func StartTCP() error {
	bind := ":3600"
	//	for _, bind := range 2 {
	log.Printf("start tcp listen addr:\"%s\"", bind)
	go tcpListen(bind)
	//	}
	return nil
}

func tcpListen(bind string) {
	addr, err := net.ResolveTCPAddr("tcp", bind)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr(\"tcp\"), %s) error(%v)", bind, err)
		panic(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("net.ListenTCP(\"tcp4\", \"%s\") error(%v)", bind, err)
		panic(err)
	}
	// free the listener resource
	defer func() {
		log.Printf("tcp addr: \"%s\" close", bind)
		if err := l.Close(); err != nil {
			log.Fatalf("listener.Close() error(%v)", err)
		}
	}()
	// init reader buffer instance
	//rb := newtcpBufCache()
	for {
		log.Println("start accept")
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatalf("listener.AcceptTCP() error(%v)", err)
			continue
		}
		if err = conn.SetKeepAlive(true); err != nil {
			log.Fatalf("conn.SetKeepAlive() error(%v)", err)
			conn.Close()
			continue
		}
		go handleTCPConn(conn)
		log.Println("accept finished")
	}
}

func handleTCPConn(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	log.Printf("<%s> handleTcpConn routine start", addr)
}

type tcpBufCache struct {
}

func newtcpBufCache() {
}
