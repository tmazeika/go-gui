package elements

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	Content string
	parentData
	surface *sdl.Surface
}

func (b *Text) GetSize(c Constraints) Size {
	font, err := ttf.OpenFont("FiraSans-Regular.ttf", 32)
	if err != nil {
		panic(err)
	}
	defer font.Close()
	if c.Max.Width.IsUnbounded() {
		b.surface, err = font.RenderUTF8Blended(b.Content, sdl.Color{A: 255})
	} else {
		b.surface, err = font.RenderUTF8BlendedWrapped(b.Content, sdl.Color{A: 255}, int(c.Max.Width))
	}
	if err != nil {
		panic(err)
	}
	return NewSize(Length(b.surface.W), Length(b.surface.H)).GrowToSatisfy(c)
}

func (b *Text) Draw(surface *sdl.Surface, r Rect) {
	defer b.surface.Free()
	if err := b.surface.Blit(nil, surface, r.ToSdl()); err != nil {
		panic(err)
	}
}
