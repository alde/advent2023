package six_test

import (
	"fmt"
	"testing"

	"alde.nu/advent2023/six"
	"github.com/stretchr/testify/assert"
)

var INPUT = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func Test_ProcessRaces(t *testing.T) {
	races := six.ProcessRaces(INPUT)

	assert.Equal(t, 7, races[0].RaceLength)
	assert.Equal(t, 9, races[0].RecordDistance)

	assert.Equal(t, 15, races[1].RaceLength)
	assert.Equal(t, 40, races[1].RecordDistance)

	assert.Equal(t, 30, races[2].RaceLength)
	assert.Equal(t, 200, races[2].RecordDistance)
}
func Test_ButtonPress(t *testing.T) {
	testData := []struct {
		pressTime        int
		raceLength       int
		expectedDistance int
	}{
		{pressTime: 0, raceLength: 7, expectedDistance: 0},
		{pressTime: 1, raceLength: 7, expectedDistance: 6},
		{pressTime: 2, raceLength: 7, expectedDistance: 10},
		{pressTime: 3, raceLength: 7, expectedDistance: 12},
		{pressTime: 4, raceLength: 7, expectedDistance: 12},
		{pressTime: 5, raceLength: 7, expectedDistance: 10},
		{pressTime: 6, raceLength: 7, expectedDistance: 6},
		{pressTime: 7, raceLength: 7, expectedDistance: 0},
		{pressTime: 2, raceLength: 15, expectedDistance: 26},
		{pressTime: 3, raceLength: 15, expectedDistance: 36},
		{pressTime: 4, raceLength: 15, expectedDistance: 44},
		{pressTime: 5, raceLength: 15, expectedDistance: 50},
	}

	for _, td := range testData {
		t.Run(fmt.Sprintf("test pressing %d ms in a %d length race", td.pressTime, td.raceLength), func(t *testing.T) {
			actual := six.Launch(td.pressTime, td.raceLength)
			assert.Equal(t, td.expectedDistance, actual)
		})
	}
}

func Test_WaysToWin(t *testing.T) {
	actual := six.WaysToWin(&six.Race{
		RaceLength: 7, RecordDistance: 9,
	})
	assert.Equal(t, 4, len(actual))
}
