package main

import (
	"fmt"
	"os"
	
	"github.com/petrvelicka/simpsolv/solver"
)

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

