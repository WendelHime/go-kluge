package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

// Point represents a point, speciying the coordinate
// and the RGB color at the same position
type Point struct {
	X int
	Y int
	R uint8
	G uint8
	B uint8
}

func newPoint(x, y int, r, g, b uint8) *Point {
	return &Point{
		X: x,
		Y: y,
		R: r,
		G: g,
		B: b,
	}
}

func randFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// GeneratePoints run through the image creating random points
func GeneratePoints(img image.Image, threshold float64) []*Point {
	points := make([]*Point, 0)

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			if uint8(r) <= 250 && uint8(g) <= 250 && uint8(b) <= 250 {
				val := randFloats(0, 100)
				if val < threshold {
					points = append(points, newPoint(x, y, uint8(r), uint8(g), uint8(b)))
				}
			}
		}
	}
	return points
}

// BuildLines create a new image and draw lines between points with
// a minimum distance
func BuildLines(w, h int, points []*Point, minDist float64) *gg.Context {
	// create empty context
	ctx := gg.NewContext(w, h)

	// iterate over points and draw a line if they
	// have a mininum distance
	for i, p := range points {
		for j := 0; j < i; j++ {
			if math.Hypot(float64(p.X-points[j].X), float64(p.Y-points[j].Y)) < minDist {
				ctx.MoveTo(float64(p.X), float64(p.Y))
				ctx.LineTo(float64(points[j].X), float64(points[j].Y))
				// ctx.SetLineWidth(0.65)
				ctx.SetRGB255(int(points[j].R), int(points[j].G), int(points[j].B))
				ctx.Stroke()
			}
		}
	}
	return ctx
}

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
	points := GeneratePoints(img, *threshold)

	// draw lines between points
	ctx := BuildLines(img.Bounds().Max.X, img.Bounds().Max.Y, points, *minDist)

	// save image
	err = ctx.SavePNG(*output)
	if err != nil {
		log.Fatalf("fail saving png: %+v", err)
	}

}
