package room

import (
	"errors"

	"github.com/dywoq/dywoqgame/graphics"
	"github.com/dywoq/dywoqgame/resource"
)

type Object struct {
	name   string
	sprite *graphics.Sprite
}

// NewObject returns a new pointer to Object with the given name and sprite.
// sprite must be not nil, otherwise, it returns a error.
func NewObject(name string, sprite *graphics.Sprite) (*Object, error) {
	if sprite == nil {
		return nil, errors.New("pointer to graphics.Sprite is nil")
	}
	o := &Object{
		name:   name,
		sprite: sprite,
	}
	return o, nil
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
