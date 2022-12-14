package app

import (
	"fmt"
	"os"
)

//generatePostProcessFixed will take the "Keypress" values from merged.tsv and properly replace the values from curated-postprocessed.tsv.
// and export them into a new tsv file, postprocessed-fixed.tsv.
func generatePostProcessFixed(mergedKeypressArr, postArr []item, dir string) {
	var fixedArr []item

	for _, item := range postArr {
		if item.Tier == "Keypress" {
			continue
		} else {
			fixedArr = append(fixedArr, item)
		}
	}
	fixedArr = append(fixedArr, mergedKeypressArr...)

	postProcessedFixedTsv, err := os.Create(fmt.Sprintf("%s/postprocessed-fixed.tsv", dir))
	if err != nil {
		fmt.Printf("failed to create postprocessed-fixed.tsv -- %v", err)
	}
	defer postProcessedFixedTsv.Close()

	for _, item := range fixedArr {
		postProcessedFixedTsv.WriteString(fmt.Sprintf("%.3f\t%.3f\t%s\t%s\n",
			item.TBegin, item.TEnd, item.Tier, item.Content))
	}
}
