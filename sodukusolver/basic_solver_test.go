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

	assert.Equal(t, []CellSolution{}, findNakedSingles(notes1), "No solution expected")

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

	expected := []CellSolution{
		CellSolution{1, 2, 1, "Naked Single", "Cell"},
		CellSolution{7, 6, 3, "Naked Single", "Cell"},
	}
	assert.Equal(t, expected, findNakedSingles(notes2), "Two solutions expected")
}
