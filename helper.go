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

func checkMessageFormat(message string) bool {
	r := regexp.MustCompile(`^BOX_[A,D,G][1,4,7],[0-2],[0-2]:[1-9]$`)
	return r.MatchString(message)
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
