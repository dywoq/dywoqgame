package room

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqgame/resource"
)

type Room struct {
	name    string
	w, h    int
	objects []*Object
}

// NewRoom returns a new pointer to Room with the given name,
// width, height and initial objects.
// May return a error if any of objects doesn't exist, the name of room is already taken,
// or the objects' kind is not resource.Object.
func NewRoom(name string, rm resource.Management, w, h int, objects []string) (*Room, error) {
	if rm.Has(name) {
		return nil, fmt.Errorf("resource named \"%s\" already exists", name)
	}
	o := []*Object{}
	if len(objects) != 0 {
		for _, obj := range objects {
			current := rm.Get(obj)
			if !rm.Has(obj) {
				return nil, fmt.Errorf("resource named \"%s\" doesn't exist", obj )
			}
			
			if current.Kind() != resource.Object {
				return nil, fmt.Errorf("expected object, not %s", current.Kind())
			}
			
			convertedObj, ok := current.(*Object)
			if !ok {
				return nil, errors.New("didn't convert resource.Resource to room.Object successfully")
			}

			o = append(o, convertedObj)
		}
	}
	return &Room{name: name, objects: o, w: w, h: h}, nil
}

// Name returns a name of the room.
func (r *Room) Name() string {
	return r.name
}

// Load loads the room into the resource manager m.
func (r *Room) Load(m resource.Management) error {
	return m.AddResource(r.name, r.Kind(), r.Fields())
}

// Delete deletes the room from the resource manager m.
func (r *Room) Delete(m resource.Management) error {
	return m.DeleteResource(r.name)
}

// Fields returns the fields of the room
// that is present in the registry of the resource manager.
func (r *Room) Fields() map[string]any {
	m := make(map[string]any)
	m["objects"] = r.objects
	m["width"] = r.w
	m["height"] = r.h
	return m
}

// Kind returns the kind of room,
// specifically resource.Room.
func (r *Room) Kind() resource.Kind {
	return resource.Room
}
