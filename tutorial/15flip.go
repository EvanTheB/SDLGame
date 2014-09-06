package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/EvanTheB/SDLGame/tools"
)



const(
    W int32 = 800
    H int32 = 600
)

var renderer *sdl.Renderer
var window *sdl.Window



func main() {
    window, renderer = tools.Start("flip", int(W), int(H))    
    defer window.Destroy()
    defer renderer.Destroy()
    
    arrow := tools.LoadTexture(renderer, "resources/arrow.png")
    back := tools.LoadTexture(renderer, "resources/background.png")
    
    angle := 0.0
    var flip = sdl.RendererFlip(sdl.FLIP_NONE)
            
    quit := false
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch t := e.(type){
            case *sdl.QuitEvent: 
                quit = true
            case *sdl.KeyDownEvent:
                switch t.Keysym.Sym{
                case sdl.K_LEFT:
                    angle += 30.0                    
                case sdl.K_RIGHT:
                    angle -= 30.0
                case sdl.K_q:
                    flip ^= sdl.FLIP_HORIZONTAL
                case sdl.K_w:
                    flip ^= sdl.FLIP_VERTICAL
                }
                
            }
        }
        
        renderer.SetDrawColor(0xFF,0xFF,0xFF,0xFF)
        renderer.Clear()    
        
        
        renderer.Copy(back, nil, nil)
        renderer.CopyEx(arrow, nil, &sdl.Rect{300,400,100,100},
                        angle, &sdl.Point{50,50}, flip)
        
        renderer.Present()
    }
}


