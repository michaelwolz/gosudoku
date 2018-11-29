package main

import (
	"flag"
	"github.com/grandcat/zeroconf"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
	Simple mDNS Server to make it possible to find
	sudoku boxes via mDNS (aka. bonjour)

	Most of the code is copied from the server example
	of the mDNS library: https://github.com/grandcat/zeroconf
	Only small changes were made for our use case.
*/

var (
	name    = flag.String("name", "SudokuFun", "The name for the service.")
	service = flag.String("service", "_sudokusolver._tcp", "Set the service type of the new service.")
	domain  = flag.String("domain", "local.", "Set the network domain. Default should be fine.")
	port    = flag.Int("port", 42424, "Set the port the service is listening to.")
)

func main() {
	flag.Parse()

	server, err := zeroconf.Register(*name, *service, *domain, *port, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	log.Println("Published service:")
	log.Println("- Name:", *name)
	log.Println("- Type:", *service)
	log.Println("- Domain:", *domain)
	log.Println("- Port:", *port)

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // Wait for kill signal

	log.Println("Shutting down.")
}
