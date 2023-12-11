package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"alde.nu/advent2023/day_01"
	"alde.nu/advent2023/day_02"
	"alde.nu/advent2023/day_03"
	"alde.nu/advent2023/day_04"
	"alde.nu/advent2023/day_05"
	"alde.nu/advent2023/day_06"
	"alde.nu/advent2023/day_07"
	"alde.nu/advent2023/day_08"
	"alde.nu/advent2023/day_09"
	"alde.nu/advent2023/day_10"
	"alde.nu/advent2023/shared"
)

func main() {

	target := flag.String("target", "all", "which day to run")
	flag.Parse()

	days := map[string]func(string){
		"one":   day_01.Run,
		"two":   day_02.Run,
		"three": day_03.Run,
		"four":  day_04.Run,
		"five":  day_05.Run,
		"six":   day_06.Run,
		"seven": day_07.Run,
		"eight": day_08.Run,
		"nine":  day_09.Run,
		"ten":   day_10.Run,
		// "eleven": day_11.Run,
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
