package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/brave-experiments/viproxy"
	"github.com/mdlayher/vsock"
)

func parseAddr(rawAddr string) net.Addr {
	var addr net.Addr
	var err error

	addr, err = net.ResolveTCPAddr("tcp", rawAddr)
	if err == nil {
		return addr
	}

	// We couldn't parse the address, so we must be dealing with AF_VSOCK.  We
	// expect an address like 3:8080.
	fields := strings.Split(rawAddr, ":")
	if len(fields) != 2 {
		log.Fatal("Looks like we're given neither AF_INET nor AF_VSOCK addr.")
	}
	cid, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal("Couldn't turn CID into integer.")
	}
	port, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatal("Couldn't turn port into integer.")
	}

	return &vsock.Addr{ContextID: uint32(cid), Port: uint32(port)}
}

func main() {
	// E.g.: IN_ADDR=127.0.0.1:8080 OUT_ADDR=3:8080 go run main.go
	rawInAddr, rawOutAddr := os.Getenv("IN_ADDR"), os.Getenv("OUT_ADDR")
	if rawInAddr == "" || rawOutAddr == "" {
		log.Fatal("Environment variables IN_ADDR and OUT_ADDR not set.")
	}

	tuple := &viproxy.Tuple{
		InAddr:  parseAddr(rawInAddr),
		OutAddr: parseAddr(rawOutAddr),
	}
	p := viproxy.NewVIProxy([]*viproxy.Tuple{tuple})
	p.Start()
	<-make(chan bool)
}
