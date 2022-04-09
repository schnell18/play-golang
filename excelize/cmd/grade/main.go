package main

import (
	"fmt"
	_ "image/png"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	sheetName := "grade"
	f.SetSheetName("Sheet1", sheetName)
	data := [][]interface{}{
		{"考试成绩统计表"},
		{"考试名称：期中考试", nil, nil, nil, "基础科目", nil, nil, "理科科目"},
		{"序号", "学号", "姓名", "班级", "数学", "英语", "语文", "化学", "生物", "物理", "总分"},
		{1, 10001, "学生A", "1班", 93, 80, 89, 86, 57, 77},
		{2, 10002, "学生B", "1班", 65, 72, 91, 75, 64, 90},
		{3, 10003, "学生C", "2班", 92, 99, 89, 86, 79, 69},
		{4, 10004, "学生D", "1班", 72, 69, 71, 84, 75, 83},
		{5, 10005, "学生E", "2班", 81, 93, 59, 76, 64, 90},
		{6, 10006, "学生F", "2班", 92, 90, 82, 96, 92, 70},
	}

	for i, row := range data {
		startCell, _ := excelize.JoinCellName("A", i+1)
		if err := f.SetSheetRow(sheetName, startCell, &row); err != nil {
			fmt.Fprintf(os.Stderr, "fail to SetSheetRow due to %v\n", err)
		}
	}

	formulaType, ref := excelize.STCellFormulaTypeShared, "K4:K9"
	err := f.SetCellFormula(sheetName, "K4", "SUM(E4:J4)", excelize.FormulaOpts{Ref: &ref, Type: &formulaType})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set cell formula due to: %v\n", err)
		os.Exit(2)
	}

	mergeCells := [][]string{
		{"A1", "K1"},
		{"A2", "D2"},
		{"E2", "G2"},
		{"H2", "J2"},
	}

	for _, ranges := range mergeCells {
		if err = f.MergeCell(sheetName, ranges[0], ranges[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Fail to merge cells due to: %v\n", err)
			os.Exit(3)
		}
	}

	style1, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Fill:      excelize.Fill{Type: "patten", Color: []string{"#34c6eb"}, Pattern: 1},
	})
	if err = f.SetCellStyle(sheetName, "A1", "A1", style1); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set cell style due to: %v\n", err)
		os.Exit(3)
	}

	style2, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	for _, ranges := range mergeCells[1:] {
		if err = f.SetCellStyle(sheetName, ranges[0], ranges[0], style2); err != nil {
			fmt.Fprintf(os.Stderr, "Fail to set cell style due to: %v\n", err)
			os.Exit(3)
		}
	}

	// if err = f.SetColWidth(sheetName, "D", "K", 6); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Fail to set col width due to: %v\n", err)
	// 	os.Exit(3)
	// }

	if err = f.AddTable(sheetName, "A3", "K9", `{
        "table_name": "table",
        "table_style": "TableStyleLight2"
    }`); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to add table due to: %v\n", err)
		os.Exit(3)
	}

	if err = f.AddChart(sheetName, "A10", `{
        "type": "col",
        "series": [
            {
                "name": "grade!$A$2",
                "categories": "grade!$C$4:$C$9",
                "values": "grade!$K$4:$K$9"
            }
        ],
        "format": {
            "x_scale": 1.3,
            "x_offset": 10,
            "y_offset": 20
        },
        "title": {
            "name": "总分柱状图"
        }
    }`); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to add chart due to: %v\n", err)
		os.Exit(3)
	}

	// "legend": {
	//     "none": true
	// },

	if err = f.AddPicture(sheetName, "D6", "images/stamp.png", `{
        "x_offset": 15,
        "y_offset": 15,
        "x_scale": 0.3,
        "y_scale": 0.3
    }`); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to add picture due to: %v\n", err)
		os.Exit(3)
	}

	if err = f.SetSheetViewOptions(sheetName, -1, excelize.ShowGridLines(false)); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set sheet view options due to: %v\n", err)
		os.Exit(3)
	}

	if err = f.SetPanes(sheetName, `{
        "freeze": true,
        "split": false,
        "x_split": 0,
        "y_split": 3,
        "top_left_cell": "A4",
        "active_pane": "bottomLeft"
    }`); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to set sheet view options due to: %v\n", err)
		os.Exit(3)
	}

	if err := f.SaveAs("grade.xlsx"); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to save grade.xlsx due to: %v\n", err)
		os.Exit(4)
	}
}
