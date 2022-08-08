package app

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func buildArrays(mergedPath string, postPath string) ([]item, []item, error) {
	MergedItems, err := readItems(mergedPath)
	if err != nil {
		return nil, nil, err
	}
	curatedPostProcessedItems, err := readItems(postPath)
	if err != nil {
		return nil, nil, err
	}
	return MergedItems, curatedPostProcessedItems, nil
}

func readItems(path string) ([]item, error) {
	file, err := os.Open(path)
	if err != nil {
		readError := fmt.Errorf("error reading file %s: %w", path, err)
		return nil, readError
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	itemsArray := make([]item, 0)

	tsvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, column := range tsvData {
		if column[0] == "tBegin" {
			continue
		} else {
			tBegin, err := strconv.ParseFloat(column[0], 64)
			if err != nil {
				fmt.Printf("parse failure -%v", err)
				os.Exit(1)
			}
			tEnd, err := strconv.ParseFloat(column[1], 64)
			if err != nil {
				fmt.Printf("parse failure -%v", err)
				os.Exit(1)
			}
			tier := column[2]
			content := column[3]
			lineItem := item{tBegin, tEnd, tier, content}
			itemsArray = append(itemsArray, lineItem)
		}

	}

	return itemsArray, nil
}
