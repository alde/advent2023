package main

import (
	"flag"

	"alde.nu/advent2023/five"
	"alde.nu/advent2023/four"
	"alde.nu/advent2023/one"
	"alde.nu/advent2023/shared"
	"alde.nu/advent2023/six"
	"alde.nu/advent2023/three"
	"alde.nu/advent2023/two"
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
}
