package elements

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Column struct {
	Alignment  Alignment
	Gap        float32
	Children   []Box
	parentData interface{}
}

func (b *Column) GetSize(c Constraints) Size {
	mostWidth := float32(0)
	usedHeight := float32(len(b.Children)-1) * b.Gap
	for _, child := range b.Children {
		size := child.GetSize(NewConstraints(0, 0, c.Max.Width, c.Max.Height-usedHeight))
		child.SetParentData(size)
		mostWidth = max(mostWidth, size.Width)
		usedHeight += size.Height
	}
	if usedHeight > c.Max.Height {
		fmt.Println("overflow in column")
	}
	return Size{
		max(c.Min.Width, min(c.Max.Width, mostWidth)),
		max(c.Min.Height, min(c.Max.Height, usedHeight)),
	}
}

func (b *Column) Draw(surface *sdl.Surface, r Rect) {
	for _, child := range b.Children {
		size := child.GetParentData().(Size)
		addX := float32(0)
		switch b.Alignment {
		case StartAlign:
		case CenterAlign:
			addX = r.Width/2 - size.Width/2
		case EndAlign:
			addX = r.Width - size.Width
		}
		child.Draw(surface, Rect{r.X + addX, r.Y, size})
		r.Y += size.Height + b.Gap
	}
}

func (b *Column) GetParentData() interface{} {
	return b.parentData
}

func (b *Column) SetParentData(data interface{}) {
	b.parentData = data
}
