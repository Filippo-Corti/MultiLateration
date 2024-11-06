package view

import (
	"bytes"
	"image"
	"image/color"
	"log/slog"
	"mlat/pkg/constants"
	"mlat/pkg/model"
	"mlat/pkg/viewmodel"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	fontFace      *text.GoTextFaceSource
	cellImage     = ebiten.NewImage(constants.CELLSIZE, constants.CELLSIZE)
	innerCellRect = image.Rect(constants.CELL_INNER_BORDER, constants.CELL_INNER_BORDER, constants.CELLSIZE-constants.CELL_INNER_BORDER, constants.CELLSIZE-constants.CELL_INNER_BORDER)
)

func init() {
	// Setup Font
	ff, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		slog.Error("couldn't load font")
	}
	fontFace = ff

	//Setup Cell Image
	cellImage.Fill(color.White)
}

type GameView struct {
	ViewModel *viewmodel.GridViewModel
}

func (gv *GameView) Update() error {
	
}

func (gv *GameView) Draw(screen *ebiten.Image) {
	gameGrid := gv.ViewModel.GetGameGrid()

	for _, row := range gameGrid {
		for _, cell := range row {
			drawCell(screen, cell)
		}
	}

}

func (gv *GameView) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT
}

func drawCell(screen *ebiten.Image, cell model.Cell) {
	outerRect := cellImage
	x := float64(cell.P.Col * constants.CELLSIZE)
	y := float64(cell.P.Row * constants.CELLSIZE)

	// Draw Rectangles
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(cellBackgroundColor(cell.State))

	innerRect := outerRect.SubImage(innerCellRect).(*ebiten.Image)
	innerRect.Fill(cellColor(cell.State))

	screen.DrawImage(outerRect, op)

	// Draw Text
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(x+float64(constants.CELLSIZE)/2, y+float64(constants.CELLSIZE)/2)
	textOp.ColorScale.ScaleWithColor(color.Black)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	text.Draw(screen, cellText(cell.Type), &text.GoTextFace{
		Source: fontFace,
		Size:   24,
	}, textOp)
}
