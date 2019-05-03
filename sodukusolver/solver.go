package sodukusolver

import (
	"fmt"
	"math"
	"time"
)

//the board, notes and exclusions are represented as a 9x9 int array
//with zero denoting no solution value in a board, and no notes in the notes array
//zero in the exclusions means no removals of notes from the notes value

//notes and exclusions are represented as 9bit integers with a 1 in each bit denoting a note value
//0x07 means that the notes values are 1,2, and 3 (00000111 in binary)

//SolutionStep describes a single solution step
type SolutionStep struct {
	Strategy    string
	Description string
	Board       [9][9]int
	Notes       [9][9]int
}

//Solution describes how the Soduku puzzle was solved
type Solution struct {
	InitialBoard [9][9]int
	InitialNotes [9][9]int
	Solved       bool
	Elapsed      time.Duration
	Steps        []SolutionStep
	//unexported and used within the solver function to maintain state
	board      [9][9]int
	notes      [9][9]int
	exclusions [9][9]int
}

//StrategyResult indicates the success or failure of a strategy
//and how many solutions were found for the strategy
type StrategyResult struct {
	success   bool
	solutions int
}

//CellStrategy a solver strategy that returns an array of AbsoluteCellSolutions
type CellStrategy func([9][9]int) []AbsoluteCellSolution

//applyCellSolutionStrategy applies a cell strategy and returns a result
func applyCellSolutionStrategy(strategy CellStrategy, solution *Solution) StrategyResult {
	startTime := time.Now()
	solutions := strategy(solution.notes)
	elapsed := time.Since(startTime)
	if len(solutions) > 0 {
		for i := 0; i < len(solutions); i++ {
			s := solutions[i]
			description := "Found " + s.strategy + " [" + fmt.Sprintf("%d", s.number) + "] at " + s.location + " " + getBoardRowLetter(s.row) + fmt.Sprintf("%d", 1+s.column)

			//update the board and notes
			solution.board[s.row][s.column] = s.number
			solution.notes = recalculateBoardNotes(solution.board, solution.exclusions)

			step := SolutionStep{Description: description, Board: solution.board, Notes: solution.notes}
			solution.Steps = append(solution.Steps, step)
		}
	}

	solution.Elapsed = solution.Elapsed + elapsed
	return StrategyResult{len(solutions) > 0, len(solutions)}
}

//ExclusionStrategy a solver strategy the eliminates/excludes notes that cannot be a cell solution
type ExclusionStrategy func([9][9]int) []CellExclusion

//applyCellExclusionStrategy applies a cell notes exclusion strategy and returns a result
func applyCellExclusionStrategy(strategy ExclusionStrategy, solution *Solution) StrategyResult {
	startTime := time.Now()
	exclusionSolutions := findNakedPairs(solution.notes)
	elapsed := time.Since(startTime)
	if len(exclusionSolutions) > 0 {
		for i := 0; i < len(exclusionSolutions); i++ {
			s := exclusionSolutions[i]
			matchingPair := "[" + getNotesAsDigitString(s.number) + "]"
			description := "Found " + s.strategy + " " + matchingPair + " in " + getCellRefsAsString(s.matches)
			description2 := "Removing " + matchingPair + " from " + getCellRefsAsString(s.exclusions)
			exclRefs := s.exclusions
			for k := 0; k < len(exclRefs); k++ {
				solution.exclusions[exclRefs[k].row][exclRefs[k].column] = solution.exclusions[exclRefs[k].row][exclRefs[k].column] | s.number
			}

			//update the board and notes
			solution.notes = recalculateBoardNotes(solution.board, solution.exclusions)

			step := SolutionStep{Description: description + ", " + description2, Board: solution.board, Notes: solution.notes}
			solution.Steps = append(solution.Steps, step)
		}
	}
	solution.Elapsed = solution.Elapsed + elapsed
	return StrategyResult{len(exclusionSolutions) > 0, len(exclusionSolutions)}
}

//initialiseSolution initialises the solution
func initialiseSolution(initialBoard [9][9]int) Solution {
	solution := Solution{InitialBoard: initialBoard, board: initialBoard}

	//initialise the notes
	solution.InitialNotes = recalculateBoardNotes(solution.board, solution.exclusions)
	solution.notes = solution.InitialNotes

	return solution
}

//SolveBoard solves the board
func SolveBoard(board [9][9]int) Solution {
	solution := initialiseSolution(board)

	// print("\nInitial Board\n")
	// printBoard(false)
	// print("\nWith Notes\n")
	// printBoard(true)
	// print("\n")

	//keep applying the different solution strategies until the puzzle is solved
	for {
		//findNakedSingles
		result := applyCellSolutionStrategy(findNakedSingles, &solution)
		if result.success {
			continue
		}

		//findHiddenSingles
		result = applyCellSolutionStrategy(findHiddenSingles, &solution)
		if result.success {
			continue
		}

		//find naked pairs
		result = applyCellExclusionStrategy(findNakedPairs, &solution)
		if result.success {
			continue
		}

		break
	}

	solution.Solved = isSolved(solution.board)

	return solution
}

//recalculateBoardNotes takes the board and generates a new set of notes
//from the current solved cells
func recalculateBoardNotes(board [9][9]int, exclusions [9][9]int) [9][9]int {
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

// isSolved return true if the board is solved
func isSolved(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
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

//containsNumberInNote indicates whether the note contains the specified number
func containsNumberInNote(note int, number int) bool {
	numberAsBit := (int)(math.Pow(2, (float64)(number-1)))
	return note&numberAsBit == numberAsBit
}
