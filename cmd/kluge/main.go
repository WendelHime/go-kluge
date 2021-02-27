package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/WendelHime/go-kluge"
	"github.com/fogleman/gg"
)

func main() {
	// setting up vars
	fp := flag.String("filepath", "/home/wotan/Pictures/ada_in_blank.png", "the filepath for a segmented image")
	threshold := flag.Float64("threshold", 0.25, "a threshold for random creation of points, the value must be in the range 0, 100.")
	minDist := flag.Float64("minDist", 70, "minimum distance of points")
	output := flag.String("output", "./output.png", "output filepath")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s -filepath images/ada_in_blank.png -threshold 0.35 -minDist 50 -output images/ada_output.png\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// load image
	img, err := gg.LoadImage(*fp)
	if err != nil {
		panic(err)
	}

	// get random points
	points := kluge.GeneratePoints(img, *threshold)

	// draw lines between points
	ctx := kluge.BuildLines(img.Bounds().Max.X, img.Bounds().Max.Y, points, *minDist)

	// save image
	err = ctx.SavePNG(*output)
	if err != nil {
		log.Fatalf("fail saving png: %+v", err)
	}

}
