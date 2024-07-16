package csv

import (
	"encoding/csv"
	"os"
	"strconv"
)

func WriteToFile(filename string, headers []string, data map[string]int) error {
	file2, err := os.Create(filename + ".csv")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	defer writer.Flush()

	writer.Write(headers)
	for pName, pScore := range data {
		writer.Write([]string{pName, strconv.Itoa(pScore)})
	}

	return nil
}

func SingleProductExtendedReport(filename string, headers []string, productName string, data map[string]int) error {
	file, err := os.OpenFile(filename+".csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)
	for ingredient, score := range data {
		writer.Write([]string{productName, ingredient, strconv.Itoa(score)})
	}

	writer.Write([]string{"", "", ""})
	writer.Write([]string{"", "", ""})

	return nil
}

func WriteToCsv(filename string, headers []string, data [][]string) error {
	file, err := os.OpenFile(filename+".csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)
	for _, scores := range data {
		writer.Write(scores)
	}

	writer.Write([]string{"", "", ""})
	writer.Write([]string{"", "", ""})

	return nil
}
