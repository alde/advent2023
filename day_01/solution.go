package day_01

import (
	"alde.nu/advent2023/shared"
)

func PartOne(input []string) *shared.Result {
	total := 0
	for _, s := range input {
		nums := shared.ExtractNumbers(s)
		total += shared.MergeNumbers(nums)
	}

	return &shared.Result{Day: "One", Task: "One", Value: total}
}

func PartTwo(input []string) *shared.Result {
	total := 0
	for _, s := range input {
		nums := shared.ExtractNumbersRedux(s)
		total += shared.MergeNumbers(nums)
	}

	return &shared.Result{Day: "One", Task: "Two", Value: total}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
