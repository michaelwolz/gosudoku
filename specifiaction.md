# Telegram-SudokuSolver Specification

## Communication

The sudoku-solver communicates via a Telegram Bot. Messages can be sent to the bot and will be redirected to all other instances of the sovler. 

## Message Format

The message format is a base64 encoded string which contains the boxname, the relative coordinate of the field and the value e.g.:
- BOX_D4,0,1:7 => Qk9YX0Q0LDAsMTo3