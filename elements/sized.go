package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sized struct {
	Width  Length
	Height Length
	Child  Box
	parentData
}

func (b *Sized) GetSize(c Constraints) Size {
	if b.Child == nil {
		return c.Prefer(b.Width, b.Height)
	}
	return b.Child.GetSize(c.Tighten(b.Width, b.Height))
}

func (b *Sized) Draw(surface *sdl.Surface, r Rect) {
	if b.Child != nil {
		b.Child.Draw(surface, r)
	}
}
