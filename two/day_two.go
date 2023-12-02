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

func (c *Cubes) Add(hand *Hand) {
	if hand.Color == "red" {
		c.Red += hand.Amount
	}
	if hand.Color == "blue" {
		c.Blue += hand.Amount
	}
	if hand.Color == "green" {
		c.Green += hand.Amount
	}
}

type Hand struct {
	Color  string
	Amount int
}

func NewHand(s string) *Hand {
	hand := &Hand{}
	split := strings.Split(s, " ")
	amount, _ := strconv.Atoi(split[0])
	hand.Amount = amount
	color := split[len(split)-1]
	hand.Color = color
	return hand
}

func makeHands(round string) []*Hand {
	res := []*Hand{}
	for _, hand := range strings.Split(round, ", ") {
		res = append(res, NewHand(hand))
	}

	return res
}

func isValidRound(round string, cubesAvailable *Cubes) bool {
	cubesInRound := &Cubes{}
	hands := makeHands(round)
	for _, hand := range hands {
		cubesInRound.Add(hand)
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
		if len(record) <= 1 {
			continue
		}
		gameId, _ := strconv.Atoi(strings.Split(record[0], " ")[1])
		allRoundsValid := true
		for _, round := range strings.Split(record[1], "; ") {
			if !isValidRound(round, cubes) {
				allRoundsValid = false
				break
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
		hands := makeHands(round)
		for _, hand := range hands {
			if hand.Color == "red" && hand.Amount >= minCubes.Red {
				minCubes.Red = hand.Amount
			}
			if hand.Color == "green" && hand.Amount >= minCubes.Green {
				minCubes.Green = hand.Amount
			}
			if hand.Color == "blue" && hand.Amount >= minCubes.Blue {
				minCubes.Blue = hand.Amount
			}
		}
	}

	return minCubes
}

func PartTwo(input []string) *shared.Result[int] {
	result := 0
	for _, game := range input {
		record := strings.Split(game, ": ")
		if len(record) <= 1 {
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
