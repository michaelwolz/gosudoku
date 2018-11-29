package gosudoku

/*
	Simple mDNS Server to make it possible to find
	sudoku boxes via mDNS (aka. bonjour)

	Most of the code is copied from the examples
	of the mDNS library: https://github.com/grandcat/zeroconf
	Only small changes were made for our use case.
*/

import (
	"context"
	"github.com/grandcat/zeroconf"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	mdnsName     = "SudokuFun"
	mdnsService  = "_sudokusolver._tcp"
	mdnsDomain   = "local."
	mdnsWaitTime = 10
	mdnsPort     int
)

func registerService(boxID *int, port *int) {
	mdnsName += strconv.Itoa(*boxID)
	mdnsPort = *port

	server, err := zeroconf.Register(mdnsName, mdnsService, mdnsDomain, mdnsPort, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	log.Println("Published mdnsService:")
	log.Println("- Name:", mdnsName)
	log.Println("- Type:", mdnsService)
	log.Println("- Domain:", mdnsDomain)
	log.Println("- Port:", mdnsPort)

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	log.Println("Shutting down.")

	// TODO: I guess server could offline, when all 9 Boxes are found.
}

func searchForClients() {
	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			log.Println(entry)
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(mdnsWaitTime))
	defer cancel()

	err = resolver.Browse(ctx, mdnsService, mdnsDomain, entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	// Wait some additional time to see debug messages on go routine shutdown.
	time.Sleep(1 * time.Second)
}
