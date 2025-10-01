package room

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqgame/graphics"
	"github.com/dywoq/dywoqgame/resource"
)

type Object struct {
	name   string
	sprite *graphics.Sprite
}

// NewObject returns a new pointer to Object with the given name and sprite name.
func NewObject(name string, rm resource.Management, sprite string) (*Object, error) {
	if rm.Has(name) {
		return nil, fmt.Errorf("resource named \"%s\" already exists", name)
	}
	s := rm.Get(sprite)
	if s == nil {
		return nil, errors.New("got nil, not resource")
	}
	if s.Kind() != resource.Sprite {
		return nil, fmt.Errorf("expected a sprite, not %s", s.Kind())
	}
	convertedS, ok := s.(*graphics.Sprite)
	if !ok {
		return nil, errors.New("not successful converting s into *graphics.Sprite")
	}
	return &Object{name: name, sprite: convertedS}, nil
}

// Name returns a name of the object.
func (o *Object) Name() string {
	return o.name
}

// Fields returns the fields of the resource
// that is present in the registry of the resource manager.
func (o *Object) Fields() map[string]any {
	m := make(map[string]any)
	m["sprite"] = o.sprite
	return m
}

// Kind returns a kind of object,
// specifically resource.Object.
func (o *Object) Kind() resource.Kind {
	return resource.Object
}

// Load loads the object into the resource manager m using m.AddResource.
func (o *Object) Load(m resource.Management) error {
	return m.AddResource(o.name, resource.Object, o.Fields())
}

// Delete deletes the object from the resource manager m using m.DeleteResource.
func (o *Object) Delete(m resource.Management) error {
	return m.DeleteResource(o.name)
}
