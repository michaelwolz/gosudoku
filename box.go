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
	rowValues      [3][9]int // Stores all values which are set in a whole row (includes values from other boxes)
	colValues      [3][9]int // Stores all values which are set in a whole column (includes values from other boxes)
}

// Initializes the field configuration from a given string
// Format: xy:v with x between 0 and 2 (column) and y between 0 and 2 (row) and value v, separated by comma
// TODO: Add error handling for malformed field configurations
func (b *box) InitializeBox(boxID *string, fieldString string) {
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

// Set field value via coordinates
func (b *box) SetFieldValue(x, y, v int) {
	// Matrix conversion, see: https://stackoverflow.com/a/14015582
	b.values[x+y*3] = v
}

// Set field value via coordinates
func (b *box) GetFieldValue(x, y int) int {
	return b.values[x+y*3]
}

// Calculate possible values for empty fields
func (b *box) CalculatePossibleValues() {
	for field, value := range b.values {
		if value == 0 {
			var impossibleValues = map[int]struct{}{}
			//addValuesToMap(b.rowValues[0], &impossibleValues)
			//addValuesToMap(b.colValues[0], &impossibleValues)
			//addValuesToMap(b.values, &impossibleValues)
			for i := 0; i < 9; i++ {
				if _, ok := impossibleValues[i]; !ok {
					b.possibleValues[field][i] = struct{}{}
				}
			}
			if len(b.possibleValues[field]) < 2 {
				var err error
				b.values[field], err = getKey(b.possibleValues[field])
				if err != nil {
					panic(err)
				}
				b.CalculatePossibleValues()
			}
		}
	}
}

// Add multiple values to map structure
func addValuesToMap(values []int, m map[int]struct{}) {
	for i := range values {
		m[i] = struct{}{}
	}
}

// Removes possible value from a field
func (b *box) removePossibleValues(field int, value int) {
	delete(b.possibleValues[field], value)
}

// Helper function to get first key of map
func getKey(m map[int]struct{}) (key int, err error) {
	for k := range m {
		return k, nil
	}
	return 0, errors.New("empty map")
}

// Helper function for checking if integer value is in slice (linear time :S)
func InSlice(slice []int, value int) bool {
	for _, x := range slice {
		if value == x {
			return true
		}
	}
	return false
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
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.GetFieldValue(0, 0), b.GetFieldValue(1, 0), b.GetFieldValue(2, 0))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.GetFieldValue(0, 1), b.GetFieldValue(1, 1), b.GetFieldValue(2, 1))
	fmt.Printf("├─────┼─────┼─────┤\n")
	fmt.Printf("│  %d  │  %d  │  %d  │\n", b.GetFieldValue(0, 2), b.GetFieldValue(1, 2), b.GetFieldValue(2, 2))
	fmt.Printf("╰─────┴─────┴─────╯\n")
}
