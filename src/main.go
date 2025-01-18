package main

import (
	"anorth/drawing_app/app"
)

func main() {
	app := app.NewApp(800, 600) // Set up the app with 800x600 resolution.
	app.Run()
}
