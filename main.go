package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"gui/elements"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	tree := elements.Padding{
		Top:    50,
		Right:  50,
		Bottom: 50,
		Left:   50,
		Child: &elements.Background{
			Color: 0xffeb7e75,
			Child: &elements.Padding{
				Top:    20,
				Right:  20,
				Bottom: 20,
				Left:   20,
				Child: &elements.Column{
					Alignment: elements.CenterAlign,
					Gap:       20,
					Children: []elements.Box{
						&elements.Background{
							Color: 0xffe8eb6a,
							Child: &elements.Text{
								Content: "Hi!",
							},
						},
						&elements.Background{
							Color: 0xffedc268,
							Child: &elements.Text{
								Content: "Hello, world!",
							},
						},
					},
				},
			},
		},
	}
	mainSize := tree.GetSize(elements.NewConstraints(0, 0, 800, 600))
	tree.Draw(surface, elements.Rect{Size: mainSize})
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				if e.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
			case *sdl.QuitEvent:
				running = false
			}
		}
	}
}
