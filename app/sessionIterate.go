package app

func SessionIterate(arr []SessionInfo, boolean bool) error {
	for _, session := range arr {
		mergedArr, curatedPostArr, err := buildArrays(session.MergedTsvPath,
			session.CuratedPostProcessedPath)
		if err != nil {
			return err
		}
		generateSessionReport(compareArr(mergedArr, curatedPostArr))
		// if boolean {
		// 	generatePostProcessFixed(mergedArr, curatedPostArr)
		// }
	}
	return nil
}
