package main

import "github.com/codingoverdrive/go-soduku/sodukusolver"

func init() {
}

func main() {
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

	sodukusolver.InitaliseBoard(board)
	sodukusolver.SolveBoard()
}
