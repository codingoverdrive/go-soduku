package sodukusolver

import "math"

//holds the current board with the intial given cells and the new solved cells
var board = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

//holds the current notes for cells that are not yet solved
//solved cells will have a value of zero for the cell notes
//notes are represented as 9bit integers with a 1 in each bit denoting a note value
//0x07 means that the notes values are 1,2, and 3 (00000111 in binary)
var notes = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

//InitaliseBoard initialises the board
func InitaliseBoard(newBoard [9][9]int) {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			board[row][column] = newBoard[row][column]
		}
	}
}

//PrintBoard outputs the current state of the board
//The showNotes parameter determines whether the notes are also displayed
func PrintBoard(showNotes bool) {

}

//recalculateBoardNotes takes the board and generates a new set of notes
//from the current solved cells
func recalculateBoardNotes(board [9][9]int) [9][9]int {
	var newNotes [9][9]int
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] == 0 {
				rowSolvedNumbers := getSolvedNumbersInNineCells(convertRowToNineCells(board, row))
				colSolvedNumbers := getSolvedNumbersInNineCells(convertColumnToNineCells(board, column))
				blockIndex := row/3 + column%3
				blockSolvedNumbers := getSolvedNumbersInNineCells(convertBlockToNineCells(board, blockIndex))
				newNotes[row][column] = 0x1ff ^ (rowSolvedNumbers | colSolvedNumbers | blockSolvedNumbers)
			} else {
				//cell already solved
				newNotes[row][column] = 0
			}
		}
	}
	return newNotes
}

func convertRowToNineCells(board [9][9]int, row int) [9]int {
	var cells [9]int
	for column := 0; column < 9; column++ {
		cells[column] = board[row][column]
	}
	return cells
}

func convertColumnToNineCells(board [9][9]int, column int) [9]int {
	var cells [9]int
	for row := 0; row < 9; row++ {
		cells[row] = board[row][column]
	}
	return cells
}

func convertBlockToNineCells(board [9][9]int, blockIndex int) [9]int {
	startRow := 3 * (blockIndex / 3)
	startColumn := 3 * (blockIndex % 3)
	var cells [9]int
	for k := 0; k < 9; k++ {
		cells[k] = board[startRow+k/3][startColumn+k%3]
	}
	return cells
}

func getSolvedNumbersInNineCells(cells [9]int) int {
	numbersSet := 0
	for i := 0; i < 9; i++ {
		//the cell will contain a single number between 1 and 9, or a zero if the cell is unsolved
		if cells[i] > 0 {
			//1 is represented by the zero bit position, hence the subtraction below
			exp := (float64)(cells[i] - 1)
			//perform an OR with the previous values already found
			numbersSet = numbersSet | (int)(math.Pow(2, exp))
		}
	}
	return numbersSet
}
