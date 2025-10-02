package resource

// Kind is a resource kind.
type Kind string

type Manager interface {
	// AddResource loads a resource into the resource manager.
	AddResource(name string, kind Resource, fields map[string]any) error

	// UnloadResource unloads a resource from the resource manager.
	UnloadResource(name string) error
}

type Resource interface {
	// Kind returns a resource kind.
	Kind() Kind

	// Name returns a name of resource.
	Name() string

	// Fields returns a fields of resource,
	// including the custom and builtin ones.
	Fields() map[string]any

	// Load loads resource into the resource manager m.
	Load(m Manager) error

	// Unload unloads resource from the resource manager m.
	Unload(m Manager) error
}

const (
	Object Kind = "object"
	Room   Kind = "room"
)

// Correct reports whether k is a valid resource kind.
func (k Kind) Correct() bool {
	switch k {
	case Object:
		return true
	case Room:
		return true
	}
	return false
}
