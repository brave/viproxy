package main

import (
	"log"
	"net"

	"github.com/brave-experiments/viproxy"
)

func resolveTCPAddr(addr string) *net.TCPAddr {
	a, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to resolve TCP address: %s", err)
	}
	return a
}

func main() {
	tuple := &viproxy.Tuple{
		InAddr:  resolveTCPAddr("127.0.0.1:8080"),
		OutAddr: resolveTCPAddr("127.0.0.1:443"),
	}
	p := viproxy.NewVIProxy([]*viproxy.Tuple{tuple})
	p.Start()
	<-make(chan bool)
}
