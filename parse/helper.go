package parse

func check_word(word string) bool {
	operators := []string {"<=", "=", ">=", "min", "max", "+", "-", "*"}
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
