package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/marcel-baur/wrdl/logic"
)

type NewGame struct {
    Letters int
}

func CreateGameCaller(c *gin.Context) {
    var request NewGame
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{ERROR_KEY: err.Error()})
        return
    }
    if request.Letters < 4 {
        c.JSON(http.StatusBadRequest, gin.H{ERROR_KEY: "Need at least four letters!"})
        return
    }
    if request.Letters > 7 {
        c.JSON(http.StatusBadRequest, gin.H{ERROR_KEY: "Not more than seven letters allowed"})
        return
    }
    game := logic.CreateGame(request.Letters)
    c.JSON(http.StatusOK, gin.H{"game": game})
}

