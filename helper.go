package gosudoku

import (
	"log"
	"net"
	"regexp"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkSoftErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Get Local IP Address (https://gist.github.com/jniltinho/9787946)
func getLocalIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	checkErr(err)

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}

// Check if string array strContains specific string
func strContains(haystack []string, needle string) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}

// Check if int array strContains specific int
func intContains(haystack []int, needle int) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}

// Check IP/Port answer from boxManager
func checkIP(ip *string) bool {
	r, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]),[0-9]+$`)
	return r.MatchString(*ip)
}

// Map of all neighbors to all boxes
var neighbors = map[string][]string{
	"BOX_A1": {"BOX_A4", "BOX_D1"},
	"BOX_A4": {"BOX_A1", "BOX_A7", "BOX_D4"},
	"BOX_A7": {"BOX_A4", "BOX_D7"},
	"BOX_D1": {"BOX_A1", "BOX_D4", "BOX_G1"},
	"BOX_D4": {"BOX_A4", "BOX_D1", "BOX_D7", "BOX_G4"},
	"BOX_D7": {"BOX_D4", "BOX_A7", "BOX_G7"},
	"BOX_G1": {"BOX_D1", "BOX_G4"},
	"BOX_G4": {"BOX_G1", "BOX_D4", "BOX_G7"},
	"BOX_G7": {"BOX_G4", "BOX_D7"},
}

var redirectNeighbors = map[string]map[string]string{
	"BOX_A1": {},
	"BOX_A4": {"BOX_A1": "BOX_A7", "BOX_A7": "BOX_A1"},
	"BOX_A7": {},
	"BOX_D1": {"BOX_A1": "BOX_G1", "BOX_G1": "BOX_A1"},
	"BOX_D4": {"BOX_A4": "BOX_G4", "BOX_D1": "BOX_D7"},
	"BOX_D7": {"BOX_A7": "BOX_G7", "BOX_G7": "BOX_A7"},
	"BOX_G1": {},
	"BOX_G4": {"BOX_G1": "BOX_G7", "BOX_G7": "BOX_G1"},
	"BOX_G7": {},
}
