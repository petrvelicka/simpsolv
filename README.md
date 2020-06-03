# simpsolv - wrapper for glpsol for solving linear programming problems

## Requirements

Go language compiler (you can get it from [the website](https://golang.org/)). After installing the compiler, the installation of this package is pretty straightforward, just install like any other go package.

## Running

You can choose from two options:

1. Reading the input from a file - run the program with the file name as a parameter
2. Inputting the data manually from stdin. If so, do not pass any parameters via command line. After entering all constraints and target function press `Ctrl+D` - to close the stdin and start solving the problem

## Input example

```
max 3 * x + 5 * y - 2 * z
2 * x + 3 * z <= 5
2 * z - 3 + y <= 8
```

All operands **must** be separated with spaces.

## Structure LPProblem

1. Variables - array of strings - variable names in any order
2. Object - Object function definition 
3. Minmax - Flag that determines if the object function should be maximized (true) or minimized (false)
4. Constraints - array of strings - constraints
5. Roots - after solving - values for variables in the same order as Variables
6. Program - _glpsol_ program text (after running `solver.Generate_program`)
7. ObjectValue - object function value
