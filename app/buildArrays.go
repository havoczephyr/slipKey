package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	itemsArray := make([]item, 0)

	for scanner.Scan() {
		lineData := strings.Split(scanner.Text(), "\t")
		tBegin, err := strconv.ParseFloat(lineData[0], 64)
		if err != nil {
			continue
		}
		tEnd, err := strconv.ParseFloat(lineData[1], 64)
		if err != nil {
			continue
		}
		tier := lineData[2]
		content := lineData[3]

		lineItem := item{tBegin, tEnd, tier, content}
		itemsArray = append(itemsArray, lineItem)
	}
	return itemsArray, nil
}
