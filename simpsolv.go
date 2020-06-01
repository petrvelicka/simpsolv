package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/petrvelicka/simpsolv/parse"
	"github.com/petrvelicka/simpsolv/solver"
)

func read_input(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}

func main() {
	args := os.Args
	var lines []string

	if len(args) > 1 {
		lines = strings.Split(read_input(args[1]), "\n")

		// remove the last pointless line
		lines = lines[:len(lines)-1]

	} else {
		fmt.Println("not implemented")
		// TODO: add ability to input data from stdin
		return
	}
	var problem solver.LPProblem
	for _, line := range lines {
		parse.ParseLine(line, &problem)
	}

	fmt.Println(problem.Variables)
	fmt.Println(problem.Object)
	fmt.Println(problem.Minmax)
	fmt.Println(problem.Constraints)
}
