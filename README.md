# go-soduku
Soduku Solver in Go

A package comprising different solver strategies for Soduku puzzles.

# Implemented Solver Strategies

The following strategies are implemented:
* Naked Singles
* Hidden Singles
* Hidden Pairs
* Naked Pairs
* Pointing Pairs

# Using the Solver package in your own application

Import the sodukusolver package
```
import "github.com/codingoverdrive/go-soduku/sodukusolver"
```

Create a board and pass it to the SolveBoard() function. Note that zero denotes an unknown/unsolved/empty cell

```
  //create a board
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
```

The solver works by applying different solution strategies repeatedly until the puzzle is solved. 
It applies the simplest strategies first, before trying more complicated strategies if those fail. 
As soon as a strategy yields a solution for a cell, or eliminates notes from unsolved cells, the solver
returns to the simplest strategy again.

```
	//keep applying the different solution strategies until the puzzle is solved
	for {
		result := applyCellSolutionStrategy(findNakedSingles, &solution)
		if result.success {
			continue
		}

		result = applyCellSolutionStrategy(findHiddenSingles, &solution)
		if result.success {
			continue
		}

		result = applyCellExclusionStrategy(findHiddenPairExclusions, &solution)
		if result.success {
			continue
		}
    ...
```

The solver will stop either when the board is solved, or the strategies yield no further solutions. 
The solution includes the final board and the steps (and strategies) taken to yield the final board state.

The solver solves most (hard) puzzles using four basic strategies in less than 0.12ms (120us) on a MacBook Pro. 

This project is a work in progress and more strategies will be added over time.


# Running the solver locally

Checkout the project from github

```
git clone https://github.com/codingoverdrive/go-soduku.git
cd cmd/solver
go run main.go
```

To solve a different puzzle, modify the board two dimensional array in `cmd/solver/main.go` or add your own tests in `sodukusolver/solver_test.go`

The 9x9 array represents a soduku board

```
Initial Soduku board
       1       2       3       4       5       6       7       8       9
   =========================================================================
   I       |       |       I       |       |       I       |       |       I
 A I       |   2   |       I       |       |       I       |       |       I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 B I   7   |       |   5   I       |   2   |       I       |       |       I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 C I   8   |       |       I   9   |       |   5   I       |   7   |       I
   I       |       |       I       |       |       I       |       |       I
   =========================================================================
   I       |       |       I       |       |       I       |       |       I
 D I   3   |   7   |       I   5   |       |   2   I   6   |       |       I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 E I   6   |   5   |       I       |       |       I       |   1   |   8   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 F I       |       |   8   I   4   |       |   6   I       |   3   |   5   I
   I       |       |       I       |       |       I       |       |       I
   =========================================================================
   I       |       |       I       |       |       I       |       |       I
 G I       |   6   |       I   1   |       |   8   I       |       |   7   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 H I       |       |       I       |   4   |       I   9   |       |   3   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 I I       |       |       I       |       |       I       |   6   |       I
   I       |       |       I       |       |       I       |       |       I
   =========================================================================

With Notes
       1       2       3       4       5       6       7       8       9
   =========================================================================
   I  1..  |       |  1.3  I  ..3  |  1.3  |  1.3  I  1.3  |  ...  |  1..  I
 A I  4..  |   2   |  4.6  I  ..6  |  ..6  |  4..  I  45.  |  45.  |  4.6  I
   I  ..9  |       |  ..9  I  78.  |  78.  |  7..  I  .8.  |  .89  |  ..9  I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |  1.3  |       I  ..3  |       |  1.3  I  1.3  |  ...  |  1..  I
 B I   7   |  4..  |   5   I  ..6  |   2   |  4..  I  4..  |  4..  |  4.6  I
   I       |  ..9  |       I  .8.  |       |  ...  I  .8.  |  .89  |  ..9  I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |  1.3  |  1.3  I       |  1.3  |       I  123  |       |  12.  I
 C I   8   |  4..  |  4.6  I   9   |  ..6  |   5   I  4..  |   7   |  4.6  I
   I       |  ...  |  ...  I       |  ...  |       I  ...  |       |  ...  I
   =========================================================================
   I       |       |  1..  I       |  1..  |       I       |  ...  |  ...  I
 D I   3   |   7   |  4..  I   5   |  ...  |   2   I   6   |  4..  |  4..  I
   I       |       |  ..9  I       |  .89  |       I       |  ..9  |  ..9  I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |  .2.  I  ..3  |  ..3  |  ..3  I  .2.  |       |       I
 E I   6   |   5   |  4..  I  ...  |  ...  |  ...  I  4..  |   1   |   8   I
   I       |       |  ..9  I  7..  |  7.9  |  7.9  I  7..  |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I  12.  |  1..  |       I       |  1..  |       I  .2.  |       |       I
 F I  ...  |  ...  |   8   I   4   |  ...  |   6   I  ...  |   3   |   5   I
   I  ..9  |  ..9  |       I       |  7.9  |       I  7..  |       |       I
   =========================================================================
   I  .2.  |       |  .23  I       |  ..3  |       I  .2.  |  .2.  |       I
 G I  45.  |   6   |  4..  I   1   |  .5.  |   8   I  45.  |  45.  |   7   I
   I  ..9  |       |  ..9  I       |  ..9  |       I  ...  |  ...  |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I  12.  |  1..  |  12.  I  .2.  |       |  ...  I       |  .2.  |       I
 H I  .5.  |  ...  |  ...  I  ..6  |   4   |  ...  I   9   |  .5.  |   3   I
   I  ...  |  .8.  |  7..  I  7..  |       |  7..  I       |  .8.  |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I  12.  |  1.3  |  123  I  .23  |  ..3  |  ..3  I  12.  |       |  12.  I
 I I  45.  |  4..  |  4..  I  ...  |  .5.  |  ...  I  45.  |   6   |  4..  I
   I  ..9  |  .89  |  7.9  I  7..  |  7.9  |  7.9  I  .8.  |       |  ...  I
   =========================================================================

Step  1 Found Naked Single [7] at H6
Step  2 Found Hidden Single [8] at D5
Step  3 Found Hidden Single [6] at H4
Step  4 Found Hidden Single [7] at I3
Step  5 Found Hidden Single [6] at B9
Step  6 Found Hidden Single [1] at D3
Step  7 Found Hidden Single [2] at I4
Step  8 Found Hidden Single [1] at F5
Step  9 Found Naked Single [9] at F2
Step 10 Found Naked Single [2] at H3
Step 11 Found Naked Single [4] at E3
Step 12 Found Naked Single [2] at F1
Step 13 Found Naked Single [7] at F7
Step 14 Found Naked Single [2] at E7
Step 15 Found Hidden Single [9] at B8
Step 16 Found Hidden Single [2] at C9
Step 17 Found Hidden Single [2] at G8
Step 18 Found Naked Single [4] at D8
Step 19 Found Naked Single [9] at D9
Step 20 Found Naked Pairs [3,6] in C3, C5, Removing [3,6] from C2, C7
Step 21 Found Naked Pairs [3,9] in E6, I6, Removing [3,9] from A6, B6
Step 22 Found Naked Pairs [1,4] in A6, A9, Removing [1,4] from A1, A7
Step 23 Found Naked Pairs [1,4] in A9, C7, Removing [1,4] from A7, B7
Step 24 Found Naked Single [9] at A1
Step 25 Found Hidden Single [9] at G3
Step 26 Found Hidden Single [3] at G5
Step 27 Found Hidden Single [3] at I2
Step 28 Found Naked Single [6] at C5
Step 29 Found Naked Single [9] at I6
Step 30 Found Naked Single [7] at A5
Step 31 Found Naked Single [3] at C3
Step 32 Found Naked Single [3] at E6
Step 33 Found Naked Single [5] at I5
Step 34 Found Naked Single [6] at A3
Step 35 Found Naked Single [7] at E4
Step 36 Found Naked Single [9] at E5
Step 37 Found Hidden Single [8] at I7
Step 38 Found Hidden Single [8] at H2
Step 39 Found Naked Single [3] at B7
Step 40 Found Naked Single [5] at H8
Step 41 Found Naked Single [5] at A7
Step 42 Found Naked Single [8] at A8
Step 43 Found Naked Single [8] at B4
Step 44 Found Naked Single [4] at G7
Step 45 Found Naked Single [1] at H1
Step 46 Found Naked Single [3] at A4
Step 47 Found Naked Single [1] at C7
Step 48 Found Naked Single [5] at G1
Step 49 Found Naked Single [4] at I1
Step 50 Found Naked Single [1] at I9
Step 51 Found Naked Single [4] at A9
Step 52 Found Naked Single [4] at C2
Step 53 Found Naked Single [1] at A6
Step 54 Found Naked Single [1] at B2
Step 55 Found Naked Single [4] at B6

       1       2       3       4       5       6       7       8       9
   =========================================================================
   I       |       |       I       |       |       I       |       |       I
 A I   9   |   2   |   6   I   3   |   7   |   1   I   5   |   8   |   4   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 B I   7   |   1   |   5   I   8   |   2   |   4   I   3   |   9   |   6   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 C I   8   |   4   |   3   I   9   |   6   |   5   I   1   |   7   |   2   I
   I       |       |       I       |       |       I       |       |       I
   =========================================================================
   I       |       |       I       |       |       I       |       |       I
 D I   3   |   7   |   1   I   5   |   8   |   2   I   6   |   4   |   9   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 E I   6   |   5   |   4   I   7   |   9   |   3   I   2   |   1   |   8   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 F I   2   |   9   |   8   I   4   |   1   |   6   I   7   |   3   |   5   I
   I       |       |       I       |       |       I       |       |       I
   =========================================================================
   I       |       |       I       |       |       I       |       |       I
 G I   5   |   6   |   9   I   1   |   3   |   8   I   4   |   2   |   7   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 H I   1   |   8   |   2   I   6   |   4   |   7   I   9   |   5   |   3   I
   I       |       |       I       |       |       I       |       |       I
   --------+-------+-------+-------+-------+-------+-------+-------+--------
   I       |       |       I       |       |       I       |       |       I
 I I   4   |   3   |   7   I   2   |   5   |   9   I   8   |   6   |   1   I
   I       |       |       I       |       |       I       |       |       I
   =========================================================================

Solved in 114.082µs
```