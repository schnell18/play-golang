// go get -u github.com/mbndr/figlet4go/...
package main

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

func main() {
	render := figlet4go.NewAsciiRender()
	renderStr, _ := render.Render("BANANA")
	fmt.Print(renderStr)

	// With render options (colors, font)
	opts := figlet4go.NewRenderOptions()
	opts.FontName = "larry3d"
	opts.FontColor = []figlet4go.Color{figlet4go.ColorYellow}
	s, _ := render.RenderOpts("BANANA", opts)
	fmt.Print(s)
}
