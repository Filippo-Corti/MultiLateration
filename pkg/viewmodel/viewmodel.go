package viewmodel

import (
	"mlat/pkg/model"
)

type GridViewModel struct {
	Game *model.GameGrid
}

func (gvm *GridViewModel) GetGameGrid() [][]model.Cell {
	return gvm.Game.Grid
}
