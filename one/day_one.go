package one

import (
	"strings"

	"alde.nu/advent2023/shared"
)

func PartOne(input []string) *shared.Result[int] {
	total := 0
	for _, s := range input {
		nums := shared.ExtractNumbers(s)
		total += shared.MergeNumbers(nums)
	}

	return &shared.Result[int]{Day: "One", Task: "One", Value: total}
}

func PartTwo(input []string) *shared.Result[int] {
	total := 0
	for _, s := range input {
		nums := shared.ExtractNumbersRedux(s)
		total += shared.MergeNumbers(nums)
	}

	return &shared.Result[int]{Day: "One", Task: "Two", Value: total}
}

func Run(input string) {
	data := shared.LoadInput(input)

	shared.PrintResult(PartOne(strings.Split(data, "\n")))
	shared.PrintResult(PartTwo(strings.Split(data, "\n")))
}
