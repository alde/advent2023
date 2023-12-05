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

type SeedRange struct {
	Start  int
	Length int
}

type Almanac struct {
	SeedList      []int
	SeedListRange []SeedRange
	Dict          map[string]*Map
}

func (a *Almanac) ResolveMappings(seed int) []int {

	soil := a.Dict["seed-to-soil"].Get(seed)
	fertilizer := a.Dict["soil-to-fertilizer"].Get(soil)
	water := a.Dict["fertilizer-to-water"].Get(fertilizer)
	light := a.Dict["water-to-light"].Get(water)
	temperature := a.Dict["light-to-temperature"].Get(light)
	humidity := a.Dict["temperature-to-humidity"].Get(temperature)
	location := a.Dict["humidity-to-location"].Get(humidity)

	return []int{
		seed, soil, fertilizer, water, light, temperature, humidity, location,
	}
}

func (a *Almanac) HasRangeContainingSeed(seed int) bool {
	for _, r := range a.SeedListRange {
		if seed >= r.Start && seed <= r.Start+r.Length {
			return true
		}
	}
	return false
}
func (a *Almanac) Reverse(location int) int {
	humidity := a.Dict["humidity-to-location"].Reverse(location)
	temp := a.Dict["temperature-to-humidity"].Reverse(humidity)
	light := a.Dict["light-to-temperature"].Reverse(temp)
	water := a.Dict["water-to-light"].Reverse(light)
	fertilizer := a.Dict["fertilizer-to-water"].Reverse(water)
	soil := a.Dict["soil-to-fertilizer"].Reverse(fertilizer)
	seed := a.Dict["seed-to-soil"].Reverse(soil)

	return seed
}

func findEndOfBlock(slice []string, startIndex int) int {
	for i := startIndex; i < len(slice); i++ {
		if strings.TrimSpace(slice[i]) == "" {
			return i
		}
	}
	return len(slice)
}

func makeMaps(input []string) map[string]*Map {
	result := make(map[string]*Map)

	for n := 1; n < len(input); n++ {
		endIndex := n
		section := ""
		if strings.HasSuffix(input[n], " map:") {
			endIndex = findEndOfBlock(input, n)
			section = strings.TrimRight(input[n], " map:")
			m := MakeMap(input[n+1 : endIndex])
			result[section] = m
		}
		n = endIndex
	}
	return result
}

func ProcessInput(input []string) *Almanac {
	almanac := &Almanac{}
	seedList := shared.ConvertToNumSlice(strings.TrimLeft(input[0], "seeds: "))
	almanac.SeedList = seedList
	seedRanges := []SeedRange{}
	for i := 0; i < len(seedList); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			Start:  seedList[i],
			Length: seedList[i+1],
		})
	}
	almanac.SeedListRange = seedRanges

	almanac.Dict = makeMaps(input)

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

func PartOne(almanac *Almanac) *shared.Result {
	result := 0

	for _, seed := range almanac.SeedList {
		path := almanac.ResolveMappings(seed)
		location := path[len(path)-1]
		if result == 0 || location < result {
			result = location
		}
	}
	return &shared.Result{Day: "Five", Task: "One", Value: result}
}

func PartTwo(almanac *Almanac) *shared.Result {
	result := math.MaxInt

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
	almanac := ProcessInput(data)

	shared.PrintResult(func() *shared.Result { return PartOne(almanac) })
	shared.PrintResult(func() *shared.Result { return PartTwo(almanac) })
}
