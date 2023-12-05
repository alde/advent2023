package five

import (
	"strings"

	"alde.nu/advent2023/shared"
)

type Range struct {
	sourceRangeStart int
	targetRangeStart int
	rangeLength      int
}

func (r *Range) Contains(source int) bool {
	return source >= r.sourceRangeStart && source < r.sourceRangeStart+r.rangeLength
}

func (r *Range) Convert(source int) int {
	if source >= r.sourceRangeStart && source < r.sourceRangeStart+r.rangeLength {
		return r.targetRangeStart + (source - r.sourceRangeStart)
	}
	return source
}

type Map struct {
	ranges []*Range
}

func (m *Map) AddRange(sourceStart int, targetStart int, length int) {
	m.ranges = append(m.ranges, &Range{
		sourceRangeStart: sourceStart,
		targetRangeStart: targetStart,
		rangeLength:      length,
	})
}

func (m *Map) Get(source int) int {
	for _, r := range m.ranges {
		if r.Contains(source) {
			return r.Convert(source)
		}
	}
	return source
}

type Almanac struct {
	SeedList              []int
	SeedToSoil            *Map
	SoilToFertilizer      *Map
	FertilizerToWater     *Map
	WaterToLight          *Map
	LightToTemperature    *Map
	TemperatureToHumidity *Map
	HumidityToLocation    *Map
}

func (a *Almanac) ResolveMappings(seed int) []int {
	soil := a.SeedToSoil.Get(seed)
	fertilizer := a.SoilToFertilizer.Get(soil)
	water := a.FertilizerToWater.Get(fertilizer)
	light := a.WaterToLight.Get(water)
	temp := a.LightToTemperature.Get(light)
	hum := a.TemperatureToHumidity.Get(temp)
	loc := a.HumidityToLocation.Get(hum)

	return []int{
		seed, soil, fertilizer, water, light, temp, hum, loc,
	}
}

func findEndOfBlock(slice []string, startIndex int) int {
	for i := startIndex; i < len(slice); i++ {
		if strings.TrimSpace(slice[i]) == "" {
			return i
		}
	}
	return len(slice)
}

func ProcessInput(input []string) Almanac {
	almanac := Almanac{}
	almanac.SeedList = shared.ConvertToNumSlice(strings.TrimLeft(input[0], "seeds: "))
	for n := 1; n < len(input); n++ {
		endIndex := n
		section := ""
		if strings.HasSuffix(input[n], " map:") {
			endIndex = findEndOfBlock(input, n)
			section = strings.TrimRight(input[n], " map:")

			m := MakeMap(input[n+1 : endIndex])
			switch section {
			case "seed-to-soil":
				almanac.SeedToSoil = m
			case "soil-to-fertilizer":
				almanac.SoilToFertilizer = m
			case "fertilizer-to-water":
				almanac.FertilizerToWater = m
			case "water-to-light":
				almanac.WaterToLight = m
			case "light-to-temperature":
				almanac.LightToTemperature = m
			case "temperature-to-humidity":
				almanac.TemperatureToHumidity = m
			case "humidity-to-location":
				almanac.HumidityToLocation = m
			}
		}
		n = endIndex
	}
	return almanac
}

func MakeMap(input []string) *Map {
	result := &Map{}

	for _, row := range input {
		parts := shared.ConvertToNumSlice(row)
		targetStart := parts[0]
		sourceStart := parts[1]
		length := parts[2]
		result.AddRange(sourceStart, targetStart, length)
	}

	return result
}

func PartOne(input []string) *shared.Result[int] {
	result := 0
	almanac := ProcessInput(input)
	for _, seed := range almanac.SeedList {
		path := almanac.ResolveMappings(seed)
		location := path[len(path)-1]
		if result == 0 || location < result {
			result = location
		}
	}
	return &shared.Result[int]{Day: "Five", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result[int] {
	result := 0

	return &shared.Result[int]{Day: "Five", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInput(input)

	shared.PrintResult(PartOne(strings.Split(data, "\n")))
	// shared.PrintResult(PartTwo(strings.Split(data, "\n")))

}
