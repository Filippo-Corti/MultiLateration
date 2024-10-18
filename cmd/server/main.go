package main

import (
	"fmt"
	"image/png"
	"log/slog"
	"net/http"
	"positioning/pkg/controllers"
)

type TemplateData struct {
	ImageSource string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/index.html")
}

func canvasHandler(w http.ResponseWriter, r *http.Request) {

	controller := controllers.NewSpaceController(1000, 800)

	controller.AddStation(200, 150)
	controller.AddStation(800, 300)
	controller.AddStation(600, 450)

	image := controller.RenderView()

	png.Encode(w, image)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ml_canvas", canvasHandler)

	slog.Info("Server is running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Loading the server", err)
		return
	}

}
