package sodukusolver

import "math"

//RelativeCellSolution represents a solution number and its index position within 9 cells
type RelativeCellSolution struct {
	index  int
	number int
}

// AbsoluteCellSolution represents a solution number in a cell and whether found in a row, column or block (location)
type AbsoluteCellSolution struct {
	row      int
	column   int
	number   int
	strategy string
	location string
}

//RelativeCellSolutions represents a solution number and its index position within 9 cells
type RelativeCellSolutions struct {
	indexes []int
	number  int
}

//CellRef defines a coordinate for a cell
type CellRef struct {
	row    int
	column int
}

//CellExclusion defines exclusions identified by a strategy
type CellExclusion struct {
	number     int
	matches    []CellRef
	exclusions []CellRef
	strategy   string
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

	//search the rows
	for row := 0; row < 9; row++ {
		rowSolutions := findHiddenSinglesInNineCells(convertRowToNineCells(notes, row))
		for i := 0; i < len(rowSolutions); i++ {
			s := rowSolutions[i]
			solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{row, s.index, s.number, "Hidden Single", "Cell"})
		}
	}

	//search the columns
	for column := 0; column < 9; column++ {
		rowSolutions := findHiddenSinglesInNineCells(convertColumnToNineCells(notes, column))
		for i := 0; i < len(rowSolutions); i++ {
			s := rowSolutions[i]
			solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{s.index, column, s.number, "Hidden Single", "Cell"})
		}
	}

	//search the blocks
	for block := 0; block < 9; block++ {
		rowSolutions := findHiddenSinglesInNineCells(convertBlockToNineCells(notes, block))
		for i := 0; i < len(rowSolutions); i++ {
			s := rowSolutions[i]

			startRow := 3 * (block / 3)
			row := startRow + s.index/3
			startColumn := 3 * (block % 3)
			column := startColumn + s.index%3

			solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{row, column, s.number, "Hidden Single", "Cell"})
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

func findNakedPairs(notes [9][9]int) []CellExclusion {
	var cellExclusions = []CellExclusion{}
	//search rows for naked pairs
	for row := 0; row < 9; row++ {
		cells := convertRowToNineCells(notes, row)
		cellSolutions := findNakedPairInNineCells(cells)
		if len(cellSolutions) == 0 {
			continue
		}

		exclusions := []int{}
		//remove the same notes from other cells in the row
		for i := 0; i < len(cellSolutions); i++ {
			s := cellSolutions[i]
			for column := 0; column < 9; column++ {
				//skip the cells that have the pairs in
				if contains(s.indexes, column) {
					continue
				}

				//check whether this note has any number from the pair
				if notes[row][column]&s.number > 0 {
					exclusions = append(exclusions, column)
				}
			}

			//only return this exclusion solution if >0 cells can have exclusions applied
			if len(exclusions) > 0 {
				matchRefs := []CellRef{}
				for i := 0; i < len(s.indexes); i++ {
					matchRefs = append(matchRefs, CellRef{row, s.indexes[i]})
				}
				exclRefs := []CellRef{}
				for i := 0; i < len(exclusions); i++ {
					exclRefs = append(exclRefs, CellRef{row, exclusions[i]})
				}
				cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, exclRefs, "Naked Pairs"})
			}
		}
	}

	//search column for naked pairs
	for column := 0; column < 9; column++ {
		cells := convertColumnToNineCells(notes, column)
		cellSolutions := findNakedPairInNineCells(cells)
		if len(cellSolutions) == 0 {
			continue
		}

		exclusions := []int{}
		//remove the same notes from other cells in the row
		for i := 0; i < len(cellSolutions); i++ {
			s := cellSolutions[i]
			for row := 0; row < 9; row++ {
				//skip the cells that have the pairs in
				if contains(s.indexes, row) {
					continue
				}

				//check whether this note has any number from the pair
				if notes[row][column]&s.number > 0 {
					exclusions = append(exclusions, row)
				}
			}

			//only return this exclusion solution if >0 cells can have exclusions applied
			if len(exclusions) > 0 {
				matchRefs := []CellRef{}
				for i := 0; i < len(s.indexes); i++ {
					matchRefs = append(matchRefs, CellRef{s.indexes[i], column})
				}
				exclRefs := []CellRef{}
				for i := 0; i < len(exclusions); i++ {
					exclRefs = append(exclRefs, CellRef{exclusions[i], column})
				}
				cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, exclRefs, "Naked Pairs"})
			}
		}
	}

	//search block for naked pairs
	for block := 0; block < 9; block++ {
		cells := convertBlockToNineCells(notes, block)
		cellSolutions := findNakedPairInNineCells(cells)
		if len(cellSolutions) == 0 {
			continue
		}

		rowOffset := 3 * (block / 3)
		columnOffset := 3 * (block % 3)

		exclusions := []int{}
		//remove the same notes from other cells in the row
		for i := 0; i < len(cellSolutions); i++ {
			s := cellSolutions[i]
			for cellIndex := 0; cellIndex < 9; cellIndex++ {

				//skip the cells that have the pairs in
				if contains(s.indexes, cellIndex) {
					continue
				}

				row := rowOffset + cellIndex/3
				column := columnOffset + cellIndex%3

				//check whether this note has any number from the pair
				if notes[row][column]&s.number > 0 {
					exclusions = append(exclusions, cellIndex)
				}
			}

			//only return this exclusion solution if >0 cells can have exclusions applied
			if len(exclusions) > 0 {
				matchRefs := []CellRef{}
				for i := 0; i < len(s.indexes); i++ {
					row := rowOffset + s.indexes[i]/3
					column := columnOffset + s.indexes[i]%3
					matchRefs = append(matchRefs, CellRef{row, column})
				}
				exclRefs := []CellRef{}
				for i := 0; i < len(exclusions); i++ {
					row := rowOffset + exclusions[i]/3
					column := columnOffset + exclusions[i]%3
					exclRefs = append(exclRefs, CellRef{row, column})
				}
				cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, exclRefs, "Naked Pairs"})
			}
		}
	}

	return cellExclusions
}

//findNakedPairInNineCells finds naked pairs in nine cells
func findNakedPairInNineCells(cells [9]int) []RelativeCellSolutions {
	var solutions = []RelativeCellSolutions{}
	var pairs = []int{}
	for i := 0; i < 9; i++ {
		if contains(pairs, cells[i]) || countNumbersInNote(cells[i]) != 2 {
			continue
		}

		//avoid process this pair again
		pairs = append(pairs, cells[i])

		//found a cell with a pair of note numbers
		var indexes = []int{i}

		//look for another pair match
		for k := i + 1; k < 9; k++ {
			if cells[i] != cells[k] {
				continue
			}
			indexes = append(indexes, k)
		}

		//discard this solutiom if more than two pairs are found
		if len(indexes) == 2 {
			solutions = append(solutions, RelativeCellSolutions{indexes, cells[i]})
		}
	}
	return solutions
}

//contains indicates whether the array contains the specified value
func contains(array []int, value int) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
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
