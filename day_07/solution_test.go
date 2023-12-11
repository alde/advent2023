package day_07_test

import (
	"fmt"
	"testing"

	"alde.nu/advent2023/day_07"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func Test_GetStrength_Normal(t *testing.T) {
	context := day_07.Context{
		Priority:     []string{},
		JacksAreWild: false,
	}
	assert.Equal(t, day_07.ONE_PAIR, day_07.GetStrength("32T3K", &context))
	assert.Equal(t, day_07.THREE_OF_A_KIND, day_07.GetStrength("T55J5", &context))
	assert.Equal(t, day_07.TWO_PAIRS, day_07.GetStrength("KK677", &context))
	assert.Equal(t, day_07.TWO_PAIRS, day_07.GetStrength("KTJJT", &context))
	assert.Equal(t, day_07.THREE_OF_A_KIND, day_07.GetStrength("QQQJA", &context))
}

func Test_GetStrengthForWildJacks(t *testing.T) {
	context := &day_07.Context{
		Priority:     []string{},
		JacksAreWild: true,
	}
	testData := []struct {
		strength int
		hand     string
	}{
		{day_07.HIGH_CARD, "12345"},
		{day_07.HIGH_CARD, "AKQT9"},
		{day_07.ONE_PAIR, "12344"},
		{day_07.ONE_PAIR, "AKQJT"},
		{day_07.TWO_PAIRS, "AAKKQ"},
		{day_07.THREE_OF_A_KIND, "AAKQJ"},
		{day_07.THREE_OF_A_KIND, "AKQJJ"},
		{day_07.FULL_HOUSE, "AAKKJ"},
		{day_07.FULL_HOUSE, "AAKKK"},
		{day_07.FOUR_OF_A_KIND, "AAKJJ"},
		{day_07.FOUR_OF_A_KIND, "AAAAK"},
		{day_07.FIVE_OF_A_KIND, "AAAAA"},
		{day_07.FIVE_OF_A_KIND, "AJJJJ"},
		{day_07.FIVE_OF_A_KIND, "AAJJJ"},
		{day_07.FIVE_OF_A_KIND, "AAAJJ"},
		{day_07.FIVE_OF_A_KIND, "AAAAJ"},
		{day_07.FIVE_OF_A_KIND, "AAAAA"},
	}

	for i, td := range testData {
		t.Run(fmt.Sprintf("Test %d - %s should be %s", i, td.hand, day_07.StrengthName(td.strength)), func(t *testing.T) {
			actual := day_07.GetStrength(td.hand, context)
			assert.Equal(t, td.strength, actual)
		})
	}
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
	parser := &day_07.Context{
		JacksAreWild: false,
		Priority:     day_07.DefultPriority,
	}
	hands := day_07.ParseHands(INPUT, parser)

	for i, e := range expected {
		assert.Equal(t, e.rank, hands[i].Rank)
		assert.Equal(t, e.cards, hands[i].Cards)
		assert.Equal(t, e.bid, hands[i].Bid)
	}
}

func Test_PartTwo(t *testing.T) {
	expected := 5905
	parser := &day_07.Context{
		JacksAreWild: true,
		Priority:     day_07.CardPrioJackWild,
	}

	hands := day_07.ParseHands(INPUT, parser)

	result := 0
	for _, h := range hands {
		result += h.Rank * h.Bid
	}
	assert.Equal(t, expected, result)
}
