package views

import (
	"image"
	"log/slog"
	"positioning/pkg/models"
	ml "positioning/pkg/multilateration"

	"github.com/fogleman/gg"
)

const STATION_RADIUS = 5
const TARGET_RADIUS = 5

type SpaceView struct {
	Space *models.Space
}

func NewSpaceView(space *models.Space) *SpaceView {
	return &SpaceView{
		Space: space,
	}
}

func (sv *SpaceView) Render(showStationsArea bool) *image.RGBA {
	image := image.NewRGBA(image.Rect(0, 0, sv.Space.Width, sv.Space.Height))
	ctx := gg.NewContextForRGBA(image)

	for _, station := range sv.Space.Stations {
		renderStation(ctx, station, showStationsArea)
	}
	renderTarget(ctx, sv.Space.TargetPosition)
	renderEstimatedTarget(ctx, sv.Space.EstimatedTargetPosition)

	slog.Info("Space View Rendered")

	return image
}

func renderStation(ctx *gg.Context, station *ml.StationData, showStationArea bool) {
	ctx.SetRGB255(0, 0, 0)
	ctx.DrawCircle(station.Position.X, station.Position.Y, STATION_RADIUS)
	ctx.Fill()

	if showStationArea {
		ctx.SetRGBA255(0, 0, 0, 150)
		ctx.DrawCircle(station.Position.X, station.Position.Y, station.DistToTarget)
		ctx.Stroke()
	}
}

func renderTarget(ctx *gg.Context, target ml.Position) {
	ctx.SetRGB255(0, 0, 255)
	ctx.DrawCircle(target.X, target.Y, STATION_RADIUS)
	ctx.Fill()
}

func renderEstimatedTarget(ctx *gg.Context, position ml.Position) {
	ctx.SetRGB255(255, 0, 0)
	ctx.SetLineWidth(3)
	ctx.DrawLine(position.X-10, position.Y-10, position.X+10, position.Y+10)
	ctx.DrawLine(position.X-10, position.Y+10, position.X+10, position.Y-10)

	ctx.Stroke()
}
