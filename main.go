package main

import (
	"fmt"
	"os"
	"slipKey/app"
)

func main() {
	// var returnError error
	fixBool, wd, err1 := app.Boot()
	if err1 != nil {
		fmt.Printf("failed to boot -%v", err1)
		os.Exit(1)
	}
	sessionInfos := app.ProcessSessions(wd)
	err2 := app.SessionIterate(sessionInfos, fixBool)
	if err2 != nil {
		fmt.Printf("failed to iterate -%v", err2)
		os.Exit(1)
	}

}
