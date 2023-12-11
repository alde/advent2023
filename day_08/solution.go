package day_08

import (
	"fmt"
	"strings"

	"alde.nu/advent2023/shared"
)

func GetDirections(dirs string) *shared.CircleList[rune] {
	cl := shared.NewCircleList[rune]()
	for _, r := range dirs {
		cl.AddNode(r)
	}

	return cl
}

type Direction struct {
	Left  string
	Right string
}

func GetNodes(nodeList []string) map[string]*Direction {
	nodes := make(map[string]*Direction)
	for _, row := range nodeList {
		if len(row) == 0 {
			continue
		}

		split1 := strings.Split(row, " = ")
		label := split1[0]
		split2 := strings.Split(split1[1], ", ")
		left := strings.TrimPrefix(split2[0], "(")
		right := strings.TrimSuffix(split2[1], ")")
		nodes[label] = &Direction{
			Left: left, Right: right,
		}
	}
	return nodes
}

func contains(es []string, node string) bool {
	for _, e := range es {
		if e == node {
			return true
		}
	}
	return false
}

func Traverse(directions *shared.CircleList[rune], nodes map[string]*Direction) int {
	return traverseFrom("AAA", directions, nodes)[0]
}

func traverseFrom(current string, directions *shared.CircleList[rune], nodes map[string]*Direction) []int {
	count := 0
	seen := []string{}
	moves := []int{}
	for {
		node, ok := nodes[current]
		if !ok {
			panic(fmt.Errorf("unknown node %s", current))
		}
		dir := directions.Pop()
		if dir == 'L' {
			current = node.Left
		} else {
			current = node.Right
		}
		count += 1
		if strings.HasSuffix(current, "Z") {
			seen = append(seen, current)
			moves = append(moves, count)
		}
		if contains(seen, current) {
			// we've looped
			return moves
		}
	}
}

func GreatestCommonDivider(a, b int) int {
	for b > 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
func LeastCommonMultiple(nums []int) int {
	innerLCM := func(a, b int) int {
		return a * (b / GreatestCommonDivider(a, b))
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = innerLCM(result, nums[i])
	}
	return result
}

func TraverseMultiple(currents []string, directions *shared.CircleList[rune], nodes map[string]*Direction) int {
	endStates := []int{}
	for _, startPos := range currents {
		dirs := directions
		moves := traverseFrom(startPos, dirs, nodes)
		endStates = append(endStates, moves...)
	}

	return LeastCommonMultiple(endStates)
}

func GetStartingPositions(nodes map[string]*Direction) []string {
	res := []string{}
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			res = append(res, k)
		}
	}

	return res
}

func PartOne(directions *shared.CircleList[rune], nodes map[string]*Direction) *shared.Result {
	result := Traverse(directions, nodes)
	return &shared.Result{Day: "eight", Task: "One", Value: result}
}

func PartTwo(starts []string, directions *shared.CircleList[rune], nodes map[string]*Direction) *shared.Result {
	result := TraverseMultiple(starts, directions, nodes)
	return &shared.Result{Day: "eight", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)
	directions := GetDirections(data[0])
	nodes := GetNodes(data[1:])

	shared.PrintResult(func() *shared.Result { return PartOne(directions, nodes) })
	starts := GetStartingPositions(nodes)
	shared.PrintResult(func() *shared.Result { return PartTwo(starts, directions, nodes) })
}
