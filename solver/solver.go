package solver

import (
	"fmt"
	"strconv"
	"os/exec"
	"strings"
	"bufio"
	"io"
	"errors"
)

type LPProblem struct {
	Variables []string
	Object string
	Minmax bool // false = min, true = max
	Constraints []string
	Roots []int
	Program []string
	ObjectValue int
}

func Generate_program(problem *LPProblem) []string {
	var program = []string {}
	// variables
	for _, val := range (*problem).Variables {
		program = append(program, gen_var(val))
	}
	// object function
	program = append(program, gen_object((*problem).Minmax, (*problem).Object))


	// constraints
	for i, val := range (*problem).Constraints {
		program = append(program, gen_constraint(val, i))
	}
	program = append(program, "solve;")
	// output of variables
	for _, val := range (*problem).Variables {
		program = append(program, gen_var_output(val))
	}

	program = append(program, gen_obj_value_output((*problem).Object))

	program = append(program, "end;")
	return program
}

func Check_glpsol() bool {
	_, err := exec.LookPath("glpsol")
	if err != nil {
		return false
	}
	return true

}

func Run_glpsol(problem *LPProblem) ([]int, error) {
	cmd := exec.Command("glpsol", "-m", "/dev/stdin")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	for _, val := range (*problem).Program {
		io.WriteString(stdin, val + "\n")
	}

	stdin.Close()
	scanner := bufio.NewScanner(stdout)

	var solved = false
	var Roots = []int{}
	for scanner.Scan() {
		m := scanner.Text()
		if m == "INTEGER OPTIMAL SOLUTION FOUND" {
			solved = true
		}
		if solved && m[:10] == "SOLVER VAR" {
			root, _ := strconv.Atoi(strings.Split(m, " ")[3])
			Roots = append(Roots, root)
		}
		if solved && m[:10] == "SOLVER OBJ" {
			objvalue, _ := strconv.Atoi(strings.Split(m, " ")[2])
			(*problem).ObjectValue = objvalue
		}
	}
	(*problem).Roots = Roots
	if (!solved) {
		return Roots, errors.New("not solved")
	} else {
		return Roots, nil
	}
}

func Solve(problem *LPProblem) ([]int, error) {
	(*problem).Program = Generate_program(problem)
	if !Check_glpsol() {
		return (*problem).Roots, errors.New("glpsol not found")
	}
	_, err := Run_glpsol(problem)
	return (*problem).Roots, err
}
