package day_10_test

import (
	"alde.nu/advent2023/day_10"

	"testing"

	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

var I2 = []string{
	".....",
	".S-7.",
	".|.|.",
	".L-J.",
	".....",
}

func Test_PartOne(t *testing.T) {
	actual := day_10.PartOne(INPUT)
	assert.Equal(t, 8, actual.Value)

	actual = day_10.PartOne(I2)
	assert.Equal(t, 4, actual.Value)
}

var I3 = []string{
	"FF7FSF7F7F7F7F7F---7",
	"L|LJ||||||||||||F--J",
	"FL-7LJLJ||||||LJL-77",
	"F--JF--7||LJLJ7F7FJ-",
	"L---JF-JLJ.||-FJLJJ7",
	"|F|F-JF---7F7-L7L|7|",
	"|FFJF7L7F-JF7|JL---7",
	"7-L-JL7||F7|L7F-7F7|",
	"L.L7LFJ|||||FJL7||LJ",
	"L7JLJL-JLJLJL--JLJ.L",
}

func Test_PartTwo(t *testing.T) {
	actual := day_10.PartTwo(I3)
	assert.Equal(t, 10, actual.Value)
}
