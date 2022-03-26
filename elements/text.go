package elements

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	Content    string
	parentData interface{}
	surface    *sdl.Surface
}

func (b *Text) GetSize(c Constraints) Size {
	font, err := ttf.OpenFont("FiraSans-Regular.ttf", 32)
	if err != nil {
		panic(err)
	}
	defer font.Close()
	surface, err := font.RenderUTF8BlendedWrapped(b.Content, sdl.Color{A: 255}, int(c.Max.Width))
	if err != nil {
		panic(err)
	}
	b.surface = surface
	return Size{
		Width:  max(c.Min.Width, float32(surface.W)),
		Height: max(c.Min.Height, min(c.Max.Height, float32(surface.H))),
	}
}

func (b *Text) Draw(surface *sdl.Surface, r Rect) {
	defer b.surface.Free()
	if err := b.surface.Blit(nil, surface, &sdl.Rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.Width),
		H: int32(r.Height),
	}); err != nil {
		panic(err)
	}
}

func (b *Text) GetParentData() interface{} {
	return b.parentData
}

func (b *Text) SetParentData(data interface{}) {
	b.parentData = data
}
