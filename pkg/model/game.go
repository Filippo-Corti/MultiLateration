package model

import "errors"

type Game struct {
	Grid     *GameGrid
	Stations []*Station
}

func NewGame(rows, cols int) *Game {
	return &Game{
		Grid:     NewGameGrid(rows, cols),
		Stations: []*Station{},
	}
}

func (g *Game) AddStation(s Station) error {

	cell := g.Grid.GetCell(s.P)

	if cell == nil || cell.Type != CellTypeEmpty {
		return errors.New("station can't be placed on non-empty or non-existing cell")
	}

	cell.Type = CellTypeStation
	g.Stations = append(g.Stations, &s)
	return nil
}

func (g *Game) ExpandAllStations() {
	for _, s := range g.Stations {
		s.expand(g.Grid)
	}
}
