#!/bin/sh

screen -dmS field1 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_A1 -maddress=127.0.0.1 -mport=4242 -lport=1330 -input=./example/sudoku1.txt
screen -dmS field2 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_D1 -maddress=127.0.0.1 -mport=4242 -lport=1331 -input=./example/sudoku2.txt
screen -dmS field3 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_G1 -maddress=127.0.0.1 -mport=4242 -lport=1332 -input=./example/sudoku3.txt
screen -dmS field4 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_A4 -maddress=127.0.0.1 -mport=4242 -lport=1333 -input=./example/sudoku4.txt
screen -dmS field5 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_D4 -maddress=127.0.0.1 -mport=4242 -lport=1334 -input=./example/sudoku5.txt
screen -dmS field6 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_G4 -maddress=127.0.0.1 -mport=4242 -lport=1335 -input=./example/sudoku6.txt
screen -dmS field7 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_A7 -maddress=127.0.0.1 -mport=4242 -lport=1336 -input=./example/sudoku7.txt
screen -dmS field8 go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_D7 -maddress=127.0.0.1 -mport=4242 -lport=1337 -input=./example/sudoku8.txt
go run cmd/sudokuSolver/sudokuSolver.go -boxID=BOX_G7 -maddress=127.0.0.1 -mport=4242 -lport=1338 -input=./example/sudoku9.txt