package app

import "fmt"

func compareArr(mergedArr, postArr []item) []keypressGap {
	mergedKeypresses := getKeypresses(mergedArr)
	postKeypresses := getKeypresses(postArr)
	keypressGaps := make([]keypressGap, 0)
	mergedIndex := 0

	for val, keypress := range postKeypresses {

		mergedStart := mergedIndex
		if keypress.Content != mergedKeypresses[mergedIndex].Content {
			fmt.Printf("compareArr Debug: keypress comparator %s, %s", keypress.Content, mergedKeypresses[mergedIndex].Content)
			mergedIndex++
		}
		gapSize := mergedIndex - mergedStart
		fmt.Printf("compareArr Debug: gapsize %d", gapSize)
		gapKeypresses := make([]item, 0)
		for i := 0; i < gapSize; i++ {
			index := i + mergedStart
			gapKeypresses = append(gapKeypresses, mergedKeypresses[index])
		}
		triggerTime := gapKeypresses[0].TBegin - mergedKeypresses[val-1].TBegin
		duration := gapKeypresses[len(gapKeypresses)-1].TEnd - gapKeypresses[0].TBegin
		gap := keypressGap{gapKeypresses, triggerTime, duration}
		fmt.Printf("compareArr Debug: triggerTime %f, duration %f", triggerTime, duration)

		keypressGaps = append(keypressGaps, gap)
		mergedIndex++
	}
	return keypressGaps
}

func getKeypresses(input []item) []item {
	outputArray := make([]item, 0)
	for _, i := range input {
		if i.Tier == "Keypress" {
			outputArray = append(outputArray, i)
		}
	}
	return outputArray
}
