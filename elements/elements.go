package elements

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Length float32

func (l Length) IsUnbounded() bool {
	return math.IsInf(float64(l), 1)
}

var Unbounded = Length(math.Inf(1))

type Distribution uint8

const (
	DistributeStart Distribution = iota
	DistributeCenter
	DistributeEnd
	DistributeSpaceAround
	DistributeSpaceBetween
	DistributeSpaceEvenly
)

type Axis uint8

const (
	AxisX Axis = iota
	AxisY
)

type Alignment uint8

const (
	AlignCenter Alignment = iota
	AlignStart
	AlignEnd
	AlignStretch
)

type Axes uint

const (
	AxesX Axes = iota
	AxesY
	AxesXY
)

type Size struct {
	Width  Length
	Height Length
}

func NewSize(width, height Length) Size {
	if width.IsUnbounded() || height.IsUnbounded() {
		panic("cannot create unbounded size")
	}
	if width < 0 || height < 0 {
		panic("cannot create size with negative dimensions")
	}
	return Size{width, height}
}

func (s Size) PlaceAtTopLeft(r Rect) Rect {
	if s.Width > r.Width() || s.Height > r.Height() {
		panic("size cannot fit in rect")
	}
	return NewRect(r.Left, r.Left+s.Width, r.Top, r.Top+s.Height)
}

func (s Size) PlaceAtCenter(r Rect) Rect {
	addX := (r.Width() - s.Width) / 2
	addY := (r.Height() - s.Height) / 2
	if addX < 0 || addY < 0 {
		panic("size cannot fit in rect")
	}
	r.Left += addX
	r.Top += addY
	return NewRect(r.Left, r.Left+s.Width, r.Top, r.Top+s.Height)
}

func (s Size) Flip() Size {
	return Size{s.Height, s.Width}
}

func (s Size) Grow(width, height Length) Size {
	return NewSize(s.Width+width, s.Height+height)
}

func (s Size) Minus(width, height Length) Size {
	return s.Grow(-width, -height)
}

func (s Size) GrowToSatisfy(c Constraints) Size {
	if s.Width > c.Max.Width || s.Height > c.Max.Height {
		panic("size is too large for constraints")
	}
	return NewSize(max(c.Min.Width, s.Width), max(c.Min.Height, s.Height))
}

func (s Size) ShrinkToSatisfy(c Constraints) Size {
	if s.Width < c.Min.Width || s.Height < c.Min.Height {
		panic("size is too small for constraints")
	}
	return NewSize(min(c.Max.Width, s.Width), min(c.Max.Height, s.Height))
}

func (s Size) MustSatisfy(c Constraints) Size {
	s.GrowToSatisfy(c)
	s.ShrinkToSatisfy(c)
	return s
}

type Constraints struct {
	Min Size
	Max Size
}

func NewConstraints(minWidth, minHeight, maxWidth, maxHeight Length) Constraints {
	c := Constraints{
		Min: Size{
			Width:  minWidth,
			Height: minHeight,
		},
		Max: Size{
			Width:  maxWidth,
			Height: maxHeight,
		},
	}
	c.validate()
	return c
}

func (c Constraints) TightenToMax() Constraints {
	return NewConstraints(c.Max.Width, c.Max.Height, c.Max.Width, c.Max.Height)
}

func (c Constraints) TightenWidthToMax() Constraints {
	c.Min.Width = c.Max.Width
	c.validate()
	return c
}

func (c Constraints) TightenHeightToMax() Constraints {
	c.Min.Height = c.Max.Height
	c.validate()
	return c
}

func (c Constraints) TightenWidthToMin() Constraints {
	c.Max.Width = c.Min.Width
	c.validate()
	return c
}

func (c Constraints) TightenHeightToMin() Constraints {
	c.Max.Height = c.Min.Height
	c.validate()
	return c
}

func (c Constraints) TightenMainToMax(axis Axis) Constraints {
	if axis == AxisX {
		c.Min.Width = c.Max.Width
	} else {
		c.Min.Height = c.Max.Height
	}
	c.validate()
	return c
}

func (c Constraints) TightenCrossToMax(axis Axis) Constraints {
	if axis == AxisX {
		c.Min.Height = c.Max.Height
	} else {
		c.Min.Width = c.Max.Width
	}
	c.validate()
	return c
}

func (c Constraints) Flip() Constraints {
	return NewConstraints(c.Min.Height, c.Min.Width, c.Max.Height, c.Max.Width)
}

func (c Constraints) validate() {
	if c.Min.Width.IsUnbounded() || c.Min.Height.IsUnbounded() {
		panic("cannot create constraints with unbounded minimums")
	}
	if c.Min.Width < 0 || c.Min.Height < 0 {
		panic("cannot create constraints with negative minimums")
	}
	if c.Max.Width < c.Min.Width || c.Max.Height < c.Min.Height {
		panic("cannot create constraints with maximums that are less than the minimums")
	}
}

func (c Constraints) Contains(width, height Length) bool {
	return width >= c.Min.Width && height >= c.Min.Height &&
		width <= c.Max.Width && height <= c.Max.Height
}

func (c Constraints) MustContain(width, height Length) {
	if !c.Contains(width, height) {
		panic("size breaks out of constraints")
	}
}

func (c Constraints) ShrinkUnbounded(size Size) Constraints {
	if c.Max.Width.IsUnbounded() {
		c.Max.Width = max(c.Min.Width, size.Width)
	}
	if c.Max.Height.IsUnbounded() {
		c.Max.Height = max(c.Min.Height, size.Height)
	}
	return c
}

func (c Constraints) ShrinkUnboundedToZero() Constraints {
	return c.ShrinkUnbounded(NewSize(0, 0))
}

func (c Constraints) Biggest() Size {
	if c.Max.Width.IsUnbounded() {
		c.Max.Width = 0
	}
	if c.Max.Height.IsUnbounded() {
		c.Max.Height = 0
	}
	return NewSize(c.Max.Width, c.Max.Height)
}

func (c Constraints) Smallest() Size {
	return NewSize(c.Min.Width, c.Min.Height)
}

func (c Constraints) Tighten(width, height Length) Constraints {
	size := c.Prefer(width, height)
	return NewConstraints(size.Width, size.Height, size.Width, size.Height)
}

func (c Constraints) Prefer(width, height Length) Size {
	return NewSize(
		max(c.Min.Width, min(c.Max.Width, width)),
		max(c.Min.Height, min(c.Max.Height, height)),
	)
}

func (c Constraints) And(other Constraints) Constraints {
	c.Min.Width = max(c.Min.Width, other.Min.Width)
	c.Min.Height = max(c.Min.Height, other.Min.Height)
	c.Max.Width = min(c.Max.Width, other.Max.Width)
	c.Max.Height = min(c.Max.Height, other.Max.Height)
	c.validate()
	return c
}

func (c Constraints) Shrink(width, height Length) Constraints {
	c.Max.Width -= width
	c.Max.Height -= height
	c.Min.Width = max(0, c.Min.Width-width)
	c.Min.Height = max(0, c.Min.Height-height)
	c.validate()
	return c
}

func (c Constraints) NoMin() Constraints {
	c.Min = Size{}
	return c
}

func (c Constraints) UnboundedWidth() Constraints {
	c.Max.Width = Unbounded
	return c
}

func (c Constraints) UnboundedHeight() Constraints {
	c.Max.Height = Unbounded
	return c
}

type Rect struct {
	Left   Length
	Right  Length
	Top    Length
	Bottom Length
}

func NewRect(left, right, top, bottom Length) Rect {
	if left.IsUnbounded() || right.IsUnbounded() || top.IsUnbounded() || bottom.IsUnbounded() {
		panic("rect cannot have unbounded sides")
	}
	if left > right || top > bottom {
		panic("rect cannot have negative dimensions")
	}
	return Rect{left, right, top, bottom}
}

func (r Rect) Flip() Rect {
	return Rect{r.Top, r.Bottom, r.Left, r.Right}
}

func (r Rect) Width() Length {
	return r.Right - r.Left
}

func (r Rect) Height() Length {
	return r.Bottom - r.Top
}

func (r Rect) Grow(left, right, top, bottom Length) Rect {
	return NewRect(r.Left-left, r.Right+right, r.Top-top, r.Bottom+bottom)
}

func (r Rect) Shrink(left, right, top, bottom Length) Rect {
	return r.Grow(-left, -right, -top, -bottom)
}

func (r Rect) ToSdl() *sdl.Rect {
	return &sdl.Rect{
		X: int32(r.Left),
		Y: int32(r.Top),
		W: int32(r.Right - r.Left),
		H: int32(r.Bottom - r.Top),
	}
}

type Box interface {
	GetSize(c Constraints) Size
	Draw(surface *sdl.Surface, r Rect)
	GetParentData() interface{}
	SetParentData(data interface{})
}

func min(a Length, b Length) Length {
	if a < b {
		return a
	}
	return b
}

func max(a Length, b Length) Length {
	if a > b {
		return a
	}
	return b
}

type parentData struct {
	data interface{}
}

func (p *parentData) GetParentData() interface{} {
	return p.data
}

func (p *parentData) SetParentData(data interface{}) {
	p.data = data
}
