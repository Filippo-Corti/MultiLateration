package main

import (
	"log/slog"
	"mlat/pkg/model"
	"mlat/pkg/view"
	"mlat/pkg/viewmodel"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	game := model.NewGameModel(10, 10)
	viewmodel := &viewmodel.GridViewModel{Game: game}
	view := &view.GameView{ViewModel: viewmodel}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(view); err != nil {
		slog.Error("Problema", "errore", err)
	}
}
