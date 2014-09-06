package tools

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "fmt"
    "os"
)

func LoadText(renderer *sdl.Renderer, font *ttf.Font, text string, textColor sdl.Color) *sdl.Texture{
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

func LoadFont(path string, size int) *ttf.Font{
    ttf.Init()
	font, _ := ttf.OpenFont("resources/lazy.ttf", 28)
	if (font == nil){
	    fmt.Fprintf(os.Stderr, "Failed to create font: %s\n", ttf.GetError())
		os.Exit(1)
	}
	return font
}
