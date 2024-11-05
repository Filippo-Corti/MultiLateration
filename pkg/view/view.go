package view

import (
	"fmt"
	"mlat/pkg/viewmodel"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameView struct {
	ViewModel *viewmodel.GridViewModel
}

func (gv *GameView) Update() error {
	return nil
}

func (gv *GameView) Draw(screen *ebiten.Image) {
	str := "";
	gameGrid := gv.ViewModel.GetGameGrid()
	for _, row := range gameGrid {
		for _, cell := range row {
			str += fmt.Sprintf("%v | ", cell)
		}
		str += "\n"
	}

	ebitenutil.DebugPrint(screen, str)
}

func (gv *GameView) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
