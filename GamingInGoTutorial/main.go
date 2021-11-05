package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800

	targetTicksPerSecond = 60
)

var delta float64

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"GoGaming",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing Window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing rendere:", err)
		return
	}
	defer renderer.Destroy()

	plr := newPlayer(renderer)
	elements = append(elements, plr)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2)
			y := float64(j) * basicEnemySize
			enemy := newBasicEnemy(renderer, vector{x, y})
			elements = append(elements, enemy)
		}
	}

	initBulletPool(renderer)

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				if err != nil {
					fmt.Println("updating element", err)
					return
				}
				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element", err)
					return
				}
			}
		}

		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions:", err)
			return
		}

		renderer.Present()

		delta = float64(time.Since(frameStartTime).Seconds()) * targetTicksPerSecond
	}

}
