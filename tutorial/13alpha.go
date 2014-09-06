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

type rgb struct{
    r, g, b uint8
}

func main() {
    
    window, renderer = tools.Start("alpha blend", int(W), int(H))    
    defer window.Destroy()
    defer renderer.Destroy()
    
    
    colors := tools.LoadTexture(renderer, "resources/color.png")
    back := tools.LoadTexture(renderer, "resources/background.png")
    var counter = uint8(255)
        
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
        
        counter--
        if counter == 0 {
            counter = 255            
        }
                            
        colors.SetAlphaMod( counter)
        renderer.Copy(back, nil, nil)
        renderer.Copy(colors, nil,nil)
        
        renderer.Present()
    }
}

