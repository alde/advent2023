package main

import (
	"flag"

	one "alde.nu/advent2023/day_1"
	two "alde.nu/advent2023/day_2"
	three "alde.nu/advent2023/day_3"
	four "alde.nu/advent2023/day_4"
	five "alde.nu/advent2023/day_5"
	six "alde.nu/advent2023/day_6"
	seven "alde.nu/advent2023/day_7"
	"alde.nu/advent2023/shared"
)

func main() {

	target := flag.String("target", "all", "which day to run")
	flag.Parse()

	if *target == "all" || *target == "one" {
		shared.PrintDay("one", func() { one.Run("inputs/ONE.txt") })
	}
	if *target == "all" || *target == "two" {
		shared.PrintDay("two", func() { two.Run("inputs/TWO.txt") })
	}
	if *target == "all" || *target == "three" {
		shared.PrintDay("three", func() { three.Run("inputs/THREE.txt") })
	}
	if *target == "all" || *target == "four" {
		shared.PrintDay("four", func() { four.Run("inputs/FOUR.txt") })
	}
	if *target == "all" || *target == "five" {
		shared.PrintDay("five", func() { five.Run("inputs/FIVE.txt") })
	}
	if *target == "all" || *target == "six" {
		shared.PrintDay("six", func() { six.Run("inputs/SIX.txt") })
	}
	if *target == "all" || *target == "seven" {
		shared.PrintDay("seven", func() { seven.Run("inputs/SEVEN.txt") })
	}
}
