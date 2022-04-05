package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Center struct {
	Child Box
	parentData
}

func (b *Center) GetSize(c Constraints) Size {
	if b.Child == nil {
		return c.Smallest()
	}
	size := b.Child.GetSize(c.NoMin())
	b.Child.SetParentData(size)
	return c.ShrinkUnbounded(size).Biggest()
}

func (b *Center) Draw(g *sdl.Renderer, r Rect) {
	if b.Child != nil {
		size := b.Child.GetParentData().(Size)
		b.Child.Draw(g, size.PlaceAtCenter(r))
	}
}
