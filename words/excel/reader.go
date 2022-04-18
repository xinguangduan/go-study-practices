package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"golang-study-practices/words/db"
	"golang-study-practices/words/db/vo"
	"strconv"
)

func ImportData() {
	excelFileName := "/Desktop/COCA20000增强版.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Errorf("file not found!")
	}
	var index int
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var cells []string
			for _, cell := range row.Cells {
				text := cell.String()
				cells = append(cells, text)
			}
			// 获取标签页(时间)
			if cells[0] == "单词列表" {
				continue
			}
			index += 1
			var freq, _ = strconv.Atoi(cells[3])
			word := vo.EnglishWords{
				Id:         index,
				WordName:   cells[0],
				SoundMark:  cells[1],
				Paraphrase: cells[2],
				Frequency:  freq,
			}

			fmt.Printf("%#v", word)
			db.StructBatchInsert(word)

			/*for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s  ", text)
			}*/
		}
	}
}
