package day_11_test

import (
	"alde.nu/advent2023/day_11"
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

func Test_ExpandGalaxy(t *testing.T) {
	galaxy := day_11.Convert(INPUT)
	actual := day_11.ExpandGalaxy(galaxy)
	expected := day_11.Convert([]string{
		"....#........",
		".........#...",
		"#............",
		".............",
		".............",
		"........#....",
		".#...........",
		"............#",
		".............",
		".............",
		".........#...",
		"#....#.......",
	})

	assert.Equal(t, expected, actual)
}
