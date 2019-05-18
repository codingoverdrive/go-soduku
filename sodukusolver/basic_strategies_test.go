package sodukusolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNakedSingles(t *testing.T) {
	notes1 := [9][9]int{
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
	assert.Equal(t, []AbsoluteCellSolution{}, findNakedSingles(notes1), "No solution expected")

	notes2 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected := []AbsoluteCellSolution{
		AbsoluteCellSolution{1, 2, 1, "Naked Single", "Cell"},
		AbsoluteCellSolution{7, 6, 3, "Naked Single", "Cell"},
	}
	assert.Equal(t, expected, findNakedSingles(notes2), "Two solutions expected")
}

func Test_findHiddenSinglesInNineCells(t *testing.T) {
	cells1 := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, []RelativeCellSolution{}, findHiddenSinglesInNineCells(cells1), "No solutions expected")

	cells2 := [9]int{0, 0, 3, 0, 0, 0, 0, 4, 6}
	expected2 := []RelativeCellSolution{
		{2, 1},
	}
	assert.Equal(t, expected2, findHiddenSinglesInNineCells(cells2), "One solution expected")

	cells3 := [9]int{12, 0, 2, 0, 0, 0, 1, 4, 6}
	expected3 := []RelativeCellSolution{
		{6, 1},
		{0, 4},
	}
	assert.Equal(t, expected3, findHiddenSinglesInNineCells(cells3), "Two solutions expected")

	cells4 := [9]int{0, 6, 0, 0, 0, 0, 2, 0, 0}
	expected4 := []RelativeCellSolution{
		{1, 3},
	}
	assert.Equal(t, expected4, findHiddenSinglesInNineCells(cells4), "One solution expected")
}

func Test_findHiddenSingles(t *testing.T) {
	notes1 := [9][9]int{
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
	assert.Equal(t, []AbsoluteCellSolution{}, findHiddenSingles(notes1), "No solution expected")

	notes2 := [9][9]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 6, 2, 0, 0, 0, 2, 2, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 2, 0, 0, 0, 2, 1, 1},
		{0, 2, 0, 0, 0, 0, 0, 2, 0},
	}
	expected2 := []AbsoluteCellSolution{
		AbsoluteCellSolution{7, 7, 1, "Hidden Single", "Cell"},
		AbsoluteCellSolution{1, 1, 3, "Hidden Single", "Cell"},
		AbsoluteCellSolution{7, 0, 1, "Hidden Single", "Cell"},
	}
	assert.ElementsMatch(t, expected2, findHiddenSingles(notes2), "Two solutions expected")
}

func Test_findNakedPairsInNineCells(t *testing.T) {
	cells1 := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, []RelativeCellSolutions{}, findNakedPairsInNineCells(cells1), "No naked pairs expected")

	cells2 := [9]int{6, 0x101, 0, 6, 6, 0, 0, 0x101, 0}
	expected2 := []RelativeCellSolutions{RelativeCellSolutions{[]int{1, 7}, 0x101, "Cell"}}
	assert.Equal(t, expected2, findNakedPairsInNineCells(cells2), "One naked pair expected")

	cells3 := [9]int{0, 0x101, 0, 6, 0, 6, 0, 0x101, 0}
	expected3 := []RelativeCellSolutions{
		RelativeCellSolutions{[]int{1, 7}, 0x101, "Cell"},
		RelativeCellSolutions{[]int{3, 5}, 6, "Cell"},
	}
	assert.ElementsMatch(t, expected3, findNakedPairsInNineCells(cells3), "Two naked pairs expected")
}

func Test_findNakedPairExclusions(t *testing.T) {
	notes1 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 3, 0, 3, 0, 0, 3, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{4, 12, 0, 0, 0, 12, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected1 := []CellExclusion{
		CellExclusion{
			number:       12,
			matches:      []CellRef{CellRef{6, 1}, CellRef{6, 5}},
			removeNumber: 12,
			exclusions:   []CellRef{CellRef{6, 0}, CellRef{6, 7}},
			strategy:     "Naked Pairs"},
	}
	assert.Equal(t, expected1, findNakedPairExclusions(notes1), "One naked pair expected")

	notes2 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 4, 0, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 12, 0, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 12, 0, 0},
		{0, 3, 0, 0, 0, 0, 8, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected2 := []CellExclusion{
		CellExclusion{
			number:       12,
			matches:      []CellRef{CellRef{2, 6}, CellRef{6, 6}},
			removeNumber: 12,
			exclusions:   []CellRef{CellRef{0, 6}, CellRef{7, 6}},
			strategy:     "Naked Pairs"},
	}
	assert.Equal(t, expected2, findNakedPairExclusions(notes2), "One naked pair expected")

	notes3 := [9][9]int{
		{3, 0, 3, 0, 0, 0, 0, 0, 0},
		{4, 12, 12, 0, 0, 0, 0, 4, 0},
		{3, 8, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected3 := []CellExclusion{
		CellExclusion{
			number:       12,
			matches:      []CellRef{CellRef{1, 1}, CellRef{1, 2}},
			removeNumber: 12,
			exclusions:   []CellRef{CellRef{1, 0}, CellRef{1, 7}},
			strategy:     "Naked Pairs"},
		CellExclusion{
			number:       12,
			matches:      []CellRef{CellRef{1, 1}, CellRef{1, 2}},
			removeNumber: 12,
			exclusions:   []CellRef{CellRef{1, 0}, CellRef{2, 1}},
			strategy:     "Naked Pairs"},
		CellExclusion{
			number:       3,
			matches:      []CellRef{CellRef{0, 0}, CellRef{2, 0}},
			removeNumber: 3,
			exclusions:   []CellRef{CellRef{7, 0}},
			strategy:     "Naked Pairs"},
	}
	assert.ElementsMatch(t, expected3, findNakedPairExclusions(notes3), "Three naked pairs expected")
}

func Test_findHiddenPairsInNineCells(t *testing.T) {
	cells1 := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, []RelativeCellSolutions{}, findHiddenPairsInNineCells(cells1), "No hidden pairs expected")

	cells2 := [9]int{0, 0x03, 0, 0x07, 0, 0, 0, 0, 0}
	expected2 := []RelativeCellSolutions{RelativeCellSolutions{[]int{1, 3}, 0x03, "Cell"}}
	assert.Equal(t, expected2, findHiddenPairsInNineCells(cells2), "One hidden pairs expected")

	cells3 := [9]int{0, 0x0b, 0, 0x07, 0, 0, 0x04, 0, 0}
	expected3 := []RelativeCellSolutions{RelativeCellSolutions{[]int{1, 3}, 0x03, "Cell"}}
	assert.Equal(t, expected3, findHiddenPairsInNineCells(cells3), "One hidden pairs expected")

	cells4 := [9]int{0, 0x109, 0, 0, 0, 0, 0, 0x105, 0}
	expected4 := []RelativeCellSolutions{RelativeCellSolutions{[]int{1, 7}, 0x101, "Cell"}}
	assert.Equal(t, expected4, findHiddenPairsInNineCells(cells4), "One hidden pair expected")

	cells5 := [9]int{0, 0x109, 0, 0x0e, 0, 0x06, 0, 0x111, 0}
	expected5 := []RelativeCellSolutions{
		RelativeCellSolutions{[]int{1, 7}, 0x101, "Cell"},
		RelativeCellSolutions{[]int{3, 5}, 0x06, "Cell"},
	}
	assert.Equal(t, expected5, findHiddenPairsInNineCells(cells5), "Two hidden pairs expected")

	//naked pairs should not be found
	cells6 := [9]int{0, 0x03, 0, 0x03, 0, 0, 0, 0, 0}
	expected6 := []RelativeCellSolutions{}
	assert.Equal(t, expected6, findHiddenPairsInNineCells(cells6), "No hidden pairs expected")

}

func Test_findHiddenPairExclusions(t *testing.T) {
	notes1 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 14, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 14, 0, 0},
		{0, 0, 0, 13, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected1 := []CellExclusion{
		CellExclusion{
			number:       6,
			matches:      []CellRef{CellRef{6, 1}, CellRef{6, 6}},
			removeNumber: 9,
			exclusions:   []CellRef{CellRef{6, 1}, CellRef{6, 6}},
			strategy:     "Hidden Pairs"},
		CellExclusion{
			number:       12,
			matches:      []CellRef{CellRef{1, 3}, CellRef{7, 3}},
			removeNumber: 3,
			exclusions:   []CellRef{CellRef{1, 3}, CellRef{7, 3}},
			strategy:     "Hidden Pairs"},
		CellExclusion{
			number:       3,
			matches:      []CellRef{CellRef{3, 6}, CellRef{4, 7}},
			removeNumber: 4,
			exclusions:   []CellRef{CellRef{3, 6}, CellRef{4, 7}},
			strategy:     "Hidden Pairs"},
	}
	assert.ElementsMatch(t, expected1, findHiddenPairExclusions(notes1), "Three hidden pair expected")
}

func Test_findPointingPairsInNineCellBlock(t *testing.T) {
	block := [9]int{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}
	assert.Equal(t, []RelativeCellSolutions{}, findPointingPairsInNineCellBlock(block), "No solution expected")

	block2 := [9]int{
		0, 9, 1,
		6, 0, 0,
		4, 0, 0,
	}
	expected2 := []RelativeCellSolutions{
		RelativeCellSolutions{[]int{1, 2}, 0x01, "Row"},
		RelativeCellSolutions{[]int{3, 6}, 0x04, "Column"},
	}
	assert.ElementsMatch(t, expected2, findPointingPairsInNineCellBlock(block2), "One solution expected")
}

func Test_findPointingPairExclusions(t *testing.T) {
	notes1 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected1 := []CellExclusion{
		CellExclusion{
			number:       1,
			matches:      []CellRef{CellRef{4, 4}, CellRef{5, 4}},
			removeNumber: 1,
			exclusions:   []CellRef{CellRef{7, 4}},
			strategy:     "Pointing Pairs"},
	}
	assert.Equal(t, expected1, findPointingPairExclusions(notes1), "One Pointing pair in row expected")

	notes2 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 5, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected2 := []CellExclusion{
		CellExclusion{
			number:       1,
			matches:      []CellRef{CellRef{4, 1}, CellRef{4, 2}},
			removeNumber: 1,
			exclusions:   []CellRef{CellRef{4, 6}},
			strategy:     "Pointing Pairs"},
	}
	assert.Equal(t, expected2, findPointingPairExclusions(notes2), "One Pointing pair in column expected")

	notes3 := [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 0, 4, 6, 0, 0, 4, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected3 := []CellExclusion{
		CellExclusion{
			number:       1,
			matches:      []CellRef{CellRef{3, 5}, CellRef{4, 5}},
			removeNumber: 1,
			exclusions:   []CellRef{CellRef{7, 5}},
			strategy:     "Pointing Pairs"},
		CellExclusion{
			number:       4,
			matches:      []CellRef{CellRef{6, 3}, CellRef{6, 4}},
			removeNumber: 4,
			exclusions:   []CellRef{CellRef{6, 1}, CellRef{6, 7}},
			strategy:     "Pointing Pairs"},
	}
	assert.Equal(t, expected3, findPointingPairExclusions(notes3), "Two Pointing pairs expected")
}

func Test_findBoxPairsInNineCellLine(t *testing.T) {
	block := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	assert.Equal(t, []RelativeCellSolutions{}, findBoxPairsInNineCellLine(block), "No solution expected")

	block2 := [9]int{0, 1, 1, 0, 0, 0, 0, 0, 0}
	expected2 := []RelativeCellSolutions{
		RelativeCellSolutions{[]int{1, 2}, 0x01, "Box Pair"},
	}
	assert.Equal(t, expected2, findBoxPairsInNineCellLine(block2), "One solution expected")

	block3 := [9]int{0, 1, 1, 0, 0, 0, 1, 0, 0}
	expected3 := []RelativeCellSolutions{}
	assert.Equal(t, expected3, findBoxPairsInNineCellLine(block3), "No solution expected")

	block4 := [9]int{0, 4, 4, 0, 0, 0, 2, 2, 0}
	expected4 := []RelativeCellSolutions{
		RelativeCellSolutions{[]int{1, 2}, 0x04, "Box Pair"},
		RelativeCellSolutions{[]int{6, 7}, 0x02, "Box Pair"},
	}
	assert.ElementsMatch(t, expected4, findBoxPairsInNineCellLine(block4), "Two solutions expected")

	block5 := [9]int{0, 0, 1, 1, 0, 0, 0, 0, 0}
	expected5 := []RelativeCellSolutions{}
	assert.Equal(t, expected5, findBoxPairsInNineCellLine(block5), "No solution expected")
}
func Test_findBoxLineReductionExclusions(t *testing.T) {
	notes1 := [9][9]int{
		{0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected1 := []CellExclusion{
		CellExclusion{
			number:       1,
			matches:      []CellRef{CellRef{0, 1}, CellRef{0, 2}},
			removeNumber: 1,
			exclusions:   []CellRef{CellRef{1, 1}, CellRef{2, 2}},
			strategy:     "Box Line Reduction"},
	}
	assert.Equal(t, expected1, findBoxLineReductionExclusions(notes1), "One box line reduction expected")

	notes2 := [9][9]int{
		{0, 1, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected2 := []CellExclusion{
		CellExclusion{
			number:       1,
			matches:      []CellRef{CellRef{0, 1}, CellRef{1, 1}},
			removeNumber: 1,
			exclusions:   []CellRef{CellRef{0, 2}, CellRef{2, 2}},
			strategy:     "Box Line Reduction"},
	}
	assert.Equal(t, expected2, findBoxLineReductionExclusions(notes2), "One box line reduction expected")

	notes3 := [9][9]int{
		{0, 1, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 2, 2, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
	}
	expected3 := []CellExclusion{
		CellExclusion{
			number:       1,
			matches:      []CellRef{CellRef{0, 1}, CellRef{1, 1}},
			removeNumber: 1,
			exclusions:   []CellRef{CellRef{0, 2}, CellRef{2, 2}},
			strategy:     "Box Line Reduction"},
		CellExclusion{
			number:       2,
			matches:      []CellRef{CellRef{6, 6}, CellRef{6, 7}},
			removeNumber: 2,
			exclusions:   []CellRef{CellRef{8, 8}},
			strategy:     "Box Line Reduction"},
	}
	assert.ElementsMatch(t, expected3, findBoxLineReductionExclusions(notes3), "Two box line reductions expected")
}
