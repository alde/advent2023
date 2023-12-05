package four

import (
	"strconv"
	"strings"

	"alde.nu/advent2023/shared"
)

type Card struct {
	Number       int
	WinningCount int
}

func intersection(inputs ...[]int) []int {
	hash := make(map[int]*int) // value, counter
	result := make([]int, 0)
	for _, slice := range inputs {
		duplicationHash := make(map[int]bool) // duplication checking for individual slice
		for _, value := range slice {
			if _, isDup := duplicationHash[value]; !isDup { // is not duplicated in slice
				if counter := hash[value]; counter != nil { // is found in hash counter map
					if *counter++; *counter >= len(inputs) { // is found in every slice
						result = append(result, value)
					}
				} else { // not found in hash counter map
					i := 1
					hash[value] = &i
				}
				duplicationHash[value] = true
			}
		}
	}
	return result
}

func ParseCard(row string) *Card {
	split1 := strings.Split(row, ":")

	cardNumber, _ := strconv.Atoi(strings.TrimLeft(split1[0], "Card "))
	split2 := strings.Split(split1[1], "|")
	winningNumbers := shared.ConvertToNumSlice(split2[0])
	playedNumbers := shared.ConvertToNumSlice(split2[1])

	winCount := len(intersection(winningNumbers, playedNumbers))

	return &Card{
		Number:       cardNumber,
		WinningCount: winCount,
	}
}

func makeCards(input []string) []*Card {
	cards := []*Card{}
	for _, row := range input {
		if len(row) <= 1 {
			continue
		}
		cards = append(cards, ParseCard(row))
	}

	return cards
}

func calculatePoints(c *Card) int {
	if c.WinningCount == 0 {
		return 0
	}

	result := 1

	for i := 1; i < c.WinningCount; i++ {
		result *= 2
	}

	return result
}

func countCards(cards []*Card, cardMap map[int]*Card, counter int) int {
	if len(cards) == 0 {
		return counter
	}

	head, tail := cards[0], cards[1:]
	for len(tail) >= 0 {
		if head.WinningCount > 0 {
			for n := 1; n <= head.WinningCount; n++ {
				tail = append(tail, cardMap[head.Number+n])
			}
		}
		counter += 1
		if len(tail) == 0 {
			break
		}
		head, tail = tail[0], tail[1:]
	}

	return counter
}

func PartOne(input []string) *shared.Result[int] {
	cards := makeCards(input)
	result := 0
	for _, c := range cards {
		result += calculatePoints(c)
	}
	return &shared.Result[int]{Day: "Four", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result[int] {
	cards := makeCards(input)
	cardMap := make(map[int]*Card)
	for _, card := range cards {
		cardMap[card.Number] = card
	}

	result := countCards(cards, cardMap, 0)

	return &shared.Result[int]{Day: "Four", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInput(input)

	shared.PrintResult(PartOne(strings.Split(data, "\n")))
	shared.PrintResult(PartTwo(strings.Split(data, "\n")))

}
