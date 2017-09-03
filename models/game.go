package models

import (
	"os"
	"github.com/michaelbui/bingo-backend/helpers"
)

type GameModel struct {

}

var (
	game *GameModel
	gameLockFile string = "./game.lock"
)

func Game() *GameModel {
	if game == nil {
		game = &GameModel{}
	}
	return game
}

func (g *GameModel) Activate() error {
	if _, err := os.Stat(gameLockFile); os.IsNotExist(err) {
		if _, err = os.Create(gameLockFile); err != nil {
			return err
		}
	}
	return nil
}

func (g *GameModel) Reset() error {
	return helpers.Database().Init()
}

func (g *GameModel) IsLocked() bool {
	_, err := os.Stat(gameLockFile)
	return !os.IsNotExist(err)
}
