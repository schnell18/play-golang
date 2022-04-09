package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	sheet := f.NewSheet("order")

	f.SetCellValue("order", "A2", "Hello universe")
	f.SetCellValue("order", "B2", 100)

	f.SetActiveSheet(sheet)
	if err := f.SaveAs("hello.xlsx"); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to save hello.xlsx due to %v\n", err.Error())
		os.Exit(3)
	}
}
