package day_03

import (
	"strconv"

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

func IsAdjacentToGear(matrix []string, row, column int) (bool, int, int) {
	for y := row - 1; y <= row+1; y++ {
		if y < 0 || y >= len(matrix) {
			continue
		}
		for x := column - 1; x <= column+1; x++ {
			if x < 0 || x >= len(matrix[y]) || (x == column && y == row) {
				continue
			}
			if matrix[y][x] == '*' {
				return true, x, y
			}
		}
	}

	return false, 0, 0
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

type Gear struct {
	X      int
	Y      int
	Number int
}

func CheckForGrears(matrix []string) []int {
	gearCandidates := []Gear{}
	for y := 0; y < len(matrix); y++ {
		tempString := ""
		isAdjacent := false
		gearX := 0
		gearY := 0
		for x := 0; x < len(matrix[y]); x++ {
			isDigit := IsDigit(matrix[y][x])
			if isDigit {
				tempString += string(matrix[y][x])
				if !isAdjacent {
					isAdjacent, gearX, gearY = IsAdjacentToGear(matrix, y, x)
				}
			}
			if !isDigit || len(matrix[y])-1 == x { // if end of row or not a digit
				if isAdjacent {
					num, _ := strconv.Atoi(tempString)
					gearCandidates = append(gearCandidates, Gear{Number: num, X: gearX, Y: gearY})
				}
				isAdjacent = false
				gearX = 0
				gearY = 0
				tempString = ""
			}

		}
	}

	gearRatios := []int{}
	for i := 0; i < len(gearCandidates); i++ {
		for j := i + 1; j < len(gearCandidates); j++ {
			if gearCandidates[i].X == gearCandidates[j].X && gearCandidates[i].Y == gearCandidates[j].Y {
				gearRatios = append(gearRatios, gearCandidates[i].Number*gearCandidates[j].Number)
				continue
			}
		}
	}

	return gearRatios
}

func PartOne(input []string) *shared.Result {
	result := 0
	for _, i := range CheckMatrix(input) {
		result += i
	}
	return &shared.Result{Day: "Three", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result {
	result := 0
	for _, i := range CheckForGrears(input) {
		result += i
	}
	return &shared.Result{Day: "Three", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
