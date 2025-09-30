package resource

// Type represents the game engine resource kind.
type Type string

const (
	Object Type = "object"
	Audio  Type = "audio"
	Sprite Type = "sprite"
)

// Resource is the game engine resource.
type Resource interface {
	// Loads resource into the registry of the engine.
	// If resource already exists, it returns a error.
	Load(Management) error

	// Unloads resource from the registry of the engine.
	// If resource is not loaded, it returns a error.
	Unload(Management) error
}
