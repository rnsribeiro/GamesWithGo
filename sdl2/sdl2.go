package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH, HEIGHT int32 = 1020, 765

type color struct {
	r, g, b byte
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*int(WIDTH) + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func main() {

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(WIDTH), int32(HEIGHT))
	if err != nil {
		fmt.Println(err)
	}
	defer tex.Destroy()

	pixels := make([]byte, WIDTH*HEIGHT*4)

	for y := 0; y < int(HEIGHT); y++ {
		for x := 0; x < int(WIDTH); x++ {
			setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
		}
	}
	tex.Update(nil, unsafe.Pointer(&pixels[0]), int(WIDTH)*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	sdl.Delay(2000)

}
