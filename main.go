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

	window, err := sdl.CreateWindow("go-gui", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	tree := &elements.Padding{
		Top:    50,
		Right:  50,
		Bottom: 50,
		Left:   50,
		Child: &elements.Background{
			Color: 0xff998833,
			Child: &elements.ClipBox{
				Radius: 70,
				Child: &elements.Background{
					Color: 0xffeb7e75,
					Child: &elements.Grow{
						Axes: elements.AxesXY,
						Child: &elements.Flex{
							Axis:         elements.AxisY,
							Alignment:    elements.AlignStart,
							Distribution: elements.DistributeSpaceBetween,
							Fit:          elements.FlexLoose,
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
								// &elements.FlexItem{
								// 	Factor: 2,
								// 	Child: &elements.Background{
								// 		Color: 0xffffcc00,
								// 		Child: &elements.Sized{
								// 			Width:  100,
								// 			Height: 100,
								// 		},
								// 	},
								// },
								// &elements.FlexItem{
								// 	Factor: 1,
								// 	Child: &elements.Background{
								// 		Color: 0xffff0000,
								// 		Child: &elements.Sized{
								// 			Width:  100,
								// 			Height: 100,
								// 		},
								// 	},
								// },
								&elements.Center{
									Child: &elements.Background{
										Color: 0xff9980f2,
										Child: &elements.Sized{
											Width:  150,
											Height: 150,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	mainSize := tree.GetSize(elements.NewConstraints(0, 0, 800, 600))
	tree.Draw(renderer, elements.NewRect(0, 0, mainSize.Width, mainSize.Height))
	renderer.Present()

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
