package gosudoku

import (
	"fmt"
	"log"
)

var initialized bool
var myBox box

// Initializes the game
func InitializeSudoku(fieldString string) {
	log.Println("Initializig Sudoku Solver!")

	myBox.InitializeBox(fieldString)
	initialized = true
	myBox.DrawBox()
}

// Starts solving algorithm
func Solve(algorithm string) {
	if !initialized {
		log.Println("Solve(): Field not initialized! Nothing to solve.")
	}

	switch algorithm {
	case "simple":
		fmt.Println("Algorithm not yet implemented")
	case "advanced":
		fmt.Println("Algorithm not yet implemented")
	case "dancinglinks":
		fmt.Println("Algorithm not yet implemented")
	default:
		fmt.Println("Unknown solving algorithm. Please use von of {\"simple\", \"advanced\", \"dancinglinks\"}")
	}
}
