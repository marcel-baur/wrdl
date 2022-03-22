package logic

import (
	"testing"
)

func TestCheckSolution(t *testing.T) {
	word := "notes"
	solution := "slosh"
	game := Game{
		Solution: solution,
	}
	want := []LetterState{"grey", "yellow", "grey", "grey", "yellow"}
	result := CheckSolution(word, game)
	if !CompareSolutions(want, result) {
        t.Fatalf(`CheckSolution("%q", "%q") = %q, want match for %q`,word, solution, result, want)
	}
}

func TestCheckAlmostCorrectSolution(t *testing.T) {
	word := "slash"
	solution := "slosh"
	game := Game{
		Solution: solution,
	}
	want := []LetterState{"green", "green", "grey", "green", "green"}
	result := CheckSolution(word, game)
	if !CompareSolutions(want, result) {
        t.Fatalf(`CheckSolution("%q", "%q") = %q, want match for %q`,word, solution, result, want)
	}
}

func TestCheckCorrectSolution(t *testing.T) {
	word := "slash"
	solution := "slash"
	game := Game{
		Solution: solution,
	}
	want := []LetterState{"green", "green", "green", "green", "green"}
	result := CheckSolution(word, game)
	if !CompareSolutions(want, result) {
        t.Fatalf(`CheckSolution("%q", "%q") = %q, want match for %q`,word, solution, result, want)
	}
}

func TestCheckDoubleLetterSolution(t *testing.T) {
	word := "lasso"
	solution := "slosh"
	game := Game{
		Solution: solution,
	}
	want := []LetterState{"yellow", "grey", "yellow", "green", "yellow"}
	result := CheckSolution(word, game)
	if !CompareSolutions(want, result) {
        t.Fatalf(`CheckSolution("%q", "%q") = %q, want match for %q`,word, solution, result, want)
	}
}

func TestCheckDoubleLetterSolutionTwo(t *testing.T) {
	word := "silos"
	solution := "slosh"
	game := Game{
		Solution: solution,
	}
	want := []LetterState{"green", "grey", "yellow", "yellow", "yellow"}
	result := CheckSolution(word, game)
	if !CompareSolutions(want, result) {
        t.Fatalf(`CheckSolution("%q", "%q") = %q, want match for %q`,word, solution, result, want)
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
