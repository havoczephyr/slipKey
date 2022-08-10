package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slipKey/app"
	"time"
)

func main() {
	fixBool, wd, err1 := app.Boot()
	if err1 != nil {
		fmt.Printf("failed to boot -%v", err1)
		return
	}
	folderName := fmt.Sprintf("SlipKey-Reports%d%d%d",
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Year())

	reportsFolder := filepath.Join(".", folderName)
	err3 := os.MkdirAll(reportsFolder, os.ModeAppend)
	if err3 != nil {
		fmt.Printf("failed to generate Reports folder -%v", err3)
		return
	}
	sessionInfos := app.ProcessSessions(wd)
	err2 := app.SessionIterate(sessionInfos, fixBool, folderName)
	if err2 != nil {
		fmt.Printf("failed to iterate -%v", err2)
		return
	}

}
