package quiz

import (
	"reflect"
	"testing"
)

var prob1 = problem{
question: "1+1",
answer: "2",
}

var prob2 = problem{
question: "2+1",
answer: "3",
}
var prob3 = problem{
question: "3+1",
answer: "4",
}

var problems = []problem{prob1, prob2, prob3}

func TestShuffle(t *testing.T){
	shuffled := problems
	shuffleQuiz(shuffled)
	original := problems

	if reflect.DeepEqual(original, shuffled) {
		t.Errorf("Quiz was not shuffled")
	}
}

func TestGetSubmission(t *testing.T) {
	getSubmission()
}
