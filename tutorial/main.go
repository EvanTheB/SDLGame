package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_image"
    "fmt"
    "os"
    //"math/rand"
)

var window *sdl.Window
var surface *sdl.Surface


func main() {
    window, surface = start()
    defer window.Destroy()
    screens := loadScreens()
    defer unloadScreens(screens)
    curScreen := screens["main"]   
    
    quit := false
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch t := e.(type){
            case *sdl.QuitEvent: 
                quit = true
            case *sdl.KeyDownEvent:
                switch t.Keysym.Sym{
                case sdl.K_UP:
                    curScreen = screens["up"]
                case sdl.K_DOWN:
                    curScreen = screens["down"]
                case sdl.K_LEFT:
                    curScreen = screens["left"]
                case sdl.K_RIGHT:
                    curScreen = screens["right"]
                case sdl.K_SPACE:
                    curScreen = screens["main"]            
                }
                
            }
        }
        //rect := sdl.Rect { rand.Int31n(800), rand.Int31n(600), 100, 100 }
        
        
        //surface.FillRect(&rect, rand.Uint32())
        var stretchRect sdl.Rect
        surface.GetClipRect(&stretchRect)
        curScreen.BlitScaled(nil, surface, &stretchRect)
        screens["png"].Blit(nil, surface, nil)
        window.UpdateSurface()
    }
}

func loadPNG(path string) *sdl.Surface{
    bmp := img.Load(path)
    if bmp == nil{
        fmt.Println( "Unable to load image %s! SDL Error: %s\n", path, img.GetError() )
    }else{
        opt := bmp.Convert(surface.Format, 0)
        if opt == nil{
            fmt.Println( "Unable to optimise image %s! SDL Error: %s\n", path, img.GetError() )
        }
        bmp.Free()
        return opt
    }
    return bmp
}

func loadBMP(path string) *sdl.Surface{
    bmp := sdl.LoadBMP(path)
    if bmp == nil{
        fmt.Println( "Unable to load image %s! SDL Error: %s\n", path, sdl.GetError() )
    }else{
        opt := bmp.Convert(surface.Format, 0)
        if opt == nil{
            fmt.Println( "Unable to optimise image %s! SDL Error: %s\n", path, sdl.GetError() )
        }
        bmp.Free()
        return opt
    }
    return bmp
}

func loadScreens() map[string]*sdl.Surface{
    ret := make(map[string]*sdl.Surface)
    ret["main"] = loadBMP("resources/press.bmp")
    ret["up"] = loadBMP("resources/up.bmp")
    ret["right"] = loadBMP("resources/right.bmp")
    ret["left"] = loadBMP("resources/left.bmp")
    ret["down"] = loadBMP("resources/down.bmp")
    ret["png"] = loadPNG("resources/mouse.png")
    return ret
}

func unloadScreens(s map[string]*sdl.Surface) {
    for _,v := range s{
        v.Free()
    }
}

func start() (*sdl.Window, *sdl.Surface){ 
    window := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
            800, 600, sdl.WINDOW_SHOWN)
    if window == nil{
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}
	
	surface := window.GetSurface()
	if surface == nil{
	    fmt.Fprintf(os.Stderr, "Failed to create surface: %s\n", sdl.GetError())
		os.Exit(1)
	}
    
    return window, surface
}
