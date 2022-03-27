package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Grow struct {
	Axes  Axes
	Child Box
	parentData
}

func (b *Grow) GetSize(c Constraints) Size {
	c = c.ShrinkUnboundedToZero()
	switch b.Axes {
	case AxesX:
		c = c.TightenWidthToMax()
	case AxesY:
		c = c.TightenHeightToMax()
	case AxesXY:
		c = c.TightenToMax()
	}
	if b.Child == nil {
		return c.Biggest()
	}
	return b.Child.GetSize(c)
}

func (b *Grow) Draw(surface *sdl.Surface, r Rect) {
	if b.Child != nil {
		b.Child.Draw(surface, r)
	}
}
