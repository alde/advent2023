package nine_test

import (
	"fmt"

	nine "alde.nu/advent2023/day_9"

	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataRowOne = []int{
	0, 3, 6, 9, 12, 15,
}

func Test_CalculateDifferences(t *testing.T) {
	actual := nine.CaluclateDifferences(testDataRowOne)
	gen1 := []int{3, 3, 3, 3, 3}
	assert.Equal(t, gen1, actual)

	actual = nine.CaluclateDifferences(gen1)
	gen2 := []int{0, 0, 0, 0}
	assert.Equal(t, gen2, actual)
}

func Test_BuildHistory(t *testing.T) {
	actual := nine.BuildHistory(testDataRowOne)
	expect := [][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}

	assert.Equal(t, expect, actual)
}

func Test_ExtrapolateFuture(t *testing.T) {
	testData := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 3, 6, 9, 12, 15, 18},
				{3, 3, 3, 3, 3, 3},
				{0, 0, 0, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{1, 3, 6, 10, 15, 21, 28},
				{2, 3, 4, 5, 6, 7},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0},
			},
		},
		{
			input: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
			expected: [][]int{
				{10, 13, 16, 21, 30, 45, 68},
				{3, 3, 5, 9, 15, 23},
				{0, 2, 4, 6, 8},
				{2, 2, 2, 2},
				{0, 0, 0},
			},
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			actual := nine.ExtrapolateFuture(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}

func Test_ExtrapolatePast(t *testing.T) {
	testData := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{-3, 0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3, 3},
				{0, 0, 0, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 1, 3, 6, 10, 15, 21},
				{1, 2, 3, 4, 5, 6},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0},
			},
		},
		{
			input: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
			expected: [][]int{
				{5, 10, 13, 16, 21, 30, 45},
				{5, 3, 3, 5, 9, 15},
				{-2, 0, 2, 4, 6},
				{2, 2, 2, 2},
				{0, 0, 0},
			},
		},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			actual := nine.ExtrapolatePast(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}

var input = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func Test_PartOne(t *testing.T) {
	actual := nine.PartOne(input)
	assert.Equal(t, 114, actual.Value)
}

func Test_PartTwo(t *testing.T) {
	actual := nine.PartTwo(input)
	assert.Equal(t, 2, actual.Value)
}
