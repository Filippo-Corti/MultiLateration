package model

var Directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Station struct {
	P           Position
	depthFactor int
	frontier    []Position
}

func NewStation(row, col int) *Station {
	position := Position{row, col}
	return &Station{
		P:           position,
		depthFactor: 0,
		frontier:    []Position{position},
	}
}

func (s *Station) expand(gameGrid *GameGrid) {
	newFrontier := []Position{}
	for _, frontierPos := range s.frontier {
		frontierCell := gameGrid.GetCell(frontierPos)

		for _, dir := range Directions {
			newPos := Position{frontierPos.Row + dir[0], frontierPos.Col + dir[1]}
			newCell := gameGrid.GetCell(newPos)

			if newCell == nil ||  newCell.State != CellStateEmpty {
				continue
			}

			newCell.State = CellStateOnFrontier

			newFrontier = append(newFrontier, newPos)
		}

		frontierCell.State = CellStateVisited
	}
	s.frontier = newFrontier
}
