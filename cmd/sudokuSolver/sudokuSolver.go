package main

import (
	"flag"
	"fmt"
	"github.com/michaelwolz/gosudoku"
	"io/ioutil"
	"os"
)

var (
	inputFile      string
	managerAddress string
	boxID          string
	lport          int
	mport          int
)

// Initializes flag configuration
func init() {
	flag.StringVar(&inputFile, "input", "", "Input Sudoku field as .txt")
	flag.StringVar(&boxID, "boxID", "", "Box Number")
	flag.IntVar(&mport, "mport", -1, "Manager port")
	flag.IntVar(&lport, "lport", -1, "Local port")
	flag.StringVar(&managerAddress, "maddress", "127.0.0.1", "Address of the box manager")
}

// Usage: sudokuSolver -input={INPUTFILE} -lport={LOCALPORT} -maddress={MANAGERADDRESS} -mport={MANAGERPORT} -boxID={BOXNUMBER}
func main() {
	flag.Parse()

	if len(inputFile) == 0 {
		fmt.Println("-input option is missing! You have to provide a Sudoku field configuration.")
		os.Exit(1)
	}

	if boxID == "" {
		fmt.Println("-boxID option is missing!")
		os.Exit(1)
	}

	if lport == -1 {
		fmt.Println("-lport option is missing!")
		os.Exit(1)
	}

	if mport == -1 {
		fmt.Println("-mport option is missing!")
		os.Exit(1)
	}

	gosudoku.InitializeSudoku(readFile(inputFile), &boxID)
	//gosudoku.LaunchTCPServer(&lport)
	gosudoku.ConnectToManager(&managerAddress, &mport, &lport)
	//gosudoku.Solve()
}

// Reads field configuration from input file
func readFile(fname string) string {
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
