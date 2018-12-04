package gosudoku

import (
	"log"
)

var initialized bool
var MyBox box

// Initializes the game
func InitializeSudoku(fieldString string, boxID *string) {
	log.Println("Initializig Sudoku Solver!")

	MyBox.InitializeBox(boxID, fieldString)
	initialized = true
	MyBox.DrawBox()
}

// Starts solving algorithm
func Solve() {
	if !initialized {
		log.Println("Solve(): Field not initialized! Nothing to solve.")
	}
}
