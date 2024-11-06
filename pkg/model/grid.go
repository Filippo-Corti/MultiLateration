package model

import (
	"fmt"
	"log/slog"
)

/* Position */

type Position struct {
	Row, Col int
}

/* CellType */

type CellType int

const (
	CellTypeEmpty CellType = iota
	CellTypeStation
)

func (t CellType) String() string {
	switch t {
	case CellTypeEmpty:
		return "Empty"
	case CellTypeStation:
		return "Station"
	default:
		return "Unknown"
	}
}

func (t CellType) getTraversingCost() int {
	switch t {
	case CellTypeEmpty:
		return 1
	case CellTypeStation:
		return 2
	default:
		slog.Error("Unrecognized Cell Type")
		return -1
	}
}

/* CellState */

type CellState int

const (
	CellStateEmpty CellState = iota
	CellStateVisited
	CellStateOnFrontier
)

type Cell struct {
	P     Position
	Type  CellType
	State CellState
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%d, %d]", c.P.Col, c.P.Row)
}

/* GameGrid */

type GameGrid struct {
	Rows int
	Cols int
	Grid [][]Cell
}

func NewGameGrid(rows, cols int) *GameGrid {
	grid := getStartingGrid(rows, cols)
	return &GameGrid{
		Rows: rows,
		Cols: cols,
		Grid: grid,
	}
}

func (gm *GameGrid) GetCell(p Position) *Cell {
	if p.Col < 0 || p.Col >= gm.Cols || p.Row < 0 || p.Row >= gm.Rows {
		return nil
	}

	return &gm.Grid[p.Row][p.Col]
}

func (gm *GameGrid) String() string {
	str := fmt.Sprintf("Game{Grid: %dx%d}", gm.Rows, gm.Cols)
	for _, row := range gm.Grid {
		for _, cell := range row {
			str += cell.String()
		}
	}
	return str
}

func getStartingGrid(rows, cols int) (grid [][]Cell) {
	grid = make([][]Cell, rows)
	for r := range rows {
		grid[r] = make([]Cell, cols)
		for c := range cols {
			cell := Cell{
				P:     Position{r, c},
				Type:  CellTypeEmpty,
				State: CellStateEmpty,
			}
			grid[r][c] = cell
		}
	}
	return
}
