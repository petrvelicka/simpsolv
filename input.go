package main

import (
	"os"
	"bufio"
	"io/ioutil"
	"fmt"
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
