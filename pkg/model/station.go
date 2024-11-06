package model

var Directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Station struct {
	P              Position
	visitedByRound [][]Position
}

func NewStation(row, col int) *Station {
	position := Position{row, col}
	return &Station{
		P:              position,
		visitedByRound: [][]Position{{position}},
	}
}

func (s *Station) expand(gameGrid *GameGrid) {
	expanding := false
	frontier := s.visitedByRound[len(s.visitedByRound)-1]
	newFrontier := []Position{}

	//slog.Info("Expanding with visited length of", "value", len(s.visitedByRound))

	// Remove Visited more than 3 expansions ago
	if len(s.visitedByRound) >= 3 {
		toUnvisit := s.visitedByRound[len(s.visitedByRound)-3]
		for _, pos := range toUnvisit {
			cell := gameGrid.GetCell(pos)
			if cell != nil {
				cell.State = CellStateEmpty
			}
		}
		s.visitedByRound = s.visitedByRound[1:]
	}

	// Expand Frontier
	for _, frontierPos := range frontier {
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

	//Eventually Restart Expansion
	if !expanding {
		s.visitedByRound = append(s.visitedByRound, []Position{s.P})
	} else {
		s.visitedByRound = append(s.visitedByRound, newFrontier)
	}
}
