package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "github.com/EvanTheB/spacegame/tools"
)




const(
    W int32 = 800
    H int32 = 600
)

var renderer *sdl.Renderer
var window *sdl.Window
var font *ttf.Font


func main() {
    window, renderer = tools.Start("font", int(W), int(H))    
    defer window.Destroy()
    defer renderer.Destroy()
    font = tools.LoadFont("resources/lazy.ttf", 28)
    
    text := tools.LoadText(renderer, font, "it works!", sdl.Color{128,128,128, 255})
        
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
        
        renderer.Copy(text, nil, nil)
                
        renderer.Present()
    }
}


