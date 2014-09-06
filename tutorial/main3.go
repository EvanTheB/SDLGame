package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "fmt"
    "os"
    //"math/rand"
)

var window *sdl.Window
var renderer *sdl.Renderer

const(
    W int = 800
    H int = 600
)
    

func main() {
    start()
    defer window.Destroy()
    defer renderer.Destroy()
    man := loadTexture("resources/foo.png", true)
    back := loadTexture("resources/background.png", false)
    defer back.Destroy()
    defer man.Destroy()
    
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
        
        
        renderer.Copy(back, nil, nil)
        renderer.Copy(man, nil, &sdl.Rect{0,0,50,100})
        renderer.SetDrawColor(0xFF,0xFF,0x00,0xFF)
        renderer.FillRect( &sdl.Rect{100,100,50,100})
        
        renderer.Present()
    }
}

func loadTexture(path string, key bool) *sdl.Texture{
    bmp := img.Load(path)
    if bmp == nil{
        fmt.Fprintf(os.Stderr, "Unable to load image %s! SDL Error: %s\n", path, img.GetError() )
    }else{
        if key{
            if bmp.SetColorKey(1, sdl.Color{0xFF,0xFF,0x00,0xFF}.Uint32()) != 0{
                fmt.Fprintf(os.Stderr,"Unable to set Color Key %s! SDL Error: %s\n", path, sdl.GetError() ) 
            }   
        }
        texture := renderer.CreateTextureFromSurface(bmp)
        if texture == nil{
            fmt.Fprintf(os.Stderr,"Unable to create texture %s! SDL Error: %s\n", path, sdl.GetError() )
        }
        bmp.Free()
        return texture
    }
    return nil
}

func start(){ 
    window = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
            W, H, sdl.WINDOW_SHOWN)
    if window == nil{
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}	
	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)     
	if renderer == nil{
	    fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(1)
	}
}
