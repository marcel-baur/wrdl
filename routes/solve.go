package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcel-baur/wrdl/logic"
	"net/http"
)

type PostSolveRequest struct {
	Hash uuid.UUID
	Word string
}

func PostSolutionCaller(c *gin.Context) {
	var request PostSolveRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ERROR_RESPONSE_KEY: err.Error()})
		return
	}

	game := logic.GMap[request.Hash]
	if game == nil {
		c.JSON(http.StatusBadRequest, gin.H{ERROR_RESPONSE_KEY: "Game not found!"})
	}
	if len(request.Word) != len(game.Solution) {
		err_string := fmt.Sprintf("The proposed solution must have length %d", len(game.Solution))
		c.JSON(http.StatusBadRequest, gin.H{ERROR_RESPONSE_KEY: err_string})
        return
	}
    if len(game.Attempts) > 6 {
        c.JSON(http.StatusOK, gin.H{"correct_solution": game.Solution})
        return
    }
    result := logic.CheckSolution(request.Word, game)
    game.Attempts = append(game.Attempts, request.Word)
    c.JSON(http.StatusOK, gin.H{"result": result})
}
