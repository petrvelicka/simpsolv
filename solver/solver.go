package solver

type LPProblem struct {
	Variables []string
	Object string
	Minmax bool // false = min, true = max
	Constraints []string
}
