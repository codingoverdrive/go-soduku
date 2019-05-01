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
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 5, 0, 2, 0, 0, 0, 0},
		{8, 0, 0, 9, 0, 5, 0, 7, 0},
		{3, 7, 0, 5, 0, 2, 6, 0, 0},
		{6, 5, 0, 0, 0, 0, 0, 1, 8},
		{0, 0, 8, 4, 0, 6, 0, 3, 5},
		{0, 6, 0, 1, 0, 8, 0, 0, 7},
		{0, 0, 0, 0, 4, 0, 9, 0, 3},
		{0, 0, 0, 0, 0, 0, 0, 6, 0},
	}
	expectedNotes := [9][9]int{
		{0x109, 0x0, 0x12d, 0xe4, 0xe5, 0x4d, 0x9d, 0x198, 0x129},
		{0x0, 0x10d, 0x0, 0xa4, 0x0, 0xd, 0x8d, 0x188, 0x129},
		{0x0, 0xd, 0x2d, 0x0, 0x25, 0x0, 0xf, 0x0, 0x2b},
		{0x0, 0x0, 0x109, 0x0, 0x181, 0x0, 0x0, 0x108, 0x108},
		{0x0, 0x0, 0x10a, 0x44, 0x144, 0x144, 0x4a, 0x0, 0x0},
		{0x103, 0x101, 0x0, 0x0, 0x141, 0x0, 0x42, 0x0, 0x0},
		{0x11a, 0x0, 0x10e, 0x0, 0x114, 0x0, 0x1a, 0x1a, 0x0},
		{0x13, 0x81, 0x43, 0x62, 0x0, 0x40, 0x0, 0x92, 0x0},
		{0x11b, 0x18d, 0x14f, 0x46, 0x154, 0x144, 0x9b, 0x0, 0xb},
	}
	notes2 := recalculateBoardNotes(board2)
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			assert.Equal(t, expectedNotes[row][column], notes2[row][column], "Unexpected notes value")
		}
	}
}

func Test_containsNumberInNote(t *testing.T) {
	type TestData struct {
		number   int
		digit    int
		expected bool
	}

	var testData = []TestData{
		TestData{0x1ff, 1, true},
		TestData{0x1ff, 2, true},
		TestData{0x1ff, 3, true},
		TestData{0x1ff, 4, true},
		TestData{0x1ff, 5, true},
		TestData{0x1ff, 6, true},
		TestData{0x1ff, 7, true},
		TestData{0x1ff, 8, true},
		TestData{0x1ff, 9, true},
		TestData{0x01, 1, true},
		TestData{0x05, 1, true},
		TestData{0x05, 3, true},
	}

	for i := 0; i < len(testData); i++ {
		number := testData[i].number
		digit := testData[i].digit
		expected := testData[i].expected
		actual := containsNumberInNote(number, digit)
		if actual != expected {
			t.Errorf("containsNumberInNote for %d in %d should be set but was not", digit, number)
		}
	}
}
