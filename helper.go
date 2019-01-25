package gosudoku

import (
	"log"
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
