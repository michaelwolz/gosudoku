package gosudoku

import (
	"log"
)

var initialized bool
var myBox box

// Initializes the game
func InitializeSudoku(fieldString string, boxID *string) {
	log.Println("Initializig Sudoku Solver!")

	myBox.initializeBox(boxID, fieldString)
	initialized = true
	myBox.DrawBox()
}

// Starts solving algorithm
func Solve() {
	if !initialized {
		log.Println("Solve(): Field not initialized! Nothing to solve.")
	}
}
