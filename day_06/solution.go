package day_06

import (
	"strings"

	"alde.nu/advent2023/shared"
)

type Race struct {
	RaceLength     int
	RecordDistance int
}

func WaysToWin(race *Race) int {
	waysToWin := 0
	for pressTime := 0; pressTime < race.RaceLength; pressTime++ {
		if Launch(pressTime, race.RaceLength) > race.RecordDistance {
			waysToWin += 1
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

func ProcessAsSingleRace(input []string) *Race {
	time := shared.DropWhitespaces(strings.Split(input[0], ":")[1])
	distance := shared.DropWhitespaces(strings.Split(input[1], ":")[1])
	return &Race{
		RaceLength:     time,
		RecordDistance: distance,
	}
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
		result *= WaysToWin(r)
	}

	return &shared.Result{Day: "Six", Task: "One", Value: result}
}

func PartTwo(race *Race) *shared.Result {
	result := WaysToWin(race)

	return &shared.Result{Day: "Six", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)
	races := ProcessRaces(data)

	shared.PrintResult(func() *shared.Result { return PartOne(races) })
	singleRace := ProcessAsSingleRace(data)
	shared.PrintResult(func() *shared.Result { return PartTwo(singleRace) })
}
