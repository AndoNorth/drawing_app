package app

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// App represents the main application.
type App struct {
	width  int
	height int
	canvas *ebiten.Image // Main drawing canvas.
}

func (a *App) Run() {
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}

// NewApp initializes the application.
func NewApp(width, height int) *App {
	canvas := ebiten.NewImage(width, height)
	canvas.Fill(color.White) // Start with a blank white canvas.
	return &App{
		width:  width,
		height: height,
		canvas: canvas,
	}
}

// Update handles game logic updates.
func (a *App) Update() error {
	// TODO: Handle tool interactions, mouse input, and dialogs.
	return nil
}

// Draw renders the game screen.
func (a *App) Draw(screen *ebiten.Image) {
	// Fill the background.
	screen.Fill(color.RGBA{240, 240, 240, 255})

	// Draw the canvas.
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(100, 50) // Properly apply the translation to the GeoM object.
	screen.DrawImage(a.canvas, opts)


	// Placeholder: Draw top menu bar.
	ebitenutil.DebugPrintAt(screen, "[ New | Save | Load ]", 10, 10)

	// Placeholder: Draw side panel.
	ebitenutil.DebugPrintAt(screen, "[ Tools ]", 10, 60)
}

// Layout specifies the screen dimensions.
func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.width, a.height
}
