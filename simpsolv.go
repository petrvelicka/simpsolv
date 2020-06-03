package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/petrvelicka/simpsolv/parse"
	"github.com/petrvelicka/simpsolv/solver"
)

func main() {
	args := os.Args
	var lines []string

	if len(args) > 1 {
		lines = strings.Split(read_input(args[1]), "\n")
	} else {
		lines = strings.Split(read_input_stdin(), "\n")
	}
	
	// remove the last pointless line
	lines = lines[:len(lines)-1]


	var problem solver.LPProblem
	for _, line := range lines {
		parse.ParseLine(line, &problem)
	}

	_, err := solver.Solve(&problem)
	if err == nil {
		fmt.Println("Solution: ")
		output_roots(&problem)
	} else {
		fmt.Fprintln(os.Stderr, "The system wasn't solved, the optimal solution doesn't exist")
	}
}
