package main

import (
	"flag"
	"go-lightning/lightning"
	"log"
	"time"
)

func main() {
	var (
		height float64
		width  float64
		file   string
		line   []lightning.Vector
	)

	flag.Float64Var(&height, "height", 720, "preffered height of output image")
	flag.Float64Var(&width, "width", 1280, "preffered width of output image")
	flag.StringVar(&file, "file", "lightning.png", "preffered filename of output image")
	flag.Parse()

	log.Printf("started: %s %.0fx%.0f", file, width, height)

	const stopDistance = 1.5
	gen := lightning.NewRandGenerator(0.25, time.Now().Unix())
	startPoint := lightning.Vector{X: width / 2, Y: 0}

	line = append(line, lightning.Build(startPoint, lightning.Vector{X: 0, Y: height}, stopDistance, gen)...)
	if width > 0 {
		line = append(line, lightning.Build(startPoint, lightning.Vector{X: width, Y: height}, stopDistance, gen)...)
	}

	log.Printf("total points: %d", len(line))
	lightning.Save(file, lightning.Draw(line))
	log.Printf("saved")
}
