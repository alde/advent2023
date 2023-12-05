package five

import (
	"math"
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
	if r.Contains(source) {
		return r.targetRangeStart + (source - r.sourceRangeStart)
	}
	return source
}

func (r *Range) ContainsTarget(target int) bool {
	return target >= r.targetRangeStart && target < r.targetRangeStart+r.rangeLength
}

func (r *Range) GetSourceForTarget(target int) int {
	if r.ContainsTarget(target) {
		return r.sourceRangeStart + (target - r.targetRangeStart)
	}
	return target
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

func (m *Map) Reverse(from int) int {
	for _, r := range m.ranges {
		if r.ContainsTarget(from) {
			return r.GetSourceForTarget(from)
		}
	}
	return from
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

type SeedRange struct {
	Start  int
	Length int
}

type AlmanacTwo struct {
	SeedListRange         []SeedRange
	SeedToSoil            *Map
	SoilToFertilizer      *Map
	FertilizerToWater     *Map
	WaterToLight          *Map
	LightToTemperature    *Map
	TemperatureToHumidity *Map
	HumidityToLocation    *Map
}

func (a *AlmanacTwo) HasRangeContainingSeed(seed int) bool {
	for _, r := range a.SeedListRange {
		if seed >= r.Start && seed <= r.Start+r.Length {
			return true
		}
	}
	return false
}
func (a *AlmanacTwo) Reverse(location int) int {
	humidity := a.HumidityToLocation.Reverse(location)
	temp := a.TemperatureToHumidity.Reverse(humidity)
	light := a.LightToTemperature.Reverse(temp)
	water := a.WaterToLight.Reverse(light)
	fertilizer := a.FertilizerToWater.Reverse(water)
	soil := a.SoilToFertilizer.Reverse(fertilizer)
	seed := a.SeedToSoil.Reverse(soil)

	return seed
}
func (a *AlmanacTwo) ResolveMappings(seed int) int {
	soil := a.SeedToSoil.Get(seed)
	fertilizer := a.SoilToFertilizer.Get(soil)
	water := a.FertilizerToWater.Get(fertilizer)
	light := a.WaterToLight.Get(water)
	temp := a.LightToTemperature.Get(light)
	hum := a.TemperatureToHumidity.Get(temp)
	loc := a.HumidityToLocation.Get(hum)

	return loc
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

func ProcessInputPartTwo(input []string) AlmanacTwo {
	almanac := AlmanacTwo{}
	seedList := shared.ConvertToNumSlice(strings.TrimLeft(input[0], "seeds: "))
	seedRanges := []SeedRange{}
	for i := 0; i < len(seedList); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			Start:  seedList[i],
			Length: seedList[i+1],
		})
	}
	almanac.SeedListRange = seedRanges

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

func PartOne(input []string) *shared.Result {
	result := 0
	almanac := ProcessInput(input)
	for _, seed := range almanac.SeedList {
		path := almanac.ResolveMappings(seed)
		location := path[len(path)-1]
		if result == 0 || location < result {
			result = location
		}
	}
	return &shared.Result{Day: "Five", Task: "One", Value: result}
}

func PartTwo(input []string) *shared.Result {
	result := math.MaxInt
	almanac := ProcessInputPartTwo(input)

	for loc := 0; loc < 200_000_000; loc++ {
		seed := almanac.Reverse(loc)
		if almanac.HasRangeContainingSeed(seed) {
			result = loc
			break
		}
	}

	return &shared.Result{Day: "Five", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
