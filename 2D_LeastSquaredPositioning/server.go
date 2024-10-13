package main

import (
	"fmt"
	"html/template"
	"image"
	"image/png"
	"net/http"

	"github.com/fogleman/gg"
)

type TemplateData struct {
	ImageSource string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		ImageSource: "/image.png",
	}

	template.Execute(w, data)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	image := image.NewRGBA(image.Rect(0, 0, 500, 500))

	ctx := gg.NewContextForRGBA(image)

	ctx.SetRGB255(255, 0, 0)
	ctx.DrawCircle(250, 250, 100)
	ctx.Stroke()

	png.Encode(w, image)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image.png", imageHandler)

	fmt.Println("Server is running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Loading the server", err)
		return
	}

}
