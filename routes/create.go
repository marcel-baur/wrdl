package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

    c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
