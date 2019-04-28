package sodukusolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//used to test the row, column and block conversions to 9 cells
var invalidTestBoard = [9][9]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 0, 0, 0, 0, 0, 8},
	{7, 8, 9, 0, 0, 0, 0, 0, 7},
	{0, 0, 0, 1, 2, 3, 0, 0, 6},
	{0, 0, 0, 4, 5, 6, 0, 0, 5},
	{0, 0, 0, 7, 8, 9, 0, 0, 4},
	{0, 0, 0, 0, 0, 0, 0, 0, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 2},
	{9, 8, 7, 6, 5, 4, 3, 2, 1},
}

func Test_convertRowToNineCells(t *testing.T) {
	//test the first row
	cells := convertRowToNineCells(invalidTestBoard, 0)
	for column := 0; column < 9; column++ {
		assert.Equal(t, (1 + column), cells[column], "Unexpected cell value")
	}
	//test the last row
	cells = convertRowToNineCells(invalidTestBoard, 8)
	for column := 0; column < 9; column++ {
		assert.Equal(t, (9 - column), cells[column], "Unexpected cell value")
	}
}

func Test_convertColumnToNineCells(t *testing.T) {
	cells := convertColumnToNineCells(invalidTestBoard, 8)
	for row := 0; row < 9; row++ {
		assert.Equal(t, (9 - row), cells[row], "Unexpected cell value")
	}
}

func Test_convertBlockToNineCells(t *testing.T) {
	cells := convertBlockToNineCells(invalidTestBoard, 0)
	for i := 0; i < 9; i++ {
		assert.Equal(t, (1 + i), cells[i], "Unexpected cell value")
	}
	cells = convertBlockToNineCells(invalidTestBoard, 4)
	for i := 0; i < 9; i++ {
		assert.Equal(t, (1 + i), cells[i], "Unexpected cell value")
	}
}

func Test_getSolvedNumbersInNineCells(t *testing.T) {
	assert.Equal(t, 0x100, getSolvedNumbersInNineCells([9]int{0, 0, 0, 9, 0, 0, 0, 0, 0}), "Unexpected notes value")
	assert.Equal(t, 0x07, getSolvedNumbersInNineCells([9]int{1, 0, 0, 2, 0, 0, 3, 0, 0}), "Unexpected notes value")
	assert.Equal(t, 0x06, getSolvedNumbersInNineCells([9]int{0, 0, 0, 2, 0, 0, 3, 0, 0}), "Unexpected notes value")
	assert.Equal(t, 0x1ff, getSolvedNumbersInNineCells([9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), "Unexpected notes value")
}

func Test_recalculateBoardNotes(t *testing.T) {
	board1 := [9][9]int{
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
	notes1 := recalculateBoardNotes(board1)
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			assert.Equal(t, 0x1ff, notes1[row][column], "Unexpected notes value")
		}

	}

	board2 := [9][9]int{
		{9, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 9, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}
	notes2 := recalculateBoardNotes(board2)
	assert.Equal(t, 0xff, notes2[0][1], "Unexpected notes value")
	assert.Equal(t, 0xff, notes2[1][8], "Unexpected notes value")
	assert.Equal(t, 0xff, notes2[5][4], "Unexpected notes value")

}
