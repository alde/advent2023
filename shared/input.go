package shared

import (
	"os"
	"strings"
)

const PS = string(os.PathSeparator)

func LoadInput(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(dir + PS + filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func LoadInputAsStringSlice(filename string) []string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(dir + PS + filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
