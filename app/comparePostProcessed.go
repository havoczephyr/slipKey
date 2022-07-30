package app

func ComparePostProcessed(mergedArr, postArr []Item) []Item {
	mergedKeypresses := getKeypresses(mergedArr)
	postKeypresses := getKeypresses(postArr)
	keypressGaps := make([]KeypressGap, 0)
	mergedIndex := 0

	for _, keypress := range postKeypresses {
		//while keypress.Content of mergedKeypress != merged[mergedIndex].Content
		//mergedIndex ++
		mergedStart := mergedIndex
		for keypress.Content != mergedKeypresses[mergedIndex].Content {
			mergedIndex++
		}
		// delta := mergedKeypresses[mergedIndex].TEnd - mergedKeypresses[mergedStart].TBegin
		gapSize := mergedIndex - mergedStart
		gapKeypresses := make([]Item, 0)
		for i := 0; i < gapSize; i++ {
			index := i + mergedStart
			gapKeypresses = append(gapKeypresses, mergedKeypresses[index])
		}
		gap := KeypressGap{}
		gap.Size()
		keypressGaps = append(keypressGaps, gap)
		mergedIndex++
	}

}

func getKeypresses(input []Item) []Item {
	outputArray := make([]Item, 0)
	for _, i := range input {
		if i.Tier == "Keypress" {
			outputArray = append(outputArray, i)
		}
	}
	return outputArray
}
