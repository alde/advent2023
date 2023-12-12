package shared

import (
	"strconv"
	"strings"
	"unicode"
)

func ConvertToNumSlice(s string) []int {
	res := []int{}
	for _, wn := range strings.Fields(s) {
		wn1, err := strconv.Atoi(wn)
		if err != nil {
			continue
		}
		res = append(res, wn1)
	}
	return res
}
func DropWhitespaces(s string) int {
	storage := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		storage += string(s[i])
	}

	num, _ := strconv.Atoi(storage)
	return num
}
func ExtractNumbers(s string) []int {
	var res = []int{}
	for _, r := range s {
		if unicode.IsNumber(r) {
			res = append(res, int(r-'0'))
		}
	}
	if len(res) == 1 {
		res = append(res, res[0])
	}
	if len(res) > 2 {
		return []int{res[0], res[len(res)-1]}
	}
	return res
}

var WordNums = []struct {
	word  string
	value int
}{
	{word: "one", value: 1},
	{word: "two", value: 2},
	{word: "three", value: 3},
	{word: "four", value: 4},
	{word: "five", value: 5},
	{word: "six", value: 6},
	{word: "seven", value: 7},
	{word: "eight", value: 8},
	{word: "nine", value: 9},
}

func ParseNumbers(str string, idx int) int {
	for _, wn := range WordNums {
		if len(str) >= idx+len(wn.word) && str[idx:idx+len(wn.word)] == wn.word {
			return wn.value
		}
	}
	return -1
}

func ExtractNumbersRedux(s string) []int {
	var res = []int{}

	for i, r := range s {
		if unicode.IsNumber(r) {
			res = append(res, int(r-'0'))
		}
		c := ParseNumbers(s, i)
		if c > 0 {
			res = append(res, c)
		}
	}
	if len(res) == 1 {
		res = append(res, res[0])
	}
	if len(res) > 2 {
		return []int{res[0], res[len(res)-1]}
	}

	return res
}

func MergeNumbers(ints []int) int {
	res := ""

	for _, i := range ints {
		res += strconv.Itoa(i)
	}

	response, _ := strconv.Atoi(res)
	return response
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
