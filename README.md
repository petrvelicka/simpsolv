# simpsolv - wrapper for glpsol for solving linear programming problems

## Requirements

Go language compiler (you can get it from [the website](golang.org)). The installation is straightforward, just install like any other go package.

## Running

You can choose from two options:

1. Reading the input from a file -- run the program with the file name as a parameter
2. (_Is broken as for now_) Inputting the data manually from stdin. If so, do not pass any parameters via command line

## Input example

```
max 3 * x + 5 * y - 2 * z
2 * x + 3 * z <= 5
2 * z - 3 + y <= 8
```

All operands **must** be separated with spaces.
