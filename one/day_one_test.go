package one_test

import (
	"strings"
	"testing"

	"alde.nu/advent2023/one"
	"github.com/stretchr/testify/assert"
)

const INPUT = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

func Test_PartOne(t *testing.T) {

	res := one.PartOne(strings.Split(INPUT, "\n"))

	assert.Equal(t, 142, res.Value)
}
