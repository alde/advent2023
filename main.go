package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	one "alde.nu/advent2023/day_1"
	two "alde.nu/advent2023/day_2"
	three "alde.nu/advent2023/day_3"
	four "alde.nu/advent2023/day_4"
	five "alde.nu/advent2023/day_5"
	six "alde.nu/advent2023/day_6"
	seven "alde.nu/advent2023/day_7"
	eight "alde.nu/advent2023/day_8"
	"alde.nu/advent2023/shared"
)

func main() {

	target := flag.String("target", "all", "which day to run")
	flag.Parse()

	days := map[string]func(string){
		"one":   one.Run,
		"two":   two.Run,
		"three": three.Run,
		"four":  four.Run,
		"five":  five.Run,
		"six":   six.Run,
		"seven": seven.Run,
		"eight": eight.Run,
	}

	if *target == "all" {
		for day, fun := range days {
			shared.PrintDay(day, func() { fun(fmt.Sprintf("inputs/%s.txt", strings.ToUpper(day))) })
		}
	} else {
		day := *target
		fun, ok := days[day]
		if !ok {
			fmt.Printf("unknown day %s\n", day)
			os.Exit(2)
		}
		shared.PrintDay(day, func() { fun(fmt.Sprintf("inputs/%s.txt", strings.ToUpper(day))) })
	}
}
