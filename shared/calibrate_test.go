package shared_test

import (
	"testing"

	"alde.nu/advent2023/shared"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	input    string
	expected []int
}

func Test_ExtractNumbers(t *testing.T) {
	tdata := []TestData{{
		input:    "abc123def",
		expected: []int{1, 3},
	}, {
		input:    "two5six",
		expected: []int{5, 5},
	},
	}

	for _, td := range tdata {
		actual := shared.ExtractNumbers(td.input)
		assert.Equal(t, actual, td.expected)
	}
}

func Test_ExtractNumbersRedux(t *testing.T) {
	tdata := []TestData{{
		input:    "abc123def",
		expected: []int{1, 3},
	}, {
		input:    "two5six",
		expected: []int{2, 6},
	},
	}

	for _, td := range tdata {
		actual := shared.ExtractNumbersRedux(td.input)
		assert.Equal(t, actual, td.expected)
	}
}

func Test_ParseNumbers(t *testing.T) {
	actual := shared.ParseNumbers("ione123def", 0)
	assert.Equal(t, actual, -1)

	actual = shared.ParseNumbers("two5six", 0)
	assert.Equal(t, actual, 2)

	actual = shared.ParseNumbers("two5six", 4)
	assert.Equal(t, actual, 6)
}
