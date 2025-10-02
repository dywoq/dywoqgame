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

	// Delete deletes a resource from the resource management m.
	Delete(m Management) error

	// Fields returns the fields of the resource
	// that is present in the registry of the resource manager.
	Fields() map[string]any

	// Kind returns a kind of the resource.
	Kind() Kind
}

// implementation of Resource interface
type iResource struct {
	name   string
	kind   Kind
	fields map[string]any
}

func (i *iResource) Name() string {
	return i.name
}

func (i *iResource) Load(m Management) error {
	return m.AddResource(i.name, i.kind, i.fields)
}

func (i *iResource) Delete(m Management) error {
	return m.DeleteResource(i.name)
}

func (i *iResource) Fields() map[string]any {
	return i.fields
}

func (i *iResource) Kind() Kind {
	return i.kind
}

// New creates a new resource with given name, kind and fields.
func New(name string, kind Kind, fields map[string]any) Resource {
	return &iResource{name: name, kind: kind, fields: fields}
}
