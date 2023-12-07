package seven_test

import (
	"testing"

	seven "alde.nu/advent2023/day_7"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func Test_GetStrength(t *testing.T) {
	assert.Equal(t, seven.ONE_PAIR, seven.GetStrength("32T3K"))
	assert.Equal(t, seven.THREE_OF_A_KIND, seven.GetStrength("T55J5"))
	assert.Equal(t, seven.TWO_PAIRS, seven.GetStrength("KK677"))
	assert.Equal(t, seven.TWO_PAIRS, seven.GetStrength("KTJJT"))
	assert.Equal(t, seven.THREE_OF_A_KIND, seven.GetStrength("QQQJA"))
}

func Test_ParseHands(t *testing.T) {
	expected := []struct {
		rank  int
		cards []string
		bid   int
	}{
		{5, []string{"Q", "Q", "Q", "J", "A"}, 483},
		{4, []string{"T", "5", "5", "J", "5"}, 684},
		{3, []string{"K", "K", "6", "7", "7"}, 28},
		{2, []string{"K", "T", "J", "J", "T"}, 220},
		{1, []string{"3", "2", "T", "3", "K"}, 765},
	}
	hands := seven.ParseHands(INPUT)

	for i, e := range expected {
		assert.Equal(t, e.rank, hands[i].Rank)
		assert.Equal(t, e.cards, hands[i].Cards)
		assert.Equal(t, e.bid, hands[i].Bid)
	}
}
