// Package room provides tools for managing and switching rooms (scenes) in an Ebiten game.
package room

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Management is a interface that allows you to manage scenes,
// and rooms.
type Management interface {
	// Register registers a new room into the manager.
	// It does nothing if the room already exists with the same name.
	Register(r Changeable)

	// Set sets the current room to room.
	// Returns an error if room wasn't found.
	Set(room string) error

	// Rooms returns a map of stored rooms.
	Rooms() map[string]Changeable

	// Update calls the room's Update method.
	Update() error

	// Draw calls the room's Draw method.
	Draw(screen *ebiten.Image)
}
