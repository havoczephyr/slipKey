package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slipKey/app"
	"time"
)

func main() {
	// var returnError error
	fixBool, wd, err1 := app.Boot()
	if err1 != nil {
		fmt.Printf("failed to boot -%v", err1)
		os.Exit(1)
	}
	folderName := fmt.Sprintf("SlipKey-Reports%d%d%d",
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Year())

	reportsFolder := filepath.Join(".", folderName)
	err3 := os.MkdirAll(reportsFolder, os.ModeAppend)
	if err3 != nil {
		fmt.Printf("failed to generate Reports folder -%v", err3)
		os.Exit(1)
	}
	sessionInfos := app.ProcessSessions(wd)
	err2 := app.SessionIterate(sessionInfos, fixBool, folderName)
	if err2 != nil {
		fmt.Printf("failed to iterate -%v", err2)
		os.Exit(1)
	}

}
