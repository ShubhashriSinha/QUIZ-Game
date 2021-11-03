package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileNme := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFileNme)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file %v \n", *csvFileNme))
		//os.Exit(1)
	}
	//_ = file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	//fmt.Println(lines)
	//fmt.Printf("Type of line is %T", lines)

	problems := parseLines(lines)
	//fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		var answer string
		fmt.Scanf("%s \n", &answer)
		fmt.Printf("Type of answer : %T and type of input: %T \n", p.a, answer)
		fmt.Printf("answer : %v and input: %v \n", p.a, answer)
		if answer == p.a {
			fmt.Println("Correct !!")
			correct++
		} else {
			fmt.Println("Incorrect!!")
		}
	}

	fmt.Printf("Your score : %d out of %d. \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
