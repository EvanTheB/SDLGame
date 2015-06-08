package main

import (
	"fmt"
	"github.com/EvanTheB/SDLGame/tools"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

const (
	SCREEN_W     int     = 1024
	SCREEN_H     int     = 800
	SCREEN_SCALE float64 = 5E11
	SCREEN_RATIO float64 = float64(SCREEN_W) / float64(SCREEN_H)
)

var renderer *sdl.Renderer
var window *sdl.Window
var sprite *sdl.Texture

func main() {
	window, renderer = tools.Start("planets", int(SCREEN_W), int(SCREEN_H))
	wideView := tools.ViewBox{
		-SCREEN_SCALE / 2 / float64(SCREEN_H) * float64(SCREEN_W),
		-SCREEN_SCALE / 2,
		SCREEN_SCALE / float64(SCREEN_H) * float64(SCREEN_W),
		SCREEN_SCALE,
	}
	defer window.Destroy()
	defer renderer.Destroy()
	sprite = tools.LoadTextureTransparent(renderer, "resources/sprites.png", 0, 0xFF, 0xFF)

	bodies := tools.GetPlanets2()
	jup := 0
	gany := 0
	for i, b := range bodies {
		if b.Name == "Jupiter (599)" {
			jup = i
		}
		if b.Name == "Ganymede (503)" {
			gany = i
		}
	}
	quit := false
	curView := &wideView
	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_UP:
					auto := tools.GetAutoView(
						[]tools.Vector{
							bodies[jup].Position,
							bodies[gany].Position,
						},
						SCREEN_RATIO)
					curView = &auto
				case sdl.K_DOWN:
					curView = &wideView
				case sdl.K_q:
					quit = true
				}
				fmt.Println("View:", curView)
			}

		}
		// Draw background
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 0, 0xFF)

		tools.UpdateBodiesSeconds(bodies, 60*60)
		DrawBodies(bodies, *curView)

		renderer.Present()
		//fmt.Println(earth.Position, sun.Position)
	}
}

func pointInView(pos float64, start float64, size float64, pixels int) int32 {
	return int32(float64(pixels) * (pos - start) / size)
}

func DrawBodies(bodies []*tools.Body, view tools.ViewBox) {
	for _, v := range bodies {
		err := renderer.Copy(sprite, &sdl.Rect{0, 0, 100, 100},
			&sdl.Rect{pointInView(v.Position.X, view.X, view.W, SCREEN_W),
				pointInView(v.Position.Y, view.Y, view.H, SCREEN_H),
				20, 20})
		// err := renderer.DrawPoint(pointInView(v.Position.X, view.X, view.W, SCREEN_W),
		// 	pointInView(v.Position.Y, view.Y, view.H, SCREEN_H))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to draw point:", sdl.GetError())
		}
	}
}
