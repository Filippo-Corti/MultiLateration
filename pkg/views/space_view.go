package views

import (
	"image"
	"log/slog"
	"positioning/pkg/models"

	"github.com/fogleman/gg"
)

const STATION_RADIUS = 5

type SpaceView struct {
	Space *models.Space
}

func NewSpaceView(space *models.Space) *SpaceView {
	return &SpaceView{
		Space: space,
	}
}

func (sv *SpaceView) Render() *image.RGBA {
	image := image.NewRGBA(image.Rect(0, 0, sv.Space.Width, sv.Space.Height))
	ctx := gg.NewContextForRGBA(image)

	ctx.SetRGB(0, 0, 0)
	for _, station := range sv.Space.Stations {
		ctx.DrawCircle(station.X, station.Y, STATION_RADIUS)
		ctx.Fill()
	}

	slog.Info("Space View Rendered")

	return image
}
