package model

import "fmt"

type Position struct {
	Row, Col int
}

type Cell struct {
	P Position
}

func (c *Cell) toString() string {
	return fmt.Sprintf("[%d, %d]", c.P.Col, c.P.Row)
}

type GameModel struct {
	Rows int
	Cols int
	Grid [][]Cell
}

func NewGameModel(rows, cols int) *GameModel {
	grid := getStartingGrid(rows, cols)
	return &GameModel{
		Rows: rows,
		Cols: cols,
		Grid: grid,
	}
}

func (gm *GameModel) toString() string {
	str := fmt.Sprintf("Game{Grid: %dx%d}", gm.Rows, gm.Cols)
	for _, row := range gm.Grid {
		for _, cell := range row {
			str += cell.toString()
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
				P: Position{r, c},
			}
			grid[r][c] = cell
		}
	}
	return
}
