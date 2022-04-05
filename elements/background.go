package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Background struct {
	Color
	Child Box
	parentData
}

func (b *Background) GetSize(c Constraints) Size {
	if b.Child == nil {
		return c.Smallest()
	}
	return b.Child.GetSize(c)
}

func (b *Background) Draw(g *sdl.Renderer, r Rect) {
	b.Color.SetDrawColorFor(g)
	if err := g.FillRectF(r.ToSdl()); err != nil {
		panic(err)
	}
	if b.Child != nil {
		b.Child.Draw(g, r)
	}
}
