package sodukusolver

import (
	"fmt"
	"math"
	"time"
)

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

//holds the current notes that have been eliminated by previous solutions
//ths is because some strategies identify the values that cannot be in unsolved cells
var exclusions = [9][9]int{
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

//StrategyResult indicates the success or failure of a strategy
//and how many solutions were found for the strategy
type StrategyResult struct {
	success   bool
	solutions int
}

//CellStrategy a solver strategy that returns an array of AbsoluteCellSolutions
type CellStrategy func([9][9]int) []AbsoluteCellSolution

//applyCellStrategy applies a cell strategy and returns a result
func applyCellStrategy(strategy CellStrategy, duration *time.Duration) StrategyResult {
	startTime := time.Now()
	solutions := strategy(notes)
	elapsed := time.Since(startTime)
	if len(solutions) > 0 {
		for i := 0; i < len(solutions); i++ {
			s := solutions[i]
			board[s.Row][s.Column] = s.Number
			print("Found ", s.Type, " [", s.Number, "] at ", s.Location, " ", getBoardRowLetter(s.Row), (1 + s.Column), "\n")
		}
		notes = recalculateBoardNotes(board)
	}
	*duration = *duration + elapsed
	return StrategyResult{len(solutions) > 0, len(solutions)}
}

//ExclusionStrategy a solver strategy the eliminates/excludes notes that cannot be a cell solution
type ExclusionStrategy func([9][9]int) []CellExclusion

//applyCellExclusionStrategy applies a cell notes exclusion strategy and returns a result
func applyCellExclusionStrategy(strategy ExclusionStrategy, duration *time.Duration) StrategyResult {
	startTime := time.Now()
	exclusionSolutions := findNakedPairs(notes)
	elapsed := time.Since(startTime)
	if len(exclusionSolutions) > 0 {
		for i := 0; i < len(exclusionSolutions); i++ {
			s := exclusionSolutions[i]
			print("Found ", s.strategy, " [", getNotesAsDigitString(s.number), "] in ", getCellRefsAsString(s.matches), "\n")
			print("Removing from ", getCellRefsAsString(s.exclusions), "\n")
			exclRefs := s.exclusions
			for k := 0; k < len(exclRefs); k++ {
				exclusions[exclRefs[k].Row][exclRefs[k].Column] = exclusions[exclRefs[k].Row][exclRefs[k].Column] | s.number
			}
		}
		notes = recalculateBoardNotes(board)
	}
	*duration = *duration + elapsed
	return StrategyResult{len(exclusionSolutions) > 0, len(exclusionSolutions)}
}

//SolveBoard solves the board
func SolveBoard(newBoard [9][9]int) {
	print("Soduku Solver\n\n")

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			board[row][column] = newBoard[row][column]
		}
	}

	//initialise the notes
	notes = recalculateBoardNotes(board)

	print("\nInitial Board\n")
	PrintBoard(false)
	print("\nWith Notes\n")
	PrintBoard(true)
	print("\n")

	//keep track how how long the (full) solution takes to compute
	var duration time.Duration

	//keep tracks of how many steps are required to solve the puzzle
	passCount := 0

	//keep applying the different solution strategies until the puzzle is solved
	for {
		passCount++
		print("Pass ", passCount, "\n")

		//findNakedSingles
		result := applyCellStrategy(findNakedSingles, &duration)
		if result.success {
			passCount += result.solutions - 1
			continue
		}

		//findHiddenSingles
		result = applyCellStrategy(findHiddenSingles, &duration)
		if result.success {
			passCount += result.solutions - 1
			continue
		}

		result = applyCellExclusionStrategy(findNakedPairs, &duration)
		if result.success {
			passCount += result.solutions - 1
			continue
		}

		break
	}
	print("\n")
	PrintBoard(true)

	if IsSolved() {
		print("\nSolved in ", fmt.Sprintf("%s", duration), "\n")
	} else {
		print("\nUnsolved after ", fmt.Sprintf("%s", duration), "\n")
		print("Stopping\n\n")
	}
}

//PrintBoard outputs the current state of the board
//The showNotes parameter determines whether the notes are also displayed
func PrintBoard(showNotes bool) {
	verticalBlockSeparator := "I"
	print("       1       2       3       4       5       6       7       8       9\n")
	for row := 0; row < 9; row++ {
		if row%3 == 0 {
			//used to differentiate the blocks horizontally
			print("   =========================================================================\n")
		} else {
			print("   --------+-------+-------+-------+-------+-------+-------+-------+--------\n")
		}

		//each row is composed of three printed lines
		//this is so that the notes can be represented as a 9x9 grid within each cell
		for z := 0; z < 3; z++ {
			if z == 1 {
				print(" ", getBoardRowLetter(row), " ")
			} else {
				print("   ")
			}
			for column := 0; column < 9; column++ {
				if column%3 == 0 {
					//used to differentiate the blocks vertically
					print(verticalBlockSeparator)
				} else {
					print("|")
				}
				print("  ")

				actualValue := board[row][column]
				if actualValue == 0 && showNotes {
					noteValue := notes[row][column]
					displayNotes(noteValue, z)
				} else {
					if z == 1 && actualValue != 0 {
						print(" ", actualValue, " ")
					} else {
						print("   ")
					}
				}

				print("  ")
			}
			print(verticalBlockSeparator)
			print("\n")
		}
	}
	print("   =========================================================================\n")
}

//getBoardRowLetter returns the row letter for the specified row
func getBoardRowLetter(row int) string {
	return []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}[row]
}

//displayNotes shows the notes within the cells as a 3x3 grid
//which grid line (0-2) is represented by the rowIndex parameter
func displayNotes(value int, rowIndex int) {
	filler := "."

	for number := 1 + (3 * rowIndex); number <= 3+(3*rowIndex); number++ {
		digitBit := (int)(math.Pow(2, (float64)(number-1)))
		if value&digitBit > 0 {
			print(number)
		} else {
			print(filler)
		}
	}
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
				blockIndex := 3*(row/3) + column/3
				blockSolvedNumbers := getSolvedNumbersInNineCells(convertBlockToNineCells(board, blockIndex))
				exclusionNumbers := exclusions[row][column]
				newNotes[row][column] = 0x1ff ^ (rowSolvedNumbers | colSolvedNumbers | blockSolvedNumbers | exclusionNumbers)
			} else {
				//cell already solved
				newNotes[row][column] = 0
			}
		}
	}
	return newNotes
}

//convertRowToNineCells converts a row to a 9 cell array
func convertRowToNineCells(board [9][9]int, row int) [9]int {
	var cells [9]int
	for column := 0; column < 9; column++ {
		cells[column] = board[row][column]
	}
	return cells
}

//convertColumnToNineCells converts a column to a 9 cell array
func convertColumnToNineCells(board [9][9]int, column int) [9]int {
	var cells [9]int
	for row := 0; row < 9; row++ {
		cells[row] = board[row][column]
	}
	return cells
}

//convertBlockToNineCells takes a 3x3 block and converts it to a 9 cell array
//the blocks are indexed from 0 reading from top left to bottom right
//the first three rows form blocks 0-2, the second three rows form blocks 3-5
//and the final three rows for blocks 6-8
func convertBlockToNineCells(board [9][9]int, blockIndex int) [9]int {
	startRow := 3 * (blockIndex / 3)
	startColumn := 3 * (blockIndex % 3)
	var cells [9]int
	for k := 0; k < 9; k++ {
		cells[k] = board[startRow+k/3][startColumn+k%3]
	}
	return cells
}

//getSolvedNumbersInNineCells identifies the solved numbers in 9 cells
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

// IsSolved return true if the board is solved
func IsSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

//getCellRefsAsString prints the cells refs in a comma separated list
func getCellRefsAsString(cellRefs []CellRef) string {
	var s string
	for i := 0; i < len(cellRefs); i++ {
		if i > 0 {
			s = s + ", "
		}
		s = s + getBoardRowLetter(cellRefs[i].Row) + fmt.Sprintf("%d", 1+cellRefs[i].Column)
	}
	return s
}

//getNotesAsDigitString gets the notes as a single comma separated string
func getNotesAsDigitString(note int) string {
	s := ""
	for i := 1; i <= 9; i++ {
		if containsNumberInNote(note, i) {
			if len(s) > 0 {
				s = s + ","
			}
			s = s + fmt.Sprintf("%d", i)
		}
	}
	return s
}

//containsNumberInNote indicates whether the note contains the specified number
func containsNumberInNote(note int, number int) bool {
	numberAsBit := (int)(math.Pow(2, (float64)(number-1)))
	return note&numberAsBit == numberAsBit
}
