package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Alignment int

const (
	StartAlign Alignment = iota
	CenterAlign
	EndAlign
)

type Size struct {
	Width  float32
	Height float32
}

type Constraints struct {
	Min Size
	Max Size
}

func NewConstraints(minWidth, minHeight, maxWidth, maxHeight float32) Constraints {
	return Constraints{
		Min: Size{
			Width:  minWidth,
			Height: minHeight,
		},
		Max: Size{
			Width:  maxWidth,
			Height: maxHeight,
		},
	}
}

type Rect struct {
	X float32
	Y float32
	Size
}

type Box interface {
	GetSize(c Constraints) Size
	Draw(surface *sdl.Surface, r Rect)
	GetParentData() interface{}
	SetParentData(data interface{})
}

func min(a float32, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a float32, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
