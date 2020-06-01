package parse

import (
		"strings"
		"strconv"

		"github.com/petrvelicka/simpsolv/solver"
)

func check_word(word string) bool {
	operators := []string {"<=", "=", ">=", "min", "max", "+", "-"}
	if exists(operators, word) {
		return false
	}
	return true
}

func exists(vars []string, value string) bool {
	for _, val := range vars {
		if value == val {
			return true
		}
	}
	return false
}

func ParseLine(line string, problem *solver.LPProblem) {
	// check if the line is object function
	lexems := strings.Split(line, " ")
	if lexems[0] == "min" || lexems[0] == "max" {
		if lexems[0] == "max" {
			(*problem).Minmax = true
		} else {
			(*problem).Minmax = false
		}
		(*problem).Object = line[4:] // min/max and a space
	} else {
		(*problem).Constraints = append((*problem).Constraints, line)
	}

	for _, lexeme := range lexems {
		if check_word(lexeme) {
			_, err := strconv.Atoi(lexeme)
			if err != nil && !exists((*problem).Variables, lexeme) {
				(*problem).Variables = append((*problem).Variables, lexeme)
			}
		}
	}
}
