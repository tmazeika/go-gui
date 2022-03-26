package elements

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Row struct {
	Alignment  Alignment
	Gap        float32
	Children   []Box
	parentData interface{}
}

func (b *Row) GetSize(c Constraints) Size {
	usedWidth := float32(len(b.Children)-1) * b.Gap
	mostHeight := float32(0)
	for _, child := range b.Children {
		size := child.GetSize(NewConstraints(0, 0, c.Max.Width-usedWidth, c.Max.Height))
		child.SetParentData(size)
		usedWidth += size.Width
		mostHeight = max(mostHeight, size.Height)
	}
	if usedWidth > c.Max.Width {
		fmt.Println("overflow in row")
	}
	return Size{
		max(c.Min.Width, min(c.Max.Width, usedWidth)),
		max(c.Min.Height, min(c.Max.Height, mostHeight)),
	}
}

func (b *Row) Draw(surface *sdl.Surface, r Rect) {
	for _, child := range b.Children {
		size := child.GetParentData().(Size)
		addY := float32(0)
		switch b.Alignment {
		case StartAlign:
		case CenterAlign:
			addY = r.Height/2 - size.Height/2
		case EndAlign:
			addY = r.Height - size.Height
		}
		child.Draw(surface, Rect{r.X, r.Y + addY, size})
		r.X += size.Width + b.Gap
	}
}

func (b *Row) GetParentData() interface{} {
	return b.parentData
}

func (b *Row) SetParentData(data interface{}) {
	b.parentData = data
}
