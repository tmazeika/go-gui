package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Padding struct {
	Top        float32
	Right      float32
	Bottom     float32
	Left       float32
	Child      Box
	parentData interface{}
}

func (b *Padding) GetSize(c Constraints) Size {
	maxWidth := c.Max.Width - b.Left - b.Right
	maxHeight := c.Max.Height - b.Top - b.Bottom
	size := b.Child.GetSize(NewConstraints(0, 0, maxWidth, maxHeight))
	b.Child.SetParentData(size)
	return Size{
		Width:  max(c.Min.Width, size.Width+b.Left+b.Right),
		Height: max(c.Min.Height, size.Height+b.Top+b.Bottom),
	}
}

func (b *Padding) Draw(surface *sdl.Surface, r Rect) {
	size := b.Child.GetParentData().(Size)
	b.Child.Draw(surface, Rect{
		X:    r.X + b.Left,
		Y:    r.Y + b.Top,
		Size: size,
	})
}

func (b *Padding) GetParentData() interface{} {
	return b.parentData
}

func (b *Padding) SetParentData(data interface{}) {
	b.parentData = data
}
