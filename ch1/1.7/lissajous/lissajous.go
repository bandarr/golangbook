//Server1 is a minimal "echo" server.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	cycles := r.URL.Query().Get("cycles")

	numcycles, err := strconv.Atoi(cycles)

	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi: %v\n", err)
	}

	lissajous(w, numcycles)
}

var palette = []color.Color{color.Black, color.RGBA{0x4C, 0xBB, 0x17, 0xff}, color.RGBA{0xD2, 0x04, 0x2D, 0xff}}

const ( // first color in palette
	blackIndex  = 0 // next color in palette
	greenIndex  = 1
	cherryIndex = 2
)

func lissajous(out io.Writer, numcycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	cycles := numcycles          // number of complete x oscillator revolutions
	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
