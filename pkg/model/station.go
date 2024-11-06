package model

var Directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Station struct {
	P         Position
	frontier  []Position
	canExpand bool
}

func NewStation(row, col int) *Station {
	position := Position{row, col}
	return &Station{
		P:         position,
		frontier:  []Position{position},
		canExpand: true,
	}
}

func (s *Station) expand(gameGrid *GameGrid) {
	if !s.canExpand {
		return
	}

	expanding := false
	newFrontier := []Position{}
	for _, frontierPos := range s.frontier {
		frontierCell := gameGrid.GetCell(frontierPos)

		for _, dir := range Directions {
			newPos := Position{frontierPos.Row + dir[0], frontierPos.Col + dir[1]}
			newCell := gameGrid.GetCell(newPos)

			if newCell == nil || newCell.State != CellStateEmpty {
				continue
			}

			newCell.State = CellStateOnFrontier

			newFrontier = append(newFrontier, newPos)
			expanding = true
		}

		frontierCell.State = CellStateVisited
	}
	s.frontier = newFrontier

	if !expanding {
		s.canExpand = false
	}
}
