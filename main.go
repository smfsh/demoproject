package main

import (
	"fmt"
	"image/color"

	"github.com/eliukblau/pixterm/pkg/ansimage"
)

var bg = color.RGBA{
	R: 0,
	G: 0,
	B: 0,
	A: 0,
}
var dm ansimage.DitheringMode = 0
var sm ansimage.ScaleMode = 0

func main() {
	fmt.Println("Hello World!")
	fmt.Println("My name is Miles")

	snyk := getImage("https://theme.zdassets.com/theme_assets/9132644/8fea1d21f84207552082e1091daebd33032edd75.png")
	fmt.Print(snyk)
}

func getImage(url string) string {
	image, err := ansimage.NewScaledFromURL(url, 60, 60, bg, sm, dm)
	if err != nil {
		panic(err)
	}
	return image.Render()
}
