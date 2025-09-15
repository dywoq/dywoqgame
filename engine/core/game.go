package core

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Window *Window
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Loop() error {
	e := ebitenGameWrapper{g}
	return ebiten.RunGame(&e)
}

type ebitenGameWrapper struct {
	game *Game
}

func (e *ebitenGameWrapper) Update() error {
	return nil
}

func (e *ebitenGameWrapper) Draw(screen *ebiten.Image) {
	// empty
}

func (e *ebitenGameWrapper) Layout(int, int) (int, int) {
	return e.game.Window.Width, e.game.Window.Height
}
