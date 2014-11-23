package main

import (
	"fmt"
	"github.com/EvanTheB/SDLGame/tools"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

const (
	W     int     = 400
	H     int     = 400
	Scale float64 = 5E11 / 400
)

var renderer *sdl.Renderer
var window *sdl.Window

func main() {
	window, renderer = tools.Start("planets", int(W), int(H))
	defer window.Destroy()
	defer renderer.Destroy()

	sun := tools.Body{1.988E30, 6.96E8, tools.Vector{0, 0}, tools.Vector{0, 0}}
	earth := tools.Body{5.972E24, 6.371E6, tools.Vector{1.5E11, 0}, tools.Vector{0, 2.9E4}}
	mars := tools.Body{6.4185e23, 6.371E6, tools.Vector{2.25e11, 0}, tools.Vector{0, 2.41e4}}
	ecc := tools.Body{1e20, 1e5, tools.Vector{2e11, 0}, tools.Vector{0, 3e4}}
	bodies := []*tools.Body{&sun, &earth, &mars, &ecc}

	fmt.Println(earth.Position.Dist(sun.Position))
	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	renderer.Clear()
	quit := false
	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}

		renderer.SetDrawColor(0, 0, 0, 0xFF)

		tools.UpdateBodiesSeconds(bodies, 60*60*24*7)
		DrawBodies(bodies)

		renderer.Present()
		//fmt.Println(earth.Position, sun.Position)
	}
}

func DrawBodies(bodies []*tools.Body) {
	for _, v := range bodies {
		if renderer.DrawPoint(int(v.Position.X/Scale)+W/2, int(v.Position.Y/Scale)+H/2) != nil {
			fmt.Fprintf(os.Stderr, "failed to draw point:", sdl.GetError())
		}
	}
}
