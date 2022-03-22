package logic

import (
	"github.com/google/uuid"
	"github.com/marcel-baur/wrdl/res"
)

type GameMap map[uuid.UUID]*Game

var GMap GameMap

type Game struct {
	Solution string
	Attempts []string
	Hash     uuid.UUID
}

func CreateGame(len int) *Game {
	solution := res.FetchWord(len)
	game := new(Game)
	game.Solution = solution
	game.Hash = uuid.New()
	GMap[game.Hash] = game
	return game
}
