package day_11

import (
	"alde.nu/advent2023/shared"
)

type Coord struct {
	X, Y int
}

func (p *Coord) equals(other *Coord) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Coord) distance(o Coord) int {
	return shared.Abs(p.X-o.X) + shared.Abs(p.Y-o.Y)
}

func ParseInput(input []string) [][]byte {
	res := [][]byte{}
	for _, row := range input {
		r := []byte{}
		for _, col := range row {
			r = append(r, byte(col))
		}
		res = append(res, r)
	}
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			res[row][col] = input[row][col]
		}
	}
	return res
}

func ExpandUniverse(universe [][]byte) ([]int, []int, []Coord) {
	emptyRows := []int{}
	emptyColumns := []int{}
	galaxies := []Coord{}

	for col := 0; col < len(universe[0]); col++ {
		isColEmpty := true
		for row := 0; row < len(universe); row++ {
			h := []byte("#")[0]
			if shared.Contains[int](emptyRows, row) == false &&
				shared.Contains[byte](universe[row], '#') == false {
				emptyRows = append(emptyRows, row)
			} else if universe[row][col] == h {
				isColEmpty = false
				galaxies = append(galaxies, Coord{X: col, Y: row})
			}
		}
		if isColEmpty {
			emptyColumns = append(emptyColumns, col)
		}
	}

	return emptyRows, emptyColumns, galaxies
}

func GalaxyDistances(multiplier int, galaxies []Coord, emptyRows, emptyCols []int) []int {
	distances := []int{}

	for i := 0; i < len(galaxies)-1; i++ {
		from := galaxies[i]

		for j := i; j < len(galaxies); j++ {
			to := galaxies[j]
			d := from.distance(to)
			minY, maxY := shared.MinMax(from.Y, to.Y)
			minX, maxX := shared.MinMax(from.X, to.X)

			for _, r := range emptyRows {
				if r > minY && r < maxY {
					d += multiplier - 1
				}
			}
			for _, c := range emptyCols {
				if c > minX && c < maxX {
					d += multiplier - 1
				}
			}
			distances = append(distances, d)
		}
	}

	return distances
}

func Solve(universe [][]byte, multiplier int) []int {
	emptyRows, emptyColumns, galaxies := ExpandUniverse(universe)
	return GalaxyDistances(multiplier, galaxies, emptyRows, emptyColumns)
}

func PartOne(universe [][]byte, multiplier int) *shared.Result {
	result := 0
	distances := Solve(universe, multiplier)

	for _, r := range distances {
		result += r
	}

	return &shared.Result{Day: "Eleven", Task: "One", Value: result}
}

func PartTwo(universe [][]byte, multiplier int) *shared.Result {
	result := 0
	distances := Solve(universe, multiplier)

	for _, r := range distances {
		result += r
	}

	return &shared.Result{Day: "Eleven", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)
	universe := ParseInput(data)

	shared.PrintResult(func() *shared.Result { return PartOne(universe, 2) })
	shared.PrintResult(func() *shared.Result { return PartTwo(universe, 1_000_000) })
}
