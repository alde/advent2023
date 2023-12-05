package main

import (
	"alde.nu/advent2023/five"
	"alde.nu/advent2023/four"
	"alde.nu/advent2023/one"
	"alde.nu/advent2023/shared"
	"alde.nu/advent2023/three"
	"alde.nu/advent2023/two"
)

func main() {
	shared.PrintDay("one", func() { one.Run("inputs/ONE.txt") })
	shared.PrintDay("two", func() { two.Run("inputs/TWO.txt") })
	shared.PrintDay("three", func() { three.Run("inputs/THREE.txt") })
	shared.PrintDay("four", func() { four.Run("inputs/FOUR.txt") })
	shared.PrintDay("five", func() { five.Run("inputs/FIVE.txt") })
}
