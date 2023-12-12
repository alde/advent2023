package day_11_test

import (
	"fmt"

	"alde.nu/advent2023/day_11"
	"alde.nu/advent2023/shared"
	"github.com/stretchr/testify/assert"

	"testing"
)

var INPUT = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}
var testUniverse = [][]byte{
	{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
	{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
	{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
	{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
}

func Test_ParseInput(t *testing.T) {
	actual := day_11.ParseInput(INPUT)

	assert.Equal(t, testUniverse, actual)
}

func Test_ExpandUniverse(t *testing.T) {

	input := [][]byte{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '#'}}
	emptyRows, emptyColumns, galaxies := day_11.ExpandUniverse(input)

	assert.Equal(t, []int{0, 1}, emptyColumns)
	assert.Equal(t, []int{0, 1}, emptyRows)
	assert.Equal(t, []day_11.Coord{{2, 2}}, galaxies)
}

func Test_Solve(t *testing.T) {
	testData := []struct {
		multiplier int
		expected   int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actual := shared.Sum(day_11.Solve(testUniverse, td.multiplier))

			assert.Equal(t, td.expected, actual)
		})
	}

	actual := day_11.PartOne(day_11.ParseInput(INPUT), 2)
	expected := 374
	assert.Equal(t, expected, actual.Value)
}
