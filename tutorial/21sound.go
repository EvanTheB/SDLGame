package main

import (
    "github.com/veandco/go-sdl2/sdl"
    
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
var (
    scratch *mix.Chunk
    high    *mix.Chunk
    medium *mix.Chunk
    low *mix.Chunk
    music *mix.Music
)


func main() {
    
    start()    
    defer window.Destroy()
    defer renderer.Destroy()
    loadAudio()
    

    
    screens := loadScreens()
    
        
    quit := false
    
    cur := screens["main"]
    
    
    for !quit{
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent(){
            switch t := e.(type){
            case *sdl.QuitEvent: 
                quit = true
            case *sdl.KeyDownEvent:
                switch t.Keysym.Sym{
                case sdl.K_1:
                    high.PlayChannel(-1, 0)
                case sdl.K_2:
                    medium.PlayChannel(-1, 0)
                case sdl.K_3:
                    low.PlayChannel(-1, 0)
                case sdl.K_4:
                    scratch.PlayChannel(-1, 0)
                case sdl.K_9:
                    if !mix.MusicPlaying(){
                        if !music.Play(1){
                            fmt.Printf("music did not play: %s", sdl.GetError())
                        }
                    }else{
                        if !mix.PausedMusic(){
                            mix.PauseMusic()    
                        }else{
                            mix.ResumeMusic()    
                        }
                        
                    }                    
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
    ret["main"] = loadTexture("resources/prompt.png")
    ret["up"] = loadTexture("resources/up.bmp")
    ret["right"] = loadTexture("resources/right.bmp")
    ret["left"] = loadTexture("resources/left.bmp")
    ret["down"] = loadTexture("resources/down.bmp")
    ret["png"] = loadTexture("resources/mouse.png")
    return ret
}

func loadAudio(){
    scratch = mix.LoadWAV("resources/scratch.wav")
    high = mix.LoadWAV("resources/high.wav")
    medium = mix.LoadWAV("resources/medium.wav")
    low = mix.LoadWAV("resources/low.wav")
    music = mix.LoadMUS("resources/beat.wav")
    if music == nil{
        fmt.Printf("failed to load music")
    }    
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
    if sdl.Init(sdl.INIT_AUDIO) < 0{
		fmt.Fprintf(os.Stderr, "Failed to init:", sdl.GetError())
		os.Exit(1)
    }
       
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
    
    mix.SetSoundFonts("resources/")    
    if !mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 2048){
        fmt.Fprintf(os.Stderr, "Failed to open audio:%s\n", sdl.GetError())
		os.Exit(1)    
    }	

}

