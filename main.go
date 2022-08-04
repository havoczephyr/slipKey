package main

import (
	"fmt"
	"slipKey/app"
)

func main() {
	var returnError error
	fixBool, wd, err1 := app.Boot()
	if err1 != nil {
		returnError = fmt.Errorf("failed to boot -%w", err1)
		panic(returnError)
	}
	sessionInfos := app.ProcessSessions(wd)
	err2 := app.SessionIterate(sessionInfos, fixBool)
	if err2 != nil {
		returnError = fmt.Errorf("failed to iterate -%w", err2)
		panic(returnError)
	}

}
