package entry

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqgame/graphics"
	"github.com/dywoq/dywoqgame/resource"
	"github.com/dywoq/dywoqgame/room"
	"github.com/hajimehoshi/ebiten/v2"
)

// Core is a entry point of the game, holding sprites,
// objects and rooms internally.
type Core struct {
	Window Window

	sprites resourcesMap[*graphics.Sprite]
	objects resourcesMap[*room.Object]
	rooms   resourcesMap[*room.Room]
	running bool
	game    *ebitenGame
}

// implementation of ebiten.Game
type ebitenGame struct {
	c *Core
}

type resourceStruct[R resource.Resource] struct {
	fields map[string]any
}

type resourcesMap[R resource.Resource] map[string]resourceStruct[R]

func (r resourcesMap[R]) has(name string) bool {
	_, ok := r[name]
	return ok
}

func (r resourcesMap[R]) get(name string) (resourceStruct[R], bool) {
	res, ok := r[name]
	return res, ok
}

func (r resourcesMap[R]) delete(name string) bool {
	if _, ok := r[name]; !ok {
		return false
	}
	delete(r, name)
	return true
}

// AddResource adds the resource to the game with the given name, kind and fields.
// Must be called before the game will be running, otherwise the error will be returned.
//
// If the game already has the resource named name, if the fields are not facing the requirements of the interface of fields,
// or the game is running, AddResource will return a error.
func (c *Core) AddResource(name string, kind resource.Kind, fields map[string]any) error {
	if c.running {
		return errors.New("the game is running")
	}
	if c.Has(name) {
		return fmt.Errorf("resource \"%s\" already exists", name)
	}
	if !c.fieldsHasExpectedInterface(kind, fields) {
		return errors.New("fields don't have the expected interface")
	}
	switch kind {
	case resource.Object:
		c.objects[name] = resourceStruct[*room.Object]{
			fields: fields,
		}
		return nil
	case resource.Room:
		c.rooms[name] = resourceStruct[*room.Room]{
			fields: fields,
		}
		return nil
	case resource.Sprite:
		c.sprites[name] = resourceStruct[*graphics.Sprite]{
			fields: fields,
		}
		return nil
	}

	return errors.New("wrong resource")
}

// DeleteResource deletes the resource named name.
// Returns a error if the resource wasn't found or the core couldn't delete the game.
func (c *Core) DeleteResource(name string) error {
	if !c.Has(name) {
		return fmt.Errorf("can't find resource \"%s\"", name)
	}
	res := c.Get(name)
	switch res.Kind() {
	case resource.Sprite:
		if ok := c.sprites.delete(name); !ok {
			return fmt.Errorf("couldn't delete sprite \"%s\"", name)
		}
		return nil

	case resource.Object:
		if ok := c.objects.delete(name); !ok {
			return fmt.Errorf("couldn't delete object \"%s\"", name)
		}
		return nil

	case resource.Room:
		if ok := c.rooms.delete(name); !ok {
			return fmt.Errorf("couldn't delete room \"%s\"", name)
		}
		return nil

	default:
		return fmt.Errorf("wrong resource kind: %s", res.Kind())
	}
}

// Get gets the resource with name,
// otherwise, if the resource wasn't found, it returns nil.
func (c *Core) Get(name string) resource.Resource {
	obj, ok := c.objects.get(name)
	if ok {
		return resource.New(name, resource.Object, obj.fields)
	}
	sprite, ok := c.sprites.get(name)
	if ok {
		return resource.New(name, resource.Sprite, sprite.fields)
	}
	room, ok := c.rooms.get(name)
	if ok {
		return resource.New(name, resource.Room, room.fields)
	}
	return nil
}

// Has checks whether the resource named name exists in the game.
func (c *Core) Has(name string) bool {
	switch {
	case c.objects.has(name):
		return true
	case c.sprites.has(name):
		return true
	case c.rooms.has(name):
		return true
	}
	return false
}

func (c *Core) Run() error {
	if c.running {
		return errors.New("the game is already running")
	}
	if c.game == nil {
		return errors.New("internal c.game is nil")
	}
	ebiten.SetWindowTitle(c.Window.Title)
	c.running = true
	return ebiten.RunGame(c.game)
}

func (c *Core) fieldsHasExpectedInterface(kind resource.Kind, fields map[string]any) bool {
	has := func(expectedFields []string, input map[string]any) bool {
		for _, expectedField := range expectedFields {
			if _, ok := fields[expectedField]; !ok {
				return false
			}
		}
		return true
	}

	switch kind {
	case resource.Object:
		expectedFields := []string{"sprite"}
		return has(expectedFields, fields)

	case resource.Sprite:
		expectedFields := []string{"size", "raw"}
		return has(expectedFields, fields)

	case resource.Room:
		expectedFields := []string{"objects", "width", "height"}
		return has(expectedFields, fields)
	}
	return false
}

func (e *ebitenGame) Update() error {
	// there's nothing yet
	return nil
}

func (e *ebitenGame) Draw(screen *ebiten.Image) {
	// there's nothing yet
}

func (e *ebitenGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.c.Window.Width, e.c.Window.Height
}
