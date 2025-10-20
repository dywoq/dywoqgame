package room

import "github.com/hajimehoshi/ebiten/v2"

// Room is an Ebiten game room, scene and state.
type Room interface {
	Update() error
	Draw(screen *ebiten.Image)
}

// Changeable represents room that can be changed by the room manager.
type Changeable interface {
	Room
	Name() string
}
