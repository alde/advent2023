package day_11

import (
	"alde.nu/advent2023/shared"
)

func Convert(input []string) [][]byte {
	res := [][]byte{}
	for _, row := range input {
		r := []byte{}
		for _, col := range row {
			r = append(r, byte(col))
		}
		res = append(res, r)
	}
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			res[row][col] = input[row][col]
		}
	}
	return res
}
func emptyRow(row []byte) bool {
	for _, c := range row {
		if c != '.' {
			return false
		}
	}
	return true
}
func emptyColumn(galaxy [][]byte, column int) bool {
	for _, row := range galaxy {
		if row[column] != '.' {
			return false
		}
	}
	return true
}
func insertColumn(expanded, galaxy [][]byte, column int) [][]byte {
	for r := 0; r < len(galaxy); r++ {
		expanded[column] = append(expanded[column], galaxy[column][r])
	}
	return expanded
}

func ExpandGalaxy(input [][]byte) [][]byte {
	expanded := [][]byte{}
	for row := 0; row < len(input); row++ {
		if emptyRow(input[row]) {
			expanded = append(expanded, input[row])
		}
		expanded = append(expanded, input[row])
	}
	for column := 0; column < len(input[0]); column++ {
		if emptyColumn(input, column) {
			expanded = insertColumn(expanded, input, column)
		}
		expanded = insertColumn(expanded, input, column)
	}

	return expanded
}
func PartOne(galaxy []string) *shared.Result {
	result := 0
	return &shared.Result{Day: "Eleven", Task: "One", Value: result}
}

func PartTwo(galaxy []string) *shared.Result {
	result := 0
	return &shared.Result{Day: "Eleven", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
