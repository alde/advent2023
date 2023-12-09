package nine

import (
	"alde.nu/advent2023/shared"
)

func allZeros(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func CaluclateDifferences(nums []int) []int {
	res := []int{}
	for n := 0; n < len(nums)-1; n++ {
		res = append(res, nums[n+1]-nums[n])
	}
	return res
}

func BuildHistory(nums []int) [][]int {
	res := [][]int{nums}
	for {
		nums = CaluclateDifferences(nums)
		res = append(res, nums)
		if allZeros(nums) {
			break
		}
		if len(res) > 100_000 {
			panic("infinite loop!?")
		}
	}
	return res
}

func Extrapolate(history [][]int) [][]int {
	for n := len(history) - 1; n >= 0; n-- {
		entry := history[n]
		newValue := 0
		if n < len(history)-1 {
			prevRow := history[n+1]
			newValue = entry[len(entry)-1] + prevRow[len(prevRow)-1]
		}

		entry = append(entry, newValue)
		history[n] = entry
	}
	return history
}

func PartOne(data []string) *shared.Result {
	result := 0

	for _, row := range data {
		nums := shared.ConvertToNumSlice(row)
		history := BuildHistory(nums)
		future := Extrapolate(history)[0]

		result += future[len(future)-1]
	}

	return &shared.Result{Day: "nine", Task: "One", Value: result}
}

func PartTwo(data []string) *shared.Result {
	result := 0
	return &shared.Result{Day: "nine", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
