package sodukusolver

import "math"

// CellSolution represents a solution number in a cell and whether found in a row, column or block (location)
type CellSolution struct {
	Row      int
	Column   int
	Number   int
	Type     string
	Location string
}

// identifies single value(s) in any cell on the board
func findNakedSingles(notes [9][9]int) []CellSolution {
	var solutions = []CellSolution{}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			note := 0x1ff & notes[row][column]

			if note == 0 {
				//no notes to check
				continue
			}

			//identify a single number in this cell's notes
			if countNumbersInNote(note) == 1 {
				solutions = append(solutions, CellSolution{row, column, getLowestNumberFromNote(note), "Naked Single", "Cell"})
			}
		}
	}
	return solutions
}

// countNumbersInNote indicates how many numbers (bits) are set in the note
func countNumbersInNote(note int) int {
	count := 0
	for i := 0; i < 9; i++ {
		if note&1 == 1 {
			count++
		}
		note = note >> 1
	}
	return count
}

// getLowestNumberFromNote returns the lowest set number (bit) from the note
func getLowestNumberFromNote(note int) int {
	for k := 1; k <= 9; k++ {
		digitAsBit := (int)(math.Pow(2, (float64)(k-1)))
		if note^digitAsBit == 0 {
			return k
		}
	}
	return 0
}
