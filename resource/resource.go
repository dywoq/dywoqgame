package resource

// Kind is a resource kind.
type Kind string

const (
	Object Kind = "object"
	Room   Kind = "room"
	Sprite Kind = "sprite"
)

type Resource interface {
	// Name returns a name of the resource.
	Name() string

	// Load loads a resource into the resource management m.
	Load(m Management) error

	// Unload unloads a resource from the resource management m.
	Delete(m Management) error

	// Fields returns the fields of the resource
	// that is present in the registry of the resource manager.
	Fields() map[string]any

	// Kind returns a kind of the resource.
	Kind() Kind
}
