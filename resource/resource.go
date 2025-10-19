package resource

import (
	"fmt"
)

type Kind string

type Table map[string]any

type Resource interface {
	// Name returns the name of the resource.
	Name() string

	// Table returns the data of the resource.
	Table() Table

	// Kind returns the kind of the resource.
	Kind() Kind
}

type Management interface {
	// Exists reports whether the resource with name exists.
	Exists(name string) bool

	// TableOf returns the table with data of the resource named name.
	// It should return nil and an error if there's no resource with name exists.
	TableOf(name string) (Table, error)

	// Load loads r into the resource manager.
	// Returns ErrExists if there's already existing resource with the same name.
	Load(r Resource) error

	// Unload unloads r from the resource manager.
	// Returns ErrNotFound if there's no resource with the name.
	Unload(r Resource) error

	// KindOf returns the kind of the resource.
	// It should return an error if there's no resource named name.
	KindOf(name string) (Kind, error)
}

// A resource kind.
const (
	Sprite Kind = "sprite"
	Object Kind = "object"
)

func ErrExists(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("resource: already exists")
	}
	return fmt.Errorf("resource: \"%s\" already exists", name)
}

func ErrNotFound(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("resource: not found")
	}
	return fmt.Errorf("resource: \"%s\" not found", name)
}

func ErrExpected(name string, expected, got Kind) error {
	return fmt.Errorf("resource: \"%s\", expected kind %v, got %v", name, expected, got)
}
