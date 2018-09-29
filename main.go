package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"sort"

	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var filters = map[string]imaging.ResampleFilter{
	"NearestNeighbor":   imaging.NearestNeighbor,
	"Box":               imaging.Box,
	"Linear":            imaging.Linear,
	"Hermite":           imaging.Hermite,
	"MitchellNetravali": imaging.MitchellNetravali,
	"CatmullRom":        imaging.CatmullRom,
	"BSpline":           imaging.BSpline,
	"Gaussian":          imaging.Gaussian,
	"Lanczos":           imaging.Lanczos,
	"Hann":              imaging.Hann,
	"Hamming":           imaging.Hamming,
	"Blackman":          imaging.Blackman,
	"Bartlett":          imaging.Bartlett,
	"Welch":             imaging.Welch,
	"Cosine":            imaging.Cosine,
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: resize [inputfile] [outputfile]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

// https://stackoverflow.com/a/38300583
func addLabel(img *image.NRGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
	}

	// Open a test image.
	src, err := imaging.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	width, height := src.Bounds().Max.X, src.Bounds().Max.Y

	filterNames := make([]string, 0, len(filters))
	for k := range filters {
		filterNames = append(filterNames, k)
	}

	sort.Strings(filterNames)

	dst := imaging.New(width*2, height*2, color.NRGBA{0, 0, 0, 0})

	for i, k := range filterNames {
		filter := filters[k]

		x := (width / 2) * (i % 4)
		y := (height / 2) * (i / 4)

		resized := imaging.Resize(src, width/2, height/2, filter)
		dst = imaging.Paste(dst, resized, image.Pt(x, y))
		addLabel(dst, x+10, y+20, k)
	}

	err = imaging.Save(dst, os.Args[2])
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

}
