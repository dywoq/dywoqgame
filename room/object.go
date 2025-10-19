package room

import (
	"github.com/dywoq/dywoqgame/resource"
)

type Object struct {
	name        string
	hasSprite   bool
	spriteTable resource.Table
}

// NewObject returns a pointer to Object.
// Returns
//
//	resource.ErrNotFound(sprite)
//
// if there's no resource called sprite.
//
// If the resource exists, but its kind is not what expected (expected is resource.Sprite),
// returns
//
//	resource.ErrExpected(sprite, resource.Sprite, kind).
func NewObject(rm resource.Management, name, sprite string) (*Object, error) {
	if !rm.Exists(sprite) {
		return nil, resource.ErrNotFound(sprite)
	}
	if kind, _ := rm.KindOf(sprite); kind != resource.Sprite {
		return nil, resource.ErrExpected(sprite, resource.Sprite, kind)
	}
	table, err := rm.TableOf(sprite)
	if err != nil {
		return nil, err
	}
	o := &Object{name, true, table}
	return o, nil
}

// NewObjectNoSprite returns a pointer to Object.
// The only difference from NewObject that
// you can load sprite into the object later with Object.Load.
func NewObjectNoSprite(name string) (*Object, error) {
	return &Object{name, false, resource.Table{}}, nil
}

// LoadSprite loads the sprite into the object.
//
// Returns
//
//	resource.ErrNotFound(sprite)
//
// if there's no resource called sprite.
//
// If the resource exists, but its kind is not what expected (expected is resource.Sprite),
// returns
//
//	resource.ErrExpected(sprite, resource.Sprite, kind)
//
// Does nothing and returns nil if object already has a sprite.
func (o *Object) LoadSprite(rm resource.Management, sprite string) error {
	if o.hasSprite {
		return nil
	}
	if !rm.Exists(sprite) {
		return resource.ErrNotFound(sprite)
	}
	if k, _ := rm.KindOf(sprite); k != resource.Sprite {
		return resource.ErrExpected(sprite, resource.Sprite, k)
	}
	t, err := rm.TableOf(sprite)
	if err != nil {
		return err
	}
	o.spriteTable = t
	o.hasSprite = true
	return nil
}

// Name returns the name of the object.
func (o *Object) Name() string { return o.name }

// HasSprite returns true if the object has a sprite.
func (o *Object) HasSprite() bool { return o.hasSprite }

// Sprite returns the metadata table of the sprite.
func (o *Object) Sprite() resource.Table { return o.spriteTable }

// Kind returns the kind of the resource,
// specifically resource.Object.
func (o *Object) Kind() resource.Kind { return resource.Object }

// Table returns the metadata of the object,
// including
func (o *Object) Table() resource.Table {
	t := resource.Table{
		"name":       o.name,
		"has_sprite": o.hasSprite,
		"sprite":     o.spriteTable,
	}
	return t
}
