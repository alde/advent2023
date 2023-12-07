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
	priority []string // wish I could put this somewhere else, but I don't want to make my own sorting
}
type Hands []*Hand

type Context struct {
	Priority     []string
	JacksAreWild bool
}

var DefultPriority = []string{
	"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2",
}
var CardPrioJackWild = []string{
	"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J",
}

func CompareCards(left string, right string, priority []string) bool {
	lp := -1
	rp := -1
	// this can be optimized
	for i, c := range priority {
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
		return CompareCards(h[i].Cards[n], h[j].Cards[n], h[i].priority)
	}
	return false
}

func Max(seen map[string]int) int {
	max := 0
	for _, v := range seen {
		if v > max {
			max = v
		}
	}
	return max
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

func StrengthName(i int) string {
	switch i {
	case HIGH_CARD:
		return "HIGH_CARD"
	case ONE_PAIR:
		return "ONE_PAIR"
	case TWO_PAIRS:
		return "TWO_PAIRS"
	case THREE_OF_A_KIND:
		return "THREE_OF_A_KIND"
	case FULL_HOUSE:
		return "FULL_HOUSE"
	case FOUR_OF_A_KIND:
		return "FOUR_OF_A_KIND"
	case FIVE_OF_A_KIND:
		return "FIVE_OF_A_KIND"
	}
	return "unknown"
}

func CountCards(cards string) map[string]int {
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

	return seen
}

func GetStrengthForWildJacks(seen map[string]int) (str int) {
	jacks, ok := seen["J"]
	if !ok {
		jacks = 0
	}

	seenLength := len(seen)
	if seenLength == 1 {
		return FIVE_OF_A_KIND
	}

	if seenLength == 2 {
		if jacks > 0 {
			return FIVE_OF_A_KIND
		}
		if Max(seen) == 3 {
			return FULL_HOUSE
		}

		return FOUR_OF_A_KIND
	}

	if seenLength == 3 {
		if jacks == 1 {
			if Max(seen) == 3 {
				return FOUR_OF_A_KIND // 3 of something, 1 of something else, 1 jack, eg AAAJT
			}
			return FULL_HOUSE // 2 of something, 1 jack, 2 of something else, eg AAJTT
		}
		if jacks == 2 || jacks == 3 {
			// 2 of something, 1 of something else, 2 jacks
			// 3 jacks, 1 of something, 1 of something else
			return FOUR_OF_A_KIND
		}

		if Max(seen) == 3 {
			// it's not 1 or 2 jacks, so 3 of something, 2 of something else but different and no jacks
			return THREE_OF_A_KIND
		}
		return TWO_PAIRS // it's not one or two jacks (as per above)
	}

	if seenLength == 4 {
		if jacks > 0 {
			return THREE_OF_A_KIND // 2 of something and one jack
		}
		return ONE_PAIR // has to be a pair of jacks or no jacks
	}
	if jacks > 0 {
		return ONE_PAIR // One jack means it has to be a pair
	}
	return HIGH_CARD
}

func GetNormalStrength(seen map[string]int) int {
	seenLength := len(seen)
	if seenLength == 1 {
		return FIVE_OF_A_KIND
	}
	if seenLength == 2 { // four of a kind or full house
		if Max(seen) == 4 {
			return FOUR_OF_A_KIND
		}
		return FULL_HOUSE

	}

	if seenLength == 3 { // two pairs or three of a kind
		if Max(seen) == 3 {
			return THREE_OF_A_KIND
		}
		return TWO_PAIRS
	}

	if seenLength == 4 { // only one pair
		return ONE_PAIR
	}

	return HIGH_CARD
}
func GetStrength(cards string, context *Context) int {
	seen := CountCards(cards)
	if context.JacksAreWild {
		return GetStrengthForWildJacks(seen)
	}
	return GetNormalStrength(seen)

}

func ParseHands(input []string, context *Context) []*Hand {
	var hands Hands
	for _, s := range input {
		split := strings.Fields(s)
		cards := strings.Split(split[0], "")
		bid, _ := strconv.Atoi(split[1])
		hands = append(hands, &Hand{
			Cards:    cards,
			Bid:      bid,
			Strength: GetStrength(strings.Join(cards, ""), context),
			priority: context.Priority,
		})
	}

	sort.Sort(hands)
	for i := 0; i < len(hands); i++ {
		hands[i].Rank = len(hands) - i
	}

	return hands
}

func PartOne(data []string) *shared.Result {
	context := &Context{
		Priority:     DefultPriority,
		JacksAreWild: false,
	}
	hands := ParseHands(data, context)
	result := 0

	for _, h := range hands {
		result += h.Rank * h.Bid
	}
	return &shared.Result{Day: "Seven", Task: "One", Value: result}
}

func PartTwo(data []string) *shared.Result {
	context := &Context{
		Priority:     CardPrioJackWild,
		JacksAreWild: true,
	}

	hands := ParseHands(data, context)

	result := 0
	for _, h := range hands {
		// fmt.Printf("hand %+v\n", h)
		result += h.Rank * h.Bid
	}

	return &shared.Result{Day: "Seven", Task: "Two", Value: result}
}

func Run(input string) {
	data := shared.LoadInputAsStringSlice(input)

	shared.PrintResult(func() *shared.Result { return PartOne(data) })
	shared.PrintResult(func() *shared.Result { return PartTwo(data) })
}
