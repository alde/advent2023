package five_test

import (
	"fmt"
	"testing"

	"alde.nu/advent2023/five"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"seeds: 79 14 55 13",
	"	",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"	",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"	",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"	",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"	",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"	",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"	",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

func Test_MakeMap(t *testing.T) {
	input := []string{
		"50 98 2",
		"52 50 48",
	}
	actual := five.MakeMap(input)
	assert.Equal(t, 1, actual.Get(1))
	assert.Equal(t, 50, actual.Get(98))
	assert.Equal(t, 51, actual.Get(99))
	for i := 0; i < 48; i++ {
		assert.Equal(t, 52+i, actual.Get(50+i))
	}
}

func Test_ProcessInput(t *testing.T) {
	actual := five.ProcessInput(INPUT)

	assert.Equal(t, actual.SeedList, []int{79, 14, 55, 13})

	testData := []struct {
		seed int
		path []int
	}{
		{seed: 79, path: []int{79, 81, 81, 81, 74, 78, 78, 82}},
		{seed: 14, path: []int{14, 14, 53, 49, 42, 42, 43, 43}},
		{seed: 55, path: []int{55, 57, 57, 53, 46, 82, 82, 86}},
		{seed: 13, path: []int{13, 13, 52, 41, 34, 34, 35, 35}},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			path := actual.ResolveMappings(td.seed)
			assert.Equal(t, td.path, path)
		})
	}
}

func Test_Reverse(t *testing.T) {
	actual := five.ProcessInputPartTwo(INPUT)

	assert.Equal(t, actual.SeedListRange, []five.SeedRange{{79, 14}, {55, 13}})

	testData := []struct {
		seed int
		path []int
	}{
		{seed: 79, path: []int{79, 81, 81, 81, 74, 78, 78, 82}},
		{seed: 14, path: []int{14, 14, 53, 49, 42, 42, 43, 43}},
		{seed: 55, path: []int{55, 57, 57, 53, 46, 82, 82, 86}},
		{seed: 13, path: []int{13, 13, 52, 41, 34, 34, 35, 35}},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			path := actual.ResolveMappings(td.seed)
			assert.Equal(t, td.path[len(td.path)-1], path)
		})
	}
}

func Test_PartOne(t *testing.T) {
	res := five.PartOne(INPUT)

	assert.Equal(t, 35, res.Value)
}

func Test_PartTwo(t *testing.T) {
	res := five.PartTwo(INPUT)

	assert.Equal(t, 46, res.Value)
}
