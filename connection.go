package gosudoku

import (
	"context"
	"flag"
	"github.com/grandcat/zeroconf"
	"log"
	"time"
)

type Connection interface {
	getRow() []int
	getCol() []int
}

// List of addresses to corresponding box numbers
var boxList = make([]string, 9)

// Use mDNS to find other boxes
func MDNS() {
	var (
		service  = "_workstation._tcp"
		domain   = "local."
		waitTime = 10
	)

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*waitTime))
	defer cancel()
	err = resolver.Browse(ctx, *service, *domain, entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	// Wait some additional time to see debug messages on go routine shutdown.
	time.Sleep(1 * time.Second)
}

func ConnectToServer(ip string, port int) {
	// TODO: Implement simple connection to a server
}

func DecentralConnection() {
	// TODO: Implement a fully decentralized version
	// Probably not gonna happen, but it would be pretty nice :D
}
