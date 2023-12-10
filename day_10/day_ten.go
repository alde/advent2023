package ten

import (
	"alde.nu/advent2023/shared"
)

type Coordinates struct {
	X, Y int
}

type Direction Coordinates

var (
	North = Direction{-1, 0}
	South = Direction{1, 0}
	West  = Direction{0, -1}
	East  = Direction{0, 1}
)

var pipes = map[byte]map[Direction]Direction{
	'|': {
		North: North,
		South: South,
	},
	'-': {
		East: East,
		West: West,
	},
	'L': {
		South: East,
		West:  North,
	},
	'J': {
		East:  North,
		South: West,
	},
	'7': {
		East:  South,
		North: West,
	},
	'F': {
		North: East,
		West:  South,
	},
}

func MakeGrid(input []string) [][]byte {
	grid := make([][]byte, len(input))
	for i := range input {
		grid[i] = []byte(input[i])
	}

	return grid
}
func FindStart(grid [][]byte) Coordinates {
	var s Coordinates

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 'S' {
				s.X = x
				s.Y = y
				return s
			}
		}
	}

	panic("no starting position found")
}

func FindLoop(s Coordinates, grid [][]byte) []Coordinates {
	for _, pipe := range "|-LJ7F" {
		grid[s.X][s.Y] = byte(pipe)
		loop := checkLoop(s, grid)
		if loop != nil {
			return loop
		}
	}

	panic("no loop found")
}

func checkLoop(s Coordinates, grid [][]byte) []Coordinates {
	cur := s
	dir := anyKey(pipes[grid[s.X][s.Y]])

	res := []Coordinates{}

	for {
		res = append(res, cur)
		newDir, ok := pipes[grid[cur.X][cur.Y]][dir]
		if !ok {
			return nil
		}

		newCoord := Coordinates{cur.X + newDir.X, cur.Y + newDir.Y}

		if newCoord.X < 0 || newCoord.X >= len(grid) || newCoord.Y < 0 || newCoord.Y >= len(grid[newCoord.X]) {
			return nil
		}
		if newCoord == s {
			if _, ok := pipes[grid[s.X][s.Y]][newDir]; !ok {
				return nil
			}
			break
		}
		cur = newCoord
		dir = newDir
	}

	return res
}

func anyKey(m map[Direction]Direction) Direction {
	for k := range m {
		return k
	}

	panic("empty map")
}

func PartOne(input []string) *shared.Result {
	grid := MakeGrid(input)
	s := FindStart(grid)
	loop := FindLoop(s, grid)

	result := len(loop) / 2
	return &shared.Result{Day: "Ten", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result {
	result := 0
	grid := MakeGrid(input)
	s := FindStart(grid)
	loop := FindLoop(s, grid)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(loop); i++ {
		cur := loop[i]
		next := loop[(i+1)%len(loop)]

		polygonArea += cur.X*next.Y - cur.Y*next.X
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	result = polygonArea - len(loop)/2 + 1

	return &shared.Result{Day: "Ten", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
