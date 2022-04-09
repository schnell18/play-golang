package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("picture.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	if err = f.AddPicture("Sheet1", "A2", "images/texlion.png", ""); err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	if err = f.AddPicture("Sheet1", "D20", "images/texlion2.jpg", `{
        "x_scale": 0.3,
        "y_scale": 0.3
    }`); err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	if err = f.AddPicture("Sheet1", "H2", "images/latexlion.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}
	// Save the spreadsheet with the origin path.
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
	}
}
