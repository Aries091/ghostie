package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var ghostImageData []byte

var ghostImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(ghostImageData))
	if err != nil {
		log.Fatal(err)
	}
	ghostImage = ebiten.NewImageFromImage(img)
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100) // Position of the ghost
	screen.DrawImage(ghostImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ghost Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
