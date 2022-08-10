package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//Cycles through every folder in the given directory. If they are folders with the prefix of -session,
// and they contain both merged and curated-postprocessed tsv files, they will be collected and returned
// as a []SessionInfo.
func ProcessSessions(dir string) []SessionInfo {
	folders, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("failed to read directory -%v", err)
	}
	sessionInfos := make([]SessionInfo, 0)
	for _, folder := range folders {
		isSession := folder.IsDir() && strings.HasPrefix(folder.Name(), "session-")
		if isSession {
			sessionPath := filepath.Join(dir, folder.Name())
			mergedPath, curatedPostPath, err := vitalCheck(sessionPath)
			if err != nil {
				fmt.Printf("vital check failed -%v", err)
			} else {
				session := SessionInfo{
					SessionName:              folder.Name(),
					SessionPath:              sessionPath,
					MergedTsvPath:            mergedPath,
					CuratedPostProcessedPath: curatedPostPath,
				}
				sessionInfos = append(sessionInfos, session)
			}
		}
	}
	return sessionInfos
}

//vitalCheck() is used to verify if merged.tsv and curated-postprocessed.tsv are available inside of an iterated session.
// Will return the directory address to both and an error, if any.
func vitalCheck(dir string) (string, string, error) {
	const MERGED_NAME string = "merged.tsv"
	const CUR_POST_NAME string = "curated-processed.tsv"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", "", err
	}

	merged := ""
	cur_post := ""
	for _, file := range files {
		if file.Name() == MERGED_NAME {
			merged = filepath.Join(dir, MERGED_NAME)
			if cur_post != "" {
				return merged, cur_post, nil
			}
		} else if file.Name() == CUR_POST_NAME {
			cur_post = filepath.Join(dir, CUR_POST_NAME)
			if merged != "" {
				return merged, cur_post, nil
			}
		}
	}
	var empty_err error
	if merged == "" && cur_post == "" {
		empty_err = fmt.Errorf("could not find files: %s and %s", MERGED_NAME, CUR_POST_NAME)
	} else if merged == "" {
		empty_err = fmt.Errorf("could not find file: %s", MERGED_NAME)
	} else {
		empty_err = fmt.Errorf("could not find file: %s", CUR_POST_NAME)
	}

	return merged, cur_post, empty_err
}
