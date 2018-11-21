package gosudoku

import "fmt"

// Initializes the game
func InitializeSudoku(fieldString string) {
	fmt.Println("Initializig Sudoku Solver!")

	var box box
	box.InitializeBox(fieldString)
	box.DrawBox()
}
