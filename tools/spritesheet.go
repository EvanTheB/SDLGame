package tools

import (
    "github.com/veandco/go-sdl2/sdl"

)

type SpriteSheet struct{
    Sheet *sdl.Texture
    Sprites []sdl.Rect
    CurrentSprite int
}

func (s *SpriteSheet) Render(renderer *sdl.Renderer, dst *sdl.Rect){
    renderer.Copy( s.Sheet, &s.Sprites[s.CurrentSprite], dst)
}

func (s *SpriteSheet) Next(){
    s.CurrentSprite = (s.CurrentSprite + 1) % len(s.Sprites)
}
