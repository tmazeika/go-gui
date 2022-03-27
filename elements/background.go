package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Background struct {
	Color uint32
	Child Box
	parentData
}

func (b *Background) GetSize(c Constraints) Size {
	if b.Child == nil {
		return c.Smallest()
	}
	return b.Child.GetSize(c)
}

func (b *Background) Draw(surface *sdl.Surface, r Rect) {
	if err := surface.FillRect(r.ToSdl(), b.Color); err != nil {
		panic(err)
	}
	if b.Child != nil {
		b.Child.Draw(surface, r)
	}
}
