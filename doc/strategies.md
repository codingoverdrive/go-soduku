# Soduku Strategies

Soduku strategies can be divided into two categories:
* Those that produce solutions to a cell
* Those that eliminate potential solutions in a cell

They can then be further divided into those that are simple to understand, and those that more advanced or difficult to understand and implement in code.

So far, five different strategies have been implemented; two of which will produce a cell solution, and the others that will eliminate possible cell solutions.

| Complexity  | Name  | Solution Type  |
|---|---|---|
| Simple | Naked Singles | Single cell solution |
| Simple | Hidden Single | Single cell solution |
| Simple | Naked Pairs | Elimination of candidate solutions |
| Simple | Hidden Pairs | Elimination of candidate solutions |
| Simple | Box Line Reduction | Elimination of candidate solutions |

Many of the strategies look for patterns of numbers in "houses" on the board. A house is a row, column or a 3x3 block (of which there are 9 in a Soduku puzzle). This means that certain coding optimisations can be made. 

Searching for pairs of numbers in notes in a row, for example, is the same as search for pairs in a column. So some of the strategies will make use of algorithms that look for patterns in nine cells; first applying the algorithm to rows, then columns and finally also blocks (if appropriate).

See the example below.

```
func findHiddenSingles(notes [9][9]int) []AbsoluteCellSolution {
	var solutions = []AbsoluteCellSolution{}

	//search the rows
	for row := 0; row < 9; row++ {
>>LOOK>>	rowSolutions := findHiddenSinglesInNineCells(convertRowToNineCells(notes, row))
		for i := 0; i < len(rowSolutions); i++ {
			s := rowSolutions[i]
    		    solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{row, s.index, s.number, "Hidden Single", "Cell"})
		}
	}

	//search the columns
	for column := 0; column < 9; column++ {
>>LOOK>>	columnSolutions := findHiddenSinglesInNineCells(convertColumnToNineCells(notes, column))
		for i := 0; i < len(columnSolutions); i++ {
			s := columnSolutions[i]
    		    solutions = appendAbsoluteCellSolution(solutions, AbsoluteCellSolution{s.index, column, s.number, "Hidden Single", "Cell"})
		}
	}
...

>>LOOK>>    Denotes where a 9 cell searcher function
            findHiddenSinglesInNineCells is being used
            to search in either rows or columns
```

No matter how hard the Soduku puzzle, it is the basic strategies that produce most of the solution steps. Only very occasionally do you have to use an advanced strategy to solve the puzzle, and normally only for a single step. After that the solver can revert to the simple or basic strategies again.

Here is a typical sets of steps to solve a hard puzzle (graded 4 out of 5 where five is deemed expert level):

```
Step  1 Found Naked Single [7] at H6
Step  2 Found Hidden Single [8] at D5
Step  3 Found Hidden Single [6] at H4
Step  4 Found Hidden Single [7] at I3
Step  5 Found Hidden Single [6] at B9
...
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
...
```
Notice how detection of (naked and hidden) singles is mostly used to solve this puzzle, with pair detection for steps 20-23, after which the steps revert to (naked and hidden) singles again.

The basic strategies implemented in this solver should yield solutions for more than 95% of the hard puzzles that you give it to solve, and in generally 50-60 steps.

It is my intention to add more advanced solving strategies over time in order to allow the solver to deal with some of the harder and diabolical Soduku puzzles.
