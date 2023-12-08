package eight_test

import (
	"testing"

	eight "alde.nu/advent2023/day_8"
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
	actual := eight.GetNodes(INPUT[1:])
	assert.Equal(t, "BBB", actual["AAA"].Left)
	assert.Equal(t, "CCC", actual["AAA"].Right)

	assert.Equal(t, "ZZZ", actual["CCC"].Left)
	assert.Equal(t, "GGG", actual["CCC"].Right)
}

func Test_GetDirections(t *testing.T) {
	actual := eight.GetDirections(INPUT[0])
	assert.Equal(t, 'R', actual.Pop())
	assert.Equal(t, 'L', actual.Pop())
	assert.Equal(t, 'R', actual.Pop())
}

func Test_PartOne(t *testing.T) {
	actual := eight.PartOne(eight.GetDirections(INPUT[0]), eight.GetNodes(INPUT[1:]))
	assert.Equal(t, 2, actual.Value)

	actual = eight.PartOne(eight.GetDirections("LLR"), eight.GetNodes([]string{
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
	nodes := eight.GetNodes(secondInput[1:])
	actual := eight.GetStartingPositions(nodes)
	assert.Equal(t, []string{"11A", "22A"}, actual)
}

func Test_PartTwo(t *testing.T) {
	actual := eight.TraverseMultiple([]string{"11A", "22A"}, eight.GetDirections(secondInput[0]), eight.GetNodes(secondInput[1:]))
	assert.Equal(t, 6, actual)
}
