package res

import (
    "math/rand"
    "github.com/marcel-baur/wrdl/utils"
)

type WordList []string

var EnglishWords WordList


func FetchWord(l int) string {
    filtered := utils.Filter(EnglishWords, func(word string) bool {
        return len(word) == l
    })
    index := rand.Intn(len(filtered))
    return filtered[index]
}
