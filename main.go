package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var numCorrect int = 0

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	timeLimit := flag.Int("limit", 30 , "the time limit for the quiz in seconds")

	shouldShuffle := flag.Bool("shuffle", false, "a value of true will shuffle the questions")

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

	if *shouldShuffle{
		shuffleQuiz(problems)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	runQuiz(problems, timer)
}

func runQuiz(problems []problem, timer *time.Timer){
	for i, problem:= range problems{
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		submissionCh := make(chan string)
		go func() {
			var submission string
			fmt.Scanf("%s\n", &submission)
			submissionCh <- submission
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", numCorrect, len(problems))
			return
		case submission := <-submissionCh:
			if submission == problem.answer {
				numCorrect++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", numCorrect, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(strings.ToLower(line[1])),
		}
	}
	return ret
}

func shuffleQuiz(problems []problem){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) {problems[i], problems[j] = problems[j], problems[i]})
}

type problem struct {
	question string
	answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
