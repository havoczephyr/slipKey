package app

func SessionIterate(arr []SessionInfo) {
	for _, session := range arr {
		BuildArrays(session.MergedTsvPath, session.CuratedPostProcessedPath)
	}
}
