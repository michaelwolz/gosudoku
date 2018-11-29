package main

import (
	"flag"
	"fmt"
	"github.com/michaelwolz/gosudoku"
	"io/ioutil"
	"os"
)

var inputFile string

// Initializes flag configuration
func init() {
	flag.StringVar(&inputFile, "input", "", "Input Sudoku field as .txt")
}

// Starts sudoku_solver
// Usage: sudoku_solder -input={INPUTFILE} [-distributed] [-box={BOXNUMBER}] [-auto]
func main() {
	flag.Parse()

	if len(inputFile) == 0 {
		fmt.Println("-input option is missing! You have to provide a Sudoku field configuration.")
		os.Exit(1)
	}

	//gosudoku.InitializeSudoku(readFile(inputFile))
	//gosudoku.Solve("simple")
	gosudoku.MDNS()
}

// Reads field configuration from input file
func readFile(fname string) string {
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
