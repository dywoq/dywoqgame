package graphics

import (
	"image"
	"io/fs"

	"github.com/dywoq/dywoqgame/resource"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	name   string
	path   string
	width  int
	height int

	raw *ebiten.Image
}

var imagesCache = map[string]*ebiten.Image{}

// NewSprite returns a new pointer to Sprite structure, with the given fs.FS,
// path and name.
// The file from path must be image.
func NewSprite(name string, path string, fs fs.FS) (*Sprite, error) {
	if raw, ok := imagesCache[path]; ok {
		return &Sprite{name, path, raw.Bounds().Dx(), raw.Bounds().Dy(), raw}, nil
	}
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	raw := ebiten.NewImageFromImage(img)
	imagesCache[path] = raw
	return &Sprite{name, path, raw.Bounds().Dx(), raw.Bounds().Dy(), raw}, nil
}

// Name returns the name of the sprite resource.
func (s *Sprite) Name() string { return s.name }

// Path returns the path to the sprite.
func (s *Sprite) Path() string { return s.path }

// Width returns the width of the sprite.
func (s *Sprite) Width() int { return s.width }

// Height returns the height of the sprite.
func (s *Sprite) Height() int { return s.height }

// Table returns the table with the sprite data,
// including name, path, width, height and raw data.
func (s *Sprite) Table() map[string]any {
	table := map[string]any{
		"name":   s.name,
		"path":   s.path,
		"width":  s.width,
		"height": s.height,
		"raw":    s.raw,
	}
	return table
}

// Kind returns the kind of the resource,
// specifically resource.Sprite.
func (s *Sprite) Kind() resource.Kind {
	return resource.Sprite
}
