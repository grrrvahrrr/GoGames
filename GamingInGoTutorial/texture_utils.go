package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func drawTexture(tex *sdl.Texture, position vector, rotation float64, renderer *sdl.Renderer) error {
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("quering texture: %v", err))
	}

	position.x -= float64(width) / 2
	position.y -= float64(height) / 2

	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)},
		&sdl.Rect{X: int32(position.x), Y: int32(position.y), W: int32(width), H: int32(height)},
		rotation,
		&sdl.Point{X: int32(width) / 2, Y: int32(height) / 2},
		sdl.FLIP_NONE)

}
