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

func (b *Text) Draw(g *sdl.Renderer, r Rect) {
	defer b.surface.Free()
	tex, err := g.CreateTextureFromSurface(b.surface)
	if err != nil {
		panic(err)
	}
	if err := g.CopyF(tex, nil, r.ToSdl()); err != nil {
		panic(err)
	}
}
