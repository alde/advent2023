package shared

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Result struct {
	Day   string
	Task  string
	Value int
}

var (
	y = color.New(color.FgYellow)
	c = color.New(color.FgCyan)
	g = color.New(color.FgGreen)
	r = color.New(color.FgRed)
	w = color.New(color.FgWhite)
	m = color.New(color.FgMagenta)
)

func PrintDay(day string, dayFunc func()) {
	res := []string{
		r.Sprint("*"),
		g.Sprint("*"),
		c.Sprint("*"),
		y.Sprint("DAY"),
		m.Sprint(strings.ToUpper(day)),
		c.Sprint("*"),
		g.Sprint("*"),
		r.Sprint("*"),
	}

	fmt.Printf("\n%s\n", strings.Join(res, " "))
	dayFunc()
}

func PrintResult(taskFunc func() *Result) {
	start := time.Now()
	result := taskFunc()
	elapsed := time.Now().Sub(start)
	res := []string{
		w.Sprint("⭐"),
		y.Sprint("Day"),
		c.Sprint(result.Day),
		y.Sprint("Task"),
		g.Sprint(result.Task),
		w.Sprint(":"),
		m.Sprintf("%v", result.Value),
		w.Sprint("["),
		r.Sprintf("%s", elapsed),
		w.Sprint("]"),
	}

	fmt.Printf("%s\n", strings.Join(res, " "))
}

func Timer[T any](partFunc func() *Result) (*Result, time.Duration) {
	T1 := time.Now()
	result := partFunc()
	elapsed := T1.Sub(time.Now())

	return result, elapsed
}
