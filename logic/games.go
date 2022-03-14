package logic

import (
    "github.com/google/uuid"
    "github.com/marcel-baur/wrdl/res"
)

type Game struct {
    Solution string
    Attempts []string
    Hash uuid.UUID
}

func CreateGame(len int) *Game {
    solution := res.FetchWord(len)
    game := new(Game)
    game.Solution = solution
    game.Hash = uuid.New()
    return game
}
