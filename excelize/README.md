# Introduction

Practical [Execlize][1] examples.

## Excel basic

Concepts:

- workbook
- sheet
- cell (coordinate, column represented by letter, row represented by number, eg: E2)

common file extension names:
- xlsx: ordinary workbook
- xlsm: workbook with macro
- xltx: workbook template
- xltm: workbook template with macro

common excel softwares:
- Microsoft Office Excel
- Apple Numbers
- WPS Office
- Apache Open Office
- LibreOffice
- Google Sheets

open source excel libraies for major languages:
- go: excelize, unidoc/unioffice, xlsx
- c: libxlsxwriter
- c++: OpenXLSX
- Java: Apache POI XSSF, easyexcel
- Rust: calamine
- PHP: PhpSpreadsheet
- JavaScript: SheetJS/js-xlsx
- Python: xlsxWriter OpenPyXL
- .NET: NPOI

## excelize features

- read/write
- save/save as
- import/export
- ordinary/stream API
- 图表工作表 (AddChartSheet)
- 工作表背景图片 (SetSheetBackground)
- 条件式样 (NewConditionalStyle)
- 备注 (AddComment)
- 数据验证 (AddDataValidation)
- 表头锁定 (SetPanes)
- 隐藏网格线 (ShowGridLines(false))
