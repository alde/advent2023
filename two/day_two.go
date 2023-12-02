package two

import (
	"strconv"
	"strings"

	"alde.nu/advent2023/shared"
)

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func (c *Cubes) Add(color string, amount int) {
	if color == "red" {
		c.Red += amount
	}
	if color == "blue" {
		c.Blue += amount
	}
	if color == "green" {
		c.Green += amount
	}
}

type Hand struct {
	Color  string
	Amount int
}

func isValidRound(round string, cubesAvailable *Cubes) bool {
	cubesInRound := &Cubes{}
	for _, hand := range strings.Split(round, ", ") {
		h := strings.Split(hand, " ")
		amount, _ := strconv.Atoi(h[0])
		color := h[len(h)-1]

		cubesInRound.Add(color, amount)
	}

	if cubesInRound.Blue > cubesAvailable.Blue {
		return false
	}
	if cubesInRound.Red > cubesAvailable.Red {
		return false
	}
	if cubesInRound.Green > cubesAvailable.Green {
		return false
	}

	return true
}

func PartOne(input []string, cubes *Cubes) *shared.Result[int] {
	result := 0
	for _, game := range input {
		record := strings.Split(game, ": ")
		if len(record) == 1 {
			continue
		}
		gameId, _ := strconv.Atoi(strings.Split(record[0], " ")[1])
		allRoundsValid := true
		for _, round := range strings.Split(record[1], "; ") {
			isValid := isValidRound(round, cubes)
			if !isValid {
				allRoundsValid = false
			}
		}
		if allRoundsValid {
			result += gameId
		}

	}
	return &shared.Result[int]{Day: "Two", Task: "One", Value: result}
}
func minReqForGame(game string) *Cubes {
	minCubes := &Cubes{}
	for _, round := range strings.Split(game, "; ") {
		for _, hand := range strings.Split(round, ", ") {
			h := strings.Split(hand, " ")
			amount, _ := strconv.Atoi(h[0])
			color := h[len(h)-1]
			if color == "red" && amount >= minCubes.Red {
				minCubes.Red = amount
			}
			if color == "green" && amount >= minCubes.Green {
				minCubes.Green = amount
			}
			if color == "blue" && amount >= minCubes.Blue {
				minCubes.Blue = amount
			}
		}
	}

	return minCubes
}
func PartTwo(input []string) *shared.Result[int] {
	result := 0
	for _, game := range input {
		record := strings.Split(game, ": ")
		if len(record) == 1 {
			continue
		}
		reqCubes := minReqForGame(record[1])

		result += reqCubes.Red * reqCubes.Green * reqCubes.Blue

	}
	return &shared.Result[int]{Day: "Two", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInput(input)

	shared.PrintResult(PartOne(strings.Split(data, "\n"), &Cubes{Red: 12, Green: 13, Blue: 14}))
	shared.PrintResult(PartTwo(strings.Split(data, "\n")))
}
