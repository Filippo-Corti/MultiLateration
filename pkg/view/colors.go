package view

import (
	"image/color"
	"mlat/pkg/model"
)

func cellColor(cell model.Cell) color.Color {
	switch {
	case cell.OnAnyFrontier:
		return color.RGBA{200, 182, 255, 1}
	case cell.VisitsCount > 0:
		return color.RGBA{255, 214, 255, 1}
	default:
		return color.RGBA{187, 208, 255, 1}
	}
}

func cellBackgroundColor(cell model.Cell) color.Color {
	switch {
	case cell.OnAnyFrontier:
		return color.RGBA{164, 134, 255, 1}
	case cell.VisitsCount > 0:
		return color.RGBA{255, 162, 255, 1}
	default:
		return color.RGBA{138, 174, 255, 1}
	}
}

func cellText(cType model.CellType) string {
	switch cType {
	case model.CellTypeStation:
		return "S"
	default:
		return ""
	}
}
