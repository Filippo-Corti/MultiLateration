package main

import (
	"log/slog"
	"mlat/pkg/constants"
	"mlat/pkg/model"
	"mlat/pkg/view"
	"mlat/pkg/viewmodel"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	game := model.NewGameGrid(constants.GRID_WIDTH, constants.GRID_HEIGHT)
	viewmodel := &viewmodel.GridViewModel{Game: game}
	view := &view.GameView{ViewModel: viewmodel}

	ebiten.SetWindowSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Just a Board (for now)")

	if err := ebiten.RunGame(view); err != nil {
		slog.Error("Ebiten Game Error", "msg", err)
	}
}
