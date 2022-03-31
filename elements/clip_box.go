package elements

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
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

func (b *ClipBox) Draw(surface *sdl.Surface, r Rect) {
	cpy, err := surface.Duplicate()
	if err != nil {
		panic(err)
	}
	if b.Child != nil {
		b.Child.Draw(cpy, r)
	}
	for y := int32(r.Top); y < int32(r.Bottom); y++ {
		for x := int32(r.Left); x < int32(r.Right); x++ {
			xLeft, xRight := r.Left+Length(b.Radius), r.Right-Length(b.Radius)
			yTop, yBottom := r.Top+Length(b.Radius), r.Bottom-Length(b.Radius)
			distTL := math.Sqrt(math.Pow(float64(x)-float64(xLeft), 2) + math.Pow(float64(y)-float64(yTop), 2))
			distTR := math.Sqrt(math.Pow(float64(x)-float64(xRight), 2) + math.Pow(float64(y)-float64(yTop), 2))
			distBR := math.Sqrt(math.Pow(float64(x)-float64(xRight), 2) + math.Pow(float64(y)-float64(yBottom), 2))
			distBL := math.Sqrt(math.Pow(float64(x)-float64(xLeft), 2) + math.Pow(float64(y)-float64(yBottom), 2))
			if distTL <= float64(b.Radius) || distTR <= float64(b.Radius) || distBR <= float64(b.Radius) || distBL <= float64(b.Radius) {
				surface.Set(int(x), int(y), cpy.At(int(x), int(y)))
			}
		}
	}
	if err := cpy.Blit(&sdl.Rect{
		X: int32(r.Left + Length(b.Radius)),
		Y: int32(r.Top),
		W: int32(r.Width() - Length(b.Radius*2)),
		H: int32(r.Height()),
	}, surface, r.Shrink(Length(b.Radius), Length(b.Radius*2), 0, 0).ToSdl()); err != nil {
		panic(err)
	}
	if err := cpy.Blit(&sdl.Rect{
		X: int32(r.Left),
		Y: int32(r.Top + Length(b.Radius)),
		W: int32(r.Width()),
		H: int32(r.Height() - Length(b.Radius*2)),
	}, surface, r.Shrink(0, 0, Length(b.Radius), Length(b.Radius*2)).ToSdl()); err != nil {
		panic(err)
	}
	// if err := cpy.Blit(&sdl.Rect{
	// 	X: int32(r.Left),
	// 	Y: int32(r.Top + Length(b.Radius)),
	// 	W: int32(r.Width()),
	// 	H: int32(r.Height() - Length(b.Radius*2)),
	// }, surface, r.ToSdl()); err != nil {
	// 	panic(err)
	// }
}
