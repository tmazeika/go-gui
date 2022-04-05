package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type FlexItem struct {
	Factor float32
	Fit    FlexFit
	Child  Box
	parentData
}

func (b *FlexItem) GetSize(c Constraints) Size {
	if b.Child == nil {
		return c.Smallest()
	}
	return b.Child.GetSize(c)
}

func (b *FlexItem) Draw(g *sdl.Renderer, r Rect) {
	if b.Child != nil {
		b.Child.Draw(g, r)
	}
}
