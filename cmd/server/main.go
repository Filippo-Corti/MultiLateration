package main

import (
	"fmt"
	"image/png"
	"log/slog"
	"net/http"

	"positioning/pkg/graphics"

)

type TemplateData struct {
	ImageSource string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/index.html")
}

func canvasHandler(w http.ResponseWriter, r *http.Request) {
	image := graphics.BuildSimpleLaterationImage()
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
