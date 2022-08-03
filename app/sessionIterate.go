package app

func SessionIterate(arr []SessionInfo, boolean bool) error {
	for _, session := range arr {
		mergedArr, curatedPostArr, err := buildArrays(session.MergedTsvPath,
			session.CuratedPostProcessedPath)
		if err != nil {
			return err
			break
		}
		generateSessionReport(comparePostProcessed(mergedArr, curatedPostArr))
		if boolean == true {
			generatePostProcessFixed(mergedArr, curatedPostArr)
		}
	}
	return nil
}
