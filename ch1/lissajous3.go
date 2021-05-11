package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete3 = []color.Color{
	color.RGBA{ R: 42, G: 187, B: 155, A: 1},
	color.RGBA{ R: 255, G: 255, B: 255, A: 1},
	color.RGBA{ R: 233, G: 212, B: 96, A: 1},
	color.Black,
	color.White,
}
const (
	greenIndex3 = 0 // first color in palette
	blueIndex3 = 1 // next color in palette
	blackIndex3 = 2 // next color in palette
	whiteIndex3 = 3 // next color in palette
)

func main() {
	lissajous3(os.Stdout)
}

func lissajous3(out io.Writer) {
	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	var colorIndex uint8 = greenIndex3
	for i := 0; i < nframes; i++ {
		rect := image.Rect(
			0,
			0,
			2*size+1,
			2*size+1,
		)
		img := image.NewPaletted(rect, pallete3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				colorIndex,
			)
		}
		if i % (nframes / len(pallete3)) == 0 {
			colorIndex++
		}
		if colorIndex % uint8(len(pallete3)) == 0 {
			colorIndex = 0
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
