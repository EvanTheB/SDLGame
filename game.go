package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/EvanTheB/SDLGame/tools"
    "math"
    "fmt"
    "os"
)



const(
    W int = 800
    H int = 800
    Scale float64 = 5E11/800
)

var renderer *sdl.Renderer
var window *sdl.Window

type Vector struct{
    X, Y float64
}

func (v1 Vector) To(v2 Vector) Vector{
    return Vector{v2.X - v1.X, v2.Y - v1.Y}
}

func (v1 Vector) Dist(v2 Vector) float64{
    return math.Sqrt(math.Pow(v2.X - v1.X, 2) + math.Pow(v2.Y - v1.Y, 2))
}

func (v1 *Vector) Add(v2 Vector){
    v1.X += v2.X
    v1.Y += v2.Y
}

func (v1 *Vector) Mul(m float64){
    v1.X *= m
    v1.Y *= m
}

type Body struct{
    Mass float64
    Radius float64
    Position Vector
    Velocity Vector
}

func main() {    
    window, renderer = tools.Start("planets", int(W), int(H))    
    defer window.Destroy()
    defer renderer.Destroy()
    
    sun := Body{ 1.988E30, 6.96E8, Vector{0,0}, Vector{0,0}}
    earth := Body{ 5.972E24, 6.371E6, Vector{1.5E11,0}, Vector{0, 2.9E4}}
    mars := Body{ 6.4185e23, 6.371E6, Vector{2.25e11,0}, Vector{0, 2.41e4}}
    ecc := Body{ 1e20, 1e5, Vector{2e11,0}, Vector{0, 3e4}}
    bodies := []*Body{&sun, &earth, &mars, &ecc}
    
    fmt.Println(earth.Position.Dist(sun.Position))
    renderer.SetDrawColor(0xFF,0xFF,0xFF,0xFF)
    renderer.Clear()  
    quit := false
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch e.(type){
            case *sdl.QuitEvent: 
                quit = true
            }
        }
             
        
        renderer.SetDrawColor(0,0,0,0xFF)
        
        for i := 0; i < 60*60*24; i++{
            UpdateBodies(bodies)
        }
        DrawBodies(bodies)
        
        
        renderer.Present()
        //fmt.Println(earth.Position, sun.Position)
    }
}

func DrawBodies(bodies []*Body){
    for _, v := range bodies{
        if renderer.DrawPoint(int(v.Position.X/Scale) + W/2, int(v.Position.Y/Scale) + H/2) != 0 {
            fmt.Fprintf(os.Stderr, "failed to draw point:", sdl.GetError())
        }
    }
}

func UpdateBodies(bodies []*Body){
    for _, a := range bodies{
        total := Vector{0, 0}
        for _, b := range bodies{
            if a == b{
                continue
            }
            dir := a.Position.To(b.Position)
            accel := 6.67E-11 * b.Mass / math.Pow(a.Position.Dist(b.Position), 3)
            dir.Mul(accel)
            total.Add(dir)            
        }
        a.Velocity.Add(total)
        a.Position.Add(a.Velocity)
    } 
}
