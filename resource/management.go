package resource

type Management interface {
	// AddResource adds the resource named named with given kind,
	// and fields.
	AddResource(name string, kind Kind, fields map[string]any) error

	// DeleteResource deletes the resource named name.
	DeleteResource(name string) error

	// Gets the resource with the name,
	// if it wasn't found, it returns nil.
	Get(name string) Resource

	// Has checks if resource manager already has a resource named name.
	Has(name string) bool
}
