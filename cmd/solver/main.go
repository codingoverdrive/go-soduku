package main

import (
	"fmt"

	"github.com/codingoverdrive/go-soduku/sodukusolver"
)

func init() {
}

func main() {
	//the Soduku puzzle board is represented as a 9x9 array of int
	var board = [9][9]int{
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

	//compute the solution
	solution := sodukusolver.SolveBoard(board)

	//show each solution step
	for i := 0; i < len(solution.Steps); i++ {
		if i == 0 {
			//show the initial state of the board
			print("Initial Soduku board\n")
			sodukusolver.PrintBoard(solution.InitialBoard, [9][9]int{}, false)
			print("\n\n")
			//show the board again, but with the notes populated as well
			print("With Notes\n")
			sodukusolver.PrintBoard(solution.InitialBoard, solution.InitialNotes, true)
			print("\n\n")
		}
		//provide the solution details for each step
		step := solution.Steps[i]
		print("Pass ", (i + 1), " ", step.Description, "\n")
	}

	//show the final state of the board
	print("\n\n")
	finalStep := solution.Steps[len(solution.Steps)-1]
	sodukusolver.PrintBoard(finalStep.Board, finalStep.Notes, false)

	if solution.Solved {
		print("\nSolved in ", fmt.Sprintf("%s", solution.Elapsed), "\n")
	} else {
		print("\nUnsolved after ", fmt.Sprintf("%s", solution.Elapsed), "\n")
		print("Stopping\n\n")
	}

}
