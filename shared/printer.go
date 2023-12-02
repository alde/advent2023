package shared

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Result[T any] struct {
	Day   string
	Task  string
	Value T
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

func PrintResult[T any](result *Result[T]) {
	res := []string{
		w.Sprint("⭐"),
		y.Sprint("Day"),
		c.Sprint(result.Day),
		y.Sprint("Task"),
		g.Sprint(result.Task),
		w.Sprint(":"),
		m.Sprintf("%v", result.Value),
	}

	fmt.Printf("%s\n", strings.Join(res, " "))
}
