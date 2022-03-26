package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Background struct {
	Color      uint32
	Child      Box
	parentData interface{}
}

func (b *Background) GetSize(c Constraints) Size {
	return b.Child.GetSize(c)
}

func (b *Background) Draw(surface *sdl.Surface, r Rect) {
	if err := surface.FillRect(&sdl.Rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.Width),
		H: int32(r.Height),
	}, b.Color); err != nil {
		panic(err)
	}
	b.Child.Draw(surface, r)
}

func (b *Background) GetParentData() interface{} {
	return b.parentData
}

func (b *Background) SetParentData(data interface{}) {
	b.parentData = data
}
