package model

var Directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Station struct {
	P              Position
	visitedByRound []map[Position]bool
}

func NewStation(row, col int) *Station {
	position := Position{row, col}
	return &Station{
		P:              position,
		visitedByRound: []map[Position]bool{{position: true}},
	}
}

func (s *Station) expand(gameGrid *GameGrid) {
	expanding := false

	previousFrontier := map[Position]bool{s.P: true}
	if len(s.visitedByRound) >= 2 {
		previousFrontier = s.visitedByRound[len(s.visitedByRound)-2]
	}
	frontier := s.visitedByRound[len(s.visitedByRound)-1]
	newFrontier := map[Position]bool{}

	// Remove Visited more than 3 expansions ago
	if len(s.visitedByRound) >= 3 {
		toUnvisit := s.visitedByRound[len(s.visitedByRound)-3]
		for pos := range toUnvisit {
			cell := gameGrid.GetCell(pos)
			if cell != nil {
				cell.VisitsCount--
			}
		}
		s.visitedByRound = s.visitedByRound[1:]
	}

	// Expand Frontier
	for frontierPos := range frontier {
		frontierCell := gameGrid.GetCell(frontierPos)

		for _, dir := range Directions {
			newPos := Position{frontierPos.Row + dir[0], frontierPos.Col + dir[1]}
			newCell := gameGrid.GetCell(newPos)

			previouslyVisited := previousFrontier[newPos]
			currentlyVisited := newFrontier[newPos]

			if newCell == nil || previouslyVisited || currentlyVisited {
				continue
			}

			newCell.OnAnyFrontier = true
			newCell.VisitsCount++

			newFrontier[newPos] = true
			expanding = true
		}

		frontierCell.OnAnyFrontier = false
	}

	//Eventually Restart Expansion
	if !expanding {
		s.visitedByRound = append(s.visitedByRound, map[Position]bool{s.P: true})
	} else {
		s.visitedByRound = append(s.visitedByRound, newFrontier)
	}
}
