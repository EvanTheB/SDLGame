package tools

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"os"
)

func LoadTextureTransparent(renderer *sdl.Renderer, path string, r, g, b uint8) *sdl.Texture {
	bmp, err := img.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load image %s! SDL Error: %s\n", path, img.GetError())
	} else {
		if bmp.SetColorKey(1, sdl.MapRGB(bmp.Format, r, g, b)) != nil {
			fmt.Fprintf(os.Stderr, "Unable to set Color Key %s! SDL Error: %s\n", path, sdl.GetError())
		}
		texture, err := renderer.CreateTextureFromSurface(bmp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create texture %s! SDL Error: %s\n", path, sdl.GetError())
		}
		bmp.Free()
		return texture
	}
	return nil
}

func LoadTexture(renderer *sdl.Renderer, path string) *sdl.Texture {
	bmp, err := img.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load image %s! SDL Error: %s\n", path, img.GetError())
	} else {
		texture, err := renderer.CreateTextureFromSurface(bmp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create texture %s! SDL Error: %s\n", path, sdl.GetError())
		}
		bmp.Free()
		return texture
	}
	return nil
}

func Start(windowText string, W, H int) (*sdl.Window, *sdl.Renderer) {
	window, err := sdl.CreateWindow(windowText, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		W, H, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(1)
	}
	return window, renderer
}
