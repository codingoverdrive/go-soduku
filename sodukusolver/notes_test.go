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

func Test_setNumber(t *testing.T) {
	assert.Equal(t, 0, setNumber(0), "Failed to set correct bit")
	assert.Equal(t, 0x01, setNumber(1), "Failed to set correct bit")
	assert.Equal(t, 0x04, setNumber(3), "Failed to set correct bit")
	assert.Equal(t, 0x100, setNumber(9), "Failed to set correct bit")
}

func Test_isNumberSet(t *testing.T) {
	assert.Equal(t, false, isNumberSet(0, 1), "Failed to test if digit set")
	assert.Equal(t, true, isNumberSet(1, 1), "Failed to test if digit set")
	assert.Equal(t, false, isNumberSet(0x06, 1), "Failed to test if digit set")
	assert.Equal(t, true, isNumberSet(0x06, 2), "Failed to test if digit set")
	assert.Equal(t, true, isNumberSet(0x06, 3), "Failed to test if digit set")
	assert.Equal(t, false, isNumberSet(0x06, 4), "Failed to test if digit set")
	assert.Equal(t, false, isNumberSet(0, 9), "Failed to test if digit set")
	assert.Equal(t, true, isNumberSet(0x100, 9), "Failed to test if digit set")
}

func Test_getCommonNumberCount(t *testing.T) {
	assert.Equal(t, 0, getCommonNumberCount(0x0, 0x0), "Failed to getCommmonNumberCount correctly")
	assert.Equal(t, 0, getCommonNumberCount(0x04, 0x02), "Failed to getCommmonNumberCount correctly")
	assert.Equal(t, 1, getCommonNumberCount(0x04, 0x05), "Failed to getCommmonNumberCount correctly")
	assert.Equal(t, 2, getCommonNumberCount(0x109, 0x105), "Failed to getCommmonNumberCount correctly")
}

func Test_isSameColumnInBlock(t *testing.T) {
	assert.False(t, isSameColumnInBlock([]int{1}), "Should be false")

	cellIndexes1 := []int{1, 5}
	assert.False(t, isSameColumnInBlock(cellIndexes1), "Should not be in same column")

	cellIndexes2 := []int{1, 4}
	assert.True(t, isSameColumnInBlock(cellIndexes2), "Should be in same column")

	cellIndexes3 := []int{1, 4, 5}
	assert.False(t, isSameColumnInBlock(cellIndexes3), "Should not be in same column")
}

func Test_isSameRowInBlock(t *testing.T) {
	assert.False(t, isSameRowInBlock([]int{1}), "Should be false")

	cellIndexes1 := []int{1, 4}
	assert.False(t, isSameRowInBlock(cellIndexes1), "Should not be in same row")

	cellIndexes2 := []int{1, 2}
	assert.True(t, isSameRowInBlock(cellIndexes2), "Should be in same row")

	cellIndexes3 := []int{1, 2, 4}
	assert.False(t, isSameRowInBlock(cellIndexes3), "Should not be in same row")
}
