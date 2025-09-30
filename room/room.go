package room

import "github.com/dywoq/dywoqgame/resource"

type Room struct {
	name    string
	w, h    int
	objects []*Object
}

// NewRoom returns a new pointer to Room with the given name,
// width, height and initial objects.
func NewRoom(name string, w, h int, objects []*Object) *Room {
	return &Room{name: name, objects: objects, w: w, h: h}
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
