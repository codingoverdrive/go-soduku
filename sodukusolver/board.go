package sodukusolver

import (
	"fmt"
	"math"
)

//PrintSolutionSteps prints the solution steps
//showBoard determines whether the board should be displayed (at each step)
func PrintSolutionSteps(solution Solution, showBoard bool) {
	for i := 0; i < len(solution.Steps); i++ {
		if i == 0 {
			//show the initial state of the board
			print("Initial Soduku board\n")
			PrintBoard(solution.InitialBoard, [9][9]int{}, false)
			print("\n")
			print("With Notes\n")
			PrintBoard(solution.InitialBoard, solution.InitialNotes, true)
			print("\n")
		}
		//provide the solution details for each step
		step := solution.Steps[i]
		if showBoard && i > 0 {
			//show the board again, but with the notes populated as well
			print("\n")
			PrintBoard(step.Board, step.Notes, true)
			print("\n")
		}
		print("Step ", fmt.Sprintf("%2d", i+1), " ", step.Description, "\n")
	}
	print("\n")
}

//PrintBoard outputs the current state of the board
//showNotes determines whether the notes are also displayed
func PrintBoard(board [9][9]int, notes [9][9]int, showNotes bool) {
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

//getBoardRowLetter returns the row letter for the specified row
func getBoardRowLetter(row int) string {
	return []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}[row]
}

//getCellRefsAsString prints the cells refs in a comma separated list
func getCellRefsAsString(cellRefs []CellRef) string {
	var s string
	for i := 0; i < len(cellRefs); i++ {
		if i > 0 {
			s = s + ", "
		}
		s = s + getBoardRowLetter(cellRefs[i].row) + fmt.Sprintf("%d", 1+cellRefs[i].column)
	}
	return s
}

//getNotesAsDigitString gets the notes as a single comma separated string
func getNotesAsDigitString(note int) string {
	s := ""
	for i := 1; i <= 9; i++ {
		if isNumberSet(note, i) {
			if len(s) > 0 {
				s = s + ","
			}
			s = s + fmt.Sprintf("%d", i)
		}
	}
	return s
}
