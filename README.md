# go-soduku
Soduku Solver in Go

A package comprising different solvers for Soduku puzzles.

# Solvers

The current implemented solvers are:
* Naked Singles
* Hidden Singles
* Naked Pairs

# Running the Solver

Checkout the project from github

```
git clone https://github.com/codingoverdrive/go-soduku.git
cd cmd/solver
go run main.go
```

To solve a different puzzle, modify the board two dimensional array in cmd/solver/main.go

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


Pass 1 Found Naked Single [7] at Cell H6
Pass 2 Found Hidden Single [8] at Cell D5
Pass 3 Found Hidden Single [6] at Cell H4
Pass 4 Found Hidden Single [7] at Cell I3
Pass 5 Found Hidden Single [6] at Cell B9
Pass 6 Found Hidden Single [1] at Cell D3
Pass 7 Found Hidden Single [2] at Cell I4
Pass 8 Found Hidden Single [1] at Cell F5
Pass 9 Found Naked Single [9] at Cell F2
Pass 10 Found Naked Single [2] at Cell H3
Pass 11 Found Naked Single [4] at Cell E3
Pass 12 Found Naked Single [2] at Cell F1
Pass 13 Found Naked Single [7] at Cell F7
Pass 14 Found Naked Single [2] at Cell E7
Pass 15 Found Hidden Single [9] at Cell B8
Pass 16 Found Hidden Single [2] at Cell C9
Pass 17 Found Hidden Single [2] at Cell G8
Pass 18 Found Naked Single [4] at Cell D8
Pass 19 Found Naked Single [9] at Cell D9
Pass 20 Found Naked Pairs [3,6] in C3, C5, Removing [3,6] from C2, C7
Pass 21 Found Naked Pairs [3,9] in E6, I6, Removing [3,9] from A6, B6
Pass 22 Found Naked Pairs [1,4] in A6, A9, Removing [1,4] from A1, A7
Pass 23 Found Naked Pairs [1,4] in A9, C7, Removing [1,4] from A7, B7
Pass 24 Found Naked Single [9] at Cell A1
Pass 25 Found Hidden Single [9] at Cell G3
Pass 26 Found Hidden Single [3] at Cell G5
Pass 27 Found Hidden Single [3] at Cell I2
Pass 28 Found Naked Single [6] at Cell C5
Pass 29 Found Naked Single [9] at Cell I6
Pass 30 Found Naked Single [7] at Cell A5
Pass 31 Found Naked Single [3] at Cell C3
Pass 32 Found Naked Single [3] at Cell E6
Pass 33 Found Naked Single [5] at Cell I5
Pass 34 Found Naked Single [6] at Cell A3
Pass 35 Found Naked Single [7] at Cell E4
Pass 36 Found Naked Single [9] at Cell E5
Pass 37 Found Hidden Single [8] at Cell I7
Pass 38 Found Hidden Single [8] at Cell H2
Pass 39 Found Naked Single [3] at Cell B7
Pass 40 Found Naked Single [5] at Cell H8
Pass 41 Found Naked Single [5] at Cell A7
Pass 42 Found Naked Single [8] at Cell A8
Pass 43 Found Naked Single [8] at Cell B4
Pass 44 Found Naked Single [4] at Cell G7
Pass 45 Found Naked Single [1] at Cell H1
Pass 46 Found Naked Single [3] at Cell A4
Pass 47 Found Naked Single [1] at Cell C7
Pass 48 Found Naked Single [5] at Cell G1
Pass 49 Found Naked Single [4] at Cell I1
Pass 50 Found Naked Single [1] at Cell I9
Pass 51 Found Naked Single [4] at Cell A9
Pass 52 Found Naked Single [4] at Cell C2
Pass 53 Found Naked Single [1] at Cell A6
Pass 54 Found Naked Single [1] at Cell B2
Pass 55 Found Naked Single [4] at Cell B6


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

Solved in 111.27µs
```