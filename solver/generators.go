package solver

func gen_var(name string) string {
	return "var " + name + " >= 0, integer;"
}

func gen_constraint(str string, counter int) string {
	return "c" + strconv.Itoa(counter) + ": " + str + ";"
}

func get_minmax_word(minmax bool) string {
	if minmax {
		return "maximize"
	}
	return "minimize"
}

func gen_object(minmax bool, str string) string {
	return get_minmax_word(minmax) + " obj: " + str + ";"
}

func gen_var_output(name string) string {
	return "printf \"SOLVER VAR " + name + " %d\\n\", "+ name + ";"
}

func gen_obj_value_output(name string) string {
	return "printf \"SOLVER OBJ: %d\\n\", " + name + ";"
}
