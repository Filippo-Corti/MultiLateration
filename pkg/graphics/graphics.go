package graphics

import (
	"image"
	"log/slog"

	"github.com/fogleman/gg"
)

func BuildSimpleLaterationImage() *image.RGBA {
	image := image.NewRGBA(image.Rect(0, 0, 1000, 800))
	ctx := gg.NewContextForRGBA(image)

	ctx.SetRGB255(255, 0, 0)
	ctx.DrawCircle(500, 400, 100)
	ctx.Stroke()

	slog.Info("Lateration Image created")

	return image
}
