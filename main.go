package main

// import "github.com/yohamta/donburi"

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WinW  = 600
	WinH  = 450
	TITLE = "Mobilizia"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Mobilizia!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WinW, WinH
}

func main() {

	ebiten.SetWindowSize(WinW, WinH)
	ebiten.SetWindowTitle(TITLE)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
