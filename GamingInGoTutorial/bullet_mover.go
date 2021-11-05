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
	mover.container.position.x += bulletSpeed * math.Cos(mover.container.rotation) * delta
	mover.container.position.y += bulletSpeed * math.Sin(mover.container.rotation) * delta
	if mover.container.position.x > screenWidth || mover.container.position.x < 0 || mover.container.position.y < 0 || mover.container.position.y > screenHeight {
		mover.container.active = false
	}

	mover.container.collisions[0].center = mover.container.position

	return nil
}

func (mover *bulletMover) onCollision(other *element) error {
	mover.container.active = false
	return nil
}
