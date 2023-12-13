package day_12_test

import (
	"fmt"
	"testing"

	"alde.nu/advent2023/day_12"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func Test_Parse(t *testing.T) {
	testData := []struct {
		record string
		groups []int
	}{
		{"???.###", []int{1, 1, 3}},
		{".??..??...?##.", []int{1, 1, 3}},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}},
		{"????.#...#...", []int{4, 1, 1}},
		{"????.######..#####.", []int{1, 6, 5}},
		{"?###????????", []int{3, 2, 1}},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actualRecord, actualGroups := day_12.Parse(INPUT[i])
			r2, g2 := day_12.Parse(INPUT[i])
			assert.Equal(t, td.record, actualRecord)
			assert.Equal(t, td.groups, actualGroups)
			assert.Equal(t, r2, actualRecord)
			assert.Equal(t, g2, actualGroups)
		})
	}
}

func Test_Parse2(t *testing.T) {
	expectedRecord := "#?.???.?###??####.?."
	expectedGroups := []int{2, 3, 10, 1}
	input := "#?.???.?###??####.?. 2,3,10,1"
	r1, g1 := day_12.Parse(input)
	assert.Equal(t, expectedRecord, r1)
	assert.Equal(t, expectedGroups, g1)
}

func Test_GetCandidates(t *testing.T) {
	testData := []struct {
		record     string
		groups     []int
		candidates [][]int
	}{
		{"???.###", []int{1, 1, 3}, [][]int{{0, 1, 2}, {0, 1, 2}, {0, 4}}},
		{".??..??...?##.", []int{1, 1, 3}, [][]int{{1, 2, 5, 6}, {1, 2, 5, 6}, {10}}},
		{"?###????????", []int{3, 2, 1}, [][]int{{1, 5, 6, 7, 8, 9}, {5, 6, 7, 8, 9, 10}, {5, 6, 7, 8, 9, 10, 11}}},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actual := day_12.GetCandidates(td.record, td.groups)
			assert.Equal(t, td.candidates, actual)
		})
	}
}

func Test_ValidArrangements(t *testing.T) {
	testData := []struct {
		validCombinations int
	}{
		{1},
		{4},
		{1},
		{1},
		{4},
		{10},
	}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actual := day_12.ValidArrangements(INPUT[i], false)
			assert.Equal(t, td.validCombinations, actual)
		})
	}
}

func Test_ValidArrangements_2(t *testing.T) {
	testData := []struct {
		data     string
		expected int
	}{{data: "#?.???.?###??####.?. 2,3,10,1", expected: 1}}
	for i, td := range testData {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			actual := day_12.ValidArrangements(td.data, false)
			assert.Equal(t, td.expected, actual)
		})
	}
}

func Test_PartOne(t *testing.T) {
	actual := day_12.PartOne(INPUT)
	assert.Equal(t, 21, actual.Value)
}
