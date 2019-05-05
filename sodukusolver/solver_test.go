package sodukusolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveBoard1(t *testing.T) {
	board := [9][9]int{
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

	expected := [9][9]int{
		{9, 2, 6, 3, 7, 1, 5, 8, 4},
		{7, 1, 5, 8, 2, 4, 3, 9, 6},
		{8, 4, 3, 9, 6, 5, 1, 7, 2},
		{3, 7, 1, 5, 8, 2, 6, 4, 9},
		{6, 5, 4, 7, 9, 3, 2, 1, 8},
		{2, 9, 8, 4, 1, 6, 7, 3, 5},
		{5, 6, 9, 1, 3, 8, 4, 2, 7},
		{1, 8, 2, 6, 4, 7, 9, 5, 3},
		{4, 3, 7, 2, 5, 9, 8, 6, 1}}

	solution := SolveBoard(board)
	assert.True(t, solution.Solved, "Board should be solved")
	assert.Equal(t, expected, solution.board, "Unexpected board solution")
}

func Test_SolveBoard2(t *testing.T) {
	board := [9][9]int{
		{0, 5, 0, 0, 0, 0, 0, 0, 3},
		{1, 8, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 5, 1, 0, 8, 0, 0},
		{0, 0, 5, 0, 0, 0, 1, 0, 2},
		{0, 3, 1, 0, 0, 0, 5, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 4, 8},
		{0, 0, 7, 8, 6, 0, 0, 3, 9},
		{0, 0, 0, 0, 0, 4, 2, 0, 0},
		{4, 0, 0, 7, 0, 9, 6, 0, 0},
	}

	expected := [9][9]int{
		{2, 5, 6, 9, 4, 8, 7, 1, 3},
		{1, 8, 4, 2, 7, 3, 9, 5, 6},
		{3, 7, 9, 5, 1, 6, 8, 2, 4},
		{8, 4, 5, 6, 3, 7, 1, 9, 2},
		{9, 3, 1, 4, 8, 2, 5, 6, 7},
		{7, 6, 2, 1, 9, 5, 3, 4, 8},
		{5, 2, 7, 8, 6, 1, 4, 3, 9},
		{6, 9, 8, 3, 5, 4, 2, 7, 1},
		{4, 1, 3, 7, 2, 9, 6, 8, 5},
	}

	solution := SolveBoard(board)
	assert.True(t, solution.Solved, "Board should be solved")
	assert.Equal(t, expected, solution.board, "Unexpected board solution")
}

func Test_SolveBoard3(t *testing.T) {
	board := [9][9]int{
		{0, 0, 8, 9, 3, 0, 0, 0, 0},
		{0, 9, 6, 0, 0, 8, 0, 0, 0},
		{7, 0, 2, 0, 0, 4, 0, 0, 0},
		{0, 0, 4, 0, 2, 5, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 5, 0},
		{0, 0, 0, 3, 4, 0, 9, 0, 0},
		{0, 0, 0, 1, 0, 0, 2, 0, 4},
		{0, 0, 0, 4, 0, 0, 3, 7, 0},
		{0, 0, 0, 0, 7, 9, 5, 0, 0},
	}

	expected := [9][9]int{
		{1, 4, 8, 9, 3, 7, 6, 2, 5},
		{5, 9, 6, 2, 1, 8, 7, 4, 3},
		{7, 3, 2, 5, 6, 4, 8, 9, 1},
		{9, 8, 4, 7, 2, 5, 1, 3, 6},
		{3, 1, 7, 8, 9, 6, 4, 5, 2},
		{2, 6, 5, 3, 4, 1, 9, 8, 7},
		{8, 7, 9, 1, 5, 3, 2, 6, 4},
		{6, 5, 1, 4, 8, 2, 3, 7, 9},
		{4, 2, 3, 6, 7, 9, 5, 1, 8},
	}

	solution := SolveBoard(board)
	assert.True(t, solution.Solved, "Board should be solved")
	assert.Equal(t, expected, solution.board, "Unexpected board solution")

}

/*
func Test_SolveBoard4(t *testing.T) {
	board := [9][9]int{
		{0, 0, 0, 0, 0, 3, 0, 0, 0},
		{6, 7, 0, 0, 0, 4, 3, 0, 5},
		{3, 0, 0, 0, 7, 0, 0, 8, 0},
		{7, 0, 2, 6, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 3},
		{1, 0, 9, 7, 0, 0, 0, 0, 6},
		{4, 0, 0, 0, 2, 0, 0, 6, 0},
		{9, 2, 0, 0, 0, 6, 7, 0, 8},
		{0, 0, 0, 0, 0, 8, 0, 0, 0},

		// {7, 0, 0, 0, 0, 5, 2, 3, 0},
		// {0, 0, 0, 0, 6, 0, 4, 0, 9},
		// {0, 0, 0, 1, 3, 0, 0, 6, 8},
		// {0, 8, 0, 5, 0, 0, 0, 0, 3},
		// {0, 0, 0, 0, 0, 0, 1, 2, 0},
		// {4, 0, 0, 0, 0, 1, 8, 0, 0},
		// {0, 2, 0, 0, 0, 0, 0, 0, 0},
		// {5, 0, 1, 0, 0, 4, 0, 0, 0},
		// {0, 6, 0, 9, 0, 0, 0, 0, 4},
	}

	expected := [9][9]int{}

	solution := SolveBoard(board)
	PrintSolutionSteps(solution, false)

	assert.True(t, solution.Solved, "Board should be solved")
	assert.Equal(t, expected, solution.board, "Unexpected board solution")

}
*/
