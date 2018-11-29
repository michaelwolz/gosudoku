package gosudoku

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type box struct {
	id     int
	values [9]int
}

// Returns true if all values of the sudoku box are filled
// TODO: THIS IS JUST WRONG!
func (b *box) IsFull() bool {
	return len(b.values) == cap(b.values)
}

// Helper function to set field value via coordinates
func (b *box) SetFieldValue(x, y, v int) {
	// Matrix conversion, see: https://stackoverflow.com/a/14015582
	b.values[x+y*3] = v
}

// Helper function to set field value via coordinates
func (b *box) GetFieldValue(x, y int) int {
	return b.values[x+y*3]
}

// Initializes the field configuration from a given string
// Format: xy:v with x between 0 and 2 (column) and y between 0 and 2 (row) and value v, separated by comma
// TODO: Add error handling for malformed field configurations
func (b *box) InitializeBox(boxID *int, fieldString string) {
	b.id = *boxID
	log.Println("Reading input configuration...")
	config := strings.Split(fieldString, ",")
	for _, el := range config {
		x, err := strconv.Atoi(string(el[0]))
		y, err := strconv.Atoi(string(el[1]))
		v, err := strconv.Atoi(string(el[3]))
		if err != nil {
			panic(err)
		}
		b.SetFieldValue(x, y, v)
	}
}

// Get row values of local box.
// row in range [0:2]
func (b *box) getRow(row int) ([]int, error) {
	if row < 0 || row > 2 {
		return nil, errors.New("row number out of range")
	}
	return b.values[row*3 : row*3+3], nil
}

// Get column values of local box
// col in range [0:2]
func (b *box) getCol(col int) ([]int, error) {
	if col < 0 || col > 2 {
		return nil, errors.New("col number out of range")
	}
	colValues := make([]int, 3)
	for i := 0; i < 3; i++ {
		colValues[i] = b.values[i*3+col]
	}
	return colValues, nil
}

// Defining ASCII constants for drawing the sudoku field
// FOR FUTURE DEVELOPMENT
const (
	h_line       = "─"
	v_line       = "│"
	cross        = "┼"
	t_bar_top    = "┬"
	t_bar_bottom = "┴"
	rt_corner    = "╭"
	rb_corner    = "╰"
	lt_corner    = "╮"
	lb_corner    = "╯"
)

// Draws box for pretty output
// TODO: Make this nice
func (b *box) DrawBox() {
	fmt.Printf("╭─────┬─────┬─────╮\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.GetFieldValue(0, 0), b.GetFieldValue(1, 0), b.GetFieldValue(2, 0))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.GetFieldValue(0, 1), b.GetFieldValue(1, 1), b.GetFieldValue(2, 1))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.GetFieldValue(0, 2), b.GetFieldValue(1, 2), b.GetFieldValue(2, 2))
	fmt.Printf("╰─────┴─────┴─────╯\n")
}
