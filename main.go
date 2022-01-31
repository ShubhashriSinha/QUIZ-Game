package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	ques string
	ans  string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFileNme := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for each question in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileNme)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file %v \n", *csvFileNme))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {

		fmt.Printf("Problem #%d: %s = \n", i+1, p.ques)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYour score : %d out of %d. \n", correct, len(problems))
			return

		case answer := <-answerCh:
			if answer == p.ans {
				fmt.Println("Correct !!")
				correct++
			} else {
				fmt.Println("Incorrect!!")
			}
		}
	}
	fmt.Printf("Your scored : %d out of %d. \n", correct, len(problems))
}
