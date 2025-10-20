package room

import (
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func ErrNotFound(name string) error {
	if len(name) == 0 {
		return errors.New("room: not found")
	}
	return fmt.Errorf("room: \"%s\" not found")
}
