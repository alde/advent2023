package six

import (
	"math"

	"alde.nu/advent2023/shared"
)

type Race struct {
	RaceLength     int
	RecordDistance int
}

func WaysToWin(race *Race) []int {
	waysToWin := []int{}
	// quarterRaceLength := race.RaceLength / 4
	// fmt.Printf("check from %d to %d", quarterRaceLength, quarterRaceLength*3)
	for pressTime := 0; pressTime < race.RaceLength; pressTime++ {
		if distance := Launch(pressTime, race.RaceLength); distance > race.RecordDistance {
			waysToWin = append(waysToWin, distance)
		}
	}

	return waysToWin
}
func ProcessRaces(input []string) []*Race {
	races := []*Race{}
	times := shared.ConvertToNumSlice(input[0])
	distances := shared.ConvertToNumSlice(input[1])
	for i := 0; i < len(times); i++ {
		races = append(races, &Race{
			RaceLength:     times[i],
			RecordDistance: distances[i],
		})
	}
	return races
}

func Launch(pressTime int, raceLength int) int {
	velocity := 0
	if pressTime == 0 {
		return 0
	}
	velocity += pressTime
	raceLength -= pressTime

	return velocity * raceLength
}

func PartOne(races []*Race) *shared.Result {
	result := 1

	for _, r := range races {
		ways := WaysToWin(r)
		result *= len(ways)
	}

	return &shared.Result{Day: "Six", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result {
	result := math.MaxInt

	return &shared.Result{Day: "Six", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)
	races := ProcessRaces(data)

	shared.PrintResult(func() *shared.Result { return PartOne(races) })
	// shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
