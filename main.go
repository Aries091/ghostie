// main.go
package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var ghostImage *ebiten.Image

func init() {
	file, err := os.Open("ghost.png")
	if err != nil {
		log.Fatal("Failed to open image file:", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Failed to decode image:", err)
	}
	ghostImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	y         float64
	velocityY float64
}

const (
	gravity   = 0.4
	jumpForce = -9.5
	groundY   = 350
	scale     = 0.5
	ghostX    = 100
)

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.y >= groundY {

		g.velocityY = jumpForce
	}

	g.velocityY += gravity
	g.y += g.velocityY

	if g.y > groundY {
		g.y = groundY
		g.velocityY = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(ghostX, g.y)
	screen.DrawImage(ghostImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ghost Game - File Load")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
