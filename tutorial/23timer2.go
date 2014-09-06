package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "github.com/veandco/go-sdl2/sdl_image"
    "fmt"
    "os"
//    "math/rand"
)



const(
    W int32 = 800
    H int32 = 600
)

var renderer *sdl.Renderer
var window *sdl.Window
var font *ttf.Font


func main() {
    
    start()    
    defer window.Destroy()
    defer renderer.Destroy()
    
    text := loadText("Press enter to reset", sdl.Color{0,0,0, 255})
    startTime := sdl.GetTicks()
    
        
    quit := false
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch t := e.(type){
            case *sdl.QuitEvent: 
                quit = true
            case *sdl.KeyDownEvent:
                if t.Keysym.Sym == sdl.K_RETURN{
                    startTime = sdl.GetTicks()
                }
            
            }
        }
        
        renderer.SetDrawColor(0xFF,0xFF,0xFF,0xFF)
        renderer.Clear()    
        timeText := loadText(fmt.Sprintf("time: %d", sdl.GetTicks() - startTime), sdl.Color{0,0,0, 255})      
        renderer.Copy(text, nil, &sdl.Rect{0,0,400,100})
        renderer.Copy(timeText, nil, &sdl.Rect{400,0,400,100})

        
        
        renderer.Present()
    }
}

func loadText(text string, textColor sdl.Color) *sdl.Texture{
    bmp := font.RenderText_Solid(text, textColor)
    if bmp == nil{
        fmt.Fprintf(os.Stderr, "Unable to load text %s! SDL Error: %s\n", ttf.GetError() )
    }else{
        
        texture := renderer.CreateTextureFromSurface(bmp)
        if texture == nil{
            fmt.Fprintf(os.Stderr,"Unable to create texture %s! SDL Error: %s\n", sdl.GetError() )
        }
        bmp.Free()
        return texture
    }
    return nil
}

func loadTexture(path string) *sdl.Texture{
    bmp := img.Load(path)
    if bmp == nil{
        fmt.Fprintf(os.Stderr, "Unable to load image %s! SDL Error: %s\n", path, img.GetError() )
    }else{
        if bmp.SetColorKey(1, sdl.Color{0xFF,0xFF,0x00,0xFF}.Uint32()) != 0{
            fmt.Fprintf(os.Stderr,"Unable to set Color Key %s! SDL Error: %s\n", path, sdl.GetError() )           
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
            int(W), int(H), sdl.WINDOW_SHOWN)
    if window == nil{
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}
		
	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC)     
	if renderer == nil{
	    fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(1)
	}	
	ttf.Init()
	font, _ = ttf.OpenFont("resources/lazy.ttf", 28)
	if (font == nil){
	    fmt.Fprintf(os.Stderr, "Failed to create font: %s\n", ttf.GetError())
		os.Exit(1)
	}	
}

