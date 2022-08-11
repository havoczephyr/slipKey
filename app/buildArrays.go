package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//buildArrays will use the directory paths of merged.tsv and curated-postprocessed.tsv and returns []item's of both and an error, if any.
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

//readItems() will read and parse an entire .tsv file and return an []item, with an error, if any.
func readItems(path string) ([]item, error) {
	file, err := os.Open(path)
	if err != nil {
		readError := fmt.Errorf("error reading file %s: %w", path, err)
		return nil, readError
	}
	defer file.Close()

	// reader := csv.NewReader(file)
	// reader.Comma = '\t'
	// reader.FieldsPerRecord = -1

	// tsvData, err := reader.ReadAll()
	// if err != nil {
	// 	return nil, err
	// }
	//csv.NewReader had issues with "" because it was presuming normal tsv conventions.

	itemsArray := make([]item, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan() //<-- skips the first line tBegin, tEnd etc.
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "\t")
		tBegin, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return nil, err
		}
		tEnd, err := strconv.ParseFloat(values[1], 64)
		if err != nil {
			return nil, err
		}
		tier := values[2]
		content := values[3]
		lineItem := item{tBegin, tEnd, tier, content}
		itemsArray = append(itemsArray, lineItem)
	}

	return itemsArray, nil
}
