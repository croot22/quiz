package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var numCorrect int = 0

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil{
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	runQuiz(problems)

	fmt.Printf("You scored %d out of %d.\n", numCorrect, len(problems))
}

func runQuiz(problems []problem){
	for i, problem:= range problems{
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var submission string
		fmt.Scanf("%s\n", &submission)
		if submission == problem.answer {
			numCorrect++
		}
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
