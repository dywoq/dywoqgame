package collision

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// CanCollide defines a type that can collide with another of the same type.
type CanCollide[T any] interface {
	Collides(other *T) bool
}

type Rect struct {
	X      float32 `json:"x,omitempty"`
	Y      float32 `json:"y,omitempty"`
	Width  float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
}

func (r *Rect) Collides(other *Rect) bool {
	if other == nil {
		return false
	}
	return r.X < other.X+other.Width &&
		r.X+r.Width > other.X &&
		r.Y < other.Y+other.Height &&
		r.Y+r.Height > other.Y
}

type Circle struct {
	X      float32 `json:"x,omitempty"`
	Y      float32 `json:"y,omitempty"`
	Radius float32 `json:"radius,omitempty"`
}

func (c *Circle) Collides(other *Circle) bool {
	if other == nil {
		return false
	}
	dx := c.X - other.X
	dy := c.Y - other.Y
	rSum := c.Radius + other.Radius
	return dx*dx+dy*dy <= rSum*rSum
}

type Image struct {
	Image *ebiten.Image `json:"image"`
	X     int           `json:"x"`
	Y     int           `json:"y"`
	mask  [][]bool      `json:"mask"`
}

// GenerateMask converts the image into a bool mask (true = opaque).
func (img *Image) GenerateMask() {
	w, h := img.Image.Bounds().Dx(), img.Image.Bounds().Dy()
	img.mask = make([][]bool, h)
	for y := range h {
		img.mask[y] = make([]bool, w)
		for x := range w {
			_, _, _, a := img.Image.At(x, y).RGBA()
			img.mask[y][x] = a != 0
		}
	}
}

func (img *Image) Collides(other *Image) bool {
	if img == nil || other == nil || img.mask == nil || other.mask == nil {
		return false
	}

	w1, h1 := len(img.mask[0]), len(img.mask)
	w2, h2 := len(other.mask[0]), len(other.mask)

	if img.X+w1 <= other.X || img.X >= other.X+w2 ||
		img.Y+h1 <= other.Y || img.Y >= other.Y+h2 {
		return false
	}

	startX := int(math.Max(float64(img.X), float64(other.X)))
	endX := int(math.Min(float64(img.X+w1), float64(other.X+w2)))
	startY := int(math.Max(float64(img.Y), float64(other.Y)))
	endY := int(math.Min(float64(img.Y+h1), float64(other.Y+h2)))

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			if img.mask[y-img.Y][x-img.X] && other.mask[y-other.Y][x-other.X] {
				return true
			}
		}
	}

	return false
}

func Check[T any](a CanCollide[T], b *T) bool {
	return a.Collides(b)
}
