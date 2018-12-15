package gosudoku

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type box struct {
	id             string
	values         [9]int
	possibleValues [9]map[int]struct{}
	rowValues      [3][]int // Stores all values which are set in a whole row (includes values from other boxConnections)
	colValues      [3][]int // Stores all values which are set in a whole column (includes values from other boxConnections)
}

// Initializes the field configuration from a given string
// Format: xy:v with x between 0 and 2 (column) and y between 0 and 2 (row) and value v, separated by comma
// TODO: Add error handling for malformed field configurations
func (b *box) initializeBox(boxID *string, fieldString string) {
	b.id = *boxID
	log.Println("Reading input configuration...")
	config := strings.Split(fieldString, ",")
	for _, el := range config {
		x, err := strconv.Atoi(string(el[0]))
		y, err := strconv.Atoi(string(el[1]))
		v, err := strconv.Atoi(string(el[3]))
		checkErr(err)
		b.setFieldValue(x, y, v)
	}

	// Set initial possible values
	for field, value := range b.values {
		if value == 0 {
			for i := 1; i < 10; i++ {
				if !intContains(b.values[:], i) {
					b.possibleValues[field][i] = struct{}{}
				}
			}
		}
	}
}

// Set field value via coordinates
func (b *box) setFieldValue(x, y, v int) {
	// Matrix conversion, see: https://stackoverflow.com/a/14015582
	b.values[x+y*3] = v
}

// Set field value via coordinates
func (b *box) getFieldValue(x, y int) int {
	return b.values[x+y*3]
}

//Helper function to return coordinates from index
func getCoordinatesForIndex(index int) (int, int) {
	return index % 3, index / 3
}

// Set row value
func (b *box) setRowValue(xcoord int, val int) {
	b.rowValues[xcoord] = append(b.rowValues[xcoord], val)
	for i := 0; i < 3; i++ {
		index := xcoord*3 + i
		delete(b.possibleValues[index], val)
		b.checkAndSet(index)
	}
}

// Set column value
func (b *box) setColValue(ycoord int, val int) {
	b.colValues[ycoord] = append(b.colValues[ycoord], val)
	for i := 0; i < 3; i++ {
		index := ycoord + i*3
		delete(b.possibleValues[index], val)
		b.checkAndSet(index)
	}
}

// Check and set possible values
func (b *box) checkAndSet(index int) {
	if len(b.possibleValues[index]) < 2 {
		var val int
		for key := range b.possibleValues[index] {
			val = key
			delete(b.possibleValues[index], key)
		}
		b.values[index] = val
		x, y := getCoordinatesForIndex(index)
		sendToNeighbors(x, y, val)
	}
}

// Removes value from possible values from all field
func (b *box) removeFromPossibleValues(val int) {
	for field := range b.possibleValues {
		if _, ok := b.possibleValues[field][val]; ok {
			delete(b.possibleValues[field], val)
			b.checkAndSet(field)
		}
	}
}

// Removes possible value from a field
func (b *box) removePossibleValues(field int, value int) {
	delete(b.possibleValues[field], value)
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

// Draws box for pretty output
func (b *box) DrawBox() {
	fmt.Printf("╭─────┬─────┬─────╮\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.getFieldValue(0, 0), b.getFieldValue(1, 0), b.getFieldValue(2, 0))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.getFieldValue(0, 1), b.getFieldValue(1, 1), b.getFieldValue(2, 1))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.getFieldValue(0, 2), b.getFieldValue(1, 2), b.getFieldValue(2, 2))
	fmt.Printf("╰─────┴─────┴─────╯\n")
}
