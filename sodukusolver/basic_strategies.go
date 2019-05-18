package sodukusolver

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
	indexes  []int
	number   int
	location string
}

//CellRef defines a coordinate for a cell
type CellRef struct {
	row    int
	column int
}

//CellExclusion defines exclusions identified by a strategy
type CellExclusion struct {
	number       int
	matches      []CellRef
	removeNumber int
	exclusions   []CellRef
	strategy     string
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
		columnSolutions := findHiddenSinglesInNineCells(convertColumnToNineCells(notes, column))
		for i := 0; i < len(columnSolutions); i++ {
			s := columnSolutions[i]
			solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{s.index, column, s.number, "Hidden Single", "Cell"})
		}
	}

	//search the blocks
	for block := 0; block < 9; block++ {
		blockSolutions := findHiddenSinglesInNineCells(convertBlockToNineCells(notes, block))
		for i := 0; i < len(blockSolutions); i++ {
			s := blockSolutions[i]

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

//findNakedPairExclusions
func findNakedPairExclusions(notes [9][9]int) []CellExclusion {
	strategyName := "Naked Pairs"
	var cellExclusions = []CellExclusion{}

	//search rows for naked pairs
	for row := 0; row < 9; row++ {
		cells := convertRowToNineCells(notes, row)
		cellSolutions := findNakedPairsInNineCells(cells)
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
				cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, s.number, exclRefs, strategyName})
			}
		}
	}

	//search column for naked pairs
	for column := 0; column < 9; column++ {
		cells := convertColumnToNineCells(notes, column)
		cellSolutions := findNakedPairsInNineCells(cells)
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
				cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, s.number, exclRefs, strategyName})
			}
		}
	}

	//search block for naked pairs
	for block := 0; block < 9; block++ {
		cells := convertBlockToNineCells(notes, block)
		cellSolutions := findNakedPairsInNineCells(cells)
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
				cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, s.number, exclRefs, strategyName})
			}
		}
	}

	return cellExclusions
}

//findNakedPairsInNineCells finds naked pairs in nine cells
func findNakedPairsInNineCells(notes [9]int) []RelativeCellSolutions {
	var solutions = []RelativeCellSolutions{}

	//keep track of the pairs marked to be ignored
	ignorePairs := []int{}

	//scan all the cells looking for matching (hidden) pairs
	for x := 0; x < 9; x++ {
		//skip cells that don't have at least two numbers in them
		if notes[x] == 0 || countNumbersInNote(notes[x]) != 2 {
			continue
		}

		for y := x + 1; y < 9; y++ {
			if notes[y] != 0 && getCommonNumberCount(notes[x], notes[y]) == 2 {
				if notes[x] != notes[y] {
					continue
				}

				commonDigits := notes[x] & notes[y]

				//don't process a pair that has already been marked for ignoring
				if contains(ignorePairs, commonDigits) {
					break
				}

				//check that this pair doesn't exist in any other cell
				foundInThirdCell := false
				for z := y + 1; z < 9; z++ {
					if notes[z]&commonDigits == commonDigits {
						foundInThirdCell = true
						break
					}
				}

				if foundInThirdCell {
					ignorePairs = append(ignorePairs, commonDigits)
					break
				} else {
					//add this pair as a solution
					solutions = append(solutions, RelativeCellSolutions{[]int{x, y}, commonDigits, "Cell"})
				}
			}
		}
	}
	return solutions
}

//findHiddenPairExclusions finds (notes) exclusions from hidden pairs
func findHiddenPairExclusions(notes [9][9]int) []CellExclusion {
	strategyName := "Hidden Pairs"
	var cellExclusions = []CellExclusion{}

	//search rows for hidden pairs
	for row := 0; row < 9; row++ {
		cells := convertRowToNineCells(notes, row)
		cellSolutions := findHiddenPairsInNineCells(cells)
		if len(cellSolutions) == 0 {
			continue
		}

		//remove the non-pair numbers from the pair of cells
		for i := 0; i < len(cellSolutions); i++ {
			s := cellSolutions[i]

			removeNumber := 0
			matchRefs := []CellRef{}
			for i := 0; i < len(s.indexes); i++ {
				removeNumber = removeNumber | notes[row][s.indexes[i]]
				matchRefs = append(matchRefs, CellRef{row, s.indexes[i]})
			}
			removeNumber = removeNumber ^ s.number
			cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, removeNumber, matchRefs, strategyName})
		}
	}

	//search column for hidden pairs
	for column := 0; column < 9; column++ {
		cells := convertColumnToNineCells(notes, column)
		cellSolutions := findHiddenPairsInNineCells(cells)
		if len(cellSolutions) == 0 {
			continue
		}

		//remove the non-pair numbers from the pair of cells
		for i := 0; i < len(cellSolutions); i++ {
			s := cellSolutions[i]

			removeNumber := 0
			matchRefs := []CellRef{}
			for i := 0; i < len(s.indexes); i++ {
				removeNumber = removeNumber | notes[s.indexes[i]][column]
				matchRefs = append(matchRefs, CellRef{s.indexes[i], column})
			}
			removeNumber = removeNumber ^ s.number
			cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, removeNumber, matchRefs, strategyName})
		}
	}

	//search block for hidden pairs
	for block := 0; block < 9; block++ {
		cells := convertBlockToNineCells(notes, block)
		cellSolutions := findHiddenPairsInNineCells(cells)
		if len(cellSolutions) == 0 {
			continue
		}

		rowOffset := 3 * (block / 3)
		columnOffset := 3 * (block % 3)

		//remove the non-pair numbers from the pair of cells
		for i := 0; i < len(cellSolutions); i++ {
			s := cellSolutions[i]

			removeNumber := 0
			matchRefs := []CellRef{}
			for i := 0; i < len(s.indexes); i++ {
				row := rowOffset + s.indexes[i]/3
				column := columnOffset + s.indexes[i]%3
				removeNumber = removeNumber | notes[row][column]
				matchRefs = append(matchRefs, CellRef{row, column})
			}
			removeNumber = removeNumber ^ s.number
			cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, removeNumber, matchRefs, strategyName})
		}
	}

	return cellExclusions
}

//findHiddenPairsInNineCells finds hidden pairs in nine cells
func findHiddenPairsInNineCells(notes [9]int) []RelativeCellSolutions {
	var solutions = []RelativeCellSolutions{}

	//scan all the cells looking for (hidden) pairs
	for x := 0; x < 9; x++ {
		//skip cells that don't have at least two numbers in them
		if notes[x] == 0 || countNumbersInNote(notes[x]) < 2 {
			continue
		}

		//look for the second hidden pair
		for y := x + 1; y < 9; y++ {
			if notes[y] != 0 && getCommonNumberCount(notes[x], notes[y]) == 2 {
				//ignore naked pairs
				if notes[x] == notes[y] {
					continue
				}

				//identify the pair of numbers
				commonDigits := notes[x] & notes[y]

				//check that no other cells have either of the common digits
				addPair := true
				for z := 0; z < 9; z++ {
					//ignore the cell pairs already found
					if z == x || z == y {
						continue
					}

					if notes[z]&commonDigits > 0 {
						addPair = false
						break
					}
				}

				//add this pair as a solution
				if addPair {
					solutions = append(solutions, RelativeCellSolutions{[]int{x, y}, commonDigits, "Cell"})
				}

			}
		}
	}
	return solutions
}

//findPointingPairExclusions finds exclusions based on pointing pairs
func findPointingPairExclusions(notes [9][9]int) []CellExclusion {
	var cellExclusions = []CellExclusion{}

	for block := 0; block < 9; block++ {
		blockNotes := convertBlockToNineCells(notes, block)
		pairSolutions := findPointingPairsInNineCellBlock(blockNotes)
		if len(pairSolutions) == 0 {
			continue
		}

		rowOffset := 3 * (block / 3)
		columnOffset := 3 * (block % 3)

		//find exclusions for each of the pointing pairs
		for i := 0; i < len(pairSolutions); i++ {
			s := pairSolutions[i]

			switch s.location {
			case "Row":
				{
					row := rowOffset + s.indexes[0]/3

					//convert the matching pairs into absolute cell refs
					matchRefs := []CellRef{}
					for cellIndex := 0; cellIndex < len(s.indexes); cellIndex++ {
						column := columnOffset + s.indexes[cellIndex]%3
						matchRefs = append(matchRefs, CellRef{row, column})
					}

					//find the exclusions pointed to by the pairs
					exclRefs := []CellRef{}
					for columnIndex := 0; columnIndex < 9; columnIndex++ {
						//skip matching columns within the 3x3 block
						cellRef := CellRef{row, columnIndex}
						if containsCellRef(matchRefs, cellRef) {
							continue
						}
						if notes[row][columnIndex]&s.number == s.number {
							exclRefs = append(exclRefs, cellRef)
						}
					}
					//only add if there are any exclusions
					if len(exclRefs) > 0 {
						cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, s.number, exclRefs, "Pointing Pairs"})
					}
				}
			case "Column":
				{
					column := columnOffset + s.indexes[0]%3

					//convert the matching pairs into absolute cell refs
					matchRefs := []CellRef{}
					for cellIndex := 0; cellIndex < len(s.indexes); cellIndex++ {
						row := rowOffset + s.indexes[cellIndex]/3
						matchRefs = append(matchRefs, CellRef{row, column})
					}

					//find the exclusions pointed to by the pairs
					exclRefs := []CellRef{}
					for rowIndex := 0; rowIndex < 9; rowIndex++ {
						//skip matching rows within the 3x3 block
						cellRef := CellRef{rowIndex, column}
						if containsCellRef(matchRefs, cellRef) {
							continue
						}
						if notes[rowIndex][column]&s.number == s.number {
							exclRefs = append(exclRefs, cellRef)
						}
					}
					//only add if there are any exclusions
					if len(exclRefs) > 0 {
						cellExclusions = append(cellExclusions, CellExclusion{s.number, matchRefs, s.number, exclRefs, "Pointing Pairs"})
					}
				}
			}
		}
	}

	return cellExclusions
}

//findPointingPairsInNineCellBlock finds pointing pairs in the block that are
//aligned in a single column or row
func findPointingPairsInNineCellBlock(notes [9]int) []RelativeCellSolutions {
	var solutions = []RelativeCellSolutions{}

	//for each digit, identify the cell indexes that contains that digit/number
	digitCells := [9][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	for digit := 0; digit < 9; digit++ {
		for cellIndex := 0; cellIndex < 9; cellIndex++ {
			if isNumberSet(notes[cellIndex], digit+1) {
				digitCells[digit] = append(digitCells[digit], cellIndex)
			}
		}

		//skip digits where there are not enough or too many cells
		if len(digitCells[digit]) <= 1 || len(digitCells[digit]) > 3 {
			continue
		}

		//check to see whether the cell indexes are aligned in a column or row
		if isSameColumnInBlock(digitCells[digit]) {
			solutions = append(solutions, RelativeCellSolutions{digitCells[digit], setNumber(digit + 1), "Column"})
			continue
		}
		if isSameRowInBlock(digitCells[digit]) {
			solutions = append(solutions, RelativeCellSolutions{digitCells[digit], setNumber(digit + 1), "Row"})
			continue
		}
	}
	return solutions
}

//findBoxLineReductionExclusions finds exclusions within blocks based on box line reductions
func findBoxLineReductionExclusions(notes [9][9]int) []CellExclusion {
	var solutions = []CellExclusion{}

	for row := 0; row < 9; row++ {
		cells := convertRowToNineCells(notes, row)
		pairSolutions := findBoxPairsInNineCellLine(cells)
		if len(pairSolutions) == 0 {
			continue
		}

		for i := 0; i < len(pairSolutions); i++ {
			s := pairSolutions[i]

			//convert the matching pairs into absolute cell refs
			matchRefs := []CellRef{}
			for cellIndex := 0; cellIndex < len(s.indexes); cellIndex++ {
				matchRefs = append(matchRefs, CellRef{row, s.indexes[cellIndex]})
			}

			//determine which block the pair sits in
			column := s.indexes[0]
			blockIndex := 3*(row/3) + column/3
			blockCells := convertBlockToNineCells(notes, blockIndex)

			rowOffset := 3 * (blockIndex / 3)
			columnOffset := 3 * (blockIndex % 3)

			//and remove the same number from any other cells in the block
			exclRefs := []CellRef{}
			for cellIndex := 0; cellIndex < 9; cellIndex++ {

				//skip cell if it doesn't contain the number pair
				if blockCells[cellIndex]&s.number == 0 {
					continue
				}

				//skip if this is one of the pair cells
				cellR := CellRef{rowOffset + cellIndex/3, columnOffset + cellIndex%3}
				if containsCellRef(matchRefs, cellR) {
					continue
				}

				exclRefs = append(exclRefs, cellR)
			}

			if len(exclRefs) > 0 {
				solutions = append(solutions, CellExclusion{s.number, matchRefs, s.number, exclRefs, "Box Line Reduction"})
			}
		}
	}

	for column := 0; column < 9; column++ {
		cells := convertColumnToNineCells(notes, column)
		pairSolutions := findBoxPairsInNineCellLine(cells)
		if len(pairSolutions) == 0 {
			continue
		}

		for i := 0; i < len(pairSolutions); i++ {
			s := pairSolutions[i]

			//convert the matching pairs into absolute cell refs
			matchRefs := []CellRef{}
			for cellIndex := 0; cellIndex < len(s.indexes); cellIndex++ {
				matchRefs = append(matchRefs, CellRef{s.indexes[cellIndex], column})
			}

			//determine which block the pair sits in
			row := s.indexes[0]
			blockIndex := 3*(row/3) + column/3
			blockCells := convertBlockToNineCells(notes, blockIndex)

			rowOffset := 3 * (blockIndex / 3)
			columnOffset := 3 * (blockIndex % 3)

			//and remove the same number from any other cells in the block
			exclRefs := []CellRef{}
			for cellIndex := 0; cellIndex < 9; cellIndex++ {

				//skip cell if it doesn't contain the number pair
				if blockCells[cellIndex]&s.number == 0 {
					continue
				}

				//skip if this is one of the pair cells
				cellR := CellRef{rowOffset + cellIndex/3, columnOffset + cellIndex%3}
				if containsCellRef(matchRefs, cellR) {
					continue
				}

				exclRefs = append(exclRefs, cellR)
			}

			if len(exclRefs) > 0 {
				solutions = append(solutions, CellExclusion{s.number, matchRefs, s.number, exclRefs, "Box Line Reduction"})
			}
		}
	}

	return solutions
}

//findBoxPairsInNineCellLine finds pairs in a row or column nine cell line
func findBoxPairsInNineCellLine(notes [9]int) []RelativeCellSolutions {
	var solutions = []RelativeCellSolutions{}

	//for each digit, identify the cell indexes that contains that digit/number
	digitCells := [9][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	for digit := 0; digit < 9; digit++ {
		for cellIndex := 0; cellIndex < 9; cellIndex++ {
			if isNumberSet(notes[cellIndex], digit+1) {
				digitCells[digit] = append(digitCells[digit], cellIndex)
			}
		}

		//skip digits where there are not enough or too many cells
		if len(digitCells[digit]) <= 1 || len(digitCells[digit]) > 3 {
			continue
		}

		//check whether the cell indexes are all in the same block
		blockIndex := digitCells[digit][0] / 3
		for i := 1; i < len(digitCells[digit]); i++ {
			//determine the block that this cell sits in
			bIndex := digitCells[digit][i] / 3
			if bIndex != blockIndex {
				blockIndex = -1
				break
			}
		}

		//cells in the same block?
		if blockIndex > -1 {
			solutions = append(solutions, RelativeCellSolutions{digitCells[digit], setNumber(digit + 1), "Box Pair"})
		}
	}
	return solutions
}
