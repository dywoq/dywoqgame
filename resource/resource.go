package resource

// Type represents the game engine resource kind.
type Type string

// A resource kind.
const (
	ObjectKind Type = "object"
	AudioKind  Type = "audio"
	SpriteKind Type = "sprite"
	RoomKind   Type = "room"
)

// Resource is the game engine resource.
type Resource interface {
	// Loads resource into the registry of the engine.
	// If resource already exists, it returns a error.
	Load() error

	// Unloads resource from the registry of the engine.
	// If resource is not loaded, it returns a error.
	Unload() error

	// Type returns type of resource.
	Type() Type
}
