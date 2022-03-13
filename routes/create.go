package routes

import (
	"net/http"
    "math/rand"
	"github.com/gin-gonic/gin"
    "github.com/marcel-baur/wrdl/res"
    "github.com/marcel-baur/wrdl/utils"

)

type NewGame struct {
    Letters int
}

func CreateGameCaller(c *gin.Context) {
    var request NewGame
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if request.Letters < 4 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Need at least four letters!"})
        return
    }
    if request.Letters > 7 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not more than seven letters allowed"})
        return
    }
    word := FetchWord(request.Letters)
    c.JSON(http.StatusOK, gin.H{"word": word})
}

func FetchWord(l int) string {
    filtered := utils.Filter(res.EnglishWords, func(word string) bool {
        return len(word) == l
    })
    index := rand.Intn(len(filtered))
    return filtered[index]
}
