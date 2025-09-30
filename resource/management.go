package resource

type Management interface {
	// AddResource adds the resource named named with given kind,
	// and fields.
	AddResource(name string, kind Kind, fields map[string]any) error

	// DeleteResource deletes the resource named name.
	DeleteResource(name string) error
}
