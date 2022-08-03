package app

func comparePostProcessed(mergedArr, postArr []item) []keypressGap {
	mergedKeypresses := getKeypresses(mergedArr)
	postKeypresses := getKeypresses(postArr)
	keypressGaps := make([]keypressGap, 0)
	mergedIndex := 0

	for _, keypress := range postKeypresses {

		mergedStart := mergedIndex
		for keypress.Content != mergedKeypresses[mergedIndex].Content {
			mergedIndex++
		}
		gapSize := mergedIndex - mergedStart
		gapKeypresses := make([]item, 0)
		for i := 0; i < gapSize; i++ {
			index := i + mergedStart
			gapKeypresses = append(gapKeypresses, mergedKeypresses[index])
		}
		gap := keypressGap{gapKeypresses}

		keypressGaps = append(keypressGaps, gap)
		mergedIndex++
	}
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
