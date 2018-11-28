package gosudoku

import (
	"fmt"
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
func (b *box) InitializeBox(fieldString string) {
	fmt.Println("Reading input configuration...")
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
