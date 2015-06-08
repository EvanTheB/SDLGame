package tools

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"os"
)

func LoadText(renderer *sdl.Renderer, font *ttf.Font, text string, textColor sdl.Color) *sdl.Texture {
	bmp, err := font.RenderUTF8_Solid(text, textColor)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load text %s! SDL Error: %s\n", ttf.GetError())
	} else {
		texture, err := renderer.CreateTextureFromSurface(bmp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create texture %s! SDL Error: %s\n", sdl.GetError())
		}
		bmp.Free()
		return texture
	}
	return nil
}

func LoadFont(path string, size int) *ttf.Font {
	ttf.Init()
	font, _ := ttf.OpenFont("resources/lazy.ttf", 28)
	if font == nil {
		fmt.Fprintf(os.Stderr, "Failed to create font: %s\n", ttf.GetError())
		os.Exit(1)
	}
	return font
}
