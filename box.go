package gosudoku

import (
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
			b.possibleValues[field] = make(map[int]struct{})
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
func (b *box) setRowValue(ycoord int, val int) {
	b.rowValues[ycoord] = append(b.rowValues[ycoord], val)
	for i := 0; i < 3; i++ {
		index := ycoord*3 + i
		b.removeFromPossibleValues(index, val)
	}
}

// Set column value
func (b *box) setColValue(xcoord int, val int) {
	b.colValues[xcoord] = append(b.colValues[xcoord], val)
	for i := 0; i < 3; i++ {
		index := xcoord + i*3
		b.removeFromPossibleValues(index, val)
	}
}

// Check and set possible values
func (b *box) checkAndSet(index int) {
	if len(b.possibleValues[index]) == 1 {
		var val int
		for key := range b.possibleValues[index] {
			val = key
			delete(b.possibleValues[index], key)
		}
		x, y := getCoordinatesForIndex(index)
		log.Println("Setting value at pos: " + strconv.Itoa(index) + " to: " + strconv.Itoa(val))
		b.values[index] = val
		b.drawBox()
		b.removeFromAllPossibleValues(val)
		sendToNeighbors(x, y, val)
	}
}

// Removes value from possible values from all field
func (b *box) removeFromPossibleValues(index, val int) {
	if b.values[index] == 0 {
		if _, ok := b.possibleValues[index][val]; ok {
			/*fmt.Print("Removing "  + strconv.Itoa(val) + " from index " + strconv.Itoa(index) + ": ")
			fmt.Print(b.possibleValues[index])
			fmt.Print(" -> ")*/
			delete(b.possibleValues[index], val)
			b.checkAndSet(index)
		}
	}
}

// Removes value from all possible values of myBox
func (b *box) removeFromAllPossibleValues(val int) {
	for index := range b.values {
		b.removeFromPossibleValues(index, val)
	}
}

// Draws box for pretty output
func (b *box) drawBox() {
	fmt.Printf("╭─────┬─────┬─────╮\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.getFieldValue(0, 0), b.getFieldValue(1, 0), b.getFieldValue(2, 0))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.getFieldValue(0, 1), b.getFieldValue(1, 1), b.getFieldValue(2, 1))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.getFieldValue(0, 2), b.getFieldValue(1, 2), b.getFieldValue(2, 2))
	fmt.Printf("╰─────┴─────┴─────╯\n")
}
