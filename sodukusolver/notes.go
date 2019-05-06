package sodukusolver

import "math"

//isNumberSet indicates whether the specified number is set in the note
func isNumberSet(note int, number int) bool {
	numberAsBit := (int)(math.Pow(2, (float64)(number-1)))
	return note&numberAsBit == numberAsBit
}

//setNumber sets the appropriate bit
func setNumber(number int) int {
	return (int)(math.Pow(2, (float64)(number-1)))
}

//getCommonNumberCount returns the number of common digits (bits set)
func getCommonNumberCount(number1 int, number2 int) int {
	count := 0
	for digit := 0; digit < 9; digit++ {
		if (number1 & number2 & 0x01) == 0x01 {
			count++
		}
		number1 = number1 >> 1
		number2 = number2 >> 1
	}

	return count
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

//isSameColumnInBlock indicates whether the cell indexes are in the
//same column for the cells of a block
func isSameColumnInBlock(cellIndexes []int) bool {
	if len(cellIndexes) < 2 {
		return false
	}
	for i := 0; i < len(cellIndexes)-1; i++ {
		if (cellIndexes[i] % 3) != (cellIndexes[i+1] % 3) {
			return false
		}
	}
	return true
}

//isSameRowInBlock indicates whether the cell indexes are in the
//same row for the cells of a block
func isSameRowInBlock(cellIndexes []int) bool {
	if len(cellIndexes) < 2 {
		return false
	}
	for i := 0; i < len(cellIndexes)-1; i++ {
		if (cellIndexes[i] / 3) != (cellIndexes[i+1] / 3) {
			return false
		}
	}
	return true
}
