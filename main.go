package main

import (
	"flag"

	"github.com/colin-404/go-syslog-listener/listener"
)

func main() {
	//根据参数是udp还是tcp
	protocol := flag.String("protocol", "udp", "udp or tcp")
	listenPort := flag.String("port", "5144", "listen port")
	flag.Parse()
	listener.ListenSyslog(*protocol, *listenPort)
}
