package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	host, container := parseHostContainerAddrs()

	p, err := NewProxy(host, container)
	if err != nil {
		fmt.Fprintf(os.Stderr, "1\n%s", err)
		os.Exit(1)
	}
	go handleStopSignals(p)

	fmt.Printf("listening on %s -> server %s ...\n", p.FrontendAddr(), p.BackendAddr())

	// Run will block until the proxy stops
	p.Run()
}

// parseHostContainerAddrs parses the flags passed on reexec to create the TCP or UDP
// net.Addrs to map the host and container ports
func parseHostContainerAddrs() (host net.Addr, container net.Addr) {
	var (
		proto      = flag.String("proto", "tcp", "proxy protocol")
		listenAddr = flag.String("listen", "", "host:port")
		serverAddr = flag.String("server", "", "host:port")
	)

	flag.Parse()

	if *listenAddr == "" || *serverAddr == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch *proto {
	case "tcp":
		host, _ = net.ResolveTCPAddr("tcp4", *listenAddr)
		container, _ = net.ResolveTCPAddr("tcp4", *serverAddr)
	case "udp":
		host, _ = net.ResolveUDPAddr("udp4", *listenAddr)
		container, _ = net.ResolveUDPAddr("udp4", *serverAddr)
	default:
		log.Fatalf("unsupported protocol %s", *proto)
	}

	return host, container
}

func handleStopSignals(p Proxy) {
	s := make(chan os.Signal, 10)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)

	for range s {
		p.Close()

		os.Exit(0)
	}
}
