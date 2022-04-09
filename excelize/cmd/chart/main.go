package main

import (
	"github.com/xuri/excelize/v2"
)

func main() {
	categories := map[string]string{
		"A2": "Small",
		"A3": "Normal",
		"A4": "Large",
		"B1": "Apple",
		"C1": "Orange",
		"D1": "Pear",
	}
	values := map[string]int{
		"B2": 2,
		"C2": 3,
		"D2": 3,
		"B3": 5,
		"C3": 2,
		"D3": 4,
		"B4": 6,
		"C4": 7,
		"D4": 8,
	}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	f.SetCellStr("Sheet1", "A10", "根据指定路径保存文件")
	if err := f.AddChart("Sheet1", "E1", `{
        "type":"col3DClustered",
        "series":[
          {"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},
          {"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},
          {"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],
          "title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		println(err.Error())
		return
	}
	// 根据指定路径保存文件
	if err := f.SaveAs("chart.xlsx"); err != nil {
		println(err.Error())
	}
}
