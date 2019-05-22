package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/codingoverdrive/go-soduku/sodukusolver"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	//get command line argument to determine what actons to take
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "perftest" {
		//this is a performance test
		solveMultipleBoards()
	} else {
		//solve a single board and show solution
		solveSingleBoard()
	}
}

// solveSingleBoard solves a single board
func solveSingleBoard() {
	//get the last board
	board := boards[len(boards)-1]

	//compute the solution
	solution := sodukusolver.SolveBoard(board)

	//show each solution step
	sodukusolver.PrintSolutionSteps(solution, false)

	//show the final state of the board
	sodukusolver.PrintBoard(solution.FinalBoard, solution.FinalNotes, true)

	if solution.Solved {
		print("\nSolved in ", fmt.Sprintf("%s", solution.Elapsed), "\n")
	} else {
		print("\nUnsolved after ", fmt.Sprintf("%s", solution.Elapsed), "\n")
		print("Stopping\n\n")
	}
}

//solveMultipleBoards solves multiple boards as part of a performance test
func solveMultipleBoards() {
	var maxIterations = 2000
	var solvedCount = 0
	startTime := time.Now()
	for i := 0; i < maxIterations; i++ {
		board := boards[randInt(0, 9)]
		solution := sodukusolver.SolveBoard(board)
		if solution.Solved {
			solvedCount++
		}
	}
	var elapsedTime = time.Since(startTime)

	//print stats
	print(fmt.Sprintf("%d", maxIterations), " boards processed in ", fmt.Sprintf("%s", elapsedTime), "\n")
	boardsPerSec := (int64(maxIterations) * int64(time.Second)) / int64(elapsedTime)

	//show solution status (solved vs unsolved boards)
	var solutionSuccess = 100 * solvedCount / maxIterations
	print(fmt.Sprintf("%d", solutionSuccess), "% solution rate\n")

	//show the solution rate per second
	print("Solutions/sec: ", boardsPerSec, "\n")

}

//randInt generates a randon integer between the min and max bounds
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// sample (hard) boards
var boardA = [9][9]int{
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

var boardB = [9][9]int{
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

var boardC = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 5, 8, 0, 1, 6, 0, 0},
	{0, 0, 0, 5, 9, 6, 0, 0, 0},
	{0, 0, 0, 9, 0, 3, 0, 0, 0},
	{0, 5, 0, 4, 0, 2, 0, 1, 0},
	{0, 0, 9, 0, 1, 0, 8, 0, 0},
	{0, 6, 4, 0, 0, 0, 1, 8, 0},
	{0, 2, 8, 1, 0, 4, 3, 6, 0},
	{0, 7, 1, 0, 6, 0, 4, 5, 0},
}

var boardD = [9][9]int{
	{7, 0, 0, 0, 0, 5, 2, 3, 0},
	{0, 0, 0, 0, 6, 0, 4, 0, 9},
	{0, 0, 0, 1, 3, 0, 0, 6, 8},
	{0, 8, 0, 5, 0, 0, 0, 0, 3},
	{0, 0, 0, 0, 0, 0, 1, 2, 0},
	{4, 0, 0, 0, 0, 1, 8, 0, 0},
	{0, 2, 0, 0, 0, 0, 0, 0, 0},
	{5, 0, 1, 0, 0, 4, 0, 0, 0},
	{0, 6, 0, 9, 0, 0, 0, 0, 4},
}

var boardE = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 4, 7, 0, 9, 3, 0, 0},
	{0, 3, 0, 1, 4, 2, 0, 7, 0},
	{5, 4, 7, 0, 0, 0, 2, 9, 3},
	{3, 1, 0, 0, 0, 0, 0, 6, 8},
	{2, 6, 8, 0, 0, 0, 7, 1, 5},
	{0, 9, 0, 2, 6, 8, 0, 5, 0},
	{0, 0, 1, 4, 0, 3, 9, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var boardF = [9][9]int{
	{0, 0, 0, 0, 0, 6, 0, 0, 0},
	{0, 2, 0, 8, 7, 9, 0, 0, 0},
	{0, 7, 0, 0, 0, 0, 2, 0, 0},
	{9, 0, 0, 0, 0, 0, 0, 2, 5},
	{1, 0, 0, 9, 0, 0, 0, 3, 0},
	{0, 6, 0, 0, 1, 0, 0, 7, 0},
	{6, 4, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 2, 7, 0, 0, 3, 1, 0},
	{0, 0, 1, 0, 8, 4, 0, 0, 0},
}

var boardG = [9][9]int{
	{6, 0, 9, 0, 0, 0, 0, 0, 0},
	{0, 4, 0, 6, 9, 0, 0, 5, 0},
	{8, 0, 0, 0, 0, 0, 1, 0, 0},
	{9, 0, 0, 0, 8, 0, 4, 0, 0},
	{0, 5, 0, 0, 0, 0, 0, 9, 0},
	{0, 0, 6, 0, 7, 0, 0, 0, 5},
	{0, 0, 1, 0, 0, 0, 0, 0, 2},
	{0, 2, 0, 0, 5, 4, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 5, 0, 8},
}

var boardH = [9][9]int{
	{0, 7, 0, 0, 0, 2, 3, 5, 0},
	{0, 0, 3, 7, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 9, 0, 0, 2},
	{0, 0, 1, 0, 0, 0, 7, 0, 4},
	{0, 4, 8, 1, 0, 0, 0, 0, 0},
	{3, 6, 7, 0, 2, 0, 0, 9, 0},
	{0, 3, 0, 8, 6, 5, 0, 4, 0},
	{0, 0, 2, 9, 7, 0, 0, 0, 3},
	{0, 0, 0, 2, 0, 0, 0, 0, 0},
}

var boardI = [9][9]int{
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

var boards = [][9][9]int{
	boardA, boardB, boardC, boardD, boardE, boardF, boardG, boardH, boardI,
}
