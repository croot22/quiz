package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer string
}

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
	numCorrect := 0
	for i, problem := range problems{
		giveQuestion(problem, i+1)
		submissionCh := make(chan string)
		go func() {
			submission := getSubmission()
			submissionCh <- submission
		}()
		select {
		case <-timer.C:
			giveFinalGrade(numCorrect, len(problems))
			return
		case submission := <-submissionCh:
			gradeSubmission(submission, problem.answer, &numCorrect)
		}
	}
	giveFinalGrade(numCorrect, len(problems))
}

func giveQuestion(problem problem, line int){
	fmt.Printf("Problem #%d: %s = \n", line, problem.question)
}

func getSubmission() string{
	var submission string
	_, err := fmt.Scanf("%s\n", &submission)
	if err != nil {
		exit("Failed to scan submitted answer")
	}
	return submission
}

func gradeSubmission(submission, answer string, numCorrect *int){
	if submission == answer {
		*numCorrect++
	}
}

func giveFinalGrade(numCorrect, numProblems int){
	fmt.Printf("You scored %d out of %d.\n", numCorrect, numProblems)
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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
