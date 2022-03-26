package logic

import (
	"testing"
)

func TestCheckSolution(t *testing.T) {
	word := "notes"
	solution := "slosh"
	want := []LetterState{"grey", "yellow", "grey", "grey", "yellow"}
    GenericTestSolution(word, solution,want, t)
}

func TestCheckAlmostCorrectSolution(t *testing.T) {
	word := "slash"
	solution := "slosh"
	want := []LetterState{"green", "green", "grey", "green", "green"}
    GenericTestSolution(word, solution,want, t)
}

func TestCheckCorrectSolution(t *testing.T) {
	word := "slash"
	solution := "slash"
	want := []LetterState{"green", "green", "green", "green", "green"}
    GenericTestSolution(word, solution,want, t)
}

func TestCheckDoubleLetterSolution(t *testing.T) {
	word := "lasso"
	solution := "slosh"
	want := []LetterState{"yellow", "grey", "yellow", "green", "yellow"}
    GenericTestSolution(word, solution,want, t)
}

func TestCheckDoubleLetterSolutionTwo(t *testing.T) {
	word := "silos"
	solution := "slosh"
	want := []LetterState{"green", "grey", "yellow", "yellow", "yellow"}
    GenericTestSolution(word, solution,want, t)
}

func GenericTestSolution(word string, solution string, want []LetterState, t *testing.T) {
	game := Game{
		Solution: solution,
	}
	result := CheckSolution(word, &game)
	if !CompareSolutions(want, result) {
		t.Fatalf(`CheckSolution("%q", "%q") = %q, want match for %q`, word, solution, result, want)
	}
}

func CompareSolutions(want []LetterState, result []LetterState) bool {
	for pos, elem := range want {
		if elem != result[pos] {
			return false
		}
	}
	return true
}
