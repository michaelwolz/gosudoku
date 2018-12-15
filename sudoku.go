package gosudoku

import (
	"log"
)

var myBox box

// Initializes the game
func InitializeSudoku(fieldString string, boxID *string) {
	log.Println("Initializig Sudoku Solver!")
	myBox.initializeBox(boxID, fieldString)
	myBox.DrawBox()
}
