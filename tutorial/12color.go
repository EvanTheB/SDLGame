package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/EvanTheB/SDLGame/tools"
    "math/rand"
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
    
    window, renderer = tools.Start("color", int(W), int(H))    
    defer window.Destroy()
    defer renderer.Destroy()
    
    
    colors := tools.LoadTexture(renderer, "resources/color.png")
    var thisRGB = rgb{255,255,255}
    var nextRGB = rgb{255,255,255}
    var counter = int(255)
        
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
            thisRGB = nextRGB
            nextRGB.r = uint8(rand.Int31n(255))
            nextRGB.g = uint8(rand.Int31n(255))
            nextRGB.b = uint8(rand.Int31n(255))
        }
                            
        colors.SetColorMod( uint8(int(thisRGB.r)*counter/255 + int(nextRGB.r)*(255-counter)/255), 
                            uint8(int(thisRGB.g)*counter/255 + int(nextRGB.g)*(255-counter)/255), 
                            uint8(int(thisRGB.b)*counter/255 + int(nextRGB.b)*(255-counter)/255),) 
        renderer.Copy(colors, nil,nil)
        
        renderer.Present()
    }
}
