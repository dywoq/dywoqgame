package graphics

import (
	"fmt"

	"github.com/dywoq/dywoqgame/resource"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var imageCache = make(map[string]*ebiten.Image)

type Sprite struct {
	Size  Size
	Image *ebiten.Image

	name string
	path string
}

// NewSprite returns a pointer to Sprite structure.
// Returns a error if name is already taken.
func NewSprite(name string, rm resource.Management, path string) (*Sprite, error) {
	if rm.Has(name) {
		return nil, fmt.Errorf("resource named %s already exists", name)
	}
	if img, ok := imageCache[path]; ok {
		return &Sprite{Size: Size{
			Width:  float32(img.Bounds().Dx()),
			Height: float32(img.Bounds().Dy()),
		}, Image: img, name: name, path: path}, nil
	}
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	imageCache[path] = img

	return &Sprite{Size: Size{
		Width:  float32(img.Bounds().Dx()),
		Height: float32(img.Bounds().Dy()),
	}, Image: img, name: name, path: path}, nil
}

// Name returns a name of the sprite.
func (s *Sprite) Name() string {
	return s.name
}

// Path returns a path of the sprite.
func (s *Sprite) Path() string {
	return s.path
}

// Fields returns a fields that come with
// the loading into the resource manager.
func (s *Sprite) Fields() map[string]any {
	fields := make(map[string]any)
	fields["size"] = s.Size
	fields["raw"] = s.Image
	return fields
}

// Load loads the sprite into the resource manager m.
func (s *Sprite) Load(m resource.Management) error {
	return m.AddResource(s.name, s.Kind(), s.Fields())
}

// Delete deletes the sprite from the resource manager.
func (s *Sprite) Delete(m resource.Management) error {
	return m.DeleteResource(s.name)
}

// Kind returns a kind of the sprite,
// specifically resource.Sprite.
func (s *Sprite) Kind() resource.Kind {
	return resource.Sprite
}
