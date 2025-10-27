// go.mod: add module; run `go get github.com/common-nighthawk/go-figure`
package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {
	// Create and print a BANANA banner
	f := figure.NewFigure("BANANA", "", true) // "" = default font, strict=true
	f.Print()
	f = figure.NewColorFigure("BANANA", "", "yellow", true)
	f.Print()
}
