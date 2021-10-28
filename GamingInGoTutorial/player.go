package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex  *sdl.Texture
	x, y float64

	lastShot time.Time
}

const (
	playerSpeed        = 0.2
	playerSize         = 105
	playerShotCooldown = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "sprites/player.bmp")
	p.x = screenWidth / 2
	p.y = screenHeight - playerSize/2

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2
	y := p.y - playerSize/2
	renderer.Copy(p.tex, &sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize}, &sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x-(playerSize/2) > 0 {
			p.x -= playerSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x+(playerSize/2) < screenWidth {
			p.x += playerSpeed
		}
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown {
			p.shoot(p.x+25, p.y-20)
			p.shoot(p.x-25, p.y-20)
			p.lastShot = time.Now()
		}
	}
}

func (p *player) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.x = x
		bul.y = y
		bul.angle = 270 * (math.Pi / 180)
	}
}
