package main

import (
	"sync"
	//"time"
)

type client struct {
	id           string
	keepAlive    uint16
	cleanSession bool
}

type clients struct {
	sync.RWMutex
	list map[string]*client
}

func newClient() {
}

func newClients() *clients {
	return &clients{
		sync.RWMutex{},
		make(map[string]*client),
	}
}

func (c *client) Start() {
	//c.cleanSession =
}

func (c *client) Receive(s *Comet) {
	for {
		select {
		default:
			//cp, err = ReadPacket()
		}
	}
}

func (c *client) Send() {
}

func (c *client) KeepAliveTimer(s *Comet) {
	for {
		//_ := time.NewTimer(time.Duration(float64(c.keepAlive)*1.5) * time.Second)
	}
}
