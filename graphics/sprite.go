package graphics

import (
	"image"
	"os"

	"github.com/dywoq/dywoqgame/resource"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	name string
	path string
	size Size
	img  *ebiten.Image
}

var imageCache = make(map[string]*ebiten.Image)

// NewSpite returns a pointer to Sprite with given name, path and position.
// Returns a image from the cache.
func NewSprite(name string, path string) (*Sprite, error) {
	if img, ok := imageCache[path]; ok {
		return &Sprite{
			name,
			path,
			Size{img.Bounds().Dx(), img.Bounds().Dy()},
			img,
		}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	nativeImg, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	img := ebiten.NewImageFromImage(nativeImg)
	imageCache[path] = img

	return &Sprite{
		name,
		path,
		Size{img.Bounds().Dx(), img.Bounds().Dy()},
		img,
	}, nil
}

// Loads a sprite into the m resource manager with m.AddResource.
func (s *Sprite) Load(m resource.Management) error {
	fields := map[string]any{
		"path":   s.path,
		"size":   s.size,
		"native": s.img,
	}
	return m.AddResource(s.name, resource.Sprite, fields)
}

// Unloads a sprite from the m resource manager with m.RemoveResource.
func (s *Sprite) Unload(m resource.Management) error {
	return m.RemoveResource(s.name)
}

// Size returns a size of the sprite.
func (s *Sprite) Size() Size {
	return s.size
}

// Path returns a path of the sprite.
func (s *Sprite) Path() string {
	return s.path
}
