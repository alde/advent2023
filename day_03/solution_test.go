package day_03_test

import (
	"fmt"
	"testing"

	"alde.nu/advent2023/day_03"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func Test_PartOne(t *testing.T) {
	res := day_03.PartOne(INPUT)

	assert.Equal(t, 4361, res.Value)
}
func Test_PartTwo(t *testing.T) {
	res := day_03.PartTwo(INPUT)

	assert.Equal(t, 467835, res.Value)
}
func Test_CheckMatrix(t *testing.T) {
	testData := []struct {
		inputString []string
		expected    []int
	}{
		{inputString: []string{"123...", "..+..."}, expected: []int{123}},
		{inputString: []string{".123..", "432.."}, expected: []int{}},
		{inputString: []string{"123+345", "......."}, expected: []int{123, 345}},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("test for %v", td.inputString), func(t *testing.T) {
			assert.Equal(t, td.expected, day_03.CheckMatrix(td.inputString))
		})
	}

	assert.Equal(t, []int{467, 35, 633, 617, 592, 755, 664, 598}, day_03.CheckMatrix(INPUT))
}

func Test_CheckForGears(t *testing.T) {
	testData := []struct {
		inputString []string
		expected    []int
	}{
		{inputString: []string{"123...", "..*...", "456..."}, expected: []int{123 * 456}},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("test for %v", td.inputString), func(t *testing.T) {
			assert.Equal(t, td.expected, day_03.CheckForGrears(td.inputString))
		})
	}

	assert.Equal(t, []int{16345, 451490}, day_03.CheckForGrears(INPUT))
}

func Test_IsDigit(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{input: "/", expected: false},
		{input: "0", expected: true},
		{input: "1", expected: true},
		{input: "2", expected: true},
		{input: "3", expected: true},
		{input: "4", expected: true},
		{input: "5", expected: true},
		{input: "6", expected: true},
		{input: "7", expected: true},
		{input: "8", expected: true},
		{input: "9", expected: true},
		{input: ":", expected: false},
		{input: "A", expected: false},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("test for %v", td.input), func(t *testing.T) {
			assert.Equal(t, td.expected, day_03.IsDigit(td.input[0]))
		})
	}
}
