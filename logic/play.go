package logic

type LetterState string

const (
	GREY   LetterState = "grey"
	GREEN  LetterState = "green"
	YELLOW LetterState = "yellow"
)

func CheckSolution(word string, game *Game) []LetterState {
	var result []LetterState
	if word == game.Solution {
		for _, _ = range word {
			result = append(result, GREEN)
		}
		return result
	}
	runified_sol := []rune(game.Solution)
	used_char := make(map[rune]int)
	occurrences := make(map[rune]int)
	for _, char := range game.Solution {
		val, ok := occurrences[char]
		if ok {
			occurrences[char] = val + 1
			continue
		}
		occurrences[char] = 1
	}
	for pos, char := range word {
		if runified_sol[pos] == char {
			result = append(result, GREEN)
			continue
		}
		if ContainsRune(runified_sol, char) {
			conf, ok := used_char[char]
			occ, _ := occurrences[char]
			if !ok {
				used_char[char] = 1
				result = append(result, YELLOW)
				continue
			}
			if occ > conf {
				used_char[char] = conf + 1
				result = append(result, YELLOW)
				continue
			}
			result = append(result, GREY)
			continue
		}
		result = append(result, GREY)
	}
	return result
}

func ContainsRune(arr []rune, item rune) bool {
	for _, x := range arr {
		if x == item {
			return true
		}
	}
	return false
}
