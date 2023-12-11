package day_01_test

import (
	"strings"
	"testing"

	"alde.nu/advent2023/day_01"
	"github.com/stretchr/testify/assert"
)

const INPUT = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

func Test_PartOne(t *testing.T) {

	res := day_01.PartOne(strings.Split(INPUT, "\n"))

	assert.Equal(t, 142, res.Value)
}
