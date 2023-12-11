package day_05_test

import (
	"fmt"
	"testing"

	"alde.nu/advent2023/day_05"
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
	actual := day_05.MakeMap(input)
	assert.Equal(t, 1, actual.Get(1))
	assert.Equal(t, 50, actual.Get(98))
	assert.Equal(t, 51, actual.Get(99))
	for i := 0; i < 48; i++ {
		assert.Equal(t, 52+i, actual.Get(50+i))
	}
}

func Test_ProcessInput(t *testing.T) {
	actual := day_05.ProcessInput(INPUT)

	assert.Equal(t, actual.SeedList, []int{79, 14, 55, 13})

	testData := []struct {
		seed     int
		location int
	}{
		{seed: 79, location: 82},
		{seed: 14, location: 43},
		{seed: 55, location: 86},
		{seed: 13, location: 35},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			location := actual.Traverse(td.seed)
			assert.Equal(t, td.location, location)
		})
	}
}

func Test_Reverse(t *testing.T) {
	actual := day_05.ProcessInput(INPUT)

	assert.Equal(t, actual.SeedListRange, []day_05.SeedRange{{79, 14}, {55, 13}})

	testData := []struct {
		actual   int
		location int
	}{
		{actual: 79, location: 82},
		{actual: 14, location: 43},
		{actual: 55, location: 86},
		{actual: 13, location: 35},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			seed := actual.Reverse(td.location)
			assert.Equal(t, td.actual, seed)
		})
	}
}

func Test_PartOne(t *testing.T) {
	almanac := day_05.ProcessInput(INPUT)
	res := day_05.PartOne(almanac)

	assert.Equal(t, 35, res.Value)
}

func Test_PartTwo(t *testing.T) {
	almanac := day_05.ProcessInput(INPUT)
	res := day_05.PartTwo(almanac)

	assert.Equal(t, 46, res.Value)
}
