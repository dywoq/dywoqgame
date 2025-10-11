// Package room provides tools for managing and switching rooms (scenes) in an Ebiten game.
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

// Manager is responsible for changing the room states in the Ebiten engine game.
type Manager struct {
	rooms   map[string]Changeable
	current Changeable
}

// NewManager returns a pointer to Manager structure.
func NewManager() *Manager {
	return &Manager{rooms: make(map[string]Changeable)}
}

// Register registers a new room into the manager.
// Does nothing if r already exists.
func (m *Manager) Register(r Changeable) {
	if _, ok := m.rooms[r.Name()]; !ok {
		m.rooms[r.Name()] = r
	}
}

// Set sets the current room to room.
// Returns an error if room wasn't found.
func (m *Manager) Set(room string) error {
	val, ok := m.rooms[room]
	if !ok {
		return fmt.Errorf("can't find %s room", room)
	}
	m.current = val
	return nil
}

// Update calls the room's Update method.
func (m *Manager) Update() error {
	if m.current == nil {
		return errors.New("current room is nil")
	}
	return m.current.Update()
}

// Draw draws the room's Draw method.
func (m *Manager) Draw(screen *ebiten.Image) {
	if m.current == nil {
		return
	}
	m.current.Draw(screen)
}

// Rooms returns a map of rooms.
func (m *Manager) Rooms() map[string]Changeable {
	return m.rooms
}
