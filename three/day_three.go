package three

import (
	"strconv"
	"strings"

	"alde.nu/advent2023/shared"
)

func IsDigit(char byte) bool {
	return char >= 48 && char <= 57
}

func IsAdjacentToSymbol(matrix []string, row, column int) bool {
	for y := row - 1; y <= row+1; y++ {
		if y < 0 || y >= len(matrix) {
			continue
		}
		for x := column - 1; x <= column+1; x++ {
			if x < 0 || x >= len(matrix[y]) || (x == column && y == row) {
				continue
			}
			if !IsDigit(matrix[y][x]) && matrix[y][x] != '.' {
				return true
			}
		}
	}

	return false
}

func CheckMatrix(matrix []string) []int {
	// for each row (y)
	// for each column (x)
	// if isDigit(m[y][x])
	// save to temp string
	// if isAdjacent to a symbol
	// mark as being a part number
	// else if is period
	// stop checking if it's a number
	// if isPartNumber
	// add to collection
	numberCollection := []int{}
	for y := 0; y < len(matrix); y++ {
		tempString := ""
		isAdjacent := false
		for x := 0; x < len(matrix[y]); x++ {
			isDigit := IsDigit(matrix[y][x])
			if isDigit {
				tempString += string(matrix[y][x])
				if !isAdjacent {
					isAdjacent = IsAdjacentToSymbol(matrix, y, x)
				}
			}
			if !isDigit || len(matrix[y])-1 == x { // if end of row or not a digit
				if isAdjacent {
					num, _ := strconv.Atoi(tempString)
					numberCollection = append(numberCollection, num)
				}
				isAdjacent = false
				tempString = ""
			}

		}
	}
	return numberCollection
}

func PartOne(input []string) *shared.Result[int] {
	result := 0
	for _, i := range CheckMatrix(input) {
		result += i
	}
	return &shared.Result[int]{Day: "Three", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result[int] {

	return &shared.Result[int]{Day: "Three", Task: "Two", Value: 0}
}

func Run(input string) {
	data := shared.LoadInput(input)

	shared.PrintResult(PartOne(strings.Split(data, "\n")))
	// shared.PrintResult(PartTwo(strings.Split(data, "\n")))
}
