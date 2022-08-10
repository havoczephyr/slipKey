package app

import "fmt"

//SessionIterate cycles through every session in []SessionInfo and initializes array construction through buildArrays(),
// sends the constructed arrays through compareArr, sends the values to generateSessionReport and if enabled,
// resolve the issue in generatePostProcessFixed. returns an error, if any.
func SessionIterate(arr []SessionInfo, boolean bool, folderStr string) error {
	for _, session := range arr {
		mergedArr, curatedPostArr, err := buildArrays(session.MergedTsvPath,
			session.CuratedPostProcessedPath)
		if err != nil {
			fmt.Printf("buildArrays failed on %s - %v", session.SessionName, err)
			continue
		}
		gaps, mergedKeypresses, err := compareArr(mergedArr, curatedPostArr)
		if err != nil {
			fmt.Printf("compareArr failed on %s - %v", session.SessionName, err)
			continue
		}
		generateSessionReport(gaps, session.SessionName, session.SessionPath, boolean, folderStr)
		if boolean {
			generatePostProcessFixed(mergedKeypresses, curatedPostArr, session.SessionPath)
		}
	}
	return nil
}
