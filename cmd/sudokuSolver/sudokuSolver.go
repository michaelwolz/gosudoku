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
	telegramToken  string
	mport          int
)

var Done = make(chan int)

// Initializes flag configuration
func init() {
	flag.StringVar(&inputFile, "input", "", "Input Sudoku field as .txt")
	flag.StringVar(&boxID, "boxID", "", "Box Number")
	flag.IntVar(&mport, "mport", -1, "Manager port")
	flag.StringVar(&managerAddress, "maddress", "127.0.0.1", "Address of the box manager")
	flag.StringVar(&telegramToken, "telegramtoken", "", "Token of the Telegram Bot")
}

// Usage: sudokuSolver -input={INPUTFILE} -telegramtoken={TOKEN} -maddress={MANAGERADDRESS} -mport={MANAGERPORT} -boxID={BOXNUMBER}
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

	if telegramToken == "" {
		fmt.Println("-telegramtoken option is missing!")
		os.Exit(1)
	}

	if mport == -1 {
		fmt.Println("-mport option is missing!")
		os.Exit(1)
	}

	go gosudoku.StartTelegramBot(telegramToken)
	gosudoku.InitializeSudoku(readFile(inputFile), &boxID)

	<-gosudoku.Done
}

// Reads field configuration from input file
func readFile(fname string) string {
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
