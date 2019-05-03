package sodukusolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSingleNumberSet(t *testing.T) {
	assert.Equal(t, 0, countNumbersInNote(0), "Failed in test for single digit")
	assert.Equal(t, 1, countNumbersInNote(1), "Failed in test for single digit")
	assert.Equal(t, 1, countNumbersInNote(2), "Failed in test for single digit")
	assert.Equal(t, 1, countNumbersInNote(4), "Failed in test for single digit")
	assert.Equal(t, 2, countNumbersInNote(5), "Failed in test for single digit")
	assert.Equal(t, 3, countNumbersInNote(7), "Failed in test for single digit")
	assert.Equal(t, 9, countNumbersInNote(0x1ff), "Failed in test for single digit")
}

func Test_getLowestNumberFromNote(t *testing.T) {
	assert.Equal(t, 0, getLowestNumberFromNote(0), "Failed to get correct number from note")
	assert.Equal(t, 1, getLowestNumberFromNote(1), "Failed to get correct number from note")
	assert.Equal(t, 2, getLowestNumberFromNote(2), "Failed to get correct number from note")
	assert.Equal(t, 3, getLowestNumberFromNote(4), "Failed to get correct number from note")
	assert.Equal(t, 9, getLowestNumberFromNote(0x100), "Failed to get correct number from note")
}

func Test_findNakedSingles(t *testing.T) {
	notes1 := [9][9]int{
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
	assert.Equal(t, []AbsoluteCellSolution{}, findNakedSingles(notes1), "No solution expected")

	notes2 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected := []AbsoluteCellSolution{
		AbsoluteCellSolution{1, 2, 1, "Naked Single", "Cell"},
		AbsoluteCellSolution{7, 6, 3, "Naked Single", "Cell"},
	}
	assert.Equal(t, expected, findNakedSingles(notes2), "Two solutions expected")
}

func Test_findHiddenSinglesInNineCells(t *testing.T) {
	cells1 := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, []RelativeCellSolution{}, findHiddenSinglesInNineCells(cells1), "No solutions expected")

	cells2 := [9]int{0, 0, 3, 0, 0, 0, 0, 4, 6}
	expected2 := []RelativeCellSolution{
		{2, 1},
	}
	assert.Equal(t, expected2, findHiddenSinglesInNineCells(cells2), "One solution expected")

	cells3 := [9]int{12, 0, 2, 0, 0, 0, 1, 4, 6}
	expected3 := []RelativeCellSolution{
		{6, 1},
		{0, 4},
	}
	assert.Equal(t, expected3, findHiddenSinglesInNineCells(cells3), "Two solutions expected")

	cells4 := [9]int{0, 6, 0, 0, 0, 0, 2, 0, 0}
	expected4 := []RelativeCellSolution{
		{1, 3},
	}
	assert.Equal(t, expected4, findHiddenSinglesInNineCells(cells4), "One solution expected")
}

func Test_findHiddenSingles(t *testing.T) {
	notes1 := [9][9]int{
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
	assert.Equal(t, []AbsoluteCellSolution{}, findHiddenSingles(notes1), "No solution expected")

	notes2 := [9][9]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 6, 2, 0, 0, 0, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 2, 0, 0, 0, 2, 1, 1},
		{0, 2, 0, 0, 0, 0, 0, 2, 0},
	}
	expected2 := []AbsoluteCellSolution{
		AbsoluteCellSolution{7, 7, 1, "Hidden Single", "Cell"},
		AbsoluteCellSolution{1, 1, 3, "Hidden Single", "Cell"},
		AbsoluteCellSolution{7, 0, 1, "Hidden Single", "Cell"},
	}
	assert.ElementsMatch(t, expected2, findHiddenSingles(notes2), "Two solutions expected")
}

func Test_findNakedPairInNineCells(t *testing.T) {
	cells1 := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, []RelativeCellSolutions{}, findNakedPairInNineCells(cells1), "No naked pairs expected")

	cells2 := [9]int{6, 0x101, 0, 6, 6, 0, 0, 0x101, 0}
	expected2 := []RelativeCellSolutions{RelativeCellSolutions{[]int{1, 7}, 0x101}}
	assert.Equal(t, expected2, findNakedPairInNineCells(cells2), "One naked pair expected")

	cells3 := [9]int{0, 0x101, 0, 6, 0, 6, 0, 0x101, 0}
	expected3 := []RelativeCellSolutions{
		RelativeCellSolutions{[]int{1, 7}, 0x101},
		RelativeCellSolutions{[]int{3, 5}, 6},
	}
	assert.ElementsMatch(t, expected3, findNakedPairInNineCells(cells3), "Two naked pairs expected")
}

func Test_findNakedPairs(t *testing.T) {
	notes1 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 3, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{4, 12, 0, 0, 0, 12, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected1 := []CellExclusion{
		CellExclusion{
			number:     12,
			matches:    []CellRef{CellRef{6, 1}, CellRef{6, 5}},
			exclusions: []CellRef{CellRef{6, 0}, CellRef{6, 7}},
			strategy:   "Naked Pairs"},
	}
	assert.Equal(t, expected1, findNakedPairs(notes1), "One naked pair expected")

	notes2 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 12, 0, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 12, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected2 := []CellExclusion{
		CellExclusion{
			number:     12,
			matches:    []CellRef{CellRef{2, 6}, CellRef{6, 6}},
			exclusions: []CellRef{CellRef{0, 6}, CellRef{7, 6}},
			strategy:   "Naked Pairs"},
	}
	assert.Equal(t, expected2, findNakedPairs(notes2), "One naked pair expected")

	notes3 := [9][9]int{
		{3, 0, 3, 0, 0, 0, 0, 0, 0},
		{4, 12, 12, 0, 0, 0, 0, 4, 0},
		{3, 8, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected3 := []CellExclusion{
		CellExclusion{
			number:     12,
			matches:    []CellRef{CellRef{1, 1}, CellRef{1, 2}},
			exclusions: []CellRef{CellRef{1, 0}, CellRef{1, 7}},
			strategy:   "Naked Pairs"},
		CellExclusion{
			number:     12,
			matches:    []CellRef{CellRef{1, 1}, CellRef{1, 2}},
			exclusions: []CellRef{CellRef{1, 0}, CellRef{2, 1}},
			strategy:   "Naked Pairs"},
		CellExclusion{
			number:     3,
			matches:    []CellRef{CellRef{0, 0}, CellRef{2, 0}},
			exclusions: []CellRef{CellRef{7, 0}},
			strategy:   "Naked Pairs"},
	}
	assert.ElementsMatch(t, expected3, findNakedPairs(notes3), "Three naked pairs expected")
}
