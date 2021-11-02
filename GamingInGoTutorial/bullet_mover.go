package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) onUpdate() error {
	mover.container.position.x += bulletSpeed * math.Cos(mover.container.rotation)
	mover.container.position.y += bulletSpeed * math.Sin(mover.container.rotation)
	if mover.container.position.x > screenWidth || mover.container.position.x < 0 || mover.container.position.y < 0 || mover.container.position.y > screenHeight {
		mover.container.active = false
	}
	return nil
}
