package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"skingenius/config"
	"skingenius/database"
	dbmodel "skingenius/database/model"
	"skingenius/pubchem/model"
	"strings"
)

const pubchemBaseUrl = "https://pubchem.ncbi.nlm.nih.gov/rest/pug_view/data/compound/%s/json"
const dirname = "pubchem/ingredients/"

func main() {

	db, err := database.NewGormClient(config.RemoteHost, config.Port, config.User, config.Password)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to establish db connection, error: %v", err))
		os.Exit(1)
	}

	//downloadAndSaveIngredients()
	parseIngredientsAndSoreToDb(db)
}

func parseIngredientsAndSoreToDb(db database.Connector) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(dirname + file.Name())
			if err != nil {
				log.Fatal(err)
			}

			respById := model.GetRecordByIdResp{}
			if unmarshallErr := json.Unmarshal(content, &respById); unmarshallErr != nil {
				fmt.Println(fmt.Sprintf("failed to unmarshall pubchem record from file %s: %v", file.Name(), unmarshallErr))
			}

			var i dbmodel.Ingredient

			i.Name = respById.Record.RecordTitle
			i.PubchemId = respById.Record.RecordNumber

			fmt.Println(fmt.Sprintf("Name: %#v", respById.Record.RecordTitle))
			fmt.Println(fmt.Sprintf("CID: %#v", respById.Record.RecordNumber))

			for _, section := range respById.Record.Section {
				for _, ss := range section.Section {
					for _, sss := range ss.Section {
						if strings.Contains(sss.TOCHeading, "CAS") {
							fmt.Println(sss.Information[0])
							i.CasNumber = sss.Information[0].Value.StringWithMarkup[0].String
						}
						if strings.Contains(sss.TOCHeading, "European Community") {
							fmt.Println(sss.Information[0])
							i.ECNumber = sss.Information[0].Value.StringWithMarkup[0].String
						}
						if strings.Contains(sss.TOCHeading, "Synonyms") {
							fmt.Println(sss.Information[0])
							for _, sn := range sss.Information[0].Value.StringWithMarkup {
								i.Synonyms = append(i.Synonyms, sn.String)
							}
						}
					}
				}
			}

			saveErr := db.SaveIngredient(&i)
			if saveErr != nil {
				fmt.Println(fmt.Sprintf("failed ot save ingredient [%s], error: %v", file.Name(), saveErr))
			}

			fmt.Println(fmt.Sprintf("%#v \n\n", i))
		}
	}
}

func downloadAndSaveIngredients() {

	ingredientsList := []string{"962", "753", "23343930", "9204", "91745", "443277", "8221", "5280343", "71386", "104935", "57417357", "14985", "31226", "5366837", "5280460", "519047", "615735", "3082978", "7417", "311", "18650", "10657", "6852218", "6049", "61811", "10230", "5280489", "440917", "8207", "6549", "176871", "5280484", "5350520", "5357148", "5462308", "8221", "9859093", "753", "9204", "11119", "31226", "5357148", "6322", "86472", "108007", "962", "11953916", "23343930", "91770", "985", "985", "5357148", "5280852", "72736", "5462308", "10467", "17817338", "5280852", "446284", "14870", "753", "85857", "9204", "5280934", "18561718", "23345", "5318514", "521889", "442431", "14985", "962", "11953916", "5357148", "5460178", "31249", "91745", "6436662", "6322", "5357148", "118678", "72736", "5462308", "8221", "11979287", "11979287", "3082978", "753", "85857", "9204", "5280934", "643682", "1548969", "6850742", "9895643", "516036", "9927494", "94510", "5360323", "11180", "5281", "8201752", "962", "11953916", "72214142", "5357148", "9568835", "9898300", "520985", "None", "5360523", "24822455"}

	//var pubchemIngredients []model.PubChemRecord

	//var jsonIngredients []string

	var skippedIngredients []string

	for _, i := range ingredientsList {
		fmt.Println(fmt.Sprintf("ID: %#v", i))

		resp, err := http.Get(fmt.Sprintf(pubchemBaseUrl, i))
		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to obtain record by CID: %s", i))
		}

		//jsonIngredients = append(jsonIngredients, string(body))

		respById := model.GetRecordByIdResp{}
		if unmarshallErr := json.Unmarshal(body, &respById); unmarshallErr != nil {
			fmt.Println(fmt.Sprintf("failed to unmarshall pubchem record: %v", unmarshallErr))
		}

		if respById.Record.RecordTitle == "" {
			skippedIngredients = append(skippedIngredients, i)
			continue
		}

		//fmt.Println(fmt.Sprintf("Record: %#v", respById.Record))
		fmt.Println(fmt.Sprintf("Name: %#v", respById.Record.RecordTitle))
		fmt.Println(fmt.Sprintf("CID: %#v", respById.Record.RecordNumber))
		saveToFile(respById.Record.RecordTitle, respById.Record.RecordNumber, string(body))
		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("Skipped ingredients: %v", skippedIngredients))

	//fmt.Println(fmt.Sprintf("Ingredients total: %d, obtained: %d", len(ingredientsList), len(jsonIngredients)))
}

func saveToFile(name string, CID int, jsonIngredient string) {

	filename := fmt.Sprintf("%s%d_%s.txt", dirname, CID, strings.ReplaceAll(name, " ", "-"))
	if fileExists(filename) {
		fmt.Println(fmt.Sprintf("file exists: %s", filename))
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(jsonIngredient)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if info == nil {
		fmt.Println("NIL: ", filename)
		return false
	}

	return !info.IsDir()
}
