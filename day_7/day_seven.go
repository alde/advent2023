package seven

import (
	"sort"
	"strconv"
	"strings"

	"alde.nu/advent2023/shared"
)

type Hand struct {
	Rank     int
	Cards    []string
	Bid      int
	Strength int
}
type Hands []*Hand

var cardPriority = []string{
	"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2",
}

func CompareCards(left string, right string) bool {
	lp := -1
	rp := -1
	// this can be optimized
	for i, c := range cardPriority {
		if c == left {
			lp = i
		}
		if c == right {
			rp = i
		}
	}

	return lp <= rp
}

func (h Hands) Len() int {
	return len(h)
}
func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h Hands) Less(i, j int) bool {
	if h[i].Strength > h[j].Strength {
		// left wins
		return true
	}
	if h[i].Strength < h[j].Strength {
		// right wins
		return false
	}
	// Strength are the same, have to compare cards
	for n := 0; n < 5; n++ { // 5 cards per hand
		if h[i].Cards[n] == h[j].Cards[n] {
			continue // if they are the same, continue to the next card in each hand
		}
		return CompareCards(h[i].Cards[n], h[j].Cards[n])
	}
	return false
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIRS       = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

func Max(seen map[string]int) int {
	max := 0
	for _, v := range seen {
		if v > max {
			max = v
		}
	}
	return max
}

func GetStrength(cards string) int {
	seen := make(map[string]int)
	for i := 0; i < len(cards); i++ {
		r := string(cards[i])
		_, ok := seen[r]
		if !ok {
			seen[r] = 1
		} else {
			seen[r] += 1
		}
	}

	if len(seen) == 1 {
		return FIVE_OF_A_KIND
	}
	if len(seen) == 2 { // four of a kind or full house
		if Max(seen) == 4 {
			return FOUR_OF_A_KIND
		}
		return FULL_HOUSE

	}

	if len(seen) == 3 { // two pairs or three of a kind
		if Max(seen) == 3 {
			return THREE_OF_A_KIND
		}
		return TWO_PAIRS
	}

	if len(seen) == 4 { // only one pair
		return ONE_PAIR
	}

	return HIGH_CARD
}

func ParseHands(input []string) []*Hand {
	hands := []*Hand{}
	for _, s := range input {
		split := strings.Fields(s)
		cards := strings.Split(split[0], "")
		bid, _ := strconv.Atoi(split[1])
		hands = append(hands, &Hand{
			Cards:    cards,
			Bid:      bid,
			Strength: GetStrength(strings.Join(cards, "")),
		})
	}

	sort.Sort(Hands(hands))
	for i := 0; i < len(hands); i++ {
		hands[i].Rank = len(hands) - i
	}

	return hands
}

func PartOne(hands Hands) *shared.Result {
	result := 0
	for _, h := range hands {
		result += h.Rank * h.Bid
	}
	return &shared.Result{Day: "Seven", Task: "One", Value: result}
}

func PartTwo(data []string) *shared.Result {
	result := 0

	return &shared.Result{Day: "Seven", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	hands := ParseHands(data)

	shared.PrintResult(func() *shared.Result { return PartOne(hands) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
