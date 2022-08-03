package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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
				session := SessionInfo{folder.Name(), mergedPath, curatedPostPath}
				sessionInfos = append(sessionInfos, session)
			}
		}
	}
	return sessionInfos
}
