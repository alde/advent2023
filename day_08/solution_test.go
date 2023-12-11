package day_08_test

import (
	"testing"

	"alde.nu/advent2023/day_08"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

func Test_GetNodes(t *testing.T) {
	actual := day_08.GetNodes(INPUT[1:])
	assert.Equal(t, "BBB", actual["AAA"].Left)
	assert.Equal(t, "CCC", actual["AAA"].Right)

	assert.Equal(t, "ZZZ", actual["CCC"].Left)
	assert.Equal(t, "GGG", actual["CCC"].Right)
}

func Test_GetDirections(t *testing.T) {
	actual := day_08.GetDirections(INPUT[0])
	assert.Equal(t, 'R', actual.Pop())
	assert.Equal(t, 'L', actual.Pop())
	assert.Equal(t, 'R', actual.Pop())
}

func Test_PartOne(t *testing.T) {
	actual := day_08.PartOne(day_08.GetDirections(INPUT[0]), day_08.GetNodes(INPUT[1:]))
	assert.Equal(t, 2, actual.Value)

	actual = day_08.PartOne(day_08.GetDirections("LLR"), day_08.GetNodes([]string{
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}))
	assert.Equal(t, 6, actual.Value)
}

var secondInput = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func Test_GetStartingPositions(t *testing.T) {
	nodes := day_08.GetNodes(secondInput[1:])
	actual := day_08.GetStartingPositions(nodes)
	assert.Equal(t, []string{"11A", "22A"}, actual)
}

func Test_PartTwo(t *testing.T) {
	actual := day_08.TraverseMultiple([]string{"11A", "22A"}, day_08.GetDirections(secondInput[0]), day_08.GetNodes(secondInput[1:]))
	assert.Equal(t, 6, actual)
}
