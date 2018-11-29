package main

import (
	"flag"
	"fmt"
	"github.com/michaelwolz/gosudoku"
	"io/ioutil"
	"os"
)

var (
	inputFile string
	boxID     int
	port      int
)

// Initializes flag configuration
func init() {
	flag.StringVar(&inputFile, "input", "", "Input Sudoku field as .txt")
	flag.IntVar(&boxID, "boxID", -1, "Box Number")
	flag.IntVar(&port, "port", -1, "Port")
}

// Usage: sudokuSolver -input={INPUTFILE} [-distributed] [-box={BOXNUMBER}] [-auto]
func main() {
	flag.Parse()

	if len(inputFile) == 0 {
		fmt.Println("-input option is missing! You have to provide a Sudoku field configuration.")
		os.Exit(1)
	}

	if boxID == -1 {
		fmt.Println("-boxID option is missing! You have to provide a box number [0-8]")
		os.Exit(1)
	}

	if port == -1 {
		fmt.Println("-port option is missing!")
		os.Exit(1)
	}

	gosudoku.InitializeSudoku(readFile(inputFile), &boxID)
	//gosudoku.FindBoxes(&boxID, &port)
	gosudoku.Solve("simple", &port)
}

// Reads field configuration from input file
func readFile(fname string) string {
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
