package lightning

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strings"
)

type Gradient struct {
	r float64
	g float64
	b float64
}

func (g *Gradient) GetColor(state float64) color.Color {
	return color.RGBA{
		R: uint8(math.Min(255, g.r*state)),
		G: uint8(math.Min(255, g.g*state)),
		B: uint8(math.Min(255, g.b*state)),
		A: 255,
	}
}

// colored
func drawLine(image *image.RGBA, line []Vector) {
	grad := Gradient{r: 250, g: 200, b: 500}
	for x := image.Rect.Min.X; x < image.Rect.Max.X; x++ {
		for y := image.Rect.Min.Y; y < image.Rect.Max.Y; y++ {
			pt := Vector{X: float64(x), Y: float64(y)}
			var minDist float64 = 1000000
			for _, v := range line {
				if d := pt.Sub(v).LengthSqr(); minDist > d {
					minDist = d
				}
			}
			image.Set(x, y, grad.GetColor(math.Pow(minDist, -0.2)))
		}
	}
}

// white line on black
func drawSimple(image *image.RGBA, line []Vector) {
	total := len(image.Pix)
	for i := 3; i < total; i++ {
		image.Pix[i] = 255
	}
	for _, v := range line {
		image.Set(int(v.X), int(v.Y), color.White)
	}
}

func Draw(line []Vector) image.Image {
	min := vector0
	max := vector0
	for _, v := range line {
		min = min.Min(v)
		max = max.Max(v)
	}
	max.Add(vector1)
	min.Sub(vector1)
	img := image.NewRGBA(image.Rect(int(min.X), int(min.Y), int(max.X), int(max.Y)))
	drawLine(img, line)
	return img
}

func Save(file string, image image.Image) {
	if !strings.HasSuffix(file, ".png") {
		file += ".png"
	}
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, image); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}
