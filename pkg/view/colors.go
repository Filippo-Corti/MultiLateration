package view

import (
	"image/color"
	"mlat/pkg/model"
)

func cellColor(state model.CellState) color.Color {
	switch state {
	case model.CellStateEmpty:
		return color.RGBA{187, 208, 255, 1}
	case model.CellStateVisited:
		return color.RGBA{255, 214, 255, 1}
	case model.CellStateOnFrontier:
		return color.RGBA{200, 182, 255, 1}
	default:
		return nil
	}
}

func cellBackgroundColor(state model.CellState) color.Color {
	switch state {
	case model.CellStateEmpty:
		return color.RGBA{138, 174, 255, 1}
	case model.CellStateVisited:
		return color.RGBA{255, 162, 255, 1}
	case model.CellStateOnFrontier:
		return color.RGBA{164, 134, 255, 1}
	default:
		return nil
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