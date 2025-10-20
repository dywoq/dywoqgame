package entry

import (
	"github.com/dywoq/dywoqgame/room"
	"github.com/hajimehoshi/ebiten/v2"
)

// RoomManager is standard room manager.
// Implements room.Management.
type RoomManager struct {
	rooms   map[string]room.Changeable
	current room.Changeable
}

// NewRoomManager returns a new pointer to RoomManager.
func NewRoomManager() *RoomManager {
	return &RoomManager{rooms: make(map[string]room.Changeable)}
}

// Register registers a new room into the manager.
// It does nothing if the room already exists with the same name.
func (r *RoomManager) Register(c room.Changeable) {
	if _, ok := r.rooms[c.Name()]; !ok {
		r.rooms[c.Name()] = c
	}
}

// Set sets the current room by name.
// Returns an error if room wasn't found.
func (r *RoomManager) Set(name string) error {
	if _, ok := r.rooms[name]; !ok {
		return room.ErrNotFound(name)
	}
	r.current = r.rooms[name]
	return nil
}

// Rooms returns a map of stored rooms.
func (r *RoomManager) Rooms() map[string]room.Changeable {
	return r.rooms
}

// Update calls the room's Update method.
func (r *RoomManager) Update() error {
	return r.current.Update()
}

// Draw calls the room's Draw method.
func (r *RoomManager) Draw(screen *ebiten.Image) {
	r.current.Draw(screen)
}
