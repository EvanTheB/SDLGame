package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/EvanTheB/spacegame/tools"
)



const(
    W int32 = 800
    H int32 = 600
)

var renderer *sdl.Renderer
var window *sdl.Window

func main() {    
    window, renderer = tools.Start("sprites", int(W), int(H))    
    defer window.Destroy()
    defer renderer.Destroy()
    
    
    sprite := tools.LoadTextureTransparent(renderer, "resources/sprites.png", 0, 0xFF, 0xFF)
    
        
    quit := false
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch e.(type){
            case *sdl.QuitEvent: 
                quit = true
            }
        }
        renderer.SetDrawColor(0xFF,0xFF,0xFF,0xFF)
        renderer.Clear()       
         
        renderer.Copy(sprite, &sdl.Rect{0,0,100,100}, &sdl.Rect{0,0,100,100})
        renderer.Copy(sprite, &sdl.Rect{0,100,100,100}, &sdl.Rect{W-100,0,100,100})
        renderer.Copy(sprite, &sdl.Rect{100,0,100,100}, &sdl.Rect{W-100,H-100,100,100})
        renderer.Copy(sprite, &sdl.Rect{100,100,100,100}, &sdl.Rect{0,H-100,100,100})
        
        renderer.Present()
    }
}




