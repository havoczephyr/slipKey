package app

import "fmt"

func compareArr(mergedArr, postArr []item) ([]keypressGap, []item, error) {
	mergedKeypresses := filterKeypresses(mergedArr)
	postKeypresses := filterKeypresses(postArr)

	postDict, err := buildPostMap(postKeypresses)
	if err != nil {
		return nil, nil, err
	}

	var prefixKeypress *item = nil
	keypressGaps := make([]keypressGap, 0)
	i := 0
	for i < len(mergedKeypresses) {
		keypress := mergedKeypresses[i]

		_, found := postDict[keypress.TBegin]
		if !found {
			missingKeypresses := make([]item, 0)
			for !found {
				missingKeypresses = append(missingKeypresses, keypress)
				i++
				keypress = mergedKeypresses[i]
				_, found = postDict[keypress.TBegin]
			}
			suffixKeypress := &keypress
			var triggerTime float64
			if i == 0 {
				triggerTime = missingKeypresses[0].TBegin
			} else {
				triggerTime = missingKeypresses[0].TBegin - prefixKeypress.TBegin
			}
			duration := missingKeypresses[len(missingKeypresses)-1].TBegin - missingKeypresses[0].TBegin
			gap := keypressGap{
				PrefixKeypress:    prefixKeypress,
				SuffixKeypress:    suffixKeypress,
				MissingKeypresses: missingKeypresses,
				TriggerTime:       triggerTime,
				Duration:          duration,
			}
			keypressGaps = append(keypressGaps, gap)
		}
		prefixKeypress = &keypress
		i++
	}
	return keypressGaps, mergedKeypresses, nil
}
func buildPostMap(postKeypresses []item) (map[float64]item, error) {
	var dict map[float64]item = make(map[float64]item)
	for _, keypress := range postKeypresses {
		dict[keypress.TBegin] = keypress
	}
	if len(dict) != len(postKeypresses) {
		err := fmt.Errorf("buildPostMap failed - could not build dict")
		return dict, err
	}
	return dict, nil
}

func filterKeypresses(input []item) []item {
	outputArray := make([]item, 0)
	for _, i := range input {
		if i.Tier == "Keypress" {
			outputArray = append(outputArray, i)
		}
	}
	return outputArray
}
