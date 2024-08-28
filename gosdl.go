package main

import (
	"local/gosdl/insignals"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	ins := insignals.New()

	err := ttf.Init()
	if err != nil {
		panic(err)
	}
	defer ttf.Quit()

	fonty, err := ttf.OpenFont("./FiraCodeNerdFont-Regular.ttf", 12)
	if err != nil {
		panic(err)
	}
	defer fonty.Close()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("the best window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// implement of rederer flags like render to texture
	rend, err := sdl.CreateRenderer(window, -1, 2)
	if err != nil {
		panic(err)
	}
	defer rend.Destroy()

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	rect2 := sdl.Rect{X: 200, Y: 200, W: 200, H: 20}

	var fg sdl.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	msgsurface, err := fonty.RenderUTF8Solid("test string <=", fg)
	if err != nil {
		panic(err)
	}
	msgtexture, err := rend.CreateTextureFromSurface(msgsurface)
	if err != nil {
		panic(err)
	}

	for ins.Quit == false && ins.Keys[sdl.SCANCODE_ESCAPE] == false {
		ins.Update()
		rend.SetDrawColor(128, 128, 128, 255)
		rend.Clear()
		rend.SetDrawColor(255, 0, 255, 255)
		rend.FillRect(&rect)
		rend.Copy(msgtexture, nil, &rect2)

		rend.Present()
	}
}
