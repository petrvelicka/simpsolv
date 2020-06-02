package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"bufio"

	"github.com/petrvelicka/simpsolv/parse"
	"github.com/petrvelicka/simpsolv/solver"
)

func read_input(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(-2)
	}
	return string(data)
}

func read_input_stdin() string {
	var data = ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}

	return data
}

func output_roots(problem *solver.LPProblem) {
	if (*problem).Roots == nil {
			fmt.Fprintln(os.Stderr, "ERROR: SOLVE THE PROBLEM FIRST")
			return
	}
	fmt.Println("VarName\t | Value")
	fmt.Println("-----------------")

	for i, _ := range (*problem).Roots {
		fmt.Printf("%s\t | %d\n", (*problem).Variables[i], (*problem).Roots[i])
	}
	fmt.Println("Object function value for that variables:", (*problem).ObjectValue)
}

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
		fmt.Println("The system wasn't solved, the optimal solution doesn't exist")
	}
}
