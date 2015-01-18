package main

import (
	"fmt"
	"log"

	"github.com/beati/sdl"
)

func main() {
	err := sdl.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := sdl.Init(sdl.InitVideo)
	if err != nil {
		return err
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL simple example", sdl.WindowPosUndefined,
		sdl.WindowPosUndefined, 800, 600, 0)
	if err != nil {
		return err
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RendererPresentVSync)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	for {
		sdl.HandleEvents()
		if !sdl.Running {
			break
		}

		if sdl.Mouse.Left {
			fmt.Println("mouse left")
		}

		if sdl.Keyboard.Up {
			fmt.Println("keyboard up")
		}
		if sdl.Keyboard.X {
			fmt.Println("keyboard x")
		}

		err = renderer.Clear()
		if err != nil {
			return err
		}

		renderer.Present()
	}

	return nil
}
