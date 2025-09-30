package resource

type Management interface {
	// AddResources adds the resource to the game,
	// with given fields and name.
	// Returns a error if the resource already exists.
	AddResource(name string, fields map[string]any) error

	// RemoveResources removers the resource with name.
	// Returns a error if the resource doesn't exist.
	RemoveResource(name string) error
}
