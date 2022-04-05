package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ClipBox struct {
	Radius float32
	Child  Box
	parentData
}

func (b *ClipBox) GetSize(c Constraints) Size {
	if b.Child == nil {
		return c.Smallest()
	}
	return b.Child.GetSize(c)
}

func (b *ClipBox) Draw(g *sdl.Renderer, r Rect) {
	w, h, err := g.GetOutputSize()
	if err != nil {
		panic(err)
	}
	fgTex, err := g.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		panic(err)
	}
	maskTex, err := g.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		panic(err)
	}
	bgTex := g.GetRenderTarget()
	if err := g.SetRenderTarget(fgTex); err != nil {
		panic(err)
	}
	if b.Child != nil {
		b.Child.Draw(g, r)
	}
	if err := g.SetRenderTarget(maskTex); err != nil {
		panic(err)
	}
	if err := DrawRoundedRectangle(g, float32(r.Left), float32(r.Top), float32(r.Width), float32(r.Height), b.Radius, 255, 255, 255, 255); err != nil {
		panic(err)
	}
	if err := g.SetRenderTarget(bgTex); err != nil {
		panic(err)
	}
	if err := g.Copy(maskTex, nil, nil); err != nil {
		panic(err)
	}

	// if err := gfx.FilledEllipseRGBA(g, int32(float32(r.Left)+b.Radius), int32(float32(r.Top)+b.Radius), int32(b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.AACircleRGBA(g, int32(float32(r.Left)+b.Radius), int32(float32(r.Top)+b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.FilledEllipseRGBA(g, int32(float32(r.Right)-b.Radius), int32(float32(r.Top)+b.Radius), int32(b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.AACircleRGBA(g, int32(float32(r.Right)-b.Radius), int32(float32(r.Top)+b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.FilledEllipseRGBA(g, int32(float32(r.Right)-b.Radius), int32(float32(r.Bottom)-b.Radius), int32(b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.AACircleRGBA(g, int32(float32(r.Right)-b.Radius), int32(float32(r.Bottom)-b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.FilledEllipseRGBA(g, int32(float32(r.Left)+b.Radius), int32(float32(r.Bottom)-b.Radius), int32(b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if err := gfx.AACircleRGBA(g, int32(float32(r.Left)+b.Radius), int32(float32(r.Bottom)-b.Radius), int32(b.Radius), 255, 255, 255, 255); !err {
	// 	panic(err)
	// }
	// if b := gfx.FilledPolygonRGBA(g, []int16{int16(float32(r.Left) + b.Radius), int16(float32(r.Right) - b.Radius), int16(float32(r.Right) - b.Radius), int16(float32(r.Left) + b.Radius)}, []int16{int16(r.Top), int16(r.Top), int16(r.Bottom), int16(r.Bottom)}, 255, 255, 255, 255); !b {
	// 	panic(b)
	// }
	// if b := gfx.FilledPolygonRGBA(g, []int16{int16(r.Left), int16(r.Right), int16(r.Right), int16(r.Left)}, []int16{int16(float32(r.Top) + b.Radius), int16(float32(r.Top) + b.Radius), int16(float32(r.Bottom) - b.Radius), int16(float32(r.Bottom) - b.Radius)}, 255, 255, 255, 255); !b {
	// 	panic(b)
	// }

	// resTex, err := g.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, w, h)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := resTex.SetBlendMode(sdl.BLENDMODE_BLEND); err != nil {
	// 	panic(err)
	// }
	// if err := g.SetRenderTarget(resTex); err != nil {
	// 	panic(err)
	// }
	// if err := maskTex.SetBlendMode(sdl.BLENDMODE_MOD); err != nil {
	// 	panic(err)
	// }
	// if err := fgTex.SetBlendMode(sdl.BLENDMODE_NONE); err != nil {
	// 	panic(err)
	// }
	// if err := g.Copy(fgTex, nil, nil); err != nil {
	// 	panic(err)
	// }
	// if err := g.Copy(maskTex, nil, nil); err != nil {
	// 	panic(err)
	// }
	// if err := g.SetRenderTarget(bgTex); err != nil {
	// 	panic(err)
	// }
	// if err := g.Copy(resTex, nil, nil); err != nil {
	// 	panic(err)
	// }
}
