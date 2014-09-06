package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "github.com/veandco/go-sdl2/sdl_image"
    "github.com/veandco/go-sdl2/sdl_mixer"
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
var haptic *sdl.Haptic
var joystick *sdl.Joystick


func main() {
    
    start()    
    defer window.Destroy()
    defer renderer.Destroy()
    
    

    
    screens := loadScreens()
    
        
    quit := false
    
    cur := screens["main"]

    
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch e.(type){
            case *sdl.QuitEvent: 
                quit = true
            case *sdl.JoyButtonEvent:
                if haptic.RumblePlay(0.75, 500) != 0{
                    fmt.Printf( "Warning: Unable to play rumble:", sdl.GetError() )
                }
            }
            
        }
        
        
        
        renderer.SetDrawColor(0xFF,0xFF,0xFF,0xFF)
        renderer.Clear()    
        
        
        renderer.Copy(cur, nil, nil)
        
        
        renderer.Present()
    }
}


func loadScreens() map[string]*sdl.Texture{
    ret := make(map[string]*sdl.Texture)
    ret["main"] = loadTexture("resources/press.bmp")
    ret["up"] = loadTexture("resources/up.bmp")
    ret["right"] = loadTexture("resources/right.bmp")
    ret["left"] = loadTexture("resources/left.bmp")
    ret["down"] = loadTexture("resources/down.bmp")
    ret["png"] = loadTexture("resources/mouse.png")
    return ret
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
    sdl.Init(sdl.INIT_JOYSTICK | sdl.INIT_HAPTIC)
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
	
	if sdl.NumJoysticks() < 1 {
        fmt.Println("no joysticks bra")
        os.Exit(1)
    }
    joystick = sdl.JoystickOpen(0)
    if joystick == nil{
        fmt.Println("Joystick failed to open", sdl.GetError())
        os.Exit(1)
    }
    haptic = sdl.HapticOpenFromJoystick(joystick)
    if haptic == nil{
        fmt.Println("Haptic failed to load", sdl.GetError())
        os.Exit(1)
    }
    haptic.RumbleInit()
}

