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

var walkClips = []sdl.Rect{
    {0, 0, 64, 205},
    {64, 0, 64, 205},
    {128, 0, 64, 205},
    {196, 0, 64, 205},
}

func main() {
    
    window, renderer = tools.Start("animated sprite", int(W), int(H))
    defer window.Destroy()
    defer renderer.Destroy()
    
    
    walkSprite := tools.SpriteSheet{ tools.LoadTexture(renderer, "resources/walk.png"),
                                       []sdl.Rect{
                                            {0, 0, 64, 205},
                                            {64, 0, 64, 205},
                                            {128, 0, 64, 205},
                                            {196, 0, 64, 205},
                                        },
                                        0,
                                 }

    
    back := tools.LoadTexture(renderer, "resources/background.png")
    
    const frameDelay = 6
    delayCount := 0
            
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
        
        delayCount++
        if delayCount >= frameDelay{
            delayCount = 0
            walkSprite.Next()
        }
        renderer.Copy(back, nil, nil)
        walkSprite.Render(renderer, nil)
        
        renderer.Present()
    }
}


