package gosudoku

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var msgProcessRegex = regexp.MustCompile(`^(BOX_[A,D,G][1,4,7]),([0-2]),([0-2]):([1-9])$`)
var msgRegex = regexp.MustCompile(`^BOX_[A,D,G][1,4,7],[0-2],[0-2]:[1-9]$`)

func processMessage(message string) {
	fmt.Println("Processing received: " + message)
	if msgRegex.MatchString(message) {
		matches := msgProcessRegex.FindStringSubmatch(message)

		val, err := strconv.Atoi(matches[4])
		checkSoftErr(err)
		if matches[1][:len(matches[1])-1] == myBox.id[:len(myBox.id)-1] {
			x, err := strconv.Atoi(matches[2])
			checkSoftErr(err)
			myBox.setColValue(x, val)
		} else {
			y, err := strconv.Atoi(matches[3])
			checkSoftErr(err)
			myBox.setRowValue(y, val)
		}
		redirectToNeighbor(message, matches[1])
	} else {
		log.Println("malformed message received " + message)
	}
}
