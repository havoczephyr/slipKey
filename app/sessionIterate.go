package app

import "fmt"

func SessionIterate(arr []SessionInfo, boolean bool, folderStr string) error {
	for _, session := range arr {
		mergedArr, curatedPostArr, err1 := buildArrays(session.MergedTsvPath,
			session.CuratedPostProcessedPath)
		if err1 != nil {
			fmt.Printf("buildArrays failed on %s - %v", session.SessionName, err1)
		}
		gaps, mergedKeypresses, err2 := compareArr(mergedArr, curatedPostArr)
		if err2 != nil {
			fmt.Printf("compareArr failed on %s - %v", session.SessionName, err2)
		}
		generateSessionReport(gaps, session.SessionName, session.SessionPath, boolean, folderStr)
		if boolean {
			generatePostProcessFixed(mergedKeypresses, curatedPostArr, session.SessionPath)
		}
	}
	return nil
}
