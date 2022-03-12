package main

import (
	"fmt"
	"github.com/marcel-baur/wrdl/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    fmt.Printf("Starting server on port 8080")
    gin.ForceConsoleColor()
    r := gin.Default()
    r.POST("/new", routes.CreateGameCaller)
    r.POST("/check", routes.PostSolutionCaller)
    r.Run()
}
