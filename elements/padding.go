package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Padding struct {
	Left   Length
	Right  Length
	Top    Length
	Bottom Length
	Child  Box
	parentData
}

func (b *Padding) GetSize(c Constraints) Size {
	padX := b.Left + b.Right
	padY := b.Top + b.Bottom
	if b.Child == nil {
		return NewSize(padX, padY).GrowToSatisfy(c)
	}
	size := b.Child.GetSize(c.Shrink(padX, padY))
	return size.Grow(padX, padY).MustSatisfy(c)
}

func (b *Padding) Draw(g *sdl.Renderer, r Rect) {
	if b.Child != nil {
		b.Child.Draw(g, r.Shrink(b.Left, b.Right, b.Top, b.Bottom))
	}
}
