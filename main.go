package main

import (
	"fmt"
    "log"
    "os"
    "bufio"
	"github.com/marcel-baur/wrdl/routes"
    "github.com/marcel-baur/wrdl/res"
	"github.com/gin-gonic/gin"
)



func main() {
    fmt.Printf("Starting server on port 8080")
    gin.ForceConsoleColor()
    fmt.Printf("Loading words...")
    LoadWords()
    fmt.Printf("Words loaded! %d words in list.", len(res.EnglishWords))
    r := gin.Default()
    r.POST("/new", routes.CreateGameCaller)
    r.POST("/check", routes.PostSolutionCaller)
    r.Run()
}

func LoadWords() {
    f, err := os.Open("res/words.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        res.EnglishWords = append(res.EnglishWords, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
