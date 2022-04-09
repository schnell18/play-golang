package main

import (
	"encoding/csv"
	"fmt"
	_ "image/png"
	"io"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	csvfile, err := os.Open("MSFT.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to open MSFT.csv due to: %v\n", err)
		os.Exit(1)
	}
	defer csvfile.Close()

	f := excelize.NewFile()
	sheetName := "stock"
	f.SetSheetName("Sheet1", sheetName)

	reader := csv.NewReader(csvfile)
	row := 1
	for {
		record, er := reader.Read()
		if er == io.EOF {
			break
		}
		if er != nil {
			fmt.Fprintf(os.Stderr, "Fail to read MSFT.csv due to: %v\n", er)
			os.Exit(2)
		}
		cell, _ := excelize.CoordinatesToCellName(1, row)
		if row == 1 {
			if err = f.SetSheetRow(sheetName, cell, &record); err != nil {
				fmt.Fprintf(os.Stderr, "Fail to set first row due to: %v\n", err)
				os.Exit(2)
			}
		} else {
			numbers := convertToSlice(record)
			if err = f.SetSheetRow(sheetName, cell, &numbers); err != nil {
				fmt.Fprintf(os.Stderr, "Fail to set row due to: %v\n", err)
				os.Exit(2)
			}

		}
		row++
	}

	if err = f.SetCellValue(sheetName, "A1", "交易日期"); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set trade date due to: %v\n", err)
		os.Exit(4)
	}

	if err = f.SetCellValue(sheetName, "B1", "开盘价"); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set open price due to: %v\n", err)
		os.Exit(4)
	}

	if err = f.SetSheetRow(sheetName, "C1", &[]string{
		"最高价", "最低价", "收盘价", "收盘调价", "交易量"}); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set header to: %v\n", err)
		os.Exit(4)
	}

	style1, _ := f.NewStyle(&excelize.Style{NumFmt: 2})
	if err = f.SetColStyle(sheetName, "B:F", style1); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set column style due to: %v\n", err)
		os.Exit(4)
	}

	style2, _ := f.NewStyle(&excelize.Style{NumFmt: 3})
	if err = f.SetColStyle(sheetName, "G", style2); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set column style due to: %v\n", err)
		os.Exit(4)
	}
	if err = f.SetColWidth(sheetName, "A", "G", 11); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set column width due to: %v\n", err)
		os.Exit(4)
	}
	if err = f.SetColVisible(sheetName, "F", false); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to hide column F due to: %v\n", err)
		os.Exit(4)
	}

	if err = f.InsertRow(sheetName, 1); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to insert row due to: %v\n", err)
		os.Exit(4)
	}

	if err = f.MergeCell(sheetName, "A1", "G1"); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to merge cells due to: %v\n", err)
		os.Exit(4)
	}

	if err = f.SetCellRichText(sheetName, "A1", []excelize.RichTextRun{
		{
			Font: &excelize.Font{
				Bold:   true,
				Family: "Times New Roman",
				Size:   20,
				Color:  "#2354e8",
			},
			Text: "MSFT\r\n",
		},
		{
			Font: &excelize.Font{
				Bold:   true,
				Family: "Microsfot Yahei",
				Size:   16,
			},
			Text: "近五年日交易价格数据",
		},
	}); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set rich text due to: %v\n", err)
		os.Exit(4)
	}

	style3, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText:   true,
			Horizontal: "center",
			Vertical:   "center",
		}})
	if err = f.SetCellStyle(sheetName, "A1", "A1", style3); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set text wrap of first row due to: %v\n", err)
		os.Exit(4)
	}
	if err = f.SetRowHeight(sheetName, 1, 60); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set height of first row due to: %v\n", err)
		os.Exit(4)
	}

	sheetId := f.NewSheet("trend")
	f.SetActiveSheet(sheetId)

	if err = f.AddChart("trend", "A1", `{
        "type": "line",
        "series": [
            {
                "name": "stock!$E$2",
                "categories": "stock!$A$3:$A$1262",
                "values": "stock!$E$3:$E$1262",
                "marker": {
                    "symbol": "none"
                }
            }
        ],
        "format": {
            "x_scale": 1.6,
            "y_scale": 1.5,
            "x_offset": 15,
            "y_offset": 10
        },
        "legend": {
            "none": true
        },
        "x_axis": {
            "tick_label_strip": 60
        },
        "title": {
            "name": "收盘价"
        }
    }`); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to add chart due to: %v\n", err)
		os.Exit(3)
	}

	if err = f.AddChart("trend", "A24", `{
        "type": "area",
        "series": [
            {
                "name": "stock!$G$2",
                "categories": "stock!$A$3:$A$1262",
                "values": "stock!$G$3:$G$1262"
            }
        ],
        "format": {
            "x_scale": 1.6,
            "y_scale": 1.5,
            "x_offset": 15
        },
        "legend": {
            "none": true
        },
        "x_axis": {
            "tick_label_strip": 60
        },
        "title": {
            "name": "成交量"
        }
    }`); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to add chart due to: %v\n", err)
		os.Exit(3)
	}

	if err := f.SaveAs("msft.xlsx"); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to save msft.xlsx due to: %v\n", err)
		os.Exit(4)
	}

}

func convertToSlice(record []string) (numbers []interface{}) {
	for _, rec := range record {
		if f, err := strconv.ParseFloat(rec, 64); err != nil {
			numbers = append(numbers, rec)
		} else {
			numbers = append(numbers, f)
		}
	}
	return
}
