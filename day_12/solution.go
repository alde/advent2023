package day_12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"alde.nu/advent2023/shared"
)

type GroupArrangements struct {
	record            string
	possiblePositions [][]int
	sizes             []int
	cache             map[string]int
}

// Part 2 redoes a lot of calculations - so keep a cache
func (ga GroupArrangements) memoizedCalculation(nesting int, start int) int {
	key := fmt.Sprintf("n%ds%d", nesting, start)
	if v, ok := ga.cache[key]; ok {
		return v
	}
	result := ga.calculation(nesting, start)
	ga.cache[key] = result
	return result
}

func (ga GroupArrangements) calculation(nesting int, start int) int {
	if len(ga.possiblePositions[nesting:]) == 1 {
		possibleCount := 0
		for _, pos := range ga.possiblePositions[nesting:][0] {
			if pos >= start && !strings.ContainsRune(ga.record[pos+ga.sizes[nesting:][0]:], '#') && !strings.ContainsRune(ga.record[start:pos], '#') {
				possibleCount++
			}
		}
		return possibleCount
	}
	possibleCount := 0
	for _, pos := range ga.possiblePositions[nesting:][0] {
		if pos >= start && !strings.ContainsRune(ga.record[start:pos], '#') {
			possibleCount += ga.memoizedCalculation(nesting+1, ga.sizes[nesting:][0]+pos+1)
		}
	}
	return possibleCount
}

func Parse(input string) (string, []int) {
	split := strings.Fields(input)

	groups := []int{}
	for _, v := range strings.Split(split[1], ",") {
		value, _ := strconv.Atoi(v)
		groups = append(groups, value)
	}
	return split[0], groups
}

func GetCandidates(record string, groupSizes []int) [][]int {
	candidates := [][]int{}
	for _, size := range groupSizes {
		regx, _ := regexp.Compile(fmt.Sprintf("([?.][#?]{%d}[?.])", size))
		padded := fmt.Sprintf(".%s.", record)
		thisGroupsCandidates := []int{}
		offset := 0
		for {
			pos := regx.FindStringIndex(padded[offset:])
			if pos == nil {
				break
			}
			thisGroupsCandidates = append(thisGroupsCandidates, pos[0]+offset)
			offset += pos[0] + 1
		}
		if len(thisGroupsCandidates) == 0 {
			panic(fmt.Errorf("no potisions available | %v %v %v", record, groupSizes, size))
		}
		candidates = append(candidates, thisGroupsCandidates)
	}
	return candidates
}

func ValidArrangements(row string, unfoldMap bool) int {
	record, groupSizes := Parse(row)

	if unfoldMap {
		// replace the record with 5 copies of itself - separated by ?
		record = fmt.Sprintf("%s?%s?%s?%s?%s", record, record, record, record, record)
		expanded := []int{}
		// Add 5 more copies of the groups
		for i := 0; i < 5; i++ {
			expanded = append(expanded, groupSizes...)
		}
		groupSizes = expanded
	}

	possiblePositions := GetCandidates(record, groupSizes)
	calculator := GroupArrangements{
		record:            record,
		possiblePositions: possiblePositions,
		sizes:             groupSizes,
		cache:             make(map[string]int),
	}

	total := calculator.memoizedCalculation(0, 0)
	return total
}

func PartOne(data []string) *shared.Result {
	result := 0

	for _, row := range data {
		arrangements := ValidArrangements(row, false)
		result += arrangements
	}

	return &shared.Result{Day: "Twelve", Task: "One", Value: result}
}

func PartTwo(data []string) *shared.Result {
	result := 0

	for _, row := range data {
		arrangements := ValidArrangements(row, true)
		result += arrangements
	}

	return &shared.Result{Day: "Twelve", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
