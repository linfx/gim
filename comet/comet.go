package main

import (
	"fmt"
	"net"
)

type Comet struct {
	clients *clients
	//listeners map[string]*Lister
}

func NewComet() *Comet {
	c := &Comet{
		clients: newClients(),
	}
	return c
}

func (c *Comet) Listener() error {
	//	l, err := net.Listen("tcp", "1883")
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//_, err := l.Accept()
	//	c.Server(conn)
	return nil
}

func (c *Comet) Stop() {
	//for	_,l:=range c.l
}

func (c *Comet) Server(conn net.Conn) {
	fmt.Println("hhh")
}
