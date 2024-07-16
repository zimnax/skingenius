package xlsx

import (
	"fmt"
	"github.com/thedatashed/xlsxreader"
	"os"
)

func ReadAnswers(path string) []string {
	var answers []string

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	e, err := xlsxreader.OpenFile(path)
	if err != nil {
		fmt.Printf("error: %s \n", err)
		return []string{}
	}
	defer e.Close()

	//fmt.Printf("Worksheets: %s \n", e.Sheets)

	for row := range e.ReadRows(e.Sheets[0]) {
		if row.Error != nil {
			fmt.Printf("error on row %d: %s \n", row.Index, row.Error)
			return []string{}
		}

		if row.Index < 10 {
			//fmt.Printf("-> %+v \n", row.Cells)

			for _, cell := range row.Cells {
				if cell.Column == "C" && cell.Row > 1 {
					answers = append(answers, cell.Value)
				}
			}
		}
	}

	return answers
}
