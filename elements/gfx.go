package elements

import (
	"errors"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func DrawRoundedRectangle(g *sdl.Renderer, x1, y1, w, h, radius float32, red, green, blue, alpha uint8) error {
	defer elapsed("DrawRoundedRectangle")()
	x2 := x1 + w
	y2 := y1 + h

	left := startToIntInclusive(x1)
	leftFull := startToInt(x1)
	leftRadiusFull := startToInt(x1 + radius)
	right := endToIntInclusive(x2)
	rightFull := endToInt(x2)
	rightRadiusFull := endToInt(x2 - radius)
	top := startToIntInclusive(y1)
	topFull := startToInt(y1)
	topRadiusFull := startToInt(y1 + radius)
	bottom := endToIntInclusive(y2)
	bottomFull := endToInt(y2)
	bottomRadiusFull := endToInt(y2 - radius)

	if err := g.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		return err
	}
	if err := g.SetDrawColor(red, green, blue, alpha); err != nil {
		return err
	}
	if err := g.FillRect(&sdl.Rect{X: leftFull, Y: topRadiusFull, W: rightFull - leftFull, H: bottomRadiusFull - topRadiusFull}); err != nil {
		return err
	}
	if err := g.FillRect(&sdl.Rect{X: leftRadiusFull, Y: topFull, W: rightRadiusFull - leftRadiusFull, H: bottomFull - topFull}); err != nil {
		return err
	}

	rr := radius * radius
	aa := 4
	aaf := float32(aa)
	aafSq := float32(aa * aa)

	for y := top; y <= bottom; y++ {
		if y == topRadiusFull {
			y = bottomRadiusFull
		}
		for x := left; x <= right; x++ {
			if x == leftRadiusFull {
				x = rightRadiusFull
			}
			hits := 0
			for _y := 0; _y < aa; _y++ {
				yf := float32(y) + float32(_y)/aaf
				for _x := 0; _x < aa; _x++ {
					xf := float32(x) + float32(_x)/aaf
					if x1 <= xf && xf < x2 && y1+radius <= yf && yf < y2-radius {
						hits++
						continue
					} else if x1+radius <= xf && xf < x2-radius && y1 <= yf && yf < y2 {
						hits++
						continue
					}
					xLeft := xf - (x1 + radius)
					yTop := yf - (y1 + radius)
					xRight := xf - (x2 - radius)
					yBottom := yf - (y2 - radius)
					if (xLeft*xLeft+yTop*yTop) < rr ||
						(xRight*xRight+yTop*yTop) < rr ||
						(xRight*xRight+yBottom*yBottom) < rr ||
						(xLeft*xLeft+yBottom*yBottom) < rr {
						hits++
					}
				}
			}
			t := float32(hits) / aafSq
			if err := g.SetDrawColor(uint8(float32(red)*t), uint8(float32(green)*t), uint8(float32(blue)*t), uint8(float32(alpha)*t)); err != nil {
				return err
			}
			if err := g.DrawPoint(x, y); err != nil {
				return err
			}
		}
	}
	return nil
}

// Adapted from: https://github.com/rtrussell/BBCSDL/blob/master/src/SDL2_gfxPrimitives.c#L4668

func AAFilledEllipseRGBA(renderer *sdl.Renderer, cx, cy, rx, ry float32, r, g, b, a uint8) error {
	var n, xi, yi int
	var s, v, x, y, dx, dy float64
	if rx <= 0 || ry <= 0 {
		return errors.New("radius cannot be negative")
	}
	if err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		return err
	}
	if rx >= ry {
		n = int(ry + 1)
		for yi = int(cy - float32(n) - 1); yi <= int(cy+float32(n)+1); yi++ {
			if float32(yi) < cy-0.5 {
				y = float64(yi)
			} else {
				y = float64(yi + 1)
			}
			s = (y - float64(cy)) / float64(ry)
			s *= s
			x = 0.5
			if s < 1 {
				x = float64(rx) * math.Sqrt(1-s)
				if x >= 0.5 {
					if err := renderer.SetDrawColor(r, g, b, a); err != nil {
						return err
					}
					if err := renderer.DrawLine(int32(float64(cx)-x+1), int32(yi), int32(float64(cx)+x-1), int32(yi)); err != nil {
						return err
					}
				}
			}
			s = 8 * float64(ry) * float64(ry)
			dy = math.Abs(y-float64(cy)) - 1
			xi = int(float64(cx) - x)
			for {
				dx = float64((cx - float32(xi) - 1) * ry / rx)
				v = s - 4*(dx-dy)*(dx-dy)
				if v < 0 {
					break
				}
				v = (math.Sqrt(v) - 2*(dx+dy)) / 4
				if v < 0 {
					break
				}
				if v > 1 {
					v = 1
				}
				if err := renderer.SetDrawColor(r, g, b, uint8(float64(a)*v)); err != nil {
					return err
				}
				if err := renderer.DrawPoint(int32(xi), int32(yi)); err != nil {
					return err
				}
				xi--
			}
			xi = int(float64(cx) + x)
			for {
				dx = float64((float32(xi) - cx) * ry / rx)
				v = s - 4*(dx-dy)*(dx-dy)
				if v < 0 {
					break
				}
				v = (math.Sqrt(v) - 2*(dx+dy)) / 4
				if v < 0 {
					break
				}
				if v > 1 {
					v = 1
				}
				if err := renderer.SetDrawColor(r, g, b, uint8(float64(a)*v)); err != nil {
					return err
				}
				if err := renderer.DrawPoint(int32(xi), int32(yi)); err != nil {
					return err
				}
				xi++
			}
		}
	} else {
		n = int(rx + 1)
		for xi = int(cx - float32(n) - 1); xi <= int(cx+float32(n)+1); xi++ {
			if float32(xi) < cx-0.5 {
				x = float64(xi)
			} else {
				x = float64(xi + 1)
			}
			s = (x - float64(cx)) / float64(rx)
			s *= s
			x = 0.5
			if s < 1 {
				y = float64(ry) * math.Sqrt(1-s)
				if y >= 0.5 {
					if err := renderer.SetDrawColor(r, g, b, a); err != nil {
						return err
					}
					if err := renderer.DrawLine(int32(xi), int32(float64(cy)-y+1), int32(xi), int32(float64(cy)+y-1)); err != nil {
						return err
					}
				}
			}
			s = 8 * float64(rx) * float64(rx)
			dx = math.Abs(x-float64(cx)) - 1
			yi = int(float64(cy) - y)
			for {
				dy = float64((cy - float32(yi) - 1) * rx / ry)
				v = s - 4*(dy-dx)*(dy-dx)
				if v < 0 {
					break
				}
				v = (math.Sqrt(v) - 2*(dy+dx)) / 4
				if v < 0 {
					break
				}
				if v > 1 {
					v = 1
				}
				if err := renderer.SetDrawColor(r, g, b, uint8(float64(a)*v)); err != nil {
					return err
				}
				if err := renderer.DrawPoint(int32(xi), int32(yi)); err != nil {
					return err
				}
				yi--
			}
			yi = int(float64(cy) + y)
			for {
				dy = float64((float32(yi) - cy) * rx / ry)
				v = s - 4*(dy-dx)*(dy-dx)
				if v < 0 {
					break
				}
				v = (math.Sqrt(v) - 2*(dy+dx)) / 4
				if v < 0 {
					break
				}
				if v > 1 {
					v = 1
				}
				if err := renderer.SetDrawColor(r, g, b, uint8(float64(a)*v)); err != nil {
					return err
				}
				if err := renderer.DrawPoint(int32(xi), int32(yi)); err != nil {
					return err
				}
				yi++
			}
		}
	}
	return nil
}

func startToInt(x float32) int32 {
	if x < 0 {
		return int32(x)
	}
	return int32(math.Ceil(float64(x)))
}

func startToIntInclusive(x float32) int32 {
	if x < 0 {
		return int32(math.Floor(float64(x)))
	}
	return int32(x)
}

func endToInt(x float32) int32 {
	if x < 0 {
		return int32(math.Floor(float64(x)))
	}
	return int32(x)
}

func endToIntInclusive(x float32) int32 {
	if x < 0 {
		return int32(x)
	}
	return int32(math.Ceil(float64(x)))
}
