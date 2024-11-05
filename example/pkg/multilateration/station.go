package multilateration

type StationData struct {
	Position     Position
	DistToTarget float64
}

func NewStationData(x, y, d float64) *StationData {
	return &StationData{
		Position:     NewPosition(x, y),
		DistToTarget: d,
	}
}
