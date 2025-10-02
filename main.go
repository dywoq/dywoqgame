package main

import (
	"image/color"

	"github.com/dywoq/dywoqgame/room"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type menu struct {
	m                     *room.Manager
	rectWidth, rectHeight float32
}

type home struct{}

func (m *menu) Draw(i *ebiten.Image) {
	vector.DrawFilledRect(i, 640/2, 480/2, m.rectWidth, m.rectHeight, color.White, false)
}

func (m *menu) Update() error {
	m.rectWidth += 2
	m.rectHeight += 2

	if m.rectWidth >= 150 {
		return m.m.Set("home")
	}

	return nil
}

func (m *menu) Name() string {
	return "menu"
}

func (h *home) Draw(i *ebiten.Image) {
	vector.DrawFilledRect(i, 640/2, 480/2, 20, 20, color.White, false)
}

func (h *home) Update() error {
	return nil
}

func (h *home) Name() string {
	return "home"
}

type game struct {
	m *room.Manager
}

func (g *game) Draw(screen *ebiten.Image) {
	g.m.Draw(screen)
}

func (g *game) Update() error {
	return g.m.Update()
}

func (g *game) Layout(int, int) (int, int) {
	return 640, 480
}

func main() {
	m := room.NewManager()
	m.Register(&menu{m, 0, 0})
	m.Register(&home{})

	err := m.Set("menu")
	if err != nil {
		panic(err)
	}

	g := &game{m}
	err = ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
