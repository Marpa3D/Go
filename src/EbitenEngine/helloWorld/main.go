// Первая программа на движке EbitenEngine
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	purpleCol := color.RGBA{208, 0, 255, 255}
	for x := 100; x < 200; x++ {
		for y := 50; y < 200; y++ {
			screen.Set(x, y, purpleCol)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Test program!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
