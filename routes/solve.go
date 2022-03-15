package routes

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
	"github.com/marcel-baur/wrdl/logic"
)

type PostSolveRequest struct {
    Hash uuid.UUID
    Word string
}

func PostSolutionCaller(c *gin.Context) {
    var request PostSolveRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{ERROR_KEY: err.Error()})
        return
    }

    game := logic.GMap[request.Hash]
    if game == nil {
        c.JSON(http.StatusBadRequest, gin.H{ERROR_KEY: "Game not found!"})
    }
    if len(request.Word) != len(game.Solution) {
        err_string := fmt.Sprintf("The proposed solution must have length %d", len(game.Solution))
        c.JSON(http.StatusBadRequest, gin.H{ERROR_KEY: err_string})
    }
}
