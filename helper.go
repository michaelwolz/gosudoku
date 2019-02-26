package gosudoku

import (
	"log"
	"strconv"
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

// Check if int array strContains specific int
func intContains(haystack []int, needle int) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}
	return false
}

// Returns x,y-Coordinates and value from a fieldConfig-String
func readFieldConfigStr(config string) (int, int, int) {
	x, err := strconv.Atoi(string(config[0]))
	y, err := strconv.Atoi(string(config[1]))
	v, err := strconv.Atoi(string(config[3]))
	checkSoftErr(err)
	return x, y, v
}
