package sodukusolver

import "math"

//RelativeCellSolution represents a solution number and its index position within 9 cells
type RelativeCellSolution struct {
	Index  int
	Number int
}

// AbsoluteCellSolution represents a solution number in a cell and whether found in a row, column or block (location)
type AbsoluteCellSolution struct {
	Row      int
	Column   int
	Number   int
	Type     string
	Location string
}

//findNakedSingles identifies single value(s) in any cell on the board
func findNakedSingles(notes [9][9]int) []AbsoluteCellSolution {
	var solutions = []AbsoluteCellSolution{}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			note := 0x1ff & notes[row][column]

			if note == 0 {
				//no notes to check
				continue
			}

			//identify a single number in this cell's notes
			if countNumbersInNote(note) == 1 {
				solutions = append(solutions, AbsoluteCellSolution{row, column, getLowestNumberFromNote(note), "Naked Single", "Cell"})
			}
		}
	}
	return solutions
}

//findHiddenSingles identifies hidden single values in an cell on the board
func findHiddenSingles(notes [9][9]int) []AbsoluteCellSolution {
	var solutions = []AbsoluteCellSolution{}

	for row := 0; row < 9; row++ {
		rowSolutions := findHiddenSinglesInNineCells(convertRowToNineCells(notes, row))
		//convert the RelativeCellSolution to an AbsoluteCellSolution
		for i := 0; i < len(rowSolutions); i++ {
			s := rowSolutions[i]
			solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{row, s.Index, s.Number, "Hidden Single", "Cell"})
		}
	}

	for column := 0; column < 9; column++ {
		rowSolutions := findHiddenSinglesInNineCells(convertColumnToNineCells(notes, column))
		for i := 0; i < len(rowSolutions); i++ {
			s := rowSolutions[i]
			solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{s.Index, column, s.Number, "Hidden Single", "Cell"})
		}
	}

	return solutions
}

//appendAbsoluteCellSolution adds a cell solution if it does not already exist in the list
func appendAbsoluteCellSolution(solutions []AbsoluteCellSolution, solution AbsoluteCellSolution) []AbsoluteCellSolution {
	for i := 0; i < len(solutions); i++ {
		if solutions[i] == solution {
			return solutions
		}
	}

	return append(solutions, solution)
}

//findHiddenSinglesInNineCells finds all the hidden single values in 9 cells
func findHiddenSinglesInNineCells(notes [9]int) []RelativeCellSolution {
	var solutions = []RelativeCellSolution{}

	digitBit := 0x01
	for number := 1; number <= 9; number++ {
		//count the number of cells that contain this number (ie have bit set)
		numberCount := 0
		foundAt := -1
		for index := 0; index < 9; index++ {
			if notes[index]&digitBit == digitBit {
				numberCount++
				foundAt = index
			}
		}

		if numberCount == 1 {
			solutions = append(solutions, RelativeCellSolution{foundAt, number})
		}

		//shift the bit to test for the next number (bit)
		digitBit = digitBit << 1
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
