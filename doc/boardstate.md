# Representing The Soduku Board State

The board is represented by a (two dimensional) 9x9 int array. Each element in the array has a value 0 - 9 where zero represents an unsolved cell, and a single value of 1 .. 9 represents the solution to a particular cell.

This is an example of a board:

```
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
```

which equates to this (display representation)

```
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
```

In addition to the board, a second 9x9 int array called `notes` is used to represent the possible values that could be in any unsolved cell. These notes are analagous to the small numbers you pencil on a Soduku board in a newspaper or book of Sokuku puzzles.

Those cells that are unsolved and have notes (in the display below) are presented by cells that have dots in them.

```
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

```

The notes for cell A1 contains the values 1,4 and 9 as possible solution values. This means that the solution will be one of those three numbers although we do know which one yet.

When a cell is already solved, the `board` cell value will be >0 and the `notes` cell value will be zero. When a cell is unsolved, then the `board` cell value will be zero and the `notes` cell value will be >0.

The non zero numbers in the `notes` array indicate the possible values for a cell where each bit in the number corresponds to those numbers. If a notes cell, for example, contains the value 0x05 (hex) or 0101 (in binary) then this denotes that the possible values or "notes" for that cell are 1 and 3. 

```
 Possible note value  9 8 7 6 5 4 3 2 1
                      -----------------
 Int value (binary)   0 0 0 0 0 0 1 0 1    ==  0x05  or 5 (in decimal)
```

Why represent the notes this way? Because it is an extremely easy way to test for numbers in notes, and to add or remove numbers when applying solution strategies. Set arithmetic can be performed very efficiently by computers, and is probably a better solution than using a 9 element array or string to represent the notes values in a cell.

So, if you wanted to add the number 3 to an existing cell's notes, you would set the third bit of the number (reading from right to left).

In code this is represented as:
```
newNote = oldNote | 0x04
```
Yes, the 0x04 looks strange, but 0x04 means 0000 0100 (in binary), and denotes that the 3rd bit (right to left) is set.

Similarly to test whether the notes cell contains the number 3 (as a possible solution value), you test whether the 3rd bit is set.

```
if note & 0x04 == 0x04 {
    //the 3rd bit denoting the number 3 is set
} else {
    //the 3rd bit is not set
}
```

The `sodukusolver/notes.go` file has some helper functions to set, test and count numbers in a notes cell which can help reduce the effort of working with bit arithmetic. 

So now we have two 9x9 arrays to represent the puzzle state; one for the board of solved values (called `board`), and one for the possible notes values (called `notes`) so far.

Due to the way the solver works and based on soduku strategies that often identify values that cannot be a solution for a cell, we actually require a third 9x9 int array to represent those impossible solution values. Again these are represented as integers whose bits represent those digits that cannot be solutions.

The notes for a cell are actually calculated by removing the `exclusion` bits from the `notes` bits. 

Assume that cell A1 has no solution, and the solver has calculated that the permissable notes values are 1, 3, and 4. This would mean that:

```
board[0][0] == 0  (no solution for this cell)
notes[0][0] == 0x0d  (or 0000 1101 in binary)
```

If after applying a strategy, we now find that the number 3 (third bit from right to left) cannot be a possible value in that cell then the following would be true:

```
board[0][0] == 0  (no solution for this cell)
notes[0][0] == 0x0d  (or 0000 1101 in binary)
exclusions[0][0] == 0x04  (or 0000 0100 in binary)
```

The actual possible notes values are therefore a combination of the `notes` and `exclusions` cell values. We want to ensure that a bit is removed from the `notes` cell if the same bit is set in the `exclusions` cell. In pure logic terms, this can be calculated as follows:

```
// starting vales
//notes[0][0] == 0x0d  (or 0000 1101 in binary)
//exclusions[0][0] == 0x04  (or 0000 0100 in binary)

// invert the bits but ignore everything to the left of the 9th bit
// reading from right to left using the 0x1ff mask
inverted[0][0] = ^notes[0][0] & 0x01ff
// 0 0000 1101 -> 1 1111 0010

// now remove any bits that have been set in the exclusions
// do this by OR'ing the values
newInvertedNotes[0][0] = inverted[0][0] | exclusions[0][0]
// 1 1111 0010 | 0 0000 0100 -> 1 1111 0110

// now invert the bits again to produce the actual new notes
newNotes = ~newInvertedNotes[0][0] & 0x01ff
// 1 1111 0110 -> 0 0000 1001

// notice how the third bit (from right to left) has been removed
```

OK, so the logic above seems *quite* complicated (and it is a little unwieldy). However the code that the solver uses to perform this is actually much simpler using OR's and one XOR (as shown below)

```
// determine all the known "solved" values for the row, column and block
// that the cell "sits" within and treat the "exclusions" as additional
// solved values in order to eliminate those too

    rowSolvedNumbers := rowSolved[row]
    colSolvedNumbers := columnSolved[column]
    blockSolvedNumbers := blockSolved[blockIndex]
    exclusionNumbers := exclusions[row][column]
    newNotes[row][column] = 0x1ff ^ (rowSolvedNumbers | colSolvedNumbers | blockSolvedNumbers | exclusionNumbers)
```

Here are the bit truth tables in case you're interested:
```
  NOT       AND            OR             XOR
 0 -> 1    0 & 0 -> 0     0 | 0 -> 0     0 ^ 0 -> 1   
 1 -> 0    0 & 1 -> 0     0 | 1 -> 1     0 ^ 1 -> 0
           1 & 0 -> 0     1 | 0 -> 1     1 ^ 0 -> 0
           1 & 1 -> 1     1 | 1 -> 1     1 ^ 1 -> 1
```

In summary, the board, notes and exclusions are represented as three 9x9 int arrays. 

The values in the `board` array are the actual decimal number of the solution while the values in the `notes` and `exclusions` arrays are the bitwise representation of the notes and exclusions respectively.